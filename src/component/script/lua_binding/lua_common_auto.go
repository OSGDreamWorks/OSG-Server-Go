package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common"
    "common/logger"
    "component/server"
    "time"
)

const luaWatchSystemSignalFuncName = "WatchSystemSignal"
const luaWriteObjFuncName = "WriteObj"
const luaSetIntervalFuncName = "SetInterval"
const luaClearIntervalFuncName = "ClearInterval"

var indexSetIntervalMethods = map[string]*lua.LFunction{

}

func Register_lua_common(L *lua.LState) {
    logger.Debug("Register_lua_common")
    DefaultScript.RegisterGlobalFunction(luaWatchSystemSignalFuncName, Register_lua_common_WatchSystemSignal)
    DefaultScript.RegisterGlobalFunction(luaWriteObjFuncName, Register_lua_common_WriteObj)
    DefaultScript.RegisterGlobalFunction(luaSetIntervalFuncName, Register_lua_common_SetInterval)
    DefaultScript.RegisterGlobalFunction(luaClearIntervalFuncName, Register_lua_common_ClearInterval)
}

func Register_lua_common_WatchSystemSignal(L *lua.LState) int {
    common.WatchSystemSignal()
    return 0
}

func Register_lua_common_WriteObj(L *lua.LState) int {
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

func Register_lua_common_SetInterval(L *lua.LState) int {
    identifier := L.CheckString(1)
    delay := L.CheckNumber(2)
    duration := L.CheckNumber(3)
    function := L.CheckFunction(4)
    indexSetIntervalMethods[identifier] = function

    go func() {
        time.Sleep(time.Duration(delay) * time.Second)
        for {
            err2 := L.CallByParam(lua.P{
                Fn: indexSetIntervalMethods[identifier],
                NRet: 1,
                Protect: true,
            })

            if err2 !=nil {
                logger.Warning("SetInterval [%v] already be clear!", identifier)
                break
            }

            time.Sleep(time.Duration(duration) * time.Second)
        }
    }()
    return 0
}

func Register_lua_common_ClearInterval(L *lua.LState) int {
    identifier := L.CheckString(1)
    indexSetIntervalMethods[identifier] = nil
    return  0
}