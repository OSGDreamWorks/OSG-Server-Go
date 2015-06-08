package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

func Register_lua_server_RpcConn(L *lua.LState) int {
    logger.Debug("Register_server_RpcConn")
    conn := &server.Conn{}
    DefaultScript.RegisterGlobalClassBegin("RpcConn", conn)
    DefaultScript.RegisterGlobalClassEnd("RpcConn")
    return 1
}

