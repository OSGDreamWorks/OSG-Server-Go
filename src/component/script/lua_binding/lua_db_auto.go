package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "component/db"
)

const luaDBInitFuncName = "DBInit"
const luaDBQueryFuncName = "DBQuery"
const luaDBWriteFuncName = "DBWrite"
const luaDBDeleteFuncName = "DBDelete"

func Register_lua_db(L *lua.LState) {
    logger.Debug("Register_lua_db")
    DefaultScript.RegisterGlobalFunction(luaDBInitFuncName, Register_lua_db_DBInit)
    DefaultScript.RegisterGlobalFunction(luaDBQueryFuncName, Register_lua_db_DBQuery)
    DefaultScript.RegisterGlobalFunction(luaDBWriteFuncName, Register_lua_db_DBWrite)
    DefaultScript.RegisterGlobalFunction(luaDBDeleteFuncName, Register_lua_db_DBDelete)
}

func Register_lua_db_DBInit(L *lua.LState) int {
    db.Init()
    return 0
}

func Register_lua_db_DBQuery(L *lua.LState) int {
    tablename := L.CheckString(1)
    key := L.CheckString(2)
    buf := []byte("")
    exist, err := db.QueryBinary(tablename, key, &buf)
    if err == nil {
        L.Push(lua.LString(string(buf)))
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(""))
    }else{
        logger.Debug("DBQuery Error %v, %v, %v", buf, exist, err)
        L.Push(lua.LString(""))
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(err.Error()))
    }
    return 3//value exist err
}

func Register_lua_db_DBWrite(L *lua.LState) int {
    tablename := L.CheckString(1)
    key := L.CheckString(2)
    value := L.CheckString(3)
    exist, err := db.WriteBinary(tablename, key, []byte(value))
    if err == nil {
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(""))
    }else{
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(err.Error()))
    }
    return 2//result err
}

func Register_lua_db_DBDelete(L *lua.LState) int {
    tablename := L.CheckString(1)
    key := L.CheckString(2)
    exist, err := db.DeleteBinary(tablename, key)
    if err == nil {
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(""))
    }else{
        L.Push(lua.LBool(exist))
        L.Push(lua.LString(err.Error()))
    }
    return 2//result err
}