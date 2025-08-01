// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/permission/object/object.proto

package permission_object

import (
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission/summary"
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

// Defines a type of resource or object that can be accessed in the system.
type Object struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Id            string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                   // Unique identifier for the object type
	Name          string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // Human-readable name (e.g., "Inspection")
	Description   string                       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // Detailed explanation of the object type
	Code          string                       `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`               // Machine-friendly code (e.g., "inspection")
	CreatedAt     *timestamppb.Timestamp       `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Permissions   []*summary.PermissionSummary `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Object) Reset() {
	*x = Object{}
	mi := &file_apfish_v1_permission_object_object_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_permission_object_object_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_apfish_v1_permission_object_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Object) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Object) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Object) GetPermissions() []*summary.PermissionSummary {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_apfish_v1_permission_object_object_proto protoreflect.FileDescriptor

const file_apfish_v1_permission_object_object_proto_rawDesc = "" +
	"\n" +
	"(apfish.v1/permission/object/object.proto\x12\x1bapfish.v1.permission.object\x1a\x1fgoogle/protobuf/timestamp.proto\x1a5apfish.v1/permission/summary/permission_summary.proto\"\xf0\x01\n" +
	"\x06Object\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x12\n" +
	"\x04code\x18\x04 \x01(\tR\x04code\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12Q\n" +
	"\vpermissions\x18\x06 \x03(\v2/.apfish.v1.permission.summary.PermissionSummaryR\vpermissionsBVZTgithub.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission/object;permission_objectb\x06proto3"

var (
	file_apfish_v1_permission_object_object_proto_rawDescOnce sync.Once
	file_apfish_v1_permission_object_object_proto_rawDescData []byte
)

func file_apfish_v1_permission_object_object_proto_rawDescGZIP() []byte {
	file_apfish_v1_permission_object_object_proto_rawDescOnce.Do(func() {
		file_apfish_v1_permission_object_object_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_permission_object_object_proto_rawDesc), len(file_apfish_v1_permission_object_object_proto_rawDesc)))
	})
	return file_apfish_v1_permission_object_object_proto_rawDescData
}

var file_apfish_v1_permission_object_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apfish_v1_permission_object_object_proto_goTypes = []any{
	(*Object)(nil),                    // 0: apfish.v1.permission.object.Object
	(*timestamppb.Timestamp)(nil),     // 1: google.protobuf.Timestamp
	(*summary.PermissionSummary)(nil), // 2: apfish.v1.permission.summary.PermissionSummary
}
var file_apfish_v1_permission_object_object_proto_depIdxs = []int32{
	1, // 0: apfish.v1.permission.object.Object.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: apfish.v1.permission.object.Object.permissions:type_name -> apfish.v1.permission.summary.PermissionSummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apfish_v1_permission_object_object_proto_init() }
func file_apfish_v1_permission_object_object_proto_init() {
	if File_apfish_v1_permission_object_object_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_permission_object_object_proto_rawDesc), len(file_apfish_v1_permission_object_object_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apfish_v1_permission_object_object_proto_goTypes,
		DependencyIndexes: file_apfish_v1_permission_object_object_proto_depIdxs,
		MessageInfos:      file_apfish_v1_permission_object_object_proto_msgTypes,
	}.Build()
	File_apfish_v1_permission_object_object_proto = out.File
	file_apfish_v1_permission_object_object_proto_goTypes = nil
	file_apfish_v1_permission_object_object_proto_depIdxs = nil
}
