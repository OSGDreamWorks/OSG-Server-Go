package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

const luaServerTypeName = "Server"

func Register_lua_server_Server(L *lua.LState) int {
    logger.Debug("Register_server_%s", luaServerTypeName)
    svc := &server.Server{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaServerTypeName, svc)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_server_Server_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexServerMethods))
    //DefaultScript.RegisterGlobalClassEnd(luaServerTypeName, svc)
    return 1
}

var indexServerMethods = map[string]lua.LGFunction{
    "ServeConn": Register_lua_server_Server_ServeConn,
}

func Register_lua_server_Server_newClass(L *lua.LState) int {
    svc := server.NewServer()
    ud := L.NewUserData()
    ud.Value = svc
    L.SetMetatable(ud, L.GetTypeMetatable(luaServerTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_server_Server_ServeConn(L *lua.LState) int {
    ud := L.CheckUserData(1)
    arg := L.CheckUserData(2)
    if v, ok := ud.Value.(*server.Server); ok {
        if c, ok := arg.Value.(*server.RpcConn); ok{
            v.ServeConn(*c)
        }
    }
    return 0
}