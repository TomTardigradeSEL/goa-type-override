// Code generated with goa v3.11.3, DO NOT EDIT.
//
// Pet Service protocol buffer definition
//
// Command:
// $ goa gen petsvc/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v4.22.0
// source: goadesign_goagen_pet_service.proto

package pet_servicepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pets "petsvc/interfaces/pets"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendDogsStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dog *pets.Dog `protobuf:"bytes,1,opt,name=dog,proto3" json:"dog,omitempty"`
}

func (x *SendDogsStreamingRequest) Reset() {
	*x = SendDogsStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDogsStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDogsStreamingRequest) ProtoMessage() {}

func (x *SendDogsStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDogsStreamingRequest.ProtoReflect.Descriptor instead.
func (*SendDogsStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendDogsStreamingRequest) GetDog() *pets.Dog {
	if x != nil {
		return x.Dog
	}
	return nil
}

type SendDogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendDogsResponse) Reset() {
	*x = SendDogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDogsResponse) ProtoMessage() {}

func (x *SendDogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDogsResponse.ProtoReflect.Descriptor instead.
func (*SendDogsResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{1}
}

type ReceiveDogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveDogsRequest) Reset() {
	*x = ReceiveDogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveDogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveDogsRequest) ProtoMessage() {}

func (x *ReceiveDogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveDogsRequest.ProtoReflect.Descriptor instead.
func (*ReceiveDogsRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{2}
}

type ReceiveDogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dog *pets.Dog `protobuf:"bytes,1,opt,name=dog,proto3" json:"dog,omitempty"`
}

func (x *ReceiveDogsResponse) Reset() {
	*x = ReceiveDogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveDogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveDogsResponse) ProtoMessage() {}

func (x *ReceiveDogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveDogsResponse.ProtoReflect.Descriptor instead.
func (*ReceiveDogsResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveDogsResponse) GetDog() *pets.Dog {
	if x != nil {
		return x.Dog
	}
	return nil
}

type SendCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *pets.Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *SendCatRequest) Reset() {
	*x = SendCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCatRequest) ProtoMessage() {}

func (x *SendCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCatRequest.ProtoReflect.Descriptor instead.
func (*SendCatRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{4}
}

func (x *SendCatRequest) GetCat() *pets.Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type SendCatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendCatResponse) Reset() {
	*x = SendCatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCatResponse) ProtoMessage() {}

func (x *SendCatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCatResponse.ProtoReflect.Descriptor instead.
func (*SendCatResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{5}
}

type ReceiveCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveCatRequest) Reset() {
	*x = ReceiveCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveCatRequest) ProtoMessage() {}

func (x *ReceiveCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveCatRequest.ProtoReflect.Descriptor instead.
func (*ReceiveCatRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{6}
}

type ReceiveCatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *pets.Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *ReceiveCatResponse) Reset() {
	*x = ReceiveCatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_pet_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveCatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveCatResponse) ProtoMessage() {}

func (x *ReceiveCatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_pet_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveCatResponse.ProtoReflect.Descriptor instead.
func (*ReceiveCatResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_pet_service_proto_rawDescGZIP(), []int{7}
}

func (x *ReceiveCatResponse) GetCat() *pets.Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

var File_goadesign_goagen_pet_service_proto protoreflect.FileDescriptor

var file_goadesign_goagen_pet_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0e, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x67, 0x73, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x03, 0x64, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x65, 0x74,
	0x73, 0x2e, 0x44, 0x6f, 0x67, 0x52, 0x03, 0x64, 0x6f, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x44, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x44, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x44,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x64,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x2e,
	0x44, 0x6f, 0x67, 0x52, 0x03, 0x64, 0x6f, 0x67, 0x22, 0x2d, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x63, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x31, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x03, 0x63,
	0x61, 0x74, 0x32, 0xc9, 0x02, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x52, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x67, 0x73, 0x12, 0x25, 0x2e,
	0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x6f, 0x67, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x44, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x44, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x44, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x61, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x12, 0x1e, 0x2e,
	0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10,
	0x5a, 0x0e, 0x2f, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_pet_service_proto_rawDescOnce sync.Once
	file_goadesign_goagen_pet_service_proto_rawDescData = file_goadesign_goagen_pet_service_proto_rawDesc
)

func file_goadesign_goagen_pet_service_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_pet_service_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_pet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_pet_service_proto_rawDescData)
	})
	return file_goadesign_goagen_pet_service_proto_rawDescData
}

var file_goadesign_goagen_pet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goadesign_goagen_pet_service_proto_goTypes = []interface{}{
	(*SendDogsStreamingRequest)(nil), // 0: pet_service.SendDogsStreamingRequest
	(*SendDogsResponse)(nil),         // 1: pet_service.SendDogsResponse
	(*ReceiveDogsRequest)(nil),       // 2: pet_service.ReceiveDogsRequest
	(*ReceiveDogsResponse)(nil),      // 3: pet_service.ReceiveDogsResponse
	(*SendCatRequest)(nil),           // 4: pet_service.SendCatRequest
	(*SendCatResponse)(nil),          // 5: pet_service.SendCatResponse
	(*ReceiveCatRequest)(nil),        // 6: pet_service.ReceiveCatRequest
	(*ReceiveCatResponse)(nil),       // 7: pet_service.ReceiveCatResponse
	(*pets.Dog)(nil),                 // 8: pets.Dog
	(*pets.Cat)(nil),                 // 9: pets.Cat
}
var file_goadesign_goagen_pet_service_proto_depIdxs = []int32{
	8, // 0: pet_service.SendDogsStreamingRequest.dog:type_name -> pets.Dog
	8, // 1: pet_service.ReceiveDogsResponse.dog:type_name -> pets.Dog
	9, // 2: pet_service.SendCatRequest.cat:type_name -> pets.Cat
	9, // 3: pet_service.ReceiveCatResponse.cat:type_name -> pets.Cat
	0, // 4: pet_service.PetService.SendDogs:input_type -> pet_service.SendDogsStreamingRequest
	2, // 5: pet_service.PetService.ReceiveDogs:input_type -> pet_service.ReceiveDogsRequest
	4, // 6: pet_service.PetService.SendCat:input_type -> pet_service.SendCatRequest
	6, // 7: pet_service.PetService.ReceiveCat:input_type -> pet_service.ReceiveCatRequest
	1, // 8: pet_service.PetService.SendDogs:output_type -> pet_service.SendDogsResponse
	3, // 9: pet_service.PetService.ReceiveDogs:output_type -> pet_service.ReceiveDogsResponse
	5, // 10: pet_service.PetService.SendCat:output_type -> pet_service.SendCatResponse
	7, // 11: pet_service.PetService.ReceiveCat:output_type -> pet_service.ReceiveCatResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_pet_service_proto_init() }
func file_goadesign_goagen_pet_service_proto_init() {
	if File_goadesign_goagen_pet_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_pet_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDogsStreamingRequest); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDogsResponse); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveDogsRequest); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveDogsResponse); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCatRequest); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCatResponse); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveCatRequest); i {
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
		file_goadesign_goagen_pet_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveCatResponse); i {
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
			RawDescriptor: file_goadesign_goagen_pet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_pet_service_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_pet_service_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_pet_service_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_pet_service_proto = out.File
	file_goadesign_goagen_pet_service_proto_rawDesc = nil
	file_goadesign_goagen_pet_service_proto_goTypes = nil
	file_goadesign_goagen_pet_service_proto_depIdxs = nil
}
