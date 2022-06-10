// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/system/idp.proto

package rpcv3

import (
	context "context"
	v3 "github.com/paralus/paralus/proto/types/systempb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IdpClient is the client API for Idp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdpClient interface {
	CreateIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error)
	GetIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error)
	ListIdps(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v3.IdpList, error)
	UpdateIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error)
	DeleteIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type idpClient struct {
	cc grpc.ClientConnInterface
}

func NewIdpClient(cc grpc.ClientConnInterface) IdpClient {
	return &idpClient{cc}
}

func (c *idpClient) CreateIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error) {
	out := new(v3.Idp)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.v3.Idp/CreateIdp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idpClient) GetIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error) {
	out := new(v3.Idp)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.v3.Idp/GetIdp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idpClient) ListIdps(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v3.IdpList, error) {
	out := new(v3.IdpList)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.v3.Idp/ListIdps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idpClient) UpdateIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*v3.Idp, error) {
	out := new(v3.Idp)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.v3.Idp/UpdateIdp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idpClient) DeleteIdp(ctx context.Context, in *v3.Idp, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.v3.Idp/DeleteIdp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdpServer is the server API for Idp service.
// All implementations should embed UnimplementedIdpServer
// for forward compatibility
type IdpServer interface {
	CreateIdp(context.Context, *v3.Idp) (*v3.Idp, error)
	GetIdp(context.Context, *v3.Idp) (*v3.Idp, error)
	ListIdps(context.Context, *emptypb.Empty) (*v3.IdpList, error)
	UpdateIdp(context.Context, *v3.Idp) (*v3.Idp, error)
	DeleteIdp(context.Context, *v3.Idp) (*emptypb.Empty, error)
}

// UnimplementedIdpServer should be embedded to have forward compatible implementations.
type UnimplementedIdpServer struct {
}

func (UnimplementedIdpServer) CreateIdp(context.Context, *v3.Idp) (*v3.Idp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdp not implemented")
}
func (UnimplementedIdpServer) GetIdp(context.Context, *v3.Idp) (*v3.Idp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdp not implemented")
}
func (UnimplementedIdpServer) ListIdps(context.Context, *emptypb.Empty) (*v3.IdpList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIdps not implemented")
}
func (UnimplementedIdpServer) UpdateIdp(context.Context, *v3.Idp) (*v3.Idp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIdp not implemented")
}
func (UnimplementedIdpServer) DeleteIdp(context.Context, *v3.Idp) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIdp not implemented")
}

// UnsafeIdpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdpServer will
// result in compilation errors.
type UnsafeIdpServer interface {
	mustEmbedUnimplementedIdpServer()
}

func RegisterIdpServer(s grpc.ServiceRegistrar, srv IdpServer) {
	s.RegisterService(&Idp_ServiceDesc, srv)
}

func _Idp_CreateIdp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Idp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServer).CreateIdp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.v3.Idp/CreateIdp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServer).CreateIdp(ctx, req.(*v3.Idp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idp_GetIdp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Idp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServer).GetIdp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.v3.Idp/GetIdp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServer).GetIdp(ctx, req.(*v3.Idp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idp_ListIdps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServer).ListIdps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.v3.Idp/ListIdps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServer).ListIdps(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idp_UpdateIdp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Idp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServer).UpdateIdp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.v3.Idp/UpdateIdp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServer).UpdateIdp(ctx, req.(*v3.Idp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idp_DeleteIdp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Idp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServer).DeleteIdp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.v3.Idp/DeleteIdp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServer).DeleteIdp(ctx, req.(*v3.Idp))
	}
	return interceptor(ctx, in, info, handler)
}

// Idp_ServiceDesc is the grpc.ServiceDesc for Idp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Idp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.v3.Idp",
	HandlerType: (*IdpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdp",
			Handler:    _Idp_CreateIdp_Handler,
		},
		{
			MethodName: "GetIdp",
			Handler:    _Idp_GetIdp_Handler,
		},
		{
			MethodName: "ListIdps",
			Handler:    _Idp_ListIdps_Handler,
		},
		{
			MethodName: "UpdateIdp",
			Handler:    _Idp_UpdateIdp_Handler,
		},
		{
			MethodName: "DeleteIdp",
			Handler:    _Idp_DeleteIdp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/idp.proto",
}
