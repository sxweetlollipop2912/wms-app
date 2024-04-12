// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: product_manager/api/api.proto

package api

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

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku             string   `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ExpiredDate     int64    `protobuf:"varint,3,opt,name=expired_date,json=expiredDate,proto3" json:"expired_date,omitempty"`
	Category        string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	ShelfNames      []string `protobuf:"bytes,5,rep,name=shelf_names,json=shelfNames,proto3" json:"shelf_names,omitempty"`
	QuantityOnShelf []int64  `protobuf:"varint,6,rep,packed,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *ImportRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ImportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportRequest) GetExpiredDate() int64 {
	if x != nil {
		return x.ExpiredDate
	}
	return 0
}

func (x *ImportRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ImportRequest) GetShelfNames() []string {
	if x != nil {
		return x.ShelfNames
	}
	return nil
}

func (x *ImportRequest) GetQuantityOnShelf() []int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return nil
}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *ImportResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ImportResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku             string   `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	ShelfNames      []string `protobuf:"bytes,2,rep,name=shelf_names,json=shelfNames,proto3" json:"shelf_names,omitempty"`
	QuantityOnShelf []int64  `protobuf:"varint,3,rep,packed,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"` // < 0 to export all
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ExportRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ExportRequest) GetShelfNames() []string {
	if x != nil {
		return x.ShelfNames
	}
	return nil
}

func (x *ExportRequest) GetQuantityOnShelf() []int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return nil
}

type ExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok              bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Reason          string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Sku             string   `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name            string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShelfNames      []string `protobuf:"bytes,5,rep,name=shelf_names,json=shelfNames,proto3" json:"shelf_names,omitempty"`
	QuantityOnShelf []int64  `protobuf:"varint,6,rep,packed,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"`
	ExpiredDate     int64    `protobuf:"varint,7,opt,name=expired_date,json=expiredDate,proto3" json:"expired_date,omitempty"`
	Category        string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *ExportResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ExportResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ExportResponse) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ExportResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExportResponse) GetShelfNames() []string {
	if x != nil {
		return x.ShelfNames
	}
	return nil
}

func (x *ExportResponse) GetQuantityOnShelf() []int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return nil
}

func (x *ExportResponse) GetExpiredDate() int64 {
	if x != nil {
		return x.ExpiredDate
	}
	return 0
}

func (x *ExportResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type GetProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok              bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Reason          string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Sku             string   `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name            string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShelfNames      []string `protobuf:"bytes,5,rep,name=shelf_names,json=shelfNames,proto3" json:"shelf_names,omitempty"`
	QuantityOnShelf []int64  `protobuf:"varint,6,rep,packed,name=quantity_on_shelf,json=quantityOnShelf,proto3" json:"quantity_on_shelf,omitempty"`
	ExpiredDate     int64    `protobuf:"varint,7,opt,name=expired_date,json=expiredDate,proto3" json:"expired_date,omitempty"`
	Category        string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_product_manager_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetProductResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetProductResponse) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *GetProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductResponse) GetShelfNames() []string {
	if x != nil {
		return x.ShelfNames
	}
	return nil
}

func (x *GetProductResponse) GetQuantityOnShelf() []int64 {
	if x != nil {
		return x.QuantityOnShelf
	}
	return nil
}

func (x *GetProductResponse) GetExpiredDate() int64 {
	if x != nil {
		return x.ExpiredDate
	}
	return 0
}

func (x *GetProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_product_manager_api_api_proto protoreflect.FileDescriptor

var file_product_manager_api_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x38, 0x0a, 0x0e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x22, 0xea, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0xee, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x82, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x06, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x33, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x75, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x33, 0x2e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x37, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a,
	0x24, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_manager_api_api_proto_rawDescOnce sync.Once
	file_product_manager_api_api_proto_rawDescData = file_product_manager_api_api_proto_rawDesc
)

func file_product_manager_api_api_proto_rawDescGZIP() []byte {
	file_product_manager_api_api_proto_rawDescOnce.Do(func() {
		file_product_manager_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_manager_api_api_proto_rawDescData)
	})
	return file_product_manager_api_api_proto_rawDescData
}

var file_product_manager_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_manager_api_api_proto_goTypes = []interface{}{
	(*ImportRequest)(nil),      // 0: simple_warehouse.project_manager.api.ImportRequest
	(*ImportResponse)(nil),     // 1: simple_warehouse.project_manager.api.ImportResponse
	(*ExportRequest)(nil),      // 2: simple_warehouse.project_manager.api.ExportRequest
	(*ExportResponse)(nil),     // 3: simple_warehouse.project_manager.api.ExportResponse
	(*GetProductRequest)(nil),  // 4: simple_warehouse.project_manager.api.GetProductRequest
	(*GetProductResponse)(nil), // 5: simple_warehouse.project_manager.api.GetProductResponse
}
var file_product_manager_api_api_proto_depIdxs = []int32{
	0, // 0: simple_warehouse.project_manager.api.ProductManager.Import:input_type -> simple_warehouse.project_manager.api.ImportRequest
	2, // 1: simple_warehouse.project_manager.api.ProductManager.Export:input_type -> simple_warehouse.project_manager.api.ExportRequest
	4, // 2: simple_warehouse.project_manager.api.ProductManager.GetProduct:input_type -> simple_warehouse.project_manager.api.GetProductRequest
	1, // 3: simple_warehouse.project_manager.api.ProductManager.Import:output_type -> simple_warehouse.project_manager.api.ImportResponse
	3, // 4: simple_warehouse.project_manager.api.ProductManager.Export:output_type -> simple_warehouse.project_manager.api.ExportResponse
	5, // 5: simple_warehouse.project_manager.api.ProductManager.GetProduct:output_type -> simple_warehouse.project_manager.api.GetProductResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_manager_api_api_proto_init() }
func file_product_manager_api_api_proto_init() {
	if File_product_manager_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_manager_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportRequest); i {
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
		file_product_manager_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
		file_product_manager_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_product_manager_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportResponse); i {
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
		file_product_manager_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_product_manager_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductResponse); i {
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
			RawDescriptor: file_product_manager_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_manager_api_api_proto_goTypes,
		DependencyIndexes: file_product_manager_api_api_proto_depIdxs,
		MessageInfos:      file_product_manager_api_api_proto_msgTypes,
	}.Build()
	File_product_manager_api_api_proto = out.File
	file_product_manager_api_api_proto_rawDesc = nil
	file_product_manager_api_api_proto_goTypes = nil
	file_product_manager_api_api_proto_depIdxs = nil
}
