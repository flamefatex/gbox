package doc

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "doc",
	Short: "proto doc generator",
	Long:  "proto doc generator",
	Args:  cobra.MinimumNArgs(0),
	Run:   run,
}

var param = &paramInfo{}

type paramInfo struct {
	Src  string // proto源文件目录
	Out  string // 输出文件目录
	Type string // proto包的根路径
}

func init() {
	Cmd.Flags().StringVarP(&param.Src, "source", "s", "./src", "proto源文件目录")
	Cmd.Flags().StringVarP(&param.Out, "out", "o", "./doc", "输出文件目录")
	Cmd.Flags().StringVarP(&param.Type, "type", "t", "all", "文档类型，all|json|markdown|html")
}
