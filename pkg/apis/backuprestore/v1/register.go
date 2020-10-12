package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Resource takes unqualified resource and returns GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

type typeInfo struct {
	PluralName   string
	ItemType     runtime.Object
	ItemListType runtime.Object
}

func newTypeInfo(pluralName string, itemType, itemListType runtime.Object) typeInfo {
	return typeInfo{
		PluralName:   pluralName,
		ItemType:     itemType,
		ItemListType: itemListType,
	}
}

// CustomResources retruns map of resources under kubera backup
// restore API group
func CustomResources() map[string]typeInfo {
	return map[string]typeInfo{
		"KuberaBackup":    newTypeInfo("kuberabackups", &KuberaBackup{}, &KuberaBackupList{}),
		"Hook":            newTypeInfo("hooks", &Hook{}, &HookList{}),
		"PodVolumeBackup": newTypeInfo("podvolumebackups", &PodVolumeBackup{}, &PodVolumeBackupList{}),
	}
}

func addKnownTypes(scheme *runtime.Scheme) error {
	for _, typeInfo := range CustomResources() {
		scheme.AddKnownTypes(SchemeGroupVersion, typeInfo.ItemType, typeInfo.ItemListType)
	}

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
