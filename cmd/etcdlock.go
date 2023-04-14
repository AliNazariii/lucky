package cmd

import (
	"github.com/spf13/cobra"
	"lucky/internal/modules/locker/etcdlock"
)

func init() {
	rootCmd.AddCommand(etcdLockCmd)
}

var etcdLockCmd = &cobra.Command{
	Use:   "etcdlock",
	Short: "Run scenario with EtcdLock algorithm",
	Long:  "This command runs and check our scenario with EtcdLock algorithm and a etcd cluster",
	Run:   etcdlock.RunEtcdLock,
}
