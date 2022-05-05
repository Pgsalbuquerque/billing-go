// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: grpc/message_service.proto

package grpc

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

// SendTokenClient is the client API for SendToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendTokenClient interface {
	RequestToken(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Token, error)
}

type sendTokenClient struct {
	cc grpc.ClientConnInterface
}

func NewSendTokenClient(cc grpc.ClientConnInterface) SendTokenClient {
	return &sendTokenClient{cc}
}

func (c *sendTokenClient) RequestToken(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/SendToken/RequestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendTokenServer is the server API for SendToken service.
// All implementations must embed UnimplementedSendTokenServer
// for forward compatibility
type SendTokenServer interface {
	RequestToken(context.Context, *ID) (*Token, error)
	mustEmbedUnimplementedSendTokenServer()
}

// UnimplementedSendTokenServer must be embedded to have forward compatible implementations.
type UnimplementedSendTokenServer struct {
}

func (UnimplementedSendTokenServer) RequestToken(context.Context, *ID) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}
func (UnimplementedSendTokenServer) mustEmbedUnimplementedSendTokenServer() {}

// UnsafeSendTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendTokenServer will
// result in compilation errors.
type UnsafeSendTokenServer interface {
	mustEmbedUnimplementedSendTokenServer()
}

func RegisterSendTokenServer(s grpc.ServiceRegistrar, srv SendTokenServer) {
	s.RegisterService(&SendToken_ServiceDesc, srv)
}

func _SendToken_RequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendTokenServer).RequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendToken/RequestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendTokenServer).RequestToken(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// SendToken_ServiceDesc is the grpc.ServiceDesc for SendToken service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendToken_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SendToken",
	HandlerType: (*SendTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestToken",
			Handler:    _SendToken_RequestToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/message_service.proto",
}

// SendIDClient is the client API for SendID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendIDClient interface {
	RequestID(ctx context.Context, in *Token, opts ...grpc.CallOption) (*ID, error)
}

type sendIDClient struct {
	cc grpc.ClientConnInterface
}

func NewSendIDClient(cc grpc.ClientConnInterface) SendIDClient {
	return &sendIDClient{cc}
}

func (c *sendIDClient) RequestID(ctx context.Context, in *Token, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/SendID/RequestID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendIDServer is the server API for SendID service.
// All implementations must embed UnimplementedSendIDServer
// for forward compatibility
type SendIDServer interface {
	RequestID(context.Context, *Token) (*ID, error)
	mustEmbedUnimplementedSendIDServer()
}

// UnimplementedSendIDServer must be embedded to have forward compatible implementations.
type UnimplementedSendIDServer struct {
}

func (UnimplementedSendIDServer) RequestID(context.Context, *Token) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestID not implemented")
}
func (UnimplementedSendIDServer) mustEmbedUnimplementedSendIDServer() {}

// UnsafeSendIDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendIDServer will
// result in compilation errors.
type UnsafeSendIDServer interface {
	mustEmbedUnimplementedSendIDServer()
}

func RegisterSendIDServer(s grpc.ServiceRegistrar, srv SendIDServer) {
	s.RegisterService(&SendID_ServiceDesc, srv)
}

func _SendID_RequestID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendIDServer).RequestID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendID/RequestID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendIDServer).RequestID(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// SendID_ServiceDesc is the grpc.ServiceDesc for SendID service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendID_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SendID",
	HandlerType: (*SendIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestID",
			Handler:    _SendID_RequestID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/message_service.proto",
}

// SendValidateClient is the client API for SendValidate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendValidateClient interface {
	RequestValidate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Validate, error)
}

type sendValidateClient struct {
	cc grpc.ClientConnInterface
}

func NewSendValidateClient(cc grpc.ClientConnInterface) SendValidateClient {
	return &sendValidateClient{cc}
}

func (c *sendValidateClient) RequestValidate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Validate, error) {
	out := new(Validate)
	err := c.cc.Invoke(ctx, "/sendValidate/RequestValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendValidateServer is the server API for SendValidate service.
// All implementations must embed UnimplementedSendValidateServer
// for forward compatibility
type SendValidateServer interface {
	RequestValidate(context.Context, *Token) (*Validate, error)
	mustEmbedUnimplementedSendValidateServer()
}

// UnimplementedSendValidateServer must be embedded to have forward compatible implementations.
type UnimplementedSendValidateServer struct {
}

func (UnimplementedSendValidateServer) RequestValidate(context.Context, *Token) (*Validate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestValidate not implemented")
}
func (UnimplementedSendValidateServer) mustEmbedUnimplementedSendValidateServer() {}

// UnsafeSendValidateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendValidateServer will
// result in compilation errors.
type UnsafeSendValidateServer interface {
	mustEmbedUnimplementedSendValidateServer()
}

func RegisterSendValidateServer(s grpc.ServiceRegistrar, srv SendValidateServer) {
	s.RegisterService(&SendValidate_ServiceDesc, srv)
}

func _SendValidate_RequestValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendValidateServer).RequestValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sendValidate/RequestValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendValidateServer).RequestValidate(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// SendValidate_ServiceDesc is the grpc.ServiceDesc for SendValidate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendValidate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sendValidate",
	HandlerType: (*SendValidateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestValidate",
			Handler:    _SendValidate_RequestValidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/message_service.proto",
}
