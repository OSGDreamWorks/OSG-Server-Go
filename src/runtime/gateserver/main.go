package main

import (
	"common/logger"
	"component/script"
)

func main() {
	logger.Info("start login server")

	script.DefaultLuaScript().ExecuteScriptFile("script/runtime/gateserver/main.lua")

	logger.Info("stop login server")
}
