// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/userpb/v3/role.proto

package userv3

import (
	v3 "github.com/RafaySystems/rcloud-base/components/common/proto/types/commonpb/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string       `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string       `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *RoleSpec    `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *v3.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Role) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Role) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Role) GetSpec() *RoleSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Role) GetStatus() *v3.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RoleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rolepermissions []string `protobuf:"bytes,1,rep,name=rolepermissions,proto3" json:"rolepermissions,omitempty"`
	IsGlobal        bool     `protobuf:"varint,2,opt,name=isGlobal,proto3" json:"isGlobal,omitempty"`
	Scope           string   `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *RoleSpec) Reset() {
	*x = RoleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSpec) ProtoMessage() {}

func (x *RoleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSpec.ProtoReflect.Descriptor instead.
func (*RoleSpec) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleSpec) GetRolepermissions() []string {
	if x != nil {
		return x.Rolepermissions
	}
	return nil
}

func (x *RoleSpec) GetIsGlobal() bool {
	if x != nil {
		return x.IsGlobal
	}
	return false
}

func (x *RoleSpec) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string           `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string           `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.ListMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Items      []*Role          `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleList) Reset() {
	*x = RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleList) ProtoMessage() {}

func (x *RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleList.ProtoReflect.Descriptor instead.
func (*RoleList) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleList) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *RoleList) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *RoleList) GetMetadata() *v3.ListMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RoleList) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_types_userpb_v3_role_proto protoreflect.FileDescriptor

var file_proto_types_userpb_v3_role_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x04,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0x92, 0x41, 0x45, 0x2a,
	0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x52, 0x6f, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x33, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x3e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a,
	0x92, 0x41, 0x27, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x19, 0x4b, 0x69, 0x6e, 0x64, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x3a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x6d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0x92, 0x41, 0x29, 0x2a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x1d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x5b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x42, 0x24, 0x92, 0x41, 0x21, 0x2a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x32, 0x19, 0x53, 0x70, 0x65,
	0x63, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x60, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x25, 0x92, 0x41, 0x22, 0x2a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x16, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x37,
	0x92, 0x41, 0x34, 0x0a, 0x32, 0x2a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x32, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0xd2, 0x01, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xd2, 0x01,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0xd2, 0x01, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xd2, 0x01, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x95, 0x02, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x59, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2f, 0x92,
	0x41, 0x2c, 0x2a, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x20, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x18, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x0f,
	0x72, 0x6f, 0x6c, 0x65, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x4b, 0x0a, 0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x2f, 0x92, 0x41, 0x2c, 0x2a, 0x08, 0x49, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x32, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x79, 0x20, 0x69, 0x66, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x72, 0x6f,
	0x6c, 0x65, 0x52, 0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0x92, 0x41, 0x19,
	0x2a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x32, 0x10, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x3a, 0x2d, 0x92, 0x41, 0x2a, 0x0a, 0x28, 0x2a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x20, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x12, 0x52, 0x6f, 0x6c,
	0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x99, 0x03, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x39, 0x92, 0x41, 0x36, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0x25, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64,
	0x32, 0x1e, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f,
	0x6c, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x78, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x33, 0x92, 0x41, 0x30, 0x2a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x22, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x59, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x24, 0x92, 0x41, 0x21, 0x2a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x40, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x1c, 0x92,
	0x41, 0x19, 0x0a, 0x17, 0x2a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x09,
	0x52, 0x6f, 0x6c, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x40, 0x01, 0x42, 0xff, 0x01, 0x0a, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x33, 0xa2, 0x02,
	0x04, 0x52, 0x44, 0x54, 0x55, 0xaa, 0x02, 0x17, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65,
	0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca,
	0x02, 0x17, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x23, 0x52, 0x61, 0x66, 0x61,
	0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72,
	0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_userpb_v3_role_proto_rawDescOnce sync.Once
	file_proto_types_userpb_v3_role_proto_rawDescData = file_proto_types_userpb_v3_role_proto_rawDesc
)

func file_proto_types_userpb_v3_role_proto_rawDescGZIP() []byte {
	file_proto_types_userpb_v3_role_proto_rawDescOnce.Do(func() {
		file_proto_types_userpb_v3_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_userpb_v3_role_proto_rawDescData)
	})
	return file_proto_types_userpb_v3_role_proto_rawDescData
}

var file_proto_types_userpb_v3_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_types_userpb_v3_role_proto_goTypes = []interface{}{
	(*Role)(nil),            // 0: rafay.dev.types.user.v3.Role
	(*RoleSpec)(nil),        // 1: rafay.dev.types.user.v3.RoleSpec
	(*RoleList)(nil),        // 2: rafay.dev.types.user.v3.RoleList
	(*v3.Metadata)(nil),     // 3: rafay.dev.types.common.v3.Metadata
	(*v3.Status)(nil),       // 4: rafay.dev.types.common.v3.Status
	(*v3.ListMetadata)(nil), // 5: rafay.dev.types.common.v3.ListMetadata
}
var file_proto_types_userpb_v3_role_proto_depIdxs = []int32{
	3, // 0: rafay.dev.types.user.v3.Role.metadata:type_name -> rafay.dev.types.common.v3.Metadata
	1, // 1: rafay.dev.types.user.v3.Role.spec:type_name -> rafay.dev.types.user.v3.RoleSpec
	4, // 2: rafay.dev.types.user.v3.Role.status:type_name -> rafay.dev.types.common.v3.Status
	5, // 3: rafay.dev.types.user.v3.RoleList.metadata:type_name -> rafay.dev.types.common.v3.ListMetadata
	0, // 4: rafay.dev.types.user.v3.RoleList.items:type_name -> rafay.dev.types.user.v3.Role
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_userpb_v3_role_proto_init() }
func file_proto_types_userpb_v3_role_proto_init() {
	if File_proto_types_userpb_v3_role_proto != nil {
		return
	}
	file_proto_types_userpb_v3_rolepermission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_userpb_v3_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_proto_types_userpb_v3_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSpec); i {
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
		file_proto_types_userpb_v3_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleList); i {
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
			RawDescriptor: file_proto_types_userpb_v3_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_userpb_v3_role_proto_goTypes,
		DependencyIndexes: file_proto_types_userpb_v3_role_proto_depIdxs,
		MessageInfos:      file_proto_types_userpb_v3_role_proto_msgTypes,
	}.Build()
	File_proto_types_userpb_v3_role_proto = out.File
	file_proto_types_userpb_v3_role_proto_rawDesc = nil
	file_proto_types_userpb_v3_role_proto_goTypes = nil
	file_proto_types_userpb_v3_role_proto_depIdxs = nil
}
