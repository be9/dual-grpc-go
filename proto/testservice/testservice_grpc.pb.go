// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: proto/testservice/testservice.proto

package testservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TestService_GetFoos_FullMethodName    = "/sandcastle.TestService/GetFoos"
	TestService_GetFoobars_FullMethodName = "/sandcastle.TestService/GetFoobars"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	GetFoos(ctx context.Context, in *GetFoosRequest, opts ...grpc.CallOption) (*GetFoosResponse, error)
	GetFoobars(ctx context.Context, in *GetFoobarsRequest, opts ...grpc.CallOption) (*GetFoobarsResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetFoos(ctx context.Context, in *GetFoosRequest, opts ...grpc.CallOption) (*GetFoosResponse, error) {
	out := new(GetFoosResponse)
	err := c.cc.Invoke(ctx, TestService_GetFoos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetFoobars(ctx context.Context, in *GetFoobarsRequest, opts ...grpc.CallOption) (*GetFoobarsResponse, error) {
	out := new(GetFoobarsResponse)
	err := c.cc.Invoke(ctx, TestService_GetFoobars_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	GetFoos(context.Context, *GetFoosRequest) (*GetFoosResponse, error)
	GetFoobars(context.Context, *GetFoobarsRequest) (*GetFoobarsResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) GetFoos(context.Context, *GetFoosRequest) (*GetFoosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoos not implemented")
}
func (UnimplementedTestServiceServer) GetFoobars(context.Context, *GetFoobarsRequest) (*GetFoobarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoobars not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_GetFoos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFoosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetFoos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetFoos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetFoos(ctx, req.(*GetFoosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetFoobars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFoobarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetFoobars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetFoobars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetFoobars(ctx, req.(*GetFoobarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sandcastle.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFoos",
			Handler:    _TestService_GetFoos_Handler,
		},
		{
			MethodName: "GetFoobars",
			Handler:    _TestService_GetFoobars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/testservice/testservice.proto",
}