/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file devpod/experimental/v1/identityprovider.proto (package devpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message devpod.experimental.v1.GetIDTokenRequest
 */
export class GetIDTokenRequest extends Message<GetIDTokenRequest> {
  /**
   * @generated from field: string workspace_id = 1;
   */
  workspaceId = "";

  /**
   * @generated from field: repeated string audience = 2;
   */
  audience: string[] = [];

  /**
   * @generated from field: string scope = 3;
   */
  scope = "";

  constructor(data?: PartialMessage<GetIDTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.experimental.v1.GetIDTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "audience", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "scope", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetIDTokenRequest {
    return new GetIDTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetIDTokenRequest {
    return new GetIDTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetIDTokenRequest {
    return new GetIDTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetIDTokenRequest | PlainMessage<GetIDTokenRequest> | undefined, b: GetIDTokenRequest | PlainMessage<GetIDTokenRequest> | undefined): boolean {
    return proto3.util.equals(GetIDTokenRequest, a, b);
  }
}

/**
 * @generated from message devpod.experimental.v1.GetIDTokenResponse
 */
export class GetIDTokenResponse extends Message<GetIDTokenResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token = "";

  constructor(data?: PartialMessage<GetIDTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.experimental.v1.GetIDTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetIDTokenResponse {
    return new GetIDTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetIDTokenResponse {
    return new GetIDTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetIDTokenResponse {
    return new GetIDTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetIDTokenResponse | PlainMessage<GetIDTokenResponse> | undefined, b: GetIDTokenResponse | PlainMessage<GetIDTokenResponse> | undefined): boolean {
    return proto3.util.equals(GetIDTokenResponse, a, b);
  }
}
