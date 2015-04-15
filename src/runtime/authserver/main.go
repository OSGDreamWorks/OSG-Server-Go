package main

import (
	server "app/authserver"
	"common/config"
	"common/logger"
	"flag"
	"net"
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

	authServer := server.NewAuthServer(authcfg)

	tsock, err := net.Listen("tcp", authcfg.AuthHost)
	if err != nil {
		logger.Fatal("net.Listen: %s", err.Error())
	}

	server.StartServices(authServer, tsock)

	server.WaitForExit(authServer)

	tsock.Close()

	logger.Info("stop auth server")
}
