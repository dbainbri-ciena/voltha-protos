// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/schema.proto

package schema

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Contains the name and content of a *.proto file
type ProtoFile struct {
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Proto                string   `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Descriptor_          []byte   `protobuf:"bytes,3,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoFile) Reset()         { *m = ProtoFile{} }
func (m *ProtoFile) String() string { return proto.CompactTextString(m) }
func (*ProtoFile) ProtoMessage()    {}
func (*ProtoFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d785a7d5fd3f7a9, []int{0}
}

func (m *ProtoFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoFile.Unmarshal(m, b)
}
func (m *ProtoFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoFile.Marshal(b, m, deterministic)
}
func (m *ProtoFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoFile.Merge(m, src)
}
func (m *ProtoFile) XXX_Size() int {
	return xxx_messageInfo_ProtoFile.Size(m)
}
func (m *ProtoFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoFile.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoFile proto.InternalMessageInfo

func (m *ProtoFile) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ProtoFile) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *ProtoFile) GetDescriptor_() []byte {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

// Proto files and compiled descriptors for this interface
type Schemas struct {
	// Proto files
	Protos []*ProtoFile `protobuf:"bytes,1,rep,name=protos,proto3" json:"protos,omitempty"`
	// Proto file name from which swagger.json shall be generated
	SwaggerFrom string `protobuf:"bytes,2,opt,name=swagger_from,json=swaggerFrom,proto3" json:"swagger_from,omitempty"`
	// Proto file name from which yang schemas shall be generated
	YangFrom             string   `protobuf:"bytes,3,opt,name=yang_from,json=yangFrom,proto3" json:"yang_from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schemas) Reset()         { *m = Schemas{} }
func (m *Schemas) String() string { return proto.CompactTextString(m) }
func (*Schemas) ProtoMessage()    {}
func (*Schemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d785a7d5fd3f7a9, []int{1}
}

func (m *Schemas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schemas.Unmarshal(m, b)
}
func (m *Schemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schemas.Marshal(b, m, deterministic)
}
func (m *Schemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schemas.Merge(m, src)
}
func (m *Schemas) XXX_Size() int {
	return xxx_messageInfo_Schemas.Size(m)
}
func (m *Schemas) XXX_DiscardUnknown() {
	xxx_messageInfo_Schemas.DiscardUnknown(m)
}

var xxx_messageInfo_Schemas proto.InternalMessageInfo

func (m *Schemas) GetProtos() []*ProtoFile {
	if m != nil {
		return m.Protos
	}
	return nil
}

func (m *Schemas) GetSwaggerFrom() string {
	if m != nil {
		return m.SwaggerFrom
	}
	return ""
}

func (m *Schemas) GetYangFrom() string {
	if m != nil {
		return m.YangFrom
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoFile)(nil), "schema.ProtoFile")
	proto.RegisterType((*Schemas)(nil), "schema.Schemas")
}

func init() { proto.RegisterFile("voltha_protos/schema.proto", fileDescriptor_9d785a7d5fd3f7a9) }

var fileDescriptor_9d785a7d5fd3f7a9 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0xa5, 0x1b, 0x6e, 0x36, 0x9b, 0x0c, 0x83, 0x48, 0xe9, 0x44, 0xe6, 0x9e, 0xe6, 0xc3, 0x12,
	0x9c, 0x7e, 0x81, 0xb0, 0xf9, 0x26, 0xb2, 0x81, 0x0f, 0x3e, 0x38, 0xd2, 0xee, 0x2e, 0x0b, 0xb6,
	0x49, 0x49, 0x63, 0x65, 0xaf, 0xfe, 0x82, 0x9f, 0xe6, 0x2f, 0xf8, 0x21, 0xd2, 0xde, 0x4c, 0x7c,
	0xcb, 0x39, 0xe7, 0x5e, 0xee, 0x39, 0x27, 0x24, 0xae, 0x4c, 0xe6, 0x76, 0x62, 0x5d, 0x58, 0xe3,
	0x4c, 0xc9, 0xcb, 0x74, 0x07, 0xb9, 0x60, 0x0d, 0xa2, 0x1d, 0x44, 0xf1, 0x85, 0x34, 0x46, 0x66,
	0xc0, 0x45, 0xa1, 0xb8, 0xd0, 0xda, 0x38, 0xe1, 0x94, 0xd1, 0x25, 0x4e, 0xc5, 0x43, 0xaf, 0x36,
	0x28, 0x79, 0xdf, 0x72, 0xc8, 0x0b, 0xb7, 0x47, 0x71, 0xfc, 0x4a, 0xc2, 0xa7, 0xfa, 0xb1, 0x50,
	0x19, 0xd0, 0x21, 0x09, 0xb7, 0x2a, 0x83, 0xb5, 0x16, 0x39, 0x44, 0xc1, 0x28, 0x98, 0x84, 0xcb,
	0xe3, 0x9a, 0x78, 0x14, 0x39, 0xd0, 0x33, 0x72, 0xd4, 0xac, 0x44, 0xad, 0x46, 0x40, 0x40, 0x2f,
	0x09, 0xd9, 0x40, 0x99, 0x5a, 0x55, 0x38, 0x63, 0xa3, 0xf6, 0x28, 0x98, 0xf4, 0x97, 0xff, 0x98,
	0xb1, 0x23, 0xdd, 0x55, 0x63, 0xb2, 0xa4, 0xd7, 0xa4, 0x83, 0x21, 0xa2, 0x60, 0xd4, 0x9e, 0xf4,
	0x66, 0xa7, 0xcc, 0x87, 0xf9, 0x33, 0xb0, 0xf4, 0x03, 0xf4, 0x8a, 0xf4, 0xcb, 0x0f, 0x21, 0x25,
	0xd8, 0xf5, 0xd6, 0x9a, 0xdc, 0x9f, 0xec, 0x79, 0x6e, 0x61, 0x4d, 0x5e, 0x7b, 0xdd, 0x0b, 0x2d,
	0x51, 0x6f, 0xa3, 0xd7, 0x9a, 0xa8, 0xc5, 0xd9, 0x33, 0x39, 0xc1, 0xab, 0x2b, 0xb0, 0x95, 0x4a,
	0x81, 0xce, 0x49, 0xf8, 0x00, 0x0e, 0x39, 0x7a, 0xce, 0xb0, 0x11, 0x76, 0x68, 0x84, 0xcd, 0xeb,
	0x46, 0xe2, 0xc1, 0xc1, 0x90, 0x77, 0x3c, 0x1e, 0x7c, 0x7e, 0xff, 0x7c, 0xb5, 0x42, 0xda, 0xf5,
	0xb5, 0xdf, 0xdf, 0xbd, 0xcc, 0xa4, 0x29, 0xde, 0x24, 0x53, 0x9a, 0x6f, 0x12, 0xa1, 0x74, 0x62,
	0xd5, 0x34, 0x55, 0xa0, 0x05, 0xc7, 0x6f, 0x9a, 0x62, 0x00, 0x56, 0xdd, 0x70, 0x69, 0xfc, 0x56,
	0x82, 0xa9, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x49, 0xda, 0xb9, 0x25, 0xcb, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaServiceClient interface {
	// Return active grpc schemas
	GetSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Schemas, error)
}

type schemaServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchemaServiceClient(cc *grpc.ClientConn) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) GetSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Schemas, error) {
	out := new(Schemas)
	err := c.cc.Invoke(ctx, "/schema.SchemaService/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
type SchemaServiceServer interface {
	// Return active grpc schemas
	GetSchema(context.Context, *empty.Empty) (*Schemas, error)
}

func RegisterSchemaServiceServer(s *grpc.Server, srv SchemaServiceServer) {
	s.RegisterService(&_SchemaService_serviceDesc, srv)
}

func _SchemaService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.SchemaService/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchema(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _SchemaService_GetSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/schema.proto",
}
