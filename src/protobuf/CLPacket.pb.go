// Code generated by protoc-gen-go.
// source: CLPacket.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CL_CheckAccount struct {
	Uid      string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Option   string `protobuf:"bytes,4,opt,name=option" json:"option,omitempty"`
	Language uint32 `protobuf:"varint,5,opt,name=language" json:"language,omitempty"`
	Udid     string `protobuf:"bytes,6,opt,name=udid" json:"udid,omitempty"`
}

func (m *CL_CheckAccount) Reset()                    { *m = CL_CheckAccount{} }
func (m *CL_CheckAccount) String() string            { return proto.CompactTextString(m) }
func (*CL_CheckAccount) ProtoMessage()               {}
func (*CL_CheckAccount) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*CL_CheckAccount)(nil), "protobuf.CL_CheckAccount")
}

var fileDescriptor1 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0xf6, 0x09, 0x48,
	0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0x0b, 0x19, 0xb9, 0xf8, 0x9d, 0x7d, 0xe2, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x1d, 0x93,
	0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x09, 0x2e, 0xf6, 0x44, 0x88, 0xa4, 0x04, 0x13, 0x58, 0x14,
	0xc6, 0x15, 0x92, 0xe2, 0xe2, 0x28, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x06,
	0x4b, 0xc1, 0xf9, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x2c, 0x60,
	0x19, 0x28, 0x0f, 0xa4, 0x27, 0x27, 0x31, 0x2f, 0xbd, 0x34, 0x31, 0x3d, 0x55, 0x82, 0x15, 0x28,
	0xc3, 0x1b, 0x04, 0xe7, 0x0b, 0x09, 0x71, 0xb1, 0x94, 0xa6, 0x00, 0x2d, 0x67, 0x03, 0xeb, 0x00,
	0xb3, 0x9d, 0x98, 0x3c, 0x18, 0x93, 0xd8, 0xc0, 0x2e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x86, 0x60, 0x2c, 0x7d, 0xca, 0x00, 0x00, 0x00,
}
