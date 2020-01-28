// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bullbear.proto

package bullbear

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

type Opinion int32

const (
	Opinion_NONE    Opinion = 0
	Opinion_BULLISH Opinion = 1
	Opinion_BEARISH Opinion = 2
)

var Opinion_name = map[int32]string{
	0: "NONE",
	1: "BULLISH",
	2: "BEARISH",
}

var Opinion_value = map[string]int32{
	"NONE":    0,
	"BULLISH": 1,
	"BEARISH": 2,
}

func (x Opinion) String() string {
	return proto.EnumName(Opinion_name, int32(x))
}

func (Opinion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{0}
}

type Resource struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	BullsCount           int32    `protobuf:"varint,3,opt,name=bulls_count,json=bullsCount,proto3" json:"bulls_count,omitempty"`
	BearsCount           int32    `protobuf:"varint,4,opt,name=bears_count,json=bearsCount,proto3" json:"bears_count,omitempty"`
	Bulls                []string `protobuf:"bytes,5,rep,name=bulls,proto3" json:"bulls,omitempty"`
	Bears                []string `protobuf:"bytes,6,rep,name=bears,proto3" json:"bears,omitempty"`
	Opinion              Opinion  `protobuf:"varint,7,opt,name=opinion,proto3,enum=Opinion" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetBullsCount() int32 {
	if m != nil {
		return m.BullsCount
	}
	return 0
}

func (m *Resource) GetBearsCount() int32 {
	if m != nil {
		return m.BearsCount
	}
	return 0
}

func (m *Resource) GetBulls() []string {
	if m != nil {
		return m.Bulls
	}
	return nil
}

func (m *Resource) GetBears() []string {
	if m != nil {
		return m.Bears
	}
	return nil
}

func (m *Resource) GetOpinion() Opinion {
	if m != nil {
		return m.Opinion
	}
	return Opinion_NONE
}

type Request struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	UserUuid             string    `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Opinion              Opinion   `protobuf:"varint,3,opt,name=opinion,proto3,enum=Opinion" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Request) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Request) GetOpinion() Opinion {
	if m != nil {
		return m.Opinion
	}
	return Opinion_NONE
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{2}
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

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ListRequest struct {
	ResourceType         string   `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceUuids        []string `protobuf:"bytes,2,rep,name=resource_uuids,json=resourceUuids,proto3" json:"resource_uuids,omitempty"`
	UserUuid             string   `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ListRequest) GetResourceUuids() []string {
	if m != nil {
		return m.ResourceUuids
	}
	return nil
}

func (m *ListRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

type ListResponse struct {
	Error                *Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resources            []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12cf591b457316ba, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ListResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterEnum("Opinion", Opinion_name, Opinion_value)
	proto.RegisterType((*Resource)(nil), "Resource")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
}

func init() { proto.RegisterFile("proto/bullbear.proto", fileDescriptor_12cf591b457316ba) }

var fileDescriptor_12cf591b457316ba = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x61, 0x8b, 0xda, 0x40,
	0x10, 0x35, 0xc6, 0x98, 0x64, 0xf4, 0x8e, 0x63, 0xb9, 0x0f, 0xcb, 0xb5, 0xe5, 0x64, 0x8b, 0x54,
	0x0a, 0x4d, 0xc1, 0xd2, 0x1f, 0x50, 0x0f, 0x69, 0x0b, 0x72, 0xc2, 0xb6, 0x7e, 0x96, 0xe8, 0x2d,
	0x45, 0xb0, 0x6e, 0xdc, 0xcd, 0x16, 0xfa, 0xff, 0xfa, 0xc3, 0xca, 0xcc, 0x66, 0x3d, 0xf3, 0x41,
	0xb8, 0x6f, 0xf3, 0xde, 0xbc, 0xcc, 0xec, 0x7b, 0x43, 0xe0, 0xb6, 0x32, 0xba, 0xd6, 0x1f, 0x37,
	0x6e, 0xbf, 0xdf, 0xa8, 0xd2, 0x14, 0x04, 0xc5, 0xbf, 0x08, 0x32, 0xa9, 0xac, 0x76, 0x66, 0xab,
	0x18, 0x83, 0x9e, 0x73, 0xbb, 0x27, 0x1e, 0x8d, 0xa2, 0x49, 0x2e, 0xa9, 0x46, 0xae, 0xfe, 0x5b,
	0x29, 0xde, 0xf5, 0x1c, 0xd6, 0xec, 0x1e, 0x06, 0x38, 0xc6, 0xae, 0xb7, 0xda, 0x1d, 0x6a, 0x1e,
	0x8f, 0xa2, 0x49, 0x22, 0x81, 0xa8, 0x07, 0x64, 0x48, 0xa0, 0x4a, 0x13, 0x04, 0xbd, 0x46, 0x80,
	0x94, 0x17, 0xdc, 0x42, 0x42, 0x72, 0x9e, 0x8c, 0xe2, 0x49, 0x2e, 0x3d, 0x20, 0x16, 0x35, 0xbc,
	0xdf, 0xb0, 0x08, 0x98, 0x80, 0x54, 0x57, 0xbb, 0xc3, 0x4e, 0x1f, 0x78, 0x3a, 0x8a, 0x26, 0xd7,
	0xd3, 0xac, 0x58, 0x7a, 0x2c, 0x43, 0x43, 0x1c, 0x21, 0x95, 0xea, 0xe8, 0x94, 0xad, 0xd9, 0x18,
	0x32, 0xd3, 0x18, 0x22, 0x23, 0x83, 0x69, 0x5e, 0x04, 0x87, 0xf2, 0xd4, 0x62, 0xaf, 0x20, 0x77,
	0x56, 0x99, 0x35, 0x19, 0xf6, 0xe6, 0x32, 0x24, 0x56, 0x68, 0xfa, 0x6c, 0x65, 0x7c, 0x69, 0xe5,
	0x67, 0x48, 0xe6, 0xc6, 0x68, 0x83, 0x09, 0x6d, 0xf5, 0x93, 0x5f, 0x96, 0x48, 0xaa, 0x19, 0x87,
	0xf4, 0xb7, 0xb2, 0xb6, 0xfc, 0x15, 0x82, 0x0b, 0x50, 0x2c, 0x29, 0xef, 0x4a, 0x1f, 0xac, 0x62,
	0xaf, 0x21, 0x51, 0x38, 0xa2, 0x79, 0x67, 0xbf, 0xa0, 0x81, 0xd2, 0x93, 0x2d, 0x23, 0xdd, 0x8b,
	0x46, 0xc4, 0x1f, 0x18, 0x2c, 0x76, 0xb6, 0x0e, 0xf6, 0xdf, 0xc2, 0x55, 0x68, 0xad, 0xe9, 0x70,
	0xfe, 0x98, 0xc3, 0x40, 0xfe, 0xc4, 0x03, 0x8e, 0xe1, 0xfa, 0x24, 0xc2, 0x00, 0x2c, 0xef, 0x52,
	0xe2, 0xa7, 0x4f, 0x31, 0x05, 0xdb, 0xce, 0x28, 0x6e, 0x67, 0x24, 0x56, 0x30, 0xf4, 0x7b, 0x5f,
	0x64, 0xe6, 0x1d, 0xe4, 0x61, 0xb6, 0x5f, 0xd6, 0x72, 0xf3, 0xdc, 0x7b, 0xff, 0x01, 0xd2, 0x26,
	0x6a, 0x96, 0x41, 0xef, 0x71, 0xf9, 0x38, 0xbf, 0xe9, 0xb0, 0x01, 0xa4, 0xb3, 0xd5, 0x62, 0xf1,
	0xfd, 0xc7, 0xb7, 0x9b, 0x88, 0xc0, 0xfc, 0x8b, 0x44, 0xd0, 0x9d, 0x1e, 0x21, 0x9b, 0xb9, 0xfd,
	0x7e, 0xa6, 0x4a, 0xc3, 0xde, 0x40, 0xfc, 0x55, 0xd5, 0xec, 0x79, 0xee, 0x1d, 0x95, 0xf4, 0x3c,
	0xd1, 0x61, 0xf7, 0xd0, 0x7f, 0x30, 0xaa, 0xac, 0x15, 0xcb, 0x8a, 0x26, 0xad, 0xb6, 0x60, 0x0c,
	0x3d, 0x74, 0xc4, 0x86, 0xc5, 0x59, 0xa0, 0x77, 0x57, 0xc5, 0xb9, 0x4d, 0xd1, 0xd9, 0xf4, 0xe9,
	0xcf, 0xf9, 0xf4, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x09, 0x1b, 0xc0, 0x26, 0x51, 0x03, 0x00, 0x00,
}
