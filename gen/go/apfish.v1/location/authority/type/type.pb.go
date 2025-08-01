// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/location/authority/type/type.proto

package location_authority_type

import (
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/authority/summary"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of contact method (e.g., "Email", "Phone").
type Type struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Id            string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // Unique type ID.
	Name          string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Human-readable name (e.g., "Email").
	NameRus       string                      `protobuf:"bytes,3,opt,name=name_rus,json=nameRus,proto3" json:"name_rus,omitempty"`
	CreatedAt     *timestamppb.Timestamp      `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // When the type was defined.
	Authorities   []*summary.AuthoritySummary `protobuf:"bytes,5,rep,name=authorities,proto3" json:"authorities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Type) Reset() {
	*x = Type{}
	mi := &file_apfish_v1_location_authority_type_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_location_authority_type_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_apfish_v1_location_authority_type_type_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetNameRus() string {
	if x != nil {
		return x.NameRus
	}
	return ""
}

func (x *Type) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Type) GetAuthorities() []*summary.AuthoritySummary {
	if x != nil {
		return x.Authorities
	}
	return nil
}

var File_apfish_v1_location_authority_type_type_proto protoreflect.FileDescriptor

const file_apfish_v1_location_authority_type_type_proto_rawDesc = "" +
	"\n" +
	",apfish.v1/location/authority/type/type.proto\x12!apfish.v1.location.authority.type\x1a\x1fgoogle/protobuf/timestamp.proto\x1a<apfish.v1/location/authority/summary/authority_summary.proto\"\xda\x01\n" +
	"\x04Type\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x19\n" +
	"\bname_rus\x18\x03 \x01(\tR\anameRus\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12X\n" +
	"\vauthorities\x18\x05 \x03(\v26.apfish.v1.location.authority.summary.AuthoritySummaryR\vauthoritiesBbZ`github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/authority/type;location_authority_typeb\x06proto3"

var (
	file_apfish_v1_location_authority_type_type_proto_rawDescOnce sync.Once
	file_apfish_v1_location_authority_type_type_proto_rawDescData []byte
)

func file_apfish_v1_location_authority_type_type_proto_rawDescGZIP() []byte {
	file_apfish_v1_location_authority_type_type_proto_rawDescOnce.Do(func() {
		file_apfish_v1_location_authority_type_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_location_authority_type_type_proto_rawDesc), len(file_apfish_v1_location_authority_type_type_proto_rawDesc)))
	})
	return file_apfish_v1_location_authority_type_type_proto_rawDescData
}

var file_apfish_v1_location_authority_type_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apfish_v1_location_authority_type_type_proto_goTypes = []any{
	(*Type)(nil),                     // 0: apfish.v1.location.authority.type.Type
	(*timestamppb.Timestamp)(nil),    // 1: google.protobuf.Timestamp
	(*summary.AuthoritySummary)(nil), // 2: apfish.v1.location.authority.summary.AuthoritySummary
}
var file_apfish_v1_location_authority_type_type_proto_depIdxs = []int32{
	1, // 0: apfish.v1.location.authority.type.Type.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: apfish.v1.location.authority.type.Type.authorities:type_name -> apfish.v1.location.authority.summary.AuthoritySummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apfish_v1_location_authority_type_type_proto_init() }
func file_apfish_v1_location_authority_type_type_proto_init() {
	if File_apfish_v1_location_authority_type_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_location_authority_type_type_proto_rawDesc), len(file_apfish_v1_location_authority_type_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apfish_v1_location_authority_type_type_proto_goTypes,
		DependencyIndexes: file_apfish_v1_location_authority_type_type_proto_depIdxs,
		MessageInfos:      file_apfish_v1_location_authority_type_type_proto_msgTypes,
	}.Build()
	File_apfish_v1_location_authority_type_type_proto = out.File
	file_apfish_v1_location_authority_type_type_proto_goTypes = nil
	file_apfish_v1_location_authority_type_type_proto_depIdxs = nil
}
