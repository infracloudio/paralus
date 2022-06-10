// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/systempb/v3/project.proto

package systemv3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v31 "github.com/paralus/paralus/proto/types/commonpb/v3"
	v3 "github.com/paralus/paralus/proto/types/userpb/v3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProjectSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default               bool                       `protobuf:"varint,1,opt,name=default,proto3" json:"default,omitempty"`
	ProjectNamespaceRoles []*v3.ProjectNamespaceRole `protobuf:"bytes,2,rep,name=projectNamespaceRoles,proto3" json:"projectNamespaceRoles,omitempty"`
	UserRoles             []*v3.UserRole             `protobuf:"bytes,3,rep,name=userRoles,proto3" json:"userRoles,omitempty"`
}

func (x *ProjectSpec) Reset() {
	*x = ProjectSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSpec) ProtoMessage() {}

func (x *ProjectSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectSpec.ProtoReflect.Descriptor instead.
func (*ProjectSpec) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectSpec) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *ProjectSpec) GetProjectNamespaceRoles() []*v3.ProjectNamespaceRole {
	if x != nil {
		return x.ProjectNamespaceRoles
	}
	return nil
}

func (x *ProjectSpec) GetUserRoles() []*v3.UserRole {
	if x != nil {
		return x.UserRoles
	}
	return nil
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string        `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string        `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v31.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *ProjectSpec  `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *v31.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_project_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Project) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Project) GetMetadata() *v31.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Project) GetSpec() *ProjectSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Project) GetStatus() *v31.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ProjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string            `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string            `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v31.ListMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Items      []*Project        `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProjectList) Reset() {
	*x = ProjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectList) ProtoMessage() {}

func (x *ProjectList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectList.ProtoReflect.Descriptor instead.
func (*ProjectList) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectList) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ProjectList) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ProjectList) GetMetadata() *v31.ListMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ProjectList) GetItems() []*Project {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_types_systempb_v3_project_proto protoreflect.FileDescriptor

var file_proto_types_systempb_v3_project_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x76,
	0x33, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x03, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x6b, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x51, 0x92, 0x41, 0x4e, 0x2a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x32,
	0x43, 0x66, 0x6c, 0x61, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x20, 0x69, 0x66, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0xb4, 0x01,
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x4d,
	0x92, 0x41, 0x4a, 0x2a, 0x15, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0x31, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2c, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2c, 0x20,
	0x72, 0x6f, 0x6c, 0x65, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x15, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x3d, 0x92,
	0x41, 0x3a, 0x2a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0x2d, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x20, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x3a, 0x33, 0x92, 0x41, 0x30, 0x0a, 0x2e, 0x2a, 0x15,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x04, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x65, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0x92, 0x41,
	0x42, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x1b,
	0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x33, 0x40, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0x92,
	0x41, 0x27, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x14, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x6a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x27, 0x92, 0x41, 0x24, 0x2a, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x42, 0x1f, 0x92, 0x41, 0x1c, 0x2a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x32, 0x14,
	0x53, 0x70, 0x65, 0x63, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x62, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x25, 0x92, 0x41, 0x22, 0x2a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x16, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3d,
	0x92, 0x41, 0x3a, 0x0a, 0x38, 0x2a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0xd2, 0x01, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0xd2, 0x01, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0xd2, 0x01, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xd2, 0x01, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xbc, 0x03,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6a, 0x0a,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x4a, 0x92, 0x41, 0x47, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x32, 0x20, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38,
	0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x33, 0x40, 0x01, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0x92, 0x41, 0x2c, 0x2a, 0x04, 0x4b, 0x69,
	0x6e, 0x64, 0x32, 0x19, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x79,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x32, 0x92, 0x41, 0x2f,
	0x2a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x1d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5f, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x23,
	0x92, 0x41, 0x20, 0x2a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x40, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x20, 0x92, 0x41, 0x1d, 0x0a,
	0x1b, 0x2a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x0c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x42, 0xfd, 0x01, 0x0a,
	0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33,
	0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x70,
	0x62, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x76, 0x33, 0xa2, 0x02, 0x04,
	0x50, 0x44, 0x54, 0x53, 0xaa, 0x02, 0x1b, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x44,
	0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x56, 0x33, 0xca, 0x02, 0x1b, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x5c, 0x44, 0x65, 0x76,
	0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33,
	0xe2, 0x02, 0x27, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1f, 0x50, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_systempb_v3_project_proto_rawDescOnce sync.Once
	file_proto_types_systempb_v3_project_proto_rawDescData = file_proto_types_systempb_v3_project_proto_rawDesc
)

func file_proto_types_systempb_v3_project_proto_rawDescGZIP() []byte {
	file_proto_types_systempb_v3_project_proto_rawDescOnce.Do(func() {
		file_proto_types_systempb_v3_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_systempb_v3_project_proto_rawDescData)
	})
	return file_proto_types_systempb_v3_project_proto_rawDescData
}

var file_proto_types_systempb_v3_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_types_systempb_v3_project_proto_goTypes = []interface{}{
	(*ProjectSpec)(nil),             // 0: paralus.dev.types.system.v3.ProjectSpec
	(*Project)(nil),                 // 1: paralus.dev.types.system.v3.Project
	(*ProjectList)(nil),             // 2: paralus.dev.types.system.v3.ProjectList
	(*v3.ProjectNamespaceRole)(nil), // 3: paralus.dev.types.user.v3.ProjectNamespaceRole
	(*v3.UserRole)(nil),             // 4: paralus.dev.types.user.v3.UserRole
	(*v31.Metadata)(nil),            // 5: paralus.dev.types.common.v3.Metadata
	(*v31.Status)(nil),              // 6: paralus.dev.types.common.v3.Status
	(*v31.ListMetadata)(nil),        // 7: paralus.dev.types.common.v3.ListMetadata
}
var file_proto_types_systempb_v3_project_proto_depIdxs = []int32{
	3, // 0: paralus.dev.types.system.v3.ProjectSpec.projectNamespaceRoles:type_name -> paralus.dev.types.user.v3.ProjectNamespaceRole
	4, // 1: paralus.dev.types.system.v3.ProjectSpec.userRoles:type_name -> paralus.dev.types.user.v3.UserRole
	5, // 2: paralus.dev.types.system.v3.Project.metadata:type_name -> paralus.dev.types.common.v3.Metadata
	0, // 3: paralus.dev.types.system.v3.Project.spec:type_name -> paralus.dev.types.system.v3.ProjectSpec
	6, // 4: paralus.dev.types.system.v3.Project.status:type_name -> paralus.dev.types.common.v3.Status
	7, // 5: paralus.dev.types.system.v3.ProjectList.metadata:type_name -> paralus.dev.types.common.v3.ListMetadata
	1, // 6: paralus.dev.types.system.v3.ProjectList.items:type_name -> paralus.dev.types.system.v3.Project
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_types_systempb_v3_project_proto_init() }
func file_proto_types_systempb_v3_project_proto_init() {
	if File_proto_types_systempb_v3_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_systempb_v3_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSpec); i {
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
		file_proto_types_systempb_v3_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_proto_types_systempb_v3_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectList); i {
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
			RawDescriptor: file_proto_types_systempb_v3_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_systempb_v3_project_proto_goTypes,
		DependencyIndexes: file_proto_types_systempb_v3_project_proto_depIdxs,
		MessageInfos:      file_proto_types_systempb_v3_project_proto_msgTypes,
	}.Build()
	File_proto_types_systempb_v3_project_proto = out.File
	file_proto_types_systempb_v3_project_proto_rawDesc = nil
	file_proto_types_systempb_v3_project_proto_goTypes = nil
	file_proto_types_systempb_v3_project_proto_depIdxs = nil
}
