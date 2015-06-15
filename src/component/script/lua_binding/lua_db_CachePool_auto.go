package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/db"
    "common/config"
)

const luaCachePoolTypeName = "CachePool"

var indexCachePoolMethods = map[string]lua.LGFunction{
    "Call": Register_lua_rpc_RpcClient_Call,
    "Close": Register_lua_rpc_RpcClient_Close,
}

func Register_lua_db_CachePool(L *lua.LState) {
    logger.Debug("Register_lua_common")
    cache := db.CachePool{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaCachePoolTypeName, cache)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_db_CachePool_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexCachePoolMethods))
    DefaultScript.RegisterGlobalClassEnd(luaCachePoolTypeName)
}

func Register_lua_db_CachePool_newClass(L *lua.LState) int {
    cfg := L.CheckString(2)
    var cacheCfg config.CacheConfig
    if err := config.ReadConfig(cfg, &cacheCfg); err != nil {
        logger.Fatal("load config failed, error is: %v", err)
    }
    cache := db.NewCachePool(cacheCfg)
    ud := L.NewUserData()
    ud.Value = cache
    L.SetMetatable(ud, L.GetTypeMetatable(luaCachePoolTypeName))
    L.Push(ud)
    return 1
}