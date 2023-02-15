// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bertybridge

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

// BridgeServiceClient is the client API for BridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeServiceClient interface {
	// ClientInvokeUnary invoke a unary method
	ClientInvokeUnary(ctx context.Context, in *ClientInvokeUnary_Request, opts ...grpc.CallOption) (*ClientInvokeUnary_Reply, error)
	// CreateStream create a stream
	CreateClientStream(ctx context.Context, in *ClientCreateStream_Request, opts ...grpc.CallOption) (*ClientCreateStream_Reply, error)
	// Send Message over the given stream
	ClientStreamSend(ctx context.Context, in *ClientStreamSend_Request, opts ...grpc.CallOption) (*ClientStreamSend_Reply, error)
	// Recv message over the given stream
	ClientStreamRecv(ctx context.Context, in *ClientStreamRecv_Request, opts ...grpc.CallOption) (*ClientStreamRecv_Reply, error)
	// Close the given stream (should not be used)
	ClientStreamClose(ctx context.Context, in *ClientStreamClose_Request, opts ...grpc.CallOption) (*ClientStreamClose_Reply, error)
	// Close the writing end of a stream and return the next reply
	ClientStreamCloseAndRecv(ctx context.Context, in *ClientStreamCloseAndRecv_Request, opts ...grpc.CallOption) (*ClientStreamCloseAndRecv_Reply, error)
}

type bridgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeServiceClient(cc grpc.ClientConnInterface) BridgeServiceClient {
	return &bridgeServiceClient{cc}
}

func (c *bridgeServiceClient) ClientInvokeUnary(ctx context.Context, in *ClientInvokeUnary_Request, opts ...grpc.CallOption) (*ClientInvokeUnary_Reply, error) {
	out := new(ClientInvokeUnary_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/ClientInvokeUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) CreateClientStream(ctx context.Context, in *ClientCreateStream_Request, opts ...grpc.CallOption) (*ClientCreateStream_Reply, error) {
	out := new(ClientCreateStream_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/CreateClientStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) ClientStreamSend(ctx context.Context, in *ClientStreamSend_Request, opts ...grpc.CallOption) (*ClientStreamSend_Reply, error) {
	out := new(ClientStreamSend_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/ClientStreamSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) ClientStreamRecv(ctx context.Context, in *ClientStreamRecv_Request, opts ...grpc.CallOption) (*ClientStreamRecv_Reply, error) {
	out := new(ClientStreamRecv_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/ClientStreamRecv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) ClientStreamClose(ctx context.Context, in *ClientStreamClose_Request, opts ...grpc.CallOption) (*ClientStreamClose_Reply, error) {
	out := new(ClientStreamClose_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/ClientStreamClose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) ClientStreamCloseAndRecv(ctx context.Context, in *ClientStreamCloseAndRecv_Request, opts ...grpc.CallOption) (*ClientStreamCloseAndRecv_Reply, error) {
	out := new(ClientStreamCloseAndRecv_Reply)
	err := c.cc.Invoke(ctx, "/berty.bridge.v1.BridgeService/ClientStreamCloseAndRecv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServiceServer is the server API for BridgeService service.
// All implementations must embed UnimplementedBridgeServiceServer
// for forward compatibility
type BridgeServiceServer interface {
	// ClientInvokeUnary invoke a unary method
	ClientInvokeUnary(context.Context, *ClientInvokeUnary_Request) (*ClientInvokeUnary_Reply, error)
	// CreateStream create a stream
	CreateClientStream(context.Context, *ClientCreateStream_Request) (*ClientCreateStream_Reply, error)
	// Send Message over the given stream
	ClientStreamSend(context.Context, *ClientStreamSend_Request) (*ClientStreamSend_Reply, error)
	// Recv message over the given stream
	ClientStreamRecv(context.Context, *ClientStreamRecv_Request) (*ClientStreamRecv_Reply, error)
	// Close the given stream (should not be used)
	ClientStreamClose(context.Context, *ClientStreamClose_Request) (*ClientStreamClose_Reply, error)
	// Close the writing end of a stream and return the next reply
	ClientStreamCloseAndRecv(context.Context, *ClientStreamCloseAndRecv_Request) (*ClientStreamCloseAndRecv_Reply, error)
	mustEmbedUnimplementedBridgeServiceServer()
}

// UnimplementedBridgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeServiceServer struct {
}

func (UnimplementedBridgeServiceServer) ClientInvokeUnary(context.Context, *ClientInvokeUnary_Request) (*ClientInvokeUnary_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientInvokeUnary not implemented")
}
func (UnimplementedBridgeServiceServer) CreateClientStream(context.Context, *ClientCreateStream_Request) (*ClientCreateStream_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientStream not implemented")
}
func (UnimplementedBridgeServiceServer) ClientStreamSend(context.Context, *ClientStreamSend_Request) (*ClientStreamSend_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientStreamSend not implemented")
}
func (UnimplementedBridgeServiceServer) ClientStreamRecv(context.Context, *ClientStreamRecv_Request) (*ClientStreamRecv_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientStreamRecv not implemented")
}
func (UnimplementedBridgeServiceServer) ClientStreamClose(context.Context, *ClientStreamClose_Request) (*ClientStreamClose_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientStreamClose not implemented")
}
func (UnimplementedBridgeServiceServer) ClientStreamCloseAndRecv(context.Context, *ClientStreamCloseAndRecv_Request) (*ClientStreamCloseAndRecv_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientStreamCloseAndRecv not implemented")
}
func (UnimplementedBridgeServiceServer) mustEmbedUnimplementedBridgeServiceServer() {}

// UnsafeBridgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeServiceServer will
// result in compilation errors.
type UnsafeBridgeServiceServer interface {
	mustEmbedUnimplementedBridgeServiceServer()
}

func RegisterBridgeServiceServer(s grpc.ServiceRegistrar, srv BridgeServiceServer) {
	s.RegisterService(&BridgeService_ServiceDesc, srv)
}

func _BridgeService_ClientInvokeUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInvokeUnary_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ClientInvokeUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/ClientInvokeUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ClientInvokeUnary(ctx, req.(*ClientInvokeUnary_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_CreateClientStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientCreateStream_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).CreateClientStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/CreateClientStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).CreateClientStream(ctx, req.(*ClientCreateStream_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_ClientStreamSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStreamSend_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ClientStreamSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/ClientStreamSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ClientStreamSend(ctx, req.(*ClientStreamSend_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_ClientStreamRecv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStreamRecv_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ClientStreamRecv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/ClientStreamRecv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ClientStreamRecv(ctx, req.(*ClientStreamRecv_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_ClientStreamClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStreamClose_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ClientStreamClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/ClientStreamClose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ClientStreamClose(ctx, req.(*ClientStreamClose_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_ClientStreamCloseAndRecv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStreamCloseAndRecv_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ClientStreamCloseAndRecv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.bridge.v1.BridgeService/ClientStreamCloseAndRecv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ClientStreamCloseAndRecv(ctx, req.(*ClientStreamCloseAndRecv_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// BridgeService_ServiceDesc is the grpc.ServiceDesc for BridgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "berty.bridge.v1.BridgeService",
	HandlerType: (*BridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientInvokeUnary",
			Handler:    _BridgeService_ClientInvokeUnary_Handler,
		},
		{
			MethodName: "CreateClientStream",
			Handler:    _BridgeService_CreateClientStream_Handler,
		},
		{
			MethodName: "ClientStreamSend",
			Handler:    _BridgeService_ClientStreamSend_Handler,
		},
		{
			MethodName: "ClientStreamRecv",
			Handler:    _BridgeService_ClientStreamRecv_Handler,
		},
		{
			MethodName: "ClientStreamClose",
			Handler:    _BridgeService_ClientStreamClose_Handler,
		},
		{
			MethodName: "ClientStreamCloseAndRecv",
			Handler:    _BridgeService_ClientStreamCloseAndRecv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bertybridge/bertybridge.proto",
}
