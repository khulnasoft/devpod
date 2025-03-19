// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: devpod/experimental/v1/identityprovider.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetIDTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string   `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Audience    []string `protobuf:"bytes,2,rep,name=audience,proto3" json:"audience,omitempty"`
	Scope       string   `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *GetIDTokenRequest) Reset() {
	*x = GetIDTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devpod_experimental_v1_identityprovider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIDTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIDTokenRequest) ProtoMessage() {}

func (x *GetIDTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devpod_experimental_v1_identityprovider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIDTokenRequest.ProtoReflect.Descriptor instead.
func (*GetIDTokenRequest) Descriptor() ([]byte, []int) {
	return file_devpod_experimental_v1_identityprovider_proto_rawDescGZIP(), []int{0}
}

func (x *GetIDTokenRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *GetIDTokenRequest) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *GetIDTokenRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type GetIDTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetIDTokenResponse) Reset() {
	*x = GetIDTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devpod_experimental_v1_identityprovider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIDTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIDTokenResponse) ProtoMessage() {}

func (x *GetIDTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devpod_experimental_v1_identityprovider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIDTokenResponse.ProtoReflect.Descriptor instead.
func (*GetIDTokenResponse) Descriptor() ([]byte, []int) {
	return file_devpod_experimental_v1_identityprovider_proto_rawDescGZIP(), []int{1}
}

func (x *GetIDTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_devpod_experimental_v1_identityprovider_proto protoreflect.FileDescriptor

var file_devpod_experimental_v1_identityprovider_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x68, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x44,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x80, 0x01,
	0x0a, 0x17, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x6b, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69,
	0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devpod_experimental_v1_identityprovider_proto_rawDescOnce sync.Once
	file_devpod_experimental_v1_identityprovider_proto_rawDescData = file_devpod_experimental_v1_identityprovider_proto_rawDesc
)

func file_devpod_experimental_v1_identityprovider_proto_rawDescGZIP() []byte {
	file_devpod_experimental_v1_identityprovider_proto_rawDescOnce.Do(func() {
		file_devpod_experimental_v1_identityprovider_proto_rawDescData = protoimpl.X.CompressGZIP(file_devpod_experimental_v1_identityprovider_proto_rawDescData)
	})
	return file_devpod_experimental_v1_identityprovider_proto_rawDescData
}

var file_devpod_experimental_v1_identityprovider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_devpod_experimental_v1_identityprovider_proto_goTypes = []interface{}{
	(*GetIDTokenRequest)(nil),  // 0: devpod.experimental.v1.GetIDTokenRequest
	(*GetIDTokenResponse)(nil), // 1: devpod.experimental.v1.GetIDTokenResponse
}
var file_devpod_experimental_v1_identityprovider_proto_depIdxs = []int32{
	0, // 0: devpod.experimental.v1.IdentityProviderService.GetIDToken:input_type -> devpod.experimental.v1.GetIDTokenRequest
	1, // 1: devpod.experimental.v1.IdentityProviderService.GetIDToken:output_type -> devpod.experimental.v1.GetIDTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_devpod_experimental_v1_identityprovider_proto_init() }
func file_devpod_experimental_v1_identityprovider_proto_init() {
	if File_devpod_experimental_v1_identityprovider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devpod_experimental_v1_identityprovider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIDTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_devpod_experimental_v1_identityprovider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIDTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_devpod_experimental_v1_identityprovider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devpod_experimental_v1_identityprovider_proto_goTypes,
		DependencyIndexes: file_devpod_experimental_v1_identityprovider_proto_depIdxs,
		MessageInfos:      file_devpod_experimental_v1_identityprovider_proto_msgTypes,
	}.Build()
	File_devpod_experimental_v1_identityprovider_proto = out.File
	file_devpod_experimental_v1_identityprovider_proto_rawDesc = nil
	file_devpod_experimental_v1_identityprovider_proto_goTypes = nil
	file_devpod_experimental_v1_identityprovider_proto_depIdxs = nil
}
