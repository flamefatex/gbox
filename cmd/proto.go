package cmd

import (
	"github.com/flamefatex/gbox/service/proto"
	"github.com/spf13/cobra"
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
