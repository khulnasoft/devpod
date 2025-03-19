// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devpod/experimental/v1/user.proto
//
package io.devpod.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class UserServiceClient(
  private val client: ProtocolClientInterface,
) : UserServiceClientInterface {
  /**
   *  GetAuthenticatedUser gets the user info.
   */
  override suspend fun getAuthenticatedUser(request: UserOuterClass.GetAuthenticatedUserRequest,
      headers: Headers): ResponseMessage<UserOuterClass.GetAuthenticatedUserResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/GetAuthenticatedUser",
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetAuthenticatedUserRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetAuthenticatedUserResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  ListSSHKeys lists the public SSH keys.
   */
  override suspend fun listSSHKeys(request: UserOuterClass.ListSSHKeysRequest, headers: Headers):
      ResponseMessage<UserOuterClass.ListSSHKeysResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/ListSSHKeys",
      io.devpod.publicapi.experimental.v1.UserOuterClass.ListSSHKeysRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.ListSSHKeysResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  CreateSSHKey adds a public SSH key.
   */
  override suspend fun createSSHKey(request: UserOuterClass.CreateSSHKeyRequest, headers: Headers):
      ResponseMessage<UserOuterClass.CreateSSHKeyResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/CreateSSHKey",
      io.devpod.publicapi.experimental.v1.UserOuterClass.CreateSSHKeyRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.CreateSSHKeyResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  GetSSHKey retrieves an ssh key by ID.
   */
  override suspend fun getSSHKey(request: UserOuterClass.GetSSHKeyRequest, headers: Headers):
      ResponseMessage<UserOuterClass.GetSSHKeyResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/GetSSHKey",
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetSSHKeyRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetSSHKeyResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  DeleteSSHKey removes a public SSH key.
   */
  override suspend fun deleteSSHKey(request: UserOuterClass.DeleteSSHKeyRequest, headers: Headers):
      ResponseMessage<UserOuterClass.DeleteSSHKeyResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/DeleteSSHKey",
      io.devpod.publicapi.experimental.v1.UserOuterClass.DeleteSSHKeyRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.DeleteSSHKeyResponse::class,
      StreamType.UNARY,
    ),
  )


  override suspend fun getGitToken(request: UserOuterClass.GetGitTokenRequest, headers: Headers):
      ResponseMessage<UserOuterClass.GetGitTokenResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/GetGitToken",
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetGitTokenRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.GetGitTokenResponse::class,
      StreamType.UNARY,
    ),
  )


  override suspend fun blockUser(request: UserOuterClass.BlockUserRequest, headers: Headers):
      ResponseMessage<UserOuterClass.BlockUserResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devpod.experimental.v1.UserService/BlockUser",
      io.devpod.publicapi.experimental.v1.UserOuterClass.BlockUserRequest::class,
      io.devpod.publicapi.experimental.v1.UserOuterClass.BlockUserResponse::class,
      StreamType.UNARY,
    ),
  )

}
