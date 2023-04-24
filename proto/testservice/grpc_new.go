package testservice

import (
	"context"
	grpcold "google.golang.org/grpc"
	grpc "google.golang.org/grpc/v2"
)

type testServiceClient2 struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient2(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient2{cc}
}

func (c *testServiceClient2) GetFoos(ctx context.Context, in *GetFoosRequest, opts ...grpcold.CallOption) (*GetFoosResponse, error) {
	out := new(GetFoosResponse)
	err := c.cc.Invoke(ctx, TestService_GetFoos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient2) GetFoobars(ctx context.Context, in *GetFoobarsRequest, opts ...grpcold.CallOption) (*GetFoobarsResponse, error) {
	out := new(GetFoobarsResponse)
	err := c.cc.Invoke(ctx, TestService_GetFoobars_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
