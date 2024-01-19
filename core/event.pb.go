// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/event.proto

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

type EventType int32

const (
	EventType_ORDER EventType = 0
	EventType_TRADE EventType = 1
	EventType_LT    EventType = 2
	EventType_OB    EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "ORDER",
		1: "TRADE",
		2: "LT",
		3: "OB",
	}
	EventType_value = map[string]int32{
		"ORDER": 0,
		"TRADE": 1,
		"LT":    2,
		"OB":    3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_proto_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ETy   EventType   `protobuf:"varint,1,opt,name=eTy,proto3,enum=gateway_draft1.EventType" json:"eTy,omitempty"`
	S     string      `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	P     float64     `protobuf:"fixed64,3,opt,name=p,proto3" json:"p,omitempty"`
	V     float64     `protobuf:"fixed64,4,opt,name=v,proto3" json:"v,omitempty"`
	IsBuy bool        `protobuf:"varint,5,opt,name=isBuy,proto3" json:"isBuy,omitempty"`
	Ts    uint64      `protobuf:"varint,6,opt,name=ts,proto3" json:"ts,omitempty"`
	Id    string      `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	Ex    string      `protobuf:"bytes,8,opt,name=ex,proto3" json:"ex,omitempty"`
	InTy  string      `protobuf:"bytes,9,opt,name=inTy,proto3" json:"inTy,omitempty"`
	OTy   OrderType   `protobuf:"varint,10,opt,name=oTy,proto3,enum=gateway_draft1.OrderType" json:"oTy,omitempty"`
	OSt   OrderStatus `protobuf:"varint,11,opt,name=oSt,proto3,enum=gateway_draft1.OrderStatus" json:"oSt,omitempty"`
	OId   string      `protobuf:"bytes,12,opt,name=oId,proto3" json:"oId,omitempty"`
	Ap1   float64     `protobuf:"fixed64,13,opt,name=ap1,proto3" json:"ap1,omitempty"`
	Ap2   float64     `protobuf:"fixed64,14,opt,name=ap2,proto3" json:"ap2,omitempty"`
	Ap3   float64     `protobuf:"fixed64,15,opt,name=ap3,proto3" json:"ap3,omitempty"`
	Ap4   float64     `protobuf:"fixed64,16,opt,name=ap4,proto3" json:"ap4,omitempty"`
	Ap5   float64     `protobuf:"fixed64,17,opt,name=ap5,proto3" json:"ap5,omitempty"`
	Av1   float64     `protobuf:"fixed64,18,opt,name=av1,proto3" json:"av1,omitempty"`
	Av2   float64     `protobuf:"fixed64,19,opt,name=av2,proto3" json:"av2,omitempty"`
	Av3   float64     `protobuf:"fixed64,20,opt,name=av3,proto3" json:"av3,omitempty"`
	Av4   float64     `protobuf:"fixed64,21,opt,name=av4,proto3" json:"av4,omitempty"`
	Av5   float64     `protobuf:"fixed64,22,opt,name=av5,proto3" json:"av5,omitempty"`
	Bp1   float64     `protobuf:"fixed64,23,opt,name=bp1,proto3" json:"bp1,omitempty"`
	Bp2   float64     `protobuf:"fixed64,24,opt,name=bp2,proto3" json:"bp2,omitempty"`
	Bp3   float64     `protobuf:"fixed64,25,opt,name=bp3,proto3" json:"bp3,omitempty"`
	Bp4   float64     `protobuf:"fixed64,26,opt,name=bp4,proto3" json:"bp4,omitempty"`
	Bp5   float64     `protobuf:"fixed64,27,opt,name=bp5,proto3" json:"bp5,omitempty"`
	Bv1   float64     `protobuf:"fixed64,28,opt,name=bv1,proto3" json:"bv1,omitempty"`
	Bv2   float64     `protobuf:"fixed64,29,opt,name=bv2,proto3" json:"bv2,omitempty"`
	Bv3   float64     `protobuf:"fixed64,30,opt,name=bv3,proto3" json:"bv3,omitempty"`
	Bv4   float64     `protobuf:"fixed64,31,opt,name=bv4,proto3" json:"bv4,omitempty"`
	Bv5   float64     `protobuf:"fixed64,32,opt,name=bv5,proto3" json:"bv5,omitempty"`
	O     float64     `protobuf:"fixed64,33,opt,name=o,proto3" json:"o,omitempty"`
	H     float64     `protobuf:"fixed64,34,opt,name=h,proto3" json:"h,omitempty"`
	L     float64     `protobuf:"fixed64,35,opt,name=l,proto3" json:"l,omitempty"`
	C     float64     `protobuf:"fixed64,36,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetETy() EventType {
	if x != nil {
		return x.ETy
	}
	return EventType_ORDER
}

func (x *Event) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *Event) GetP() float64 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *Event) GetV() float64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *Event) GetIsBuy() bool {
	if x != nil {
		return x.IsBuy
	}
	return false
}

func (x *Event) GetTs() uint64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *Event) GetInTy() string {
	if x != nil {
		return x.InTy
	}
	return ""
}

func (x *Event) GetOTy() OrderType {
	if x != nil {
		return x.OTy
	}
	return OrderType_MARKET
}

func (x *Event) GetOSt() OrderStatus {
	if x != nil {
		return x.OSt
	}
	return OrderStatus_NEW
}

func (x *Event) GetOId() string {
	if x != nil {
		return x.OId
	}
	return ""
}

func (x *Event) GetAp1() float64 {
	if x != nil {
		return x.Ap1
	}
	return 0
}

func (x *Event) GetAp2() float64 {
	if x != nil {
		return x.Ap2
	}
	return 0
}

func (x *Event) GetAp3() float64 {
	if x != nil {
		return x.Ap3
	}
	return 0
}

func (x *Event) GetAp4() float64 {
	if x != nil {
		return x.Ap4
	}
	return 0
}

func (x *Event) GetAp5() float64 {
	if x != nil {
		return x.Ap5
	}
	return 0
}

func (x *Event) GetAv1() float64 {
	if x != nil {
		return x.Av1
	}
	return 0
}

func (x *Event) GetAv2() float64 {
	if x != nil {
		return x.Av2
	}
	return 0
}

func (x *Event) GetAv3() float64 {
	if x != nil {
		return x.Av3
	}
	return 0
}

func (x *Event) GetAv4() float64 {
	if x != nil {
		return x.Av4
	}
	return 0
}

func (x *Event) GetAv5() float64 {
	if x != nil {
		return x.Av5
	}
	return 0
}

func (x *Event) GetBp1() float64 {
	if x != nil {
		return x.Bp1
	}
	return 0
}

func (x *Event) GetBp2() float64 {
	if x != nil {
		return x.Bp2
	}
	return 0
}

func (x *Event) GetBp3() float64 {
	if x != nil {
		return x.Bp3
	}
	return 0
}

func (x *Event) GetBp4() float64 {
	if x != nil {
		return x.Bp4
	}
	return 0
}

func (x *Event) GetBp5() float64 {
	if x != nil {
		return x.Bp5
	}
	return 0
}

func (x *Event) GetBv1() float64 {
	if x != nil {
		return x.Bv1
	}
	return 0
}

func (x *Event) GetBv2() float64 {
	if x != nil {
		return x.Bv2
	}
	return 0
}

func (x *Event) GetBv3() float64 {
	if x != nil {
		return x.Bv3
	}
	return 0
}

func (x *Event) GetBv4() float64 {
	if x != nil {
		return x.Bv4
	}
	return 0
}

func (x *Event) GetBv5() float64 {
	if x != nil {
		return x.Bv5
	}
	return 0
}

func (x *Event) GetO() float64 {
	if x != nil {
		return x.O
	}
	return 0
}

func (x *Event) GetH() float64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *Event) GetL() float64 {
	if x != nil {
		return x.L
	}
	return 0
}

func (x *Event) GetC() float64 {
	if x != nil {
		return x.C
	}
	return 0
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61,
	0x66, 0x74, 0x31, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x05, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x03, 0x65, 0x54, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x65, 0x54, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x75, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x75, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x54, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x54,
	0x79, 0x12, 0x2b, 0x0a, 0x03, 0x6f, 0x54, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6f, 0x54, 0x79, 0x12, 0x2d,
	0x0a, 0x03, 0x6f, 0x53, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x6f, 0x53, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x31, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x70,
	0x31, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x61, 0x70, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x33, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x61, 0x70, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x34, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x61, 0x70, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x35, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x70, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x31,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x76, 0x32, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x76, 0x32, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x76, 0x33, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x76, 0x33, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x76, 0x34, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x76,
	0x34, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x35, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x61, 0x76, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x62, 0x70, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x32, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x62, 0x70, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x33, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x70, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x34,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x70, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x70, 0x35, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x70, 0x35, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x76, 0x31, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x76, 0x31, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x76, 0x32, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x76,
	0x32, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x76, 0x33, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x62, 0x76, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x76, 0x34, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x62, 0x76, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x76, 0x35, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x62, 0x76, 0x35, 0x12, 0x0c, 0x0a, 0x01, 0x6f, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x22, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x68, 0x12, 0x0c, 0x0a, 0x01, 0x6c, 0x18, 0x23, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x24, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x63, 0x2a,
	0x31, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x44, 0x45,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x42,
	0x10, 0x03, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x6e, 0x6d, 0x61, 0x68, 0x61, 0x72, 0x61, 0x6a, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_event_proto_goTypes = []interface{}{
	(EventType)(0),   // 0: gateway_draft1.EventType
	(*Event)(nil),    // 1: gateway_draft1.Event
	(OrderType)(0),   // 2: gateway_draft1.OrderType
	(OrderStatus)(0), // 3: gateway_draft1.OrderStatus
}
var file_proto_event_proto_depIdxs = []int32{
	0, // 0: gateway_draft1.Event.eTy:type_name -> gateway_draft1.EventType
	2, // 1: gateway_draft1.Event.oTy:type_name -> gateway_draft1.OrderType
	3, // 2: gateway_draft1.Event.oSt:type_name -> gateway_draft1.OrderStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	file_proto_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		EnumInfos:         file_proto_event_proto_enumTypes,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}
