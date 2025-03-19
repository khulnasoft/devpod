// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devpod/components/public-api/go/v1"
)

var _ EnvironmentVariableServiceHandler = (*ProxyEnvironmentVariableServiceHandler)(nil)

type ProxyEnvironmentVariableServiceHandler struct {
	Client v1.EnvironmentVariableServiceClient
	UnimplementedEnvironmentVariableServiceHandler
}

func (s *ProxyEnvironmentVariableServiceHandler) ListUserEnvironmentVariables(ctx context.Context, req *connect_go.Request[v1.ListUserEnvironmentVariablesRequest]) (*connect_go.Response[v1.ListUserEnvironmentVariablesResponse], error) {
	resp, err := s.Client.ListUserEnvironmentVariables(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) UpdateUserEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.UpdateUserEnvironmentVariableRequest]) (*connect_go.Response[v1.UpdateUserEnvironmentVariableResponse], error) {
	resp, err := s.Client.UpdateUserEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) CreateUserEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.CreateUserEnvironmentVariableRequest]) (*connect_go.Response[v1.CreateUserEnvironmentVariableResponse], error) {
	resp, err := s.Client.CreateUserEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) DeleteUserEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.DeleteUserEnvironmentVariableRequest]) (*connect_go.Response[v1.DeleteUserEnvironmentVariableResponse], error) {
	resp, err := s.Client.DeleteUserEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) ListConfigurationEnvironmentVariables(ctx context.Context, req *connect_go.Request[v1.ListConfigurationEnvironmentVariablesRequest]) (*connect_go.Response[v1.ListConfigurationEnvironmentVariablesResponse], error) {
	resp, err := s.Client.ListConfigurationEnvironmentVariables(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) UpdateConfigurationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.UpdateConfigurationEnvironmentVariableRequest]) (*connect_go.Response[v1.UpdateConfigurationEnvironmentVariableResponse], error) {
	resp, err := s.Client.UpdateConfigurationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) CreateConfigurationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.CreateConfigurationEnvironmentVariableRequest]) (*connect_go.Response[v1.CreateConfigurationEnvironmentVariableResponse], error) {
	resp, err := s.Client.CreateConfigurationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) DeleteConfigurationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.DeleteConfigurationEnvironmentVariableRequest]) (*connect_go.Response[v1.DeleteConfigurationEnvironmentVariableResponse], error) {
	resp, err := s.Client.DeleteConfigurationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) ListOrganizationEnvironmentVariables(ctx context.Context, req *connect_go.Request[v1.ListOrganizationEnvironmentVariablesRequest]) (*connect_go.Response[v1.ListOrganizationEnvironmentVariablesResponse], error) {
	resp, err := s.Client.ListOrganizationEnvironmentVariables(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) UpdateOrganizationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.UpdateOrganizationEnvironmentVariableRequest]) (*connect_go.Response[v1.UpdateOrganizationEnvironmentVariableResponse], error) {
	resp, err := s.Client.UpdateOrganizationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) CreateOrganizationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.CreateOrganizationEnvironmentVariableRequest]) (*connect_go.Response[v1.CreateOrganizationEnvironmentVariableResponse], error) {
	resp, err := s.Client.CreateOrganizationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) DeleteOrganizationEnvironmentVariable(ctx context.Context, req *connect_go.Request[v1.DeleteOrganizationEnvironmentVariableRequest]) (*connect_go.Response[v1.DeleteOrganizationEnvironmentVariableResponse], error) {
	resp, err := s.Client.DeleteOrganizationEnvironmentVariable(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyEnvironmentVariableServiceHandler) ResolveWorkspaceEnvironmentVariables(ctx context.Context, req *connect_go.Request[v1.ResolveWorkspaceEnvironmentVariablesRequest]) (*connect_go.Response[v1.ResolveWorkspaceEnvironmentVariablesResponse], error) {
	resp, err := s.Client.ResolveWorkspaceEnvironmentVariables(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
