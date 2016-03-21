// Code generated by protoc-gen-go.
// source: PB_PacketCommon.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Packet struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Cmd            uint32 `protobuf:"varint,2,opt,name=cmd" json:"cmd,omitempty"`
	SerializedData []byte `protobuf:"bytes,3,opt,name=serialized_data,proto3" json:"serialized_data,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type RpcErrorResponse struct {
	Cmd  uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *RpcErrorResponse) Reset()                    { *m = RpcErrorResponse{} }
func (m *RpcErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*RpcErrorResponse) ProtoMessage()               {}
func (*RpcErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func init() {
	proto.RegisterType((*Packet)(nil), "protobuf.Packet")
	proto.RegisterType((*RpcErrorResponse)(nil), "protobuf.RpcErrorResponse")
}

var fileDescriptor6 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0d, 0x70, 0x8a, 0x0f,
	0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x71, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0xc1, 0x5c, 0x6c, 0x10, 0x79, 0x21, 0x3e,
	0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x20, 0x4b, 0x48, 0x80, 0x8b,
	0x39, 0x39, 0x37, 0x45, 0x82, 0x09, 0x28, 0xc0, 0x1b, 0x04, 0x62, 0x0a, 0xa9, 0x73, 0xf1, 0x17,
	0xa7, 0x16, 0x65, 0x26, 0xe6, 0x64, 0x56, 0xa5, 0xa6, 0xc4, 0xa7, 0x24, 0x96, 0x24, 0x4a, 0x30,
	0x03, 0x65, 0x79, 0x82, 0xf8, 0x10, 0xc2, 0x2e, 0x40, 0x51, 0x25, 0x0b, 0x2e, 0x81, 0xa0, 0x82,
	0x64, 0xd7, 0xa2, 0xa2, 0xfc, 0xa2, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x98, 0x71,
	0x8c, 0x08, 0xe3, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0xc0, 0x36, 0x70, 0x06, 0x81, 0xd9,
	0x4e, 0x4c, 0x1e, 0x8c, 0x49, 0x6c, 0x60, 0xc7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x0a, 0x25, 0x02, 0xbc, 0x00, 0x00, 0x00,
}
