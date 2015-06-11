package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

const luaProtoBufConnTypeName = "RpcConn"

var indexProtoBufConnMethods = map[string]lua.LGFunction{
    "SetResultServer": Register_lua_server_ProtoBufConn_SetResultServer,
    "WriteObj": Register_lua_server_ProtoBufConn_Call,
    "Call": Register_lua_server_ProtoBufConn_Call,
}

func Register_lua_server_ProtoBufConn(L *lua.LState) {
    logger.Debug("Register_server_%s", luaProtoBufConnTypeName)
    conn := &server.ProtoBufConn{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaProtoBufConnTypeName, conn)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_server_ProtoBufConn_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexProtoBufConnMethods))
    DefaultScript.RegisterGlobalClassEnd(luaProtoBufConnTypeName)
}

func Register_lua_server_ProtoBufConn_newClass(L *lua.LState) int {
    svc := &server.ProtoBufConn{}
    ud := L.NewUserData()
    ud.Value = svc
    L.SetMetatable(ud, L.GetTypeMetatable(luaProtoBufConnTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_server_ProtoBufConn_SetResultServer(L *lua.LState) int {
    ud := L.CheckUserData(1)
    resultServer:= L.CheckString(2)
    if v, ok := ud.Value.(*server.ProtoBufConn); ok {
        v.SetResultServer(resultServer)
    }
    return 0
}

func Register_lua_server_ProtoBufConn_Call(L *lua.LState) int {
    ud := L.CheckUserData(1)
    method := L.CheckString(2)
    buffer := L.CheckString(3)
    if  v, ok := ud.Value.(*server.ProtoBufConn); ok {
        err := v.Call(method, []byte(buffer))
        if err != nil {
            logger.Error("lua_server_ProtoBufConn_WriteObj Error : %s", err.Error())
        }
    }else if  v, ok := ud.Value.(*server.RpcConn); ok {
        err := (*v).Call(method,[]byte(buffer))
        if err != nil {
            logger.Error("lua_server_ProtoBufConn_WriteObj Error : %s", err.Error())
        }
    }
    return 0
}

