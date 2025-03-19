// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: devpod/v1/installation.proto

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

// InstallationServiceClient is the client API for InstallationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstallationServiceClient interface {
	// GetInstallationWorkspaceDefaultImage returns the default image for current
	// Gitpod Installation.
	GetInstallationWorkspaceDefaultImage(ctx context.Context, in *GetInstallationWorkspaceDefaultImageRequest, opts ...grpc.CallOption) (*GetInstallationWorkspaceDefaultImageResponse, error)
	// ListBlockedRepositories lists blocked repositories.
	ListBlockedRepositories(ctx context.Context, in *ListBlockedRepositoriesRequest, opts ...grpc.CallOption) (*ListBlockedRepositoriesResponse, error)
	// CreateBlockedRepository creates a new blocked repository.
	CreateBlockedRepository(ctx context.Context, in *CreateBlockedRepositoryRequest, opts ...grpc.CallOption) (*CreateBlockedRepositoryResponse, error)
	// DeleteBlockedRepository deletes a blocked repository.
	DeleteBlockedRepository(ctx context.Context, in *DeleteBlockedRepositoryRequest, opts ...grpc.CallOption) (*DeleteBlockedRepositoryResponse, error)
	// ListBlockedEmailDomains lists blocked email domains.
	ListBlockedEmailDomains(ctx context.Context, in *ListBlockedEmailDomainsRequest, opts ...grpc.CallOption) (*ListBlockedEmailDomainsResponse, error)
	// CreateBlockedEmailDomain creates a new blocked email domain.
	CreateBlockedEmailDomain(ctx context.Context, in *CreateBlockedEmailDomainRequest, opts ...grpc.CallOption) (*CreateBlockedEmailDomainResponse, error)
	// GetOnboardingState returns the onboarding state of the installation.
	GetOnboardingState(ctx context.Context, in *GetOnboardingStateRequest, opts ...grpc.CallOption) (*GetOnboardingStateResponse, error)
	// GetInstallationConfiguration returns configuration of the installation.
	GetInstallationConfiguration(ctx context.Context, in *GetInstallationConfigurationRequest, opts ...grpc.CallOption) (*GetInstallationConfigurationResponse, error)
}

type installationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstallationServiceClient(cc grpc.ClientConnInterface) InstallationServiceClient {
	return &installationServiceClient{cc}
}

func (c *installationServiceClient) GetInstallationWorkspaceDefaultImage(ctx context.Context, in *GetInstallationWorkspaceDefaultImageRequest, opts ...grpc.CallOption) (*GetInstallationWorkspaceDefaultImageResponse, error) {
	out := new(GetInstallationWorkspaceDefaultImageResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/GetInstallationWorkspaceDefaultImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) ListBlockedRepositories(ctx context.Context, in *ListBlockedRepositoriesRequest, opts ...grpc.CallOption) (*ListBlockedRepositoriesResponse, error) {
	out := new(ListBlockedRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/ListBlockedRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) CreateBlockedRepository(ctx context.Context, in *CreateBlockedRepositoryRequest, opts ...grpc.CallOption) (*CreateBlockedRepositoryResponse, error) {
	out := new(CreateBlockedRepositoryResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/CreateBlockedRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) DeleteBlockedRepository(ctx context.Context, in *DeleteBlockedRepositoryRequest, opts ...grpc.CallOption) (*DeleteBlockedRepositoryResponse, error) {
	out := new(DeleteBlockedRepositoryResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/DeleteBlockedRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) ListBlockedEmailDomains(ctx context.Context, in *ListBlockedEmailDomainsRequest, opts ...grpc.CallOption) (*ListBlockedEmailDomainsResponse, error) {
	out := new(ListBlockedEmailDomainsResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/ListBlockedEmailDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) CreateBlockedEmailDomain(ctx context.Context, in *CreateBlockedEmailDomainRequest, opts ...grpc.CallOption) (*CreateBlockedEmailDomainResponse, error) {
	out := new(CreateBlockedEmailDomainResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/CreateBlockedEmailDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) GetOnboardingState(ctx context.Context, in *GetOnboardingStateRequest, opts ...grpc.CallOption) (*GetOnboardingStateResponse, error) {
	out := new(GetOnboardingStateResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/GetOnboardingState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installationServiceClient) GetInstallationConfiguration(ctx context.Context, in *GetInstallationConfigurationRequest, opts ...grpc.CallOption) (*GetInstallationConfigurationResponse, error) {
	out := new(GetInstallationConfigurationResponse)
	err := c.cc.Invoke(ctx, "/devpod.v1.InstallationService/GetInstallationConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstallationServiceServer is the server API for InstallationService service.
// All implementations must embed UnimplementedInstallationServiceServer
// for forward compatibility
type InstallationServiceServer interface {
	// GetInstallationWorkspaceDefaultImage returns the default image for current
	// Gitpod Installation.
	GetInstallationWorkspaceDefaultImage(context.Context, *GetInstallationWorkspaceDefaultImageRequest) (*GetInstallationWorkspaceDefaultImageResponse, error)
	// ListBlockedRepositories lists blocked repositories.
	ListBlockedRepositories(context.Context, *ListBlockedRepositoriesRequest) (*ListBlockedRepositoriesResponse, error)
	// CreateBlockedRepository creates a new blocked repository.
	CreateBlockedRepository(context.Context, *CreateBlockedRepositoryRequest) (*CreateBlockedRepositoryResponse, error)
	// DeleteBlockedRepository deletes a blocked repository.
	DeleteBlockedRepository(context.Context, *DeleteBlockedRepositoryRequest) (*DeleteBlockedRepositoryResponse, error)
	// ListBlockedEmailDomains lists blocked email domains.
	ListBlockedEmailDomains(context.Context, *ListBlockedEmailDomainsRequest) (*ListBlockedEmailDomainsResponse, error)
	// CreateBlockedEmailDomain creates a new blocked email domain.
	CreateBlockedEmailDomain(context.Context, *CreateBlockedEmailDomainRequest) (*CreateBlockedEmailDomainResponse, error)
	// GetOnboardingState returns the onboarding state of the installation.
	GetOnboardingState(context.Context, *GetOnboardingStateRequest) (*GetOnboardingStateResponse, error)
	// GetInstallationConfiguration returns configuration of the installation.
	GetInstallationConfiguration(context.Context, *GetInstallationConfigurationRequest) (*GetInstallationConfigurationResponse, error)
	mustEmbedUnimplementedInstallationServiceServer()
}

// UnimplementedInstallationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstallationServiceServer struct {
}

func (UnimplementedInstallationServiceServer) GetInstallationWorkspaceDefaultImage(context.Context, *GetInstallationWorkspaceDefaultImageRequest) (*GetInstallationWorkspaceDefaultImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstallationWorkspaceDefaultImage not implemented")
}
func (UnimplementedInstallationServiceServer) ListBlockedRepositories(context.Context, *ListBlockedRepositoriesRequest) (*ListBlockedRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlockedRepositories not implemented")
}
func (UnimplementedInstallationServiceServer) CreateBlockedRepository(context.Context, *CreateBlockedRepositoryRequest) (*CreateBlockedRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlockedRepository not implemented")
}
func (UnimplementedInstallationServiceServer) DeleteBlockedRepository(context.Context, *DeleteBlockedRepositoryRequest) (*DeleteBlockedRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlockedRepository not implemented")
}
func (UnimplementedInstallationServiceServer) ListBlockedEmailDomains(context.Context, *ListBlockedEmailDomainsRequest) (*ListBlockedEmailDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlockedEmailDomains not implemented")
}
func (UnimplementedInstallationServiceServer) CreateBlockedEmailDomain(context.Context, *CreateBlockedEmailDomainRequest) (*CreateBlockedEmailDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlockedEmailDomain not implemented")
}
func (UnimplementedInstallationServiceServer) GetOnboardingState(context.Context, *GetOnboardingStateRequest) (*GetOnboardingStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnboardingState not implemented")
}
func (UnimplementedInstallationServiceServer) GetInstallationConfiguration(context.Context, *GetInstallationConfigurationRequest) (*GetInstallationConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstallationConfiguration not implemented")
}
func (UnimplementedInstallationServiceServer) mustEmbedUnimplementedInstallationServiceServer() {}

// UnsafeInstallationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstallationServiceServer will
// result in compilation errors.
type UnsafeInstallationServiceServer interface {
	mustEmbedUnimplementedInstallationServiceServer()
}

func RegisterInstallationServiceServer(s grpc.ServiceRegistrar, srv InstallationServiceServer) {
	s.RegisterService(&InstallationService_ServiceDesc, srv)
}

func _InstallationService_GetInstallationWorkspaceDefaultImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstallationWorkspaceDefaultImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).GetInstallationWorkspaceDefaultImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/GetInstallationWorkspaceDefaultImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).GetInstallationWorkspaceDefaultImage(ctx, req.(*GetInstallationWorkspaceDefaultImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_ListBlockedRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlockedRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).ListBlockedRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/ListBlockedRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).ListBlockedRepositories(ctx, req.(*ListBlockedRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_CreateBlockedRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockedRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).CreateBlockedRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/CreateBlockedRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).CreateBlockedRepository(ctx, req.(*CreateBlockedRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_DeleteBlockedRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockedRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).DeleteBlockedRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/DeleteBlockedRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).DeleteBlockedRepository(ctx, req.(*DeleteBlockedRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_ListBlockedEmailDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlockedEmailDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).ListBlockedEmailDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/ListBlockedEmailDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).ListBlockedEmailDomains(ctx, req.(*ListBlockedEmailDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_CreateBlockedEmailDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockedEmailDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).CreateBlockedEmailDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/CreateBlockedEmailDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).CreateBlockedEmailDomain(ctx, req.(*CreateBlockedEmailDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_GetOnboardingState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnboardingStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).GetOnboardingState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/GetOnboardingState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).GetOnboardingState(ctx, req.(*GetOnboardingStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallationService_GetInstallationConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstallationConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationServiceServer).GetInstallationConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devpod.v1.InstallationService/GetInstallationConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationServiceServer).GetInstallationConfiguration(ctx, req.(*GetInstallationConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstallationService_ServiceDesc is the grpc.ServiceDesc for InstallationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstallationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devpod.v1.InstallationService",
	HandlerType: (*InstallationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInstallationWorkspaceDefaultImage",
			Handler:    _InstallationService_GetInstallationWorkspaceDefaultImage_Handler,
		},
		{
			MethodName: "ListBlockedRepositories",
			Handler:    _InstallationService_ListBlockedRepositories_Handler,
		},
		{
			MethodName: "CreateBlockedRepository",
			Handler:    _InstallationService_CreateBlockedRepository_Handler,
		},
		{
			MethodName: "DeleteBlockedRepository",
			Handler:    _InstallationService_DeleteBlockedRepository_Handler,
		},
		{
			MethodName: "ListBlockedEmailDomains",
			Handler:    _InstallationService_ListBlockedEmailDomains_Handler,
		},
		{
			MethodName: "CreateBlockedEmailDomain",
			Handler:    _InstallationService_CreateBlockedEmailDomain_Handler,
		},
		{
			MethodName: "GetOnboardingState",
			Handler:    _InstallationService_GetOnboardingState_Handler,
		},
		{
			MethodName: "GetInstallationConfiguration",
			Handler:    _InstallationService_GetInstallationConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devpod/v1/installation.proto",
}
