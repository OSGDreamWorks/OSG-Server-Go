package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "sync"
)

const luaRWMutexTypeName = "RWMutex"

var indexRWMutexMethods = map[string]lua.LGFunction{
    "Lock": Register_lua_sync_RWMutex_Lock,
    "Unlock": Register_lua_sync_RWMutex_Unlock,
}

func Register_lua_sync_RWMutex(L *lua.LState) {
    logger.Debug("Register_sync_%s", luaRWMutexTypeName)
    mutex := &sync.RWMutex{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaRWMutexTypeName, mutex)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_sync_RWMutex_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_sync_RWMutex_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaRWMutexTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexRWMutexMethods))
    DefaultScript.RegisterGlobalClassEnd(luaRWMutexTypeName)
}

func Register_lua_sync_RWMutex_newClass(L *lua.LState) int {
    mutex := sync.RWMutex{}
    ud := L.NewUserData()
    ud.Value = &mutex
    L.SetMetatable(ud, L.GetTypeMetatable(luaServerTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_sync_RWMutex_Lock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*sync.RWMutex); ok {
        v.Lock()
    }
    return 0
}

func Register_lua_sync_RWMutex_Unlock(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*sync.RWMutex); ok {
        v.Unlock()
    }
    return 0
}