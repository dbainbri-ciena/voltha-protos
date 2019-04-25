// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_gemport_body.proto

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

type GemportsConfigData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ItfRef               string   `protobuf:"bytes,3,opt,name=itf_ref,json=itfRef,proto3" json:"itf_ref,omitempty"`
	TrafficClass         uint32   `protobuf:"varint,4,opt,name=traffic_class,json=trafficClass,proto3" json:"traffic_class,omitempty"`
	AesIndicator         bool     `protobuf:"varint,5,opt,name=aes_indicator,json=aesIndicator,proto3" json:"aes_indicator,omitempty"`
	TcontRef             string   `protobuf:"bytes,6,opt,name=tcont_ref,json=tcontRef,proto3" json:"tcont_ref,omitempty"`
	GemportId            uint32   `protobuf:"varint,7,opt,name=gemport_id,json=gemportId,proto3" json:"gemport_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GemportsConfigData) Reset()         { *m = GemportsConfigData{} }
func (m *GemportsConfigData) String() string { return proto.CompactTextString(m) }
func (*GemportsConfigData) ProtoMessage()    {}
func (*GemportsConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebce82f415fe9d29, []int{0}
}

func (m *GemportsConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GemportsConfigData.Unmarshal(m, b)
}
func (m *GemportsConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GemportsConfigData.Marshal(b, m, deterministic)
}
func (m *GemportsConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GemportsConfigData.Merge(m, src)
}
func (m *GemportsConfigData) XXX_Size() int {
	return xxx_messageInfo_GemportsConfigData.Size(m)
}
func (m *GemportsConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_GemportsConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_GemportsConfigData proto.InternalMessageInfo

func (m *GemportsConfigData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GemportsConfigData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GemportsConfigData) GetItfRef() string {
	if m != nil {
		return m.ItfRef
	}
	return ""
}

func (m *GemportsConfigData) GetTrafficClass() uint32 {
	if m != nil {
		return m.TrafficClass
	}
	return 0
}

func (m *GemportsConfigData) GetAesIndicator() bool {
	if m != nil {
		return m.AesIndicator
	}
	return false
}

func (m *GemportsConfigData) GetTcontRef() string {
	if m != nil {
		return m.TcontRef
	}
	return ""
}

func (m *GemportsConfigData) GetGemportId() uint32 {
	if m != nil {
		return m.GemportId
	}
	return 0
}

type GemportsOperData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GemportId            uint32   `protobuf:"varint,3,opt,name=gemport_id,json=gemportId,proto3" json:"gemport_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GemportsOperData) Reset()         { *m = GemportsOperData{} }
func (m *GemportsOperData) String() string { return proto.CompactTextString(m) }
func (*GemportsOperData) ProtoMessage()    {}
func (*GemportsOperData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebce82f415fe9d29, []int{1}
}

func (m *GemportsOperData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GemportsOperData.Unmarshal(m, b)
}
func (m *GemportsOperData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GemportsOperData.Marshal(b, m, deterministic)
}
func (m *GemportsOperData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GemportsOperData.Merge(m, src)
}
func (m *GemportsOperData) XXX_Size() int {
	return xxx_messageInfo_GemportsOperData.Size(m)
}
func (m *GemportsOperData) XXX_DiscardUnknown() {
	xxx_messageInfo_GemportsOperData.DiscardUnknown(m)
}

var xxx_messageInfo_GemportsOperData proto.InternalMessageInfo

func (m *GemportsOperData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GemportsOperData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GemportsOperData) GetGemportId() uint32 {
	if m != nil {
		return m.GemportId
	}
	return 0
}

func init() {
	proto.RegisterType((*GemportsConfigData)(nil), "bbf_fiber.GemportsConfigData")
	proto.RegisterType((*GemportsOperData)(nil), "bbf_fiber.GemportsOperData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_gemport_body.proto", fileDescriptor_ebce82f415fe9d29)
}

var fileDescriptor_ebce82f415fe9d29 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xfe, 0x4f, 0x68, 0x41, 0x02, 0x62, 0xa8, 0x14, 0x4a, 0xdd, 0x0c, 0x42, 0xdb,
	0x85, 0xe2, 0x03, 0xb4, 0x82, 0x74, 0x25, 0xcc, 0x52, 0x84, 0x70, 0x33, 0x49, 0xa6, 0x17, 0x3a,
	0x93, 0x92, 0x44, 0xc1, 0x87, 0xf4, 0x15, 0x5c, 0xfb, 0x04, 0xae, 0x65, 0x32, 0x1d, 0xb5, 0x82,
	0x0b, 0x77, 0xe1, 0x3b, 0x37, 0xe7, 0xe4, 0xe6, 0x90, 0xcb, 0x67, 0xb3, 0xf3, 0x5b, 0xe0, 0x7b,
	0x6b, 0xbc, 0x71, 0x4b, 0x21, 0x34, 0xd7, 0x28, 0x94, 0xe5, 0xb9, 0x2a, 0xf6, 0xc6, 0x7a, 0x2e,
	0x8c, 0x7c, 0x59, 0x04, 0x95, 0xc6, 0x5f, 0xea, 0x98, 0x1d, 0x5f, 0x2b, 0x94, 0x87, 0x7a, 0x68,
	0x9c, 0xfc, 0x65, 0xe8, 0x33, 0x53, 0xfe, 0xb4, 0x9b, 0xbd, 0x45, 0x84, 0xde, 0xd5, 0x29, 0x6e,
	0x6d, 0x4a, 0x8d, 0xf9, 0x2d, 0x78, 0xa0, 0xa7, 0xa4, 0x85, 0x92, 0x45, 0xd3, 0x28, 0x89, 0x57,
	0xdd, 0xf7, 0x8f, 0xd7, 0x49, 0x94, 0xb6, 0x50, 0x52, 0x4a, 0x3a, 0x25, 0x14, 0x8a, 0xb5, 0x2a,
	0x21, 0x0d, 0x67, 0x7a, 0x46, 0xfa, 0xe8, 0x35, 0xb7, 0x4a, 0xb3, 0x76, 0xc0, 0x3d, 0xf4, 0x3a,
	0x55, 0x9a, 0x5e, 0x90, 0x91, 0xb7, 0xa0, 0x35, 0x66, 0x3c, 0xdb, 0x81, 0x73, 0xac, 0x33, 0x8d,
	0x92, 0x51, 0x3a, 0x3c, 0xc0, 0x75, 0xc5, 0xaa, 0x21, 0x50, 0x8e, 0x63, 0x29, 0x31, 0x03, 0x6f,
	0x2c, 0xeb, 0x4e, 0xa3, 0x64, 0x90, 0x0e, 0x41, 0xb9, 0x4d, 0xc3, 0xe8, 0x39, 0x89, 0xeb, 0x87,
	0x57, 0x21, 0xbd, 0x10, 0x32, 0x08, 0xa0, 0x8a, 0x99, 0x10, 0xd2, 0x7c, 0x13, 0x4a, 0xd6, 0x0f,
	0x19, 0xf1, 0x81, 0x6c, 0xe4, 0xec, 0x91, 0x9c, 0x34, 0xfb, 0xdd, 0xef, 0x95, 0xfd, 0xef, 0x76,
	0xc7, 0xee, 0xed, 0x5f, 0xee, 0xab, 0x9b, 0x87, 0xeb, 0x1c, 0xfd, 0xf6, 0x49, 0x2c, 0x32, 0x53,
	0x2c, 0xa5, 0x00, 0x2c, 0x85, 0xc5, 0x79, 0x86, 0xaa, 0x84, 0x65, 0x5d, 0xc2, 0xfc, 0x50, 0x42,
	0x6e, 0xbe, 0x7b, 0x10, 0xbd, 0x00, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x81, 0x19,
	0x94, 0xfa, 0x01, 0x00, 0x00,
}
