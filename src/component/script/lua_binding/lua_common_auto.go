package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common"
    "common/logger"
)

const luaWatchSystemSignalFuncName = "WatchSystemSignal"

func Register_lua_common(L *lua.LState) {
    logger.Debug("Register_lua_common")
    DefaultScript.RegisterGlobalFunction(luaWatchSystemSignalFuncName, Register_lua_common_WatchSystemSignal)
}

func Register_lua_common_WatchSystemSignal(L *lua.LState) int {
    common.WatchSystemSignal()
    return 0
}