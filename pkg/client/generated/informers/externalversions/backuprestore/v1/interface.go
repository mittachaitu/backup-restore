// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Hooks returns a HookInformer.
	Hooks() HookInformer
	// KuberaBackups returns a KuberaBackupInformer.
	KuberaBackups() KuberaBackupInformer
	// PodVolumeBackups returns a PodVolumeBackupInformer.
	PodVolumeBackups() PodVolumeBackupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Hooks returns a HookInformer.
func (v *version) Hooks() HookInformer {
	return &hookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KuberaBackups returns a KuberaBackupInformer.
func (v *version) KuberaBackups() KuberaBackupInformer {
	return &kuberaBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodVolumeBackups returns a PodVolumeBackupInformer.
func (v *version) PodVolumeBackups() PodVolumeBackupInformer {
	return &podVolumeBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}