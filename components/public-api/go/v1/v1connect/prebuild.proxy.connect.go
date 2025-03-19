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

var _ PrebuildServiceHandler = (*ProxyPrebuildServiceHandler)(nil)

type ProxyPrebuildServiceHandler struct {
	Client v1.PrebuildServiceClient
	UnimplementedPrebuildServiceHandler
}

func (s *ProxyPrebuildServiceHandler) StartPrebuild(ctx context.Context, req *connect_go.Request[v1.StartPrebuildRequest]) (*connect_go.Response[v1.StartPrebuildResponse], error) {
	resp, err := s.Client.StartPrebuild(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyPrebuildServiceHandler) CancelPrebuild(ctx context.Context, req *connect_go.Request[v1.CancelPrebuildRequest]) (*connect_go.Response[v1.CancelPrebuildResponse], error) {
	resp, err := s.Client.CancelPrebuild(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyPrebuildServiceHandler) GetPrebuild(ctx context.Context, req *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error) {
	resp, err := s.Client.GetPrebuild(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyPrebuildServiceHandler) ListPrebuilds(ctx context.Context, req *connect_go.Request[v1.ListPrebuildsRequest]) (*connect_go.Response[v1.ListPrebuildsResponse], error) {
	resp, err := s.Client.ListPrebuilds(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyPrebuildServiceHandler) ListOrganizationPrebuilds(ctx context.Context, req *connect_go.Request[v1.ListOrganizationPrebuildsRequest]) (*connect_go.Response[v1.ListOrganizationPrebuildsResponse], error) {
	resp, err := s.Client.ListOrganizationPrebuilds(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
