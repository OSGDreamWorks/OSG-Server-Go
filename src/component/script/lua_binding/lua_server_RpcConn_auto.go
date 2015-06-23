package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

const luaRpcConnTypeName = "RpcConn"

var indexRpcConnMethods = map[string]lua.LGFunction{
    "SetResultServer": Register_lua_server_RpcConn_SetResultServer,
    "IsWebConn": Register_lua_server_RpcConn_IsWebConn,
    "WriteObj": Register_lua_server_RpcConn_Call,
    "Call": Register_lua_server_RpcConn_Call,
    "GetId": Register_lua_server_RpcConn_GetId,
    "Lock": Register_lua_server_RpcConn_Lock,
    "Unlock": Register_lua_server_RpcConn_Unlock,
}

func Register_lua_server_RpcConn(L *lua.LState) {
    logger.Debug("Register_server_%s", luaRpcConnTypeName)
    conn := &server.ProtoBufConn{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaRpcConnTypeName, conn)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_server_RpcConn_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_server_RpcConn_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaRpcConnTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexRpcConnMethods))
    DefaultScript.RegisterGlobalClassEnd(luaRpcConnTypeName)
}

func Register_lua_server_RpcConn_newClass(L *lua.LState) int {
    svc := (*server.RpcConn)(nil)
    ud := L.NewUserData()
    ud.Value = &svc
    L.SetMetatable(ud, L.GetTypeMetatable(luaRpcConnTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_server_RpcConn_SetResultServer(L *lua.LState) int {
    ud := L.CheckUserData(1)
    resultServer:= L.CheckString(2)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
       (*v).SetResultServer(resultServer)
    }
    return 0
}

func Register_lua_server_RpcConn_IsWebConn(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
        L.Push(lua.LBool((*v).IsWebConn()))
    }
    return 1
}

func Register_lua_server_RpcConn_Call(L *lua.LState) int {
    ud := L.CheckUserData(1)
    method := L.CheckString(2)
    buffer := L.CheckString(3)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
        err := (*v).Call(method,[]byte(buffer))
        if err != nil {
            logger.Error("lua_server_ProtoBufConn_WriteObj Error : %s", err.Error())
        }
    }
    return 0
}

func Register_lua_server_RpcConn_GetId(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
        cid := (*v).GetId()
        L.Push(lua.LNumber(cid))
    }
    return 1
}

func Register_lua_server_RpcConn_Lock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
        (*v).Lock()
    }
    return 0
}

func Register_lua_server_RpcConn_Unlock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if  v, ok := ud.Value.(*server.RpcConn); ok {
        (*v).Unlock()
    }
    return 0
}

