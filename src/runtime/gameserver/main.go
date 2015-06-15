package main

import (
	"common/logger"
	"component/script"
)

func main() {
	logger.Info("start game server")

	script.DefaultLuaScript().ExecuteScriptFile("script/runtime/gameserver/main.lua")

	logger.Info("stop game server")
}
