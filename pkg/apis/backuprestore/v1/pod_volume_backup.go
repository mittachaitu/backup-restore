package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=podvolumebackup

type PodVolumeBackup struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodVolumeBackupSpec `json:"spec,omitempty"`

	Status PodVolumeBackupStatus `json:"status,omitempty"`
}

type PodVolumeBackupSpec struct {
	PreHook HookSpecName

	AppBackupConfig AppBackupConfig

	PostHook HookSpecName

	// BackupStorageLocation is the name of the backup storage location
	BackupStorageLocation string
}

type PodVolumeBackupStatus struct {
	// Phase represents the current phase of pod volume backup
	Phase PodVolumeBackupPhase

	// StartTimestamp records the starttime when backup is instantiated
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// CompletedTimestamp records the completedtime when the backup is completed
	CompletedTimestamp *metav1.Time `json:"completedTimestam,omitempty"`

	PreHookStatus HookStatus

	PostHookStatus HookStatus

	ActionHookStatus HookStatus

	// Progress holds the total number of bytes of the volume and the current
	// number of backed up bytes. This can be used to display progress information
	// about the backup operation.
	Progress PodVolumeOperationProgress `json:"progress,omitempty"`
}

type PodVolumeBackupPhase string

const (
	PodVolumeBackupPhaseNew        PodVolumeBackupPhase = "New"
	PodVolumeBackupPhaseInProgress PodVolumeBackupPhase = "InProgress"
	PodVolumeBackupPhaseCompleted  PodVolumeBackupPhase = "Completed"
	PodVolumeBackupPhaseFailed     PodVolumeBackupPhase = "Failed"
)

// PodVolumeOperationProgress represents the progress of a
// PodVolumeBackup/Restore (restic) operation
type PodVolumeOperationProgress struct {
	// +optional
	TotalBytes int64 `json:"totalBytes,omitempty"`

	// +optional
	BytesDone int64 `json:"bytesDone,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=podvolumebackup

type PodVolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PodVolumeBackup `json:"items"`
}
