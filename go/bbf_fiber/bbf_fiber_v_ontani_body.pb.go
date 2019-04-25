// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_v_ontani_body.proto

package bbf_fiber

import (
	fmt "fmt"
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

type VOntaniConfigData struct {
	ParentRef              string   `protobuf:"bytes,1,opt,name=parent_ref,json=parentRef,proto3" json:"parent_ref,omitempty"`
	ExpectedSerialNumber   string   `protobuf:"bytes,2,opt,name=expected_serial_number,json=expectedSerialNumber,proto3" json:"expected_serial_number,omitempty"`
	ExpectedRegistrationId string   `protobuf:"bytes,3,opt,name=expected_registration_id,json=expectedRegistrationId,proto3" json:"expected_registration_id,omitempty"`
	PreferredChanpair      string   `protobuf:"bytes,4,opt,name=preferred_chanpair,json=preferredChanpair,proto3" json:"preferred_chanpair,omitempty"`
	ProtectionChanpair     string   `protobuf:"bytes,5,opt,name=protection_chanpair,json=protectionChanpair,proto3" json:"protection_chanpair,omitempty"`
	UpstreamChannelSpeed   uint32   `protobuf:"varint,6,opt,name=upstream_channel_speed,json=upstreamChannelSpeed,proto3" json:"upstream_channel_speed,omitempty"`
	OnuId                  uint32   `protobuf:"varint,7,opt,name=onu_id,json=onuId,proto3" json:"onu_id,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *VOntaniConfigData) Reset()         { *m = VOntaniConfigData{} }
func (m *VOntaniConfigData) String() string { return proto.CompactTextString(m) }
func (*VOntaniConfigData) ProtoMessage()    {}
func (*VOntaniConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_166371fff1eb9424, []int{0}
}

func (m *VOntaniConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VOntaniConfigData.Unmarshal(m, b)
}
func (m *VOntaniConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VOntaniConfigData.Marshal(b, m, deterministic)
}
func (m *VOntaniConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VOntaniConfigData.Merge(m, src)
}
func (m *VOntaniConfigData) XXX_Size() int {
	return xxx_messageInfo_VOntaniConfigData.Size(m)
}
func (m *VOntaniConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_VOntaniConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_VOntaniConfigData proto.InternalMessageInfo

func (m *VOntaniConfigData) GetParentRef() string {
	if m != nil {
		return m.ParentRef
	}
	return ""
}

func (m *VOntaniConfigData) GetExpectedSerialNumber() string {
	if m != nil {
		return m.ExpectedSerialNumber
	}
	return ""
}

func (m *VOntaniConfigData) GetExpectedRegistrationId() string {
	if m != nil {
		return m.ExpectedRegistrationId
	}
	return ""
}

func (m *VOntaniConfigData) GetPreferredChanpair() string {
	if m != nil {
		return m.PreferredChanpair
	}
	return ""
}

func (m *VOntaniConfigData) GetProtectionChanpair() string {
	if m != nil {
		return m.ProtectionChanpair
	}
	return ""
}

func (m *VOntaniConfigData) GetUpstreamChannelSpeed() uint32 {
	if m != nil {
		return m.UpstreamChannelSpeed
	}
	return 0
}

func (m *VOntaniConfigData) GetOnuId() uint32 {
	if m != nil {
		return m.OnuId
	}
	return 0
}

func init() {
	proto.RegisterType((*VOntaniConfigData)(nil), "bbf_fiber.VOntaniConfigData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_v_ontani_body.proto", fileDescriptor_166371fff1eb9424)
}

var fileDescriptor_166371fff1eb9424 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd1, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x06, 0x60, 0xda, 0xef, 0x6b, 0xa5, 0x01, 0x0f, 0x8d, 0xb5, 0xec, 0x45, 0x28, 0x9e, 0x0a,
	0xd2, 0x2e, 0x62, 0x45, 0xcf, 0xd6, 0x4b, 0x2f, 0x0a, 0x5b, 0xf0, 0xe0, 0x25, 0x24, 0x9b, 0xd9,
	0x6d, 0xb0, 0x9d, 0x84, 0xd9, 0x6c, 0xd1, 0x5f, 0xe9, 0x5f, 0x92, 0xce, 0xda, 0xad, 0xd7, 0x79,
	0xde, 0x77, 0x48, 0x18, 0x71, 0xb3, 0xf7, 0xdb, 0xb8, 0xd1, 0x2a, 0x90, 0x8f, 0xbe, 0x4a, 0x8d,
	0x29, 0x54, 0xe1, 0x0c, 0x90, 0xda, 0x2b, 0x8f, 0x51, 0xa3, 0x53, 0xc6, 0xdb, 0xaf, 0x39, 0xb3,
	0x1c, 0xb4, 0x7c, 0xfd, 0xdd, 0x15, 0xc3, 0xb7, 0x57, 0x4e, 0x2c, 0x3d, 0x16, 0xae, 0x7c, 0xd6,
	0x51, 0xcb, 0x2b, 0x21, 0x82, 0x26, 0xc0, 0xa8, 0x08, 0x8a, 0xa4, 0x33, 0xe9, 0x4c, 0x07, 0xd9,
	0xa0, 0x99, 0x64, 0x50, 0xc8, 0x85, 0x18, 0xc3, 0x67, 0x80, 0x3c, 0x82, 0x55, 0x15, 0x90, 0xd3,
	0x5b, 0x85, 0xf5, 0xce, 0x00, 0x25, 0x5d, 0x8e, 0x8e, 0x8e, 0xba, 0x66, 0x7c, 0x61, 0x93, 0x8f,
	0x22, 0x69, 0x5b, 0x04, 0xa5, 0xab, 0x22, 0xe9, 0xe8, 0x3c, 0x2a, 0x67, 0x93, 0x7f, 0xdc, 0x6b,
	0xb7, 0x66, 0x7f, 0x78, 0x65, 0xe5, 0x4c, 0xc8, 0x40, 0x50, 0x00, 0x11, 0x58, 0x95, 0x6f, 0x34,
	0x06, 0xed, 0x28, 0xf9, 0xcf, 0x9d, 0x61, 0x2b, 0xcb, 0x5f, 0x90, 0xa9, 0xb8, 0x38, 0xfc, 0x13,
	0x72, 0xde, 0xde, 0xe6, 0x7b, 0x9c, 0x97, 0x27, 0x6a, 0x0b, 0x0b, 0x31, 0xae, 0x43, 0x15, 0x09,
	0xf4, 0x8e, 0xe3, 0x08, 0x5b, 0x55, 0x05, 0x00, 0x9b, 0xf4, 0x27, 0x9d, 0xe9, 0x79, 0x36, 0x3a,
	0xea, 0xb2, 0xc1, 0xf5, 0xc1, 0xe4, 0xa5, 0xe8, 0x7b, 0xac, 0x0f, 0xaf, 0x3f, 0xe3, 0x54, 0xcf,
	0x63, 0xbd, 0xb2, 0x4f, 0x0f, 0xef, 0xf7, 0xa5, 0x0f, 0x1f, 0xe5, 0xdc, 0x61, 0x6a, 0x8d, 0x76,
	0x68, 0xc8, 0xcd, 0x72, 0x07, 0xa8, 0xd3, 0xe6, 0x46, 0xb3, 0xe6, 0x46, 0xf3, 0xfd, 0x6d, 0x5a,
	0xfa, 0xd3, 0xa5, 0x4c, 0x9f, 0xe7, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x5f, 0xe5,
	0xb3, 0xcb, 0x01, 0x00, 0x00,
}
