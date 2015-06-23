package main

import (
	server "app/dbserver"
	"common/config"
	"common/logger"
	"flag"
	"common"
)

var (
	dbConfigFile = flag.String("c", "etc/dbBase.json", "config file name for the dbserver")
)

func main() {
	logger.Info("start db server")

	var dbcfg config.DBConfig
	if err := config.ReadConfig(*dbConfigFile, &dbcfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	server.CreateServices(dbcfg)

	common.WatchSystemSignal()

	logger.Info("stop db server")
}
