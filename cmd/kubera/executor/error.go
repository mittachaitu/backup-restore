package executor

import "k8s.io/client-go/tools/cache"

func (c *backupController) getBackupEventHandlers() cache.ResourceEventHandler {
	if !s.config.isNonAdminMode {
		return cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				// No need to verify type assertion because obj will
				// be definetly of type KuberaBackup
				backup := obj.(*kuberaapis.KuberaBackup)

				c.queue.Add(key)
			},
		}
	}
	return cache.ResourceEventHandlerFuncs{}
}
