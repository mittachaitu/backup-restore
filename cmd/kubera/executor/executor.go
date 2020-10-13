package executor

import (
	"flag"

	"github.com/mayadata.io/kubera-backup-restore/pkg/client"
	server "github.com/mayadata.io/kubera-backup-restore/pkg/server/cmd"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// NewCommand returns new instance of cobra cli command
func NewCommand() *cobra.Command {
	// Load the configuration from system if exists
	config, err := client.LoadConfig()
	if err != nil {
		klog.Warningf("Failed to read config file: %v", err)
	}

	c := &cobra.Command{
		Use:   "kubera-protect.",
		Short: "Backup and Restore Kubernetes resources.",
		Long: `Kubera-protect is an agent to handle disaster
recovery for Kubernetes resources. It provides way to backup
kubernetes resources, applications and its data.`}

	f := client.NewFactory(config)
	// f.BindFlags(c.PersistentFlags())

	c.AddCommand(
		server.NewCommand(f),
	)

	klog.InitFlags(flag.CommandLine)
	c.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	return c
}
