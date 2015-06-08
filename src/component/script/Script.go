package script

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
)

type Script struct {
    state           *lua.LState
}

func NewScript() *Script {
    return &Script{state : lua.NewState()}
}

func (self *Script) Close() {
    self.state.Close()
}

func (self *Script) RegisterGlobalClassBegin(name string, value interface{}) {
    mt := self.state.NewTypeMetatable(name)
    self.state.SetGlobal(name, mt)
    ud := self.state.NewUserData()
    ud.Value = value
    self.state.SetMetatable(ud, self.state.GetTypeMetatable(name))
}

func (self *Script) RegisterGlobalClassFunction(name string, fun string, f lua.LGFunction) {
    mt := self.state.GetTypeMetatable(name)
    self.state.SetField(mt, fun, self.state.NewFunction(f))
}

func (self *Script) RegisterGlobalClassEnd(name string, value interface{}) {
}

func (self *Script) RegisterGlobalFunction(name string, f lua.LGFunction) {
    self.state.SetGlobal(name, self.state.NewFunction(f))
}

func (self *Script) ExecuteString(codes string) {
    if err := self.state.DoString(codes); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", codes, err.Error())
    }
}

func (self *Script) ExecuteScriptFile(file string) {
    if err := self.state.DoFile(file); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", file, err.Error())
    }
}