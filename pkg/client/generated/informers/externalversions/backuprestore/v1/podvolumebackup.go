// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	backuprestorev1 "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	versioned "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/clientset/versioned"
	internalinterfaces "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/listers/backuprestore/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodVolumeBackupInformer provides access to a shared informer and lister for
// PodVolumeBackups.
type PodVolumeBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PodVolumeBackupLister
}

type podVolumeBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodVolumeBackupInformer constructs a new informer for PodVolumeBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodVolumeBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodVolumeBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodVolumeBackupInformer constructs a new informer for PodVolumeBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodVolumeBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackuprestoreV1().PodVolumeBackups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackuprestoreV1().PodVolumeBackups(namespace).Watch(context.TODO(), options)
			},
		},
		&backuprestorev1.PodVolumeBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *podVolumeBackupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodVolumeBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podVolumeBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&backuprestorev1.PodVolumeBackup{}, f.defaultInformer)
}

func (f *podVolumeBackupInformer) Lister() v1.PodVolumeBackupLister {
	return v1.NewPodVolumeBackupLister(f.Informer().GetIndexer())
}
