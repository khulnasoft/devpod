// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devpod/experimental/v1/tokens.proto
//
package io.devpod.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.ResponseMessage

public interface TokensServiceClientInterface {
  /**
   *  CreatePersonalAccessTokenRequest creates a new token.
   */
  public suspend fun createPersonalAccessToken(request: Tokens.CreatePersonalAccessTokenRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.CreatePersonalAccessTokenResponse>

  /**
   *  ListPersonalAccessTokens returns token by ID.
   */
  public suspend fun getPersonalAccessToken(request: Tokens.GetPersonalAccessTokenRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.GetPersonalAccessTokenResponse>

  /**
   *  ListPersonalAccessTokens returns a list of tokens.
   */
  public suspend fun listPersonalAccessTokens(request: Tokens.ListPersonalAccessTokensRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.ListPersonalAccessTokensResponse>

  /**
   *  RegeneratePersonalAccessToken generates a new token and replaces the previous one.
   */
  public suspend
      fun regeneratePersonalAccessToken(request: Tokens.RegeneratePersonalAccessTokenRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.RegeneratePersonalAccessTokenResponse>

  /**
   *  UpdatePersonalAccessToken updates writable properties of a PersonalAccessToken.
   */
  public suspend fun updatePersonalAccessToken(request: Tokens.UpdatePersonalAccessTokenRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.UpdatePersonalAccessTokenResponse>

  /**
   *  DeletePersonalAccessToken removes token by ID.
   */
  public suspend fun deletePersonalAccessToken(request: Tokens.DeletePersonalAccessTokenRequest,
      headers: Headers = emptyMap()): ResponseMessage<Tokens.DeletePersonalAccessTokenResponse>
}
