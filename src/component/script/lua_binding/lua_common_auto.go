package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common"
    "common/logger"
    "component/server"
)

const luaWatchSystemSignalFuncName = "WatchSystemSignal"
const luaWriteObjFuncName = "WriteObj"

func Register_lua_common(L *lua.LState) {
    logger.Debug("Register_lua_common")
    DefaultScript.RegisterGlobalFunction(luaWatchSystemSignalFuncName, Register_lua_common_WatchSystemSignal)
    DefaultScript.RegisterGlobalFunction(luaWriteObjFuncName, Register_lua_common_WriteObj)
}

func Register_lua_common_WatchSystemSignal(L *lua.LState) int {
    common.WatchSystemSignal()
    return 0
}

func Register_lua_common_WriteObj(L *lua.LState) int {
    logger.Debug("lua_common_WriteObj")
    ud := L.CheckUserData(1)
    buffer := L.CheckString(2)
    if v, ok := ud.Value.(*server.ProtoBufConn); ok {
        err := v.WriteObj([]byte(buffer))
        if err != nil {
            logger.Error("lua_server_ProtoBufConn_WriteObj Error : %s", err.Error())
        }
    }
    return 0
}