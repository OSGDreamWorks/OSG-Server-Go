package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
)

const luaNewFunctionName = "new"
const luaIndexTypeName = "__index"

func RegisterAllModules(L *lua.LState) int {
    logger.Debug("osg module Loader")
    Register_lua_json(L)
    Register_lua_common(L)
    Register_lua_server_ProtoBufConn(L)
    Register_lua_server_Server(L)
    return 0
}


type LuaScript struct {
    state           *lua.LState
}

var DefaultScript LuaScript

func NewScript() *LuaScript {
    l := lua.NewState()
    DefaultScript = LuaScript{state : l}
    l.PreloadModule("osg", RegisterAllModules)
    return &DefaultScript
}

func (self *LuaScript) Close() {
    self.state.Close()
}

func (self *LuaScript) RegisterGlobalClassBegin(name string, value interface{}) *lua.LTable {
    mt := self.state.NewTypeMetatable(name)
    self.state.SetGlobal(name, mt)
    return mt
}

func (self *LuaScript) RegisterGlobalClassFunction(mt *lua.LTable, fun string, v lua.LValue) {
    self.state.SetField(mt, fun, v)
}

func (self *LuaScript) RegisterGlobalClassEnd(name string) {
}

func (self *LuaScript) RegisterGlobalFunction(name string, f lua.LGFunction) {
    self.state.SetGlobal(name, self.state.NewFunction(f))
}

func (self *LuaScript) ExecuteString(codes string) {
    if err := self.state.DoString(codes); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", codes, err.Error())
    }
}

func (self *LuaScript) ExecuteScriptFile(file string) {
    if err := self.state.DoFile(file); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", file, err.Error())
    }
}

func (self *LuaScript) ExecuteFunction(fn lua.LValue, n_ret int, args ...lua.LValue) (err error, ret lua.LValue) {
    err = self.state.CallByParam(lua.P{
        Fn: fn,
        NRet: n_ret,
        Protect: true,
    }, args ...)

    ret = self.state.Get(-n_ret)
    self.state.Pop(n_ret)
    return
}
