// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: v2/hcore/tunnelservice/tunnel_service.proto

package tunnelservice

import (
	context "context"
	common "github.com/hiddify/hiddify-core/v2/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TunnelService_Start_FullMethodName  = "/tunnelservice.TunnelService/Start"
	TunnelService_Stop_FullMethodName   = "/tunnelservice.TunnelService/Stop"
	TunnelService_Status_FullMethodName = "/tunnelservice.TunnelService/Status"
	TunnelService_Exit_FullMethodName   = "/tunnelservice.TunnelService/Exit"
)

// TunnelServiceClient is the client API for TunnelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TunnelServiceClient interface {
	Start(ctx context.Context, in *TunnelStartRequest, opts ...grpc.CallOption) (*TunnelResponse, error)
	Stop(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
	Status(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
	Exit(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
}

type tunnelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelServiceClient(cc grpc.ClientConnInterface) TunnelServiceClient {
	return &tunnelServiceClient{cc}
}

func (c *tunnelServiceClient) Start(ctx context.Context, in *TunnelStartRequest, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Stop(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Status(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Exit(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Exit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TunnelServiceServer is the server API for TunnelService service.
// All implementations must embed UnimplementedTunnelServiceServer
// for forward compatibility.
type TunnelServiceServer interface {
	Start(context.Context, *TunnelStartRequest) (*TunnelResponse, error)
	Stop(context.Context, *common.Empty) (*TunnelResponse, error)
	Status(context.Context, *common.Empty) (*TunnelResponse, error)
	Exit(context.Context, *common.Empty) (*TunnelResponse, error)
	mustEmbedUnimplementedTunnelServiceServer()
}

// UnimplementedTunnelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTunnelServiceServer struct{}

func (UnimplementedTunnelServiceServer) Start(context.Context, *TunnelStartRequest) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedTunnelServiceServer) Stop(context.Context, *common.Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTunnelServiceServer) Status(context.Context, *common.Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTunnelServiceServer) Exit(context.Context, *common.Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedTunnelServiceServer) mustEmbedUnimplementedTunnelServiceServer() {}
func (UnimplementedTunnelServiceServer) testEmbeddedByValue()                       {}

// UnsafeTunnelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TunnelServiceServer will
// result in compilation errors.
type UnsafeTunnelServiceServer interface {
	mustEmbedUnimplementedTunnelServiceServer()
}

func RegisterTunnelServiceServer(s grpc.ServiceRegistrar, srv TunnelServiceServer) {
	// If the following call pancis, it indicates UnimplementedTunnelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TunnelService_ServiceDesc, srv)
}

func _TunnelService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TunnelStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Start(ctx, req.(*TunnelStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Stop(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Status(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Exit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Exit(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TunnelService_ServiceDesc is the grpc.ServiceDesc for TunnelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TunnelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tunnelservice.TunnelService",
	HandlerType: (*TunnelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _TunnelService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TunnelService_Stop_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TunnelService_Status_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _TunnelService_Exit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2/hcore/tunnelservice/tunnel_service.proto",
}