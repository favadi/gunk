// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/src/util/echo.gunk

/*
Package util is a generated protocol buffer package.

package util contains a simple Echo service.

It is generated from these files:
	testdata/src/util/echo.gunk
	testdata/src/util/types.gunk

It has these top-level messages:
	CheckStatusResponse
*/
package util

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import util_imp "testdata/src/util/imp"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty from public import google/protobuf/empty.proto
type Empty google_protobuf.Empty

func (m *Empty) Reset()         { (*google_protobuf.Empty)(m).Reset() }
func (m *Empty) String() string { return (*google_protobuf.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// CheckStatusResponse is the response for a check status.
type CheckStatusResponse struct {
	Status Status `protobuf:"varint,1,,name=Status,enum=util.Status" json:"Status,omitempty"`
}

func (m *CheckStatusResponse) Reset()                    { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()               {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckStatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

func init() {
	proto.RegisterType((*CheckStatusResponse)(nil), "util.CheckStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Util service

type UtilClient interface {
	// Echo echoes a message.
	Echo(ctx context.Context, in *util_imp.Message, opts ...grpc.CallOption) (*util_imp.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilClient struct {
	cc *grpc.ClientConn
}

func NewUtilClient(cc *grpc.ClientConn) UtilClient {
	return &utilClient{cc}
}

func (c *utilClient) Echo(ctx context.Context, in *util_imp.Message, opts ...grpc.CallOption) (*util_imp.Message, error) {
	out := new(util_imp.Message)
	err := grpc.Invoke(ctx, "/util.Util/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilClient) CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := grpc.Invoke(ctx, "/util.Util/CheckStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Util service

type UtilServer interface {
	// Echo echoes a message.
	Echo(context.Context, *util_imp.Message) (*util_imp.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(context.Context, *google_protobuf.Empty) (*CheckStatusResponse, error)
}

func RegisterUtilServer(s *grpc.Server, srv UtilServer) {
	s.RegisterService(&_Util_serviceDesc, srv)
}

func _Util_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(util_imp.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).Echo(ctx, req.(*util_imp.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Util_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).CheckStatus(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Util_serviceDesc = grpc.ServiceDesc{
	ServiceName: "util.Util",
	HandlerType: (*UtilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Util_Echo_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Util_CheckStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/src/util/echo.gunk",
}

func init() { proto.RegisterFile("testdata/src/util/echo.gunk", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x3d, 0x4b, 0x05, 0x31,
	0x10, 0x54, 0x38, 0xae, 0x88, 0x22, 0x18, 0x41, 0x34, 0x4f, 0x2c, 0xac, 0x2c, 0x64, 0x83, 0xcf,
	0x3f, 0x20, 0xca, 0x2b, 0x05, 0x51, 0x6c, 0xec, 0xf2, 0xe2, 0x9a, 0x3b, 0xee, 0x23, 0xe1, 0x76,
	0x23, 0xdc, 0xbf, 0x97, 0x24, 0x1c, 0x08, 0x67, 0x91, 0x62, 0x66, 0x98, 0x99, 0xcc, 0x8a, 0x0d,
	0x23, 0xf1, 0x97, 0x61, 0xa3, 0x69, 0xb2, 0x3a, 0x72, 0xdb, 0x6b, 0xb4, 0x8d, 0x07, 0x17, 0xc7,
	0x4e, 0x56, 0x09, 0xab, 0x8d, 0xf3, 0xde, 0xf5, 0xa8, 0xc3, 0xe4, 0xd9, 0xef, 0xe3, 0xb7, 0xc6,
	0x21, 0xf0, 0x0c, 0x19, 0xaa, 0xab, 0xb5, 0x9f, 0xe7, 0x80, 0x94, 0x03, 0xd4, 0xf5, 0x5a, 0x6d,
	0x87, 0x90, 0x5e, 0xd6, 0x6f, 0xee, 0xc5, 0xd9, 0x73, 0x83, 0xb6, 0x7b, 0x67, 0xc3, 0x91, 0xde,
	0x90, 0x82, 0x1f, 0x09, 0xa5, 0x12, 0x75, 0x61, 0x2e, 0x0e, 0x6f, 0x4f, 0xb6, 0xc7, 0x90, 0x8c,
	0x50, 0x98, 0xed, 0x8f, 0xa8, 0x3e, 0xb8, 0xed, 0xe5, 0x9d, 0xa8, 0x76, 0xb6, 0xf1, 0xf2, 0x14,
	0x96, 0x58, 0x78, 0x41, 0x22, 0xe3, 0x50, 0xad, 0x29, 0xf9, 0x28, 0x8e, 0xfe, 0x14, 0xc9, 0x73,
	0x28, 0x9b, 0x60, 0xd9, 0x04, 0xbb, 0xb4, 0x49, 0x5d, 0x96, 0xaa, 0x7f, 0xfe, 0xf4, 0x54, 0x7f,
	0xe6, 0x6b, 0xbc, 0x1e, 0xec, 0xeb, 0x6c, 0x7a, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x09,
	0x6d, 0xc4, 0x3b, 0x01, 0x00, 0x00,
}