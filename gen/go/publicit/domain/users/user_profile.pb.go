// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/domain/users/user_profile.proto

package users

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/publicit/protogen/gen/go/publicit/domain/files"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserGender int32

const (
	UserGender_UNKNOWN_GENDER UserGender = 0
	UserGender_MALE           UserGender = 1
	UserGender_FEMALE         UserGender = 2
	UserGender_OTHER          UserGender = 3
)

// Enum value maps for UserGender.
var (
	UserGender_name = map[int32]string{
		0: "UNKNOWN_GENDER",
		1: "MALE",
		2: "FEMALE",
		3: "OTHER",
	}
	UserGender_value = map[string]int32{
		"UNKNOWN_GENDER": 0,
		"MALE":           1,
		"FEMALE":         2,
		"OTHER":          3,
	}
)

func (x UserGender) Enum() *UserGender {
	p := new(UserGender)
	*p = x
	return p
}

func (x UserGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGender) Descriptor() protoreflect.EnumDescriptor {
	return file_publicit_domain_users_user_profile_proto_enumTypes[0].Descriptor()
}

func (UserGender) Type() protoreflect.EnumType {
	return &file_publicit_domain_users_user_profile_proto_enumTypes[0]
}

func (x UserGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGender.Descriptor instead.
func (UserGender) EnumDescriptor() ([]byte, []int) {
	return file_publicit_domain_users_user_profile_proto_rawDescGZIP(), []int{0}
}

// UserProfile stores the user information we need to create market segments of people.
// It works also as a validation mechanism for identity match.
type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName   string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Dob         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=dob,proto3" json:"dob,omitempty"`
	PhoneNumber string                 `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Gender      UserGender             `protobuf:"varint,8,opt,name=gender,proto3,enum=protos.publicit.domain.users.UserGender" json:"gender,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsCompleted bool                   `protobuf:"varint,11,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicit_domain_users_user_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_publicit_domain_users_user_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_publicit_domain_users_user_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProfile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserProfile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserProfile) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *UserProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserProfile) GetGender() UserGender {
	if x != nil {
		return x.Gender
	}
	return UserGender_UNKNOWN_GENDER
}

func (x *UserProfile) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserProfile) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserProfile) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

var File_publicit_domain_users_user_profile_proto protoreflect.FileDescriptor

var file_publicit_domain_users_user_profile_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4a, 0x04, 0x08, 0x07,
	0x10, 0x08, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x2a, 0x41, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x03, 0x42, 0x83, 0x02, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f,
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
	file_publicit_domain_users_user_profile_proto_rawDescOnce sync.Once
	file_publicit_domain_users_user_profile_proto_rawDescData = file_publicit_domain_users_user_profile_proto_rawDesc
)

func file_publicit_domain_users_user_profile_proto_rawDescGZIP() []byte {
	file_publicit_domain_users_user_profile_proto_rawDescOnce.Do(func() {
		file_publicit_domain_users_user_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_domain_users_user_profile_proto_rawDescData)
	})
	return file_publicit_domain_users_user_profile_proto_rawDescData
}

var file_publicit_domain_users_user_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publicit_domain_users_user_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_publicit_domain_users_user_profile_proto_goTypes = []interface{}{
	(UserGender)(0),               // 0: protos.publicit.domain.users.UserGender
	(*UserProfile)(nil),           // 1: protos.publicit.domain.users.UserProfile
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_publicit_domain_users_user_profile_proto_depIdxs = []int32{
	2, // 0: protos.publicit.domain.users.UserProfile.dob:type_name -> google.protobuf.Timestamp
	0, // 1: protos.publicit.domain.users.UserProfile.gender:type_name -> protos.publicit.domain.users.UserGender
	2, // 2: protos.publicit.domain.users.UserProfile.created_at:type_name -> google.protobuf.Timestamp
	2, // 3: protos.publicit.domain.users.UserProfile.updated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_publicit_domain_users_user_profile_proto_init() }
func file_publicit_domain_users_user_profile_proto_init() {
	if File_publicit_domain_users_user_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicit_domain_users_user_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
			RawDescriptor: file_publicit_domain_users_user_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_domain_users_user_profile_proto_goTypes,
		DependencyIndexes: file_publicit_domain_users_user_profile_proto_depIdxs,
		EnumInfos:         file_publicit_domain_users_user_profile_proto_enumTypes,
		MessageInfos:      file_publicit_domain_users_user_profile_proto_msgTypes,
	}.Build()
	File_publicit_domain_users_user_profile_proto = out.File
	file_publicit_domain_users_user_profile_proto_rawDesc = nil
	file_publicit_domain_users_user_profile_proto_goTypes = nil
	file_publicit_domain_users_user_profile_proto_depIdxs = nil
}
