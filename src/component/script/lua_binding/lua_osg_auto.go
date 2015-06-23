package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "reflect"
    "protobuf"
)

const luaNewFunctionName = "new"
const luaIndexTypeName = "__index"

func RegisterOsgModule(L *lua.LState) int {
    logger.Debug("osg module Loader")
    Register_lua_json(L)
    Register_lua_common(L)
    Register_lua_db(L)
    Register_lua_db_CachePool(L)
    Register_lua_rpc_RpcClient(L)
    Register_lua_rpc_RpcServer(L)
    Register_lua_server_RpcConn(L)
    Register_lua_server_Server(L)
    return 0
}


type LuaScript struct {
    state           *lua.LState
    pbMap           map[string]reflect.Type
}

var DefaultScript LuaScript

func NewScript() *LuaScript {

    l := lua.NewState()
    DefaultScript = LuaScript{state : l, pbMap :  make(map[string]reflect.Type)}

    DefaultScript.suitablePbMap()

    l.PreloadModule("pb", RegisterPbModule)
    l.PreloadModule("protobuf", RegisterProtobufModule)
    l.PreloadModule("osg", RegisterOsgModule)

    return &DefaultScript
}

func (self *LuaScript) suitablePbMap() {
    self.pbMap["LA_CheckAccount"] = reflect.TypeOf(protobuf.LA_CheckAccount{})
    self.pbMap["AL_CheckAccountResult"] = reflect.TypeOf(protobuf.AL_CheckAccountResult{})
}

func (self *LuaScript) GetPbType(name string) reflect.Type {
    if v, ok := self.pbMap[name]; ok {
        return v
    }
    logger.Error("not regist %v in suitablePbMap...", name)
    return nil
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

func (self *LuaScript) RegisterGlobalClassField(mt *lua.LTable, fun string, v lua.LValue) {
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
