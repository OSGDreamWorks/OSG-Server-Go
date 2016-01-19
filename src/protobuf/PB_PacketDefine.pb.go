// Code generated by protoc-gen-go.
// source: PB_PacketDefine.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	PB_PacketDefine.proto

It has these top-level messages:
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CL_Protocol int32

const (
	CL_Protocol_eCL_PacketBegin CL_Protocol = 33554432
	// ----------------------------
	CL_Protocol_eCL_Connected    CL_Protocol = 33554432
	CL_Protocol_eCL_Disconnected CL_Protocol = 33554433
	CL_Protocol_eCL_CheckAccount CL_Protocol = 33554434
	// ----------------------------
	CL_Protocol_eCL_PacketEnd CL_Protocol = 34603008
)

var CL_Protocol_name = map[int32]string{
	33554432: "eCL_PacketBegin",
	// Duplicate value: 33554432: "eCL_Connected",
	33554433: "eCL_Disconnected",
	33554434: "eCL_CheckAccount",
	34603008: "eCL_PacketEnd",
}
var CL_Protocol_value = map[string]int32{
	"eCL_PacketBegin":  33554432,
	"eCL_Connected":    33554432,
	"eCL_Disconnected": 33554433,
	"eCL_CheckAccount": 33554434,
	"eCL_PacketEnd":    34603008,
}

func (x CL_Protocol) Enum() *CL_Protocol {
	p := new(CL_Protocol)
	*p = x
	return p
}
func (x CL_Protocol) String() string {
	return proto.EnumName(CL_Protocol_name, int32(x))
}
func (x CL_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CL_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CL_Protocol_value, data, "CL_Protocol")
	if err != nil {
		return err
	}
	*x = CL_Protocol(value)
	return nil
}

type LC_Protocol int32

const (
	LC_Protocol_eLC_PacketBegin LC_Protocol = 34603008
	// ----------------------------
	LC_Protocol_eLC_Connected          LC_Protocol = 34603008
	LC_Protocol_eLC_Disconnected       LC_Protocol = 34603009
	LC_Protocol_eLC_CheckAccountResult LC_Protocol = 34603010
	// ----------------------------
	LC_Protocol_eLC_PacketEnd LC_Protocol = 35651584
)

var LC_Protocol_name = map[int32]string{
	34603008: "eLC_PacketBegin",
	// Duplicate value: 34603008: "eLC_Connected",
	34603009: "eLC_Disconnected",
	34603010: "eLC_CheckAccountResult",
	35651584: "eLC_PacketEnd",
}
var LC_Protocol_value = map[string]int32{
	"eLC_PacketBegin":        34603008,
	"eLC_Connected":          34603008,
	"eLC_Disconnected":       34603009,
	"eLC_CheckAccountResult": 34603010,
	"eLC_PacketEnd":          35651584,
}

func (x LC_Protocol) Enum() *LC_Protocol {
	p := new(LC_Protocol)
	*p = x
	return p
}
func (x LC_Protocol) String() string {
	return proto.EnumName(LC_Protocol_name, int32(x))
}
func (x LC_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LC_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LC_Protocol_value, data, "LC_Protocol")
	if err != nil {
		return err
	}
	*x = LC_Protocol(value)
	return nil
}

type CS_Protocol int32

const (
	CS_Protocol_eCS_PacketBegin CS_Protocol = 50331648
	// ----------------------------
	CS_Protocol_eCS_Connected    CS_Protocol = 50331648
	CS_Protocol_eCS_Disconnected CS_Protocol = 50331649
	CS_Protocol_eCS_CheckSession CS_Protocol = 50331650
	CS_Protocol_eCS_Ping         CS_Protocol = 50331651
	// ----------------------------
	CS_Protocol_eCS_PacketEnd CS_Protocol = 83886080
)

var CS_Protocol_name = map[int32]string{
	50331648: "eCS_PacketBegin",
	// Duplicate value: 50331648: "eCS_Connected",
	50331649: "eCS_Disconnected",
	50331650: "eCS_CheckSession",
	50331651: "eCS_Ping",
	83886080: "eCS_PacketEnd",
}
var CS_Protocol_value = map[string]int32{
	"eCS_PacketBegin":  50331648,
	"eCS_Connected":    50331648,
	"eCS_Disconnected": 50331649,
	"eCS_CheckSession": 50331650,
	"eCS_Ping":         50331651,
	"eCS_PacketEnd":    83886080,
}

func (x CS_Protocol) Enum() *CS_Protocol {
	p := new(CS_Protocol)
	*p = x
	return p
}
func (x CS_Protocol) String() string {
	return proto.EnumName(CS_Protocol_name, int32(x))
}
func (x CS_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CS_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CS_Protocol_value, data, "CS_Protocol")
	if err != nil {
		return err
	}
	*x = CS_Protocol(value)
	return nil
}

type SC_Protocol int32

const (
	SC_Protocol_eSC_PacketBegin SC_Protocol = 83886080
	// ----------------------------
	SC_Protocol_eSC_Connected          SC_Protocol = 83886080
	SC_Protocol_eSC_Disconnected       SC_Protocol = 83886081
	SC_Protocol_eSC_CheckSessionResult SC_Protocol = 83886082
	SC_Protocol_eCS_PingResult         SC_Protocol = 83886083
	// ----------------------------
	SC_Protocol_eSC_PacketEnd SC_Protocol = 117440512
)

var SC_Protocol_name = map[int32]string{
	83886080: "eSC_PacketBegin",
	// Duplicate value: 83886080: "eSC_Connected",
	83886081:  "eSC_Disconnected",
	83886082:  "eSC_CheckSessionResult",
	83886083:  "eCS_PingResult",
	117440512: "eSC_PacketEnd",
}
var SC_Protocol_value = map[string]int32{
	"eSC_PacketBegin":        83886080,
	"eSC_Connected":          83886080,
	"eSC_Disconnected":       83886081,
	"eSC_CheckSessionResult": 83886082,
	"eCS_PingResult":         83886083,
	"eSC_PacketEnd":          117440512,
}

func (x SC_Protocol) Enum() *SC_Protocol {
	p := new(SC_Protocol)
	*p = x
	return p
}
func (x SC_Protocol) String() string {
	return proto.EnumName(SC_Protocol_name, int32(x))
}
func (x SC_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SC_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SC_Protocol_value, data, "SC_Protocol")
	if err != nil {
		return err
	}
	*x = SC_Protocol(value)
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.CL_Protocol", CL_Protocol_name, CL_Protocol_value)
	proto.RegisterEnum("protobuf.LC_Protocol", LC_Protocol_name, LC_Protocol_value)
	proto.RegisterEnum("protobuf.CS_Protocol", CS_Protocol_name, CS_Protocol_value)
	proto.RegisterEnum("protobuf.SC_Protocol", SC_Protocol_name, SC_Protocol_value)
}
