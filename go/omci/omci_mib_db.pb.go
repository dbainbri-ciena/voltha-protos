// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/omci_mib_db.proto

package omci

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "gopkg.in/dbainbri-ciena/voltha-protos.v2/go/common"
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

type OpenOmciEventType_OpenOmciEventType int32

const (
	OpenOmciEventType_state_change OpenOmciEventType_OpenOmciEventType = 0
)

var OpenOmciEventType_OpenOmciEventType_name = map[int32]string{
	0: "state_change",
}

var OpenOmciEventType_OpenOmciEventType_value = map[string]int32{
	"state_change": 0,
}

func (x OpenOmciEventType_OpenOmciEventType) String() string {
	return proto.EnumName(OpenOmciEventType_OpenOmciEventType_name, int32(x))
}

func (OpenOmciEventType_OpenOmciEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{6, 0}
}

type MibAttributeData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MibAttributeData) Reset()         { *m = MibAttributeData{} }
func (m *MibAttributeData) String() string { return proto.CompactTextString(m) }
func (*MibAttributeData) ProtoMessage()    {}
func (*MibAttributeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{0}
}

func (m *MibAttributeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibAttributeData.Unmarshal(m, b)
}
func (m *MibAttributeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibAttributeData.Marshal(b, m, deterministic)
}
func (m *MibAttributeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibAttributeData.Merge(m, src)
}
func (m *MibAttributeData) XXX_Size() int {
	return xxx_messageInfo_MibAttributeData.Size(m)
}
func (m *MibAttributeData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibAttributeData.DiscardUnknown(m)
}

var xxx_messageInfo_MibAttributeData proto.InternalMessageInfo

func (m *MibAttributeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MibAttributeData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MibInstanceData struct {
	InstanceId           uint32              `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Created              string              `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Modified             string              `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	Attributes           []*MibAttributeData `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MibInstanceData) Reset()         { *m = MibInstanceData{} }
func (m *MibInstanceData) String() string { return proto.CompactTextString(m) }
func (*MibInstanceData) ProtoMessage()    {}
func (*MibInstanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{1}
}

func (m *MibInstanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibInstanceData.Unmarshal(m, b)
}
func (m *MibInstanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibInstanceData.Marshal(b, m, deterministic)
}
func (m *MibInstanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibInstanceData.Merge(m, src)
}
func (m *MibInstanceData) XXX_Size() int {
	return xxx_messageInfo_MibInstanceData.Size(m)
}
func (m *MibInstanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibInstanceData.DiscardUnknown(m)
}

var xxx_messageInfo_MibInstanceData proto.InternalMessageInfo

func (m *MibInstanceData) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *MibInstanceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *MibInstanceData) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *MibInstanceData) GetAttributes() []*MibAttributeData {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type MibClassData struct {
	ClassId              uint32             `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Instances            []*MibInstanceData `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MibClassData) Reset()         { *m = MibClassData{} }
func (m *MibClassData) String() string { return proto.CompactTextString(m) }
func (*MibClassData) ProtoMessage()    {}
func (*MibClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{2}
}

func (m *MibClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibClassData.Unmarshal(m, b)
}
func (m *MibClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibClassData.Marshal(b, m, deterministic)
}
func (m *MibClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibClassData.Merge(m, src)
}
func (m *MibClassData) XXX_Size() int {
	return xxx_messageInfo_MibClassData.Size(m)
}
func (m *MibClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibClassData.DiscardUnknown(m)
}

var xxx_messageInfo_MibClassData proto.InternalMessageInfo

func (m *MibClassData) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *MibClassData) GetInstances() []*MibInstanceData {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ManagedEntity struct {
	ClassId              uint32   `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedEntity) Reset()         { *m = ManagedEntity{} }
func (m *ManagedEntity) String() string { return proto.CompactTextString(m) }
func (*ManagedEntity) ProtoMessage()    {}
func (*ManagedEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{3}
}

func (m *ManagedEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedEntity.Unmarshal(m, b)
}
func (m *ManagedEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedEntity.Marshal(b, m, deterministic)
}
func (m *ManagedEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedEntity.Merge(m, src)
}
func (m *ManagedEntity) XXX_Size() int {
	return xxx_messageInfo_ManagedEntity.Size(m)
}
func (m *ManagedEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedEntity proto.InternalMessageInfo

func (m *ManagedEntity) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *ManagedEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MessageType struct {
	MessageType          uint32   `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageType) Reset()         { *m = MessageType{} }
func (m *MessageType) String() string { return proto.CompactTextString(m) }
func (*MessageType) ProtoMessage()    {}
func (*MessageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{4}
}

func (m *MessageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageType.Unmarshal(m, b)
}
func (m *MessageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageType.Marshal(b, m, deterministic)
}
func (m *MessageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageType.Merge(m, src)
}
func (m *MessageType) XXX_Size() int {
	return xxx_messageInfo_MessageType.Size(m)
}
func (m *MessageType) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageType.DiscardUnknown(m)
}

var xxx_messageInfo_MessageType proto.InternalMessageInfo

func (m *MessageType) GetMessageType() uint32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type MibDeviceData struct {
	DeviceId             string           `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Created              string           `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	LastSyncTime         string           `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	MibDataSync          uint32           `protobuf:"varint,4,opt,name=mib_data_sync,json=mibDataSync,proto3" json:"mib_data_sync,omitempty"`
	Version              uint32           `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Classes              []*MibClassData  `protobuf:"bytes,6,rep,name=classes,proto3" json:"classes,omitempty"`
	ManagedEntities      []*ManagedEntity `protobuf:"bytes,7,rep,name=managed_entities,json=managedEntities,proto3" json:"managed_entities,omitempty"`
	MessageTypes         []*MessageType   `protobuf:"bytes,8,rep,name=message_types,json=messageTypes,proto3" json:"message_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MibDeviceData) Reset()         { *m = MibDeviceData{} }
func (m *MibDeviceData) String() string { return proto.CompactTextString(m) }
func (*MibDeviceData) ProtoMessage()    {}
func (*MibDeviceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{5}
}

func (m *MibDeviceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MibDeviceData.Unmarshal(m, b)
}
func (m *MibDeviceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MibDeviceData.Marshal(b, m, deterministic)
}
func (m *MibDeviceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MibDeviceData.Merge(m, src)
}
func (m *MibDeviceData) XXX_Size() int {
	return xxx_messageInfo_MibDeviceData.Size(m)
}
func (m *MibDeviceData) XXX_DiscardUnknown() {
	xxx_messageInfo_MibDeviceData.DiscardUnknown(m)
}

var xxx_messageInfo_MibDeviceData proto.InternalMessageInfo

func (m *MibDeviceData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *MibDeviceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *MibDeviceData) GetLastSyncTime() string {
	if m != nil {
		return m.LastSyncTime
	}
	return ""
}

func (m *MibDeviceData) GetMibDataSync() uint32 {
	if m != nil {
		return m.MibDataSync
	}
	return 0
}

func (m *MibDeviceData) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MibDeviceData) GetClasses() []*MibClassData {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *MibDeviceData) GetManagedEntities() []*ManagedEntity {
	if m != nil {
		return m.ManagedEntities
	}
	return nil
}

func (m *MibDeviceData) GetMessageTypes() []*MessageType {
	if m != nil {
		return m.MessageTypes
	}
	return nil
}

type OpenOmciEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenOmciEventType) Reset()         { *m = OpenOmciEventType{} }
func (m *OpenOmciEventType) String() string { return proto.CompactTextString(m) }
func (*OpenOmciEventType) ProtoMessage()    {}
func (*OpenOmciEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{6}
}

func (m *OpenOmciEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenOmciEventType.Unmarshal(m, b)
}
func (m *OpenOmciEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenOmciEventType.Marshal(b, m, deterministic)
}
func (m *OpenOmciEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenOmciEventType.Merge(m, src)
}
func (m *OpenOmciEventType) XXX_Size() int {
	return xxx_messageInfo_OpenOmciEventType.Size(m)
}
func (m *OpenOmciEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenOmciEventType.DiscardUnknown(m)
}

var xxx_messageInfo_OpenOmciEventType proto.InternalMessageInfo

type OpenOmciEvent struct {
	Type                 OpenOmciEventType_OpenOmciEventType `protobuf:"varint,1,opt,name=type,proto3,enum=omci_v2.OpenOmciEventType_OpenOmciEventType" json:"type,omitempty"`
	Data                 string                              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *OpenOmciEvent) Reset()         { *m = OpenOmciEvent{} }
func (m *OpenOmciEvent) String() string { return proto.CompactTextString(m) }
func (*OpenOmciEvent) ProtoMessage()    {}
func (*OpenOmciEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa402a2df36dcc1, []int{7}
}

func (m *OpenOmciEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenOmciEvent.Unmarshal(m, b)
}
func (m *OpenOmciEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenOmciEvent.Marshal(b, m, deterministic)
}
func (m *OpenOmciEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenOmciEvent.Merge(m, src)
}
func (m *OpenOmciEvent) XXX_Size() int {
	return xxx_messageInfo_OpenOmciEvent.Size(m)
}
func (m *OpenOmciEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenOmciEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OpenOmciEvent proto.InternalMessageInfo

func (m *OpenOmciEvent) GetType() OpenOmciEventType_OpenOmciEventType {
	if m != nil {
		return m.Type
	}
	return OpenOmciEventType_state_change
}

func (m *OpenOmciEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterEnum("omci_v2.OpenOmciEventType_OpenOmciEventType", OpenOmciEventType_OpenOmciEventType_name, OpenOmciEventType_OpenOmciEventType_value)
	proto.RegisterType((*MibAttributeData)(nil), "omci_v2.MibAttributeData")
	proto.RegisterType((*MibInstanceData)(nil), "omci_v2.MibInstanceData")
	proto.RegisterType((*MibClassData)(nil), "omci_v2.MibClassData")
	proto.RegisterType((*ManagedEntity)(nil), "omci_v2.ManagedEntity")
	proto.RegisterType((*MessageType)(nil), "omci_v2.MessageType")
	proto.RegisterType((*MibDeviceData)(nil), "omci_v2.MibDeviceData")
	proto.RegisterType((*OpenOmciEventType)(nil), "omci_v2.OpenOmciEventType")
	proto.RegisterType((*OpenOmciEvent)(nil), "omci_v2.OpenOmciEvent")
}

func init() { proto.RegisterFile("voltha_protos/omci_mib_db.proto", fileDescriptor_4fa402a2df36dcc1) }

var fileDescriptor_4fa402a2df36dcc1 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0x92, 0xa6, 0x4d, 0x3a, 0x89, 0xdb, 0x7c, 0xab, 0x82, 0xb6, 0x95, 0x2a, 0x2a, 0x0b,
	0x50, 0x0f, 0x34, 0x41, 0xe1, 0x80, 0xe0, 0xd4, 0xa6, 0x2d, 0x52, 0x11, 0x51, 0x25, 0xd3, 0x13,
	0x17, 0x6b, 0x6d, 0x0f, 0xee, 0x88, 0xec, 0x3a, 0xca, 0x6e, 0x2d, 0x45, 0xe2, 0x7f, 0xf1, 0x2b,
	0xfa, 0x27, 0x38, 0x71, 0xe5, 0xd2, 0x33, 0xf2, 0xda, 0x8e, 0x1d, 0x82, 0x10, 0x37, 0xbf, 0xb7,
	0x33, 0x6f, 0x66, 0xdf, 0xb3, 0x16, 0x9e, 0xa4, 0xc9, 0xd4, 0xdc, 0x0a, 0x7f, 0x36, 0x4f, 0x4c,
	0xa2, 0x87, 0x89, 0x0c, 0xc9, 0x97, 0x14, 0xf8, 0x51, 0x30, 0xb0, 0x14, 0x6b, 0x5b, 0x2a, 0x1d,
	0x1d, 0xf0, 0xd5, 0x4a, 0x89, 0x46, 0xe4, 0x25, 0xee, 0x39, 0xf4, 0x27, 0x14, 0x9c, 0x19, 0x33,
	0xa7, 0xe0, 0xce, 0xe0, 0x85, 0x30, 0x82, 0xed, 0x43, 0x4b, 0x09, 0x89, 0xbc, 0x71, 0xd4, 0x38,
	0xde, 0x1e, 0x6f, 0xfe, 0x78, 0xb8, 0x3f, 0x6c, 0x78, 0x96, 0x62, 0x7b, 0xb0, 0x99, 0x8a, 0xe9,
	0x1d, 0xf2, 0x66, 0x76, 0xe6, 0xe5, 0xc0, 0xfd, 0xd6, 0x80, 0xdd, 0x09, 0x05, 0x57, 0x4a, 0x1b,
	0xa1, 0xc2, 0x5c, 0xe4, 0x39, 0x74, 0xa9, 0xc0, 0x3e, 0x45, 0x56, 0xcb, 0x29, 0xb5, 0xa0, 0x3c,
	0xb9, 0x8a, 0x18, 0x87, 0x76, 0x38, 0x47, 0x61, 0x30, 0x2a, 0x34, 0x4b, 0xc8, 0x0e, 0xa0, 0x23,
	0x93, 0x88, 0x3e, 0x13, 0x46, 0x7c, 0xc3, 0x1e, 0x2d, 0x31, 0x7b, 0x07, 0x20, 0xca, 0x9d, 0x35,
	0x6f, 0x1d, 0x6d, 0x1c, 0x77, 0x47, 0xfb, 0x83, 0xe2, 0xba, 0x83, 0xdf, 0x6f, 0x34, 0xee, 0x7e,
	0x7f, 0xb8, 0x3f, 0xdc, 0xca, 0xaf, 0xe5, 0xd5, 0x3a, 0xdd, 0xaf, 0xd0, 0x9b, 0x50, 0x70, 0x3e,
	0x15, 0x5a, 0xdb, 0xad, 0x8f, 0xa0, 0x13, 0x66, 0x60, 0x6d, 0xe5, 0xb6, 0xa5, 0xaf, 0x22, 0xf6,
	0x1e, 0xb6, 0xcb, 0xed, 0x35, 0x6f, 0xda, 0xc1, 0xbc, 0x3e, 0xb8, 0x6e, 0xc2, 0x98, 0x65, 0x73,
	0x9d, 0x15, 0x27, 0xbc, 0xaa, 0xdd, 0xfd, 0x00, 0xce, 0x44, 0x28, 0x11, 0x63, 0x74, 0xa9, 0x0c,
	0x99, 0xc5, 0x3f, 0x8c, 0x2f, 0xb3, 0x69, 0xae, 0x65, 0xe3, 0xbe, 0x86, 0xee, 0x04, 0xb5, 0x16,
	0x31, 0xde, 0x2c, 0x66, 0xc8, 0x8e, 0xa1, 0x27, 0x73, 0xe8, 0x9b, 0xc5, 0x0c, 0x57, 0xf5, 0xba,
	0xb2, 0xaa, 0x74, 0x7f, 0x36, 0xc1, 0x99, 0x50, 0x70, 0x81, 0x29, 0x15, 0xe1, 0xb9, 0xb0, 0x1d,
	0x59, 0x54, 0x2e, 0xb2, 0x1c, 0xd5, 0xc9, 0xf9, 0xbf, 0x06, 0xf7, 0x14, 0x76, 0xa6, 0x42, 0x1b,
	0x5f, 0x2f, 0x54, 0xe8, 0x1b, 0x92, 0x58, 0xc4, 0xd7, 0xcb, 0xd8, 0x8f, 0x0b, 0x15, 0xde, 0x90,
	0x44, 0xe6, 0x82, 0x63, 0x7f, 0x56, 0x61, 0x84, 0xad, 0xe4, 0xad, 0x6c, 0x41, 0xaf, 0x2b, 0x29,
	0xc8, 0x76, 0xc8, 0xea, 0xb2, 0x19, 0x29, 0xce, 0x35, 0x25, 0x8a, 0x6f, 0xda, 0xd3, 0x12, 0xb2,
	0x53, 0xc8, 0x2d, 0x41, 0xcd, 0xb7, 0x6c, 0x08, 0x8f, 0xea, 0x21, 0x2c, 0x03, 0x1d, 0xef, 0x66,
	0x09, 0x40, 0x65, 0xab, 0x57, 0xb6, 0xb1, 0x33, 0xe8, 0xcb, 0xdc, 0x7c, 0x1f, 0x33, 0xf7, 0x09,
	0x35, 0x6f, 0x5b, 0xa9, 0xc7, 0x95, 0x54, 0x3d, 0x1d, 0x6f, 0x57, 0xd6, 0x20, 0xa1, 0x66, 0x6f,
	0xc0, 0xa9, 0x5b, 0xac, 0x79, 0xc7, 0xf6, 0xef, 0x55, 0xfd, 0x95, 0xcb, 0x5e, 0xaf, 0x66, 0xb9,
	0x76, 0xdf, 0xc2, 0xff, 0xd7, 0x33, 0x54, 0xd7, 0x32, 0xa4, 0xcb, 0x14, 0x95, 0xb1, 0x41, 0x3c,
	0xfb, 0x03, 0xc9, 0xfa, 0xd0, 0xd3, 0x46, 0x18, 0xf4, 0xc3, 0x5b, 0xa1, 0x62, 0xec, 0xff, 0xe7,
	0x22, 0x38, 0x2b, 0x65, 0xec, 0x14, 0x5a, 0xcb, 0x88, 0x77, 0x46, 0x2f, 0x96, 0xe3, 0xd7, 0xc4,
	0xd6, 0x19, 0xcf, 0x76, 0x32, 0x06, 0xad, 0x2c, 0x88, 0x22, 0x49, 0xfb, 0x3d, 0x1e, 0x7d, 0x7a,
	0x19, 0x27, 0xb3, 0x2f, 0xf1, 0x80, 0xd4, 0x30, 0x0a, 0x04, 0xa9, 0x60, 0x4e, 0x27, 0x21, 0xa1,
	0x12, 0xc3, 0xfc, 0x39, 0x39, 0xc9, 0x9f, 0x93, 0x41, 0x3a, 0x1a, 0xc6, 0x89, 0x7d, 0x7e, 0x82,
	0x2d, 0x4b, 0xbd, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x38, 0xf1, 0x8a, 0xa2, 0x9b, 0x04, 0x00,
	0x00,
}
