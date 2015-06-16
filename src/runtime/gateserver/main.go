package main

import (
	"common/logger"
	"component/script"
)

func main() {
	logger.Info("start login server")

	script.DefaultLuaScript().ExecuteScriptFile("script/runtime/loginserver/main.lua")

	logger.Info("stop login server")
}
