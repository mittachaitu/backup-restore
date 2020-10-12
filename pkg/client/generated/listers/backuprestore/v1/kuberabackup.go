// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KuberaBackupLister helps list KuberaBackups.
// All objects returned here must be treated as read-only.
type KuberaBackupLister interface {
	// List lists all KuberaBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KuberaBackup, err error)
	// KuberaBackups returns an object that can list and get KuberaBackups.
	KuberaBackups(namespace string) KuberaBackupNamespaceLister
	KuberaBackupListerExpansion
}

// kuberaBackupLister implements the KuberaBackupLister interface.
type kuberaBackupLister struct {
	indexer cache.Indexer
}

// NewKuberaBackupLister returns a new KuberaBackupLister.
func NewKuberaBackupLister(indexer cache.Indexer) KuberaBackupLister {
	return &kuberaBackupLister{indexer: indexer}
}

// List lists all KuberaBackups in the indexer.
func (s *kuberaBackupLister) List(selector labels.Selector) (ret []*v1.KuberaBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KuberaBackup))
	})
	return ret, err
}

// KuberaBackups returns an object that can list and get KuberaBackups.
func (s *kuberaBackupLister) KuberaBackups(namespace string) KuberaBackupNamespaceLister {
	return kuberaBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KuberaBackupNamespaceLister helps list and get KuberaBackups.
// All objects returned here must be treated as read-only.
type KuberaBackupNamespaceLister interface {
	// List lists all KuberaBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KuberaBackup, err error)
	// Get retrieves the KuberaBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.KuberaBackup, error)
	KuberaBackupNamespaceListerExpansion
}

// kuberaBackupNamespaceLister implements the KuberaBackupNamespaceLister
// interface.
type kuberaBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KuberaBackups in the indexer for a given namespace.
func (s kuberaBackupNamespaceLister) List(selector labels.Selector) (ret []*v1.KuberaBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KuberaBackup))
	})
	return ret, err
}

// Get retrieves the KuberaBackup from the indexer for a given namespace and name.
func (s kuberaBackupNamespaceLister) Get(name string) (*v1.KuberaBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kuberabackup"), name)
	}
	return obj.(*v1.KuberaBackup), nil
}
