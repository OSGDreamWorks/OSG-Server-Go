package script

import (
    "component/script/lua_binding"
)

var pScript *lua_binding.LuaScript

func DefaultLuaScript() *lua_binding.LuaScript {
    if pScript == nil {
        pScript = lua_binding.NewScript()
    }
    return pScript
}