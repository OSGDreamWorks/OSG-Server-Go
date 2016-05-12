package main

import (
	server "app/connector"
	"common/logger"
	"flag"
	"common/config"
	"common"
)

var (
	svrConfigFile = flag.String("c", "etc/gameserver.json", "config file name for the game server")
)

// http://127.0.0.1:7980/?method=UpdateScript&script=js for UpdateScript
func main() {
	logger.Info("start game server")

	var cfg config.SvrConfig
	if err := config.ReadConfig(*svrConfigFile, &cfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	server.CreateConnectorServerForClient(cfg)

	common.WatchSystemSignal()

	logger.Info("stop game server")
}
