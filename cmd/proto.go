package cmd

import (
	"github.com/flamefatex/gbox/service/proto"
	"github.com/spf13/cobra"
)

var ProtoCmd = &cobra.Command{
	Use:   "proto",
	Short: "proto gen ",
	Long:  "proto gen ....description",
	Args:  cobra.MinimumNArgs(0),
	Run:   ProtoDo,
}

var ProtoParam = &proto.Param{}

func init() {
	ProtoCmd.Flags().StringVarP(&ProtoParam.Src, "source", "s", "", "proto source dir")
	RootCmd.AddCommand(ProtoCmd)
}

func ProtoDo(cmd *cobra.Command, args []string) {
	proto.Do(cmd, args, ProtoParam)
}
