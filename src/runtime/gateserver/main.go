package main

import (
	"common/logger"
	"flag"
	"os"
	"syscall"
	"common"
	"component/script"
)

var (
	gateConfigFile = flag.String("c", "etc/gateserver.json", "config file name for the game server")
)

func main() {
	logger.Info("start gate server")

//	var cfg config.GateConfig
//	if err := config.ReadConfig(*gateConfigFile, &cfg); err != nil {
//		logger.Fatal("load config failed, error is: %v", err)
//		return
//	}
//
//	go gateserver.CreateGateServicesForServer(cfg)
//	go gateserver.CreateGateServicesForClient(cfg)
	script.DefaultLuaScript().ExecuteScriptFile("script/app/gateserver/main.lua")
	script.DefaultLuaScript().ExecuteString("print_r(_G)")

	handler := func(s os.Signal, arg interface{}) {
		logger.Info("handle signal: %v\n", s)
		logger.Info("stop game server")
		os.Exit(0)
	}

	handlerArray := []os.Signal{syscall.SIGINT,
		syscall.SIGILL,
		syscall.SIGFPE,
		syscall.SIGSEGV,
		syscall.SIGTERM,
		syscall.SIGABRT}

	common.WatchSystemSignal(&handlerArray, handler)

	logger.Info("stop gate server")
}
