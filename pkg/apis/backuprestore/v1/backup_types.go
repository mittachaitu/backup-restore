package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=kuberabackup

// KuberaBackup represents KuberaBackup custom resource
type KuberaBackup struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KuberaBackupSpec `json:"spec"`

	Status KubeBackupStatus `json:"status,omitempty"`
}

// KuberaBackupSpec contains specifications to instantiate backup
type KuberaBackupSpec struct {
	// IncludeResources is used to specify the resources which
	// needs to be included in backup. User can specify * in all
	// fields to perform backup on all resources
	IncludeResources []K8sResource

	// ExcludeResources can be used to specify the resources that
	// needs to be excluded from backup
	ExcludeResources []K8sResource

	// Helm/Vanilla operator file path which needs to be backed up
	// along with selections
	// NOTE: If the operator has PersistentVolume then data also will
	// be backup
	// TODO: More thought are required to handle changes
	OperatorFile string

	// PreHooks holds the information that needs to be executed
	// before taking an application dump (or) snapshot. User can
	// specify keys as: deployment.apps/ns/deployment-name,
	// statefuleset.apps/ns/statefulset-name and
	// value can be: namespace/hookSpecName
	PreHooks map[string]HookSpecName

	// PostHooks holds the information that needs to be executed
	// after taking an application dump/snapshot. User can specify
	// key as:deployment.apps/ns/deployment-name,
	//        statefulset.apps/ns/statefulset-name and
	// value can be: namespace/hookSpecName
	PostHooks map[string]HookSpecName

	// BackupConfig contains detailed specifications how to backup
	// the data
	BackupConfig BackupConfig

	// VolumeSnapshotLocationName contains the information where to
	// perform backup the data
	VolumeSnapshotLocationName string
}

// K8sResource represents to select the K8s resources
type K8sResource struct {

	// Namespace defines the namespace of the resources
	Namespace string

	// GVKName holds the Group, Version and Kind resource
	GVKName []GVKName

	// Selector is used to filter the K8s resource based on labels
	// selector can be specific to particular namespace
	Selector *metav1.LabelSelector
}

// GVKName holds the group, version, kind of particular resource
type GVKName struct {

	// Group holds the group name of K8s resource
	Group string

	// Version holds the version of K8s resource
	// For K8s resource user no need to specify it explicitly
	Version string

	// Kind is a kind of k8s resource
	Kind string

	// Name holds the name of the resource name
	Name string
}

// BackupConfig specifies the how data needs to be backedup
type BackupConfig struct {
	// AppLevelBackupConfig contains the configuration to perform
	// backup when application exist
	AppLevelBackupConfig AppLevelBackupConfig

	// VolumeLevelBackupConfig contains the configuration to perform
	// backup even without application… much suitable case when
	// application are not in running state
	VolumeLevelBackupConfig VolumeLevelBackupConfig
}

// AppLevelBackupConfig can be specified only if the
// volumes are consuming by application
type AppLevelBackupConfig struct {
	// NOTE: Here possible values for map[string]value possible
	// keys are deployment.apps/<deployment-name>,
	// statefulset.apps/<sts-name>, job-name, pod-name

	// AppBackupConfig defines the app specific configuration to
	// perform app consistent backups. Map is required to map the
	// deployment/sts(application) and way of how to backup application
	// data in app consistent way.
	// - If key is deployment.apps/<deployment-name> then specific
	//   commands will be executed in any one of the application pods
	//   belongs to particular deployment
	// - If the key is statefulset.apps/<sts-name> then specific commands
	//   will be executed in all the application pods belongs to particular
	//   statefulset
	AppBackupConfig map[string]AppBackupConfig

	// FileSystemBackup will backup the data present on the entire
	// mount point/volume. Users can specify list of deployment-names
	// (or) statefulset names as key
	FileSystemBackup map[string]FileSystemBackupConfig
}

// AppBackupConfig holds application specific configuration
// to perform backup
type AppBackupConfig struct {

	// Action will help to create backup by executing app related
	// commands/script. Ex: Pause the IOs; take snapshot/dump; resume
	// the IOs HookSpecName will be <namespace/ResourceName>
	Action HookSpecName

	// Regular expressions for paths/paths for the files that needs
	// to backup to a remote location. If user don’t provide the
	// DataToBackup then program will assume that Action output will
	// contains the path of files to backup with comma separated
	// values. After performing backup to remote location process
	// will delete the files
	DataToBackup *DataToBackup

	// It should be fully registry + image name
	// Credentials are required if it is a private repository
	AppSpecifigImage string

	// RunInBackground enables backup to run in background
	// agent will continue to perform backup with other application
	RunInBackground bool
}

// DataToBackup specifies contains the information about application
// specific data
type DataToBackup struct {
	// Paths specifies path that needs to backup in object location
	Paths []string

	// IncludeActionResult will backup the data present in
	// paths provided by action output
	IncludeActionResult bool
}

// FileSystemBackupConfig contains configuration to take backup
// of entier application data
// Separate CR will be created with information so sidecar will
// reconcile for CR and upload the data into cloud
type FileSystemBackupConfig struct {

	// ExcludeMountPoints will exclude the mount points from backup
	ExcludeMountPoints []string
}

// VolumeLevelBackupConfig contains the specifications
// to take backup at volume level
type VolumeLevelBackupConfig struct {

	// Embedded CSIBackupConfig into volumelevelbackup configuration
	CSIBackupConfig

	// Configurations required to take StorageProvider way of backup
	// Here key can have following possible values
	// Ex: deployment.apps/<deployment-name>, statefulset.apps/<sts-name> and <persistentvolume/name>, <pvc/name>
	// Step1: If key is deployment/stateful set then execute defined pre-hook
	// Step2: Take snapshot(by calling plugin related function)
	// Step3: Execute post snapshot hooks
	// Step4: Trigger backup call

	// TODO: Revisit when we have more clear thoughts on it.
	// One way is we can form
	// map[string]string:{ “storageClassName”: “pluginName” }
	StorageProviderBkpConfig map[string]StorageProviderBackupConfig
}

// CSIBackupConfig will backup the resources in K8s native way
type CSIBackupConfig struct {

	// ResourceSpecificBackupConfig holds the map of
	// application(deployment or sts)/pvc-name
	// If user needs to select different snapshot class name for
	// particular volume/set of volumes belongs to application he can
	// specify particularly
	ResourceSpecificCSIBackupConfig map[string]ResourceSpecificCSIBackupConfig
}

type ResourceSpecificCSIBackupConfig struct {

	// VolumeSnapshotClassName is snapshot class that required to create
	// CSI snapshot
	VolumeSnapshotClassName string

	// TimeOut for successful creation of volumesnapshot
	timeout *metav1.Time
}

// StorageProviderBackupConfig will execute PreHooks
// (It may help to quiesce the IOs) then it will execute
// relative plugin interfaces to take snapshot
type StorageProviderBackupConfig struct {

	// plugin will execute related interfaces to perform backup
	// Plugin related configuration
	pluginConfig string
}

type KubeBackupStatus struct {

	// Overall status of the backup
	Phase BackupPhase

	// Possible key values will be deployment.apps/ns/name, statefulset.apps/ns/name, pod/ns/name
	AppLevelBackupStatus map[string][]BackupConfigStatus

	// possible key values are persistentvolume/pvcname
	VolumeLevelBackupStatus map[string]VolumeLevelBackupStatus

	Progress BackupProgress

	// ClusterInfo will be contains the Kubernetes related version
	ClusterInfo K8sVersion
}

type BackupProgress struct {
	completedCount int

	PendingCount int

	totalCount int
}

type BackupConfigStatus struct {
	PreHookStatus HookStatus

	PostHookStatus HookStatus

	Action HookStatus

	Error string

	Warnings []string
}

type VolumeLevelBackupStatus struct {
	Status BackupConfigStatus

	SnapshotName string
}

type K8sVersion struct {
	Version string
}

type BackupPhase string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=kuberabackup

type KuberaBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KuberaBackup `json:"items"`
}
