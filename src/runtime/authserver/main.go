package main

import (
	server "app/authserver"
	"common/config"
	"common/logger"
	"flag"
	"common"
)

var (
	authConfigFile = flag.String("c", "etc/authserver.json", "config file name for the auth server")
)

func main() {
	logger.Info("start auth server")

	var authcfg config.AuthConfig
	if err := config.ReadConfig(*authConfigFile, &authcfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	server.CreateServices(authcfg)

	common.WatchSystemSignal()

	logger.Info("stop auth server")
}
