package controllers

import (
	kuberaapis "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	"github.com/mayadata.io/kubera-backup-restore/pkg/client"
	clientset "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/clientset/versioned"

	informers "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/informers/externalversions/backuprestore/v1"
	listers "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/listers/backuprestore/v1"
	kuberadiscovery "github.com/mayadata.io/kubera-backup-restore/pkg/discovery"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	// "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	kbclient "sigs.k8s.io/controller-runtime/pkg/client"
	// "k8s.io/client-go/tools/cache"
)

type backupController struct {
	*genericController
	// objectGetter    objectGetterInterface
	discoveryHelper kuberadiscovery.Helper
	dynamicFactory  client.DynamicFactory
	lister          listers.KuberaBackupLister
	kubeClient      kubernetes.Interface
	kuberaClient    clientset.Interface
	kbclient        kbclient.Client
	isNonAdminMode  bool
}

// NewBackupController returns new instance of backupcontroller
func NewBackupController(
	informer informers.KuberaBackupInformer,
	kuberaClient clientset.Interface,
	kubeClient kubernetes.Interface,
	discoveryHelper kuberadiscovery.Helper,
	dynamicFactory client.DynamicFactory,
	kbclient kbclient.Client,
	logger logrus.FieldLogger,
	resourceEventHandler cache.ResourceEventHandler,
	isNonAdminMode bool,
) Interface {
	c := &backupController{
		genericController: newGenericController("backup", logger),
		discoveryHelper:   discoveryHelper,
		dynamicFactory:    dynamicFactory,
		lister:            informer.Lister(),
		kubeClient:        kubeClient,
		kuberaClient:      kuberaClient,
		kbclient:          kbclient,
	}
	c.syncHandler = c.sync
	informer.Informer().AddEventHandler(resourceEventHandler)

	return c
}

func (c *backupController) sync(obj interface{}) error {
	backupObj, ok := obj.(*kuberaapis.KuberaBackup)
	if !ok {
		return errors.Errorf("failed to assert into backup object")
	}

	c.logger.Infof("Backup object %+v", backupObj)
	return nil
}
