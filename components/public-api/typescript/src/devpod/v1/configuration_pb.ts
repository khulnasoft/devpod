/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file devpod/v1/configuration.proto (package devpod.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { PaginationRequest, PaginationResponse } from "./pagination_pb.js";
import { Sort } from "./sorting_pb.js";

/**
 * @generated from enum devpod.v1.PrebuildTriggerStrategy
 */
export enum PrebuildTriggerStrategy {
  /**
   * Default value. Implicitly applies to webhoook-based activation
   *
   * @generated from enum value: PREBUILD_TRIGGER_STRATEGY_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * Default value for newly enabled prebuilds.
   *
   * @generated from enum value: PREBUILD_TRIGGER_STRATEGY_ACTIVITY_BASED = 1;
   */
  ACTIVITY_BASED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(PrebuildTriggerStrategy)
proto3.util.setEnumType(PrebuildTriggerStrategy, "devpod.v1.PrebuildTriggerStrategy", [
  { no: 0, name: "PREBUILD_TRIGGER_STRATEGY_UNSPECIFIED" },
  { no: 1, name: "PREBUILD_TRIGGER_STRATEGY_ACTIVITY_BASED" },
]);

/**
 * @generated from enum devpod.v1.BranchMatchingStrategy
 */
export enum BranchMatchingStrategy {
  /**
   * @generated from enum value: BRANCH_MATCHING_STRATEGY_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: BRANCH_MATCHING_STRATEGY_DEFAULT_BRANCH = 1;
   */
  DEFAULT_BRANCH = 1,

  /**
   * @generated from enum value: BRANCH_MATCHING_STRATEGY_ALL_BRANCHES = 2;
   */
  ALL_BRANCHES = 2,

  /**
   * @generated from enum value: BRANCH_MATCHING_STRATEGY_MATCHED_BRANCHES = 3;
   */
  MATCHED_BRANCHES = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(BranchMatchingStrategy)
proto3.util.setEnumType(BranchMatchingStrategy, "devpod.v1.BranchMatchingStrategy", [
  { no: 0, name: "BRANCH_MATCHING_STRATEGY_UNSPECIFIED" },
  { no: 1, name: "BRANCH_MATCHING_STRATEGY_DEFAULT_BRANCH" },
  { no: 2, name: "BRANCH_MATCHING_STRATEGY_ALL_BRANCHES" },
  { no: 3, name: "BRANCH_MATCHING_STRATEGY_MATCHED_BRANCHES" },
]);

/**
 * @generated from message devpod.v1.Configuration
 */
export class Configuration extends Message<Configuration> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string organization_id = 2;
   */
  organizationId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string clone_url = 4;
   */
  cloneUrl = "";

  /**
   * @generated from field: google.protobuf.Timestamp creation_time = 5;
   */
  creationTime?: Timestamp;

  /**
   * @generated from field: devpod.v1.PrebuildSettings prebuild_settings = 6;
   */
  prebuildSettings?: PrebuildSettings;

  /**
   * @generated from field: devpod.v1.WorkspaceSettings workspace_settings = 7;
   */
  workspaceSettings?: WorkspaceSettings;

  constructor(data?: PartialMessage<Configuration>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.Configuration";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "organization_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "clone_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "creation_time", kind: "message", T: Timestamp },
    { no: 6, name: "prebuild_settings", kind: "message", T: PrebuildSettings },
    { no: 7, name: "workspace_settings", kind: "message", T: WorkspaceSettings },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Configuration {
    return new Configuration().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Configuration {
    return new Configuration().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Configuration {
    return new Configuration().fromJsonString(jsonString, options);
  }

  static equals(a: Configuration | PlainMessage<Configuration> | undefined, b: Configuration | PlainMessage<Configuration> | undefined): boolean {
    return proto3.util.equals(Configuration, a, b);
  }
}

/**
 * @generated from message devpod.v1.PrebuildSettings
 */
export class PrebuildSettings extends Message<PrebuildSettings> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled = false;

  /**
   * @generated from field: string branch_matching_pattern = 2;
   */
  branchMatchingPattern = "";

  /**
   * @generated from field: devpod.v1.BranchMatchingStrategy branch_strategy = 3;
   */
  branchStrategy = BranchMatchingStrategy.UNSPECIFIED;

  /**
   * @generated from field: int32 prebuild_interval = 4;
   */
  prebuildInterval = 0;

  /**
   * @generated from field: string workspace_class = 5;
   */
  workspaceClass = "";

  /**
   * @generated from field: devpod.v1.PrebuildTriggerStrategy trigger_strategy = 6;
   */
  triggerStrategy = PrebuildTriggerStrategy.UNSPECIFIED;

  /**
   * @generated from field: devpod.v1.PrebuildCloneSettings clone_settings = 7;
   */
  cloneSettings?: PrebuildCloneSettings;

  constructor(data?: PartialMessage<PrebuildSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.PrebuildSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "branch_matching_pattern", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "branch_strategy", kind: "enum", T: proto3.getEnumType(BranchMatchingStrategy) },
    { no: 4, name: "prebuild_interval", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "workspace_class", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "trigger_strategy", kind: "enum", T: proto3.getEnumType(PrebuildTriggerStrategy) },
    { no: 7, name: "clone_settings", kind: "message", T: PrebuildCloneSettings },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PrebuildSettings {
    return new PrebuildSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PrebuildSettings {
    return new PrebuildSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PrebuildSettings {
    return new PrebuildSettings().fromJsonString(jsonString, options);
  }

  static equals(a: PrebuildSettings | PlainMessage<PrebuildSettings> | undefined, b: PrebuildSettings | PlainMessage<PrebuildSettings> | undefined): boolean {
    return proto3.util.equals(PrebuildSettings, a, b);
  }
}

/**
 * @generated from message devpod.v1.PrebuildCloneSettings
 */
export class PrebuildCloneSettings extends Message<PrebuildCloneSettings> {
  /**
   * full_clone determines if the entire repository should be cloned, instead of with `--depth=1`
   *
   * @generated from field: bool full_clone = 1;
   */
  fullClone = false;

  constructor(data?: PartialMessage<PrebuildCloneSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.PrebuildCloneSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "full_clone", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PrebuildCloneSettings {
    return new PrebuildCloneSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PrebuildCloneSettings {
    return new PrebuildCloneSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PrebuildCloneSettings {
    return new PrebuildCloneSettings().fromJsonString(jsonString, options);
  }

  static equals(a: PrebuildCloneSettings | PlainMessage<PrebuildCloneSettings> | undefined, b: PrebuildCloneSettings | PlainMessage<PrebuildCloneSettings> | undefined): boolean {
    return proto3.util.equals(PrebuildCloneSettings, a, b);
  }
}

/**
 * @generated from message devpod.v1.WorkspaceSettings
 */
export class WorkspaceSettings extends Message<WorkspaceSettings> {
  /**
   * @generated from field: string workspace_class = 1;
   */
  workspaceClass = "";

  /**
   * @generated from field: repeated string restricted_workspace_classes = 2;
   */
  restrictedWorkspaceClasses: string[] = [];

  /**
   * @generated from field: repeated string restricted_editor_names = 3;
   */
  restrictedEditorNames: string[] = [];

  /**
   * Enable automatic authentication for docker daemon with all credentials specified in GITPOD_IMAGE_AUTH
   *
   * @generated from field: bool enable_dockerd_authentication = 4;
   */
  enableDockerdAuthentication = false;

  constructor(data?: PartialMessage<WorkspaceSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.WorkspaceSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workspace_class", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "restricted_workspace_classes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "restricted_editor_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "enable_dockerd_authentication", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkspaceSettings {
    return new WorkspaceSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkspaceSettings {
    return new WorkspaceSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkspaceSettings {
    return new WorkspaceSettings().fromJsonString(jsonString, options);
  }

  static equals(a: WorkspaceSettings | PlainMessage<WorkspaceSettings> | undefined, b: WorkspaceSettings | PlainMessage<WorkspaceSettings> | undefined): boolean {
    return proto3.util.equals(WorkspaceSettings, a, b);
  }
}

/**
 * @generated from message devpod.v1.CreateConfigurationRequest
 */
export class CreateConfigurationRequest extends Message<CreateConfigurationRequest> {
  /**
   * @generated from field: string organization_id = 1;
   */
  organizationId = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string clone_url = 3;
   */
  cloneUrl = "";

  constructor(data?: PartialMessage<CreateConfigurationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.CreateConfigurationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "clone_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConfigurationRequest {
    return new CreateConfigurationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConfigurationRequest {
    return new CreateConfigurationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConfigurationRequest {
    return new CreateConfigurationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateConfigurationRequest | PlainMessage<CreateConfigurationRequest> | undefined, b: CreateConfigurationRequest | PlainMessage<CreateConfigurationRequest> | undefined): boolean {
    return proto3.util.equals(CreateConfigurationRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.CreateConfigurationResponse
 */
export class CreateConfigurationResponse extends Message<CreateConfigurationResponse> {
  /**
   * @generated from field: devpod.v1.Configuration configuration = 1;
   */
  configuration?: Configuration;

  constructor(data?: PartialMessage<CreateConfigurationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.CreateConfigurationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration", kind: "message", T: Configuration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConfigurationResponse {
    return new CreateConfigurationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConfigurationResponse {
    return new CreateConfigurationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConfigurationResponse {
    return new CreateConfigurationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateConfigurationResponse | PlainMessage<CreateConfigurationResponse> | undefined, b: CreateConfigurationResponse | PlainMessage<CreateConfigurationResponse> | undefined): boolean {
    return proto3.util.equals(CreateConfigurationResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.GetConfigurationRequest
 */
export class GetConfigurationRequest extends Message<GetConfigurationRequest> {
  /**
   * @generated from field: string configuration_id = 1;
   */
  configurationId = "";

  constructor(data?: PartialMessage<GetConfigurationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.GetConfigurationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConfigurationRequest {
    return new GetConfigurationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConfigurationRequest {
    return new GetConfigurationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConfigurationRequest {
    return new GetConfigurationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetConfigurationRequest | PlainMessage<GetConfigurationRequest> | undefined, b: GetConfigurationRequest | PlainMessage<GetConfigurationRequest> | undefined): boolean {
    return proto3.util.equals(GetConfigurationRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.GetConfigurationResponse
 */
export class GetConfigurationResponse extends Message<GetConfigurationResponse> {
  /**
   * @generated from field: devpod.v1.Configuration configuration = 1;
   */
  configuration?: Configuration;

  constructor(data?: PartialMessage<GetConfigurationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.GetConfigurationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration", kind: "message", T: Configuration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConfigurationResponse {
    return new GetConfigurationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConfigurationResponse {
    return new GetConfigurationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConfigurationResponse {
    return new GetConfigurationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetConfigurationResponse | PlainMessage<GetConfigurationResponse> | undefined, b: GetConfigurationResponse | PlainMessage<GetConfigurationResponse> | undefined): boolean {
    return proto3.util.equals(GetConfigurationResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.ListConfigurationsRequest
 */
export class ListConfigurationsRequest extends Message<ListConfigurationsRequest> {
  /**
   * @generated from field: string organization_id = 1;
   */
  organizationId = "";

  /**
   * @generated from field: string search_term = 2;
   */
  searchTerm = "";

  /**
   * @generated from field: devpod.v1.PaginationRequest pagination = 3;
   */
  pagination?: PaginationRequest;

  /**
   * Configurations can be sorted by "name" OR "creationTime"
   *
   * @generated from field: repeated devpod.v1.Sort sort = 4;
   */
  sort: Sort[] = [];

  /**
   * Will filter for configurations that have prebuilds enabled/disabled, or both if not set.
   *
   * @generated from field: optional bool prebuilds_enabled = 5;
   */
  prebuildsEnabled?: boolean;

  constructor(data?: PartialMessage<ListConfigurationsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.ListConfigurationsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "search_term", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "pagination", kind: "message", T: PaginationRequest },
    { no: 4, name: "sort", kind: "message", T: Sort, repeated: true },
    { no: 5, name: "prebuilds_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListConfigurationsRequest {
    return new ListConfigurationsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListConfigurationsRequest {
    return new ListConfigurationsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListConfigurationsRequest {
    return new ListConfigurationsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListConfigurationsRequest | PlainMessage<ListConfigurationsRequest> | undefined, b: ListConfigurationsRequest | PlainMessage<ListConfigurationsRequest> | undefined): boolean {
    return proto3.util.equals(ListConfigurationsRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.ListConfigurationsResponse
 */
export class ListConfigurationsResponse extends Message<ListConfigurationsResponse> {
  /**
   * @generated from field: repeated devpod.v1.Configuration configurations = 1;
   */
  configurations: Configuration[] = [];

  /**
   * @generated from field: devpod.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;

  constructor(data?: PartialMessage<ListConfigurationsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.ListConfigurationsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configurations", kind: "message", T: Configuration, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PaginationResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListConfigurationsResponse {
    return new ListConfigurationsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListConfigurationsResponse {
    return new ListConfigurationsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListConfigurationsResponse {
    return new ListConfigurationsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListConfigurationsResponse | PlainMessage<ListConfigurationsResponse> | undefined, b: ListConfigurationsResponse | PlainMessage<ListConfigurationsResponse> | undefined): boolean {
    return proto3.util.equals(ListConfigurationsResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.UpdateConfigurationRequest
 */
export class UpdateConfigurationRequest extends Message<UpdateConfigurationRequest> {
  /**
   * @generated from field: string configuration_id = 1;
   */
  configurationId = "";

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  /**
   * @generated from field: optional devpod.v1.UpdateConfigurationRequest.PrebuildSettings prebuild_settings = 3;
   */
  prebuildSettings?: UpdateConfigurationRequest_PrebuildSettings;

  /**
   * @generated from field: optional devpod.v1.UpdateConfigurationRequest.WorkspaceSettings workspace_settings = 4;
   */
  workspaceSettings?: UpdateConfigurationRequest_WorkspaceSettings;

  constructor(data?: PartialMessage<UpdateConfigurationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.UpdateConfigurationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "prebuild_settings", kind: "message", T: UpdateConfigurationRequest_PrebuildSettings, opt: true },
    { no: 4, name: "workspace_settings", kind: "message", T: UpdateConfigurationRequest_WorkspaceSettings, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConfigurationRequest {
    return new UpdateConfigurationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest {
    return new UpdateConfigurationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest {
    return new UpdateConfigurationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConfigurationRequest | PlainMessage<UpdateConfigurationRequest> | undefined, b: UpdateConfigurationRequest | PlainMessage<UpdateConfigurationRequest> | undefined): boolean {
    return proto3.util.equals(UpdateConfigurationRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.UpdateConfigurationRequest.PrebuildSettings
 */
export class UpdateConfigurationRequest_PrebuildSettings extends Message<UpdateConfigurationRequest_PrebuildSettings> {
  /**
   * @generated from field: optional bool enabled = 1;
   */
  enabled?: boolean;

  /**
   * @generated from field: optional string branch_matching_pattern = 2;
   */
  branchMatchingPattern?: string;

  /**
   * @generated from field: optional devpod.v1.BranchMatchingStrategy branch_strategy = 3;
   */
  branchStrategy?: BranchMatchingStrategy;

  /**
   * @generated from field: optional int32 prebuild_interval = 4;
   */
  prebuildInterval?: number;

  /**
   * @generated from field: optional string workspace_class = 5;
   */
  workspaceClass?: string;

  /**
   * @generated from field: optional devpod.v1.PrebuildTriggerStrategy trigger_strategy = 6;
   */
  triggerStrategy?: PrebuildTriggerStrategy;

  /**
   * @generated from field: optional devpod.v1.PrebuildCloneSettings clone_settings = 7;
   */
  cloneSettings?: PrebuildCloneSettings;

  constructor(data?: PartialMessage<UpdateConfigurationRequest_PrebuildSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.UpdateConfigurationRequest.PrebuildSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "branch_matching_pattern", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "branch_strategy", kind: "enum", T: proto3.getEnumType(BranchMatchingStrategy), opt: true },
    { no: 4, name: "prebuild_interval", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "workspace_class", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "trigger_strategy", kind: "enum", T: proto3.getEnumType(PrebuildTriggerStrategy), opt: true },
    { no: 7, name: "clone_settings", kind: "message", T: PrebuildCloneSettings, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConfigurationRequest_PrebuildSettings {
    return new UpdateConfigurationRequest_PrebuildSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest_PrebuildSettings {
    return new UpdateConfigurationRequest_PrebuildSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest_PrebuildSettings {
    return new UpdateConfigurationRequest_PrebuildSettings().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConfigurationRequest_PrebuildSettings | PlainMessage<UpdateConfigurationRequest_PrebuildSettings> | undefined, b: UpdateConfigurationRequest_PrebuildSettings | PlainMessage<UpdateConfigurationRequest_PrebuildSettings> | undefined): boolean {
    return proto3.util.equals(UpdateConfigurationRequest_PrebuildSettings, a, b);
  }
}

/**
 * @generated from message devpod.v1.UpdateConfigurationRequest.WorkspaceSettings
 */
export class UpdateConfigurationRequest_WorkspaceSettings extends Message<UpdateConfigurationRequest_WorkspaceSettings> {
  /**
   * @generated from field: optional string workspace_class = 1;
   */
  workspaceClass?: string;

  /**
   * restricted_workspace_classes specifies the workspace classes that are NOT allowed to be used in this configuration.
   * If empty, all workspace classes are allowed.
   * Only updates if update_restricted_workspace_classes is true.
   *
   * @generated from field: repeated string restricted_workspace_classes = 2;
   */
  restrictedWorkspaceClasses: string[] = [];

  /**
   * Specifies whether restricted_workspace_classes should be updated.
   *
   * @generated from field: optional bool update_restricted_workspace_classes = 3;
   */
  updateRestrictedWorkspaceClasses?: boolean;

  /**
   * restricted_editor_names specifies the editor names that are NOT allowed to be used in this configuration.
   * If empty, all editors are allowed.
   * Only updates if update_restricted_editor_names is true.
   *
   * @generated from field: repeated string restricted_editor_names = 4;
   */
  restrictedEditorNames: string[] = [];

  /**
   * Specifies whether restricted_editor_names should be updated.
   *
   * @generated from field: optional bool update_restricted_editor_names = 5;
   */
  updateRestrictedEditorNames?: boolean;

  /**
   * Enable automatic authentication for docker daemon with all credentials specified in GITPOD_IMAGE_AUTH
   *
   * @generated from field: optional bool enable_dockerd_authentication = 6;
   */
  enableDockerdAuthentication?: boolean;

  constructor(data?: PartialMessage<UpdateConfigurationRequest_WorkspaceSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.UpdateConfigurationRequest.WorkspaceSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workspace_class", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "restricted_workspace_classes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "update_restricted_workspace_classes", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "restricted_editor_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "update_restricted_editor_names", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 6, name: "enable_dockerd_authentication", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConfigurationRequest_WorkspaceSettings {
    return new UpdateConfigurationRequest_WorkspaceSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest_WorkspaceSettings {
    return new UpdateConfigurationRequest_WorkspaceSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConfigurationRequest_WorkspaceSettings {
    return new UpdateConfigurationRequest_WorkspaceSettings().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConfigurationRequest_WorkspaceSettings | PlainMessage<UpdateConfigurationRequest_WorkspaceSettings> | undefined, b: UpdateConfigurationRequest_WorkspaceSettings | PlainMessage<UpdateConfigurationRequest_WorkspaceSettings> | undefined): boolean {
    return proto3.util.equals(UpdateConfigurationRequest_WorkspaceSettings, a, b);
  }
}

/**
 * @generated from message devpod.v1.UpdateConfigurationResponse
 */
export class UpdateConfigurationResponse extends Message<UpdateConfigurationResponse> {
  /**
   * @generated from field: devpod.v1.Configuration configuration = 1;
   */
  configuration?: Configuration;

  constructor(data?: PartialMessage<UpdateConfigurationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.UpdateConfigurationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration", kind: "message", T: Configuration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConfigurationResponse {
    return new UpdateConfigurationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConfigurationResponse {
    return new UpdateConfigurationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConfigurationResponse {
    return new UpdateConfigurationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConfigurationResponse | PlainMessage<UpdateConfigurationResponse> | undefined, b: UpdateConfigurationResponse | PlainMessage<UpdateConfigurationResponse> | undefined): boolean {
    return proto3.util.equals(UpdateConfigurationResponse, a, b);
  }
}

/**
 * @generated from message devpod.v1.DeleteConfigurationRequest
 */
export class DeleteConfigurationRequest extends Message<DeleteConfigurationRequest> {
  /**
   * @generated from field: string configuration_id = 1;
   */
  configurationId = "";

  constructor(data?: PartialMessage<DeleteConfigurationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.DeleteConfigurationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configuration_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConfigurationRequest {
    return new DeleteConfigurationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConfigurationRequest {
    return new DeleteConfigurationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConfigurationRequest {
    return new DeleteConfigurationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteConfigurationRequest | PlainMessage<DeleteConfigurationRequest> | undefined, b: DeleteConfigurationRequest | PlainMessage<DeleteConfigurationRequest> | undefined): boolean {
    return proto3.util.equals(DeleteConfigurationRequest, a, b);
  }
}

/**
 * @generated from message devpod.v1.DeleteConfigurationResponse
 */
export class DeleteConfigurationResponse extends Message<DeleteConfigurationResponse> {
  constructor(data?: PartialMessage<DeleteConfigurationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devpod.v1.DeleteConfigurationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConfigurationResponse {
    return new DeleteConfigurationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConfigurationResponse {
    return new DeleteConfigurationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConfigurationResponse {
    return new DeleteConfigurationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteConfigurationResponse | PlainMessage<DeleteConfigurationResponse> | undefined, b: DeleteConfigurationResponse | PlainMessage<DeleteConfigurationResponse> | undefined): boolean {
    return proto3.util.equals(DeleteConfigurationResponse, a, b);
  }
}
