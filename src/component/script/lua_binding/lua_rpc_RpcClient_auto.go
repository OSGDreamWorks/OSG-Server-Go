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
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_rpc_RpcClient_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaRpcClientTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexRpcClientMethods))
    DefaultScript.RegisterGlobalClassEnd(luaRpcClientTypeName)
}

func Register_lua_rpc_RpcClient_newClass(L *lua.LState) int {
    addr:= L.CheckString(1)
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

    if L.GetTop() > 4 {
        argstyp:= L.CheckString(5)
        reptyp:= L.CheckString(6)
        //logger.Debug("Register_lua_rpc_RpcClient_Call(%v,%v,%v,%v,%v,%v)", ud, method, args, "", argstyp, reptyp)
        if v, ok := ud.Value.(*rpc.Client); ok {
            typArgs := DefaultScript.GetPbType(argstyp)
            typRep := DefaultScript.GetPbType(reptyp)
            valueArgs := reflect.New(typArgs)
            valueRep := reflect.New(typRep)
            if valueRep.Interface() != nil && valueArgs.Interface() != nil {
                if value, ok := (valueArgs.Interface()).(proto.Message); ok {
                    proto.Unmarshal([]byte(args), value)
                    repMsg := valueRep.Interface()
                    v.Call(method, value, repMsg)
                    rep, err := proto.Marshal(repMsg.(proto.Message))
                    if err != nil {
                        logger.Debug("Register_lua_rpc_RpcClient_Call : Marshal Error %v ", valueRep.Interface())
                        return 0
                    }
                    L.Replace(4, lua.LString(string(rep)))
                    L.Push(lua.LString(string(rep)))
                    //logger.Debug("Register_lua_rpc_RpcClient_Call (%d): rep %v ", L.GetTop(), string(rep))
                    return 1
                }else {
                    logger.Error("Register_lua_rpc_RpcClient_Call Error type : %v", valueArgs.Interface())
                }
            }else {
                logger.Error("Register_lua_rpc_RpcClient_Call Error : valueArgs %v, valueRep %v", valueArgs.Interface(), valueRep.Interface())
            }
        }
    }else {
        if v, ok := ud.Value.(*rpc.Client); ok {
            req := []byte(args)
            rep := []byte("")
            v.Call(method, &req, &rep)
            L.Replace(4, lua.LString(string(rep)))
            L.Push(lua.LString(string(rep)))
            //logger.Debug("Register_lua_rpc_RpcClient_Call (%d): rep %v ", L.GetTop(), string(rep))
            return 1
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