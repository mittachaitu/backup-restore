// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	v1 "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodVolumeBackupLister helps list PodVolumeBackups.
// All objects returned here must be treated as read-only.
type PodVolumeBackupLister interface {
	// List lists all PodVolumeBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodVolumeBackup, err error)
	// PodVolumeBackups returns an object that can list and get PodVolumeBackups.
	PodVolumeBackups(namespace string) PodVolumeBackupNamespaceLister
	PodVolumeBackupListerExpansion
}

// podVolumeBackupLister implements the PodVolumeBackupLister interface.
type podVolumeBackupLister struct {
	indexer cache.Indexer
}

// NewPodVolumeBackupLister returns a new PodVolumeBackupLister.
func NewPodVolumeBackupLister(indexer cache.Indexer) PodVolumeBackupLister {
	return &podVolumeBackupLister{indexer: indexer}
}

// List lists all PodVolumeBackups in the indexer.
func (s *podVolumeBackupLister) List(selector labels.Selector) (ret []*v1.PodVolumeBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodVolumeBackup))
	})
	return ret, err
}

// PodVolumeBackups returns an object that can list and get PodVolumeBackups.
func (s *podVolumeBackupLister) PodVolumeBackups(namespace string) PodVolumeBackupNamespaceLister {
	return podVolumeBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodVolumeBackupNamespaceLister helps list and get PodVolumeBackups.
// All objects returned here must be treated as read-only.
type PodVolumeBackupNamespaceLister interface {
	// List lists all PodVolumeBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodVolumeBackup, err error)
	// Get retrieves the PodVolumeBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PodVolumeBackup, error)
	PodVolumeBackupNamespaceListerExpansion
}

// podVolumeBackupNamespaceLister implements the PodVolumeBackupNamespaceLister
// interface.
type podVolumeBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodVolumeBackups in the indexer for a given namespace.
func (s podVolumeBackupNamespaceLister) List(selector labels.Selector) (ret []*v1.PodVolumeBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodVolumeBackup))
	})
	return ret, err
}

// Get retrieves the PodVolumeBackup from the indexer for a given namespace and name.
func (s podVolumeBackupNamespaceLister) Get(name string) (*v1.PodVolumeBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("podvolumebackup"), name)
	}
	return obj.(*v1.PodVolumeBackup), nil
}