package script

import (
    "component/script/js_binding"
    "component/script/lua_binding"
)

var pJsScript *js_binding.JsScript
var pLusScript *lua_binding.LuaScript

func DefaultJsScript() *js_binding.JsScript {
    if pJsScript == nil {
        pJsScript = js_binding.NewScript()
    }
    return pJsScript
}
func DefaultLuaScript() *lua_binding.LuaScript {
    if pLusScript == nil {
        pLusScript = lua_binding.NewScript()
    }
    return pLusScript
}