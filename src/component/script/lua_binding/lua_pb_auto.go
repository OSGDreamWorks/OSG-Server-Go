package lua_binding

import (
    "github.com/yuin/gopher-lua"
    "common/logger"
    "code.google.com/p/goprotobuf/proto"
    "encoding/binary"
    "bytes"
)

const luaPbTypeName = "pb"
const luaIOStringTypeName = "__pb_IOString"

func RegisterPbModule(L *lua.LState) int {
    logger.Debug("pb module Loader")

    mt := DefaultScript.RegisterGlobalClassBegin(luaIOStringTypeName, &IOString{})
    DefaultScript.RegisterGlobalClassFunction(mt, "__create", L.NewFunction(new_iostring))
    DefaultScript.RegisterGlobalClassFunction(mt, "__cname", lua.LString(luaIOStringTypeName))
    DefaultScript.RegisterGlobalClassFunction(mt, "__ctype", lua.LNumber(1))
    DefaultScript.RegisterGlobalClassFunction(mt, "__index", L.SetFuncs(L.NewTable(), indexIOStringMethods))
    DefaultScript.RegisterGlobalClassEnd(luaIOStringTypeName)

    // register functions to the table
    mod := L.SetFuncs(L.NewTable(), indexPbMethods)
    // returns the module
    L.Push(mod)

    return 1
}

func RegisterProtobufModule(L *lua.LState) int {
    logger.Debug("protobuf module Loader")
    DefaultScript.ExecuteScriptFile("script/protobuf/descriptor.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/text_format.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/containers.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/listener.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/type_checkers.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/wire_format.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/encoder.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/decoder.lua")
    DefaultScript.ExecuteScriptFile("script/protobuf/protobuf.lua")
    return 0
}

type IOString struct {
    size int
    buf []byte
}

var indexPbMethods = map[string]lua.LGFunction{
    "varint_encoder": varint_encoder,
    "signed_varint_encoder": signed_varint_encoder,
    "read_tag": read_tag,
    "struct_pack": struct_pack,
    "struct_unpack": struct_unpack,
    "varint_decoder": varint_decoder,
    "signed_varint_decoder": signed_varint_decoder,
    "zig_zag_decode32": zig_zag_decode32,
    "zig_zag_encode32": zig_zag_encode32,
    "zig_zag_decode64": zig_zag_decode64,
    "zig_zag_encode64": zig_zag_encode64,
    "new_iostring": new_iostring,
}

var indexIOStringMethods = map[string]lua.LGFunction{
    "__tostring": iostring_str,
    "__len": iostring_len,
    "write": iostring_write,
    "sub": iostring_sub,
    "clear": iostring_clear,
}

func sizeVarint(buf []byte) (n int) {
    for shift := uint(0); shift < 64; shift += 7 {
        if n >= len(buf) {
            return 0
        }
        b := uint64(buf[n])
        n++
        if (b & 0x80) == 0 {
            return n
        }
    }
    return 0
}

func varint_encoder(L *lua.LState) int {
    l_value := L.CheckNumber(2)
    value := uint64(l_value)

    b := proto.EncodeVarint(value)

    L.SetTop(1)
    L.Push(lua.LString(string(b)))
    L.Call(1, 0)
    return 0
}

func signed_varint_encoder(L *lua.LState) int {
    l_value := L.CheckNumber(2)
    value := int64(l_value)

    var b []byte
    b = proto.EncodeVarint(uint64(value))

    L.SetTop(1)
    L.Push(lua.LString(string(b)))
    L.Call(1, 0)
    return 0
}

func read_tag(L *lua.LState) int {

    buffer := L.CheckString(1)
    pos := L.CheckInt(2)
    b_value := []byte(buffer)

    len := sizeVarint(b_value[pos:])

    L.Push(lua.LString(string(b_value[pos:pos+len])))
    L.Push(lua.LNumber(len + pos))

    return 2
}

func struct_pack(L *lua.LState) int {

    format := byte(L.CheckInt(2))
    l_value := L.CheckNumber(3)
    b := proto.NewBuffer(nil)
    b.Reset()
    L.SetTop(1)
        switch format {
            case 'i':
                b.EncodeFixed32(uint64(int32(l_value)))
                L.Push(lua.LString(string(b.Bytes())))
            case 'q':
                b.EncodeFixed64(uint64(int64(l_value)))
                L.Push(lua.LString(string(b.Bytes())))
            case 'f':
                b.EncodeFixed32(uint64(float32(l_value)))
                L.Push(lua.LString(string(b.Bytes())))
            case 'd':
                b.EncodeFixed64(uint64(float64(l_value)))
                L.Push(lua.LString(string(b.Bytes())))
            case 'I':
                b.EncodeFixed32(uint64(uint32(l_value)))
                L.Push(lua.LString(string(b.Bytes())))
            case 'Q':
                b.EncodeFixed64(uint64(l_value))
                L.Push(lua.LString(string(b.Bytes())))
            default:
                L.Error(lua.LString("Unknown, format"), 0)
        }
    L.Call(1, 0)

    return 0
}

func struct_unpack(L *lua.LState) int {

    format := byte(L.CheckInt(1))
    l_value := L.CheckString(2)
    pos := L.CheckInt(3)

    b_value := []byte(l_value)

    b := proto.NewBuffer(b_value[pos:])

    switch format {
        case 'i':
            value,_ := b.DecodeFixed32()
            L.Push(lua.LNumber(int32(value)))
        case 'q':
            value,_ := b.DecodeFixed64()
            L.Push(lua.LNumber(int64(value)))
        case 'f':
            value,_ := b.DecodeFixed32()
            L.Push(lua.LNumber(float32(value)))
        case 'd':
            value,_ := b.DecodeFixed64()
            L.Push(lua.LNumber(float64(value)))
        case 'I':
            value,_ := b.DecodeFixed32()
            L.Push(lua.LNumber(uint32(value)))
        case 'Q':
            value,_ := b.DecodeFixed64()
            L.Push(lua.LNumber(uint64(value)))
        default:
            L.Error(lua.LString("Unknown, format"), 0)
    }

    return 1
}

func varint_decoder(L *lua.LState) int {

    buffer := L.CheckString(1)
    pos := L.CheckInt(2)

    b_value := []byte(buffer)

    value, len := proto.DecodeVarint(b_value[pos:])

    L.Push(lua.LNumber(value))
    L.Push(lua.LNumber(len + pos))

    return 2
}

func signed_varint_decoder(L *lua.LState) int {

    buffer := L.CheckString(1)
    pos := L.CheckInt(2)

    b_value := []byte(buffer)

    value, len := proto.DecodeVarint((b_value[pos:]))

    L.Push(lua.LNumber(int64(value)))
    L.Push(lua.LNumber(len + pos))

    return 2
}

func zig_zag_decode32(L *lua.LState) int {
    n := L.CheckInt(1)
    b_buf := bytes.NewBuffer([]byte{})
    binary.Write(b_buf, binary.LittleEndian, n)
    b := proto.NewBuffer(b_buf.Bytes())
    value, err := b.DecodeZigzag32()
    if err == nil {
        L.Push(lua.LNumber(value))
        return 1
    }

    return 0
}

func zig_zag_encode32(L *lua.LState) int {
    b := proto.NewBuffer(nil)
    b.Reset()
    err := b.EncodeZigzag32(uint64(L.CheckInt(1)))
    if err == nil {
        x, _ :=binary.ReadUvarint(bytes.NewBuffer(b.Bytes()))
        L.Push(lua.LNumber(x))
        return 1
    }
    return 0
}

func zig_zag_decode64(L *lua.LState) int {
    n := L.CheckNumber(1)
    b_buf := bytes.NewBuffer([]byte{})
    binary.Write(b_buf, binary.LittleEndian, n)
    b := proto.NewBuffer(b_buf.Bytes())
    value, err := b.DecodeZigzag64()
    if err == nil {
        L.Push(lua.LNumber(value))
        return 1
    }

    return 0
}

func zig_zag_encode64(L *lua.LState) int {
    b := proto.NewBuffer(nil)
    b.Reset()
    err := b.EncodeZigzag64(uint64(L.CheckNumber(1)))
    if err == nil {
        x, _ :=binary.ReadUvarint(bytes.NewBuffer(b.Bytes()))
        L.Push(lua.LNumber(x))
        return 1
    }
    return 0
}

func new_iostring(L *lua.LState) int {
    io_str := &IOString{}
    ud := L.NewUserData()
    ud.Value = io_str
    L.SetMetatable(ud, L.GetTypeMetatable(luaIOStringTypeName))
    L.Push(ud)
    return 1
}

func iostring_str(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*IOString); ok {
        L.Push(lua.LString(string(v.buf)))
    }
    return 1
}

func iostring_len(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*IOString); ok {
        L.Push(lua.LNumber(v.size))
    }
    return 1
}

func iostring_write(L *lua.LState) int {
    ud := L.CheckUserData(1)
    l_value := L.CheckString(2)
    if v, ok := ud.Value.(*IOString); ok {
        b := bytes.NewBuffer(nil)
        b.Reset()
        b.Write(v.buf)
        b.Write([]byte(l_value))
        v.buf = b.Bytes()
        v.size = b.Len()
    }
    return 0
}

func iostring_sub(L *lua.LState) int {
    ud := L.CheckUserData(1)
    begin := int(L.CheckNumber(2))
    end := int(L.CheckNumber(3))
    if v, ok := ud.Value.(*IOString); ok {
        if begin > end || end > v.size {
            L.Error(lua.LString("Out of range"), 0)
        }else {
            L.Push(lua.LString(string(v.buf[begin-1:end-begin])))
        }
    }
    return 1
}

func iostring_clear(L *lua.LState) int {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*IOString); ok {
        v.size = 0
    }
    return 0
}