// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/bbf_fiber_types.proto

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

type AuthMethodType int32

const (
	AuthMethodType_SERIAL_NUMBER   AuthMethodType = 0
	AuthMethodType_LOID            AuthMethodType = 1
	AuthMethodType_REGISTRATION_ID AuthMethodType = 2
	AuthMethodType_OMCI            AuthMethodType = 3
	AuthMethodType_DOT1X           AuthMethodType = 4
)

var AuthMethodType_name = map[int32]string{
	0: "SERIAL_NUMBER",
	1: "LOID",
	2: "REGISTRATION_ID",
	3: "OMCI",
	4: "DOT1X",
}

var AuthMethodType_value = map[string]int32{
	"SERIAL_NUMBER":   0,
	"LOID":            1,
	"REGISTRATION_ID": 2,
	"OMCI":            3,
	"DOT1X":           4,
}

func (x AuthMethodType) String() string {
	return proto.EnumName(AuthMethodType_name, int32(x))
}

func (AuthMethodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{0}
}

type RamanMitigationType int32

const (
	RamanMitigationType_RAMAN_NONE   RamanMitigationType = 0
	RamanMitigationType_RAMAN_MILLER RamanMitigationType = 1
	RamanMitigationType_RAMAN_8B10B  RamanMitigationType = 2
)

var RamanMitigationType_name = map[int32]string{
	0: "RAMAN_NONE",
	1: "RAMAN_MILLER",
	2: "RAMAN_8B10B",
}

var RamanMitigationType_value = map[string]int32{
	"RAMAN_NONE":   0,
	"RAMAN_MILLER": 1,
	"RAMAN_8B10B":  2,
}

func (x RamanMitigationType) String() string {
	return proto.EnumName(RamanMitigationType_name, int32(x))
}

func (RamanMitigationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{1}
}

type PonIdOdnClassType int32

const (
	PonIdOdnClassType_CLASS_A      PonIdOdnClassType = 0
	PonIdOdnClassType_CLASS_B      PonIdOdnClassType = 1
	PonIdOdnClassType_CLASS_B_PLUS PonIdOdnClassType = 2
	PonIdOdnClassType_CLASS_C      PonIdOdnClassType = 3
	PonIdOdnClassType_CLASS_C_PLUS PonIdOdnClassType = 4
	PonIdOdnClassType_CLASS_AUTO   PonIdOdnClassType = 255
)

var PonIdOdnClassType_name = map[int32]string{
	0:   "CLASS_A",
	1:   "CLASS_B",
	2:   "CLASS_B_PLUS",
	3:   "CLASS_C",
	4:   "CLASS_C_PLUS",
	255: "CLASS_AUTO",
}

var PonIdOdnClassType_value = map[string]int32{
	"CLASS_A":      0,
	"CLASS_B":      1,
	"CLASS_B_PLUS": 2,
	"CLASS_C":      3,
	"CLASS_C_PLUS": 4,
	"CLASS_AUTO":   255,
}

func (x PonIdOdnClassType) String() string {
	return proto.EnumName(PonIdOdnClassType_name, int32(x))
}

func (PonIdOdnClassType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{2}
}

type ChannelpairSpeedType int32

const (
	ChannelpairSpeedType_UNPLANNED_CP_SPEED  ChannelpairSpeedType = 0
	ChannelpairSpeedType_DOWN_10_UP_10       ChannelpairSpeedType = 1
	ChannelpairSpeedType_DOWN_10_UP_2DOT5    ChannelpairSpeedType = 2
	ChannelpairSpeedType_DOWN_2DOT5_UP_2DOT5 ChannelpairSpeedType = 3
)

var ChannelpairSpeedType_name = map[int32]string{
	0: "UNPLANNED_CP_SPEED",
	1: "DOWN_10_UP_10",
	2: "DOWN_10_UP_2DOT5",
	3: "DOWN_2DOT5_UP_2DOT5",
}

var ChannelpairSpeedType_value = map[string]int32{
	"UNPLANNED_CP_SPEED":  0,
	"DOWN_10_UP_10":       1,
	"DOWN_10_UP_2DOT5":    2,
	"DOWN_2DOT5_UP_2DOT5": 3,
}

func (x ChannelpairSpeedType) String() string {
	return proto.EnumName(ChannelpairSpeedType_name, int32(x))
}

func (ChannelpairSpeedType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{3}
}

type ChannelpairType int32

const (
	ChannelpairType_CHANNELPAIR     ChannelpairType = 0
	ChannelpairType_CHANNELPAIR_XGS ChannelpairType = 1
)

var ChannelpairType_name = map[int32]string{
	0: "CHANNELPAIR",
	1: "CHANNELPAIR_XGS",
}

var ChannelpairType_value = map[string]int32{
	"CHANNELPAIR":     0,
	"CHANNELPAIR_XGS": 1,
}

func (x ChannelpairType) String() string {
	return proto.EnumName(ChannelpairType_name, int32(x))
}

func (ChannelpairType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{4}
}

type TypeBRoleType int32

const (
	TypeBRoleType_PRIMARY   TypeBRoleType = 0
	TypeBRoleType_SECONDARY TypeBRoleType = 1
)

var TypeBRoleType_name = map[int32]string{
	0: "PRIMARY",
	1: "SECONDARY",
}

var TypeBRoleType_value = map[string]int32{
	"PRIMARY":   0,
	"SECONDARY": 1,
}

func (x TypeBRoleType) String() string {
	return proto.EnumName(TypeBRoleType_name, int32(x))
}

func (TypeBRoleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{5}
}

type TypeBRoleStateType int32

const (
	TypeBRoleStateType_ACTIVE  TypeBRoleStateType = 0
	TypeBRoleStateType_STANDBY TypeBRoleStateType = 1
)

var TypeBRoleStateType_name = map[int32]string{
	0: "ACTIVE",
	1: "STANDBY",
}

var TypeBRoleStateType_value = map[string]int32{
	"ACTIVE":  0,
	"STANDBY": 1,
}

func (x TypeBRoleStateType) String() string {
	return proto.EnumName(TypeBRoleStateType_name, int32(x))
}

func (TypeBRoleStateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dd3e241efaba3dc, []int{6}
}

func init() {
	proto.RegisterEnum("bbf_fiber_types.AuthMethodType", AuthMethodType_name, AuthMethodType_value)
	proto.RegisterEnum("bbf_fiber_types.RamanMitigationType", RamanMitigationType_name, RamanMitigationType_value)
	proto.RegisterEnum("bbf_fiber_types.PonIdOdnClassType", PonIdOdnClassType_name, PonIdOdnClassType_value)
	proto.RegisterEnum("bbf_fiber_types.ChannelpairSpeedType", ChannelpairSpeedType_name, ChannelpairSpeedType_value)
	proto.RegisterEnum("bbf_fiber_types.ChannelpairType", ChannelpairType_name, ChannelpairType_value)
	proto.RegisterEnum("bbf_fiber_types.TypeBRoleType", TypeBRoleType_name, TypeBRoleType_value)
	proto.RegisterEnum("bbf_fiber_types.TypeBRoleStateType", TypeBRoleStateType_name, TypeBRoleStateType_value)
}

func init() {
	proto.RegisterFile("voltha_protos/bbf_fiber_types.proto", fileDescriptor_2dd3e241efaba3dc)
}

var fileDescriptor_2dd3e241efaba3dc = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x9b, 0xae, 0x0c, 0x76, 0x4b, 0x97, 0x3b, 0x77, 0x82, 0x0f, 0xc0, 0x5b, 0x50, 0xff,
	0x81, 0xa6, 0xf1, 0xea, 0xfc, 0xd1, 0x66, 0x29, 0x71, 0x22, 0x3b, 0x65, 0x1b, 0x2f, 0x96, 0x43,
	0xb3, 0x36, 0xa2, 0xc4, 0x51, 0x1b, 0x26, 0xed, 0xd3, 0x83, 0x92, 0x0c, 0x5a, 0xf1, 0x78, 0x7e,
	0xe7, 0xfa, 0x5c, 0xe9, 0xf8, 0xc2, 0x87, 0x27, 0xb3, 0xad, 0x37, 0x5a, 0x55, 0x3b, 0x53, 0x9b,
	0xfd, 0x2c, 0xcb, 0x1e, 0xd5, 0x63, 0x91, 0xe5, 0x3b, 0x55, 0x3f, 0x57, 0xf9, 0x7e, 0xda, 0x62,
	0x62, 0xff, 0x87, 0x9d, 0x3b, 0x38, 0xa7, 0xbf, 0xea, 0x4d, 0x94, 0xd7, 0x1b, 0xb3, 0x4a, 0x9f,
	0xab, 0x9c, 0x5c, 0xc0, 0x48, 0x06, 0x82, 0xd1, 0x50, 0xf1, 0x65, 0xe4, 0x06, 0x02, 0x7b, 0xe4,
	0x0d, 0x0c, 0xc2, 0x98, 0xf9, 0x68, 0x91, 0x31, 0xd8, 0x22, 0xb8, 0x61, 0x32, 0x15, 0x34, 0x65,
	0x31, 0x57, 0xcc, 0xc7, 0x7e, 0x63, 0xc7, 0x91, 0xc7, 0xf0, 0x84, 0x9c, 0xc1, 0x2b, 0x3f, 0x4e,
	0x17, 0xf7, 0x38, 0x70, 0x6e, 0x61, 0x2c, 0xf4, 0x4f, 0x5d, 0x46, 0x45, 0x5d, 0xac, 0x75, 0x5d,
	0x98, 0xb2, 0x4d, 0x3f, 0x07, 0x10, 0x34, 0xa2, 0x5c, 0xf1, 0x98, 0x07, 0xd8, 0x23, 0x08, 0x6f,
	0x3b, 0x1d, 0xb1, 0x30, 0x0c, 0x04, 0x5a, 0xc4, 0x86, 0x61, 0x47, 0xbe, 0xb8, 0x8b, 0xb9, 0x8b,
	0x7d, 0xc7, 0xc0, 0x45, 0x62, 0x4a, 0xb6, 0x8a, 0x57, 0xa5, 0xb7, 0xd5, 0xfb, 0x7d, 0x9b, 0x33,
	0x84, 0xd7, 0x5e, 0x48, 0xa5, 0x54, 0x14, 0x7b, 0x07, 0xe1, 0xa2, 0xd5, 0x24, 0xbe, 0x08, 0x95,
	0x84, 0x4b, 0x89, 0xfd, 0x83, 0xed, 0xe1, 0xc9, 0xc1, 0xf6, 0x3a, 0x7b, 0x40, 0x6c, 0x80, 0x97,
	0xa8, 0x65, 0x1a, 0xe3, 0x6f, 0xcb, 0xa9, 0xe0, 0xd2, 0xdb, 0xe8, 0xb2, 0xcc, 0xb7, 0x95, 0x2e,
	0x76, 0xb2, 0xca, 0xf3, 0xae, 0x99, 0x77, 0x40, 0x96, 0x3c, 0x09, 0x29, 0xe7, 0x81, 0xaf, 0xbc,
	0x44, 0xc9, 0x24, 0x08, 0x7c, 0xec, 0x35, 0x8d, 0xf9, 0xf1, 0x1d, 0x57, 0x8b, 0xb9, 0x5a, 0x26,
	0x6a, 0x31, 0x47, 0x8b, 0x5c, 0x02, 0x1e, 0xa1, 0x4f, 0x7e, 0x9c, 0x5e, 0x61, 0x9f, 0xbc, 0x87,
	0x71, 0x4b, 0x5b, 0x7d, 0x30, 0x4e, 0x9c, 0x6b, 0xb0, 0x8f, 0x36, 0xb6, 0xcb, 0x6c, 0x18, 0x7a,
	0xb7, 0xcd, 0xaa, 0x30, 0xa1, 0xac, 0xf9, 0x84, 0x31, 0xd8, 0x47, 0x40, 0xdd, 0xdf, 0x48, 0xb4,
	0x9c, 0x8f, 0x30, 0x6a, 0xa6, 0x5d, 0x61, 0xb6, 0xf9, 0xdf, 0x5e, 0x12, 0xc1, 0x22, 0x2a, 0x1e,
	0xb0, 0x47, 0x46, 0x70, 0x26, 0x03, 0x2f, 0xe6, 0x7e, 0x23, 0x2d, 0x67, 0x02, 0xe4, 0xdf, 0xb0,
	0xac, 0x75, 0xdd, 0xbd, 0x00, 0x38, 0xa5, 0x5e, 0xca, 0xbe, 0x06, 0x5d, 0x91, 0x32, 0xa5, 0xdc,
	0x77, 0x1f, 0xd0, 0x72, 0xaf, 0xbf, 0x5d, 0xad, 0x4d, 0xf5, 0x63, 0x3d, 0x2d, 0xca, 0xd9, 0x2a,
	0xd3, 0x45, 0x99, 0xed, 0x8a, 0xc9, 0xf7, 0x22, 0x2f, 0xf5, 0xac, 0x3b, 0xb5, 0x49, 0x77, 0x6a,
	0xd3, 0xa7, 0xc5, 0x6c, 0x6d, 0x0e, 0x07, 0x97, 0x9d, 0xb6, 0xfc, 0xf3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x01, 0x47, 0x87, 0x92, 0x02, 0x00, 0x00,
}
