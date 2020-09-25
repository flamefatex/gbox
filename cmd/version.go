package cmd

import (
	"github.com/spf13/cobra"

	"github.com/flamefatex/gbox/service/version"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Long:  "gbox version info",
	Args:  cobra.MinimumNArgs(0),
	Run:   version.Do,
}

func init() {
	RootCmd.AddCommand(cmdVersion)
}
