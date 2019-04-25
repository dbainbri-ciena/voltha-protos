// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_traffic_descriptor_profile_body.proto

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

type AdditionalBwEligibilityIndicatorType int32

const (
	AdditionalBwEligibilityIndicatorType_ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NONE                AdditionalBwEligibilityIndicatorType = 0
	AdditionalBwEligibilityIndicatorType_ADDITIONAL_BW_ELIGIBILITY_INDICATOR_BEST_EFFORT_SHARING AdditionalBwEligibilityIndicatorType = 1
	AdditionalBwEligibilityIndicatorType_ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NON_ASSURED_SHARING AdditionalBwEligibilityIndicatorType = 2
)

var AdditionalBwEligibilityIndicatorType_name = map[int32]string{
	0: "ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NONE",
	1: "ADDITIONAL_BW_ELIGIBILITY_INDICATOR_BEST_EFFORT_SHARING",
	2: "ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NON_ASSURED_SHARING",
}

var AdditionalBwEligibilityIndicatorType_value = map[string]int32{
	"ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NONE":                0,
	"ADDITIONAL_BW_ELIGIBILITY_INDICATOR_BEST_EFFORT_SHARING": 1,
	"ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NON_ASSURED_SHARING": 2,
}

func (x AdditionalBwEligibilityIndicatorType) String() string {
	return proto.EnumName(AdditionalBwEligibilityIndicatorType_name, int32(x))
}

func (AdditionalBwEligibilityIndicatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_af348b9355722b7c, []int{0}
}

type TrafficDescriptorProfileData struct {
	Id                               string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                             string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FixedBandwidth                   uint64                               `protobuf:"varint,3,opt,name=fixed_bandwidth,json=fixedBandwidth,proto3" json:"fixed_bandwidth,omitempty"`
	AssuredBandwidth                 uint64                               `protobuf:"varint,4,opt,name=assured_bandwidth,json=assuredBandwidth,proto3" json:"assured_bandwidth,omitempty"`
	MaximumBandwidth                 uint64                               `protobuf:"varint,5,opt,name=maximum_bandwidth,json=maximumBandwidth,proto3" json:"maximum_bandwidth,omitempty"`
	Priority                         uint32                               `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight                           uint32                               `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
	AdditionalBwEligibilityIndicator AdditionalBwEligibilityIndicatorType `protobuf:"varint,8,opt,name=additional_bw_eligibility_indicator,json=additionalBwEligibilityIndicator,proto3,enum=bbf_fiber.AdditionalBwEligibilityIndicatorType" json:"additional_bw_eligibility_indicator,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                             `json:"-"`
	XXX_unrecognized                 []byte                               `json:"-"`
	XXX_sizecache                    int32                                `json:"-"`
}

func (m *TrafficDescriptorProfileData) Reset()         { *m = TrafficDescriptorProfileData{} }
func (m *TrafficDescriptorProfileData) String() string { return proto.CompactTextString(m) }
func (*TrafficDescriptorProfileData) ProtoMessage()    {}
func (*TrafficDescriptorProfileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_af348b9355722b7c, []int{0}
}

func (m *TrafficDescriptorProfileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficDescriptorProfileData.Unmarshal(m, b)
}
func (m *TrafficDescriptorProfileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficDescriptorProfileData.Marshal(b, m, deterministic)
}
func (m *TrafficDescriptorProfileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficDescriptorProfileData.Merge(m, src)
}
func (m *TrafficDescriptorProfileData) XXX_Size() int {
	return xxx_messageInfo_TrafficDescriptorProfileData.Size(m)
}
func (m *TrafficDescriptorProfileData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficDescriptorProfileData.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficDescriptorProfileData proto.InternalMessageInfo

func (m *TrafficDescriptorProfileData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TrafficDescriptorProfileData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TrafficDescriptorProfileData) GetFixedBandwidth() uint64 {
	if m != nil {
		return m.FixedBandwidth
	}
	return 0
}

func (m *TrafficDescriptorProfileData) GetAssuredBandwidth() uint64 {
	if m != nil {
		return m.AssuredBandwidth
	}
	return 0
}

func (m *TrafficDescriptorProfileData) GetMaximumBandwidth() uint64 {
	if m != nil {
		return m.MaximumBandwidth
	}
	return 0
}

func (m *TrafficDescriptorProfileData) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TrafficDescriptorProfileData) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TrafficDescriptorProfileData) GetAdditionalBwEligibilityIndicator() AdditionalBwEligibilityIndicatorType {
	if m != nil {
		return m.AdditionalBwEligibilityIndicator
	}
	return AdditionalBwEligibilityIndicatorType_ADDITIONAL_BW_ELIGIBILITY_INDICATOR_NONE
}

func init() {
	proto.RegisterEnum("bbf_fiber.AdditionalBwEligibilityIndicatorType", AdditionalBwEligibilityIndicatorType_name, AdditionalBwEligibilityIndicatorType_value)
	proto.RegisterType((*TrafficDescriptorProfileData)(nil), "bbf_fiber.TrafficDescriptorProfileData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_traffic_descriptor_profile_body.proto", fileDescriptor_af348b9355722b7c)
}

var fileDescriptor_af348b9355722b7c = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x4d, 0xb7, 0x5b, 0x77, 0x07, 0x5c, 0xeb, 0x80, 0x12, 0x16, 0x85, 0xa2, 0x82, 0x41,
	0xdd, 0x16, 0x54, 0xf4, 0x20, 0x1e, 0x12, 0x93, 0x5d, 0x07, 0x4a, 0x22, 0xd3, 0x88, 0xe8, 0x65,
	0x98, 0xc9, 0x4c, 0xda, 0x07, 0x49, 0xa6, 0x4c, 0xa7, 0x76, 0x7b, 0x11, 0xbf, 0xa0, 0x7e, 0x0e,
	0x3f, 0x81, 0x67, 0xd9, 0xa4, 0x9b, 0xda, 0x93, 0xbd, 0xcd, 0x7b, 0xff, 0xdf, 0xef, 0x7f, 0x18,
	0x1e, 0x7a, 0xf7, 0x4d, 0x17, 0x76, 0xc6, 0xd9, 0xdc, 0x68, 0xab, 0x17, 0x23, 0x21, 0x72, 0x96,
	0x83, 0x50, 0x86, 0x59, 0xc3, 0xf3, 0x1c, 0x32, 0x26, 0xd5, 0x22, 0x33, 0x30, 0xb7, 0xda, 0x5c,
	0x31, 0x39, 0x14, 0x8a, 0x09, 0x2d, 0xd7, 0xc3, 0x5a, 0xc0, 0xc7, 0xad, 0x70, 0xea, 0xee, 0x36,
	0x95, 0xca, 0xf2, 0x06, 0x7a, 0xf8, 0xe3, 0x00, 0xdd, 0x4f, 0x9b, 0xba, 0xb0, 0x6d, 0xfb, 0xd8,
	0x94, 0x85, 0xdc, 0x72, 0x7c, 0x17, 0x75, 0x40, 0xba, 0xce, 0xc0, 0xf1, 0x8e, 0x83, 0xc3, 0xdf,
	0x7f, 0x7e, 0x3e, 0x70, 0x68, 0x07, 0x24, 0xc6, 0xa8, 0x5b, 0xf1, 0x52, 0xb9, 0x9d, 0xab, 0x80,
	0xd6, 0x6f, 0xfc, 0x04, 0xdd, 0xce, 0xe1, 0x52, 0x49, 0x26, 0x78, 0x25, 0x57, 0x20, 0xed, 0xcc,
	0x3d, 0x18, 0x38, 0x5e, 0x97, 0x9e, 0xd4, 0xeb, 0xe0, 0x7a, 0x8b, 0x9f, 0xa1, 0x3b, 0x7c, 0xb1,
	0x58, 0x9a, 0x1d, 0xb4, 0x5b, 0xa3, 0xfd, 0x4d, 0xb0, 0x03, 0x97, 0xfc, 0x12, 0xca, 0x65, 0xf9,
	0x0f, 0x7c, 0xd8, 0xc0, 0x9b, 0x60, 0x0b, 0x9f, 0xa2, 0xa3, 0xb9, 0x01, 0x6d, 0xc0, 0xae, 0xdd,
	0xde, 0xc0, 0xf1, 0x6e, 0xd1, 0x76, 0xc6, 0xf7, 0x50, 0x6f, 0xa5, 0x60, 0x3a, 0xb3, 0xee, 0xcd,
	0x3a, 0xd9, 0x4c, 0xf8, 0x3b, 0x7a, 0xc4, 0xa5, 0x04, 0x0b, 0xba, 0xe2, 0x05, 0x13, 0x2b, 0xa6,
	0x0a, 0x98, 0x82, 0x80, 0x02, 0xec, 0x9a, 0x41, 0x25, 0x21, 0xe3, 0x56, 0x1b, 0xf7, 0x68, 0xe0,
	0x78, 0x27, 0x2f, 0x46, 0xc3, 0xf6, 0x57, 0x87, 0x7e, 0x6b, 0x05, 0xab, 0x68, 0xeb, 0x90, 0x6b,
	0x25, 0x5d, 0xcf, 0x15, 0x1d, 0xf0, 0xff, 0x50, 0x4f, 0x7f, 0x39, 0xe8, 0xf1, 0x3e, 0x55, 0xf8,
	0x39, 0xf2, 0xfc, 0x30, 0x24, 0x29, 0x49, 0x62, 0x7f, 0xcc, 0x82, 0xcf, 0x2c, 0x1a, 0x93, 0x0b,
	0x12, 0x90, 0x31, 0x49, 0xbf, 0x30, 0x12, 0x87, 0xe4, 0xbd, 0x9f, 0x26, 0x94, 0xc5, 0x49, 0x1c,
	0xf5, 0x6f, 0xe0, 0xb7, 0xe8, 0xcd, 0x3e, 0x74, 0x10, 0x4d, 0x52, 0x16, 0x9d, 0x9f, 0x27, 0x34,
	0x65, 0x93, 0x0f, 0x3e, 0x25, 0xf1, 0x45, 0xdf, 0xd9, 0x57, 0x8e, 0x93, 0x98, 0xf9, 0x93, 0xc9,
	0x27, 0x1a, 0x85, 0xad, 0xdc, 0x09, 0x5e, 0x7f, 0x7d, 0x35, 0x05, 0x3b, 0x5b, 0x8a, 0x61, 0xa6,
	0xcb, 0x91, 0x14, 0x1c, 0x2a, 0x61, 0xe0, 0x2c, 0x03, 0x55, 0xf1, 0x51, 0x73, 0x89, 0x67, 0x9b,
	0x4b, 0x9c, 0xea, 0xed, 0x59, 0x8b, 0x5e, 0xbd, 0x7c, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0x57, 0xe3, 0xc1, 0xf8, 0x02, 0x00, 0x00,
}
