package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=hook

type Hook struct {
	metav1.TypeMeta

	metav1.ObjectMeta

	// Spec holds the information that needs to exececute
	Spec HookSpec

	// Status holds the information that needs to
	Status HookStatus
}

// HookSpec defines
type HookSpec struct {

	// Command or script that needs to be executed inside the pod
	Command string

	// ContainerName is the name of the container which informs where
	// to execute the command
	ContainerName string

	// IgnoreError state how process need to behave if it encounters
	// an error while executing hook
	IgnoreError bool

	// Timeout states maximum amount of time process needs to wait for
	// command completion. Timeouts will be treated as errors
	TimeOut *metav1.Duration
}

type HookStatus struct {

	// ErrorMsg records timeout errors or internal errors if encounters
	ErrorMsg string

	// StdOutput will holds the stdout of hook execution
	StdOutput string

	// StdError will holds the stderr of hook execution
	StdError string

	// Phase will contains the status of hook execution
	Phase HookPhase
}

// HookPhase is a string that represents the status of hook execution
type HookPhase string

const (
	// HookInProgress states hook execution is in progress
	HookInProgress HookPhase = "InProgress"

	// HookCompleted states hook execution is completed
	HookCompleted HookPhase = "Completed"

	// HookFailed states hook execution is failed
	HookFailed HookPhase = "Failed"
)

type HookSpecName string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=hook

type HookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Hook `json:"items"`
}
