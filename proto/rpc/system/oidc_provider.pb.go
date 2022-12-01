// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/rpc/system/oidc_provider.proto

package systemv3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v3 "github.com/paralus/paralus/proto/types/systempb/v3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_rpc_system_oidc_provider_proto protoreflect.FileDescriptor

var file_proto_rpc_system_oidc_provider_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfc, 0x06, 0x0a, 0x13, 0x4f, 0x49,
	0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xd3, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x49, 0x44, 0x43,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x67,
	0x92, 0x41, 0x3f, 0x4a, 0x3d, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x36, 0x0a, 0x34, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x6f, 0x69, 0x64, 0x63,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x9b, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x49, 0x44,
	0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e,
	0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0xa1, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x1a, 0x2a, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xcd, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x49, 0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x29, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x49,
	0x44, 0x43, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x74, 0x92, 0x41, 0x3f, 0x4a, 0x3d, 0x0a, 0x03, 0x32, 0x30, 0x34, 0x12, 0x36,
	0x0a, 0x34, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x6f, 0x69, 0x64, 0x63, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x69, 0x73,
	0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x2a, 0x2a, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0xd9, 0x04, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x4f, 0x69, 0x64, 0x63,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3b, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x76, 0x33, 0xa2, 0x02, 0x04, 0x50, 0x44, 0x52, 0x53, 0xaa, 0x02, 0x19, 0x50,
	0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x19, 0x50, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x25, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x5c, 0x44,
	0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x50,
	0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x52, 0x70, 0x63,
	0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x56, 0x33, 0x92, 0x41, 0xe4, 0x02,
	0x12, 0x2b, 0x0a, 0x15, 0x4f, 0x49, 0x64, 0x43, 0x20, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x20, 0x44, 0x65, 0x76, 0x32, 0x03, 0x33, 0x2e, 0x30, 0x2a, 0x01, 0x02,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12,
	0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30,
	0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12,
	0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x3a, 0x0a, 0x27, 0x0a, 0x0a, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x19, 0x08, 0x02, 0x1a, 0x13, 0x58, 0x2d, 0x50,
	0x41, 0x52, 0x41, 0x4c, 0x55, 0x53, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x45, 0x59, 0x49, 0x44,
	0x20, 0x02, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_rpc_system_oidc_provider_proto_goTypes = []interface{}{
	(*v3.OIDCProvider)(nil),     // 0: paralus.dev.types.system.v3.OIDCProvider
	(*emptypb.Empty)(nil),       // 1: google.protobuf.Empty
	(*v3.OIDCProviderList)(nil), // 2: paralus.dev.types.system.v3.OIDCProviderList
}
var file_proto_rpc_system_oidc_provider_proto_depIdxs = []int32{
	0, // 0: paralus.dev.rpc.system.v3.OIDCProviderService.CreateOIDCProvider:input_type -> paralus.dev.types.system.v3.OIDCProvider
	0, // 1: paralus.dev.rpc.system.v3.OIDCProviderService.GetOIDCProvider:input_type -> paralus.dev.types.system.v3.OIDCProvider
	1, // 2: paralus.dev.rpc.system.v3.OIDCProviderService.ListOIDCProvider:input_type -> google.protobuf.Empty
	0, // 3: paralus.dev.rpc.system.v3.OIDCProviderService.UpdateOIDCProvider:input_type -> paralus.dev.types.system.v3.OIDCProvider
	0, // 4: paralus.dev.rpc.system.v3.OIDCProviderService.DeleteOIDCProvider:input_type -> paralus.dev.types.system.v3.OIDCProvider
	0, // 5: paralus.dev.rpc.system.v3.OIDCProviderService.CreateOIDCProvider:output_type -> paralus.dev.types.system.v3.OIDCProvider
	0, // 6: paralus.dev.rpc.system.v3.OIDCProviderService.GetOIDCProvider:output_type -> paralus.dev.types.system.v3.OIDCProvider
	2, // 7: paralus.dev.rpc.system.v3.OIDCProviderService.ListOIDCProvider:output_type -> paralus.dev.types.system.v3.OIDCProviderList
	0, // 8: paralus.dev.rpc.system.v3.OIDCProviderService.UpdateOIDCProvider:output_type -> paralus.dev.types.system.v3.OIDCProvider
	1, // 9: paralus.dev.rpc.system.v3.OIDCProviderService.DeleteOIDCProvider:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_system_oidc_provider_proto_init() }
func file_proto_rpc_system_oidc_provider_proto_init() {
	if File_proto_rpc_system_oidc_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rpc_system_oidc_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_system_oidc_provider_proto_goTypes,
		DependencyIndexes: file_proto_rpc_system_oidc_provider_proto_depIdxs,
	}.Build()
	File_proto_rpc_system_oidc_provider_proto = out.File
	file_proto_rpc_system_oidc_provider_proto_rawDesc = nil
	file_proto_rpc_system_oidc_provider_proto_goTypes = nil
	file_proto_rpc_system_oidc_provider_proto_depIdxs = nil
}
