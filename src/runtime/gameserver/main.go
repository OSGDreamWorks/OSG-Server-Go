package main

import (
	server "app/connector"
	"common/logger"
	"component/script"
	"flag"
	"common/config"
	"common"
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

	script.DefaultJsScript().ExecuteScriptFile("script/js/common/logger.js");
	script.DefaultJsScript().ExecuteScriptFile("script/js/main.js");

	server.CreateConnectorServerForClient(cfg)

	common.WatchSystemSignal()
	//script.DefaultLuaScript().ExecuteScriptFile("script/runtime/gameserver/main.lua")

	logger.Info("stop game server")
}
