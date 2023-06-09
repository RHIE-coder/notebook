// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: protobuf/heartbeat.proto

package prouf

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

// HerBeaClient is the client API for HerBea service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HerBeaClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type herBeaClient struct {
	cc grpc.ClientConnInterface
}

func NewHerBeaClient(cc grpc.ClientConnInterface) HerBeaClient {
	return &herBeaClient{cc}
}

func (c *herBeaClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/heartbeat.HerBea/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *herBeaClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/heartbeat.HerBea/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HerBeaServer is the server API for HerBea service.
// All implementations must embed UnimplementedHerBeaServer
// for forward compatibility
type HerBeaServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedHerBeaServer()
}

// UnimplementedHerBeaServer must be embedded to have forward compatible implementations.
type UnimplementedHerBeaServer struct {
}

func (UnimplementedHerBeaServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedHerBeaServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedHerBeaServer) mustEmbedUnimplementedHerBeaServer() {}

// UnsafeHerBeaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HerBeaServer will
// result in compilation errors.
type UnsafeHerBeaServer interface {
	mustEmbedUnimplementedHerBeaServer()
}

func RegisterHerBeaServer(s grpc.ServiceRegistrar, srv HerBeaServer) {
	s.RegisterService(&HerBea_ServiceDesc, srv)
}

func _HerBea_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HerBeaServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat.HerBea/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HerBeaServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HerBea_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HerBeaServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat.HerBea/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HerBeaServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HerBea_ServiceDesc is the grpc.ServiceDesc for HerBea service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HerBea_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heartbeat.HerBea",
	HandlerType: (*HerBeaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _HerBea_Status_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _HerBea_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/heartbeat.proto",
}
