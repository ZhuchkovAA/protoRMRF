// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/user/role/role_service.proto

package user_role

import (
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/user/role/summary"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleRequest) Reset() {
	*x = RoleRequest{}
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRequest) ProtoMessage() {}

func (x *RoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRequest.ProtoReflect.Descriptor instead.
func (*RoleRequest) Descriptor() ([]byte, []int) {
	return file_apfish_v1_user_role_role_service_proto_rawDescGZIP(), []int{0}
}

func (x *RoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleSummaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *summary.RoleSummary   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleSummaryResponse) Reset() {
	*x = RoleSummaryResponse{}
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSummaryResponse) ProtoMessage() {}

func (x *RoleSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSummaryResponse.ProtoReflect.Descriptor instead.
func (*RoleSummaryResponse) Descriptor() ([]byte, []int) {
	return file_apfish_v1_user_role_role_service_proto_rawDescGZIP(), []int{1}
}

func (x *RoleSummaryResponse) GetRole() *summary.RoleSummary {
	if x != nil {
		return x.Role
	}
	return nil
}

type ListRolesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                      // Page number (1-based). Default: 1.
	PerPage       int32                  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"` // Items per page (default: 20, max: 100).
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRolesRequest) Reset() {
	*x = ListRolesRequest{}
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesRequest) ProtoMessage() {}

func (x *ListRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesRequest.ProtoReflect.Descriptor instead.
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return file_apfish_v1_user_role_role_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListRolesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRolesRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ListRoles     []*summary.RoleSummary `protobuf:"bytes,1,rep,name=list_roles,json=listRoles,proto3" json:"list_roles,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_user_role_role_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesResponse.ProtoReflect.Descriptor instead.
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return file_apfish_v1_user_role_role_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListRolesResponse) GetListRoles() []*summary.RoleSummary {
	if x != nil {
		return x.ListRoles
	}
	return nil
}

func (x *ListRolesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_apfish_v1_user_role_role_service_proto protoreflect.FileDescriptor

const file_apfish_v1_user_role_role_service_proto_rawDesc = "" +
	"\n" +
	"&apfish.v1/user/role/role_service.proto\x12\x13apfish.v1.user.role\x1a\x1eapfish.v1/user/role/role.proto\x1a.apfish.v1/user/role/summary/role_summary.proto\"\x1d\n" +
	"\vRoleRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"S\n" +
	"\x13RoleSummaryResponse\x12<\n" +
	"\x04role\x18\x01 \x01(\v2(.apfish.v1.user.role.summary.RoleSummaryR\x04role\"A\n" +
	"\x10ListRolesRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x19\n" +
	"\bper_page\x18\x02 \x01(\x05R\aperPage\"r\n" +
	"\x11ListRolesResponse\x12G\n" +
	"\n" +
	"list_roles\x18\x01 \x03(\v2(.apfish.v1.user.role.summary.RoleSummaryR\tlistRoles\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total2\x8f\x02\n" +
	"\vRoleService\x12F\n" +
	"\aGetRole\x12 .apfish.v1.user.role.RoleRequest\x1a\x19.apfish.v1.user.role.Role\x12\\\n" +
	"\x0eGetRoleSummary\x12 .apfish.v1.user.role.RoleRequest\x1a(.apfish.v1.user.role.RoleSummaryResponse\x12Z\n" +
	"\tListRoles\x12%.apfish.v1.user.role.ListRolesRequest\x1a&.apfish.v1.user.role.ListRolesResponseBFZDgithub.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/user/role;user_roleb\x06proto3"

var (
	file_apfish_v1_user_role_role_service_proto_rawDescOnce sync.Once
	file_apfish_v1_user_role_role_service_proto_rawDescData []byte
)

func file_apfish_v1_user_role_role_service_proto_rawDescGZIP() []byte {
	file_apfish_v1_user_role_role_service_proto_rawDescOnce.Do(func() {
		file_apfish_v1_user_role_role_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_user_role_role_service_proto_rawDesc), len(file_apfish_v1_user_role_role_service_proto_rawDesc)))
	})
	return file_apfish_v1_user_role_role_service_proto_rawDescData
}

var file_apfish_v1_user_role_role_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apfish_v1_user_role_role_service_proto_goTypes = []any{
	(*RoleRequest)(nil),         // 0: apfish.v1.user.role.RoleRequest
	(*RoleSummaryResponse)(nil), // 1: apfish.v1.user.role.RoleSummaryResponse
	(*ListRolesRequest)(nil),    // 2: apfish.v1.user.role.ListRolesRequest
	(*ListRolesResponse)(nil),   // 3: apfish.v1.user.role.ListRolesResponse
	(*summary.RoleSummary)(nil), // 4: apfish.v1.user.role.summary.RoleSummary
	(*Role)(nil),                // 5: apfish.v1.user.role.Role
}
var file_apfish_v1_user_role_role_service_proto_depIdxs = []int32{
	4, // 0: apfish.v1.user.role.RoleSummaryResponse.role:type_name -> apfish.v1.user.role.summary.RoleSummary
	4, // 1: apfish.v1.user.role.ListRolesResponse.list_roles:type_name -> apfish.v1.user.role.summary.RoleSummary
	0, // 2: apfish.v1.user.role.RoleService.GetRole:input_type -> apfish.v1.user.role.RoleRequest
	0, // 3: apfish.v1.user.role.RoleService.GetRoleSummary:input_type -> apfish.v1.user.role.RoleRequest
	2, // 4: apfish.v1.user.role.RoleService.ListRoles:input_type -> apfish.v1.user.role.ListRolesRequest
	5, // 5: apfish.v1.user.role.RoleService.GetRole:output_type -> apfish.v1.user.role.Role
	1, // 6: apfish.v1.user.role.RoleService.GetRoleSummary:output_type -> apfish.v1.user.role.RoleSummaryResponse
	3, // 7: apfish.v1.user.role.RoleService.ListRoles:output_type -> apfish.v1.user.role.ListRolesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apfish_v1_user_role_role_service_proto_init() }
func file_apfish_v1_user_role_role_service_proto_init() {
	if File_apfish_v1_user_role_role_service_proto != nil {
		return
	}
	file_apfish_v1_user_role_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_user_role_role_service_proto_rawDesc), len(file_apfish_v1_user_role_role_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apfish_v1_user_role_role_service_proto_goTypes,
		DependencyIndexes: file_apfish_v1_user_role_role_service_proto_depIdxs,
		MessageInfos:      file_apfish_v1_user_role_role_service_proto_msgTypes,
	}.Build()
	File_apfish_v1_user_role_role_service_proto = out.File
	file_apfish_v1_user_role_role_service_proto_goTypes = nil
	file_apfish_v1_user_role_role_service_proto_depIdxs = nil
}
