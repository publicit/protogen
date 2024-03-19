// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: publicit/domain/geolocations/address.proto

package geolocations

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	City         string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country      string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	CountryShort string `protobuf:"bytes,4,opt,name=country_short,json=countryShort,proto3" json:"country_short,omitempty"`
	Full         string `protobuf:"bytes,5,opt,name=full,proto3" json:"full,omitempty"`
	State        string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Street       string `protobuf:"bytes,7,opt,name=street,proto3" json:"street,omitempty"`
	StreetNumber string `protobuf:"bytes,8,opt,name=street_number,json=streetNumber,proto3" json:"street_number,omitempty"`
	Zipcode      string `protobuf:"bytes,9,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicit_domain_geolocations_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_publicit_domain_geolocations_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_publicit_domain_geolocations_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCountryShort() string {
	if x != nil {
		return x.CountryShort
	}
	return ""
}

func (x *Address) GetFull() string {
	if x != nil {
		return x.Full
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetStreetNumber() string {
	if x != nil {
		return x.StreetNumber
	}
	return ""
}

func (x *Address) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

var File_publicit_domain_geolocations_address_proto protoreflect.FileDescriptor

var file_publicit_domain_geolocations_address_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x67, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xf7, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75,
	0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6c, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x42, 0xa9, 0x02, 0x0a, 0x27,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x6f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x04, 0x50, 0x50, 0x44, 0x47,
	0xaa, 0x02, 0x23, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x23, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c,
	0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x2f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5c, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x26, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x47, 0x65, 0x6f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicit_domain_geolocations_address_proto_rawDescOnce sync.Once
	file_publicit_domain_geolocations_address_proto_rawDescData = file_publicit_domain_geolocations_address_proto_rawDesc
)

func file_publicit_domain_geolocations_address_proto_rawDescGZIP() []byte {
	file_publicit_domain_geolocations_address_proto_rawDescOnce.Do(func() {
		file_publicit_domain_geolocations_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicit_domain_geolocations_address_proto_rawDescData)
	})
	return file_publicit_domain_geolocations_address_proto_rawDescData
}

var file_publicit_domain_geolocations_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_publicit_domain_geolocations_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: protos.publicit.domain.geolocations.Address
}
var file_publicit_domain_geolocations_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publicit_domain_geolocations_address_proto_init() }
func file_publicit_domain_geolocations_address_proto_init() {
	if File_publicit_domain_geolocations_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicit_domain_geolocations_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_publicit_domain_geolocations_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publicit_domain_geolocations_address_proto_goTypes,
		DependencyIndexes: file_publicit_domain_geolocations_address_proto_depIdxs,
		MessageInfos:      file_publicit_domain_geolocations_address_proto_msgTypes,
	}.Build()
	File_publicit_domain_geolocations_address_proto = out.File
	file_publicit_domain_geolocations_address_proto_rawDesc = nil
	file_publicit_domain_geolocations_address_proto_goTypes = nil
	file_publicit_domain_geolocations_address_proto_depIdxs = nil
}