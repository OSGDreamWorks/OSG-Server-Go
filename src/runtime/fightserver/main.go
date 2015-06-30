package main

import (
    "common/logger"
    "component/script"
    "flag"
    "fmt"
)
var (
    svrConfigFile = flag.String("c", "etc/gameserver.json", "config file name for the fight server")
    fightServerId = flag.Uint64("n", 1, "config id for the fight server")
)

func main() {
    logger.Info("start fight server")

    script.DefaultLuaScript().ExecuteString(fmt.Sprintf("_G.svrConfigFile, _G.fightServerId = \"%s\", %d", *svrConfigFile, *fightServerId))
    script.DefaultLuaScript().ExecuteScriptFile("script/runtime/fightserver/main.lua")

    logger.Info("stop fight server")
}
