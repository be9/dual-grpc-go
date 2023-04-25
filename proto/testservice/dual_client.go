package testservice

import (
	"context"
	"google.golang.org/grpc"
)

type dualClient struct {
	c1, c2 TestServiceClient
	// coinTosser returns false: c1 is used. true: c2 is used
	coinTosser func() bool
}

func NewDualClient(c1, c2 TestServiceClient, coinTosser func() bool) TestServiceClient {
	return &dualClient{c1: c1, c2: c2, coinTosser: coinTosser}
}

func (d *dualClient) GetFoos(ctx context.Context, in *GetFoosRequest, opts ...grpc.CallOption) (*GetFoosResponse, error) {
	if len(opts) > 0 || !d.coinTosser() {
		return d.c1.GetFoos(ctx, in, opts...)
	}
	return d.c2.GetFoos(ctx, in, opts...)
}

func (d *dualClient) GetFoobars(ctx context.Context, in *GetFoobarsRequest, opts ...grpc.CallOption) (*GetFoobarsResponse, error) {
	if len(opts) > 0 || !d.coinTosser() {
		return d.c1.GetFoobars(ctx, in, opts...)
	}
	return d.c2.GetFoobars(ctx, in, opts...)
}
