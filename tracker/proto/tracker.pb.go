// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: tracker.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingCode string `protobuf:"bytes,1,opt,name=tracking_code,json=trackingCode,proto3" json:"tracking_code,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

// Event doc...
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Status string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Place  string                 `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[1]
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
	return file_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Event) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Event) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingCode string   `protobuf:"bytes,1,opt,name=tracking_code,json=trackingCode,proto3" json:"tracking_code,omitempty"`
	Events       []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

func (x *SearchResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type SearchBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingCodes []string `protobuf:"bytes,1,rep,name=tracking_codes,json=trackingCodes,proto3" json:"tracking_codes,omitempty"`
}

func (x *SearchBatchRequest) Reset() {
	*x = SearchBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBatchRequest) ProtoMessage() {}

func (x *SearchBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBatchRequest.ProtoReflect.Descriptor instead.
func (*SearchBatchRequest) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{3}
}

func (x *SearchBatchRequest) GetTrackingCodes() []string {
	if x != nil {
		return x.TrackingCodes
	}
	return nil
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{4}
}

func (x *Credentials) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Credentials) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ExpiresIn int64  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{5}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Token) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_tracker_proto protoreflect.FileDescriptor

var file_tracker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x65, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x22, 0x5b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x3b, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x73, 0x73, 0x22, 0x3c, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x32, 0x84, 0x02, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x59,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x12, 0x5d, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x63, 0x68, 0x69, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x70, 0x6f, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tracker_proto_rawDescOnce sync.Once
	file_tracker_proto_rawDescData = file_tracker_proto_rawDesc
)

func file_tracker_proto_rawDescGZIP() []byte {
	file_tracker_proto_rawDescOnce.Do(func() {
		file_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracker_proto_rawDescData)
	})
	return file_tracker_proto_rawDescData
}

var file_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tracker_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),         // 0: proto.SearchRequest
	(*Event)(nil),                 // 1: proto.Event
	(*SearchResponse)(nil),        // 2: proto.SearchResponse
	(*SearchBatchRequest)(nil),    // 3: proto.SearchBatchRequest
	(*Credentials)(nil),           // 4: proto.Credentials
	(*Token)(nil),                 // 5: proto.Token
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_tracker_proto_depIdxs = []int32{
	6, // 0: proto.Event.date:type_name -> google.protobuf.Timestamp
	1, // 1: proto.SearchResponse.events:type_name -> proto.Event
	4, // 2: proto.Tracker.Login:input_type -> proto.Credentials
	0, // 3: proto.Tracker.Search:input_type -> proto.SearchRequest
	3, // 4: proto.Tracker.SearchBatch:input_type -> proto.SearchBatchRequest
	5, // 5: proto.Tracker.Login:output_type -> proto.Token
	2, // 6: proto.Tracker.Search:output_type -> proto.SearchResponse
	2, // 7: proto.Tracker.SearchBatch:output_type -> proto.SearchResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tracker_proto_init() }
func file_tracker_proto_init() {
	if File_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchBatchRequest); i {
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
		file_tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracker_proto_goTypes,
		DependencyIndexes: file_tracker_proto_depIdxs,
		MessageInfos:      file_tracker_proto_msgTypes,
	}.Build()
	File_tracker_proto = out.File
	file_tracker_proto_rawDesc = nil
	file_tracker_proto_goTypes = nil
	file_tracker_proto_depIdxs = nil
}
