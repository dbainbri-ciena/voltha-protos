// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_tcont_body.proto

package bbf_fiber

import (
	fmt "fmt"
	_ "github.com/dbainbri-ciena/voltha-protos/go/common"
	proto "github.com/golang/protobuf/proto"
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

type TcontsConfigData struct {
	Id                          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	InterfaceReference          string   `protobuf:"bytes,3,opt,name=interface_reference,json=interfaceReference,proto3" json:"interface_reference,omitempty"`
	TrafficDescriptorProfileRef string   `protobuf:"bytes,4,opt,name=traffic_descriptor_profile_ref,json=trafficDescriptorProfileRef,proto3" json:"traffic_descriptor_profile_ref,omitempty"`
	AllocId                     uint32   `protobuf:"varint,5,opt,name=alloc_id,json=allocId,proto3" json:"alloc_id,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *TcontsConfigData) Reset()         { *m = TcontsConfigData{} }
func (m *TcontsConfigData) String() string { return proto.CompactTextString(m) }
func (*TcontsConfigData) ProtoMessage()    {}
func (*TcontsConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_28ea482896a3c108, []int{0}
}

func (m *TcontsConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcontsConfigData.Unmarshal(m, b)
}
func (m *TcontsConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcontsConfigData.Marshal(b, m, deterministic)
}
func (m *TcontsConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcontsConfigData.Merge(m, src)
}
func (m *TcontsConfigData) XXX_Size() int {
	return xxx_messageInfo_TcontsConfigData.Size(m)
}
func (m *TcontsConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_TcontsConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_TcontsConfigData proto.InternalMessageInfo

func (m *TcontsConfigData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcontsConfigData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcontsConfigData) GetInterfaceReference() string {
	if m != nil {
		return m.InterfaceReference
	}
	return ""
}

func (m *TcontsConfigData) GetTrafficDescriptorProfileRef() string {
	if m != nil {
		return m.TrafficDescriptorProfileRef
	}
	return ""
}

func (m *TcontsConfigData) GetAllocId() uint32 {
	if m != nil {
		return m.AllocId
	}
	return 0
}

type TcontsOperData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AllocId              uint32   `protobuf:"varint,3,opt,name=alloc_id,json=allocId,proto3" json:"alloc_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcontsOperData) Reset()         { *m = TcontsOperData{} }
func (m *TcontsOperData) String() string { return proto.CompactTextString(m) }
func (*TcontsOperData) ProtoMessage()    {}
func (*TcontsOperData) Descriptor() ([]byte, []int) {
	return fileDescriptor_28ea482896a3c108, []int{1}
}

func (m *TcontsOperData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcontsOperData.Unmarshal(m, b)
}
func (m *TcontsOperData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcontsOperData.Marshal(b, m, deterministic)
}
func (m *TcontsOperData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcontsOperData.Merge(m, src)
}
func (m *TcontsOperData) XXX_Size() int {
	return xxx_messageInfo_TcontsOperData.Size(m)
}
func (m *TcontsOperData) XXX_DiscardUnknown() {
	xxx_messageInfo_TcontsOperData.DiscardUnknown(m)
}

var xxx_messageInfo_TcontsOperData proto.InternalMessageInfo

func (m *TcontsOperData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcontsOperData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcontsOperData) GetAllocId() uint32 {
	if m != nil {
		return m.AllocId
	}
	return 0
}

func init() {
	proto.RegisterType((*TcontsConfigData)(nil), "bbf_fiber.TcontsConfigData")
	proto.RegisterType((*TcontsOperData)(nil), "bbf_fiber.TcontsOperData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_tcont_body.proto", fileDescriptor_28ea482896a3c108)
}

var fileDescriptor_28ea482896a3c108 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0xd9, 0xbe, 0x3c, 0x8f, 0x0d, 0x28, 0x12, 0x11, 0x62, 0x45, 0x29, 0x3d, 0xf5, 0xd2,
	0xee, 0x41, 0xf1, 0xe6, 0xa5, 0xed, 0xc5, 0x93, 0xb2, 0x88, 0x07, 0x2f, 0x21, 0x2f, 0x93, 0x76,
	0x60, 0x37, 0x59, 0xb2, 0x51, 0xf0, 0x43, 0xea, 0xe7, 0xf0, 0x13, 0x78, 0x96, 0xcd, 0xd6, 0xb5,
	0x3d, 0x28, 0x78, 0x4b, 0xf2, 0xfb, 0xcd, 0x4c, 0x86, 0x3f, 0x99, 0x3c, 0xbb, 0x3c, 0xac, 0x05,
	0x2f, 0xbd, 0x0b, 0xae, 0x4a, 0xa5, 0x34, 0xdc, 0xa0, 0x04, 0xcf, 0x83, 0x72, 0x36, 0x70, 0xe9,
	0xf4, 0xcb, 0x2c, 0x32, 0x3a, 0x68, 0xd9, 0x90, 0xed, 0x16, 0x15, 0x10, 0x44, 0x23, 0x0d, 0xaf,
	0x7f, 0x6c, 0xe7, 0x85, 0x31, 0xa8, 0xb8, 0x86, 0x4a, 0x79, 0x2c, 0x83, 0xf3, 0xb5, 0x63, 0x30,
	0x87, 0xad, 0x19, 0xe3, 0xb7, 0x84, 0x1c, 0xde, 0xd7, 0x83, 0xab, 0x85, 0xb3, 0x06, 0x57, 0x4b,
	0x11, 0x04, 0x3d, 0x26, 0x1d, 0xd4, 0x2c, 0x19, 0x25, 0x93, 0xc1, 0xbc, 0xff, 0xfe, 0xf1, 0x7a,
	0x96, 0x64, 0x1d, 0xd4, 0x94, 0x92, 0x9e, 0x15, 0x05, 0xb0, 0x4e, 0x0d, 0xb2, 0x78, 0xa6, 0x29,
	0x39, 0x42, 0x1b, 0xc0, 0x1b, 0xa1, 0x80, 0x7b, 0x30, 0xe0, 0xc1, 0x2a, 0x60, 0xdd, 0xa8, 0xd0,
	0x16, 0x65, 0x5f, 0x84, 0x2e, 0xc8, 0xf9, 0x2f, 0x3f, 0xf3, 0x60, 0x58, 0x2f, 0xd6, 0x9e, 0x6e,
	0xac, 0x65, 0x2b, 0xdd, 0x35, 0x4e, 0x06, 0x86, 0x9e, 0x90, 0x3d, 0x91, 0xe7, 0x4e, 0x71, 0xd4,
	0xac, 0x3f, 0x4a, 0x26, 0xfb, 0xd9, 0xff, 0x78, 0xbf, 0xd1, 0xe3, 0x07, 0x72, 0xd0, 0xec, 0x73,
	0x5b, 0x82, 0xff, 0xeb, 0x36, 0xdb, 0x7d, 0xbb, 0x3b, 0x7d, 0xe7, 0x57, 0x8f, 0x97, 0x2b, 0x0c,
	0xeb, 0x27, 0x39, 0x53, 0xae, 0x48, 0xb5, 0x14, 0x68, 0xa5, 0xc7, 0xa9, 0x42, 0xb0, 0x22, 0x6d,
	0x32, 0x98, 0x6e, 0x32, 0x58, 0xb9, 0xef, 0x18, 0xe4, 0xbf, 0xf8, 0x78, 0xf1, 0x19, 0x00, 0x00,
	0xff, 0xff, 0xab, 0xc2, 0x3d, 0x58, 0xf7, 0x01, 0x00, 0x00,
}
