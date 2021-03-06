// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppBackupConfig) DeepCopyInto(out *AppBackupConfig) {
	*out = *in
	if in.DataToBackup != nil {
		in, out := &in.DataToBackup, &out.DataToBackup
		*out = new(DataToBackup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppBackupConfig.
func (in *AppBackupConfig) DeepCopy() *AppBackupConfig {
	if in == nil {
		return nil
	}
	out := new(AppBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppLevelBackupConfig) DeepCopyInto(out *AppLevelBackupConfig) {
	*out = *in
	if in.AppBackupConfig != nil {
		in, out := &in.AppBackupConfig, &out.AppBackupConfig
		*out = make(map[string]AppBackupConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.FileSystemBackup != nil {
		in, out := &in.FileSystemBackup, &out.FileSystemBackup
		*out = make(map[string]FileSystemBackupConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppLevelBackupConfig.
func (in *AppLevelBackupConfig) DeepCopy() *AppLevelBackupConfig {
	if in == nil {
		return nil
	}
	out := new(AppLevelBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfig) DeepCopyInto(out *BackupConfig) {
	*out = *in
	in.AppLevelBackupConfig.DeepCopyInto(&out.AppLevelBackupConfig)
	in.VolumeLevelBackupConfig.DeepCopyInto(&out.VolumeLevelBackupConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfig.
func (in *BackupConfig) DeepCopy() *BackupConfig {
	if in == nil {
		return nil
	}
	out := new(BackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfigStatus) DeepCopyInto(out *BackupConfigStatus) {
	*out = *in
	out.PreHookStatus = in.PreHookStatus
	out.PostHookStatus = in.PostHookStatus
	out.Action = in.Action
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfigStatus.
func (in *BackupConfigStatus) DeepCopy() *BackupConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BackupConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupProgress) DeepCopyInto(out *BackupProgress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupProgress.
func (in *BackupProgress) DeepCopy() *BackupProgress {
	if in == nil {
		return nil
	}
	out := new(BackupProgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIBackupConfig) DeepCopyInto(out *CSIBackupConfig) {
	*out = *in
	if in.ResourceSpecificCSIBackupConfig != nil {
		in, out := &in.ResourceSpecificCSIBackupConfig, &out.ResourceSpecificCSIBackupConfig
		*out = make(map[string]ResourceSpecificCSIBackupConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIBackupConfig.
func (in *CSIBackupConfig) DeepCopy() *CSIBackupConfig {
	if in == nil {
		return nil
	}
	out := new(CSIBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataToBackup) DeepCopyInto(out *DataToBackup) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataToBackup.
func (in *DataToBackup) DeepCopy() *DataToBackup {
	if in == nil {
		return nil
	}
	out := new(DataToBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemBackupConfig) DeepCopyInto(out *FileSystemBackupConfig) {
	*out = *in
	if in.ExcludeMountPoints != nil {
		in, out := &in.ExcludeMountPoints, &out.ExcludeMountPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemBackupConfig.
func (in *FileSystemBackupConfig) DeepCopy() *FileSystemBackupConfig {
	if in == nil {
		return nil
	}
	out := new(FileSystemBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GVKName) DeepCopyInto(out *GVKName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GVKName.
func (in *GVKName) DeepCopy() *GVKName {
	if in == nil {
		return nil
	}
	out := new(GVKName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookList) DeepCopyInto(out *HookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookList.
func (in *HookList) DeepCopy() *HookList {
	if in == nil {
		return nil
	}
	out := new(HookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookSpec) DeepCopyInto(out *HookSpec) {
	*out = *in
	if in.TimeOut != nil {
		in, out := &in.TimeOut, &out.TimeOut
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookSpec.
func (in *HookSpec) DeepCopy() *HookSpec {
	if in == nil {
		return nil
	}
	out := new(HookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookStatus) DeepCopyInto(out *HookStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookStatus.
func (in *HookStatus) DeepCopy() *HookStatus {
	if in == nil {
		return nil
	}
	out := new(HookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sResource) DeepCopyInto(out *K8sResource) {
	*out = *in
	if in.GVKName != nil {
		in, out := &in.GVKName, &out.GVKName
		*out = make([]GVKName, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sResource.
func (in *K8sResource) DeepCopy() *K8sResource {
	if in == nil {
		return nil
	}
	out := new(K8sResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sVersion) DeepCopyInto(out *K8sVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sVersion.
func (in *K8sVersion) DeepCopy() *K8sVersion {
	if in == nil {
		return nil
	}
	out := new(K8sVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeBackupStatus) DeepCopyInto(out *KubeBackupStatus) {
	*out = *in
	if in.AppLevelBackupStatus != nil {
		in, out := &in.AppLevelBackupStatus, &out.AppLevelBackupStatus
		*out = make(map[string][]BackupConfigStatus, len(*in))
		for key, val := range *in {
			var outVal []BackupConfigStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]BackupConfigStatus, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.VolumeLevelBackupStatus != nil {
		in, out := &in.VolumeLevelBackupStatus, &out.VolumeLevelBackupStatus
		*out = make(map[string]VolumeLevelBackupStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Progress = in.Progress
	out.ClusterInfo = in.ClusterInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeBackupStatus.
func (in *KubeBackupStatus) DeepCopy() *KubeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(KubeBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberaBackup) DeepCopyInto(out *KuberaBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberaBackup.
func (in *KuberaBackup) DeepCopy() *KuberaBackup {
	if in == nil {
		return nil
	}
	out := new(KuberaBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberaBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberaBackupList) DeepCopyInto(out *KuberaBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KuberaBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberaBackupList.
func (in *KuberaBackupList) DeepCopy() *KuberaBackupList {
	if in == nil {
		return nil
	}
	out := new(KuberaBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberaBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberaBackupSpec) DeepCopyInto(out *KuberaBackupSpec) {
	*out = *in
	if in.IncludeResources != nil {
		in, out := &in.IncludeResources, &out.IncludeResources
		*out = make([]K8sResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExcludeResources != nil {
		in, out := &in.ExcludeResources, &out.ExcludeResources
		*out = make([]K8sResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PreHooks != nil {
		in, out := &in.PreHooks, &out.PreHooks
		*out = make(map[string]HookSpecName, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PostHooks != nil {
		in, out := &in.PostHooks, &out.PostHooks
		*out = make(map[string]HookSpecName, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.BackupConfig.DeepCopyInto(&out.BackupConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberaBackupSpec.
func (in *KuberaBackupSpec) DeepCopy() *KuberaBackupSpec {
	if in == nil {
		return nil
	}
	out := new(KuberaBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeBackup) DeepCopyInto(out *PodVolumeBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeBackup.
func (in *PodVolumeBackup) DeepCopy() *PodVolumeBackup {
	if in == nil {
		return nil
	}
	out := new(PodVolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodVolumeBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeBackupList) DeepCopyInto(out *PodVolumeBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodVolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeBackupList.
func (in *PodVolumeBackupList) DeepCopy() *PodVolumeBackupList {
	if in == nil {
		return nil
	}
	out := new(PodVolumeBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodVolumeBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeBackupSpec) DeepCopyInto(out *PodVolumeBackupSpec) {
	*out = *in
	in.AppBackupConfig.DeepCopyInto(&out.AppBackupConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeBackupSpec.
func (in *PodVolumeBackupSpec) DeepCopy() *PodVolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PodVolumeBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeBackupStatus) DeepCopyInto(out *PodVolumeBackupStatus) {
	*out = *in
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		*out = (*in).DeepCopy()
	}
	out.PreHookStatus = in.PreHookStatus
	out.PostHookStatus = in.PostHookStatus
	out.ActionHookStatus = in.ActionHookStatus
	out.Progress = in.Progress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeBackupStatus.
func (in *PodVolumeBackupStatus) DeepCopy() *PodVolumeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PodVolumeBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeOperationProgress) DeepCopyInto(out *PodVolumeOperationProgress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeOperationProgress.
func (in *PodVolumeOperationProgress) DeepCopy() *PodVolumeOperationProgress {
	if in == nil {
		return nil
	}
	out := new(PodVolumeOperationProgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecificCSIBackupConfig) DeepCopyInto(out *ResourceSpecificCSIBackupConfig) {
	*out = *in
	if in.timeout != nil {
		in, out := &in.timeout, &out.timeout
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecificCSIBackupConfig.
func (in *ResourceSpecificCSIBackupConfig) DeepCopy() *ResourceSpecificCSIBackupConfig {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecificCSIBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProviderBackupConfig) DeepCopyInto(out *StorageProviderBackupConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProviderBackupConfig.
func (in *StorageProviderBackupConfig) DeepCopy() *StorageProviderBackupConfig {
	if in == nil {
		return nil
	}
	out := new(StorageProviderBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeLevelBackupConfig) DeepCopyInto(out *VolumeLevelBackupConfig) {
	*out = *in
	in.CSIBackupConfig.DeepCopyInto(&out.CSIBackupConfig)
	if in.StorageProviderBkpConfig != nil {
		in, out := &in.StorageProviderBkpConfig, &out.StorageProviderBkpConfig
		*out = make(map[string]StorageProviderBackupConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeLevelBackupConfig.
func (in *VolumeLevelBackupConfig) DeepCopy() *VolumeLevelBackupConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeLevelBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeLevelBackupStatus) DeepCopyInto(out *VolumeLevelBackupStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeLevelBackupStatus.
func (in *VolumeLevelBackupStatus) DeepCopy() *VolumeLevelBackupStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeLevelBackupStatus)
	in.DeepCopyInto(out)
	return out
}
