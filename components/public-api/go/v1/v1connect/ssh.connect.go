// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: devpod/v1/ssh.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devpod/components/public-api/go/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SSHServiceName is the fully-qualified name of the SSHService service.
	SSHServiceName = "devpod.v1.SSHService"
)

// SSHServiceClient is a client for the devpod.v1.SSHService service.
type SSHServiceClient interface {
	// ListSSHPublicKeys returns all the ssh public keys for the
	// authenticated user.
	ListSSHPublicKeys(context.Context, *connect_go.Request[v1.ListSSHPublicKeysRequest]) (*connect_go.Response[v1.ListSSHPublicKeysResponse], error)
	// CreateSSHPublicKeys creates an ssh public key for the
	// authenticated user.
	CreateSSHPublicKey(context.Context, *connect_go.Request[v1.CreateSSHPublicKeyRequest]) (*connect_go.Response[v1.CreateSSHPublicKeyResponse], error)
	// DeleteSSHPublicKeys deletes an ssh public key for the
	// authenticated user.
	DeleteSSHPublicKey(context.Context, *connect_go.Request[v1.DeleteSSHPublicKeyRequest]) (*connect_go.Response[v1.DeleteSSHPublicKeyResponse], error)
}

// NewSSHServiceClient constructs a client for the devpod.v1.SSHService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSSHServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SSHServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sSHServiceClient{
		listSSHPublicKeys: connect_go.NewClient[v1.ListSSHPublicKeysRequest, v1.ListSSHPublicKeysResponse](
			httpClient,
			baseURL+"/devpod.v1.SSHService/ListSSHPublicKeys",
			opts...,
		),
		createSSHPublicKey: connect_go.NewClient[v1.CreateSSHPublicKeyRequest, v1.CreateSSHPublicKeyResponse](
			httpClient,
			baseURL+"/devpod.v1.SSHService/CreateSSHPublicKey",
			opts...,
		),
		deleteSSHPublicKey: connect_go.NewClient[v1.DeleteSSHPublicKeyRequest, v1.DeleteSSHPublicKeyResponse](
			httpClient,
			baseURL+"/devpod.v1.SSHService/DeleteSSHPublicKey",
			opts...,
		),
	}
}

// sSHServiceClient implements SSHServiceClient.
type sSHServiceClient struct {
	listSSHPublicKeys  *connect_go.Client[v1.ListSSHPublicKeysRequest, v1.ListSSHPublicKeysResponse]
	createSSHPublicKey *connect_go.Client[v1.CreateSSHPublicKeyRequest, v1.CreateSSHPublicKeyResponse]
	deleteSSHPublicKey *connect_go.Client[v1.DeleteSSHPublicKeyRequest, v1.DeleteSSHPublicKeyResponse]
}

// ListSSHPublicKeys calls devpod.v1.SSHService.ListSSHPublicKeys.
func (c *sSHServiceClient) ListSSHPublicKeys(ctx context.Context, req *connect_go.Request[v1.ListSSHPublicKeysRequest]) (*connect_go.Response[v1.ListSSHPublicKeysResponse], error) {
	return c.listSSHPublicKeys.CallUnary(ctx, req)
}

// CreateSSHPublicKey calls devpod.v1.SSHService.CreateSSHPublicKey.
func (c *sSHServiceClient) CreateSSHPublicKey(ctx context.Context, req *connect_go.Request[v1.CreateSSHPublicKeyRequest]) (*connect_go.Response[v1.CreateSSHPublicKeyResponse], error) {
	return c.createSSHPublicKey.CallUnary(ctx, req)
}

// DeleteSSHPublicKey calls devpod.v1.SSHService.DeleteSSHPublicKey.
func (c *sSHServiceClient) DeleteSSHPublicKey(ctx context.Context, req *connect_go.Request[v1.DeleteSSHPublicKeyRequest]) (*connect_go.Response[v1.DeleteSSHPublicKeyResponse], error) {
	return c.deleteSSHPublicKey.CallUnary(ctx, req)
}

// SSHServiceHandler is an implementation of the devpod.v1.SSHService service.
type SSHServiceHandler interface {
	// ListSSHPublicKeys returns all the ssh public keys for the
	// authenticated user.
	ListSSHPublicKeys(context.Context, *connect_go.Request[v1.ListSSHPublicKeysRequest]) (*connect_go.Response[v1.ListSSHPublicKeysResponse], error)
	// CreateSSHPublicKeys creates an ssh public key for the
	// authenticated user.
	CreateSSHPublicKey(context.Context, *connect_go.Request[v1.CreateSSHPublicKeyRequest]) (*connect_go.Response[v1.CreateSSHPublicKeyResponse], error)
	// DeleteSSHPublicKeys deletes an ssh public key for the
	// authenticated user.
	DeleteSSHPublicKey(context.Context, *connect_go.Request[v1.DeleteSSHPublicKeyRequest]) (*connect_go.Response[v1.DeleteSSHPublicKeyResponse], error)
}

// NewSSHServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSSHServiceHandler(svc SSHServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/devpod.v1.SSHService/ListSSHPublicKeys", connect_go.NewUnaryHandler(
		"/devpod.v1.SSHService/ListSSHPublicKeys",
		svc.ListSSHPublicKeys,
		opts...,
	))
	mux.Handle("/devpod.v1.SSHService/CreateSSHPublicKey", connect_go.NewUnaryHandler(
		"/devpod.v1.SSHService/CreateSSHPublicKey",
		svc.CreateSSHPublicKey,
		opts...,
	))
	mux.Handle("/devpod.v1.SSHService/DeleteSSHPublicKey", connect_go.NewUnaryHandler(
		"/devpod.v1.SSHService/DeleteSSHPublicKey",
		svc.DeleteSSHPublicKey,
		opts...,
	))
	return "/devpod.v1.SSHService/", mux
}

// UnimplementedSSHServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSSHServiceHandler struct{}

func (UnimplementedSSHServiceHandler) ListSSHPublicKeys(context.Context, *connect_go.Request[v1.ListSSHPublicKeysRequest]) (*connect_go.Response[v1.ListSSHPublicKeysResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devpod.v1.SSHService.ListSSHPublicKeys is not implemented"))
}

func (UnimplementedSSHServiceHandler) CreateSSHPublicKey(context.Context, *connect_go.Request[v1.CreateSSHPublicKeyRequest]) (*connect_go.Response[v1.CreateSSHPublicKeyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devpod.v1.SSHService.CreateSSHPublicKey is not implemented"))
}

func (UnimplementedSSHServiceHandler) DeleteSSHPublicKey(context.Context, *connect_go.Request[v1.DeleteSSHPublicKeyRequest]) (*connect_go.Response[v1.DeleteSSHPublicKeyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devpod.v1.SSHService.DeleteSSHPublicKey is not implemented"))
}
