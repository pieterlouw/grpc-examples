// Code generated by protoc-gen-go.
// source: strings.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	strings.proto

It has these top-level messages:
	ReverseRequest
	ReverseResponse
	ConcatRequest
	ConcatResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ReverseRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *ReverseRequest) Reset()                    { *m = ReverseRequest{} }
func (m *ReverseRequest) String() string            { return proto1.CompactTextString(m) }
func (*ReverseRequest) ProtoMessage()               {}
func (*ReverseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReverseRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReverseResponse struct {
	Reversed string `protobuf:"bytes,1,opt,name=reversed" json:"reversed,omitempty"`
}

func (m *ReverseResponse) Reset()                    { *m = ReverseResponse{} }
func (m *ReverseResponse) String() string            { return proto1.CompactTextString(m) }
func (*ReverseResponse) ProtoMessage()               {}
func (*ReverseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReverseResponse) GetReversed() string {
	if m != nil {
		return m.Reversed
	}
	return ""
}

type ConcatRequest struct {
	Value1 string `protobuf:"bytes,1,opt,name=value1" json:"value1,omitempty"`
	Value2 string `protobuf:"bytes,2,opt,name=value2" json:"value2,omitempty"`
}

func (m *ConcatRequest) Reset()                    { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string            { return proto1.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()               {}
func (*ConcatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConcatRequest) GetValue1() string {
	if m != nil {
		return m.Value1
	}
	return ""
}

func (m *ConcatRequest) GetValue2() string {
	if m != nil {
		return m.Value2
	}
	return ""
}

type ConcatResponse struct {
	Concatenated string `protobuf:"bytes,1,opt,name=concatenated" json:"concatenated,omitempty"`
}

func (m *ConcatResponse) Reset()                    { *m = ConcatResponse{} }
func (m *ConcatResponse) String() string            { return proto1.CompactTextString(m) }
func (*ConcatResponse) ProtoMessage()               {}
func (*ConcatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConcatResponse) GetConcatenated() string {
	if m != nil {
		return m.Concatenated
	}
	return ""
}

func init() {
	proto1.RegisterType((*ReverseRequest)(nil), "proto.ReverseRequest")
	proto1.RegisterType((*ReverseResponse)(nil), "proto.ReverseResponse")
	proto1.RegisterType((*ConcatRequest)(nil), "proto.ConcatRequest")
	proto1.RegisterType((*ConcatResponse)(nil), "proto.ConcatResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StringsService service

type StringsServiceClient interface {
	Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error)
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatResponse, error)
}

type stringsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStringsServiceClient(cc *grpc.ClientConn) StringsServiceClient {
	return &stringsServiceClient{cc}
}

func (c *stringsServiceClient) Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error) {
	out := new(ReverseResponse)
	err := grpc.Invoke(ctx, "/proto.StringsService/Reverse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringsServiceClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatResponse, error) {
	out := new(ConcatResponse)
	err := grpc.Invoke(ctx, "/proto.StringsService/Concat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StringsService service

type StringsServiceServer interface {
	Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error)
	Concat(context.Context, *ConcatRequest) (*ConcatResponse, error)
}

func RegisterStringsServiceServer(s *grpc.Server, srv StringsServiceServer) {
	s.RegisterService(&_StringsService_serviceDesc, srv)
}

func _StringsService_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringsServiceServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StringsService/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringsServiceServer).Reverse(ctx, req.(*ReverseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringsService_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringsServiceServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StringsService/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringsServiceServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StringsService",
	HandlerType: (*StringsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reverse",
			Handler:    _StringsService_Reverse_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _StringsService_Concat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strings.proto",
}

func init() { proto1.RegisterFile("strings.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x8a, 0x83, 0x30,
	0x18, 0x84, 0xd7, 0x05, 0xdd, 0xdd, 0x9f, 0xd5, 0x85, 0xa0, 0x22, 0x9e, 0x96, 0x1c, 0x4a, 0x2f,
	0x15, 0x6a, 0x0b, 0x85, 0x5e, 0x7a, 0xe8, 0x1b, 0xe8, 0x13, 0x58, 0xfb, 0x53, 0x84, 0x92, 0xd8,
	0x24, 0xfa, 0x06, 0x7d, 0xef, 0x42, 0x12, 0x43, 0xf5, 0x14, 0xe6, 0x63, 0x32, 0xf3, 0x0f, 0x84,
	0x52, 0x89, 0x8e, 0xdd, 0x64, 0xd1, 0x0b, 0xae, 0x38, 0xf1, 0xf5, 0x43, 0x57, 0x10, 0x55, 0x38,
	0xa2, 0x90, 0x58, 0xe1, 0x63, 0x40, 0xa9, 0x48, 0x0c, 0xfe, 0xd8, 0xdc, 0x07, 0xcc, 0xbc, 0x7f,
	0x6f, 0xfd, 0x53, 0x19, 0x41, 0x37, 0xf0, 0xe7, 0x7c, 0xb2, 0xe7, 0x4c, 0x22, 0xc9, 0xe1, 0x5b,
	0x18, 0x74, 0xb5, 0x5e, 0xa7, 0xe9, 0x09, 0xc2, 0x33, 0x67, 0x6d, 0xa3, 0xa6, 0xd4, 0x14, 0x02,
	0x1d, 0xb4, 0xb5, 0x56, 0xab, 0x1c, 0x2f, 0xb3, 0xcf, 0x37, 0x5e, 0xd2, 0x3d, 0x44, 0x53, 0x80,
	0xad, 0xa3, 0xf0, 0xdb, 0x6a, 0x82, 0xac, 0x51, 0xae, 0x72, 0xc6, 0xca, 0xa7, 0x07, 0x51, 0x6d,
	0x66, 0xd6, 0x28, 0xc6, 0xae, 0x45, 0x72, 0x84, 0x2f, 0x7b, 0x38, 0x49, 0xcc, 0xf4, 0x62, 0x3e,
	0x38, 0x4f, 0x97, 0xd8, 0x14, 0xd2, 0x0f, 0x72, 0x80, 0xc0, 0x1c, 0x41, 0x62, 0xeb, 0x99, 0x8d,
	0xca, 0x93, 0x05, 0x9d, 0x3e, 0x5e, 0x02, 0xcd, 0x77, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36,
	0x0e, 0x4d, 0x2a, 0x74, 0x01, 0x00, 0x00,
}
