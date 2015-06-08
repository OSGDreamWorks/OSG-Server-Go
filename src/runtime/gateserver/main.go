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
	"github.com/Shopify/go-lua"
)

var (
	gateConfigFile = flag.String("c", "etc/gateserver.json", "config file name for the game server")
)

func Hello(state *lua.State) int {
	num := state.Top()
	param, err := state.ToString(1)
	if(err) {
		logger.Debug("testing Func param:%d, string : %s", num, param)
	}else {
		logger.Debug("testing Func param:%d, err : %v", num, err)
	}
	state.PushGoFunction(Hello)
	return 1
}

func main() {
	logger.Info("start gate server")

	lusScript := script.NewScript()

	lusScript.RegisterGlobalFunction("Hello", Hello)
	lusScript.ExecuteString("print('testing C API')")
	lusScript.ExecuteScriptFile("script/app/test.lua")
	lusScript.TestTable("ATestTB")
	lusScript.ExecuteString("print_r(_G)")
	lusScript.ExecuteString("log(ATestTB.ok)")

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
