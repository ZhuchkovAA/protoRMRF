// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/permission/permission.proto

package permission

import (
	summary2 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/authority/summary"
	summary3 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/district/summary"
	summary4 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/location/port/summary"
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission/action/summary"
	summary1 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission/object/summary"
	summary5 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/ship/summary"
	summary7 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/user/role/summary"
	summary6 "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/user/summary"
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

// Defines an access permission in the system.
// Permissions combine an Action and Object to create granular access controls (e.g., "read:inspection").
type Permission struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Id            string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                   // Unique identifier for the permission
	Action        *summary.ActionSummary       `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`           // The operation this permission allows (e.g., "read", "create")
	Object        *summary1.ObjectSummary      `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`           // The resource this permission applies to (e.g., "inspection", "user")
	Description   string                       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"` // Human-readable explanation of the permission
	Name          string                       `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`               // Display name (e.g., "Read Inspections")
	Code          string                       `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`               // Machine-readable identifier (e.g., "inspection:read")
	CreatedAt     *timestamppb.Timestamp       `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Authorities   []*summary2.AuthoritySummary `protobuf:"bytes,8,rep,name=authorities,proto3" json:"authorities,omitempty"`
	Districts     []*summary3.DistrictSummary  `protobuf:"bytes,9,rep,name=districts,proto3" json:"districts,omitempty"`
	Ports         []*summary4.PortSummary      `protobuf:"bytes,10,rep,name=ports,proto3" json:"ports,omitempty"`
	Ships         []*summary5.ShipSummary      `protobuf:"bytes,11,rep,name=ships,proto3" json:"ships,omitempty"`
	Users         []*summary6.UserSummary      `protobuf:"bytes,12,rep,name=users,proto3" json:"users,omitempty"`
	Roles         []*summary7.RoleSummary      `protobuf:"bytes,13,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_apfish_v1_permission_permission_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_permission_permission_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_apfish_v1_permission_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetAction() *summary.ActionSummary {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Permission) GetObject() *summary1.ObjectSummary {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Permission) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Permission) GetAuthorities() []*summary2.AuthoritySummary {
	if x != nil {
		return x.Authorities
	}
	return nil
}

func (x *Permission) GetDistricts() []*summary3.DistrictSummary {
	if x != nil {
		return x.Districts
	}
	return nil
}

func (x *Permission) GetPorts() []*summary4.PortSummary {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Permission) GetShips() []*summary5.ShipSummary {
	if x != nil {
		return x.Ships
	}
	return nil
}

func (x *Permission) GetUsers() []*summary6.UserSummary {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Permission) GetRoles() []*summary7.RoleSummary {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_apfish_v1_permission_permission_proto protoreflect.FileDescriptor

const file_apfish_v1_permission_permission_proto_rawDesc = "" +
	"\n" +
	"%apfish.v1/permission/permission.proto\x12\x14apfish.v1.permission\x1a\x1fgoogle/protobuf/timestamp.proto\x1a8apfish.v1/permission/action/summary/action_summary.proto\x1a8apfish.v1/permission/object/summary/object_summary.proto\x1a<apfish.v1/location/authority/summary/authority_summary.proto\x1a:apfish.v1/location/district/summary/district_summary.proto\x1a2apfish.v1/location/port/summary/port_summary.proto\x1a)apfish.v1/ship/summary/ship_summary.proto\x1a)apfish.v1/user/summary/user_summary.proto\x1a.apfish.v1/user/role/summary/role_summary.proto\"\xe1\x05\n" +
	"\n" +
	"Permission\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12J\n" +
	"\x06action\x18\x02 \x01(\v22.apfish.v1.permission.action.summary.ActionSummaryR\x06action\x12J\n" +
	"\x06object\x18\x03 \x01(\v22.apfish.v1.permission.object.summary.ObjectSummaryR\x06object\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12\x12\n" +
	"\x04code\x18\x06 \x01(\tR\x04code\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12X\n" +
	"\vauthorities\x18\b \x03(\v26.apfish.v1.location.authority.summary.AuthoritySummaryR\vauthorities\x12R\n" +
	"\tdistricts\x18\t \x03(\v24.apfish.v1.location.district.summary.DistrictSummaryR\tdistricts\x12B\n" +
	"\x05ports\x18\n" +
	" \x03(\v2,.apfish.v1.location.port.summary.PortSummaryR\x05ports\x129\n" +
	"\x05ships\x18\v \x03(\v2#.apfish.v1.ship.summary.ShipSummaryR\x05ships\x129\n" +
	"\x05users\x18\f \x03(\v2#.apfish.v1.user.summary.UserSummaryR\x05users\x12>\n" +
	"\x05roles\x18\r \x03(\v2(.apfish.v1.user.role.summary.RoleSummaryR\x05rolesBHZFgithub.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/permission;permissionb\x06proto3"

var (
	file_apfish_v1_permission_permission_proto_rawDescOnce sync.Once
	file_apfish_v1_permission_permission_proto_rawDescData []byte
)

func file_apfish_v1_permission_permission_proto_rawDescGZIP() []byte {
	file_apfish_v1_permission_permission_proto_rawDescOnce.Do(func() {
		file_apfish_v1_permission_permission_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_permission_permission_proto_rawDesc), len(file_apfish_v1_permission_permission_proto_rawDesc)))
	})
	return file_apfish_v1_permission_permission_proto_rawDescData
}

var file_apfish_v1_permission_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apfish_v1_permission_permission_proto_goTypes = []any{
	(*Permission)(nil),                // 0: apfish.v1.permission.Permission
	(*summary.ActionSummary)(nil),     // 1: apfish.v1.permission.action.summary.ActionSummary
	(*summary1.ObjectSummary)(nil),    // 2: apfish.v1.permission.object.summary.ObjectSummary
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
	(*summary2.AuthoritySummary)(nil), // 4: apfish.v1.location.authority.summary.AuthoritySummary
	(*summary3.DistrictSummary)(nil),  // 5: apfish.v1.location.district.summary.DistrictSummary
	(*summary4.PortSummary)(nil),      // 6: apfish.v1.location.port.summary.PortSummary
	(*summary5.ShipSummary)(nil),      // 7: apfish.v1.ship.summary.ShipSummary
	(*summary6.UserSummary)(nil),      // 8: apfish.v1.user.summary.UserSummary
	(*summary7.RoleSummary)(nil),      // 9: apfish.v1.user.role.summary.RoleSummary
}
var file_apfish_v1_permission_permission_proto_depIdxs = []int32{
	1, // 0: apfish.v1.permission.Permission.action:type_name -> apfish.v1.permission.action.summary.ActionSummary
	2, // 1: apfish.v1.permission.Permission.object:type_name -> apfish.v1.permission.object.summary.ObjectSummary
	3, // 2: apfish.v1.permission.Permission.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: apfish.v1.permission.Permission.authorities:type_name -> apfish.v1.location.authority.summary.AuthoritySummary
	5, // 4: apfish.v1.permission.Permission.districts:type_name -> apfish.v1.location.district.summary.DistrictSummary
	6, // 5: apfish.v1.permission.Permission.ports:type_name -> apfish.v1.location.port.summary.PortSummary
	7, // 6: apfish.v1.permission.Permission.ships:type_name -> apfish.v1.ship.summary.ShipSummary
	8, // 7: apfish.v1.permission.Permission.users:type_name -> apfish.v1.user.summary.UserSummary
	9, // 8: apfish.v1.permission.Permission.roles:type_name -> apfish.v1.user.role.summary.RoleSummary
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_apfish_v1_permission_permission_proto_init() }
func file_apfish_v1_permission_permission_proto_init() {
	if File_apfish_v1_permission_permission_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_permission_permission_proto_rawDesc), len(file_apfish_v1_permission_permission_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apfish_v1_permission_permission_proto_goTypes,
		DependencyIndexes: file_apfish_v1_permission_permission_proto_depIdxs,
		MessageInfos:      file_apfish_v1_permission_permission_proto_msgTypes,
	}.Build()
	File_apfish_v1_permission_permission_proto = out.File
	file_apfish_v1_permission_permission_proto_goTypes = nil
	file_apfish_v1_permission_permission_proto_depIdxs = nil
}
