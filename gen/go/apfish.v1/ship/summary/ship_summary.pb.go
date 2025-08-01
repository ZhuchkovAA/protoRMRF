// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/ship/summary/ship_summary.proto

package ship_summary

import (
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

type ShipSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Imo           int32                  `protobuf:"varint,2,opt,name=imo,proto3" json:"imo,omitempty"`
	TypeId        string                 `protobuf:"bytes,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	ClassId       string                 `protobuf:"bytes,4,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	AuthorityId   string                 `protobuf:"bytes,5,opt,name=authority_id,json=authorityId,proto3" json:"authority_id,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Mmsi          int32                  `protobuf:"varint,7,opt,name=mmsi,proto3" json:"mmsi,omitempty"`
	CallSignId    string                 `protobuf:"bytes,8,opt,name=call_sign_id,json=callSignId,proto3" json:"call_sign_id,omitempty"`
	CallSignValue string                 `protobuf:"bytes,9,opt,name=call_sign_value,json=callSignValue,proto3" json:"call_sign_value,omitempty"`
	DateBuild     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=date_build,json=dateBuild,proto3" json:"date_build,omitempty"`
	DeadWeight    int32                  `protobuf:"varint,11,opt,name=dead_weight,json=deadWeight,proto3" json:"dead_weight,omitempty"`
	Tonnage       int32                  `protobuf:"varint,12,opt,name=tonnage,proto3" json:"tonnage,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShipSummary) Reset() {
	*x = ShipSummary{}
	mi := &file_apfish_v1_ship_summary_ship_summary_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShipSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipSummary) ProtoMessage() {}

func (x *ShipSummary) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_ship_summary_ship_summary_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipSummary.ProtoReflect.Descriptor instead.
func (*ShipSummary) Descriptor() ([]byte, []int) {
	return file_apfish_v1_ship_summary_ship_summary_proto_rawDescGZIP(), []int{0}
}

func (x *ShipSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShipSummary) GetImo() int32 {
	if x != nil {
		return x.Imo
	}
	return 0
}

func (x *ShipSummary) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *ShipSummary) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *ShipSummary) GetAuthorityId() string {
	if x != nil {
		return x.AuthorityId
	}
	return ""
}

func (x *ShipSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipSummary) GetMmsi() int32 {
	if x != nil {
		return x.Mmsi
	}
	return 0
}

func (x *ShipSummary) GetCallSignId() string {
	if x != nil {
		return x.CallSignId
	}
	return ""
}

func (x *ShipSummary) GetCallSignValue() string {
	if x != nil {
		return x.CallSignValue
	}
	return ""
}

func (x *ShipSummary) GetDateBuild() *timestamppb.Timestamp {
	if x != nil {
		return x.DateBuild
	}
	return nil
}

func (x *ShipSummary) GetDeadWeight() int32 {
	if x != nil {
		return x.DeadWeight
	}
	return 0
}

func (x *ShipSummary) GetTonnage() int32 {
	if x != nil {
		return x.Tonnage
	}
	return 0
}

func (x *ShipSummary) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_apfish_v1_ship_summary_ship_summary_proto protoreflect.FileDescriptor

const file_apfish_v1_ship_summary_ship_summary_proto_rawDesc = "" +
	"\n" +
	")apfish.v1/ship/summary/ship_summary.proto\x12\x16apfish.v1.ship.summary\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa9\x03\n" +
	"\vShipSummary\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03imo\x18\x02 \x01(\x05R\x03imo\x12\x17\n" +
	"\atype_id\x18\x03 \x01(\tR\x06typeId\x12\x19\n" +
	"\bclass_id\x18\x04 \x01(\tR\aclassId\x12!\n" +
	"\fauthority_id\x18\x05 \x01(\tR\vauthorityId\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04name\x12\x12\n" +
	"\x04mmsi\x18\a \x01(\x05R\x04mmsi\x12 \n" +
	"\fcall_sign_id\x18\b \x01(\tR\n" +
	"callSignId\x12&\n" +
	"\x0fcall_sign_value\x18\t \x01(\tR\rcallSignValue\x129\n" +
	"\n" +
	"date_build\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tdateBuild\x12\x1f\n" +
	"\vdead_weight\x18\v \x01(\x05R\n" +
	"deadWeight\x12\x18\n" +
	"\atonnage\x18\f \x01(\x05R\atonnage\x129\n" +
	"\n" +
	"created_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAtBLZJgithub.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/ship/summary;ship_summaryb\x06proto3"

var (
	file_apfish_v1_ship_summary_ship_summary_proto_rawDescOnce sync.Once
	file_apfish_v1_ship_summary_ship_summary_proto_rawDescData []byte
)

func file_apfish_v1_ship_summary_ship_summary_proto_rawDescGZIP() []byte {
	file_apfish_v1_ship_summary_ship_summary_proto_rawDescOnce.Do(func() {
		file_apfish_v1_ship_summary_ship_summary_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_ship_summary_ship_summary_proto_rawDesc), len(file_apfish_v1_ship_summary_ship_summary_proto_rawDesc)))
	})
	return file_apfish_v1_ship_summary_ship_summary_proto_rawDescData
}

var file_apfish_v1_ship_summary_ship_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apfish_v1_ship_summary_ship_summary_proto_goTypes = []any{
	(*ShipSummary)(nil),           // 0: apfish.v1.ship.summary.ShipSummary
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_apfish_v1_ship_summary_ship_summary_proto_depIdxs = []int32{
	1, // 0: apfish.v1.ship.summary.ShipSummary.date_build:type_name -> google.protobuf.Timestamp
	1, // 1: apfish.v1.ship.summary.ShipSummary.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apfish_v1_ship_summary_ship_summary_proto_init() }
func file_apfish_v1_ship_summary_ship_summary_proto_init() {
	if File_apfish_v1_ship_summary_ship_summary_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_ship_summary_ship_summary_proto_rawDesc), len(file_apfish_v1_ship_summary_ship_summary_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apfish_v1_ship_summary_ship_summary_proto_goTypes,
		DependencyIndexes: file_apfish_v1_ship_summary_ship_summary_proto_depIdxs,
		MessageInfos:      file_apfish_v1_ship_summary_ship_summary_proto_msgTypes,
	}.Build()
	File_apfish_v1_ship_summary_ship_summary_proto = out.File
	file_apfish_v1_ship_summary_ship_summary_proto_goTypes = nil
	file_apfish_v1_ship_summary_ship_summary_proto_depIdxs = nil
}
