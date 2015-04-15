package main

import (
	server "app/dbserver"
	"common/config"
	"common/logger"
	"flag"
	"net"
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

	dbServer := server.NewDBServer(dbcfg)

	tsock, err := net.Listen("tcp", dbcfg.DBHost)
	if err != nil {
		logger.Fatal("net.Listen: %s", err.Error())
	}

	server.StartServices(dbServer, tsock)

	server.WaitForExit(dbServer)

	tsock.Close()

	logger.Info("stop db server")
}
