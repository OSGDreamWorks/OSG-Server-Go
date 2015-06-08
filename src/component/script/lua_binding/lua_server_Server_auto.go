package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/server"
)

func Register_lua_server_Server(L *lua.LState) int {
    logger.Debug("Register_server_Server")
    svc := &server.Server{}
    DefaultScript.RegisterGlobalClassBegin("Server", svc)
    DefaultScript.RegisterGlobalClassEnd("Server")
    return 1
}