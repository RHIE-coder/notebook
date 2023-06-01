// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: protobuf/middle-beat.proto

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

// MiddleBeatClient is the client API for MiddleBeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddleBeatClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type middleBeatClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddleBeatClient(cc grpc.ClientConnInterface) MiddleBeatClient {
	return &middleBeatClient{cc}
}

func (c *middleBeatClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/middlebeat.MiddleBeat/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleBeatClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/middlebeat.MiddleBeat/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddleBeatServer is the server API for MiddleBeat service.
// All implementations must embed UnimplementedMiddleBeatServer
// for forward compatibility
type MiddleBeatServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedMiddleBeatServer()
}

// UnimplementedMiddleBeatServer must be embedded to have forward compatible implementations.
type UnimplementedMiddleBeatServer struct {
}

func (UnimplementedMiddleBeatServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedMiddleBeatServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedMiddleBeatServer) mustEmbedUnimplementedMiddleBeatServer() {}

// UnsafeMiddleBeatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddleBeatServer will
// result in compilation errors.
type UnsafeMiddleBeatServer interface {
	mustEmbedUnimplementedMiddleBeatServer()
}

func RegisterMiddleBeatServer(s grpc.ServiceRegistrar, srv MiddleBeatServer) {
	s.RegisterService(&MiddleBeat_ServiceDesc, srv)
}

func _MiddleBeat_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleBeatServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/middlebeat.MiddleBeat/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleBeatServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiddleBeat_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleBeatServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/middlebeat.MiddleBeat/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleBeatServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MiddleBeat_ServiceDesc is the grpc.ServiceDesc for MiddleBeat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MiddleBeat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "middlebeat.MiddleBeat",
	HandlerType: (*MiddleBeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _MiddleBeat_Status_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _MiddleBeat_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/middle-beat.proto",
}
