package etcdlock

import (
	"fmt"
	"github.com/spf13/cobra"
	"lucky/internal/config"
	"lucky/internal/modules/sandbox"
)

func RunEtcdLock(cmd *cobra.Command, args []string) {
	configs := config.GetConfig("lucky", "config.yaml")

	etcdLockModule, err := New(configs.Etcd)
	if err != nil {
		err := fmt.Errorf("can't create etcd lock module: %v", err)
		fmt.Println(err)
	}

	sandboxModule := sandbox.New(etcdLockModule)
	sandboxModule.Run()
}
