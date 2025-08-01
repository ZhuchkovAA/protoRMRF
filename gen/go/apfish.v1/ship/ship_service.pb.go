// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: apfish.v1/ship/ship_service.proto

package ship

import (
	summary "github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/ship/summary"
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

type ShipRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShipRequest) Reset() {
	*x = ShipRequest{}
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipRequest) ProtoMessage() {}

func (x *ShipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipRequest.ProtoReflect.Descriptor instead.
func (*ShipRequest) Descriptor() ([]byte, []int) {
	return file_apfish_v1_ship_ship_service_proto_rawDescGZIP(), []int{0}
}

func (x *ShipRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ShipSummaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ship          *summary.ShipSummary   `protobuf:"bytes,1,opt,name=ship,proto3" json:"ship,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShipSummaryResponse) Reset() {
	*x = ShipSummaryResponse{}
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShipSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipSummaryResponse) ProtoMessage() {}

func (x *ShipSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipSummaryResponse.ProtoReflect.Descriptor instead.
func (*ShipSummaryResponse) Descriptor() ([]byte, []int) {
	return file_apfish_v1_ship_ship_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShipSummaryResponse) GetShip() *summary.ShipSummary {
	if x != nil {
		return x.Ship
	}
	return nil
}

type ListShipsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                      // Page number (1-based). Default: 1.
	PerPage       int32                  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"` // Items per page (default: 20, max: 100).
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListShipsRequest) Reset() {
	*x = ListShipsRequest{}
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShipsRequest) ProtoMessage() {}

func (x *ListShipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShipsRequest.ProtoReflect.Descriptor instead.
func (*ListShipsRequest) Descriptor() ([]byte, []int) {
	return file_apfish_v1_ship_ship_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListShipsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListShipsRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListShipsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ListShips     []*summary.ShipSummary `protobuf:"bytes,1,rep,name=list_ships,json=listShips,proto3" json:"list_ships,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListShipsResponse) Reset() {
	*x = ListShipsResponse{}
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShipsResponse) ProtoMessage() {}

func (x *ListShipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apfish_v1_ship_ship_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShipsResponse.ProtoReflect.Descriptor instead.
func (*ListShipsResponse) Descriptor() ([]byte, []int) {
	return file_apfish_v1_ship_ship_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListShipsResponse) GetListShips() []*summary.ShipSummary {
	if x != nil {
		return x.ListShips
	}
	return nil
}

func (x *ListShipsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_apfish_v1_ship_ship_service_proto protoreflect.FileDescriptor

const file_apfish_v1_ship_ship_service_proto_rawDesc = "" +
	"\n" +
	"!apfish.v1/ship/ship_service.proto\x12\x0eapfish.v1.ship\x1a\x19apfish.v1/ship/ship.proto\x1a)apfish.v1/ship/summary/ship_summary.proto\"\x1d\n" +
	"\vShipRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"N\n" +
	"\x13ShipSummaryResponse\x127\n" +
	"\x04ship\x18\x01 \x01(\v2#.apfish.v1.ship.summary.ShipSummaryR\x04ship\"A\n" +
	"\x10ListShipsRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x19\n" +
	"\bper_page\x18\x02 \x01(\x05R\aperPage\"m\n" +
	"\x11ListShipsResponse\x12B\n" +
	"\n" +
	"list_ships\x18\x01 \x03(\v2#.apfish.v1.ship.summary.ShipSummaryR\tlistShips\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total2\xf1\x01\n" +
	"\vShipService\x12<\n" +
	"\aGetShip\x12\x1b.apfish.v1.ship.ShipRequest\x1a\x14.apfish.v1.ship.Ship\x12R\n" +
	"\x0eGetShipSummary\x12\x1b.apfish.v1.ship.ShipRequest\x1a#.apfish.v1.ship.ShipSummaryResponse\x12P\n" +
	"\tListShips\x12 .apfish.v1.ship.ListShipsRequest\x1a!.apfish.v1.ship.ListShipsResponseB<Z:github.com/ZhuchkovAA/protoRMRF/gen/go/apfish.v1/ship;shipb\x06proto3"

var (
	file_apfish_v1_ship_ship_service_proto_rawDescOnce sync.Once
	file_apfish_v1_ship_ship_service_proto_rawDescData []byte
)

func file_apfish_v1_ship_ship_service_proto_rawDescGZIP() []byte {
	file_apfish_v1_ship_ship_service_proto_rawDescOnce.Do(func() {
		file_apfish_v1_ship_ship_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apfish_v1_ship_ship_service_proto_rawDesc), len(file_apfish_v1_ship_ship_service_proto_rawDesc)))
	})
	return file_apfish_v1_ship_ship_service_proto_rawDescData
}

var file_apfish_v1_ship_ship_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apfish_v1_ship_ship_service_proto_goTypes = []any{
	(*ShipRequest)(nil),         // 0: apfish.v1.ship.ShipRequest
	(*ShipSummaryResponse)(nil), // 1: apfish.v1.ship.ShipSummaryResponse
	(*ListShipsRequest)(nil),    // 2: apfish.v1.ship.ListShipsRequest
	(*ListShipsResponse)(nil),   // 3: apfish.v1.ship.ListShipsResponse
	(*summary.ShipSummary)(nil), // 4: apfish.v1.ship.summary.ShipSummary
	(*Ship)(nil),                // 5: apfish.v1.ship.Ship
}
var file_apfish_v1_ship_ship_service_proto_depIdxs = []int32{
	4, // 0: apfish.v1.ship.ShipSummaryResponse.ship:type_name -> apfish.v1.ship.summary.ShipSummary
	4, // 1: apfish.v1.ship.ListShipsResponse.list_ships:type_name -> apfish.v1.ship.summary.ShipSummary
	0, // 2: apfish.v1.ship.ShipService.GetShip:input_type -> apfish.v1.ship.ShipRequest
	0, // 3: apfish.v1.ship.ShipService.GetShipSummary:input_type -> apfish.v1.ship.ShipRequest
	2, // 4: apfish.v1.ship.ShipService.ListShips:input_type -> apfish.v1.ship.ListShipsRequest
	5, // 5: apfish.v1.ship.ShipService.GetShip:output_type -> apfish.v1.ship.Ship
	1, // 6: apfish.v1.ship.ShipService.GetShipSummary:output_type -> apfish.v1.ship.ShipSummaryResponse
	3, // 7: apfish.v1.ship.ShipService.ListShips:output_type -> apfish.v1.ship.ListShipsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apfish_v1_ship_ship_service_proto_init() }
func file_apfish_v1_ship_ship_service_proto_init() {
	if File_apfish_v1_ship_ship_service_proto != nil {
		return
	}
	file_apfish_v1_ship_ship_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apfish_v1_ship_ship_service_proto_rawDesc), len(file_apfish_v1_ship_ship_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apfish_v1_ship_ship_service_proto_goTypes,
		DependencyIndexes: file_apfish_v1_ship_ship_service_proto_depIdxs,
		MessageInfos:      file_apfish_v1_ship_ship_service_proto_msgTypes,
	}.Build()
	File_apfish_v1_ship_ship_service_proto = out.File
	file_apfish_v1_ship_ship_service_proto_goTypes = nil
	file_apfish_v1_ship_ship_service_proto_depIdxs = nil
}
