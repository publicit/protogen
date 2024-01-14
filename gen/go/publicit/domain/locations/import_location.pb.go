// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/domain/locations/import_location.proto

package locations

import (
	reflect "reflect"
	sync "sync"

	files "github.com/publicit/protogen/gen/go/publicit/domain/files"
	users "github.com/publicit/protogen/gen/go/publicit/domain/users"
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

type ImportStatus int32

const (
	ImportStatus_UNKNOWN_IMPORT_STATUS ImportStatus = 0
	ImportStatus_SCHEDULED             ImportStatus = 1
	ImportStatus_PROCESSING            ImportStatus = 2
	ImportStatus_SUCCESS               ImportStatus = 3
	ImportStatus_FAILURE               ImportStatus = 4
)

// Enum value maps for ImportStatus.
var (
	ImportStatus_name = map[int32]string{
		0: "UNKNOWN_IMPORT_STATUS",
		1: "SCHEDULED",
		2: "PROCESSING",
		3: "SUCCESS",
		4: "FAILURE",
	}
	ImportStatus_value = map[string]int32{
		"UNKNOWN_IMPORT_STATUS": 0,
		"SCHEDULED":             1,
		"PROCESSING":            2,
		"SUCCESS":               3,
		"FAILURE":               4,
	}
)

func (x ImportStatus) Enum() *ImportStatus {
	p := new(ImportStatus)
	*p = x
	return p
}

func (x ImportStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_publicit_domain_locations_import_location_proto_enumTypes[0].Descriptor()
}

func (ImportStatus) Type() protoreflect.EnumType {
	return &file_publicit_domain_locations_import_location_proto_enumTypes[0]
}

func (x ImportStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportStatus.Descriptor instead.
func (ImportStatus) EnumDescriptor() ([]byte, []int) {
	return file_publicit_domain_locations_import_location_proto_rawDescGZIP(), []int{0}
}

type ImportLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User    *users.User            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Status  ImportStatus           `protobuf:"varint,4,opt,name=status,proto3,enum=protos.publicit.domain.locations.ImportStatus" json:"status,omitempty"`
	File    *files.File            `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	Error   string                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ImportLocation) Reset() {
	*x = ImportLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicit_domain_locations_import_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportLocation) ProtoMessage() {}

func (x *ImportLocation) ProtoReflect() protoreflect.Message {
	mi := &file_publicit_domain_locations_import_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportLocation.ProtoReflect.Descriptor instead.
func (*ImportLocation) Descriptor() ([]byte, []int) {
	return file_publicit_domain_locations_import_location_proto_rawDescGZIP(), []int{0}
}

func (x *ImportLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImportLocation) GetUser() *users.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ImportLocation) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ImportLocation) GetStatus() ImportStatus {
	if x != nil {
		return x.Status
	}
	return ImportStatus_UNKNOWN_IMPORT_STATUS
}

func (x *ImportLocation) GetFile() *files.File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *ImportLocation) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_publicit_domain_locations_import_location_proto protoreflect.FileDescriptor

var file_publicit_domain_locations_import_location_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x0e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x36, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a,
	0x62, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x19, 0x0a, 0x15, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52,
	0x45, 0x10, 0x04, 0x42, 0x9e, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x13, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0xa2, 0x02, 0x04, 0x50, 0x50, 0x44, 0x4c, 0xaa, 0x02, 0x20, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x20, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2,
	0x02, 0x2c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x23, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicit_domain_locations_import_location_proto_rawDescOnce sync.Once
	file_publicit_domain_locations_import_location_proto_rawDescData = file_publicit_domain_locations_import_location_proto_rawDesc
)

func file_publicit_domain_locations_import_location_proto_rawDescGZIP() []byte {
	file_publicit_domain_locations_import_location_proto_rawDescOnce.Do(func() {
		file_publicit_domain_locations_import_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_domain_locations_import_location_proto_rawDescData)
	})
	return file_publicit_domain_locations_import_location_proto_rawDescData
}

var file_publicit_domain_locations_import_location_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publicit_domain_locations_import_location_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_publicit_domain_locations_import_location_proto_goTypes = []interface{}{
	(ImportStatus)(0),             // 0: protos.publicit.domain.locations.ImportStatus
	(*ImportLocation)(nil),        // 1: protos.publicit.domain.locations.ImportLocation
	(*users.User)(nil),            // 2: protos.publicit.domain.users.User
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*files.File)(nil),            // 4: protos.publicit.domain.files.File
}
var file_publicit_domain_locations_import_location_proto_depIdxs = []int32{
	2, // 0: protos.publicit.domain.locations.ImportLocation.user:type_name -> protos.publicit.domain.users.User
	3, // 1: protos.publicit.domain.locations.ImportLocation.created:type_name -> google.protobuf.Timestamp
	0, // 2: protos.publicit.domain.locations.ImportLocation.status:type_name -> protos.publicit.domain.locations.ImportStatus
	4, // 3: protos.publicit.domain.locations.ImportLocation.file:type_name -> protos.publicit.domain.files.File
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_publicit_domain_locations_import_location_proto_init() }
func file_publicit_domain_locations_import_location_proto_init() {
	if File_publicit_domain_locations_import_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicit_domain_locations_import_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportLocation); i {
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
			RawDescriptor: file_publicit_domain_locations_import_location_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_domain_locations_import_location_proto_goTypes,
		DependencyIndexes: file_publicit_domain_locations_import_location_proto_depIdxs,
		EnumInfos:         file_publicit_domain_locations_import_location_proto_enumTypes,
		MessageInfos:      file_publicit_domain_locations_import_location_proto_msgTypes,
	}.Build()
	File_publicit_domain_locations_import_location_proto = out.File
	file_publicit_domain_locations_import_location_proto_rawDesc = nil
	file_publicit_domain_locations_import_location_proto_goTypes = nil
	file_publicit_domain_locations_import_location_proto_depIdxs = nil
}
