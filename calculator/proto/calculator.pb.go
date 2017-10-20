// Code generated by protoc-gen-go.
// source: calculator.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	calculator.proto

It has these top-level messages:
	SumRequest
	SumResponse
	MultiplyRequest
	MultiplyResponse
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

type SumRequest struct {
	Sequence []int64 `protobuf:"varint,1,rep,packed,name=sequence" json:"sequence,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto1.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SumRequest) GetSequence() []int64 {
	if m != nil {
		return m.Sequence
	}
	return nil
}

type SumResponse struct {
	Result int64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto1.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MultiplyRequest struct {
	Multiplicand int64 `protobuf:"varint,1,opt,name=multiplicand" json:"multiplicand,omitempty"`
	Multiplier   int64 `protobuf:"varint,2,opt,name=multiplier" json:"multiplier,omitempty"`
}

func (m *MultiplyRequest) Reset()                    { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string            { return proto1.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()               {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MultiplyRequest) GetMultiplicand() int64 {
	if m != nil {
		return m.Multiplicand
	}
	return 0
}

func (m *MultiplyRequest) GetMultiplier() int64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type MultiplyResponse struct {
	Product int64 `protobuf:"varint,1,opt,name=product" json:"product,omitempty"`
}

func (m *MultiplyResponse) Reset()                    { *m = MultiplyResponse{} }
func (m *MultiplyResponse) String() string            { return proto1.CompactTextString(m) }
func (*MultiplyResponse) ProtoMessage()               {}
func (*MultiplyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MultiplyResponse) GetProduct() int64 {
	if m != nil {
		return m.Product
	}
	return 0
}

func init() {
	proto1.RegisterType((*SumRequest)(nil), "proto.SumRequest")
	proto1.RegisterType((*SumResponse)(nil), "proto.SumResponse")
	proto1.RegisterType((*MultiplyRequest)(nil), "proto.MultiplyRequest")
	proto1.RegisterType((*MultiplyResponse)(nil), "proto.MultiplyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalculatorService service

type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/proto.CalculatorService/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := grpc.Invoke(ctx, "/proto.CalculatorService/Multiply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculatorService service

type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CalculatorService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}

func init() { proto1.RegisterFile("calculator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x31, 0x16, 0xa5, 0x3a, 0x90, 0x68, 0x6f, 0x28, 0x56, 0x06, 0x54, 0x59, 0x42, 0xca,
	0x80, 0x32, 0xc0, 0xcc, 0xc4, 0xcc, 0x92, 0x88, 0x0f, 0x10, 0xdc, 0x1b, 0x22, 0x39, 0xb1, 0xf1,
	0x1f, 0x24, 0x56, 0x3e, 0x39, 0xc2, 0xb1, 0xdb, 0x42, 0xa7, 0xe4, 0xfd, 0xfc, 0xee, 0xf9, 0x9e,
	0x61, 0xa5, 0x7a, 0xad, 0xa2, 0xee, 0x83, 0x71, 0x8d, 0x75, 0x26, 0x18, 0xbc, 0x48, 0x1f, 0x59,
	0x03, 0x74, 0x71, 0x6c, 0xe9, 0x23, 0x92, 0x0f, 0x58, 0xc1, 0xd2, 0xff, 0xfe, 0x4e, 0x8a, 0x04,
	0xdb, 0xf2, 0x9a, 0xb7, 0x7b, 0x2d, 0xef, 0xe1, 0x2a, 0x39, 0xbd, 0x35, 0x93, 0x27, 0xdc, 0xc0,
	0xc2, 0x91, 0x8f, 0x3a, 0x08, 0xb6, 0x65, 0x35, 0x6f, 0xb3, 0x92, 0x6f, 0x70, 0xf3, 0x1a, 0x75,
	0x18, 0xac, 0xfe, 0x2a, 0xa9, 0x12, 0xae, 0xc7, 0x19, 0x0d, 0xaa, 0x9f, 0x76, 0x79, 0xe0, 0x0f,
	0xc3, 0x3b, 0x80, 0xa2, 0xc9, 0x89, 0xf3, 0xe4, 0x38, 0x22, 0xf2, 0x01, 0x56, 0x87, 0xd8, 0xbc,
	0x82, 0x80, 0x4b, 0xeb, 0xcc, 0x2e, 0xaa, 0xb2, 0x43, 0x91, 0x8f, 0xdf, 0x0c, 0xd6, 0x2f, 0xfb,
	0xc6, 0x1d, 0xb9, 0xcf, 0x41, 0x11, 0x36, 0xc0, 0xbb, 0x38, 0xe2, 0x7a, 0x7e, 0x81, 0xe6, 0xd0,
	0xbb, 0xc2, 0x63, 0x34, 0xa7, 0xcb, 0x33, 0x7c, 0x86, 0x65, 0xb9, 0x13, 0x37, 0xd9, 0xf1, 0xaf,
	0x5b, 0x75, 0x7b, 0xc2, 0xcb, 0xf8, 0xfb, 0x22, 0x9d, 0x3c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x88, 0x7f, 0x53, 0xa9, 0x7c, 0x01, 0x00, 0x00,
}
