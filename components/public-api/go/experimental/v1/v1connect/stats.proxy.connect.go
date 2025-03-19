// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devpod/components/public-api/go/experimental/v1"
)

var _ StatsServiceHandler = (*ProxyStatsServiceHandler)(nil)

type ProxyStatsServiceHandler struct {
	Client v1.StatsServiceClient
	UnimplementedStatsServiceHandler
}

func (s *ProxyStatsServiceHandler) GetUserStats(ctx context.Context, req *connect_go.Request[v1.GetUserStatsRequest]) (*connect_go.Response[v1.GetUserStatsResponse], error) {
	resp, err := s.Client.GetUserStats(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
