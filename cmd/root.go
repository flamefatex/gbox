package cmd

import (
	"fmt"
	"os"

	"github.com/flamefatex/gbox/cmd/proto"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "gbox",
	Version: "version",
}

func Execute() {
	// 绑定
	RootCmd.AddCommand(proto.ProtoCmd)

	// 执行
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
