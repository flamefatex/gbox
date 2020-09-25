package cmd

import (
	"github.com/spf13/cobra"

	"github.com/flamefatex/gbox/service/proto"
)

var cmdProto = &cobra.Command{
	Use:   "proto",
	Short: "proto gen ",
	Long:  "proto gen ....description",
	Args:  cobra.MinimumNArgs(0),
	Run:   proto.Do,
}

func init() {
	RootCmd.AddCommand(cmdProto)
}
