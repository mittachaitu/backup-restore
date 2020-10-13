package controllers

import (
	kuberaapis "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	"k8s.io/client-go/tools/cache"
)

func (c *backupController) getBackupEventHandlers() cache.ResourceEventHandler {
	if !c.isNonAdminMode {
		return cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				// No need to verify type assertion because obj will
				// be definetly of type KuberaBackup
				backup := obj.(*kuberaapis.KuberaBackup)

				c.queue.Add(backup)
			},
		}
	}
	return cache.ResourceEventHandlerFuncs{}
}
