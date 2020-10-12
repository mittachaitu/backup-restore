// Package v1 contains API Schema definitions for the kubera backuprestore v1 API group
// +kubebuilder:object:generate=true
// +groupName=backuprestore.kubera.io
package v1

import (
	backuprestore "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	apiVersion = "v1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: backuprestore.GroupName, Version: apiVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
