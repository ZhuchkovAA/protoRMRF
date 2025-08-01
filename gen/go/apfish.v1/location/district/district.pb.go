// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/location/district/district.proto

package location_district

import (
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/port/summary"
	summary1 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission/summary"
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

type District struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Id            string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt     *timestamppb.Timestamp        `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Ports         []*summary.PortSummary        `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Permissions   []*summary1.PermissionSummary `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *District) Reset() {
	*x = District{}
	mi := &file_apfish_v1_location_district_district_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *District) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*District) ProtoMessage() {}

func (x *District) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_location_district_district_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use District.ProtoReflect.Descriptor instead.
func (*District) Descriptor() ([]byte, []int) {
	return file_apfish_v1_location_district_district_proto_rawDescGZIP(), []int{0}
}

func (x *District) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *District) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *District) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *District) GetPorts() []*summary.PortSummary {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *District) GetPermissions() []*summary1.PermissionSummary {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_apfish_v1_location_district_district_proto protoreflect.FileDescriptor

const file_apfish_v1_location_district_district_proto_rawDesc = "" +
	"\n" +
	"*apfish.v1/location/district/district.proto\x12\x1bapfish.v1.location.district\x1a\x1fgoogle/protobuf/timestamp.proto\x1a2apfish.v1/location/port/summary/port_summary.proto\x1a5apfish.v1/permission/summary/permission_summary.proto\"\x80\x02\n" +
	"\bDistrict\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12B\n" +
	"\x05ports\x18\x04 \x03(\v2,.apfish.v1.location.port.summary.PortSummaryR\x05ports\x12Q\n" +
	"\vpermissions\x18\x05 \x03(\v2/.apfish.v1.permission.summary.PermissionSummaryR\vpermissionsBVZTgithub.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/district;location_districtb\x06proto3"

var (
	file_apfish_v1_location_district_district_proto_rawDescOnce sync.Once
	file_apfish_v1_location_district_district_proto_rawDescData []byte
)

func file_apfish_v1_location_district_district_proto_rawDescGZIP() []byte {
	file_apfish_v1_location_district_district_proto_rawDescOnce.Do(func() {
		file_apfish_v1_location_district_district_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_location_district_district_proto_rawDesc), len(file_apfish_v1_location_district_district_proto_rawDesc)))
	})
	return file_apfish_v1_location_district_district_proto_rawDescData
}

var file_apfish_v1_location_district_district_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apfish_v1_location_district_district_proto_goTypes = []any{
	(*District)(nil),                   // 0: apfish.v1.location.district.District
	(*timestamppb.Timestamp)(nil),      // 1: google.protobuf.Timestamp
	(*summary.PortSummary)(nil),        // 2: apfish.v1.location.port.summary.PortSummary
	(*summary1.PermissionSummary)(nil), // 3: apfish.v1.permission.summary.PermissionSummary
}
var file_apfish_v1_location_district_district_proto_depIdxs = []int32{
	1, // 0: apfish.v1.location.district.District.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: apfish.v1.location.district.District.ports:type_name -> apfish.v1.location.port.summary.PortSummary
	3, // 2: apfish.v1.location.district.District.permissions:type_name -> apfish.v1.permission.summary.PermissionSummary
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apfish_v1_location_district_district_proto_init() }
func file_apfish_v1_location_district_district_proto_init() {
	if File_apfish_v1_location_district_district_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_location_district_district_proto_rawDesc), len(file_apfish_v1_location_district_district_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apfish_v1_location_district_district_proto_goTypes,
		DependencyIndexes: file_apfish_v1_location_district_district_proto_depIdxs,
		MessageInfos:      file_apfish_v1_location_district_district_proto_msgTypes,
	}.Build()
	File_apfish_v1_location_district_district_proto = out.File
	file_apfish_v1_location_district_district_proto_goTypes = nil
	file_apfish_v1_location_district_district_proto_depIdxs = nil
}
