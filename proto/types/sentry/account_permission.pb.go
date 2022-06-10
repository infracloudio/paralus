// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/sentry/account_permission.proto

package sentry

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *PermissionURL) Reset() {
	*x = PermissionURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_account_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionURL) ProtoMessage() {}

func (x *PermissionURL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_account_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionURL.ProtoReflect.Descriptor instead.
func (*PermissionURL) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_account_permission_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PermissionURL) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type AccountPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID      string           `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	ProjectID      string           `protobuf:"bytes,2,opt,name=projectID,proto3" json:"projectID,omitempty"`
	OrganizationID string           `protobuf:"bytes,3,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	PartnerID      string           `protobuf:"bytes,4,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	RoleName       string           `protobuf:"bytes,5,opt,name=roleName,proto3" json:"roleName,omitempty"`
	IsGlobal       bool             `protobuf:"varint,6,opt,name=isGlobal,proto3" json:"isGlobal,omitempty"`
	Scope          string           `protobuf:"bytes,7,opt,name=scope,proto3" json:"scope,omitempty"`
	PermissionName string           `protobuf:"bytes,8,opt,name=permissionName,proto3" json:"permissionName,omitempty"`
	BaseURL        string           `protobuf:"bytes,9,opt,name=baseURL,proto3" json:"baseURL,omitempty"`
	Urls           []*PermissionURL `protobuf:"bytes,10,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *AccountPermission) Reset() {
	*x = AccountPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_account_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPermission) ProtoMessage() {}

func (x *AccountPermission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_account_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPermission.ProtoReflect.Descriptor instead.
func (*AccountPermission) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_account_permission_proto_rawDescGZIP(), []int{1}
}

func (x *AccountPermission) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *AccountPermission) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *AccountPermission) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *AccountPermission) GetPartnerID() string {
	if x != nil {
		return x.PartnerID
	}
	return ""
}

func (x *AccountPermission) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *AccountPermission) GetIsGlobal() bool {
	if x != nil {
		return x.IsGlobal
	}
	return false
}

func (x *AccountPermission) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AccountPermission) GetPermissionName() string {
	if x != nil {
		return x.PermissionName
	}
	return ""
}

func (x *AccountPermission) GetBaseURL() string {
	if x != nil {
		return x.BaseURL
	}
	return ""
}

func (x *AccountPermission) GetUrls() []*PermissionURL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type SSOAccountGroupProjectRoleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName              string                 `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	RoleName              string                 `protobuf:"bytes,3,opt,name=roleName,proto3" json:"roleName,omitempty"`
	ProjectID             string                 `protobuf:"bytes,4,opt,name=projectID,proto3" json:"projectID,omitempty"`
	ProjectName           string                 `protobuf:"bytes,5,opt,name=projectName,proto3" json:"projectName,omitempty"`
	Group                 string                 `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	AccountOrganizationID string                 `protobuf:"bytes,7,opt,name=accountOrganizationID,proto3" json:"accountOrganizationID,omitempty"`
	OrganizationID        string                 `protobuf:"bytes,8,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	PartnerID             string                 `protobuf:"bytes,9,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	Scope                 string                 `protobuf:"bytes,10,opt,name=scope,proto3" json:"scope,omitempty"`
	LastLogin             *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=lastLogin,proto3" json:"lastLogin,omitempty"`
	CreatedAt             *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	FirstName             string                 `protobuf:"bytes,13,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName              string                 `protobuf:"bytes,14,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Phone                 string                 `protobuf:"bytes,15,opt,name=phone,proto3" json:"phone,omitempty"`
	Name                  string                 `protobuf:"bytes,16,opt,name=name,proto3" json:"name,omitempty"`
	LastLogout            *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=lastLogout,proto3" json:"lastLogout,omitempty"`
}

func (x *SSOAccountGroupProjectRoleData) Reset() {
	*x = SSOAccountGroupProjectRoleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_account_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSOAccountGroupProjectRoleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSOAccountGroupProjectRoleData) ProtoMessage() {}

func (x *SSOAccountGroupProjectRoleData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_account_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSOAccountGroupProjectRoleData.ProtoReflect.Descriptor instead.
func (*SSOAccountGroupProjectRoleData) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_account_permission_proto_rawDescGZIP(), []int{2}
}

func (x *SSOAccountGroupProjectRoleData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetAccountOrganizationID() string {
	if x != nil {
		return x.AccountOrganizationID
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetPartnerID() string {
	if x != nil {
		return x.PartnerID
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetLastLogin() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLogin
	}
	return nil
}

func (x *SSOAccountGroupProjectRoleData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SSOAccountGroupProjectRoleData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSOAccountGroupProjectRoleData) GetLastLogout() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLogout
	}
	return nil
}

var File_proto_types_sentry_account_permission_proto protoreflect.FileDescriptor

var file_proto_types_sentry_account_permission_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70,
	0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0xe2, 0x02, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x3b, 0x0a,
	0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0xe4, 0x04, 0x0a, 0x1e, 0x53,
	0x53, 0x4f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x15, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x42, 0xe9, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x16, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xa2, 0x02, 0x04, 0x50, 0x44,
	0x54, 0x53, 0xaa, 0x02, 0x18, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x76,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xca, 0x02, 0x18,
	0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xe2, 0x02, 0x24, 0x50, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_sentry_account_permission_proto_rawDescOnce sync.Once
	file_proto_types_sentry_account_permission_proto_rawDescData = file_proto_types_sentry_account_permission_proto_rawDesc
)

func file_proto_types_sentry_account_permission_proto_rawDescGZIP() []byte {
	file_proto_types_sentry_account_permission_proto_rawDescOnce.Do(func() {
		file_proto_types_sentry_account_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_sentry_account_permission_proto_rawDescData)
	})
	return file_proto_types_sentry_account_permission_proto_rawDescData
}

var file_proto_types_sentry_account_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_types_sentry_account_permission_proto_goTypes = []interface{}{
	(*PermissionURL)(nil),                  // 0: paralus.dev.types.sentry.PermissionURL
	(*AccountPermission)(nil),              // 1: paralus.dev.types.sentry.AccountPermission
	(*SSOAccountGroupProjectRoleData)(nil), // 2: paralus.dev.types.sentry.SSOAccountGroupProjectRoleData
	(*timestamppb.Timestamp)(nil),          // 3: google.protobuf.Timestamp
}
var file_proto_types_sentry_account_permission_proto_depIdxs = []int32{
	0, // 0: paralus.dev.types.sentry.AccountPermission.urls:type_name -> paralus.dev.types.sentry.PermissionURL
	3, // 1: paralus.dev.types.sentry.SSOAccountGroupProjectRoleData.lastLogin:type_name -> google.protobuf.Timestamp
	3, // 2: paralus.dev.types.sentry.SSOAccountGroupProjectRoleData.createdAt:type_name -> google.protobuf.Timestamp
	3, // 3: paralus.dev.types.sentry.SSOAccountGroupProjectRoleData.lastLogout:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_types_sentry_account_permission_proto_init() }
func file_proto_types_sentry_account_permission_proto_init() {
	if File_proto_types_sentry_account_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_sentry_account_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionURL); i {
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
		file_proto_types_sentry_account_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPermission); i {
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
		file_proto_types_sentry_account_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSOAccountGroupProjectRoleData); i {
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
			RawDescriptor: file_proto_types_sentry_account_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_sentry_account_permission_proto_goTypes,
		DependencyIndexes: file_proto_types_sentry_account_permission_proto_depIdxs,
		MessageInfos:      file_proto_types_sentry_account_permission_proto_msgTypes,
	}.Build()
	File_proto_types_sentry_account_permission_proto = out.File
	file_proto_types_sentry_account_permission_proto_rawDesc = nil
	file_proto_types_sentry_account_permission_proto_goTypes = nil
	file_proto_types_sentry_account_permission_proto_depIdxs = nil
}
