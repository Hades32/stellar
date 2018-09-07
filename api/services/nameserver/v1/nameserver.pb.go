// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto

package nameserver // import "github.com/ehazlett/stellar/api/services/nameserver/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RecordType int32

const (
	RecordType_UNKNOWN RecordType = 0
	RecordType_A       RecordType = 1
	RecordType_CNAME   RecordType = 2
	RecordType_MX      RecordType = 3
	RecordType_TXT     RecordType = 4
	RecordType_SRV     RecordType = 5
)

var RecordType_name = map[int32]string{
	0: "UNKNOWN",
	1: "A",
	2: "CNAME",
	3: "MX",
	4: "TXT",
	5: "SRV",
}
var RecordType_value = map[string]int32{
	"UNKNOWN": 0,
	"A":       1,
	"CNAME":   2,
	"MX":      3,
	"TXT":     4,
	"SRV":     5,
}

func (x RecordType) String() string {
	return proto.EnumName(RecordType_name, int32(x))
}
func (RecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{0}
}

type LookupRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (dst *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(dst, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Record struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=stellar.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Options              *types.Any `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{1}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (dst *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(dst, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Record) GetOptions() *types.Any {
	if m != nil {
		return m.Options
	}
	return nil
}

type LookupResponse struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{2}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (dst *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(dst, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LookupResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{3}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{4}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type CreateRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{5}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type DeleteRequest struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=stellar.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nameserver_44b666dd0ec1475c, []int{6}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "stellar.services.nameserver.v1.LookupRequest")
	proto.RegisterType((*Record)(nil), "stellar.services.nameserver.v1.Record")
	proto.RegisterType((*LookupResponse)(nil), "stellar.services.nameserver.v1.LookupResponse")
	proto.RegisterType((*ListRequest)(nil), "stellar.services.nameserver.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "stellar.services.nameserver.v1.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "stellar.services.nameserver.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.nameserver.v1.DeleteRequest")
	proto.RegisterEnum("stellar.services.nameserver.v1.RecordType", RecordType_name, RecordType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NameserverClient is the client API for Nameserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameserverClient interface {
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type nameserverClient struct {
	cc *grpc.ClientConn
}

func NewNameserverClient(cc *grpc.ClientConn) NameserverClient {
	return &nameserverClient{cc}
}

func (c *nameserverClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameserverServer is the server API for Nameserver service.
type NameserverServer interface {
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*types.Empty, error)
	Delete(context.Context, *DeleteRequest) (*types.Empty, error)
}

func RegisterNameserverServer(s *grpc.Server, srv NameserverServer) {
	s.RegisterService(&_Nameserver_serviceDesc, srv)
}

func _Nameserver_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nameserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.nameserver.v1.Nameserver",
	HandlerType: (*NameserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Nameserver_Lookup_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Nameserver_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Nameserver_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Nameserver_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto", fileDescriptor_nameserver_44b666dd0ec1475c)
}

var fileDescriptor_nameserver_44b666dd0ec1475c = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0xa4, 0x69, 0xca, 0xde, 0xda, 0x25, 0x0c, 0x45, 0x62, 0x04, 0x29, 0x01, 0xa5, 0xac,
	0x3a, 0x61, 0xeb, 0xa3, 0x20, 0xd6, 0x5a, 0x7c, 0x70, 0x37, 0x2b, 0xb1, 0x6a, 0xf1, 0x2d, 0xad,
	0x77, 0xb3, 0xc1, 0x34, 0x93, 0x4d, 0x26, 0x85, 0xf8, 0x39, 0x7e, 0xa6, 0x4f, 0x92, 0x99, 0xc4,
	0xb4, 0x8a, 0xdb, 0xb2, 0xd0, 0xb7, 0x39, 0xdc, 0x73, 0xee, 0x99, 0x9b, 0x73, 0x27, 0xf0, 0x2e,
	0x08, 0xf9, 0x75, 0xbe, 0xa0, 0x4b, 0xb6, 0x72, 0xf0, 0xda, 0xff, 0x11, 0x21, 0xe7, 0x4e, 0xc6,
	0x31, 0x8a, 0xfc, 0xd4, 0xf1, 0x93, 0xd0, 0xc9, 0x30, 0x5d, 0x87, 0x4b, 0xcc, 0x9c, 0xd8, 0x5f,
	0x61, 0x09, 0x30, 0x75, 0xd6, 0x67, 0x1b, 0x88, 0x26, 0x29, 0xe3, 0x8c, 0x3c, 0xaa, 0x44, 0xb4,
	0x16, 0xd0, 0x0d, 0xca, 0xfa, 0xcc, 0xea, 0x07, 0x2c, 0x60, 0x82, 0xea, 0x94, 0x27, 0xa9, 0xb2,
	0x1e, 0x06, 0x8c, 0x05, 0x11, 0x3a, 0x02, 0x2d, 0xf2, 0x2b, 0x07, 0x57, 0x09, 0x2f, 0xaa, 0xe2,
	0x83, 0xbf, 0x8b, 0x7e, 0x5c, 0x95, 0xec, 0xc7, 0xd0, 0x3b, 0x67, 0xec, 0x7b, 0x9e, 0x78, 0x78,
	0x93, 0x63, 0xc6, 0x49, 0x1f, 0xda, 0x37, 0x39, 0xa6, 0x85, 0xa9, 0x0c, 0x94, 0xe1, 0xb1, 0x27,
	0x81, 0xfd, 0x53, 0x01, 0xdd, 0xc3, 0x25, 0x4b, 0xbf, 0x91, 0x57, 0xa0, 0xf1, 0x22, 0x41, 0x51,
	0x3f, 0x19, 0x9d, 0xd2, 0xdb, 0xaf, 0x4b, 0xa5, 0x6a, 0x56, 0x24, 0xe8, 0x09, 0x1d, 0x21, 0xa0,
	0x95, 0x0c, 0x53, 0x15, 0xfd, 0xc5, 0xb9, 0x34, 0x5d, 0xfb, 0x51, 0x8e, 0x66, 0x4b, 0x9a, 0x0a,
	0x40, 0x28, 0x74, 0x58, 0xc2, 0x43, 0x16, 0x67, 0xa6, 0x36, 0x50, 0x86, 0xdd, 0x51, 0x9f, 0xca,
	0x41, 0x68, 0x3d, 0x08, 0x1d, 0xc7, 0x85, 0x57, 0x93, 0xec, 0x2b, 0x38, 0xa9, 0x67, 0xc9, 0x12,
	0x16, 0x67, 0x8d, 0x97, 0xb2, 0xe1, 0xf5, 0x1a, 0x3a, 0xa9, 0xb8, 0x53, 0x66, 0xaa, 0x83, 0xd6,
	0xb0, 0x3b, 0x7a, 0xb2, 0xdf, 0x08, 0x5e, 0x2d, 0xb3, 0x7b, 0xd0, 0x3d, 0x0f, 0x33, 0x5e, 0x7d,
	0x31, 0xfb, 0x03, 0xdc, 0x93, 0xb0, 0x32, 0xdd, 0x30, 0x50, 0xee, 0x66, 0x80, 0xd0, 0x9b, 0xa4,
	0xe8, 0x73, 0xac, 0x43, 0x39, 0xcc, 0x1c, 0x4b, 0xe8, 0xbd, 0xc5, 0x08, 0x1b, 0x9b, 0x03, 0x44,
	0x7b, 0x3a, 0x05, 0x68, 0x78, 0xa4, 0x0b, 0x9d, 0x4f, 0xee, 0x7b, 0xf7, 0xf2, 0x8b, 0x6b, 0x1c,
	0x91, 0x36, 0x28, 0x63, 0x43, 0x21, 0xc7, 0xd0, 0x9e, 0xb8, 0xe3, 0x8b, 0xa9, 0xa1, 0x12, 0x1d,
	0xd4, 0x8b, 0xb9, 0xd1, 0x22, 0x1d, 0x68, 0xcd, 0xe6, 0x33, 0x43, 0x2b, 0x0f, 0x1f, 0xbd, 0xcf,
	0x46, 0x7b, 0xf4, 0x4b, 0x05, 0x70, 0xff, 0xb8, 0x93, 0x00, 0x74, 0x19, 0x35, 0x79, 0xbe, 0xeb,
	0x96, 0x5b, 0xeb, 0x6d, 0xd1, 0x7d, 0xe9, 0x55, 0x98, 0x3e, 0x68, 0x65, 0xb8, 0xe4, 0xe9, 0x4e,
	0x5d, 0xb3, 0x11, 0xd6, 0xb3, 0xfd, 0xc8, 0x95, 0xc5, 0x25, 0xe8, 0x32, 0xed, 0xdd, 0xb3, 0x6c,
	0x6d, 0x85, 0x75, 0xff, 0x9f, 0xe7, 0x30, 0x2d, 0x1f, 0x7d, 0xd9, 0x50, 0xe6, 0xba, 0xbb, 0xe1,
	0x56, 0xfe, 0xff, 0x6b, 0xf8, 0x66, 0xf2, 0x75, 0x7c, 0xb7, 0xbf, 0xdb, 0xcb, 0x06, 0xcd, 0x8f,
	0x16, 0xba, 0x68, 0xfb, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x71, 0xe8, 0xc1, 0x2b,
	0x05, 0x00, 0x00,
}
