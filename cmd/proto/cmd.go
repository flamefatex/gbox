package proto

import (
	"github.com/spf13/cobra"
)

var ProtoCmd = &cobra.Command{
	Use:   "proto",
	Short: "proto gen ",
	Long:  "proto gen ....description",
	Args:  cobra.MinimumNArgs(0),
	Run:   Do,
}

var param = &Param{}

type Param struct {
	Src         string // proto源文件目录
	Out         string // 输出文件目录
	PackageRoot string // proto包的根路径
}

func init() {
	ProtoCmd.Flags().StringVarP(&param.Src, "source", "s", "./src", "proto源文件目录")
	ProtoCmd.Flags().StringVarP(&param.Out, "out", "o", "./goout", "输出文件目录")
	ProtoCmd.Flags().StringVarP(&param.PackageRoot, "package_root", "p", "", "proto包的根路径，example:/github.com/flamefatex/protos/goout")
}
