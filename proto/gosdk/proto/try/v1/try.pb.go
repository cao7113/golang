// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: proto/try/v1/try.proto

package tryv1

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

type TryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Score uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	// Types that are assignable to Gender:
	//	*TryRequest_Male
	//	*TryRequest_Female
	Gender isTryRequest_Gender `protobuf_oneof:"gender"`
}

func (x *TryRequest) Reset() {
	*x = TryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TryRequest) ProtoMessage() {}

func (x *TryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TryRequest.ProtoReflect.Descriptor instead.
func (*TryRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{0}
}

func (x *TryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TryRequest) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (m *TryRequest) GetGender() isTryRequest_Gender {
	if m != nil {
		return m.Gender
	}
	return nil
}

func (x *TryRequest) GetMale() bool {
	if x, ok := x.GetGender().(*TryRequest_Male); ok {
		return x.Male
	}
	return false
}

func (x *TryRequest) GetFemale() bool {
	if x, ok := x.GetGender().(*TryRequest_Female); ok {
		return x.Female
	}
	return false
}

type isTryRequest_Gender interface {
	isTryRequest_Gender()
}

type TryRequest_Male struct {
	Male bool `protobuf:"varint,3,opt,name=male,proto3,oneof"`
}

type TryRequest_Female struct {
	Female bool `protobuf:"varint,4,opt,name=female,proto3,oneof"`
}

func (*TryRequest_Male) isTryRequest_Gender() {}

func (*TryRequest_Female) isTryRequest_Gender() {}

type TryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TryResponse) Reset() {
	*x = TryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TryResponse) ProtoMessage() {}

func (x *TryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TryResponse.ProtoReflect.Descriptor instead.
func (*TryResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{1}
}

func (x *TryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N    uint64 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{2}
}

func (x *FibonacciRequest) GetN() uint64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *FibonacciRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  uint64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	TakenMs int64  `protobuf:"varint,2,opt,name=taken_ms,json=takenMs,proto3" json:"taken_ms,omitempty"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{3}
}

func (x *FibonacciResponse) GetResult() uint64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *FibonacciResponse) GetTakenMs() int64 {
	if x != nil {
		return x.TakenMs
	}
	return 0
}

type TimeoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeoutInMs uint32 `protobuf:"varint,1,opt,name=timeout_in_ms,json=timeoutInMs,proto3" json:"timeout_in_ms,omitempty"`
}

func (x *TimeoutRequest) Reset() {
	*x = TimeoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutRequest) ProtoMessage() {}

func (x *TimeoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeoutRequest.ProtoReflect.Descriptor instead.
func (*TimeoutRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{4}
}

func (x *TimeoutRequest) GetTimeoutInMs() uint32 {
	if x != nil {
		return x.TimeoutInMs
	}
	return 0
}

type TimeoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TimeoutResponse) Reset() {
	*x = TimeoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutResponse) ProtoMessage() {}

func (x *TimeoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeoutResponse.ProtoReflect.Descriptor instead.
func (*TimeoutResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{5}
}

func (x *TimeoutResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NMs uint32 `protobuf:"varint,1,opt,name=n_ms,json=nMs,proto3" json:"n_ms,omitempty"`
}

func (x *SlowRequest) Reset() {
	*x = SlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowRequest) ProtoMessage() {}

func (x *SlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowRequest.ProtoReflect.Descriptor instead.
func (*SlowRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{6}
}

func (x *SlowRequest) GetNMs() uint32 {
	if x != nil {
		return x.NMs
	}
	return 0
}

type SlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SlowResponse) Reset() {
	*x = SlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowResponse) ProtoMessage() {}

func (x *SlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowResponse.ProtoReflect.Descriptor instead.
func (*SlowResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{7}
}

func (x *SlowResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type FatalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *FatalRequest) Reset() {
	*x = FatalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FatalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FatalRequest) ProtoMessage() {}

func (x *FatalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FatalRequest.ProtoReflect.Descriptor instead.
func (*FatalRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{8}
}

func (x *FatalRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type FatalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FatalResponse) Reset() {
	*x = FatalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FatalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FatalResponse) ProtoMessage() {}

func (x *FatalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FatalResponse.ProtoReflect.Descriptor instead.
func (*FatalResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{9}
}

func (x *FatalResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{10}
}

func (x *Error) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DetailErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *DetailErrorRequest) Reset() {
	*x = DetailErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailErrorRequest) ProtoMessage() {}

func (x *DetailErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailErrorRequest.ProtoReflect.Descriptor instead.
func (*DetailErrorRequest) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{11}
}

func (x *DetailErrorRequest) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DetailErrorRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type DetailErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DetailErrorResponse) Reset() {
	*x = DetailErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_try_v1_try_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailErrorResponse) ProtoMessage() {}

func (x *DetailErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_try_v1_try_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailErrorResponse.ProtoReflect.Descriptor instead.
func (*DetailErrorResponse) Descriptor() ([]byte, []int) {
	return file_proto_try_v1_try_proto_rawDescGZIP(), []int{12}
}

func (x *DetailErrorResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_try_v1_try_proto protoreflect.FileDescriptor

var file_proto_try_v1_try_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x0a, 0x54, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04,
	0x6d, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x54, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x34, 0x0a, 0x10, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x01, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x46, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e,
	0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x22,
	0x34, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x5f,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x49, 0x6e, 0x4d, 0x73, 0x22, 0x23, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0b, 0x53, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x6e, 0x5f, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e, 0x4d, 0x73, 0x22, 0x20, 0x0a, 0x0c,
	0x53, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22,
	0x0a, 0x0c, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x22, 0x21, 0x0a, 0x0d, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x12,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x27, 0x0a, 0x13, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0xb3, 0x03, 0x0a, 0x0a, 0x54, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x54, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x77, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x46, 0x61, 0x74, 0x61, 0x6c,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x61, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x74, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6f, 0x37, 0x31, 0x31, 0x33, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x72, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_try_v1_try_proto_rawDescOnce sync.Once
	file_proto_try_v1_try_proto_rawDescData = file_proto_try_v1_try_proto_rawDesc
)

func file_proto_try_v1_try_proto_rawDescGZIP() []byte {
	file_proto_try_v1_try_proto_rawDescOnce.Do(func() {
		file_proto_try_v1_try_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_try_v1_try_proto_rawDescData)
	})
	return file_proto_try_v1_try_proto_rawDescData
}

var file_proto_try_v1_try_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_try_v1_try_proto_goTypes = []interface{}{
	(*TryRequest)(nil),          // 0: proto.try.v1.TryRequest
	(*TryResponse)(nil),         // 1: proto.try.v1.TryResponse
	(*FibonacciRequest)(nil),    // 2: proto.try.v1.FibonacciRequest
	(*FibonacciResponse)(nil),   // 3: proto.try.v1.FibonacciResponse
	(*TimeoutRequest)(nil),      // 4: proto.try.v1.TimeoutRequest
	(*TimeoutResponse)(nil),     // 5: proto.try.v1.TimeoutResponse
	(*SlowRequest)(nil),         // 6: proto.try.v1.SlowRequest
	(*SlowResponse)(nil),        // 7: proto.try.v1.SlowResponse
	(*FatalRequest)(nil),        // 8: proto.try.v1.FatalRequest
	(*FatalResponse)(nil),       // 9: proto.try.v1.FatalResponse
	(*Error)(nil),               // 10: proto.try.v1.Error
	(*DetailErrorRequest)(nil),  // 11: proto.try.v1.DetailErrorRequest
	(*DetailErrorResponse)(nil), // 12: proto.try.v1.DetailErrorResponse
}
var file_proto_try_v1_try_proto_depIdxs = []int32{
	0,  // 0: proto.try.v1.TryService.Try:input_type -> proto.try.v1.TryRequest
	4,  // 1: proto.try.v1.TryService.Timeout:input_type -> proto.try.v1.TimeoutRequest
	6,  // 2: proto.try.v1.TryService.Slow:input_type -> proto.try.v1.SlowRequest
	2,  // 3: proto.try.v1.TryService.Fibonacci:input_type -> proto.try.v1.FibonacciRequest
	11, // 4: proto.try.v1.TryService.DetailError:input_type -> proto.try.v1.DetailErrorRequest
	8,  // 5: proto.try.v1.TryService.Fatal:input_type -> proto.try.v1.FatalRequest
	1,  // 6: proto.try.v1.TryService.Try:output_type -> proto.try.v1.TryResponse
	5,  // 7: proto.try.v1.TryService.Timeout:output_type -> proto.try.v1.TimeoutResponse
	7,  // 8: proto.try.v1.TryService.Slow:output_type -> proto.try.v1.SlowResponse
	3,  // 9: proto.try.v1.TryService.Fibonacci:output_type -> proto.try.v1.FibonacciResponse
	12, // 10: proto.try.v1.TryService.DetailError:output_type -> proto.try.v1.DetailErrorResponse
	9,  // 11: proto.try.v1.TryService.Fatal:output_type -> proto.try.v1.FatalResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_try_v1_try_proto_init() }
func file_proto_try_v1_try_proto_init() {
	if File_proto_try_v1_try_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_try_v1_try_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TryRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TryResponse); i {
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
		file_proto_try_v1_try_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciResponse); i {
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
		file_proto_try_v1_try_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutResponse); i {
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
		file_proto_try_v1_try_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowResponse); i {
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
		file_proto_try_v1_try_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FatalRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FatalResponse); i {
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
		file_proto_try_v1_try_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_proto_try_v1_try_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailErrorRequest); i {
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
		file_proto_try_v1_try_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailErrorResponse); i {
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
	file_proto_try_v1_try_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TryRequest_Male)(nil),
		(*TryRequest_Female)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_try_v1_try_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_try_v1_try_proto_goTypes,
		DependencyIndexes: file_proto_try_v1_try_proto_depIdxs,
		MessageInfos:      file_proto_try_v1_try_proto_msgTypes,
	}.Build()
	File_proto_try_v1_try_proto = out.File
	file_proto_try_v1_try_proto_rawDesc = nil
	file_proto_try_v1_try_proto_goTypes = nil
	file_proto_try_v1_try_proto_depIdxs = nil
}
