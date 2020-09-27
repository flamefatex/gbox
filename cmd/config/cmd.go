package config

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "show config",
	Long:  "show config",
	Args:  cobra.MinimumNArgs(0),
	Run:   run,
}

func init() {
}
