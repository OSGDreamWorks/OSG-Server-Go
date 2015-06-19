package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/rpc"
)

const luaRpcServerTypeName = "RpcServer"

var indexRpcServerMethods = map[string]lua.LGFunction{
    "Register": Register_lua_rpc_RpcServer_Register,
    "ListenAndServe": Register_lua_rpc_RpcServer_ListenAndServe,
    "Lock": Register_lua_rpc_RpcServer_Lock,
    "Unlock": Register_lua_rpc_RpcServer_Unlock,
}

func Register_lua_rpc_RpcServer(L *lua.LState) {
    logger.Debug("Register_server_%s", luaRpcServerTypeName)
    svc := &rpc.Server{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaRpcServerTypeName, svc)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_rpc_RpcServer_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_rpc_RpcServer_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaRpcServerTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexRpcServerMethods))
    DefaultScript.RegisterGlobalClassEnd(luaRpcServerTypeName)
}

func Register_lua_rpc_RpcServer_newClass(L *lua.LState) int {
    svc := rpc.NewServer()
    svc.SetLuaState(L)
    ud := L.NewUserData()
    ud.Value = svc
    L.SetMetatable(ud, L.GetTypeMetatable(luaRpcServerTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_rpc_RpcServer_Register(L *lua.LState) int {
    ud := L.CheckUserData(1)
    arg := L.CheckTable(2)
    if v, ok := ud.Value.(*rpc.Server); ok {
        v.RegisterFromLua(arg)
    }
    return 0
}

func Register_lua_rpc_RpcServer_ListenAndServe(L *lua.LState) int {
    ud := L.CheckUserData(1)
    tcpAddr := L.CheckString(2)
    if v, ok := ud.Value.(*rpc.Server); ok {
        v.ListenAndServe(tcpAddr)
    }
    return 0
}

func Register_lua_rpc_RpcServer_Lock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*rpc.Server); ok {
        v.Lock()
    }
    return 0
}

func Register_lua_rpc_RpcServer_Unlock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*rpc.Server); ok {
        v.Unlock()
    }
    return 0
}