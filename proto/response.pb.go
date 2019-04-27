// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package ictBridge

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FindTransactionsByAddressResponse struct {
	Transaction          []*Transaction `protobuf:"bytes,1,rep,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FindTransactionsByAddressResponse) Reset()         { *m = FindTransactionsByAddressResponse{} }
func (m *FindTransactionsByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*FindTransactionsByAddressResponse) ProtoMessage()    {}
func (*FindTransactionsByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *FindTransactionsByAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTransactionsByAddressResponse.Unmarshal(m, b)
}
func (m *FindTransactionsByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTransactionsByAddressResponse.Marshal(b, m, deterministic)
}
func (m *FindTransactionsByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTransactionsByAddressResponse.Merge(m, src)
}
func (m *FindTransactionsByAddressResponse) XXX_Size() int {
	return xxx_messageInfo_FindTransactionsByAddressResponse.Size(m)
}
func (m *FindTransactionsByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTransactionsByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindTransactionsByAddressResponse proto.InternalMessageInfo

func (m *FindTransactionsByAddressResponse) GetTransaction() []*Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type FindTransactionsByTagResponse struct {
	Transaction          []*Transaction `protobuf:"bytes,1,rep,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FindTransactionsByTagResponse) Reset()         { *m = FindTransactionsByTagResponse{} }
func (m *FindTransactionsByTagResponse) String() string { return proto.CompactTextString(m) }
func (*FindTransactionsByTagResponse) ProtoMessage()    {}
func (*FindTransactionsByTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
}

func (m *FindTransactionsByTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTransactionsByTagResponse.Unmarshal(m, b)
}
func (m *FindTransactionsByTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTransactionsByTagResponse.Marshal(b, m, deterministic)
}
func (m *FindTransactionsByTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTransactionsByTagResponse.Merge(m, src)
}
func (m *FindTransactionsByTagResponse) XXX_Size() int {
	return xxx_messageInfo_FindTransactionsByTagResponse.Size(m)
}
func (m *FindTransactionsByTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTransactionsByTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindTransactionsByTagResponse proto.InternalMessageInfo

func (m *FindTransactionsByTagResponse) GetTransaction() []*Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type FindTransactionByHashResponse struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindTransactionByHashResponse) Reset()         { *m = FindTransactionByHashResponse{} }
func (m *FindTransactionByHashResponse) String() string { return proto.CompactTextString(m) }
func (*FindTransactionByHashResponse) ProtoMessage()    {}
func (*FindTransactionByHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{2}
}

func (m *FindTransactionByHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTransactionByHashResponse.Unmarshal(m, b)
}
func (m *FindTransactionByHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTransactionByHashResponse.Marshal(b, m, deterministic)
}
func (m *FindTransactionByHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTransactionByHashResponse.Merge(m, src)
}
func (m *FindTransactionByHashResponse) XXX_Size() int {
	return xxx_messageInfo_FindTransactionByHashResponse.Size(m)
}
func (m *FindTransactionByHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTransactionByHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindTransactionByHashResponse proto.InternalMessageInfo

func (m *FindTransactionByHashResponse) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type PollEffectResponse struct {
	Effect               string   `protobuf:"bytes,1,opt,name=effect,proto3" json:"effect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollEffectResponse) Reset()         { *m = PollEffectResponse{} }
func (m *PollEffectResponse) String() string { return proto.CompactTextString(m) }
func (*PollEffectResponse) ProtoMessage()    {}
func (*PollEffectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{3}
}

func (m *PollEffectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollEffectResponse.Unmarshal(m, b)
}
func (m *PollEffectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollEffectResponse.Marshal(b, m, deterministic)
}
func (m *PollEffectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollEffectResponse.Merge(m, src)
}
func (m *PollEffectResponse) XXX_Size() int {
	return xxx_messageInfo_PollEffectResponse.Size(m)
}
func (m *PollEffectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PollEffectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PollEffectResponse proto.InternalMessageInfo

func (m *PollEffectResponse) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

type TakeEffectResponse struct {
	Effect               string   `protobuf:"bytes,1,opt,name=effect,proto3" json:"effect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeEffectResponse) Reset()         { *m = TakeEffectResponse{} }
func (m *TakeEffectResponse) String() string { return proto.CompactTextString(m) }
func (*TakeEffectResponse) ProtoMessage()    {}
func (*TakeEffectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{4}
}

func (m *TakeEffectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeEffectResponse.Unmarshal(m, b)
}
func (m *TakeEffectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeEffectResponse.Marshal(b, m, deterministic)
}
func (m *TakeEffectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeEffectResponse.Merge(m, src)
}
func (m *TakeEffectResponse) XXX_Size() int {
	return xxx_messageInfo_TakeEffectResponse.Size(m)
}
func (m *TakeEffectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeEffectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TakeEffectResponse proto.InternalMessageInfo

func (m *TakeEffectResponse) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func init() {
	proto.RegisterType((*FindTransactionsByAddressResponse)(nil), "FindTransactionsByAddressResponse")
	proto.RegisterType((*FindTransactionsByTagResponse)(nil), "FindTransactionsByTagResponse")
	proto.RegisterType((*FindTransactionByHashResponse)(nil), "FindTransactionByHashResponse")
	proto.RegisterType((*PollEffectResponse)(nil), "PollEffectResponse")
	proto.RegisterType((*TakeEffectResponse)(nil), "TakeEffectResponse")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0x81, 0x70, 0x94, 0x82, 0xb9, 0x14, 0xdd, 0x32, 0xf3, 0x52, 0x42, 0x8a, 0x12, 0xf3, 0x8a,
	0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x9d, 0x2a, 0x1d, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b,
	0x83, 0xa0, 0xfa, 0x84, 0xf4, 0xb8, 0xb8, 0x4b, 0x10, 0x0a, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8,
	0x8d, 0x78, 0xf4, 0x90, 0x34, 0x05, 0x21, 0x2b, 0x50, 0xf2, 0xe7, 0x92, 0xc5, 0x34, 0x34, 0x24,
	0x31, 0x9d, 0x8a, 0x06, 0x3a, 0x55, 0x7a, 0x24, 0x16, 0x67, 0xe0, 0x36, 0x90, 0x11, 0xbf, 0x81,
	0x3a, 0x5c, 0x42, 0x01, 0xf9, 0x39, 0x39, 0xae, 0x69, 0x69, 0xa9, 0xc9, 0x25, 0x70, 0x53, 0xc4,
	0xb8, 0xd8, 0x52, 0xc1, 0x22, 0x60, 0x03, 0x38, 0x83, 0xa0, 0x3c, 0x90, 0xea, 0x90, 0xc4, 0xec,
	0x54, 0xe2, 0x54, 0x27, 0xb1, 0x81, 0x43, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x64,
	0x01, 0x37, 0x78, 0x01, 0x00, 0x00,
}