// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/messages/queues/queue.proto

package queues

import (
	reflect "reflect"
	sync "sync"

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

type QueueNames int32

const (
	QueueNames_QUEUE_UNKNOWN QueueNames = 0
	QueueNames_QUEUE_ISSUES  QueueNames = 1
)

// Enum value maps for QueueNames.
var (
	QueueNames_name = map[int32]string{
		0: "QUEUE_UNKNOWN",
		1: "QUEUE_ISSUES",
	}
	QueueNames_value = map[string]int32{
		"QUEUE_UNKNOWN": 0,
		"QUEUE_ISSUES":  1,
	}
)

func (x QueueNames) Enum() *QueueNames {
	p := new(QueueNames)
	*p = x
	return p
}

func (x QueueNames) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueueNames) Descriptor() protoreflect.EnumDescriptor {
	return file_publicit_messages_queues_queue_proto_enumTypes[0].Descriptor()
}

func (QueueNames) Type() protoreflect.EnumType {
	return &file_publicit_messages_queues_queue_proto_enumTypes[0]
}

func (x QueueNames) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueueNames.Descriptor instead.
func (QueueNames) EnumDescriptor() ([]byte, []int) {
	return file_publicit_messages_queues_queue_proto_rawDescGZIP(), []int{0}
}

var File_publicit_messages_queues_queue_proto protoreflect.FileDescriptor

var file_publicit_messages_queues_queue_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x31, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x51, 0x55, 0x45,
	0x55, 0x45, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x53, 0x10, 0x01, 0x42, 0x8f, 0x02, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0xa2,
	0x02, 0x04, 0x50, 0x50, 0x4d, 0x51, 0xaa, 0x02, 0x1f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0xca, 0x02, 0x1f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x5c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0xe2, 0x02, 0x2b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x3a, 0x3a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x3a, 0x3a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicit_messages_queues_queue_proto_rawDescOnce sync.Once
	file_publicit_messages_queues_queue_proto_rawDescData = file_publicit_messages_queues_queue_proto_rawDesc
)

func file_publicit_messages_queues_queue_proto_rawDescGZIP() []byte {
	file_publicit_messages_queues_queue_proto_rawDescOnce.Do(func() {
		file_publicit_messages_queues_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_messages_queues_queue_proto_rawDescData)
	})
	return file_publicit_messages_queues_queue_proto_rawDescData
}

var file_publicit_messages_queues_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publicit_messages_queues_queue_proto_goTypes = []interface{}{
	(QueueNames)(0), // 0: protos.publicit.messages.queues.QueueNames
}
var file_publicit_messages_queues_queue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publicit_messages_queues_queue_proto_init() }
func file_publicit_messages_queues_queue_proto_init() {
	if File_publicit_messages_queues_queue_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publicit_messages_queues_queue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_messages_queues_queue_proto_goTypes,
		DependencyIndexes: file_publicit_messages_queues_queue_proto_depIdxs,
		EnumInfos:         file_publicit_messages_queues_queue_proto_enumTypes,
	}.Build()
	File_publicit_messages_queues_queue_proto = out.File
	file_publicit_messages_queues_queue_proto_rawDesc = nil
	file_publicit_messages_queues_queue_proto_goTypes = nil
	file_publicit_messages_queues_queue_proto_depIdxs = nil
}
