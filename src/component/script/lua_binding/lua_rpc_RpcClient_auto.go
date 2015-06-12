package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/rpc"
    "net"
    "reflect"
    "code.google.com/p/goprotobuf/proto"
)

const luaRpcClientTypeName = "RpcClient"

var indexRpcClientMethods = map[string]lua.LGFunction{
    "Call": Register_lua_rpc_RpcClient_Call,
    "Close": Register_lua_rpc_RpcClient_Close,
}

func Register_lua_rpc_RpcClient(L *lua.LState) {
    logger.Debug("Register_server_%s", luaRpcClientTypeName)
    cli := rpc.Client{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaRpcClientTypeName, cli)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_rpc_RpcClient_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexRpcClientMethods))
    DefaultScript.RegisterGlobalClassEnd(luaRpcClientTypeName)
}

func Register_lua_rpc_RpcClient_newClass(L *lua.LState) int {
    addr:= L.CheckString(2)
    rpcConn, err := net.Dial("tcp", addr)
    if err != nil {
        logger.Fatal("connect rpc server failed %s", err.Error())
        L.Push(lua.LNil)
        return 1
    }
    ud := L.NewUserData()
    ud.Value = rpc.NewClient(rpcConn)
    L.SetMetatable(ud, L.GetTypeMetatable(luaRpcClientTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_rpc_RpcClient_Call(L *lua.LState) int {
    ud := L.CheckUserData(1)
    method:= L.CheckString(2)
    args:= L.CheckString(3)
    argstyp:= L.CheckString(4)
    reptyp:= L.CheckString(5)
    if v, ok := ud.Value.(*rpc.Client); ok {
        typArgs := DefaultScript.GetPbType(argstyp)
        typRep := DefaultScript.GetPbType(reptyp)
        valueArgs := reflect.New(typArgs)
        valueRep := reflect.New(typRep)
        if valueRep.Interface() != nil && valueArgs.Interface() != nil {
            if value, ok := (valueArgs.Interface()).(proto.Message); ok {
                proto.Unmarshal([]byte(args), value)
                v.Call(method, &value, valueRep)
            }else {
                logger.Error("Register_lua_rpc_RpcClient_Call Error type : %v", valueArgs.Interface())
            }
        }else {
            logger.Error("Register_lua_rpc_RpcClient_Call Error : valueArgs %v, valueRep %v", valueArgs.Interface(), valueRep.Interface())
        }
    }
    return 0
}

func Register_lua_rpc_RpcClient_Close(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*rpc.Client); ok {
        v.Close()
    }
    return 0
}