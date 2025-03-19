// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: devpod/v1/prebuild.proto

package v1

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

// PrebuildServiceClient is the client API for PrebuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrebuildServiceClient interface {
	StartPrebuild(ctx context.Context, in *StartPrebuildRequest, opts ...grpc.CallOption) (*StartPrebuildResponse, error)
	CancelPrebuild(ctx context.Context, in *CancelPrebuildRequest, opts ...grpc.CallOption) (*CancelPrebuildResponse, error)
	GetPrebuild(ctx context.Context, in *GetPrebuildRequest, opts ...grpc.CallOption) (*GetPrebuildResponse, error)
	ListPrebuilds(ctx context.Context, in *ListPrebuildsRequest, opts ...grpc.CallOption) (*ListPrebuildsResponse, error)
	WatchPrebuild(ctx context.Context, in *WatchPrebuildRequest, opts ...grpc.CallOption) (PrebuildService_WatchPrebuildClient, error)
	// ListOrganizationPrebuilds lists all prebuilds of an organization
	ListOrganizationPrebuilds(ctx context.Context, in *ListOrganizationPrebuildsRequest, opts ...grpc.CallOption) (*ListOrganizationPrebuildsResponse, error)
}

type prebuildServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrebuildServiceClient(cc grpc.ClientConnInterface) PrebuildServiceClient {
	return &prebuildServiceClient{cc}
}

func (c *prebuildServiceClient) StartPrebuild(ctx context.Context, in *StartPrebuildRequest, opts ...grpc.CallOption) (*StartPrebuildResponse, error) {
	out := new(StartPrebuildResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.PrebuildService/StartPrebuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prebuildServiceClient) CancelPrebuild(ctx context.Context, in *CancelPrebuildRequest, opts ...grpc.CallOption) (*CancelPrebuildResponse, error) {
	out := new(CancelPrebuildResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.PrebuildService/CancelPrebuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prebuildServiceClient) GetPrebuild(ctx context.Context, in *GetPrebuildRequest, opts ...grpc.CallOption) (*GetPrebuildResponse, error) {
	out := new(GetPrebuildResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.PrebuildService/GetPrebuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prebuildServiceClient) ListPrebuilds(ctx context.Context, in *ListPrebuildsRequest, opts ...grpc.CallOption) (*ListPrebuildsResponse, error) {
	out := new(ListPrebuildsResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.PrebuildService/ListPrebuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prebuildServiceClient) WatchPrebuild(ctx context.Context, in *WatchPrebuildRequest, opts ...grpc.CallOption) (PrebuildService_WatchPrebuildClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrebuildService_ServiceDesc.Streams[0], "/devpod.v1.PrebuildService/WatchPrebuild", opts...)
	if err != nil {
		return nil, err
	}
	x := &prebuildServiceWatchPrebuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrebuildService_WatchPrebuildClient interface {
	Recv() (*WatchPrebuildResponse, error)
	grpc.ClientStream
}

type prebuildServiceWatchPrebuildClient struct {
	grpc.ClientStream
}

func (x *prebuildServiceWatchPrebuildClient) Recv() (*WatchPrebuildResponse, error) {
	m := new(WatchPrebuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *prebuildServiceClient) ListOrganizationPrebuilds(ctx context.Context, in *ListOrganizationPrebuildsRequest, opts ...grpc.CallOption) (*ListOrganizationPrebuildsResponse, error) {
	out := new(ListOrganizationPrebuildsResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.PrebuildService/ListOrganizationPrebuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrebuildServiceServer is the server API for PrebuildService service.
// All implementations must embed UnimplementedPrebuildServiceServer
// for forward compatibility
type PrebuildServiceServer interface {
	StartPrebuild(context.Context, *StartPrebuildRequest) (*StartPrebuildResponse, error)
	CancelPrebuild(context.Context, *CancelPrebuildRequest) (*CancelPrebuildResponse, error)
	GetPrebuild(context.Context, *GetPrebuildRequest) (*GetPrebuildResponse, error)
	ListPrebuilds(context.Context, *ListPrebuildsRequest) (*ListPrebuildsResponse, error)
	WatchPrebuild(*WatchPrebuildRequest, PrebuildService_WatchPrebuildServer) error
	// ListOrganizationPrebuilds lists all prebuilds of an organization
	ListOrganizationPrebuilds(context.Context, *ListOrganizationPrebuildsRequest) (*ListOrganizationPrebuildsResponse, error)
	mustEmbedUnimplementedPrebuildServiceServer()
}

// UnimplementedPrebuildServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrebuildServiceServer struct {
}

func (UnimplementedPrebuildServiceServer) StartPrebuild(context.Context, *StartPrebuildRequest) (*StartPrebuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPrebuild not implemented")
}
func (UnimplementedPrebuildServiceServer) CancelPrebuild(context.Context, *CancelPrebuildRequest) (*CancelPrebuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPrebuild not implemented")
}
func (UnimplementedPrebuildServiceServer) GetPrebuild(context.Context, *GetPrebuildRequest) (*GetPrebuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrebuild not implemented")
}
func (UnimplementedPrebuildServiceServer) ListPrebuilds(context.Context, *ListPrebuildsRequest) (*ListPrebuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrebuilds not implemented")
}
func (UnimplementedPrebuildServiceServer) WatchPrebuild(*WatchPrebuildRequest, PrebuildService_WatchPrebuildServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPrebuild not implemented")
}
func (UnimplementedPrebuildServiceServer) ListOrganizationPrebuilds(context.Context, *ListOrganizationPrebuildsRequest) (*ListOrganizationPrebuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationPrebuilds not implemented")
}
func (UnimplementedPrebuildServiceServer) mustEmbedUnimplementedPrebuildServiceServer() {}

// UnsafePrebuildServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrebuildServiceServer will
// result in compilation errors.
type UnsafePrebuildServiceServer interface {
	mustEmbedUnimplementedPrebuildServiceServer()
}

func RegisterPrebuildServiceServer(s grpc.ServiceRegistrar, srv PrebuildServiceServer) {
	s.RegisterService(&PrebuildService_ServiceDesc, srv)
}

func _PrebuildService_StartPrebuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPrebuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildServiceServer).StartPrebuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.PrebuildService/StartPrebuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildServiceServer).StartPrebuild(ctx, req.(*StartPrebuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrebuildService_CancelPrebuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPrebuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildServiceServer).CancelPrebuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.PrebuildService/CancelPrebuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildServiceServer).CancelPrebuild(ctx, req.(*CancelPrebuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrebuildService_GetPrebuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrebuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildServiceServer).GetPrebuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.PrebuildService/GetPrebuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildServiceServer).GetPrebuild(ctx, req.(*GetPrebuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrebuildService_ListPrebuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPrebuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildServiceServer).ListPrebuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.PrebuildService/ListPrebuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildServiceServer).ListPrebuilds(ctx, req.(*ListPrebuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrebuildService_WatchPrebuild_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPrebuildRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrebuildServiceServer).WatchPrebuild(m, &prebuildServiceWatchPrebuildServer{stream})
}

type PrebuildService_WatchPrebuildServer interface {
	Send(*WatchPrebuildResponse) error
	grpc.ServerStream
}

type prebuildServiceWatchPrebuildServer struct {
	grpc.ServerStream
}

func (x *prebuildServiceWatchPrebuildServer) Send(m *WatchPrebuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PrebuildService_ListOrganizationPrebuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationPrebuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildServiceServer).ListOrganizationPrebuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.PrebuildService/ListOrganizationPrebuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildServiceServer).ListOrganizationPrebuilds(ctx, req.(*ListOrganizationPrebuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrebuildService_ServiceDesc is the grpc.ServiceDesc for PrebuildService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrebuildService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devpod.v1.PrebuildService",
	HandlerType: (*PrebuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPrebuild",
			Handler:    _PrebuildService_StartPrebuild_Handler,
		},
		{
			MethodName: "CancelPrebuild",
			Handler:    _PrebuildService_CancelPrebuild_Handler,
		},
		{
			MethodName: "GetPrebuild",
			Handler:    _PrebuildService_GetPrebuild_Handler,
		},
		{
			MethodName: "ListPrebuilds",
			Handler:    _PrebuildService_ListPrebuilds_Handler,
		},
		{
			MethodName: "ListOrganizationPrebuilds",
			Handler:    _PrebuildService_ListOrganizationPrebuilds_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPrebuild",
			Handler:       _PrebuildService_WatchPrebuild_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "devpod/v1/prebuild.proto",
}
