package etcdlock

import (
	"github.com/spf13/cobra"
	"lucky/internal/config"
	"lucky/internal/modules/sandbox"
)

func RunEtcdLock(cmd *cobra.Command, args []string) {
	configs := config.GetConfig("lucky", "config.yaml")

	etcdLockModule := New(configs.Etcd)

	sandboxModule := sandbox.New(etcdLockModule)
	sandboxModule.Run()
}
