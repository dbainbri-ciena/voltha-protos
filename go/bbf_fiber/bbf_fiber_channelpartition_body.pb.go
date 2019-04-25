// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_channelpartition_body.proto

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

type ChannelpartitionConfigData struct {
	ChannelgroupRef           string         `protobuf:"bytes,1,opt,name=channelgroup_ref,json=channelgroupRef,proto3" json:"channelgroup_ref,omitempty"`
	FecDownstream             bool           `protobuf:"varint,2,opt,name=fec_downstream,json=fecDownstream,proto3" json:"fec_downstream,omitempty"`
	ClosestOntDistance        uint32         `protobuf:"varint,3,opt,name=closest_ont_distance,json=closestOntDistance,proto3" json:"closest_ont_distance,omitempty"`
	DifferentialFiberDistance uint32         `protobuf:"varint,4,opt,name=differential_fiber_distance,json=differentialFiberDistance,proto3" json:"differential_fiber_distance,omitempty"`
	AuthenticationMethod      AuthMethodType `protobuf:"varint,5,opt,name=authentication_method,json=authenticationMethod,proto3,enum=bbf_fiber_types.AuthMethodType" json:"authentication_method,omitempty"`
	MulticastAesIndicator     bool           `protobuf:"varint,7,opt,name=multicast_aes_indicator,json=multicastAesIndicator,proto3" json:"multicast_aes_indicator,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}       `json:"-"`
	XXX_unrecognized          []byte         `json:"-"`
	XXX_sizecache             int32          `json:"-"`
}

func (m *ChannelpartitionConfigData) Reset()         { *m = ChannelpartitionConfigData{} }
func (m *ChannelpartitionConfigData) String() string { return proto.CompactTextString(m) }
func (*ChannelpartitionConfigData) ProtoMessage()    {}
func (*ChannelpartitionConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11291f9bcb6c1a5, []int{0}
}

func (m *ChannelpartitionConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelpartitionConfigData.Unmarshal(m, b)
}
func (m *ChannelpartitionConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelpartitionConfigData.Marshal(b, m, deterministic)
}
func (m *ChannelpartitionConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelpartitionConfigData.Merge(m, src)
}
func (m *ChannelpartitionConfigData) XXX_Size() int {
	return xxx_messageInfo_ChannelpartitionConfigData.Size(m)
}
func (m *ChannelpartitionConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelpartitionConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelpartitionConfigData proto.InternalMessageInfo

func (m *ChannelpartitionConfigData) GetChannelgroupRef() string {
	if m != nil {
		return m.ChannelgroupRef
	}
	return ""
}

func (m *ChannelpartitionConfigData) GetFecDownstream() bool {
	if m != nil {
		return m.FecDownstream
	}
	return false
}

func (m *ChannelpartitionConfigData) GetClosestOntDistance() uint32 {
	if m != nil {
		return m.ClosestOntDistance
	}
	return 0
}

func (m *ChannelpartitionConfigData) GetDifferentialFiberDistance() uint32 {
	if m != nil {
		return m.DifferentialFiberDistance
	}
	return 0
}

func (m *ChannelpartitionConfigData) GetAuthenticationMethod() AuthMethodType {
	if m != nil {
		return m.AuthenticationMethod
	}
	return AuthMethodType_SERIAL_NUMBER
}

func (m *ChannelpartitionConfigData) GetMulticastAesIndicator() bool {
	if m != nil {
		return m.MulticastAesIndicator
	}
	return false
}

func init() {
	proto.RegisterType((*ChannelpartitionConfigData)(nil), "bbf_fiber.ChannelpartitionConfigData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_channelpartition_body.proto", fileDescriptor_a11291f9bcb6c1a5)
}

var fileDescriptor_a11291f9bcb6c1a5 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0xfb, 0xe9, 0x05, 0xbb, 0x45, 0xd8, 0x54, 0x75, 0x0f, 0x15, 0x2d, 0x05, 0xf5,
	0x50, 0xd9, 0xd4, 0xd0, 0x63, 0xc1, 0xb1, 0x09, 0xe4, 0x10, 0x02, 0xc2, 0xa7, 0x5c, 0x96, 0xd5,
	0x6a, 0x56, 0x5a, 0x90, 0x76, 0xc4, 0xee, 0x28, 0xc1, 0x7f, 0x3a, 0xbf, 0x21, 0x48, 0x56, 0x14,
	0xc7, 0x90, 0xeb, 0xfb, 0x3c, 0xef, 0x2e, 0x33, 0xc3, 0xd6, 0x77, 0x58, 0x52, 0x21, 0x78, 0x6d,
	0x91, 0xd0, 0x2d, 0xd3, 0x54, 0x71, 0xa5, 0x53, 0xb0, 0x5c, 0x16, 0xc2, 0x18, 0x28, 0x6b, 0x61,
	0x49, 0x93, 0x46, 0xc3, 0x53, 0xcc, 0x0e, 0x71, 0xa7, 0xf9, 0xe3, 0x41, 0x5b, 0xfc, 0x7c, 0xad,
	0x4f, 0x87, 0x1a, 0xdc, 0xd1, 0xff, 0xf1, 0x30, 0x62, 0x8b, 0xed, 0xd9, 0x7b, 0x5b, 0x34, 0x4a,
	0xe7, 0x3b, 0x41, 0xc2, 0xff, 0xcd, 0x3e, 0xf7, 0xbf, 0xe5, 0x16, 0x9b, 0x9a, 0x5b, 0x50, 0x81,
	0x17, 0x7a, 0xd1, 0x38, 0xf9, 0x74, 0x9a, 0x27, 0xa0, 0xfc, 0x5f, 0x6c, 0xaa, 0x40, 0xf2, 0x0c,
	0xef, 0x8d, 0x23, 0x0b, 0xa2, 0x0a, 0x46, 0xa1, 0x17, 0x7d, 0x4c, 0x26, 0x0a, 0xe4, 0x6e, 0x08,
	0xfd, 0x15, 0x9b, 0xc9, 0x12, 0x1d, 0x38, 0xe2, 0x68, 0x88, 0x67, 0xda, 0x91, 0x30, 0x12, 0x82,
	0x37, 0xa1, 0x17, 0x4d, 0x12, 0xbf, 0x67, 0x37, 0x86, 0x76, 0x3d, 0xf1, 0xff, 0xb3, 0x6f, 0x99,
	0x56, 0x0a, 0x2c, 0x18, 0xd2, 0xa2, 0xec, 0x87, 0x18, 0x8a, 0x6f, 0xbb, 0xe2, 0xd7, 0x53, 0xe5,
	0xb2, 0x35, 0x86, 0xfe, 0x9e, 0xcd, 0x45, 0x43, 0x45, 0x8b, 0xa4, 0xe8, 0xf6, 0x55, 0x01, 0x15,
	0x98, 0x05, 0xef, 0x42, 0x2f, 0x9a, 0xfe, 0xfd, 0x1e, 0x9f, 0x6f, 0x66, 0xd3, 0x50, 0x71, 0xdd,
	0x29, 0xfb, 0x43, 0x0d, 0xc9, 0xec, 0x65, 0xfb, 0x48, 0xfc, 0x7f, 0xec, 0x4b, 0xd5, 0x94, 0x6d,
	0xe8, 0x88, 0x0b, 0x70, 0x5c, 0x9b, 0xac, 0x35, 0xd0, 0x06, 0x1f, 0xba, 0xb9, 0xe7, 0x03, 0xde,
	0x80, 0xbb, 0x7a, 0x82, 0x17, 0xab, 0xdb, 0x38, 0xd7, 0x54, 0x34, 0x69, 0x2c, 0xb1, 0x5a, 0x62,
	0x0d, 0x46, 0xa2, 0xcd, 0x96, 0xc7, 0x5b, 0xfd, 0xe9, 0x6f, 0x95, 0xe3, 0xf3, 0xb9, 0xd2, 0xf7,
	0x5d, 0xb8, 0x7e, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xe6, 0xaa, 0x16, 0x10, 0x02, 0x00, 0x00,
}
