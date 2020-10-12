package client

import (
	"os"

	kbclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	kuberaapi "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	clientset "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/clientset/versioned"
)

// Factory knows how to create a Kubera Client and Kubernetes client.
type Factory interface {
	// BindFlags binds common flags (--kubeconfig, --namespace) to the passed-in FlagSet.
	BindFlags(flags *pflag.FlagSet)
	// Client returns a Kubera Client. It uses the following priority to specify the cluster
	// configuration: --kubeconfig flag, KUBECONFIG environment variable, in-cluster configuration.
	Client() (clientset.Interface, error)
	// KubeClient returns a Kubernetes client. It uses the following priority to specify the cluster
	// configuration: --kubeconfig flag, KUBECONFIG environment variable, in-cluster configuration.
	KubeClient() (kubernetes.Interface, error)
	// DynamicClient returns a Kubernetes dynamic client. It uses the following priority to specify the cluster
	// configuration: --kubeconfig flag, KUBECONFIG environment variable, in-cluster configuration.
	DynamicClient() (dynamic.Interface, error)
	// KubebuilderClient returns a Kubernetes dynamic client. It uses the following priority to specify the cluster
	// configuration: --kubeconfig flag, KUBECONFIG environment variable, in-cluster configuration.
	KubebuilderClient() (kbclient.Client, error)
	// SetClientQPS sets the Queries Per Second for a client.
	SetClientQPS(float32)
	// SetClientBurst sets the Burst for a client.
	SetClientBurst(int)
	// ClientConfig returns a rest.Config struct used for client-go clients.
	ClientConfig() (*rest.Config, error)
	// Namespace returns the namespace which the Factory will create clients for.
	Namespace() string
}

type factory struct {
	config      *rest.Config
	flags       *pflag.FlagSet
	kubeconfig  string
	kubecontext string
	namespace   string
	clientQPS   float32
	clientBurst int
}

// NewFactory returns a Factory.
func NewFactory(config KuberaConfig) Factory {
	f := &factory{
		flags: pflag.NewFlagSet("", pflag.ContinueOnError),
	}

	f.namespace = os.Getenv("KUBERA_NAMESPACE")
	if config.Namespace() != "" {
		f.namespace = config.Namespace()
	}

	// We didn't get the namespace via env var or config file, so use the default.
	// Command line flags will override when BindFlags is called.
	if f.namespace == "" {
		f.namespace = kuberaapi.KuberaNamespace
	}

	f.flags.StringVar(&f.kubeconfig, "kubeconfig", "", "Path to the kubeconfig file to use to talk to the Kubernetes apiserver. If unset, try the environment variable KUBECONFIG, as well as in-cluster configuration")
	f.flags.StringVarP(&f.namespace, "namespace", "n", f.namespace, "The namespace in which Kubera backup restore should operate")
	f.flags.StringVar(&f.kubecontext, "kubecontext", "", "The context to use to talk to the Kubernetes apiserver. If unset defaults to whatever your current-context is (kubectl config current-context)")

	return f
}

func (f *factory) BindFlags(flags *pflag.FlagSet) {
	flags.AddFlagSet(f.flags)
}

func (f *factory) ClientConfig() (*rest.Config, error) {
	if f.config != nil {
		return f.config, nil
	}
	config, err := Config(f.kubeconfig, f.kubecontext, f.clientQPS, f.clientBurst)
	if err != nil {
		return nil, err
	}
	f.config = config
	return config, nil
}

func (f *factory) Client() (clientset.Interface, error) {
	clientConfig, err := f.ClientConfig()
	if err != nil {
		return nil, err
	}

	kuberaClient, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return kuberaClient, nil
}

func (f *factory) KubeClient() (kubernetes.Interface, error) {
	clientConfig, err := f.ClientConfig()
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return kubeClient, nil
}

func (f *factory) DynamicClient() (dynamic.Interface, error) {
	clientConfig, err := f.ClientConfig()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dynamicClient, nil
}

func (f *factory) KubebuilderClient() (kbclient.Client, error) {
	clientConfig, err := f.ClientConfig()
	if err != nil {
		return nil, err
	}

	scheme := runtime.NewScheme()
	kuberaapi.AddToScheme(scheme)
	kubebuilderClient, err := kbclient.New(clientConfig, kbclient.Options{
		Scheme: scheme,
	})

	return kubebuilderClient, nil
}

func (f *factory) SetClientQPS(qps float32) {
	f.clientQPS = qps
}

func (f *factory) SetClientBurst(burst int) {
	f.clientBurst = burst
}

func (f *factory) Namespace() string {
	return f.namespace
}
