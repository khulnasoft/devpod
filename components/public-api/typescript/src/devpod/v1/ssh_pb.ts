/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file devpod/v1/ssh.proto (package devpod.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { PaginationRequest } from "./pagination_pb.js";

/**
 * @generated from message devpod.v1.SSHPublicKey
 */
export class SSHPublicKey extends Message<SSHPublicKey> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string key = 3;
   */
  key = "";

  /**
   * @generated from field: string fingerprint = 4;
   */
  fingerprint = "";

  /**
   * @generated from field: google.protobuf.Timestamp creation_time = 5;
   */
  creationTime?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp last_used_time = 6;
   */
  lastUsedTime?: Timestamp;

  constructor(data?: PartialMessage<SSHPublicKey>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.SSHPublicKey";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "fingerprint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "creation_time", kind: "message", T: Timestamp },
    { no: 6, name: "last_used_time", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SSHPublicKey {
    return new SSHPublicKey().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SSHPublicKey {
    return new SSHPublicKey().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SSHPublicKey {
    return new SSHPublicKey().fromJsonString(jsonString, options);
  }

  static equals(a: SSHPublicKey | PlainMessage<SSHPublicKey> | undefined, b: SSHPublicKey | PlainMessage<SSHPublicKey> | undefined): boolean {
    return proto3.util.equals(SSHPublicKey, a, b);
  }
}

/**
 * @generated from message devpod.v1.ListSSHPublicKeysRequest
 */
export class ListSSHPublicKeysRequest extends Message<ListSSHPublicKeysRequest> {
  /**
   * @generated from field: devpod.v1.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  constructor(data?: PartialMessage<ListSSHPublicKeysRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.ListSSHPublicKeysRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSSHPublicKeysRequest {
    return new ListSSHPublicKeysRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSSHPublicKeysRequest {
    return new ListSSHPublicKeysRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSSHPublicKeysRequest {
    return new ListSSHPublicKeysRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSSHPublicKeysRequest | PlainMessage<ListSSHPublicKeysRequest> | undefined, b: ListSSHPublicKeysRequest | PlainMessage<ListSSHPublicKeysRequest> | undefined): boolean {
    return proto3.util.equals(ListSSHPublicKeysRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.ListSSHPublicKeysResponse
 */
export class ListSSHPublicKeysResponse extends Message<ListSSHPublicKeysResponse> {
  /**
   * @generated from field: repeated devpod.v1.SSHPublicKey ssh_keys = 1;
   */
  sshKeys: SSHPublicKey[] = [];

  /**
   * @generated from field: devpod.v1.PaginationRequest pagination = 2;
   */
  pagination?: PaginationRequest;

  constructor(data?: PartialMessage<ListSSHPublicKeysResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.ListSSHPublicKeysResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ssh_keys", kind: "message", T: SSHPublicKey, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PaginationRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSSHPublicKeysResponse {
    return new ListSSHPublicKeysResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSSHPublicKeysResponse {
    return new ListSSHPublicKeysResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSSHPublicKeysResponse {
    return new ListSSHPublicKeysResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSSHPublicKeysResponse | PlainMessage<ListSSHPublicKeysResponse> | undefined, b: ListSSHPublicKeysResponse | PlainMessage<ListSSHPublicKeysResponse> | undefined): boolean {
    return proto3.util.equals(ListSSHPublicKeysResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.CreateSSHPublicKeyRequest
 */
export class CreateSSHPublicKeyRequest extends Message<CreateSSHPublicKeyRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string key = 2;
   */
  key = "";

  constructor(data?: PartialMessage<CreateSSHPublicKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.CreateSSHPublicKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSSHPublicKeyRequest {
    return new CreateSSHPublicKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSSHPublicKeyRequest {
    return new CreateSSHPublicKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSSHPublicKeyRequest {
    return new CreateSSHPublicKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSSHPublicKeyRequest | PlainMessage<CreateSSHPublicKeyRequest> | undefined, b: CreateSSHPublicKeyRequest | PlainMessage<CreateSSHPublicKeyRequest> | undefined): boolean {
    return proto3.util.equals(CreateSSHPublicKeyRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.CreateSSHPublicKeyResponse
 */
export class CreateSSHPublicKeyResponse extends Message<CreateSSHPublicKeyResponse> {
  /**
   * @generated from field: devpod.v1.SSHPublicKey ssh_key = 1;
   */
  sshKey?: SSHPublicKey;

  constructor(data?: PartialMessage<CreateSSHPublicKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.CreateSSHPublicKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ssh_key", kind: "message", T: SSHPublicKey },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSSHPublicKeyResponse {
    return new CreateSSHPublicKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSSHPublicKeyResponse {
    return new CreateSSHPublicKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSSHPublicKeyResponse {
    return new CreateSSHPublicKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSSHPublicKeyResponse | PlainMessage<CreateSSHPublicKeyResponse> | undefined, b: CreateSSHPublicKeyResponse | PlainMessage<CreateSSHPublicKeyResponse> | undefined): boolean {
    return proto3.util.equals(CreateSSHPublicKeyResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.DeleteSSHPublicKeyRequest
 */
export class DeleteSSHPublicKeyRequest extends Message<DeleteSSHPublicKeyRequest> {
  /**
   * @generated from field: string ssh_key_id = 1;
   */
  sshKeyId = "";

  constructor(data?: PartialMessage<DeleteSSHPublicKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.DeleteSSHPublicKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ssh_key_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSSHPublicKeyRequest {
    return new DeleteSSHPublicKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSSHPublicKeyRequest {
    return new DeleteSSHPublicKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSSHPublicKeyRequest {
    return new DeleteSSHPublicKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSSHPublicKeyRequest | PlainMessage<DeleteSSHPublicKeyRequest> | undefined, b: DeleteSSHPublicKeyRequest | PlainMessage<DeleteSSHPublicKeyRequest> | undefined): boolean {
    return proto3.util.equals(DeleteSSHPublicKeyRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.DeleteSSHPublicKeyResponse
 */
export class DeleteSSHPublicKeyResponse extends Message<DeleteSSHPublicKeyResponse> {
  constructor(data?: PartialMessage<DeleteSSHPublicKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.DeleteSSHPublicKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSSHPublicKeyResponse {
    return new DeleteSSHPublicKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSSHPublicKeyResponse {
    return new DeleteSSHPublicKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSSHPublicKeyResponse {
    return new DeleteSSHPublicKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSSHPublicKeyResponse | PlainMessage<DeleteSSHPublicKeyResponse> | undefined, b: DeleteSSHPublicKeyResponse | PlainMessage<DeleteSSHPublicKeyResponse> | undefined): boolean {
    return proto3.util.equals(DeleteSSHPublicKeyResponse, a, b);
  }
}
