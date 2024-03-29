// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/order.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderType int32

const (
	OrderType_MARKET            OrderType = 0
	OrderType_LIMIT             OrderType = 1
	OrderType_STOP              OrderType = 2
	OrderType_STOP_LIMIT        OrderType = 3
	OrderType_CANCEL            OrderType = 4
	OrderType_REPLACE           OrderType = 5
	OrderType_REJECT            OrderType = 6
	OrderType_CANCEL_REPLACE    OrderType = 7
	OrderType_CANCEL_ALL        OrderType = 8
	OrderType_TAKE_PROFIT       OrderType = 9
	OrderType_TAKE_PROFIT_LIMIT OrderType = 10
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0:  "MARKET",
		1:  "LIMIT",
		2:  "STOP",
		3:  "STOP_LIMIT",
		4:  "CANCEL",
		5:  "REPLACE",
		6:  "REJECT",
		7:  "CANCEL_REPLACE",
		8:  "CANCEL_ALL",
		9:  "TAKE_PROFIT",
		10: "TAKE_PROFIT_LIMIT",
	}
	OrderType_value = map[string]int32{
		"MARKET":            0,
		"LIMIT":             1,
		"STOP":              2,
		"STOP_LIMIT":        3,
		"CANCEL":            4,
		"REPLACE":           5,
		"REJECT":            6,
		"CANCEL_REPLACE":    7,
		"CANCEL_ALL":        8,
		"TAKE_PROFIT":       9,
		"TAKE_PROFIT_LIMIT": 10,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

type OrderStatus int32

const (
	OrderStatus_NEW                    OrderStatus = 0
	OrderStatus_PARTIALLY_FILLED       OrderStatus = 1
	OrderStatus_FILLED                 OrderStatus = 2
	OrderStatus_CANCELED               OrderStatus = 3
	OrderStatus_PENDING_CANCEL         OrderStatus = 4
	OrderStatus_REJECTED               OrderStatus = 5
	OrderStatus_EXPIRED                OrderStatus = 6
	OrderStatus_PENDING_REPLACE        OrderStatus = 7
	OrderStatus_REPLACED               OrderStatus = 8
	OrderStatus_PENDING_NEW            OrderStatus = 9
	OrderStatus_ACCEPTED_FOR_BIDDING   OrderStatus = 10
	OrderStatus_PENDING_CANCEL_REPLACE OrderStatus = 11
	OrderStatus_UNKNOWN                OrderStatus = 12
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0:  "NEW",
		1:  "PARTIALLY_FILLED",
		2:  "FILLED",
		3:  "CANCELED",
		4:  "PENDING_CANCEL",
		5:  "REJECTED",
		6:  "EXPIRED",
		7:  "PENDING_REPLACE",
		8:  "REPLACED",
		9:  "PENDING_NEW",
		10: "ACCEPTED_FOR_BIDDING",
		11: "PENDING_CANCEL_REPLACE",
		12: "UNKNOWN",
	}
	OrderStatus_value = map[string]int32{
		"NEW":                    0,
		"PARTIALLY_FILLED":       1,
		"FILLED":                 2,
		"CANCELED":               3,
		"PENDING_CANCEL":         4,
		"REJECTED":               5,
		"EXPIRED":                6,
		"PENDING_REPLACE":        7,
		"REPLACED":               8,
		"PENDING_NEW":            9,
		"ACCEPTED_FOR_BIDDING":   10,
		"PENDING_CANCEL_REPLACE": 11,
		"UNKNOWN":                12,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61,
	0x66, 0x74, 0x31, 0x2a, 0xad, 0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x41, 0x4b, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x54, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11,
	0x54, 0x41, 0x4b, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x54, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x0a, 0x2a, 0xec, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x09, 0x12,
	0x18, 0x0a, 0x14, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x4c,
	0x41, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x0c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x6e, 0x6d, 0x61, 0x68, 0x61, 0x72, 0x61, 0x6a, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_order_proto_goTypes = []interface{}{
	(OrderType)(0),   // 0: gateway_draft1.OrderType
	(OrderStatus)(0), // 1: gateway_draft1.OrderStatus
}
var file_proto_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		EnumInfos:         file_proto_order_proto_enumTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
