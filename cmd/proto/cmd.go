package proto

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "proto",
	Short: "proto generator",
	Long:  "proto generator, generate go grpc-gateway validators",
	Args:  cobra.MinimumNArgs(0),
	Run:   run,
}

var param = &paramInfo{}

type paramInfo struct {
	Src         string // proto源文件目录
	Out         string // 输出文件目录
	PackageRoot string // proto包的根路径
}

func init() {
	Cmd.Flags().StringVarP(&param.Src, "source", "s", "./src", "proto源文件目录")
	Cmd.Flags().StringVarP(&param.Out, "out", "o", "./goout", "输出文件目录")
	Cmd.Flags().StringVarP(&param.PackageRoot, "package_root", "p", "", "proto包的根路径，example:/github.com/flamefatex/protos/goout")
}
