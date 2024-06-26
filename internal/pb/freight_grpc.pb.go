// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/freight.proto

package pb

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

// FreightServiceClient is the client API for FreightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightServiceClient interface {
	CalculateFreight(ctx context.Context, in *CalculateFreightRequest, opts ...grpc.CallOption) (*CalculateFreightResponse, error)
}

type freightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightServiceClient(cc grpc.ClientConnInterface) FreightServiceClient {
	return &freightServiceClient{cc}
}

func (c *freightServiceClient) CalculateFreight(ctx context.Context, in *CalculateFreightRequest, opts ...grpc.CallOption) (*CalculateFreightResponse, error) {
	out := new(CalculateFreightResponse)
	err := c.cc.Invoke(ctx, "/pb.FreightService/CalculateFreight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightServiceServer is the server API for FreightService service.
// All implementations must embed UnimplementedFreightServiceServer
// for forward compatibility
type FreightServiceServer interface {
	CalculateFreight(context.Context, *CalculateFreightRequest) (*CalculateFreightResponse, error)
	mustEmbedUnimplementedFreightServiceServer()
}

// UnimplementedFreightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightServiceServer struct {
}

func (UnimplementedFreightServiceServer) CalculateFreight(context.Context, *CalculateFreightRequest) (*CalculateFreightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateFreight not implemented")
}
func (UnimplementedFreightServiceServer) mustEmbedUnimplementedFreightServiceServer() {}

// UnsafeFreightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightServiceServer will
// result in compilation errors.
type UnsafeFreightServiceServer interface {
	mustEmbedUnimplementedFreightServiceServer()
}

func RegisterFreightServiceServer(s grpc.ServiceRegistrar, srv FreightServiceServer) {
	s.RegisterService(&FreightService_ServiceDesc, srv)
}

func _FreightService_CalculateFreight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateFreightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).CalculateFreight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FreightService/CalculateFreight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).CalculateFreight(ctx, req.(*CalculateFreightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightService_ServiceDesc is the grpc.ServiceDesc for FreightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FreightService",
	HandlerType: (*FreightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateFreight",
			Handler:    _FreightService_CalculateFreight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/freight.proto",
}
