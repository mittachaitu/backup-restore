package cmd

import (
	"context"
	"os"
	"time"

	"github.com/mayadata.io/kubera-backup-restore/pkg/client"
	"github.com/mayadata.io/kubera-backup-restore/pkg/signals"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	kuberaapis "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	clientset "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/clientset/versioned"
	informers "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/informers/externalversions"
	controllers "github.com/mayadata.io/kubera-backup-restore/pkg/controllers"
	kuberadiscovery "github.com/mayadata.io/kubera-backup-restore/pkg/discovery"
	kbclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// defaultNonAdminMode sets to false to state server will
	// will running in admin mode
	defaultNonAdminMode = false
)

type serverConfig struct {
	isNonAdminMode bool
}

type controllerRunInfo struct {
	controller   controllers.Interface
	numOfWorkers int
}

func NewCommand(factory client.Factory) *cobra.Command {
	var (
		config = serverConfig{
			isNonAdminMode: defaultNonAdminMode,
		}
	)
	command := &cobra.Command{
		Use:   "kubera-protect agent",
		Short: "Run the kubera-protect agent",
		Long:  "Run the kubera-protect agent",
		Run: func(c *cobra.Command, args []string) {

			// we log to stdout
			logrus.SetOutput(os.Stdout)
			logger := logrus.New()
			logger.Level = logrus.InfoLevel
			logger.Out = os.Stdout

			logger.Infof("Kubera-Protect started on Non Admin Mode(%t)", config.isNonAdminMode)
			s, err := newServer(factory, config, logger)
			CheckError(err)

			CheckError(s.run())
		},
	}
	command.Flags().BoolVar(&config.isNonAdminMode, "non-admin-mode", config.isNonAdminMode, "Non Admin mode runs kubera-protect to run in non admin mode with limited permissions")

	return command
}

type server struct {
	namespace             string
	kubeClientConfig      *rest.Config
	kubeClient            kubernetes.Interface
	kuberaClient          clientset.Interface
	dynamicClient         dynamic.Interface
	discoveryClient       discovery.DiscoveryInterface
	discoveryHelper       kuberadiscovery.Helper
	kbClient              kbclient.Client
	sharedInformerFactory informers.SharedInformerFactory
	context               context.Context
	cancelFunc            context.CancelFunc
	logger                logrus.FieldLogger
	config                serverConfig
	manager               manager.Manager
}

// newServer will be an newInstance of server
func newServer(f client.Factory, config serverConfig, logger *logrus.Logger) (*server, error) {
	kubeClient, err := f.KubeClient()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get kubernetes client")
	}

	kuberaClient, err := f.Client()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get kubera client")
	}

	dynamicClient, err := f.DynamicClient()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get dynamic client")
	}

	clientConfig, err := f.ClientConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get client configration")
	}

	kbClient, err := f.KubebuilderClient()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get kubebuilder client")
	}

	scheme := runtime.NewScheme()
	err = kuberaapis.AddToScheme(scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add scheme")
	}

	// cancelFunc is not deferred here because if it was, then ctx would immediately
	// be cancelled once this function exited, making it useless to any informers using later.
	// That, in turn, causes the velero server to halt when the first informer tries to use it (probably restic's).
	// Therefore, we must explicitly call it on the error paths in this function.
	ctx, cancelFunc := context.WithCancel(context.Background())

	manager, err := ctrl.NewManager(clientConfig, ctrl.Options{
		Scheme: scheme})
	if err != nil {
		cancelFunc()
		return nil, err
	}

	s := &server{
		namespace:        f.Namespace(),
		kubeClientConfig: clientConfig,
		kubeClient:       kubeClient,
		kuberaClient:     kuberaClient,
		dynamicClient:    dynamicClient,
		discoveryClient:  kuberaClient.Discovery(),
		kbClient:         kbClient,
		// We can constructs a new instance of sharedInformerFactory for all namespaces by using NewSharedInformerFactory
		sharedInformerFactory: informers.NewSharedInformerFactoryWithOptions(kuberaClient, 0, informers.WithNamespace(f.Namespace())),
		logger:                logger,
		manager:               manager,
		config:                config,
		context:               ctx,
		cancelFunc:            cancelFunc,
	}
	return s, nil
}

func (s *server) run() error {
	signals.CancelOnShutdown(s.cancelFunc, s.logger)

	// Verify whether configured namespace exist or not
	if err := s.verifyNamespaceExistence(s.namespace); err != nil {
		return err
	}

	err := s.initDiscoveryHelper()
	if err != nil {
		return err
	}

	if err := s.runControllers(); err != nil {
		return err
	}

	return nil
}

func (s *server) verifyNamespaceExistence(ns string) error {
	s.logger.WithField("namespace", ns).Info("Checking existentce of namespace")
	if _, err := s.kubeClient.CoreV1().Namespaces().Get(s.context, ns, metav1.GetOptions{}); err != nil {
		return errors.Wrapf(err, "failed to verify namespace existence")
	}

	return nil
}

func (s *server) initDiscoveryHelper() error {
	discoveryHelper, err := kuberadiscovery.NewHelper(s.discoveryClient)
	if err != nil {
		return errors.Wrapf(err, "failed to build discovery helper")
	}
	s.discoveryHelper = discoveryHelper

	go wait.Until(
		func() {
			err := discoveryHelper.LoadResources()
			if err != nil {
				s.logger.WithError(err).Error("error reloading discovery resources")
			}
		},
		5*time.Second,
		s.context.Done(),
	)

	return nil
}

func (s *server) runControllers() error {
	var backupResourceEventHandler cache.ResourceEventHandler

	dynamicFactory := client.NewDynamicFactory(s.dynamicClient)
	s.logger.Infof("Starting controllers")

	// informers := k8sinformers.NewSharedInformerFactory(s.kubeClient, 0)
	backupController = controllers.NewBackupController(
		s.sharedInformerFactory.Backuprestore().V1().KuberaBackups(),
		s.kuberaClient,
		s.kubeClient,
		s.discoveryHelper,
		dynamicFactory,
		s.kbClient,
		s.logger,
		backupResourceEventHandler,
		s.config.isNonAdminMode,
	)
	return nil
}

// getBackupInformer return configmap informer or backup informer based
// on flag
// func (s *server) getBackupInformer() genericinformers.GenericInformer {
// 	if s.config.isNonAdminMode {
// 		// return configmap informer
// 	}
// 	// return backup informer
// }
