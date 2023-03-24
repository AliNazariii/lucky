package redlock

import (
	"github.com/spf13/cobra"
	"lucky/internal/config"
	"lucky/internal/modules/sandbox"
)

func RunRedLock(cmd *cobra.Command, args []string) {
	configs := config.GetConfig("lucky", "config.yaml")

	redLockModule := New(configs.Redis)

	sandboxModule := sandbox.New(redLockModule)
	sandboxModule.Run()
}
