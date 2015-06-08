package main

import (
	"app/gateserver"
	"common/logger"
	"common/config"
	"flag"
	"os"
	"syscall"
	"common"
	"component/script"
	"github.com/yuin/gopher-lua"
)

var (
	gateConfigFile = flag.String("c", "etc/gateserver.json", "config file name for the game server")
)

func main() {
	logger.Info("start gate server")

	var cfg config.GateConfig
	if err := config.ReadConfig(*gateConfigFile, &cfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	go gateserver.CreateGateServicesForServer(cfg)
	go gateserver.CreateGateServicesForClient(cfg)

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
