// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/domain/users/role.proto

package users

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoleName int32

const (
	RoleName_administrators  RoleName = 0
	RoleName_users           RoleName = 1
	RoleName_campaign_owners RoleName = 2
)

// Enum value maps for RoleName.
var (
	RoleName_name = map[int32]string{
		0: "administrators",
		1: "users",
		2: "campaign_owners",
	}
	RoleName_value = map[string]int32{
		"administrators":  0,
		"users":           1,
		"campaign_owners": 2,
	}
)

func (x RoleName) Enum() *RoleName {
	p := new(RoleName)
	*p = x
	return p
}

func (x RoleName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleName) Descriptor() protoreflect.EnumDescriptor {
	return file_publicit_domain_users_role_proto_enumTypes[0].Descriptor()
}

func (RoleName) Type() protoreflect.EnumType {
	return &file_publicit_domain_users_role_proto_enumTypes[0]
}

func (x RoleName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleName.Descriptor instead.
func (RoleName) EnumDescriptor() ([]byte, []int) {
	return file_publicit_domain_users_role_proto_rawDescGZIP(), []int{0}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicit_domain_users_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_publicit_domain_users_role_proto_msgTypes[0]
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
	return file_publicit_domain_users_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_publicit_domain_users_role_proto protoreflect.FileDescriptor

var file_publicit_domain_users_role_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x4c, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x3e,
	0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x10, 0x02, 0x42, 0xfc,
	0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0xa2, 0x02, 0x04, 0x50, 0x50,
	0x44, 0x55, 0xaa, 0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0xca, 0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0xe2, 0x02, 0x28, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x3a, 0x3a,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicit_domain_users_role_proto_rawDescOnce sync.Once
	file_publicit_domain_users_role_proto_rawDescData = file_publicit_domain_users_role_proto_rawDesc
)

func file_publicit_domain_users_role_proto_rawDescGZIP() []byte {
	file_publicit_domain_users_role_proto_rawDescOnce.Do(func() {
		file_publicit_domain_users_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_domain_users_role_proto_rawDescData)
	})
	return file_publicit_domain_users_role_proto_rawDescData
}

var file_publicit_domain_users_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publicit_domain_users_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_publicit_domain_users_role_proto_goTypes = []interface{}{
	(RoleName)(0), // 0: protos.publicit.domain.users.RoleName
	(*Role)(nil),  // 1: protos.publicit.domain.users.Role
}
var file_publicit_domain_users_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publicit_domain_users_role_proto_init() }
func file_publicit_domain_users_role_proto_init() {
	if File_publicit_domain_users_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicit_domain_users_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publicit_domain_users_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_domain_users_role_proto_goTypes,
		DependencyIndexes: file_publicit_domain_users_role_proto_depIdxs,
		EnumInfos:         file_publicit_domain_users_role_proto_enumTypes,
		MessageInfos:      file_publicit_domain_users_role_proto_msgTypes,
	}.Build()
	File_publicit_domain_users_role_proto = out.File
	file_publicit_domain_users_role_proto_rawDesc = nil
	file_publicit_domain_users_role_proto_goTypes = nil
	file_publicit_domain_users_role_proto_depIdxs = nil
}
