package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/db"
    "common/config"
    "github.com/garyburd/redigo/redis"
)

const luaCachePoolTypeName = "CachePool"

var indexCachePoolMethods = map[string]lua.LGFunction{
    "Do": Register_lua_db_CachePool_Do,
}

func Register_lua_db_CachePool(L *lua.LState) {
    logger.Debug("Register_lua_common")
    cache := db.CachePool{}
    mt := DefaultScript.RegisterGlobalClassBegin(luaCachePoolTypeName, cache)
    DefaultScript.RegisterGlobalClassFunction(mt, "new", L.NewFunction(Register_lua_db_CachePool_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(Register_lua_db_CachePool_newClass))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaCachePoolTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexCachePoolMethods))
    DefaultScript.RegisterGlobalClassEnd(luaCachePoolTypeName)
}

func Register_lua_db_CachePool_newClass(L *lua.LState) int {
    cfg := L.CheckString(2)
    var cacheCfg config.CacheConfig
    if err := config.ReadConfig(cfg, &cacheCfg); err != nil {
        logger.Fatal("load config failed, error is: %v", err)
    }
    logger.Info("Init Cache %v", cacheCfg)
    cache := db.NewCachePool(cacheCfg)
    ud := L.NewUserData()
    ud.Value = cache
    L.SetMetatable(ud, L.GetTypeMetatable(luaCachePoolTypeName))
    L.Push(ud)
    return 1
}

func Register_lua_db_CachePool_Do(L *lua.LState) int {
    ud := L.CheckUserData(1)
    cmd:= L.CheckString(2)
    arg1:= L.CheckString(3)
    var value []byte
    var err error
    if  v, ok := ud.Value.(*db.CachePool); ok {
        if L.GetTop() == 4 {
            arg2 := L.CheckString(4)
            value, err = redis.Bytes(v.Do(cmd, arg1, arg2))
        }else {
            value, err = redis.Bytes(v.Do(cmd, arg1))
        }
    }
    if err == nil {
        L.Push(lua.LString(string(value)))
        L.Push(lua.LString(""))
    }else{
        if err != nil {
            L.Push(lua.LString(""))
            L.Push(lua.LString(err.Error()))
        }else {
            L.Push(lua.LString(""))
            L.Push(lua.LString("not string type value"))
        }
        logger.Error("Register_lua_db_CachePool_Do Error : %v, %v", value, err)
    }

    return 2
}