// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: controller/api/resources/policies/v1/policy.proto

package policies

import (
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The id of the storage policy.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The scope id of this policy. This must be defined for policy creation, but
	// is otherwise output only.
	ScopeId string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. Scope information for this resource.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Optional user-set description for identification purposes.
	Description *wrapperspb.StringValue `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The time this resource was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_time,proto3" json:"created_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=updated_time,proto3" json:"updated_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The policy's type.
	Type string `protobuf:"bytes,80,opt,name=type,proto3" json:"type,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Version is used in mutation requests, after the initial creation, to ensure
	// this resource has not changed. The mutation will fail if the version does
	// not match the latest known good version.
	Version uint32 `protobuf:"varint,90,opt,name=version,proto3" json:"version,omitempty" class:"public"` // @gotags: `class:"public"`
	// The attributes that are applicable to each policy type.
	//
	// Types that are assignable to Attrs:
	//
	//	*Policy_Attributes
	//	*Policy_StoragePolicyAttributes
	Attrs isPolicy_Attrs `protobuf_oneof:"attrs"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_policies_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Policy) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Policy) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Policy) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Policy) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Policy) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Policy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Policy) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (m *Policy) GetAttrs() isPolicy_Attrs {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (x *Policy) GetAttributes() *structpb.Struct {
	if x, ok := x.GetAttrs().(*Policy_Attributes); ok {
		return x.Attributes
	}
	return nil
}

func (x *Policy) GetStoragePolicyAttributes() *StoragePolicyAttributes {
	if x, ok := x.GetAttrs().(*Policy_StoragePolicyAttributes); ok {
		return x.StoragePolicyAttributes
	}
	return nil
}

func (x *Policy) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

type isPolicy_Attrs interface {
	isPolicy_Attrs()
}

type Policy_Attributes struct {
	Attributes *structpb.Struct `protobuf:"bytes,100,opt,name=attributes,proto3,oneof"`
}

type Policy_StoragePolicyAttributes struct {
	StoragePolicyAttributes *StoragePolicyAttributes `protobuf:"bytes,101,opt,name=storage_policy_attributes,json=storagePolicyAttributes,proto3,oneof"`
}

func (*Policy_Attributes) isPolicy_Attrs() {}

func (*Policy_StoragePolicyAttributes) isPolicy_Attrs() {}

type StoragePolicyAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetainFor   *StoragePolicyRetainFor   `protobuf:"bytes,10,opt,name=retain_for,proto3" json:"retain_for,omitempty"`
	DeleteAfter *StoragePolicyDeleteAfter `protobuf:"bytes,20,opt,name=delete_after,proto3" json:"delete_after,omitempty"`
}

func (x *StoragePolicyAttributes) Reset() {
	*x = StoragePolicyAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoragePolicyAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoragePolicyAttributes) ProtoMessage() {}

func (x *StoragePolicyAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoragePolicyAttributes.ProtoReflect.Descriptor instead.
func (*StoragePolicyAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_policies_v1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *StoragePolicyAttributes) GetRetainFor() *StoragePolicyRetainFor {
	if x != nil {
		return x.RetainFor
	}
	return nil
}

func (x *StoragePolicyAttributes) GetDeleteAfter() *StoragePolicyDeleteAfter {
	if x != nil {
		return x.DeleteAfter
	}
	return nil
}

type StoragePolicyRetainFor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// days is the number of days for which a session recording will be
	// retained. Must be provided.
	Days int32 `protobuf:"varint,10,opt,name=days,proto3" json:"days,omitempty" class:"public"` // @gotags: `class:"public"`
	// overridable signals whether this storage policy's retention duration can be
	// overridden.
	Overridable *wrapperspb.BoolValue `protobuf:"bytes,20,opt,name=overridable,proto3" json:"overridable,omitempty" class:"public"` // @gotags: `class:"public"`
}

func (x *StoragePolicyRetainFor) Reset() {
	*x = StoragePolicyRetainFor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoragePolicyRetainFor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoragePolicyRetainFor) ProtoMessage() {}

func (x *StoragePolicyRetainFor) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoragePolicyRetainFor.ProtoReflect.Descriptor instead.
func (*StoragePolicyRetainFor) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_policies_v1_policy_proto_rawDescGZIP(), []int{2}
}

func (x *StoragePolicyRetainFor) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *StoragePolicyRetainFor) GetOverridable() *wrapperspb.BoolValue {
	if x != nil {
		return x.Overridable
	}
	return nil
}

type StoragePolicyDeleteAfter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// days is the number of days after which a session recording will be
	// automatically deleted.
	Days int32 `protobuf:"varint,10,opt,name=days,proto3" json:"days,omitempty" class:"public"` // @gotags: `class:"public"`
	// overridable signals whether this storage policy's deletion policy can be
	// overridden.
	Overridable *wrapperspb.BoolValue `protobuf:"bytes,20,opt,name=overridable,proto3" json:"overridable,omitempty" class:"public"` // @gotags: `class:"public"`
}

func (x *StoragePolicyDeleteAfter) Reset() {
	*x = StoragePolicyDeleteAfter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoragePolicyDeleteAfter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoragePolicyDeleteAfter) ProtoMessage() {}

func (x *StoragePolicyDeleteAfter) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_policies_v1_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoragePolicyDeleteAfter.ProtoReflect.Descriptor instead.
func (*StoragePolicyDeleteAfter) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_policies_v1_policy_proto_rawDescGZIP(), []int{3}
}

func (x *StoragePolicyDeleteAfter) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *StoragePolicyDeleteAfter) GetOverridable() *wrapperspb.BoolValue {
	if x != nil {
		return x.Overridable
	}
	return nil
}

var File_controller_api_resources_policies_v1_policy_proto protoreflect.FileDescriptor

var file_controller_api_resources_policies_v1_policy_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf8, 0x05, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x46, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x14, 0xa0, 0xda, 0x29, 0x01, 0xc2,
	0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x22, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd,
	0x29, 0x1a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x42, 0x0f, 0xa0, 0xda, 0x29, 0x01, 0x9a, 0xe3, 0x29, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x1f, 0xa0, 0xda, 0x29, 0x01, 0x9a, 0xe3, 0x29, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x48, 0x00, 0x52, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xac, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22, 0xdb, 0x01, 0x0a,
	0x17, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x12, 0x62, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6e, 0x46, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x2f, 0xc2, 0xdd, 0x29, 0x2b, 0x0a, 0x1a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x6f, 0x72,
	0x2e, 0x64, 0x61, 0x79, 0x73, 0x12, 0x0d, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x6f, 0x72,
	0x44, 0x61, 0x79, 0x73, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x7f, 0x0a, 0x0b, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x41, 0xc2, 0xdd, 0x29,
	0x3d, 0x0a, 0x21, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x44,
	0x61, 0x79, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0b,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x18,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x33, 0xc2, 0xdd, 0x29, 0x2f, 0x0a, 0x1c, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x79, 0x73, 0x12, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x61, 0x79, 0x73, 0x52, 0x04, 0x64, 0x61, 0x79,
	0x73, 0x12, 0x83, 0x01, 0x0a, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x45, 0xc2, 0xdd, 0x29, 0x41, 0x0a, 0x23, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x61, 0x79, 0x73, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x3b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_policies_v1_policy_proto_rawDescOnce sync.Once
	file_controller_api_resources_policies_v1_policy_proto_rawDescData = file_controller_api_resources_policies_v1_policy_proto_rawDesc
)

func file_controller_api_resources_policies_v1_policy_proto_rawDescGZIP() []byte {
	file_controller_api_resources_policies_v1_policy_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_policies_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_policies_v1_policy_proto_rawDescData)
	})
	return file_controller_api_resources_policies_v1_policy_proto_rawDescData
}

var file_controller_api_resources_policies_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_api_resources_policies_v1_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),                   // 0: controller.api.resources.policies.v1.Policy
	(*StoragePolicyAttributes)(nil),  // 1: controller.api.resources.policies.v1.StoragePolicyAttributes
	(*StoragePolicyRetainFor)(nil),   // 2: controller.api.resources.policies.v1.StoragePolicyRetainFor
	(*StoragePolicyDeleteAfter)(nil), // 3: controller.api.resources.policies.v1.StoragePolicyDeleteAfter
	(*scopes.ScopeInfo)(nil),         // 4: controller.api.resources.scopes.v1.ScopeInfo
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*structpb.Struct)(nil),          // 7: google.protobuf.Struct
	(*wrapperspb.BoolValue)(nil),     // 8: google.protobuf.BoolValue
}
var file_controller_api_resources_policies_v1_policy_proto_depIdxs = []int32{
	4,  // 0: controller.api.resources.policies.v1.Policy.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	5,  // 1: controller.api.resources.policies.v1.Policy.name:type_name -> google.protobuf.StringValue
	5,  // 2: controller.api.resources.policies.v1.Policy.description:type_name -> google.protobuf.StringValue
	6,  // 3: controller.api.resources.policies.v1.Policy.created_time:type_name -> google.protobuf.Timestamp
	6,  // 4: controller.api.resources.policies.v1.Policy.updated_time:type_name -> google.protobuf.Timestamp
	7,  // 5: controller.api.resources.policies.v1.Policy.attributes:type_name -> google.protobuf.Struct
	1,  // 6: controller.api.resources.policies.v1.Policy.storage_policy_attributes:type_name -> controller.api.resources.policies.v1.StoragePolicyAttributes
	2,  // 7: controller.api.resources.policies.v1.StoragePolicyAttributes.retain_for:type_name -> controller.api.resources.policies.v1.StoragePolicyRetainFor
	3,  // 8: controller.api.resources.policies.v1.StoragePolicyAttributes.delete_after:type_name -> controller.api.resources.policies.v1.StoragePolicyDeleteAfter
	8,  // 9: controller.api.resources.policies.v1.StoragePolicyRetainFor.overridable:type_name -> google.protobuf.BoolValue
	8,  // 10: controller.api.resources.policies.v1.StoragePolicyDeleteAfter.overridable:type_name -> google.protobuf.BoolValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_controller_api_resources_policies_v1_policy_proto_init() }
func file_controller_api_resources_policies_v1_policy_proto_init() {
	if File_controller_api_resources_policies_v1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_policies_v1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_controller_api_resources_policies_v1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoragePolicyAttributes); i {
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
		file_controller_api_resources_policies_v1_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoragePolicyRetainFor); i {
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
		file_controller_api_resources_policies_v1_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoragePolicyDeleteAfter); i {
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
	file_controller_api_resources_policies_v1_policy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Policy_Attributes)(nil),
		(*Policy_StoragePolicyAttributes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_api_resources_policies_v1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_policies_v1_policy_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_policies_v1_policy_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_policies_v1_policy_proto_msgTypes,
	}.Build()
	File_controller_api_resources_policies_v1_policy_proto = out.File
	file_controller_api_resources_policies_v1_policy_proto_rawDesc = nil
	file_controller_api_resources_policies_v1_policy_proto_goTypes = nil
	file_controller_api_resources_policies_v1_policy_proto_depIdxs = nil
}
