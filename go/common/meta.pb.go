// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/meta.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Access int32

const (
	// read-write, stored attribute
	Access_CONFIG Access = 0
	// read-only field, stored with the model, covered by its hash
	Access_READ_ONLY Access = 1
	// A read-only attribute that is not stored in the model, not covered
	// by its hash, its value is filled real-time upon each request.
	Access_REAL_TIME Access = 2
)

var Access_name = map[int32]string{
	0: "CONFIG",
	1: "READ_ONLY",
	2: "REAL_TIME",
}

var Access_value = map[string]int32{
	"CONFIG":    0,
	"READ_ONLY": 1,
	"REAL_TIME": 2,
}

func (x Access) String() string {
	return proto.EnumName(Access_name, int32(x))
}

func (Access) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96b320e8a67781f3, []int{0}
}

type ChildNode struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChildNode) Reset()         { *m = ChildNode{} }
func (m *ChildNode) String() string { return proto.CompactTextString(m) }
func (*ChildNode) ProtoMessage()    {}
func (*ChildNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b320e8a67781f3, []int{0}
}

func (m *ChildNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChildNode.Unmarshal(m, b)
}
func (m *ChildNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChildNode.Marshal(b, m, deterministic)
}
func (m *ChildNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildNode.Merge(m, src)
}
func (m *ChildNode) XXX_Size() int {
	return xxx_messageInfo_ChildNode.Size(m)
}
func (m *ChildNode) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildNode.DiscardUnknown(m)
}

var xxx_messageInfo_ChildNode proto.InternalMessageInfo

func (m *ChildNode) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

var E_ChildNode = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ChildNode)(nil),
	Field:         7761772,
	Name:          "voltha.child_node",
	Tag:           "bytes,7761772,opt,name=child_node",
	Filename:      "voltha_protos/meta.proto",
}

var E_Access = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Access)(nil),
	Field:         7761773,
	Name:          "voltha.access",
	Tag:           "varint,7761773,opt,name=access,enum=voltha.Access",
	Filename:      "voltha_protos/meta.proto",
}

func init() {
	proto.RegisterEnum("voltha.Access", Access_name, Access_value)
	proto.RegisterType((*ChildNode)(nil), "voltha.ChildNode")
	proto.RegisterExtension(E_ChildNode)
	proto.RegisterExtension(E_Access)
}

func init() { proto.RegisterFile("voltha_protos/meta.proto", fileDescriptor_96b320e8a67781f3) }

var fileDescriptor_96b320e8a67781f3 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdf, 0xbe, 0x83, 0x42, 0xff, 0xe2, 0xa8, 0x3d, 0x15, 0x61, 0x50, 0x3c, 0x0d, 0x61,
	0x09, 0x56, 0x4f, 0xbb, 0xcd, 0xb9, 0xe9, 0x60, 0xb6, 0x50, 0xbc, 0xe8, 0xa5, 0xb4, 0x69, 0xec,
	0xc2, 0xda, 0xfc, 0x4b, 0x13, 0x07, 0x7e, 0x54, 0x2f, 0x7e, 0x02, 0xfd, 0x0e, 0xd2, 0x66, 0xf5,
	0xea, 0xed, 0xc9, 0x93, 0x3c, 0x3f, 0x7e, 0x04, 0xfc, 0x03, 0x56, 0x7a, 0x97, 0xa5, 0x4d, 0x8b,
	0x1a, 0x15, 0xad, 0xb9, 0xce, 0x48, 0x9f, 0x3d, 0xdb, 0xdc, 0x9c, 0x07, 0x25, 0x62, 0x59, 0x71,
	0xda, 0xb7, 0xf9, 0xdb, 0x2b, 0x2d, 0xb8, 0x62, 0xad, 0x68, 0x34, 0xb6, 0xe6, 0xe5, 0xc5, 0x04,
	0x9c, 0xe5, 0x4e, 0x54, 0x45, 0x84, 0x05, 0xf7, 0x5c, 0x18, 0xed, 0xf9, 0xbb, 0x6f, 0x05, 0xd6,
	0xd4, 0x49, 0xba, 0x78, 0x19, 0x82, 0xbd, 0x60, 0x8c, 0x2b, 0xe5, 0x01, 0xd8, 0xcb, 0x38, 0x5a,
	0x6f, 0xee, 0xdd, 0x7f, 0xde, 0x29, 0x38, 0xc9, 0x6a, 0x71, 0x97, 0xc6, 0xd1, 0xf6, 0xd9, 0xb5,
	0x8e, 0xc7, 0x6d, 0xfa, 0xb4, 0x79, 0x5c, 0xb9, 0xff, 0xe7, 0x09, 0x00, 0xeb, 0x90, 0xa9, 0xec,
	0x98, 0x13, 0x62, 0x1c, 0xc8, 0xe0, 0x40, 0xd6, 0x82, 0x57, 0x45, 0xdc, 0x68, 0x81, 0x52, 0xf9,
	0x5f, 0x9f, 0x1f, 0xa3, 0xc0, 0x9a, 0x9e, 0x84, 0x67, 0xc4, 0x38, 0x93, 0x5f, 0x9d, 0xc4, 0x61,
	0x43, 0x9c, 0x3f, 0x80, 0x9d, 0x19, 0x8f, 0x3f, 0x78, 0xdf, 0x86, 0x37, 0x0e, 0xc7, 0x03, 0xcf,
	0xf8, 0x27, 0xc7, 0xfd, 0xed, 0xcd, 0x4b, 0x58, 0x62, 0xb3, 0x2f, 0x89, 0x90, 0xb4, 0xc8, 0x33,
	0x21, 0xf3, 0x56, 0xcc, 0x98, 0xe0, 0x32, 0xa3, 0x66, 0x30, 0x33, 0xdf, 0x49, 0x0e, 0x57, 0xb4,
	0x44, 0xca, 0xb0, 0xae, 0x51, 0xe6, 0x76, 0x5f, 0x5e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0x05, 0xa6, 0x4f, 0x73, 0x01, 0x00, 0x00,
}
