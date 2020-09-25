package main

import (
	"github.com/flamefatex/config"
	"github.com/flamefatex/gbox/cmd"
	"github.com/flamefatex/gbox/service/version"

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
	version.SetVersionInfo(&version.VersionInfo{
		ServiceName: serviceName,
		Version:     Version,
		GitCommit:   GitCommit,
	})

	cmd.Execute()
}
