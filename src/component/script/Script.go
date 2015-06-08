package script

import (
    "github.com/Shopify/go-lua"
    "common/logger"
)

type Script struct {
    state           *lua.State
}

func NewScript() *Script {
    l := lua.NewState()
    lua.OpenLibraries(l)
    return &Script{state : l}
}

func (self *Script) close() {
}

func (self *Script) newTable() {
    self.state.NewTable()
}

func (self *Script) setTable(index int) {
    self.state.SetTable(index)
}

func (self *Script) TestTable(name string) {
    self.newTable()
    self.state.PushFString("ok")
    self.state.PushFString("%s", "111")
    self.setTable(-3)
    self.state.PushInteger(2)
    self.state.PushFString("%s", "222")
    self.setTable(-3)
    self.state.SetGlobal(name)
}

func (self *Script) RegisterGlobalFunction(name string, f lua.Function) {
    self.state.Register(name,f)
}

func (self *Script) ExecuteString(codes string) {
    if err := lua.DoString(self.state, codes); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", codes, err.Error())
    }
}

func (self *Script) ExecuteScriptFile(file string) {
    if err := lua.DoFile(self.state, file); err != nil {
        logger.Fatal("script: ExecuteScriptFile %s, Err : %s", file, err.Error())
    }
}