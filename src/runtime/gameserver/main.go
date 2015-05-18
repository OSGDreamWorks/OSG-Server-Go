package main

import (
	"app/connector"
	"common/logger"
	"common/config"
	"flag"
)

var (
	svrConfigFile = flag.String("c", "etc/gameserver.json", "config file name for the game server")
)

func main() {
	logger.Info("start game server")

	var cfg config.SvrConfig
	if err := config.ReadConfig(*svrConfigFile, &cfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	connector.CreateConnectorServerForClient(&cfg)

	logger.Info("stop game server")
}
