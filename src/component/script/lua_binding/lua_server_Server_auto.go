package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

const luaServerTypeName = "Server"

var indexServerMethods = map[string]lua.LGFunction{
    "Register": Register_lua_server_Server_Register,
    "ListenAndServe": Register_lua_server_Server_ListenAndServe,
    "ServeConn": Register_lua_server_Server_ServeConn,
    "Lock": Register_lua_server_Server_Lock,
    "Unlock": Register_lua_server_Server_Unlock,
    "RegCallBackOnConn": Register_lua_server_Server_RegCallBackOnConn,
    "RegCallBackOnDisConn": Register_lua_server_Server_RegCallBackOnDisConn,
    "RegCallBackOnCallBefore": Register_lua_server_Server_RegCallBackOnCallBefore,
    "RegCallBackOnCallAfter": Register_lua_server_Server_RegCallBackOnCallAfter,
}

func Register_lua_server_Server(L *lua.LState) {
    logger.Debug("Register_server_%s", luaServerTypeName)
    svc := &server.Server{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaServerTypeName, svc)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_server_Server_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_server_Server_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaServerTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexServerMethods))
    DefaultScript.RegisterGlobalClassEnd(luaServerTypeName)
}

func Register_lua_server_Server_newClass(L *lua.LState) int {
    svc := server.NewServer()
    svc.SetLuaState(L)
    ud := L.NewUserData()
    ud.Value = svc
    L.SetMetatable(ud, L.GetTypeMetatable(luaServerTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_server_Server_Register(L *lua.LState) int {
    ud := L.CheckUserData(1)
    arg := L.CheckTable(2)
    if v, ok := ud.Value.(*server.Server); ok {
        v.RegisterFromLua(arg)
    }
    return 0
}

func Register_lua_server_Server_ListenAndServe(L *lua.LState) int {
    ud := L.CheckUserData(1)
    tcpAddr := L.CheckString(2)
    httpAddr := L.CheckString(3)
    if v, ok := ud.Value.(*server.Server); ok {
        v.ListenAndServe(tcpAddr, httpAddr)
    }
    return 0
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

func Register_lua_server_Server_Lock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*server.Server); ok {
        v.Lock()
    }
    return 0
}

func Register_lua_server_Server_Unlock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*server.Server); ok {
        v.Unlock()
    }
    return 0
}

func Register_lua_server_Server_RegCallBackOnConn(L *lua.LState) int {
    ud := L.CheckUserData(1)
    luaFn := L.CheckFunction(2)
    if v, ok := ud.Value.(*server.Server); ok {
        v.RegCallBackOnConn(
            func(conn server.RpcConn) {
                udConn := L.NewUserData()
                udConn.Value = &conn
                L.SetMetatable(udConn, L.GetTypeMetatable(luaRpcConnTypeName))

                err2 := L.CallByParam(lua.P{
                    Fn: luaFn,
                    NRet: 0,
                    Protect: true,
                }, udConn)

                if err2 !=nil {
                    logger.Error("RegCallBackOnConn Error : %s", err2.Error())
                }
            },
        )
    }
    return 0
}

func Register_lua_server_Server_RegCallBackOnDisConn(L *lua.LState) int {
    ud := L.CheckUserData(1)
    luaFn := L.CheckFunction(2)
    if v, ok := ud.Value.(*server.Server); ok {
        v.RegCallBackOnDisConn(
        func(conn server.RpcConn) {
            udConn := L.NewUserData()
            udConn.Value = &conn
            L.SetMetatable(udConn, L.GetTypeMetatable(luaRpcConnTypeName))

            err2 := L.CallByParam(lua.P{
                Fn: luaFn,
                NRet: 0,
                Protect: true,
            }, udConn)

            if err2 !=nil {
                logger.Error("RegCallBackOnConn Error : %s", err2.Error())
            }
        },
        )
    }
    return 0
}

func Register_lua_server_Server_RegCallBackOnCallBefore(L *lua.LState) int {
    ud := L.CheckUserData(1)
    luaFn := L.CheckFunction(2)
    if v, ok := ud.Value.(*server.Server); ok {
        v.RegCallBackOnCallBefore(
        func(conn server.RpcConn) {
            udConn := L.NewUserData()
            udConn.Value = &conn
            L.SetMetatable(udConn, L.GetTypeMetatable(luaRpcConnTypeName))

            err2 := L.CallByParam(lua.P{
                Fn: luaFn,
                NRet: 0,
                Protect: true,
            }, udConn)

            if err2 !=nil {
                logger.Error("RegCallBackOnConn Error : %s", err2.Error())
            }
        },
        )
    }
    return 0
}

func Register_lua_server_Server_RegCallBackOnCallAfter(L *lua.LState) int {
    ud := L.CheckUserData(1)
    luaFn := L.CheckFunction(2)
    if v, ok := ud.Value.(*server.Server); ok {
        v.RegCallBackOnCallAfter(
        func(conn server.RpcConn) {
            udConn := L.NewUserData()
            udConn.Value = &conn
            L.SetMetatable(udConn, L.GetTypeMetatable(luaRpcConnTypeName))

            err2 := L.CallByParam(lua.P{
                Fn: luaFn,
                NRet: 0,
                Protect: true,
            }, udConn)

            if err2 !=nil {
                logger.Error("RegCallBackOnConn Error : %s", err2.Error())
            }
        },
        )
    }
    return 0
}