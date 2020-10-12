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

// HookInformer provides access to a shared informer and lister for
// Hooks.
type HookInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.HookLister
}

type hookInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHookInformer constructs a new informer for Hook type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHookInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHookInformer constructs a new informer for Hook type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackuprestoreV1().Hooks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackuprestoreV1().Hooks(namespace).Watch(context.TODO(), options)
			},
		},
		&backuprestorev1.Hook{},
		resyncPeriod,
		indexers,
	)
}

func (f *hookInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHookInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hookInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&backuprestorev1.Hook{}, f.defaultInformer)
}

func (f *hookInformer) Lister() v1.HookLister {
	return v1.NewHookLister(f.Informer().GetIndexer())
}
