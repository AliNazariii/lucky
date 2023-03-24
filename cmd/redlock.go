package cmd

import (
	"github.com/spf13/cobra"
	"lucky/internal/modules/locker/redlock"
)

func init() {
	rootCmd.AddCommand(redLockCmd)
}

var redLockCmd = &cobra.Command{
	Use:   "redlock",
	Short: "Run scenario with RedLock algorithm",
	Long:  "This command runs and check our scenario with RedLock algorithm and a redis cluster",
	Run:   redlock.RunRedLock,
}
