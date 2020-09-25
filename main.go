package main

import (
	"fmt"

	"github.com/flamefatex/config"
	"github.com/flamefatex/log"

	"github.com/flamefatex/gbox/cmd"
)

var (
	Version     string
	GitCommit   string
	serviceName string = "gbox" // 服务名称
)

func main() {
	// config
	config.Init(serviceName)

	// log
	logConfig := &log.Config{
		ServiceName:    serviceName,
		Level:          config.Config().GetString("log.level"),
		EnableConsole:  config.Config().GetBool("log.enable_console"),
		EnableRotation: false,
	}
	log.InitLogger(log.NewZapLogger, logConfig)

	//  version
	version := fmt.Sprintf("%s, build %s", Version, GitCommit)

	cmd.RootCmd.Version = version
	cmd.Execute()
}
