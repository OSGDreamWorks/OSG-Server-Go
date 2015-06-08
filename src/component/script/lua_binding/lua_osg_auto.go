package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
)

func RegisterAllModules(L *lua.LState) int {
    logger.Debug("osg module Loader")
    Register_lua_server_RpcConn(L)
    Register_lua_server_Server(L)
    return 1
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

func (self *LuaScript) RegisterGlobalClassBegin(name string, value interface{}) {
    mt := self.state.NewTypeMetatable(name)
    self.state.SetGlobal(name, mt)
    ud := self.state.NewUserData()
    ud.Value = value
    self.state.SetMetatable(ud, self.state.GetTypeMetatable(name))
}

func (self *LuaScript) RegisterGlobalClassFunction(name string, fun string, f lua.LGFunction) {
    mt := self.state.GetTypeMetatable(name)
    self.state.SetField(mt, fun, self.state.NewFunction(f))
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
