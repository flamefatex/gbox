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
	ProtoCmd.Flags().StringVarP(&ProtoParam.Src, "source", "s", "./src", "proto源文件目录")
	ProtoCmd.Flags().StringVarP(&ProtoParam.Out, "out", "o", "./goout", "输出文件目录")
	ProtoCmd.Flags().StringVarP(&ProtoParam.PackageRoot, "package_root", "p", "", "proto包的根路径，example:/github.com/flamefatex/protos/goout")
	RootCmd.AddCommand(ProtoCmd)
}

func ProtoDo(cmd *cobra.Command, args []string) {
	proto.Do(cmd, args, ProtoParam)
}
