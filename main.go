package main

import (
	"github.com/flamefatex/config"
	"github.com/flamefatex/gbox/cmd"
	"github.com/flamefatex/log"
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

	// version
	log.Infof("ServiceName: %s", serviceName)
	log.Infof("Version: %s", Version)
	log.Infof("GitCommit: %s", GitCommit)

	//
	cmd.Execute()
}
