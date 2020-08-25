// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_client.proto

package client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Message_TextMessage_
	Value                isMessage_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type isMessage_Value interface {
	isMessage_Value()
}

type Message_TextMessage_ struct {
	TextMessage *Message_TextMessage `protobuf:"bytes,3,opt,name=text_message,json=textMessage,proto3,oneof"`
}

func (*Message_TextMessage_) isMessage_Value() {}

func (m *Message) GetValue() isMessage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetTextMessage() *Message_TextMessage {
	if x, ok := m.GetValue().(*Message_TextMessage_); ok {
		return x.TextMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_TextMessage_)(nil),
	}
}

type Message_TextMessage struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_TextMessage) Reset()         { *m = Message_TextMessage{} }
func (m *Message_TextMessage) String() string { return proto.CompactTextString(m) }
func (*Message_TextMessage) ProtoMessage()    {}
func (*Message_TextMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{0, 0}
}

func (m *Message_TextMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_TextMessage.Unmarshal(m, b)
}
func (m *Message_TextMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_TextMessage.Marshal(b, m, deterministic)
}
func (m *Message_TextMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_TextMessage.Merge(m, src)
}
func (m *Message_TextMessage) XXX_Size() int {
	return xxx_messageInfo_Message_TextMessage.Size(m)
}
func (m *Message_TextMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_TextMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Message_TextMessage proto.InternalMessageInfo

func (m *Message_TextMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Error struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessageRequest struct {
	ConversationId       int64    `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Messages             *Message `protobuf:"bytes,2,opt,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{2}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetConversationId() int64 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *SendMessageRequest) GetMessages() *Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type SendMessageResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{3}
}

func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(m, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

func (m *SendMessageResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type WaitMessageRequest struct {
	ConversationId       int64    `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ConfirmationId       string   `protobuf:"bytes,2,opt,name=confirmation_id,json=confirmationId,proto3" json:"confirmation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitMessageRequest) Reset()         { *m = WaitMessageRequest{} }
func (m *WaitMessageRequest) String() string { return proto.CompactTextString(m) }
func (*WaitMessageRequest) ProtoMessage()    {}
func (*WaitMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{4}
}

func (m *WaitMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitMessageRequest.Unmarshal(m, b)
}
func (m *WaitMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitMessageRequest.Marshal(b, m, deterministic)
}
func (m *WaitMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitMessageRequest.Merge(m, src)
}
func (m *WaitMessageRequest) XXX_Size() int {
	return xxx_messageInfo_WaitMessageRequest.Size(m)
}
func (m *WaitMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitMessageRequest proto.InternalMessageInfo

func (m *WaitMessageRequest) GetConversationId() int64 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *WaitMessageRequest) GetConfirmationId() string {
	if m != nil {
		return m.ConfirmationId
	}
	return ""
}

type WaitMessageResponse struct {
	TimeoutSec           int64      `protobuf:"varint,1,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	Messages             []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Error                *Error     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WaitMessageResponse) Reset()         { *m = WaitMessageResponse{} }
func (m *WaitMessageResponse) String() string { return proto.CompactTextString(m) }
func (*WaitMessageResponse) ProtoMessage()    {}
func (*WaitMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{5}
}

func (m *WaitMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitMessageResponse.Unmarshal(m, b)
}
func (m *WaitMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitMessageResponse.Marshal(b, m, deterministic)
}
func (m *WaitMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitMessageResponse.Merge(m, src)
}
func (m *WaitMessageResponse) XXX_Size() int {
	return xxx_messageInfo_WaitMessageResponse.Size(m)
}
func (m *WaitMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitMessageResponse proto.InternalMessageInfo

func (m *WaitMessageResponse) GetTimeoutSec() int64 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *WaitMessageResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *WaitMessageResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type CloseConversationRequest struct {
	ConversationId       int64    `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Cause                string   `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseConversationRequest) Reset()         { *m = CloseConversationRequest{} }
func (m *CloseConversationRequest) String() string { return proto.CompactTextString(m) }
func (*CloseConversationRequest) ProtoMessage()    {}
func (*CloseConversationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{6}
}

func (m *CloseConversationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConversationRequest.Unmarshal(m, b)
}
func (m *CloseConversationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConversationRequest.Marshal(b, m, deterministic)
}
func (m *CloseConversationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConversationRequest.Merge(m, src)
}
func (m *CloseConversationRequest) XXX_Size() int {
	return xxx_messageInfo_CloseConversationRequest.Size(m)
}
func (m *CloseConversationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConversationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConversationRequest proto.InternalMessageInfo

func (m *CloseConversationRequest) GetConversationId() int64 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *CloseConversationRequest) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type CloseConversationResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseConversationResponse) Reset()         { *m = CloseConversationResponse{} }
func (m *CloseConversationResponse) String() string { return proto.CompactTextString(m) }
func (*CloseConversationResponse) ProtoMessage()    {}
func (*CloseConversationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f033c3cf63e42ba, []int{7}
}

func (m *CloseConversationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConversationResponse.Unmarshal(m, b)
}
func (m *CloseConversationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConversationResponse.Marshal(b, m, deterministic)
}
func (m *CloseConversationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConversationResponse.Merge(m, src)
}
func (m *CloseConversationResponse) XXX_Size() int {
	return xxx_messageInfo_CloseConversationResponse.Size(m)
}
func (m *CloseConversationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConversationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConversationResponse proto.InternalMessageInfo

func (m *CloseConversationResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "client.Message")
	proto.RegisterType((*Message_TextMessage)(nil), "client.Message.TextMessage")
	proto.RegisterType((*Error)(nil), "client.Error")
	proto.RegisterType((*SendMessageRequest)(nil), "client.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "client.SendMessageResponse")
	proto.RegisterType((*WaitMessageRequest)(nil), "client.WaitMessageRequest")
	proto.RegisterType((*WaitMessageResponse)(nil), "client.WaitMessageResponse")
	proto.RegisterType((*CloseConversationRequest)(nil), "client.CloseConversationRequest")
	proto.RegisterType((*CloseConversationResponse)(nil), "client.CloseConversationResponse")
}

func init() { proto.RegisterFile("chat_client.proto", fileDescriptor_6f033c3cf63e42ba) }

var fileDescriptor_6f033c3cf63e42ba = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x8f, 0x93, 0x40,
	0x14, 0x2f, 0x20, 0x5b, 0xf7, 0xa1, 0xdd, 0x74, 0xd6, 0x03, 0xb2, 0x07, 0xd9, 0xf1, 0x60, 0x13,
	0x93, 0x26, 0xe2, 0xcd, 0xd3, 0xc6, 0x46, 0xb3, 0x7b, 0xf0, 0x42, 0x4d, 0x8c, 0x5e, 0x08, 0xc2,
	0x5b, 0x1d, 0x43, 0x99, 0xca, 0x0c, 0x15, 0xcf, 0x5e, 0xfc, 0x0e, 0x7e, 0x59, 0xd3, 0x99, 0xa1,
	0xb0, 0x6d, 0x6d, 0x52, 0x6f, 0xbc, 0x7f, 0xbf, 0x3f, 0xef, 0x4d, 0x80, 0x71, 0xf6, 0x35, 0x95,
	0x49, 0x56, 0x30, 0x2c, 0xe5, 0x74, 0x59, 0x71, 0xc9, 0xc9, 0x89, 0x8e, 0xe8, 0x1f, 0x0b, 0x86,
	0xef, 0x50, 0x88, 0xf4, 0x0b, 0x92, 0x11, 0xd8, 0x2c, 0xf7, 0xad, 0xd0, 0x9a, 0x38, 0xb1, 0xcd,
	0x72, 0x42, 0xe0, 0x9e, 0xfc, 0xb9, 0x44, 0xdf, 0x0e, 0xad, 0xc9, 0x69, 0xac, 0xbe, 0xc9, 0x15,
	0x3c, 0x90, 0xd8, 0xc8, 0x64, 0xa1, 0x67, 0x7c, 0x27, 0xb4, 0x26, 0x5e, 0x74, 0x31, 0x35, 0xe0,
	0x06, 0x6a, 0xfa, 0x1e, 0x1b, 0x69, 0xbe, 0xaf, 0x07, 0xb1, 0x27, 0xbb, 0x30, 0xb8, 0x04, 0xaf,
	0x57, 0x55, 0x24, 0xd8, 0x48, 0x45, 0xbb, 0x26, 0xc1, 0x46, 0xbe, 0x1e, 0x82, 0xbb, 0x4a, 0x8b,
	0x1a, 0xe9, 0x0b, 0x70, 0xdf, 0x54, 0x15, 0xaf, 0x7a, 0xd2, 0x4e, 0x95, 0x34, 0x1f, 0x86, 0xad,
	0x02, 0xad, 0xae, 0x0d, 0xe9, 0x37, 0x20, 0x73, 0x2c, 0x73, 0x03, 0x1f, 0xe3, 0xf7, 0x1a, 0x85,
	0x24, 0xcf, 0xe0, 0x2c, 0xe3, 0xe5, 0x0a, 0x2b, 0x91, 0x4a, 0xc6, 0xcb, 0x64, 0xe3, 0x73, 0xd4,
	0x4f, 0xdf, 0xe4, 0xe4, 0x39, 0xdc, 0x37, 0x48, 0x42, 0x21, 0x7b, 0xd1, 0xd9, 0x96, 0xb7, 0x78,
	0xd3, 0x40, 0x5f, 0xc1, 0xf9, 0x1d, 0x2e, 0xb1, 0xe4, 0xa5, 0x40, 0xf2, 0x14, 0x5c, 0x5c, 0xab,
	0x56, 0x14, 0x5e, 0xf4, 0xb0, 0x05, 0x50, 0x56, 0x62, 0x5d, 0xa3, 0xb7, 0x40, 0x3e, 0xa4, 0x4c,
	0xfe, 0xaf, 0x4e, 0xdd, 0x78, 0xcb, 0xaa, 0xc5, 0xa6, 0x51, 0x2f, 0x62, 0xd4, 0x4f, 0xdf, 0xe4,
	0xf4, 0xb7, 0x05, 0xe7, 0x77, 0x88, 0x8c, 0xc8, 0x27, 0xe0, 0x49, 0xb6, 0x40, 0x5e, 0xcb, 0x44,
	0x60, 0x66, 0x58, 0xc0, 0xa4, 0xe6, 0x98, 0x6d, 0x6d, 0xc2, 0x39, 0xb8, 0x89, 0xce, 0xb2, 0x73,
	0xc0, 0xf2, 0x47, 0xf0, 0x67, 0x05, 0x17, 0x38, 0xeb, 0x59, 0x39, 0xda, 0xf8, 0x23, 0x70, 0xb3,
	0xb4, 0x16, 0xed, 0xdd, 0x75, 0x40, 0xaf, 0xe0, 0xf1, 0x1e, 0xe8, 0x23, 0xee, 0x11, 0xfd, 0xb2,
	0x61, 0xfc, 0xb6, 0xe0, 0x3f, 0x66, 0xaa, 0x36, 0xc7, 0x6a, 0xc5, 0x32, 0x24, 0xd7, 0xe0, 0xf5,
	0x2e, 0x4c, 0x82, 0x76, 0x74, 0xf7, 0x89, 0x05, 0x17, 0x7b, 0x6b, 0x5a, 0x02, 0x1d, 0xac, 0x91,
	0x7a, 0x67, 0xe8, 0x90, 0x76, 0x1f, 0x41, 0x87, 0xb4, 0xe7, 0x6e, 0x74, 0x40, 0x3e, 0xc1, 0x78,
	0xc7, 0x2b, 0x09, 0xdb, 0x99, 0x7f, 0x6d, 0x38, 0xb8, 0x3c, 0xd0, 0xd1, 0x62, 0x7f, 0x3e, 0x51,
	0x7f, 0x87, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x10, 0xe2, 0x1f, 0x32, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlowClientServiceClient is the client API for FlowClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowClientServiceClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	WaitMessage(ctx context.Context, in *WaitMessageRequest, opts ...grpc.CallOption) (*WaitMessageResponse, error)
	CloseConversation(ctx context.Context, in *CloseConversationRequest, opts ...grpc.CallOption) (*CloseConversationResponse, error)
}

type flowClientServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlowClientServiceClient(cc *grpc.ClientConn) FlowClientServiceClient {
	return &flowClientServiceClient{cc}
}

func (c *flowClientServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/client.FlowClientService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowClientServiceClient) WaitMessage(ctx context.Context, in *WaitMessageRequest, opts ...grpc.CallOption) (*WaitMessageResponse, error) {
	out := new(WaitMessageResponse)
	err := c.cc.Invoke(ctx, "/client.FlowClientService/WaitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowClientServiceClient) CloseConversation(ctx context.Context, in *CloseConversationRequest, opts ...grpc.CallOption) (*CloseConversationResponse, error) {
	out := new(CloseConversationResponse)
	err := c.cc.Invoke(ctx, "/client.FlowClientService/CloseConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowClientServiceServer is the server API for FlowClientService service.
type FlowClientServiceServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	WaitMessage(context.Context, *WaitMessageRequest) (*WaitMessageResponse, error)
	CloseConversation(context.Context, *CloseConversationRequest) (*CloseConversationResponse, error)
}

// UnimplementedFlowClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlowClientServiceServer struct {
}

func (*UnimplementedFlowClientServiceServer) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedFlowClientServiceServer) WaitMessage(ctx context.Context, req *WaitMessageRequest) (*WaitMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitMessage not implemented")
}
func (*UnimplementedFlowClientServiceServer) CloseConversation(ctx context.Context, req *CloseConversationRequest) (*CloseConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConversation not implemented")
}

func RegisterFlowClientServiceServer(s *grpc.Server, srv FlowClientServiceServer) {
	s.RegisterService(&_FlowClientService_serviceDesc, srv)
}

func _FlowClientService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowClientServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.FlowClientService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowClientServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowClientService_WaitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowClientServiceServer).WaitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.FlowClientService/WaitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowClientServiceServer).WaitMessage(ctx, req.(*WaitMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowClientService_CloseConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowClientServiceServer).CloseConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.FlowClientService/CloseConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowClientServiceServer).CloseConversation(ctx, req.(*CloseConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlowClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.FlowClientService",
	HandlerType: (*FlowClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _FlowClientService_SendMessage_Handler,
		},
		{
			MethodName: "WaitMessage",
			Handler:    _FlowClientService_WaitMessage_Handler,
		},
		{
			MethodName: "CloseConversation",
			Handler:    _FlowClientService_CloseConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat_client.proto",
}
