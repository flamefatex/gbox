package proto

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "proto",
	Short: "proto generator",
	Long:  "proto generator, generate go grpc-gateway validators",
	Args:  cobra.MinimumNArgs(0),
	Run:   Run,
}

var Param = &ParamInfo{}

type ParamInfo struct {
	Src         string // proto源文件目录
	Out         string // 输出文件目录
	PackageRoot string // proto包的根路径
}

func init() {
	Cmd.Flags().StringVarP(&Param.Src, "source", "s", "./src", "proto源文件目录")
	Cmd.Flags().StringVarP(&Param.Out, "out", "o", "./goout", "输出文件目录")
	Cmd.Flags().StringVarP(&Param.PackageRoot, "package_root", "p", "", "proto包的根路径，example:/github.com/flamefatex/protos/goout")
}
