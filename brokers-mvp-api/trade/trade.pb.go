// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.1
// source: trade/trade.proto

package trade

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BuyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentName string `protobuf:"bytes,1,opt,name=instrument_name,json=instrumentName,proto3" json:"instrument_name,omitempty"`
	OrderType      string `protobuf:"bytes,2,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Quantity       int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TransactionId  string `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *BuyRequest) Reset() {
	*x = BuyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyRequest) ProtoMessage() {}

func (x *BuyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyRequest.ProtoReflect.Descriptor instead.
func (*BuyRequest) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{0}
}

func (x *BuyRequest) GetInstrumentName() string {
	if x != nil {
		return x.InstrumentName
	}
	return ""
}

func (x *BuyRequest) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *BuyRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *BuyRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type TradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseType string `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
}

func (x *TradeResponse) Reset() {
	*x = TradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeResponse) ProtoMessage() {}

func (x *TradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeResponse.ProtoReflect.Descriptor instead.
func (*TradeResponse) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{1}
}

func (x *TradeResponse) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

type SellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentName string `protobuf:"bytes,1,opt,name=instrument_name,json=instrumentName,proto3" json:"instrument_name,omitempty"`
	OrderType      string `protobuf:"bytes,2,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Quantity       int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TransactionId  string `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *SellRequest) Reset() {
	*x = SellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellRequest) ProtoMessage() {}

func (x *SellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellRequest.ProtoReflect.Descriptor instead.
func (*SellRequest) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{2}
}

func (x *SellRequest) GetInstrumentName() string {
	if x != nil {
		return x.InstrumentName
	}
	return ""
}

func (x *SellRequest) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *SellRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SellRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{3}
}

func (x *CancelRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseType string `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{5}
}

func (x *Orders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentName string `protobuf:"bytes,1,opt,name=instrument_name,json=instrumentName,proto3" json:"instrument_name,omitempty"`
	OrderType      string `protobuf:"bytes,2,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetInstrumentName() string {
	if x != nil {
		return x.InstrumentName
	}
	return ""
}

func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

type InstrumentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instruments []*Instrument `protobuf:"bytes,1,rep,name=instruments,proto3" json:"instruments,omitempty"`
}

func (x *InstrumentsResponse) Reset() {
	*x = InstrumentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstrumentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentsResponse) ProtoMessage() {}

func (x *InstrumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentsResponse.ProtoReflect.Descriptor instead.
func (*InstrumentsResponse) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{7}
}

func (x *InstrumentsResponse) GetInstruments() []*Instrument {
	if x != nil {
		return x.Instruments
	}
	return nil
}

type Instrument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol     string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	MarketName string `protobuf:"bytes,2,opt,name=market_name,json=marketName,proto3" json:"market_name,omitempty"`
}

func (x *Instrument) Reset() {
	*x = Instrument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_trade_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instrument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instrument) ProtoMessage() {}

func (x *Instrument) ProtoReflect() protoreflect.Message {
	mi := &file_trade_trade_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instrument.ProtoReflect.Descriptor instead.
func (*Instrument) Descriptor() ([]byte, []int) {
	return file_trade_trade_proto_rawDescGZIP(), []int{8}
}

func (x *Instrument) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Instrument) GetMarketName() string {
	if x != nil {
		return x.MarketName
	}
	return ""
}

var File_trade_trade_proto protoreflect.FileDescriptor

var file_trade_trade_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x34, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x13,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xa9, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x42, 0x75, 0x79, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x64, 0x6f, 0x63,
	0x75, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x2d, 0x6d, 0x76, 0x70, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x3b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_trade_proto_rawDescOnce sync.Once
	file_trade_trade_proto_rawDescData = file_trade_trade_proto_rawDesc
)

func file_trade_trade_proto_rawDescGZIP() []byte {
	file_trade_trade_proto_rawDescOnce.Do(func() {
		file_trade_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_trade_proto_rawDescData)
	})
	return file_trade_trade_proto_rawDescData
}

var file_trade_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_trade_trade_proto_goTypes = []interface{}{
	(*BuyRequest)(nil),          // 0: trade.BuyRequest
	(*TradeResponse)(nil),       // 1: trade.TradeResponse
	(*SellRequest)(nil),         // 2: trade.SellRequest
	(*CancelRequest)(nil),       // 3: trade.CancelRequest
	(*Response)(nil),            // 4: trade.Response
	(*Orders)(nil),              // 5: trade.Orders
	(*Order)(nil),               // 6: trade.Order
	(*InstrumentsResponse)(nil), // 7: trade.InstrumentsResponse
	(*Instrument)(nil),          // 8: trade.Instrument
	(*empty.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_trade_trade_proto_depIdxs = []int32{
	6, // 0: trade.Orders.orders:type_name -> trade.Order
	8, // 1: trade.InstrumentsResponse.instruments:type_name -> trade.Instrument
	0, // 2: trade.TradingService.Buy:input_type -> trade.BuyRequest
	2, // 3: trade.TradingService.Sell:input_type -> trade.SellRequest
	3, // 4: trade.TradingService.Cancel:input_type -> trade.CancelRequest
	9, // 5: trade.TradingService.ListOrders:input_type -> google.protobuf.Empty
	9, // 6: trade.TradingService.ListInstruments:input_type -> google.protobuf.Empty
	1, // 7: trade.TradingService.Buy:output_type -> trade.TradeResponse
	1, // 8: trade.TradingService.Sell:output_type -> trade.TradeResponse
	4, // 9: trade.TradingService.Cancel:output_type -> trade.Response
	5, // 10: trade.TradingService.ListOrders:output_type -> trade.Orders
	7, // 11: trade.TradingService.ListInstruments:output_type -> trade.InstrumentsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trade_trade_proto_init() }
func file_trade_trade_proto_init() {
	if File_trade_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyRequest); i {
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
		file_trade_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeResponse); i {
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
		file_trade_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellRequest); i {
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
		file_trade_trade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_trade_trade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_trade_trade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
		file_trade_trade_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_trade_trade_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstrumentsResponse); i {
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
		file_trade_trade_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instrument); i {
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
			RawDescriptor: file_trade_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trade_trade_proto_goTypes,
		DependencyIndexes: file_trade_trade_proto_depIdxs,
		MessageInfos:      file_trade_trade_proto_msgTypes,
	}.Build()
	File_trade_trade_proto = out.File
	file_trade_trade_proto_rawDesc = nil
	file_trade_trade_proto_goTypes = nil
	file_trade_trade_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TradingServiceClient is the client API for TradingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TradingServiceClient interface {
	Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*TradeResponse, error)
	Sell(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*TradeResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*Response, error)
	ListOrders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Orders, error)
	ListInstruments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InstrumentsResponse, error)
}

type tradingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingServiceClient(cc grpc.ClientConnInterface) TradingServiceClient {
	return &tradingServiceClient{cc}
}

func (c *tradingServiceClient) Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*TradeResponse, error) {
	out := new(TradeResponse)
	err := c.cc.Invoke(ctx, "/trade.TradingService/Buy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) Sell(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*TradeResponse, error) {
	out := new(TradeResponse)
	err := c.cc.Invoke(ctx, "/trade.TradingService/Sell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trade.TradingService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) ListOrders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/trade.TradingService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) ListInstruments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InstrumentsResponse, error) {
	out := new(InstrumentsResponse)
	err := c.cc.Invoke(ctx, "/trade.TradingService/ListInstruments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServiceServer is the server API for TradingService service.
type TradingServiceServer interface {
	Buy(context.Context, *BuyRequest) (*TradeResponse, error)
	Sell(context.Context, *SellRequest) (*TradeResponse, error)
	Cancel(context.Context, *CancelRequest) (*Response, error)
	ListOrders(context.Context, *empty.Empty) (*Orders, error)
	ListInstruments(context.Context, *empty.Empty) (*InstrumentsResponse, error)
}

// UnimplementedTradingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTradingServiceServer struct {
}

func (*UnimplementedTradingServiceServer) Buy(context.Context, *BuyRequest) (*TradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buy not implemented")
}
func (*UnimplementedTradingServiceServer) Sell(context.Context, *SellRequest) (*TradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (*UnimplementedTradingServiceServer) Cancel(context.Context, *CancelRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedTradingServiceServer) ListOrders(context.Context, *empty.Empty) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (*UnimplementedTradingServiceServer) ListInstruments(context.Context, *empty.Empty) (*InstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstruments not implemented")
}

func RegisterTradingServiceServer(s *grpc.Server, srv TradingServiceServer) {
	s.RegisterService(&_TradingService_serviceDesc, srv)
}

func _TradingService_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.TradingService/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).Buy(ctx, req.(*BuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.TradingService/Sell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).Sell(ctx, req.(*SellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.TradingService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.TradingService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).ListOrders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_ListInstruments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).ListInstruments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.TradingService/ListInstruments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).ListInstruments(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trade.TradingService",
	HandlerType: (*TradingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Buy",
			Handler:    _TradingService_Buy_Handler,
		},
		{
			MethodName: "Sell",
			Handler:    _TradingService_Sell_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _TradingService_Cancel_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _TradingService_ListOrders_Handler,
		},
		{
			MethodName: "ListInstruments",
			Handler:    _TradingService_ListInstruments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trade/trade.proto",
}
