/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file devpod/experimental/v1/user.proto (package devpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { BlockUserRequest, BlockUserResponse, CreateSSHKeyRequest, CreateSSHKeyResponse, DeleteSSHKeyRequest, DeleteSSHKeyResponse, GetAuthenticatedUserRequest, GetAuthenticatedUserResponse, GetGitTokenRequest, GetGitTokenResponse, GetSSHKeyRequest, GetSSHKeyResponse, ListSSHKeysRequest, ListSSHKeysResponse } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service devpod.experimental.v1.UserService
 */
export const UserService = {
  typeName: "devpod.experimental.v1.UserService",
  methods: {
    /**
     * GetAuthenticatedUser gets the user info.
     *
     * @generated from rpc devpod.experimental.v1.UserService.GetAuthenticatedUser
     */
    getAuthenticatedUser: {
      name: "GetAuthenticatedUser",
      I: GetAuthenticatedUserRequest,
      O: GetAuthenticatedUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListSSHKeys lists the public SSH keys.
     *
     * @generated from rpc devpod.experimental.v1.UserService.ListSSHKeys
     */
    listSSHKeys: {
      name: "ListSSHKeys",
      I: ListSSHKeysRequest,
      O: ListSSHKeysResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateSSHKey adds a public SSH key.
     *
     * @generated from rpc devpod.experimental.v1.UserService.CreateSSHKey
     */
    createSSHKey: {
      name: "CreateSSHKey",
      I: CreateSSHKeyRequest,
      O: CreateSSHKeyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetSSHKey retrieves an ssh key by ID.
     *
     * @generated from rpc devpod.experimental.v1.UserService.GetSSHKey
     */
    getSSHKey: {
      name: "GetSSHKey",
      I: GetSSHKeyRequest,
      O: GetSSHKeyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteSSHKey removes a public SSH key.
     *
     * @generated from rpc devpod.experimental.v1.UserService.DeleteSSHKey
     */
    deleteSSHKey: {
      name: "DeleteSSHKey",
      I: DeleteSSHKeyRequest,
      O: DeleteSSHKeyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc devpod.experimental.v1.UserService.GetGitToken
     */
    getGitToken: {
      name: "GetGitToken",
      I: GetGitTokenRequest,
      O: GetGitTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc devpod.experimental.v1.UserService.BlockUser
     */
    blockUser: {
      name: "BlockUser",
      I: BlockUserRequest,
      O: BlockUserResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
