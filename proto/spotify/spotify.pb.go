// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/spotify/spotify.proto

package go_micro_srv_spotify

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

type PlayRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayRequest) Reset()         { *m = PlayRequest{} }
func (m *PlayRequest) String() string { return proto.CompactTextString(m) }
func (*PlayRequest) ProtoMessage()    {}
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5585db63b8ef2524, []int{0}
}

func (m *PlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayRequest.Unmarshal(m, b)
}
func (m *PlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayRequest.Marshal(b, m, deterministic)
}
func (m *PlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayRequest.Merge(m, src)
}
func (m *PlayRequest) XXX_Size() int {
	return xxx_messageInfo_PlayRequest.Size(m)
}
func (m *PlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayRequest proto.InternalMessageInfo

func (m *PlayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PlayResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayResponse) Reset()         { *m = PlayResponse{} }
func (m *PlayResponse) String() string { return proto.CompactTextString(m) }
func (*PlayResponse) ProtoMessage()    {}
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5585db63b8ef2524, []int{1}
}

func (m *PlayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayResponse.Unmarshal(m, b)
}
func (m *PlayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayResponse.Marshal(b, m, deterministic)
}
func (m *PlayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayResponse.Merge(m, src)
}
func (m *PlayResponse) XXX_Size() int {
	return xxx_messageInfo_PlayResponse.Size(m)
}
func (m *PlayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlayResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5585db63b8ef2524, []int{2}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5585db63b8ef2524, []int{3}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PlayRequest)(nil), "go.micro.srv.spotify.PlayRequest")
	proto.RegisterType((*PlayResponse)(nil), "go.micro.srv.spotify.PlayResponse")
	proto.RegisterType((*StopRequest)(nil), "go.micro.srv.spotify.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "go.micro.srv.spotify.StopResponse")
}

func init() { proto.RegisterFile("proto/spotify/spotify.proto", fileDescriptor_5585db63b8ef2524) }

var fileDescriptor_5585db63b8ef2524 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2e, 0xc8, 0x2f, 0xc9, 0x4c, 0xab, 0x84, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0x91,
	0xf4, 0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0x3d, 0xa8, 0x9c, 0x92,
	0x22, 0x17, 0x77, 0x40, 0x4e, 0x62, 0x65, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10,
	0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xc4,
	0xc7, 0xc5, 0x03, 0x51, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xc4, 0xcb, 0xc5, 0x1d, 0x5c,
	0x92, 0x5f, 0x00, 0xd5, 0x02, 0x92, 0x86, 0x70, 0x21, 0xd2, 0x46, 0xab, 0x19, 0xb9, 0xd8, 0x83,
	0x21, 0xa6, 0x0b, 0xf9, 0x73, 0xb1, 0x80, 0xb4, 0x0a, 0x29, 0xea, 0x61, 0xb3, 0x5c, 0x0f, 0xc9,
	0x66, 0x29, 0x25, 0x7c, 0x4a, 0xa0, 0x36, 0x33, 0x80, 0x0c, 0x04, 0x59, 0x86, 0xcb, 0x40, 0x24,
	0x77, 0xe1, 0x32, 0x10, 0xd9, 0xad, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xc0, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x76, 0x8a, 0xda, 0x3b, 0x01, 0x00, 0x00,
}
