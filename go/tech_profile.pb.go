// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/tech_profile.proto

package tech_profile

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Direction int32

const (
	Direction_UPSTREAM      Direction = 0
	Direction_DOWNSTREAM    Direction = 1
	Direction_BIDIRECTIONAL Direction = 2
)

var Direction_name = map[int32]string{
	0: "UPSTREAM",
	1: "DOWNSTREAM",
	2: "BIDIRECTIONAL",
}

var Direction_value = map[string]int32{
	"UPSTREAM":      0,
	"DOWNSTREAM":    1,
	"BIDIRECTIONAL": 2,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{0}
}

type SchedulingPolicy int32

const (
	SchedulingPolicy_WRR            SchedulingPolicy = 0
	SchedulingPolicy_StrictPriority SchedulingPolicy = 1
	SchedulingPolicy_Hybrid         SchedulingPolicy = 2
)

var SchedulingPolicy_name = map[int32]string{
	0: "WRR",
	1: "StrictPriority",
	2: "Hybrid",
}

var SchedulingPolicy_value = map[string]int32{
	"WRR":            0,
	"StrictPriority": 1,
	"Hybrid":         2,
}

func (x SchedulingPolicy) String() string {
	return proto.EnumName(SchedulingPolicy_name, int32(x))
}

func (SchedulingPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{1}
}

type AdditionalBW int32

const (
	AdditionalBW_AdditionalBW_None       AdditionalBW = 0
	AdditionalBW_AdditionalBW_NA         AdditionalBW = 1
	AdditionalBW_AdditionalBW_BestEffort AdditionalBW = 2
	AdditionalBW_AdditionalBW_Auto       AdditionalBW = 3
)

var AdditionalBW_name = map[int32]string{
	0: "AdditionalBW_None",
	1: "AdditionalBW_NA",
	2: "AdditionalBW_BestEffort",
	3: "AdditionalBW_Auto",
}

var AdditionalBW_value = map[string]int32{
	"AdditionalBW_None":       0,
	"AdditionalBW_NA":         1,
	"AdditionalBW_BestEffort": 2,
	"AdditionalBW_Auto":       3,
}

func (x AdditionalBW) String() string {
	return proto.EnumName(AdditionalBW_name, int32(x))
}

func (AdditionalBW) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{2}
}

type DiscardPolicy int32

const (
	DiscardPolicy_TailDrop  DiscardPolicy = 0
	DiscardPolicy_WTailDrop DiscardPolicy = 1
	DiscardPolicy_Red       DiscardPolicy = 2
	DiscardPolicy_WRed      DiscardPolicy = 3
)

var DiscardPolicy_name = map[int32]string{
	0: "TailDrop",
	1: "WTailDrop",
	2: "Red",
	3: "WRed",
}

var DiscardPolicy_value = map[string]int32{
	"TailDrop":  0,
	"WTailDrop": 1,
	"Red":       2,
	"WRed":      3,
}

func (x DiscardPolicy) String() string {
	return proto.EnumName(DiscardPolicy_name, int32(x))
}

func (DiscardPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{3}
}

type InferredAdditionBWIndication int32

const (
	InferredAdditionBWIndication_InferredAdditionBWIndication_None       InferredAdditionBWIndication = 0
	InferredAdditionBWIndication_InferredAdditionBWIndication_Assured    InferredAdditionBWIndication = 1
	InferredAdditionBWIndication_InferredAdditionBWIndication_BestEffort InferredAdditionBWIndication = 2
)

var InferredAdditionBWIndication_name = map[int32]string{
	0: "InferredAdditionBWIndication_None",
	1: "InferredAdditionBWIndication_Assured",
	2: "InferredAdditionBWIndication_BestEffort",
}

var InferredAdditionBWIndication_value = map[string]int32{
	"InferredAdditionBWIndication_None":       0,
	"InferredAdditionBWIndication_Assured":    1,
	"InferredAdditionBWIndication_BestEffort": 2,
}

func (x InferredAdditionBWIndication) String() string {
	return proto.EnumName(InferredAdditionBWIndication_name, int32(x))
}

func (InferredAdditionBWIndication) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{4}
}

type SchedulerConfig struct {
	Direction            Direction        `protobuf:"varint,1,opt,name=direction,proto3,enum=tech_profile.Direction" json:"direction,omitempty"`
	AdditionalBw         AdditionalBW     `protobuf:"varint,2,opt,name=additional_bw,json=additionalBw,proto3,enum=tech_profile.AdditionalBW" json:"additional_bw,omitempty"`
	Priority             uint32           `protobuf:"fixed32,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight               uint32           `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	SchedPolicy          SchedulingPolicy `protobuf:"varint,5,opt,name=sched_policy,json=schedPolicy,proto3,enum=tech_profile.SchedulingPolicy" json:"sched_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SchedulerConfig) Reset()         { *m = SchedulerConfig{} }
func (m *SchedulerConfig) String() string { return proto.CompactTextString(m) }
func (*SchedulerConfig) ProtoMessage()    {}
func (*SchedulerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{0}
}

func (m *SchedulerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerConfig.Unmarshal(m, b)
}
func (m *SchedulerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerConfig.Marshal(b, m, deterministic)
}
func (m *SchedulerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerConfig.Merge(m, src)
}
func (m *SchedulerConfig) XXX_Size() int {
	return xxx_messageInfo_SchedulerConfig.Size(m)
}
func (m *SchedulerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerConfig proto.InternalMessageInfo

func (m *SchedulerConfig) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_UPSTREAM
}

func (m *SchedulerConfig) GetAdditionalBw() AdditionalBW {
	if m != nil {
		return m.AdditionalBw
	}
	return AdditionalBW_AdditionalBW_None
}

func (m *SchedulerConfig) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SchedulerConfig) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *SchedulerConfig) GetSchedPolicy() SchedulingPolicy {
	if m != nil {
		return m.SchedPolicy
	}
	return SchedulingPolicy_WRR
}

type TrafficShapingInfo struct {
	Cir                  uint32                       `protobuf:"fixed32,1,opt,name=cir,proto3" json:"cir,omitempty"`
	Cbs                  uint32                       `protobuf:"fixed32,2,opt,name=cbs,proto3" json:"cbs,omitempty"`
	Pir                  uint32                       `protobuf:"fixed32,3,opt,name=pir,proto3" json:"pir,omitempty"`
	Pbs                  uint32                       `protobuf:"fixed32,4,opt,name=pbs,proto3" json:"pbs,omitempty"`
	Gir                  uint32                       `protobuf:"fixed32,5,opt,name=gir,proto3" json:"gir,omitempty"`
	AddBwInd             InferredAdditionBWIndication `protobuf:"varint,6,opt,name=add_bw_ind,json=addBwInd,proto3,enum=tech_profile.InferredAdditionBWIndication" json:"add_bw_ind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TrafficShapingInfo) Reset()         { *m = TrafficShapingInfo{} }
func (m *TrafficShapingInfo) String() string { return proto.CompactTextString(m) }
func (*TrafficShapingInfo) ProtoMessage()    {}
func (*TrafficShapingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{1}
}

func (m *TrafficShapingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficShapingInfo.Unmarshal(m, b)
}
func (m *TrafficShapingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficShapingInfo.Marshal(b, m, deterministic)
}
func (m *TrafficShapingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficShapingInfo.Merge(m, src)
}
func (m *TrafficShapingInfo) XXX_Size() int {
	return xxx_messageInfo_TrafficShapingInfo.Size(m)
}
func (m *TrafficShapingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficShapingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficShapingInfo proto.InternalMessageInfo

func (m *TrafficShapingInfo) GetCir() uint32 {
	if m != nil {
		return m.Cir
	}
	return 0
}

func (m *TrafficShapingInfo) GetCbs() uint32 {
	if m != nil {
		return m.Cbs
	}
	return 0
}

func (m *TrafficShapingInfo) GetPir() uint32 {
	if m != nil {
		return m.Pir
	}
	return 0
}

func (m *TrafficShapingInfo) GetPbs() uint32 {
	if m != nil {
		return m.Pbs
	}
	return 0
}

func (m *TrafficShapingInfo) GetGir() uint32 {
	if m != nil {
		return m.Gir
	}
	return 0
}

func (m *TrafficShapingInfo) GetAddBwInd() InferredAdditionBWIndication {
	if m != nil {
		return m.AddBwInd
	}
	return InferredAdditionBWIndication_InferredAdditionBWIndication_None
}

type TrafficScheduler struct {
	Direction            Direction           `protobuf:"varint,1,opt,name=direction,proto3,enum=tech_profile.Direction" json:"direction,omitempty"`
	AllocId              uint32              `protobuf:"fixed32,2,opt,name=alloc_id,json=allocId,proto3" json:"alloc_id,omitempty"`
	Scheduler            *SchedulerConfig    `protobuf:"bytes,3,opt,name=scheduler,proto3" json:"scheduler,omitempty"`
	TrafficShapingInfo   *TrafficShapingInfo `protobuf:"bytes,4,opt,name=traffic_shaping_info,json=trafficShapingInfo,proto3" json:"traffic_shaping_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TrafficScheduler) Reset()         { *m = TrafficScheduler{} }
func (m *TrafficScheduler) String() string { return proto.CompactTextString(m) }
func (*TrafficScheduler) ProtoMessage()    {}
func (*TrafficScheduler) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{2}
}

func (m *TrafficScheduler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficScheduler.Unmarshal(m, b)
}
func (m *TrafficScheduler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficScheduler.Marshal(b, m, deterministic)
}
func (m *TrafficScheduler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficScheduler.Merge(m, src)
}
func (m *TrafficScheduler) XXX_Size() int {
	return xxx_messageInfo_TrafficScheduler.Size(m)
}
func (m *TrafficScheduler) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficScheduler.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficScheduler proto.InternalMessageInfo

func (m *TrafficScheduler) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_UPSTREAM
}

func (m *TrafficScheduler) GetAllocId() uint32 {
	if m != nil {
		return m.AllocId
	}
	return 0
}

func (m *TrafficScheduler) GetScheduler() *SchedulerConfig {
	if m != nil {
		return m.Scheduler
	}
	return nil
}

func (m *TrafficScheduler) GetTrafficShapingInfo() *TrafficShapingInfo {
	if m != nil {
		return m.TrafficShapingInfo
	}
	return nil
}

type TrafficSchedulers struct {
	IntfId               uint32              `protobuf:"fixed32,1,opt,name=intf_id,json=intfId,proto3" json:"intf_id,omitempty"`
	OnuId                uint32              `protobuf:"fixed32,2,opt,name=onu_id,json=onuId,proto3" json:"onu_id,omitempty"`
	UniId                uint32              `protobuf:"fixed32,4,opt,name=uni_id,json=uniId,proto3" json:"uni_id,omitempty"`
	PortNo               uint32              `protobuf:"fixed32,5,opt,name=port_no,json=portNo,proto3" json:"port_no,omitempty"`
	TrafficScheds        []*TrafficScheduler `protobuf:"bytes,3,rep,name=traffic_scheds,json=trafficScheds,proto3" json:"traffic_scheds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TrafficSchedulers) Reset()         { *m = TrafficSchedulers{} }
func (m *TrafficSchedulers) String() string { return proto.CompactTextString(m) }
func (*TrafficSchedulers) ProtoMessage()    {}
func (*TrafficSchedulers) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{3}
}

func (m *TrafficSchedulers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSchedulers.Unmarshal(m, b)
}
func (m *TrafficSchedulers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSchedulers.Marshal(b, m, deterministic)
}
func (m *TrafficSchedulers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSchedulers.Merge(m, src)
}
func (m *TrafficSchedulers) XXX_Size() int {
	return xxx_messageInfo_TrafficSchedulers.Size(m)
}
func (m *TrafficSchedulers) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSchedulers.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSchedulers proto.InternalMessageInfo

func (m *TrafficSchedulers) GetIntfId() uint32 {
	if m != nil {
		return m.IntfId
	}
	return 0
}

func (m *TrafficSchedulers) GetOnuId() uint32 {
	if m != nil {
		return m.OnuId
	}
	return 0
}

func (m *TrafficSchedulers) GetUniId() uint32 {
	if m != nil {
		return m.UniId
	}
	return 0
}

func (m *TrafficSchedulers) GetPortNo() uint32 {
	if m != nil {
		return m.PortNo
	}
	return 0
}

func (m *TrafficSchedulers) GetTrafficScheds() []*TrafficScheduler {
	if m != nil {
		return m.TrafficScheds
	}
	return nil
}

type TailDropDiscardConfig struct {
	QueueSize            uint32   `protobuf:"fixed32,1,opt,name=queue_size,json=queueSize,proto3" json:"queue_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TailDropDiscardConfig) Reset()         { *m = TailDropDiscardConfig{} }
func (m *TailDropDiscardConfig) String() string { return proto.CompactTextString(m) }
func (*TailDropDiscardConfig) ProtoMessage()    {}
func (*TailDropDiscardConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{4}
}

func (m *TailDropDiscardConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TailDropDiscardConfig.Unmarshal(m, b)
}
func (m *TailDropDiscardConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TailDropDiscardConfig.Marshal(b, m, deterministic)
}
func (m *TailDropDiscardConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TailDropDiscardConfig.Merge(m, src)
}
func (m *TailDropDiscardConfig) XXX_Size() int {
	return xxx_messageInfo_TailDropDiscardConfig.Size(m)
}
func (m *TailDropDiscardConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TailDropDiscardConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TailDropDiscardConfig proto.InternalMessageInfo

func (m *TailDropDiscardConfig) GetQueueSize() uint32 {
	if m != nil {
		return m.QueueSize
	}
	return 0
}

type RedDiscardConfig struct {
	MinThreshold         uint32   `protobuf:"fixed32,1,opt,name=min_threshold,json=minThreshold,proto3" json:"min_threshold,omitempty"`
	MaxThreshold         uint32   `protobuf:"fixed32,2,opt,name=max_threshold,json=maxThreshold,proto3" json:"max_threshold,omitempty"`
	MaxProbability       uint32   `protobuf:"fixed32,3,opt,name=max_probability,json=maxProbability,proto3" json:"max_probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedDiscardConfig) Reset()         { *m = RedDiscardConfig{} }
func (m *RedDiscardConfig) String() string { return proto.CompactTextString(m) }
func (*RedDiscardConfig) ProtoMessage()    {}
func (*RedDiscardConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{5}
}

func (m *RedDiscardConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedDiscardConfig.Unmarshal(m, b)
}
func (m *RedDiscardConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedDiscardConfig.Marshal(b, m, deterministic)
}
func (m *RedDiscardConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedDiscardConfig.Merge(m, src)
}
func (m *RedDiscardConfig) XXX_Size() int {
	return xxx_messageInfo_RedDiscardConfig.Size(m)
}
func (m *RedDiscardConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RedDiscardConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RedDiscardConfig proto.InternalMessageInfo

func (m *RedDiscardConfig) GetMinThreshold() uint32 {
	if m != nil {
		return m.MinThreshold
	}
	return 0
}

func (m *RedDiscardConfig) GetMaxThreshold() uint32 {
	if m != nil {
		return m.MaxThreshold
	}
	return 0
}

func (m *RedDiscardConfig) GetMaxProbability() uint32 {
	if m != nil {
		return m.MaxProbability
	}
	return 0
}

type WRedDiscardConfig struct {
	Green                *RedDiscardConfig `protobuf:"bytes,1,opt,name=green,proto3" json:"green,omitempty"`
	Yellow               *RedDiscardConfig `protobuf:"bytes,2,opt,name=yellow,proto3" json:"yellow,omitempty"`
	Red                  *RedDiscardConfig `protobuf:"bytes,3,opt,name=red,proto3" json:"red,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WRedDiscardConfig) Reset()         { *m = WRedDiscardConfig{} }
func (m *WRedDiscardConfig) String() string { return proto.CompactTextString(m) }
func (*WRedDiscardConfig) ProtoMessage()    {}
func (*WRedDiscardConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{6}
}

func (m *WRedDiscardConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WRedDiscardConfig.Unmarshal(m, b)
}
func (m *WRedDiscardConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WRedDiscardConfig.Marshal(b, m, deterministic)
}
func (m *WRedDiscardConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WRedDiscardConfig.Merge(m, src)
}
func (m *WRedDiscardConfig) XXX_Size() int {
	return xxx_messageInfo_WRedDiscardConfig.Size(m)
}
func (m *WRedDiscardConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WRedDiscardConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WRedDiscardConfig proto.InternalMessageInfo

func (m *WRedDiscardConfig) GetGreen() *RedDiscardConfig {
	if m != nil {
		return m.Green
	}
	return nil
}

func (m *WRedDiscardConfig) GetYellow() *RedDiscardConfig {
	if m != nil {
		return m.Yellow
	}
	return nil
}

func (m *WRedDiscardConfig) GetRed() *RedDiscardConfig {
	if m != nil {
		return m.Red
	}
	return nil
}

type DiscardConfig struct {
	DiscardPolicy DiscardPolicy `protobuf:"varint,1,opt,name=discard_policy,json=discardPolicy,proto3,enum=tech_profile.DiscardPolicy" json:"discard_policy,omitempty"`
	// Types that are valid to be assigned to DiscardConfig:
	//	*DiscardConfig_TailDropDiscardConfig
	//	*DiscardConfig_RedDiscardConfig
	//	*DiscardConfig_WredDiscardConfig
	DiscardConfig        isDiscardConfig_DiscardConfig `protobuf_oneof:"discard_config"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *DiscardConfig) Reset()         { *m = DiscardConfig{} }
func (m *DiscardConfig) String() string { return proto.CompactTextString(m) }
func (*DiscardConfig) ProtoMessage()    {}
func (*DiscardConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{7}
}

func (m *DiscardConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardConfig.Unmarshal(m, b)
}
func (m *DiscardConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardConfig.Marshal(b, m, deterministic)
}
func (m *DiscardConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardConfig.Merge(m, src)
}
func (m *DiscardConfig) XXX_Size() int {
	return xxx_messageInfo_DiscardConfig.Size(m)
}
func (m *DiscardConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardConfig proto.InternalMessageInfo

func (m *DiscardConfig) GetDiscardPolicy() DiscardPolicy {
	if m != nil {
		return m.DiscardPolicy
	}
	return DiscardPolicy_TailDrop
}

type isDiscardConfig_DiscardConfig interface {
	isDiscardConfig_DiscardConfig()
}

type DiscardConfig_TailDropDiscardConfig struct {
	TailDropDiscardConfig *TailDropDiscardConfig `protobuf:"bytes,2,opt,name=tail_drop_discard_config,json=tailDropDiscardConfig,proto3,oneof"`
}

type DiscardConfig_RedDiscardConfig struct {
	RedDiscardConfig *RedDiscardConfig `protobuf:"bytes,3,opt,name=red_discard_config,json=redDiscardConfig,proto3,oneof"`
}

type DiscardConfig_WredDiscardConfig struct {
	WredDiscardConfig *WRedDiscardConfig `protobuf:"bytes,4,opt,name=wred_discard_config,json=wredDiscardConfig,proto3,oneof"`
}

func (*DiscardConfig_TailDropDiscardConfig) isDiscardConfig_DiscardConfig() {}

func (*DiscardConfig_RedDiscardConfig) isDiscardConfig_DiscardConfig() {}

func (*DiscardConfig_WredDiscardConfig) isDiscardConfig_DiscardConfig() {}

func (m *DiscardConfig) GetDiscardConfig() isDiscardConfig_DiscardConfig {
	if m != nil {
		return m.DiscardConfig
	}
	return nil
}

func (m *DiscardConfig) GetTailDropDiscardConfig() *TailDropDiscardConfig {
	if x, ok := m.GetDiscardConfig().(*DiscardConfig_TailDropDiscardConfig); ok {
		return x.TailDropDiscardConfig
	}
	return nil
}

func (m *DiscardConfig) GetRedDiscardConfig() *RedDiscardConfig {
	if x, ok := m.GetDiscardConfig().(*DiscardConfig_RedDiscardConfig); ok {
		return x.RedDiscardConfig
	}
	return nil
}

func (m *DiscardConfig) GetWredDiscardConfig() *WRedDiscardConfig {
	if x, ok := m.GetDiscardConfig().(*DiscardConfig_WredDiscardConfig); ok {
		return x.WredDiscardConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DiscardConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DiscardConfig_TailDropDiscardConfig)(nil),
		(*DiscardConfig_RedDiscardConfig)(nil),
		(*DiscardConfig_WredDiscardConfig)(nil),
	}
}

type TrafficQueue struct {
	Direction            Direction        `protobuf:"varint,1,opt,name=direction,proto3,enum=tech_profile.Direction" json:"direction,omitempty"`
	GemportId            uint32           `protobuf:"fixed32,2,opt,name=gemport_id,json=gemportId,proto3" json:"gemport_id,omitempty"`
	PbitMap              string           `protobuf:"bytes,3,opt,name=pbit_map,json=pbitMap,proto3" json:"pbit_map,omitempty"`
	AesEncryption        bool             `protobuf:"varint,4,opt,name=aes_encryption,json=aesEncryption,proto3" json:"aes_encryption,omitempty"`
	SchedPolicy          SchedulingPolicy `protobuf:"varint,5,opt,name=sched_policy,json=schedPolicy,proto3,enum=tech_profile.SchedulingPolicy" json:"sched_policy,omitempty"`
	Priority             uint32           `protobuf:"fixed32,6,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight               uint32           `protobuf:"fixed32,7,opt,name=weight,proto3" json:"weight,omitempty"`
	DiscardPolicy        DiscardPolicy    `protobuf:"varint,8,opt,name=discard_policy,json=discardPolicy,proto3,enum=tech_profile.DiscardPolicy" json:"discard_policy,omitempty"`
	DiscardConfig        *DiscardConfig   `protobuf:"bytes,9,opt,name=discard_config,json=discardConfig,proto3" json:"discard_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TrafficQueue) Reset()         { *m = TrafficQueue{} }
func (m *TrafficQueue) String() string { return proto.CompactTextString(m) }
func (*TrafficQueue) ProtoMessage()    {}
func (*TrafficQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{8}
}

func (m *TrafficQueue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficQueue.Unmarshal(m, b)
}
func (m *TrafficQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficQueue.Marshal(b, m, deterministic)
}
func (m *TrafficQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficQueue.Merge(m, src)
}
func (m *TrafficQueue) XXX_Size() int {
	return xxx_messageInfo_TrafficQueue.Size(m)
}
func (m *TrafficQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficQueue.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficQueue proto.InternalMessageInfo

func (m *TrafficQueue) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_UPSTREAM
}

func (m *TrafficQueue) GetGemportId() uint32 {
	if m != nil {
		return m.GemportId
	}
	return 0
}

func (m *TrafficQueue) GetPbitMap() string {
	if m != nil {
		return m.PbitMap
	}
	return ""
}

func (m *TrafficQueue) GetAesEncryption() bool {
	if m != nil {
		return m.AesEncryption
	}
	return false
}

func (m *TrafficQueue) GetSchedPolicy() SchedulingPolicy {
	if m != nil {
		return m.SchedPolicy
	}
	return SchedulingPolicy_WRR
}

func (m *TrafficQueue) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TrafficQueue) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TrafficQueue) GetDiscardPolicy() DiscardPolicy {
	if m != nil {
		return m.DiscardPolicy
	}
	return DiscardPolicy_TailDrop
}

func (m *TrafficQueue) GetDiscardConfig() *DiscardConfig {
	if m != nil {
		return m.DiscardConfig
	}
	return nil
}

type TrafficQueues struct {
	IntfId               uint32          `protobuf:"fixed32,1,opt,name=intf_id,json=intfId,proto3" json:"intf_id,omitempty"`
	OnuId                uint32          `protobuf:"fixed32,2,opt,name=onu_id,json=onuId,proto3" json:"onu_id,omitempty"`
	UniId                uint32          `protobuf:"fixed32,4,opt,name=uni_id,json=uniId,proto3" json:"uni_id,omitempty"`
	PortNo               uint32          `protobuf:"fixed32,5,opt,name=port_no,json=portNo,proto3" json:"port_no,omitempty"`
	TrafficQueues        []*TrafficQueue `protobuf:"bytes,6,rep,name=traffic_queues,json=trafficQueues,proto3" json:"traffic_queues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TrafficQueues) Reset()         { *m = TrafficQueues{} }
func (m *TrafficQueues) String() string { return proto.CompactTextString(m) }
func (*TrafficQueues) ProtoMessage()    {}
func (*TrafficQueues) Descriptor() ([]byte, []int) {
	return fileDescriptor_d019a68bffe14cae, []int{9}
}

func (m *TrafficQueues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficQueues.Unmarshal(m, b)
}
func (m *TrafficQueues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficQueues.Marshal(b, m, deterministic)
}
func (m *TrafficQueues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficQueues.Merge(m, src)
}
func (m *TrafficQueues) XXX_Size() int {
	return xxx_messageInfo_TrafficQueues.Size(m)
}
func (m *TrafficQueues) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficQueues.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficQueues proto.InternalMessageInfo

func (m *TrafficQueues) GetIntfId() uint32 {
	if m != nil {
		return m.IntfId
	}
	return 0
}

func (m *TrafficQueues) GetOnuId() uint32 {
	if m != nil {
		return m.OnuId
	}
	return 0
}

func (m *TrafficQueues) GetUniId() uint32 {
	if m != nil {
		return m.UniId
	}
	return 0
}

func (m *TrafficQueues) GetPortNo() uint32 {
	if m != nil {
		return m.PortNo
	}
	return 0
}

func (m *TrafficQueues) GetTrafficQueues() []*TrafficQueue {
	if m != nil {
		return m.TrafficQueues
	}
	return nil
}

func init() {
	proto.RegisterEnum("tech_profile.Direction", Direction_name, Direction_value)
	proto.RegisterEnum("tech_profile.SchedulingPolicy", SchedulingPolicy_name, SchedulingPolicy_value)
	proto.RegisterEnum("tech_profile.AdditionalBW", AdditionalBW_name, AdditionalBW_value)
	proto.RegisterEnum("tech_profile.DiscardPolicy", DiscardPolicy_name, DiscardPolicy_value)
	proto.RegisterEnum("tech_profile.InferredAdditionBWIndication", InferredAdditionBWIndication_name, InferredAdditionBWIndication_value)
	proto.RegisterType((*SchedulerConfig)(nil), "tech_profile.SchedulerConfig")
	proto.RegisterType((*TrafficShapingInfo)(nil), "tech_profile.TrafficShapingInfo")
	proto.RegisterType((*TrafficScheduler)(nil), "tech_profile.TrafficScheduler")
	proto.RegisterType((*TrafficSchedulers)(nil), "tech_profile.TrafficSchedulers")
	proto.RegisterType((*TailDropDiscardConfig)(nil), "tech_profile.TailDropDiscardConfig")
	proto.RegisterType((*RedDiscardConfig)(nil), "tech_profile.RedDiscardConfig")
	proto.RegisterType((*WRedDiscardConfig)(nil), "tech_profile.WRedDiscardConfig")
	proto.RegisterType((*DiscardConfig)(nil), "tech_profile.DiscardConfig")
	proto.RegisterType((*TrafficQueue)(nil), "tech_profile.TrafficQueue")
	proto.RegisterType((*TrafficQueues)(nil), "tech_profile.TrafficQueues")
}

func init() { proto.RegisterFile("voltha_protos/tech_profile.proto", fileDescriptor_d019a68bffe14cae) }

var fileDescriptor_d019a68bffe14cae = []byte{
	// 1072 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0x8d, 0x7f, 0x4e, 0x6c, 0x77, 0x33, 0x25, 0xc4, 0xa4, 0x0d, 0x04, 0x97, 0xaa,
	0x91, 0x91, 0x1a, 0x14, 0xa0, 0x37, 0x45, 0xaa, 0xec, 0x26, 0x52, 0x2c, 0xd1, 0x34, 0x9d, 0x04,
	0xf9, 0x8e, 0xd5, 0x78, 0x67, 0x6c, 0x8f, 0xb4, 0x99, 0x59, 0x66, 0xc7, 0x38, 0xe9, 0x15, 0x37,
	0xbc, 0x05, 0xb7, 0xbc, 0x00, 0xdc, 0x20, 0x9e, 0x88, 0x17, 0xe0, 0x1e, 0xcd, 0xec, 0xae, 0xed,
	0x5d, 0x9b, 0x14, 0x2a, 0xb8, 0x9b, 0xf3, 0xed, 0x37, 0x67, 0xce, 0x37, 0xe7, 0x67, 0x16, 0xf6,
	0xbf, 0x97, 0x81, 0x9e, 0x10, 0x2f, 0x54, 0x52, 0xcb, 0xe8, 0x50, 0x33, 0x7f, 0x62, 0xd6, 0x23,
	0x1e, 0xb0, 0x27, 0x16, 0x43, 0xf5, 0x65, 0x6c, 0xf7, 0xc1, 0x58, 0xca, 0x71, 0xc0, 0x0e, 0x49,
	0xc8, 0x0f, 0x89, 0x10, 0x52, 0x13, 0xcd, 0xa5, 0x88, 0x62, 0x6e, 0xfb, 0x87, 0x22, 0xdc, 0xbd,
	0xf0, 0x27, 0x8c, 0x4e, 0x03, 0xa6, 0x5e, 0x48, 0x31, 0xe2, 0x63, 0xf4, 0x25, 0xd4, 0x28, 0x57,
	0xcc, 0x37, 0xbc, 0x96, 0xb3, 0xef, 0x1c, 0x34, 0x8f, 0x76, 0x9e, 0x64, 0xce, 0x39, 0x4e, 0x3f,
	0xe3, 0x05, 0x13, 0x3d, 0x87, 0x06, 0xa1, 0x94, 0x9b, 0x35, 0x09, 0xbc, 0xe1, 0xac, 0x55, 0xb4,
	0x5b, 0x77, 0xb3, 0x5b, 0xbb, 0x73, 0x4a, 0x6f, 0x80, 0xeb, 0x8b, 0x0d, 0xbd, 0x19, 0xda, 0x85,
	0x6a, 0xa8, 0xb8, 0x54, 0x5c, 0xdf, 0xb4, 0x4a, 0xfb, 0xce, 0x41, 0x05, 0xcf, 0x6d, 0xf4, 0x3e,
	0x94, 0x67, 0x8c, 0x8f, 0x27, 0xba, 0x75, 0xc7, 0x7e, 0x49, 0x2c, 0xd4, 0x85, 0x7a, 0x64, 0xc2,
	0xf7, 0x42, 0x19, 0x70, 0xff, 0xa6, 0xb5, 0x61, 0xcf, 0xfc, 0x30, 0x7b, 0x66, 0x22, 0x90, 0x8b,
	0xf1, 0xb9, 0x65, 0xe1, 0x4d, 0xbb, 0x27, 0x36, 0xda, 0xbf, 0x39, 0x80, 0x2e, 0x15, 0x19, 0x8d,
	0xb8, 0x7f, 0x31, 0x21, 0x21, 0x17, 0xe3, 0xbe, 0x18, 0x49, 0xe4, 0x42, 0xc9, 0xe7, 0xca, 0xea,
	0xaf, 0x60, 0xb3, 0xb4, 0xc8, 0x30, 0xb2, 0xb2, 0x0c, 0x32, 0x8c, 0x0c, 0x12, 0x72, 0x95, 0x04,
	0x6b, 0x96, 0x16, 0x19, 0x46, 0x49, 0x90, 0x66, 0x69, 0x90, 0x31, 0x57, 0x36, 0xb0, 0x0a, 0x36,
	0x4b, 0x74, 0x0a, 0x40, 0x28, 0xf5, 0x86, 0x33, 0x8f, 0x0b, 0xda, 0x2a, 0xdb, 0x88, 0x3b, 0xd9,
	0x88, 0xfb, 0x62, 0xc4, 0x94, 0x62, 0x34, 0xbd, 0xad, 0xde, 0xa0, 0x2f, 0x28, 0xf7, 0x6d, 0xea,
	0x70, 0x95, 0x50, 0xda, 0x9b, 0xf5, 0x05, 0x6d, 0xff, 0xe9, 0x80, 0x9b, 0x86, 0x9e, 0x26, 0xf1,
	0x5d, 0xd3, 0xf7, 0x01, 0x54, 0x49, 0x10, 0x48, 0xdf, 0xe3, 0x34, 0x91, 0x58, 0xb1, 0x76, 0x9f,
	0xa2, 0x67, 0x50, 0x8b, 0x52, 0xf7, 0x56, 0xec, 0xe6, 0xd1, 0xde, 0xda, 0x1b, 0x4e, 0x4b, 0x08,
	0x2f, 0xf8, 0x08, 0xc3, 0x7b, 0x3a, 0x0e, 0xd1, 0x8b, 0xe2, 0xeb, 0xf5, 0xb8, 0x18, 0x49, 0x7b,
	0x45, 0x9b, 0x47, 0xfb, 0x59, 0x3f, 0xab, 0x79, 0xc0, 0x48, 0xaf, 0x60, 0xed, 0xdf, 0x1d, 0xd8,
	0xca, 0xeb, 0x8e, 0xd0, 0x0e, 0x54, 0xb8, 0xd0, 0x23, 0x23, 0x20, 0xce, 0x5a, 0xd9, 0x98, 0x7d,
	0x8a, 0xb6, 0xa1, 0x2c, 0xc5, 0x74, 0x21, 0x6c, 0x43, 0x8a, 0x69, 0x0c, 0x4f, 0x05, 0x37, 0x70,
	0x9c, 0xae, 0x8d, 0xa9, 0xe0, 0x7d, 0x6a, 0xdc, 0x84, 0x52, 0x69, 0x4f, 0xc8, 0x24, 0x69, 0x65,
	0x63, 0x9e, 0x49, 0x74, 0x02, 0xcd, 0xb9, 0x12, 0x73, 0x6a, 0xd4, 0x2a, 0xed, 0x97, 0x0e, 0x36,
	0xf3, 0xd5, 0x96, 0x0f, 0x0c, 0x37, 0xf4, 0x12, 0x12, 0xb5, 0x9f, 0xc2, 0xf6, 0x25, 0xe1, 0xc1,
	0xb1, 0x92, 0xe1, 0x31, 0x8f, 0x7c, 0xa2, 0x68, 0xd2, 0x77, 0x7b, 0x00, 0xdf, 0x4d, 0xd9, 0x94,
	0x79, 0x11, 0x7f, 0xc3, 0x12, 0x09, 0x35, 0x8b, 0x5c, 0xf0, 0x37, 0xac, 0xfd, 0xa3, 0x03, 0x2e,
	0x66, 0x34, 0xbb, 0xe7, 0x21, 0x34, 0xae, 0xb8, 0xf0, 0xf4, 0x44, 0xb1, 0x68, 0x22, 0x83, 0x54,
	0x79, 0xfd, 0x8a, 0x8b, 0xcb, 0x14, 0xb3, 0x24, 0x72, 0xbd, 0x44, 0x2a, 0x26, 0x24, 0x72, 0xbd,
	0x20, 0x3d, 0x86, 0xbb, 0x86, 0x14, 0x2a, 0x39, 0x24, 0x43, 0x1e, 0x2c, 0x9a, 0xb0, 0x79, 0x45,
	0xae, 0xcf, 0x17, 0x68, 0xfb, 0x57, 0x07, 0xb6, 0x06, 0x2b, 0x81, 0x7c, 0x01, 0x1b, 0x63, 0xc5,
	0x58, 0x5c, 0x71, 0x2b, 0x77, 0x92, 0xa7, 0xe3, 0x98, 0x8c, 0x9e, 0x42, 0xf9, 0x86, 0x05, 0x81,
	0x8c, 0x87, 0xc5, 0xdb, 0xb7, 0x25, 0x6c, 0xf4, 0x19, 0x94, 0x14, 0xa3, 0x49, 0x2d, 0xbe, 0x6d,
	0x93, 0xa1, 0xb6, 0xff, 0x28, 0x42, 0x23, 0x1b, 0x71, 0x0f, 0x9a, 0x34, 0x06, 0xd2, 0xe1, 0x11,
	0x37, 0xcb, 0xfd, 0x7c, 0xb3, 0x58, 0x4e, 0x32, 0x39, 0x1a, 0x74, 0xd9, 0x44, 0xdf, 0x42, 0x4b,
	0x13, 0x1e, 0x78, 0x54, 0xc9, 0xd0, 0x4b, 0xbd, 0xf9, 0xd6, 0x7f, 0xa2, 0xe8, 0x61, 0xae, 0x38,
	0xd6, 0x65, 0xfe, 0xb4, 0x80, 0xb7, 0xf5, 0xda, 0x92, 0x38, 0x03, 0xa4, 0x18, 0xcd, 0x7b, 0xfe,
	0x47, 0xb2, 0x4f, 0x0b, 0xd8, 0x55, 0xf9, 0x2c, 0xbd, 0x86, 0x7b, 0xb3, 0x35, 0x0e, 0xe3, 0x5e,
	0xfc, 0x28, 0xeb, 0x70, 0xb0, 0xc6, 0xe3, 0xd6, 0x2c, 0xef, 0xb2, 0xe7, 0x2e, 0xae, 0x31, 0xf6,
	0xd6, 0xfe, 0xb9, 0x04, 0xf5, 0xa4, 0x09, 0x5e, 0x9b, 0xea, 0x7d, 0xd7, 0x89, 0xb4, 0x07, 0x30,
	0x66, 0x57, 0xb6, 0x17, 0xe7, 0xad, 0x5b, 0x4b, 0x90, 0x3e, 0x35, 0x03, 0x2b, 0x1c, 0x72, 0xed,
	0x5d, 0x91, 0xd0, 0xde, 0x48, 0x0d, 0x57, 0x8c, 0xfd, 0x92, 0x84, 0xe8, 0x11, 0x34, 0x09, 0x8b,
	0x3c, 0x26, 0x7c, 0x75, 0x13, 0xda, 0x53, 0x8d, 0xc2, 0x2a, 0x6e, 0x10, 0x16, 0x9d, 0xcc, 0xc1,
	0xff, 0xe0, 0xf1, 0xc8, 0xbc, 0x59, 0xe5, 0xbf, 0x7d, 0xb3, 0x2a, 0x99, 0x37, 0x6b, 0xb5, 0xf0,
	0xaa, 0xff, 0xba, 0xf0, 0x7a, 0xf9, 0x5b, 0x6f, 0xd5, 0x6c, 0x0e, 0xd7, 0xfb, 0x48, 0x1a, 0x21,
	0xf5, 0x11, 0x9b, 0xed, 0x5f, 0x1c, 0x68, 0x2c, 0xe7, 0xe9, 0xff, 0x9f, 0xa0, 0xdd, 0xc5, 0x04,
	0xb5, 0x73, 0x2d, 0x6a, 0x95, 0xed, 0x04, 0xdd, 0x5d, 0x3b, 0x41, 0x6d, 0x50, 0xf3, 0xe9, 0x19,
	0x87, 0xd8, 0xf9, 0x0a, 0x6a, 0xf3, 0x62, 0x41, 0x75, 0xa8, 0x7e, 0x73, 0x7e, 0x71, 0x89, 0x4f,
	0xba, 0x2f, 0xdd, 0x02, 0x6a, 0x02, 0x1c, 0xbf, 0x1a, 0x9c, 0x25, 0xb6, 0x83, 0xb6, 0xa0, 0xd1,
	0xeb, 0x1f, 0xf7, 0xf1, 0xc9, 0x8b, 0xcb, 0xfe, 0xab, 0xb3, 0xee, 0xd7, 0x6e, 0xb1, 0xf3, 0x0c,
	0xdc, 0x7c, 0x3e, 0x51, 0x05, 0x4a, 0x03, 0x8c, 0xdd, 0x02, 0x42, 0xd0, 0xbc, 0xd0, 0x8a, 0xfb,
	0xfa, 0x3c, 0xc9, 0xa0, 0xeb, 0x20, 0x80, 0xf2, 0xe9, 0xcd, 0x50, 0x71, 0xea, 0x16, 0x3b, 0x02,
	0xea, 0xcb, 0x7f, 0x2f, 0x68, 0x1b, 0xb6, 0x96, 0x6d, 0xef, 0x4c, 0x0a, 0xe6, 0x16, 0xd0, 0x3d,
	0xb8, 0x9b, 0x85, 0xbb, 0xae, 0x83, 0xee, 0xc3, 0x4e, 0x06, 0xec, 0xb1, 0x48, 0x9f, 0x8c, 0x46,
	0x52, 0x69, 0xb7, 0xb8, 0xe2, 0xa8, 0x3b, 0xd5, 0xd2, 0x2d, 0x75, 0x9e, 0xcf, 0x27, 0x56, 0x12,
	0x69, 0x1d, 0xaa, 0xe9, 0xfc, 0x70, 0x0b, 0xa8, 0x01, 0xb5, 0xc1, 0xdc, 0x74, 0x8c, 0x0c, 0xcc,
	0xa8, 0x5b, 0x44, 0x55, 0xb8, 0x63, 0x5a, 0xd7, 0x2d, 0x75, 0x7e, 0x72, 0xe0, 0xc1, 0x6d, 0x7f,
	0x12, 0xe8, 0x11, 0x7c, 0x7c, 0xdb, 0xf7, 0x54, 0xd1, 0x01, 0x7c, 0x72, 0x2b, 0xad, 0x1b, 0x45,
	0x53, 0xc5, 0xa8, 0xeb, 0xa0, 0x4f, 0xe1, 0xf1, 0xad, 0xcc, 0x65, 0xd9, 0xc3, 0xb2, 0xfd, 0x05,
	0xfd, 0xfc, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x2b, 0xaf, 0x34, 0xd2, 0x0a, 0x00, 0x00,
}
