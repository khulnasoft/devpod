// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: devpod/v1/configuration.proto

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

// ConfigurationServiceClient is the client API for ConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationServiceClient interface {
	// Creates a new configuration.
	CreateConfiguration(ctx context.Context, in *CreateConfigurationRequest, opts ...grpc.CallOption) (*CreateConfigurationResponse, error)
	// Retrieves a configuration.
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	// Lists configurations.
	ListConfigurations(ctx context.Context, in *ListConfigurationsRequest, opts ...grpc.CallOption) (*ListConfigurationsResponse, error)
	// Updates a configuration.
	UpdateConfiguration(ctx context.Context, in *UpdateConfigurationRequest, opts ...grpc.CallOption) (*UpdateConfigurationResponse, error)
	// Deletes a configuration.
	DeleteConfiguration(ctx context.Context, in *DeleteConfigurationRequest, opts ...grpc.CallOption) (*DeleteConfigurationResponse, error)
}

type configurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationServiceClient(cc grpc.ClientConnInterface) ConfigurationServiceClient {
	return &configurationServiceClient{cc}
}

func (c *configurationServiceClient) CreateConfiguration(ctx context.Context, in *CreateConfigurationRequest, opts ...grpc.CallOption) (*CreateConfigurationResponse, error) {
	out := new(CreateConfigurationResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.ConfigurationService/CreateConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.ConfigurationService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) ListConfigurations(ctx context.Context, in *ListConfigurationsRequest, opts ...grpc.CallOption) (*ListConfigurationsResponse, error) {
	out := new(ListConfigurationsResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.ConfigurationService/ListConfigurations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) UpdateConfiguration(ctx context.Context, in *UpdateConfigurationRequest, opts ...grpc.CallOption) (*UpdateConfigurationResponse, error) {
	out := new(UpdateConfigurationResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.ConfigurationService/UpdateConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) DeleteConfiguration(ctx context.Context, in *DeleteConfigurationRequest, opts ...grpc.CallOption) (*DeleteConfigurationResponse, error) {
	out := new(DeleteConfigurationResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.ConfigurationService/DeleteConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServiceServer is the server API for ConfigurationService service.
// All implementations must embed UnimplementedConfigurationServiceServer
// for forward compatibility
type ConfigurationServiceServer interface {
	// Creates a new configuration.
	CreateConfiguration(context.Context, *CreateConfigurationRequest) (*CreateConfigurationResponse, error)
	// Retrieves a configuration.
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	// Lists configurations.
	ListConfigurations(context.Context, *ListConfigurationsRequest) (*ListConfigurationsResponse, error)
	// Updates a configuration.
	UpdateConfiguration(context.Context, *UpdateConfigurationRequest) (*UpdateConfigurationResponse, error)
	// Deletes a configuration.
	DeleteConfiguration(context.Context, *DeleteConfigurationRequest) (*DeleteConfigurationResponse, error)
	mustEmbedUnimplementedConfigurationServiceServer()
}

// UnimplementedConfigurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigurationServiceServer struct {
}

func (UnimplementedConfigurationServiceServer) CreateConfiguration(context.Context, *CreateConfigurationRequest) (*CreateConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfiguration not implemented")
}
func (UnimplementedConfigurationServiceServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedConfigurationServiceServer) ListConfigurations(context.Context, *ListConfigurationsRequest) (*ListConfigurationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigurations not implemented")
}
func (UnimplementedConfigurationServiceServer) UpdateConfiguration(context.Context, *UpdateConfigurationRequest) (*UpdateConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfiguration not implemented")
}
func (UnimplementedConfigurationServiceServer) DeleteConfiguration(context.Context, *DeleteConfigurationRequest) (*DeleteConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfiguration not implemented")
}
func (UnimplementedConfigurationServiceServer) mustEmbedUnimplementedConfigurationServiceServer() {}

// UnsafeConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationServiceServer will
// result in compilation errors.
type UnsafeConfigurationServiceServer interface {
	mustEmbedUnimplementedConfigurationServiceServer()
}

func RegisterConfigurationServiceServer(s grpc.ServiceRegistrar, srv ConfigurationServiceServer) {
	s.RegisterService(&ConfigurationService_ServiceDesc, srv)
}

func _ConfigurationService_CreateConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).CreateConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.ConfigurationService/CreateConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).CreateConfiguration(ctx, req.(*CreateConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.ConfigurationService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_ListConfigurations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigurationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).ListConfigurations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.ConfigurationService/ListConfigurations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).ListConfigurations(ctx, req.(*ListConfigurationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_UpdateConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).UpdateConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.ConfigurationService/UpdateConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).UpdateConfiguration(ctx, req.(*UpdateConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_DeleteConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).DeleteConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.ConfigurationService/DeleteConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).DeleteConfiguration(ctx, req.(*DeleteConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigurationService_ServiceDesc is the grpc.ServiceDesc for ConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devpod.v1.ConfigurationService",
	HandlerType: (*ConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConfiguration",
			Handler:    _ConfigurationService_CreateConfiguration_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _ConfigurationService_GetConfiguration_Handler,
		},
		{
			MethodName: "ListConfigurations",
			Handler:    _ConfigurationService_ListConfigurations_Handler,
		},
		{
			MethodName: "UpdateConfiguration",
			Handler:    _ConfigurationService_UpdateConfiguration_Handler,
		},
		{
			MethodName: "DeleteConfiguration",
			Handler:    _ConfigurationService_DeleteConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devpod/v1/configuration.proto",
}
