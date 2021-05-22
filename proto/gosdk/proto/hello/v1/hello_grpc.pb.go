// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hellov1

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Try(ctx context.Context, in *TryRequest, opts ...grpc.CallOption) (*TryResponse, error)
	TryContext(ctx context.Context, in *TryContextRequest, opts ...grpc.CallOption) (*TryContextResponse, error)
	TryTimeout(ctx context.Context, in *TryTimeoutRequest, opts ...grpc.CallOption) (*TryTimeoutResponse, error)
	Slow(ctx context.Context, in *SlowRequest, opts ...grpc.CallOption) (*SlowResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.hello.v1.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Try(ctx context.Context, in *TryRequest, opts ...grpc.CallOption) (*TryResponse, error) {
	out := new(TryResponse)
	err := c.cc.Invoke(ctx, "/proto.hello.v1.HelloService/Try", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) TryContext(ctx context.Context, in *TryContextRequest, opts ...grpc.CallOption) (*TryContextResponse, error) {
	out := new(TryContextResponse)
	err := c.cc.Invoke(ctx, "/proto.hello.v1.HelloService/TryContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) TryTimeout(ctx context.Context, in *TryTimeoutRequest, opts ...grpc.CallOption) (*TryTimeoutResponse, error) {
	out := new(TryTimeoutResponse)
	err := c.cc.Invoke(ctx, "/proto.hello.v1.HelloService/TryTimeout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Slow(ctx context.Context, in *SlowRequest, opts ...grpc.CallOption) (*SlowResponse, error) {
	out := new(SlowResponse)
	err := c.cc.Invoke(ctx, "/proto.hello.v1.HelloService/Slow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	Try(context.Context, *TryRequest) (*TryResponse, error)
	TryContext(context.Context, *TryContextRequest) (*TryContextResponse, error)
	TryTimeout(context.Context, *TryTimeoutRequest) (*TryTimeoutResponse, error)
	Slow(context.Context, *SlowRequest) (*SlowResponse, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) Try(context.Context, *TryRequest) (*TryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Try not implemented")
}
func (UnimplementedHelloServiceServer) TryContext(context.Context, *TryContextRequest) (*TryContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryContext not implemented")
}
func (UnimplementedHelloServiceServer) TryTimeout(context.Context, *TryTimeoutRequest) (*TryTimeoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryTimeout not implemented")
}
func (UnimplementedHelloServiceServer) Slow(context.Context, *SlowRequest) (*SlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Slow not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.hello.v1.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Try_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Try(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.hello.v1.HelloService/Try",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Try(ctx, req.(*TryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_TryContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TryContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).TryContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.hello.v1.HelloService/TryContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).TryContext(ctx, req.(*TryContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_TryTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TryTimeoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).TryTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.hello.v1.HelloService/TryTimeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).TryTimeout(ctx, req.(*TryTimeoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Slow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Slow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.hello.v1.HelloService/Slow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Slow(ctx, req.(*SlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.hello.v1.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
		{
			MethodName: "Try",
			Handler:    _HelloService_Try_Handler,
		},
		{
			MethodName: "TryContext",
			Handler:    _HelloService_TryContext_Handler,
		},
		{
			MethodName: "TryTimeout",
			Handler:    _HelloService_TryTimeout_Handler,
		},
		{
			MethodName: "Slow",
			Handler:    _HelloService_Slow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello/v1/hello.proto",
}
