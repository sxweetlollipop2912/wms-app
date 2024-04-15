// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: transaction_manager/api/api.proto

package api

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Action          int32                `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	Sku             string               `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	ShelfName       string               `protobuf:"bytes,4,opt,name=shelf_name,json=shelfName,proto3" json:"shelf_name,omitempty"`
	QuantityOnShelf int64                `protobuf:"varint,5,opt,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"`
	AuthorName      string               `protobuf:"bytes,6,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Transaction) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *Transaction) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Transaction) GetShelfName() string {
	if x != nil {
		return x.ShelfName
	}
	return ""
}

func (x *Transaction) GetQuantityOnShelf() int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return 0
}

func (x *Transaction) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action          int32  `protobuf:"varint,1,opt,name=action,proto3" json:"action,omitempty"`
	Sku             string `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	ShelfName       string `protobuf:"bytes,3,opt,name=shelf_name,json=shelfName,proto3" json:"shelf_name,omitempty"`
	QuantityOnShelf int64  `protobuf:"varint,4,opt,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"`
	AuthorName      string `protobuf:"bytes,5,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRequest) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *InsertRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *InsertRequest) GetShelfName() string {
	if x != nil {
		return x.ShelfName
	}
	return ""
}

func (x *InsertRequest) GetQuantityOnShelf() int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return 0
}

func (x *InsertRequest) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

type GetDuringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetDuringRequest) Reset() {
	*x = GetDuringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDuringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDuringRequest) ProtoMessage() {}

func (x *GetDuringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDuringRequest.ProtoReflect.Descriptor instead.
func (*GetDuringRequest) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetDuringRequest) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *GetDuringRequest) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type GetDuringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetDuringResponse) Reset() {
	*x = GetDuringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDuringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDuringResponse) ProtoMessage() {}

func (x *GetDuringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDuringResponse.ProtoReflect.Descriptor instead.
func (*GetDuringResponse) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetDuringResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type GetByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorName string `protobuf:"bytes,1,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
}

func (x *GetByUserRequest) Reset() {
	*x = GetByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserRequest) ProtoMessage() {}

func (x *GetByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserRequest.ProtoReflect.Descriptor instead.
func (*GetByUserRequest) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetByUserRequest) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

type GetByUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetByUserResponse) Reset() {
	*x = GetByUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserResponse) ProtoMessage() {}

func (x *GetByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserResponse.ProtoReflect.Descriptor instead.
func (*GetByUserResponse) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetByUserResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_manager_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_manager_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_transaction_manager_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *AddUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_transaction_manager_api_api_proto protoreflect.FileDescriptor

var file_transaction_manager_api_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x28, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0xe2, 0x03, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x06, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x37, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x86, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x3a, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_manager_api_api_proto_rawDescOnce sync.Once
	file_transaction_manager_api_api_proto_rawDescData = file_transaction_manager_api_api_proto_rawDesc
)

func file_transaction_manager_api_api_proto_rawDescGZIP() []byte {
	file_transaction_manager_api_api_proto_rawDescOnce.Do(func() {
		file_transaction_manager_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_manager_api_api_proto_rawDescData)
	})
	return file_transaction_manager_api_api_proto_rawDescData
}

var file_transaction_manager_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_transaction_manager_api_api_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: simple_warehouse.transaction_manager.api.Transaction
	(*InsertRequest)(nil),       // 1: simple_warehouse.transaction_manager.api.InsertRequest
	(*GetDuringRequest)(nil),    // 2: simple_warehouse.transaction_manager.api.GetDuringRequest
	(*GetDuringResponse)(nil),   // 3: simple_warehouse.transaction_manager.api.GetDuringResponse
	(*GetByUserRequest)(nil),    // 4: simple_warehouse.transaction_manager.api.GetByUserRequest
	(*GetByUserResponse)(nil),   // 5: simple_warehouse.transaction_manager.api.GetByUserResponse
	(*AddUserRequest)(nil),      // 6: simple_warehouse.transaction_manager.api.AddUserRequest
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_transaction_manager_api_api_proto_depIdxs = []int32{
	7, // 0: simple_warehouse.transaction_manager.api.Transaction.date:type_name -> google.protobuf.Timestamp
	7, // 1: simple_warehouse.transaction_manager.api.GetDuringRequest.start_date:type_name -> google.protobuf.Timestamp
	7, // 2: simple_warehouse.transaction_manager.api.GetDuringRequest.end_date:type_name -> google.protobuf.Timestamp
	0, // 3: simple_warehouse.transaction_manager.api.GetDuringResponse.transactions:type_name -> simple_warehouse.transaction_manager.api.Transaction
	0, // 4: simple_warehouse.transaction_manager.api.GetByUserResponse.transactions:type_name -> simple_warehouse.transaction_manager.api.Transaction
	1, // 5: simple_warehouse.transaction_manager.api.TransactionManager.Insert:input_type -> simple_warehouse.transaction_manager.api.InsertRequest
	2, // 6: simple_warehouse.transaction_manager.api.TransactionManager.GetDuring:input_type -> simple_warehouse.transaction_manager.api.GetDuringRequest
	4, // 7: simple_warehouse.transaction_manager.api.TransactionManager.GetByUser:input_type -> simple_warehouse.transaction_manager.api.GetByUserRequest
	6, // 8: simple_warehouse.transaction_manager.api.TransactionManager.AddUser:input_type -> simple_warehouse.transaction_manager.api.AddUserRequest
	8, // 9: simple_warehouse.transaction_manager.api.TransactionManager.Insert:output_type -> google.protobuf.Empty
	3, // 10: simple_warehouse.transaction_manager.api.TransactionManager.GetDuring:output_type -> simple_warehouse.transaction_manager.api.GetDuringResponse
	5, // 11: simple_warehouse.transaction_manager.api.TransactionManager.GetByUser:output_type -> simple_warehouse.transaction_manager.api.GetByUserResponse
	8, // 12: simple_warehouse.transaction_manager.api.TransactionManager.AddUser:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transaction_manager_api_api_proto_init() }
func file_transaction_manager_api_api_proto_init() {
	if File_transaction_manager_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_manager_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_transaction_manager_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_transaction_manager_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDuringRequest); i {
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
		file_transaction_manager_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDuringResponse); i {
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
		file_transaction_manager_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserRequest); i {
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
		file_transaction_manager_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserResponse); i {
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
		file_transaction_manager_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
			RawDescriptor: file_transaction_manager_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_manager_api_api_proto_goTypes,
		DependencyIndexes: file_transaction_manager_api_api_proto_depIdxs,
		MessageInfos:      file_transaction_manager_api_api_proto_msgTypes,
	}.Build()
	File_transaction_manager_api_api_proto = out.File
	file_transaction_manager_api_api_proto_rawDesc = nil
	file_transaction_manager_api_api_proto_goTypes = nil
	file_transaction_manager_api_api_proto_depIdxs = nil
}
