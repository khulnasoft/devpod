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

var _ SCMServiceHandler = (*ProxySCMServiceHandler)(nil)

type ProxySCMServiceHandler struct {
	Client v1.SCMServiceClient
	UnimplementedSCMServiceHandler
}

func (s *ProxySCMServiceHandler) SearchSCMTokens(ctx context.Context, req *connect_go.Request[v1.SearchSCMTokensRequest]) (*connect_go.Response[v1.SearchSCMTokensResponse], error) {
	resp, err := s.Client.SearchSCMTokens(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxySCMServiceHandler) GuessTokenScopes(ctx context.Context, req *connect_go.Request[v1.GuessTokenScopesRequest]) (*connect_go.Response[v1.GuessTokenScopesResponse], error) {
	resp, err := s.Client.GuessTokenScopes(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxySCMServiceHandler) SearchRepositories(ctx context.Context, req *connect_go.Request[v1.SearchRepositoriesRequest]) (*connect_go.Response[v1.SearchRepositoriesResponse], error) {
	resp, err := s.Client.SearchRepositories(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxySCMServiceHandler) ListSuggestedRepositories(ctx context.Context, req *connect_go.Request[v1.ListSuggestedRepositoriesRequest]) (*connect_go.Response[v1.ListSuggestedRepositoriesResponse], error) {
	resp, err := s.Client.ListSuggestedRepositories(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
