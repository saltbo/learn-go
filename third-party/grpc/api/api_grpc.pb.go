// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RsLocaldClient is the client API for RsLocald service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RsLocaldClient interface {
	Login(ctx context.Context, in *LoginBody, opts ...grpc.CallOption) (*LoginReply, error)
	Listen(ctx context.Context, opts ...grpc.CallOption) (RsLocald_ListenClient, error)
	SendResponse(ctx context.Context, in *ProxyResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rsLocaldClient struct {
	cc grpc.ClientConnInterface
}

func NewRsLocaldClient(cc grpc.ClientConnInterface) RsLocaldClient {
	return &rsLocaldClient{cc}
}

func (c *rsLocaldClient) Login(ctx context.Context, in *LoginBody, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.RsLocald/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rsLocaldClient) Listen(ctx context.Context, opts ...grpc.CallOption) (RsLocald_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &RsLocald_ServiceDesc.Streams[0], "/api.RsLocald/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &rsLocaldListenClient{stream}
	return x, nil
}

type RsLocald_ListenClient interface {
	Send(*ProxyResponse) error
	Recv() (*ProxyRequest, error)
	grpc.ClientStream
}

type rsLocaldListenClient struct {
	grpc.ClientStream
}

func (x *rsLocaldListenClient) Send(m *ProxyResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rsLocaldListenClient) Recv() (*ProxyRequest, error) {
	m := new(ProxyRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rsLocaldClient) SendResponse(ctx context.Context, in *ProxyResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.RsLocald/SendResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RsLocaldServer is the server API for RsLocald service.
// All implementations must embed UnimplementedRsLocaldServer
// for forward compatibility
type RsLocaldServer interface {
	Login(context.Context, *LoginBody) (*LoginReply, error)
	Listen(RsLocald_ListenServer) error
	SendResponse(context.Context, *ProxyResponse) (*emptypb.Empty, error)
	mustEmbedUnimplementedRsLocaldServer()
}

// UnimplementedRsLocaldServer must be embedded to have forward compatible implementations.
type UnimplementedRsLocaldServer struct {
}

func (UnimplementedRsLocaldServer) Login(context.Context, *LoginBody) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRsLocaldServer) Listen(RsLocald_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedRsLocaldServer) SendResponse(context.Context, *ProxyResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResponse not implemented")
}
func (UnimplementedRsLocaldServer) mustEmbedUnimplementedRsLocaldServer() {}

// UnsafeRsLocaldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RsLocaldServer will
// result in compilation errors.
type UnsafeRsLocaldServer interface {
	mustEmbedUnimplementedRsLocaldServer()
}

func RegisterRsLocaldServer(s grpc.ServiceRegistrar, srv RsLocaldServer) {
	s.RegisterService(&RsLocald_ServiceDesc, srv)
}

func _RsLocald_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsLocaldServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RsLocald/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsLocaldServer).Login(ctx, req.(*LoginBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _RsLocald_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RsLocaldServer).Listen(&rsLocaldListenServer{stream})
}

type RsLocald_ListenServer interface {
	Send(*ProxyRequest) error
	Recv() (*ProxyResponse, error)
	grpc.ServerStream
}

type rsLocaldListenServer struct {
	grpc.ServerStream
}

func (x *rsLocaldListenServer) Send(m *ProxyRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rsLocaldListenServer) Recv() (*ProxyResponse, error) {
	m := new(ProxyResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RsLocald_SendResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsLocaldServer).SendResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RsLocald/SendResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsLocaldServer).SendResponse(ctx, req.(*ProxyResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// RsLocald_ServiceDesc is the grpc.ServiceDesc for RsLocald service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RsLocald_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RsLocald",
	HandlerType: (*RsLocaldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RsLocald_Login_Handler,
		},
		{
			MethodName: "SendResponse",
			Handler:    _RsLocald_SendResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _RsLocald_Listen_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
