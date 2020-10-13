package main

import (
	executor "github.com/mayadata.io/kubera-backup-restore/cmd/kubera/executor"
	"k8s.io/klog/v2"
)

func main() {
	defer klog.Flush()
	err := executor.NewCommand().Execute()
	kubera.CheckError(err)
}
