// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: message.proto

package pb

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

type TextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Textreq string `protobuf:"bytes,1,opt,name=textreq,proto3" json:"textreq,omitempty"`
}

func (x *TextRequest) Reset() {
	*x = TextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextRequest) ProtoMessage() {}

func (x *TextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextRequest.ProtoReflect.Descriptor instead.
func (*TextRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *TextRequest) GetTextreq() string {
	if x != nil {
		return x.Textreq
	}
	return ""
}

type TextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Textres string `protobuf:"bytes,1,opt,name=textres,proto3" json:"textres,omitempty"`
}

func (x *TextResponse) Reset() {
	*x = TextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextResponse) ProtoMessage() {}

func (x *TextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextResponse.ProtoReflect.Descriptor instead.
func (*TextResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *TextResponse) GetTextres() string {
	if x != nil {
		return x.Textres
	}
	return ""
}

type BytesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytesreq []byte `protobuf:"bytes,1,opt,name=bytesreq,proto3" json:"bytesreq,omitempty"`
}

func (x *BytesRequest) Reset() {
	*x = BytesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesRequest) ProtoMessage() {}

func (x *BytesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesRequest.ProtoReflect.Descriptor instead.
func (*BytesRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *BytesRequest) GetBytesreq() []byte {
	if x != nil {
		return x.Bytesreq
	}
	return nil
}

type BytesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytesres []byte `protobuf:"bytes,1,opt,name=bytesres,proto3" json:"bytesres,omitempty"`
}

func (x *BytesResponse) Reset() {
	*x = BytesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesResponse) ProtoMessage() {}

func (x *BytesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesResponse.ProtoReflect.Descriptor instead.
func (*BytesResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *BytesResponse) GetBytesres() []byte {
	if x != nil {
		return x.Bytesres
	}
	return nil
}

type RandomData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C string `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E string `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	F string `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G string `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
}

func (x *RandomData) Reset() {
	*x = RandomData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomData) ProtoMessage() {}

func (x *RandomData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomData.ProtoReflect.Descriptor instead.
func (*RandomData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *RandomData) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *RandomData) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *RandomData) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *RandomData) GetD() string {
	if x != nil {
		return x.D
	}
	return ""
}

func (x *RandomData) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

func (x *RandomData) GetF() string {
	if x != nil {
		return x.F
	}
	return ""
}

func (x *RandomData) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

type BigData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*RandomData `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *BigData) Reset() {
	*x = BigData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigData) ProtoMessage() {}

func (x *BigData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigData.ProtoReflect.Descriptor instead.
func (*BigData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *BigData) GetContent() []*RandomData {
	if x != nil {
		return x.Content
	}
	return nil
}

type BigDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bigdatareq    *BigData `protobuf:"bytes,1,opt,name=bigdatareq,proto3" json:"bigdatareq,omitempty"`
	Returnbigdata bool     `protobuf:"varint,2,opt,name=returnbigdata,proto3" json:"returnbigdata,omitempty"`
}

func (x *BigDataRequest) Reset() {
	*x = BigDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigDataRequest) ProtoMessage() {}

func (x *BigDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigDataRequest.ProtoReflect.Descriptor instead.
func (*BigDataRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *BigDataRequest) GetBigdatareq() *BigData {
	if x != nil {
		return x.Bigdatareq
	}
	return nil
}

func (x *BigDataRequest) GetReturnbigdata() bool {
	if x != nil {
		return x.Returnbigdata
	}
	return false
}

type BigDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bigdatares *BigData `protobuf:"bytes,1,opt,name=bigdatares,proto3" json:"bigdatares,omitempty"`
}

func (x *BigDataResponse) Reset() {
	*x = BigDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigDataResponse) ProtoMessage() {}

func (x *BigDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigDataResponse.ProtoReflect.Descriptor instead.
func (*BigDataResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *BigDataResponse) GetBigdatares() *BigData {
	if x != nil {
		return x.Bigdatares
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x27, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x65, 0x78, 0x74, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x65, 0x78, 0x74, 0x72, 0x65, 0x71, 0x22, 0x28, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x78, 0x74,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x78, 0x74, 0x72,
	0x65, 0x73, 0x22, 0x2a, 0x0a, 0x0c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x72, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x72, 0x65, 0x71, 0x22, 0x2b,
	0x0a, 0x0d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x72, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0a, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x66, 0x12, 0x0c, 0x0a,
	0x01, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67, 0x22, 0x30, 0x0a, 0x07, 0x42,
	0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x60, 0x0a,
	0x0e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x0a, 0x62, 0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x62,
	0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x72, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x62, 0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x62, 0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x3b, 0x0a, 0x0f, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x62, 0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0a, 0x62, 0x69, 0x67, 0x64, 0x61, 0x74, 0x61, 0x72, 0x65, 0x73, 0x32, 0x38, 0x0a, 0x0b,
	0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x54,
	0x65, 0x78, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x3c, 0x0a, 0x0c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46,
	0x75, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x44, 0x0a, 0x0e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x0f, 0x2e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []interface{}{
	(*TextRequest)(nil),     // 0: TextRequest
	(*TextResponse)(nil),    // 1: TextResponse
	(*BytesRequest)(nil),    // 2: BytesRequest
	(*BytesResponse)(nil),   // 3: BytesResponse
	(*RandomData)(nil),      // 4: RandomData
	(*BigData)(nil),         // 5: BigData
	(*BigDataRequest)(nil),  // 6: BigDataRequest
	(*BigDataResponse)(nil), // 7: BigDataResponse
}
var file_message_proto_depIdxs = []int32{
	4, // 0: BigData.content:type_name -> RandomData
	5, // 1: BigDataRequest.bigdatareq:type_name -> BigData
	5, // 2: BigDataResponse.bigdatares:type_name -> BigData
	0, // 3: TextService.TextFunc:input_type -> TextRequest
	2, // 4: BytesService.BytesFunc:input_type -> BytesRequest
	6, // 5: BigDataService.BigDataFunc:input_type -> BigDataRequest
	1, // 6: TextService.TextFunc:output_type -> TextResponse
	3, // 7: BytesService.BytesFunc:output_type -> BytesResponse
	7, // 8: BigDataService.BigDataFunc:output_type -> BigDataResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextResponse); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesResponse); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomData); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigData); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigDataRequest); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigDataResponse); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
