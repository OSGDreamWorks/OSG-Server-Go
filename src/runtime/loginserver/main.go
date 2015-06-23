package main

import (
	server "app/loginserver"
	"common/logger"
	"flag"
	"common/config"
	"common"
)

var (
	loginConfigFile = flag.String("c", "etc/loginserver.json", "config file name for the login server")
	authConfigFile = flag.String("a", "etc/authserver.json", "config file name for the auth server")
)

func main() {
	logger.Info("start login server")

	var lgcfg config.LoginConfig
	if err := config.ReadConfig(*loginConfigFile, &lgcfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	var authcfg config.AuthConfig
	if err := config.ReadConfig(*authConfigFile, &authcfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	server.CreateServices(lgcfg, authcfg)

	common.WatchSystemSignal()

	logger.Info("stop login server")
}
