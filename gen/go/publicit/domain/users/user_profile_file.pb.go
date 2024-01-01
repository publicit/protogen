// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/domain/users/user_profile_file.proto

package users

import (
	reflect "reflect"
	sync "sync"

	files "github.com/publicit/protogen/gen/go/publicit/domain/files"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileType int32

const (
	FileType_UNKNOWN_FILE_TYPE FileType = 0
	FileType_INE_ID_BACK       FileType = 1
	FileType_INE_ID_FRONT      FileType = 2
	FileType_CURP_ID           FileType = 3
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "UNKNOWN_FILE_TYPE",
		1: "INE_ID_BACK",
		2: "INE_ID_FRONT",
		3: "CURP_ID",
	}
	FileType_value = map[string]int32{
		"UNKNOWN_FILE_TYPE": 0,
		"INE_ID_BACK":       1,
		"INE_ID_FRONT":      2,
		"CURP_ID":           3,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_publicit_domain_users_user_profile_file_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_publicit_domain_users_user_profile_file_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_publicit_domain_users_user_profile_file_proto_rawDescGZIP(), []int{0}
}

// UserProfileFile represents a relationship between a user profile and a scanned file.
type UserProfileFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	File        *files.File  `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	UserProfile *UserProfile `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	FileType    FileType     `protobuf:"varint,4,opt,name=file_type,json=fileType,proto3,enum=protos.publicit.domain.users.FileType" json:"file_type,omitempty"`
	// MrzData is the resulting serialized JSON after processing a profile file.
	MrzData string `protobuf:"bytes,5,opt,name=mrz_data,json=mrzData,proto3" json:"mrz_data,omitempty"`
}

func (x *UserProfileFile) Reset() {
	*x = UserProfileFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicit_domain_users_user_profile_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileFile) ProtoMessage() {}

func (x *UserProfileFile) ProtoReflect() protoreflect.Message {
	mi := &file_publicit_domain_users_user_profile_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileFile.ProtoReflect.Descriptor instead.
func (*UserProfileFile) Descriptor() ([]byte, []int) {
	return file_publicit_domain_users_user_profile_file_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfileFile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfileFile) GetFile() *files.File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UserProfileFile) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

func (x *UserProfileFile) GetFileType() FileType {
	if x != nil {
		return x.FileType
	}
	return FileType_UNKNOWN_FILE_TYPE
}

func (x *UserProfileFile) GetMrzData() string {
	if x != nil {
		return x.MrzData
	}
	return ""
}

var File_publicit_domain_users_user_profile_file_proto protoreflect.FileDescriptor

var file_publicit_domain_users_user_profile_file_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x72, 0x7a,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x72, 0x7a,
	0x44, 0x61, 0x74, 0x61, 0x2a, 0x51, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x45, 0x5f, 0x49,
	0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x45, 0x5f,
	0x49, 0x44, 0x5f, 0x46, 0x52, 0x4f, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55,
	0x52, 0x50, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x42, 0x87, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x14, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0xa2,
	0x02, 0x04, 0x50, 0x50, 0x44, 0x55, 0xaa, 0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0xca, 0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x73, 0xe2, 0x02, 0x28, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicit_domain_users_user_profile_file_proto_rawDescOnce sync.Once
	file_publicit_domain_users_user_profile_file_proto_rawDescData = file_publicit_domain_users_user_profile_file_proto_rawDesc
)

func file_publicit_domain_users_user_profile_file_proto_rawDescGZIP() []byte {
	file_publicit_domain_users_user_profile_file_proto_rawDescOnce.Do(func() {
		file_publicit_domain_users_user_profile_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_domain_users_user_profile_file_proto_rawDescData)
	})
	return file_publicit_domain_users_user_profile_file_proto_rawDescData
}

var file_publicit_domain_users_user_profile_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publicit_domain_users_user_profile_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_publicit_domain_users_user_profile_file_proto_goTypes = []interface{}{
	(FileType)(0),           // 0: protos.publicit.domain.users.FileType
	(*UserProfileFile)(nil), // 1: protos.publicit.domain.users.UserProfileFile
	(*files.File)(nil),      // 2: protos.publicit.domain.files.File
	(*UserProfile)(nil),     // 3: protos.publicit.domain.users.UserProfile
}
var file_publicit_domain_users_user_profile_file_proto_depIdxs = []int32{
	2, // 0: protos.publicit.domain.users.UserProfileFile.file:type_name -> protos.publicit.domain.files.File
	3, // 1: protos.publicit.domain.users.UserProfileFile.user_profile:type_name -> protos.publicit.domain.users.UserProfile
	0, // 2: protos.publicit.domain.users.UserProfileFile.file_type:type_name -> protos.publicit.domain.users.FileType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_publicit_domain_users_user_profile_file_proto_init() }
func file_publicit_domain_users_user_profile_file_proto_init() {
	if File_publicit_domain_users_user_profile_file_proto != nil {
		return
	}
	file_publicit_domain_users_user_profile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_publicit_domain_users_user_profile_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfileFile); i {
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
			RawDescriptor: file_publicit_domain_users_user_profile_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_domain_users_user_profile_file_proto_goTypes,
		DependencyIndexes: file_publicit_domain_users_user_profile_file_proto_depIdxs,
		EnumInfos:         file_publicit_domain_users_user_profile_file_proto_enumTypes,
		MessageInfos:      file_publicit_domain_users_user_profile_file_proto_msgTypes,
	}.Build()
	File_publicit_domain_users_user_profile_file_proto = out.File
	file_publicit_domain_users_user_profile_file_proto_rawDesc = nil
	file_publicit_domain_users_user_profile_file_proto_goTypes = nil
	file_publicit_domain_users_user_profile_file_proto_depIdxs = nil
}
