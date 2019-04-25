// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_multicast_distribution_set_body.proto

package bbf_fiber

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "gopkg.in/dbainbri-ciena/voltha-protos.v1/go/common"
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

type AllMulticastVlans int32

const (
	AllMulticastVlans_ALL_MULTICAST_VLANS AllMulticastVlans = 0
)

var AllMulticastVlans_name = map[int32]string{
	0: "ALL_MULTICAST_VLANS",
}

var AllMulticastVlans_value = map[string]int32{
	"ALL_MULTICAST_VLANS": 0,
}

func (x AllMulticastVlans) String() string {
	return proto.EnumName(AllMulticastVlans_name, int32(x))
}

func (AllMulticastVlans) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7c15927c509ff5c, []int{0}
}

type VlanList struct {
	MulticastVlanId      []uint32 `protobuf:"varint,1,rep,packed,name=multicast_vlan_id,json=multicastVlanId,proto3" json:"multicast_vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlanList) Reset()         { *m = VlanList{} }
func (m *VlanList) String() string { return proto.CompactTextString(m) }
func (*VlanList) ProtoMessage()    {}
func (*VlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7c15927c509ff5c, []int{0}
}

func (m *VlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanList.Unmarshal(m, b)
}
func (m *VlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanList.Marshal(b, m, deterministic)
}
func (m *VlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanList.Merge(m, src)
}
func (m *VlanList) XXX_Size() int {
	return xxx_messageInfo_VlanList.Size(m)
}
func (m *VlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanList.DiscardUnknown(m)
}

var xxx_messageInfo_VlanList proto.InternalMessageInfo

func (m *VlanList) GetMulticastVlanId() []uint32 {
	if m != nil {
		return m.MulticastVlanId
	}
	return nil
}

type MulticastDistributionSetData struct {
	Id                  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MulticastGemportRef string `protobuf:"bytes,3,opt,name=multicast_gemport_ref,json=multicastGemportRef,proto3" json:"multicast_gemport_ref,omitempty"`
	// Types that are valid to be assigned to MulticastVlans:
	//	*MulticastDistributionSetData_AllMulticastVlans
	//	*MulticastDistributionSetData_VlanList
	MulticastVlans       isMulticastDistributionSetData_MulticastVlans `protobuf_oneof:"multicast_vlans"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *MulticastDistributionSetData) Reset()         { *m = MulticastDistributionSetData{} }
func (m *MulticastDistributionSetData) String() string { return proto.CompactTextString(m) }
func (*MulticastDistributionSetData) ProtoMessage()    {}
func (*MulticastDistributionSetData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7c15927c509ff5c, []int{1}
}

func (m *MulticastDistributionSetData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MulticastDistributionSetData.Unmarshal(m, b)
}
func (m *MulticastDistributionSetData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MulticastDistributionSetData.Marshal(b, m, deterministic)
}
func (m *MulticastDistributionSetData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulticastDistributionSetData.Merge(m, src)
}
func (m *MulticastDistributionSetData) XXX_Size() int {
	return xxx_messageInfo_MulticastDistributionSetData.Size(m)
}
func (m *MulticastDistributionSetData) XXX_DiscardUnknown() {
	xxx_messageInfo_MulticastDistributionSetData.DiscardUnknown(m)
}

var xxx_messageInfo_MulticastDistributionSetData proto.InternalMessageInfo

func (m *MulticastDistributionSetData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MulticastDistributionSetData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MulticastDistributionSetData) GetMulticastGemportRef() string {
	if m != nil {
		return m.MulticastGemportRef
	}
	return ""
}

type isMulticastDistributionSetData_MulticastVlans interface {
	isMulticastDistributionSetData_MulticastVlans()
}

type MulticastDistributionSetData_AllMulticastVlans struct {
	AllMulticastVlans AllMulticastVlans `protobuf:"varint,4,opt,name=all_multicast_vlans,json=allMulticastVlans,proto3,enum=bbf_fiber.AllMulticastVlans,oneof"`
}

type MulticastDistributionSetData_VlanList struct {
	VlanList *VlanList `protobuf:"bytes,5,opt,name=vlan_list,json=vlanList,proto3,oneof"`
}

func (*MulticastDistributionSetData_AllMulticastVlans) isMulticastDistributionSetData_MulticastVlans() {
}

func (*MulticastDistributionSetData_VlanList) isMulticastDistributionSetData_MulticastVlans() {}

func (m *MulticastDistributionSetData) GetMulticastVlans() isMulticastDistributionSetData_MulticastVlans {
	if m != nil {
		return m.MulticastVlans
	}
	return nil
}

func (m *MulticastDistributionSetData) GetAllMulticastVlans() AllMulticastVlans {
	if x, ok := m.GetMulticastVlans().(*MulticastDistributionSetData_AllMulticastVlans); ok {
		return x.AllMulticastVlans
	}
	return AllMulticastVlans_ALL_MULTICAST_VLANS
}

func (m *MulticastDistributionSetData) GetVlanList() *VlanList {
	if x, ok := m.GetMulticastVlans().(*MulticastDistributionSetData_VlanList); ok {
		return x.VlanList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MulticastDistributionSetData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MulticastDistributionSetData_AllMulticastVlans)(nil),
		(*MulticastDistributionSetData_VlanList)(nil),
	}
}

func init() {
	proto.RegisterEnum("bbf_fiber.AllMulticastVlans", AllMulticastVlans_name, AllMulticastVlans_value)
	proto.RegisterType((*VlanList)(nil), "bbf_fiber.VlanList")
	proto.RegisterType((*MulticastDistributionSetData)(nil), "bbf_fiber.MulticastDistributionSetData")
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_multicast_distribution_set_body.proto", fileDescriptor_a7c15927c509ff5c)
}

var fileDescriptor_a7c15927c509ff5c = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0xae, 0xd2, 0x40,
	0x18, 0xa5, 0xbd, 0xf7, 0x1a, 0x18, 0xa3, 0xd8, 0x69, 0x88, 0x0d, 0xc1, 0x84, 0xb0, 0x6a, 0x88,
	0xb4, 0xb1, 0xfe, 0xad, 0x5c, 0x14, 0x49, 0x84, 0xa4, 0xb0, 0x28, 0xc8, 0xc2, 0xcd, 0x64, 0x86,
	0x4e, 0xeb, 0xc4, 0x69, 0x87, 0x74, 0x86, 0x26, 0xbe, 0x87, 0xaf, 0xe5, 0x7b, 0xf8, 0x04, 0xae,
	0x0d, 0xad, 0x94, 0x02, 0x8b, 0xbb, 0x9b, 0x7c, 0xe7, 0x7c, 0xe7, 0x9c, 0x6f, 0x0e, 0xf8, 0x54,
	0x08, 0xae, 0xbe, 0x63, 0xb4, 0xcf, 0x85, 0x12, 0xd2, 0x25, 0x24, 0x46, 0x31, 0x23, 0x34, 0x47,
	0xe9, 0x81, 0x2b, 0xb6, 0xc3, 0x52, 0xa1, 0x88, 0x49, 0x95, 0x33, 0x72, 0x50, 0x4c, 0x64, 0x48,
	0x52, 0x85, 0x88, 0x88, 0x7e, 0x3a, 0xe5, 0x02, 0xec, 0xd4, 0x0b, 0x7d, 0xeb, 0x52, 0x29, 0xa5,
	0x0a, 0x57, 0xa4, 0xfe, 0xbb, 0xc7, 0x3d, 0x12, 0x9a, 0xee, 0x45, 0xde, 0x94, 0x1e, 0x7d, 0x00,
	0xed, 0x2d, 0xc7, 0x59, 0xc0, 0xa4, 0x82, 0x63, 0x60, 0x9c, 0xb9, 0x05, 0xc7, 0x19, 0x62, 0x91,
	0xa5, 0x0d, 0xef, 0xec, 0x67, 0x61, 0xb7, 0x06, 0x8e, 0xec, 0x45, 0x34, 0xfa, 0xa5, 0x83, 0xc1,
	0xf2, 0x34, 0x9b, 0x35, 0xb2, 0xaf, 0xa9, 0x9a, 0x61, 0x85, 0x61, 0x0f, 0xe8, 0xe5, 0xb6, 0x66,
	0x77, 0xa6, 0x0f, 0x7f, 0xfe, 0xfe, 0x7e, 0xa5, 0x85, 0x3a, 0x8b, 0x20, 0x04, 0xf7, 0x19, 0x4e,
	0xa9, 0xa5, 0x1f, 0x81, 0xb0, 0x7c, 0x43, 0x0f, 0xf4, 0x6e, 0x33, 0xe6, 0x34, 0xb6, 0xee, 0x4a,
	0x92, 0x59, 0x83, 0x5f, 0x2a, 0x2c, 0xa4, 0x31, 0x5c, 0x01, 0x13, 0x73, 0x8e, 0x2e, 0xf3, 0x4a,
	0xeb, 0x7e, 0xa8, 0xd9, 0xcf, 0xbd, 0x81, 0x53, 0x5f, 0xef, 0xf8, 0x9c, 0x2f, 0x9b, 0xd9, 0xe5,
	0xbc, 0x15, 0x1a, 0xf8, 0x7a, 0x08, 0x3d, 0xd0, 0x29, 0x2f, 0xe6, 0x4c, 0x2a, 0xeb, 0x61, 0xa8,
	0xd9, 0x4f, 0x3d, 0xb3, 0xa1, 0x72, 0xfa, 0xa3, 0x79, 0x2b, 0x6c, 0x17, 0xff, 0xdf, 0x53, 0x03,
	0x74, 0xaf, 0xfc, 0xc7, 0xaf, 0x81, 0x71, 0x63, 0x08, 0x5f, 0x02, 0xd3, 0x0f, 0x02, 0xb4, 0xfc,
	0x1a, 0x6c, 0x16, 0x9f, 0xfd, 0xf5, 0x06, 0x6d, 0x03, 0x7f, 0xb5, 0x7e, 0xd1, 0x9a, 0x7e, 0xfc,
	0xf6, 0x3e, 0x11, 0xfb, 0x1f, 0x89, 0xc3, 0x32, 0x37, 0x22, 0x98, 0x65, 0x24, 0x67, 0x93, 0x1d,
	0xa3, 0x19, 0x76, 0xab, 0x32, 0x27, 0x55, 0x99, 0x4e, 0xf1, 0xc6, 0x4d, 0xc4, 0xb9, 0x52, 0xf2,
	0xa4, 0x9c, 0xbf, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x68, 0x4e, 0x7c, 0x58, 0x02, 0x00,
	0x00,
}
