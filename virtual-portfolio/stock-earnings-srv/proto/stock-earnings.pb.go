// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/stock-earnings.proto

package stock_earnings

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

type ListRequest struct {
	StockUuids           []string `protobuf:"bytes,1,rep,name=stock_uuids,json=stockUuids,proto3" json:"stock_uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b069cf41b8d2194, []int{0}
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

func (m *ListRequest) GetStockUuids() []string {
	if m != nil {
		return m.StockUuids
	}
	return nil
}

type Earning struct {
	StockUuid            string   `protobuf:"bytes,1,opt,name=stock_uuid,json=stockUuid,proto3" json:"stock_uuid,omitempty"`
	Date                 int64    `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Earning) Reset()         { *m = Earning{} }
func (m *Earning) String() string { return proto.CompactTextString(m) }
func (*Earning) ProtoMessage()    {}
func (*Earning) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b069cf41b8d2194, []int{1}
}

func (m *Earning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Earning.Unmarshal(m, b)
}
func (m *Earning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Earning.Marshal(b, m, deterministic)
}
func (m *Earning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Earning.Merge(m, src)
}
func (m *Earning) XXX_Size() int {
	return xxx_messageInfo_Earning.Size(m)
}
func (m *Earning) XXX_DiscardUnknown() {
	xxx_messageInfo_Earning.DiscardUnknown(m)
}

var xxx_messageInfo_Earning proto.InternalMessageInfo

func (m *Earning) GetStockUuid() string {
	if m != nil {
		return m.StockUuid
	}
	return ""
}

func (m *Earning) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type ListResponse struct {
	Earnings             []*Earning `protobuf:"bytes,1,rep,name=earnings,proto3" json:"earnings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b069cf41b8d2194, []int{2}
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

func (m *ListResponse) GetEarnings() []*Earning {
	if m != nil {
		return m.Earnings
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*Earning)(nil), "Earning")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
}

func init() { proto.RegisterFile("proto/stock-earnings.proto", fileDescriptor_8b069cf41b8d2194) }

var fileDescriptor_8b069cf41b8d2194 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2e, 0xc9, 0x4f, 0xce, 0xd6, 0x4d, 0x4d, 0x2c, 0xca, 0xcb, 0xcc, 0x4b, 0x2f,
	0xd6, 0x03, 0x0b, 0x2a, 0xe9, 0x71, 0x71, 0xfb, 0x64, 0x16, 0x97, 0x04, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0xc9, 0x73, 0x71, 0x83, 0x95, 0xc5, 0x97, 0x96, 0x66, 0xa6, 0x14, 0x4b, 0x30,
	0x2a, 0x30, 0x6b, 0x70, 0x06, 0x71, 0x81, 0x85, 0x42, 0x41, 0x22, 0x4a, 0x36, 0x5c, 0xec, 0xae,
	0x10, 0x13, 0x84, 0x64, 0xb9, 0xb8, 0x10, 0x6a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38,
	0xe1, 0x4a, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98,
	0x83, 0xc0, 0x6c, 0x25, 0x13, 0x2e, 0x1e, 0x88, 0x6d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0x2a, 0x5c, 0x1c, 0x30, 0xf7, 0x80, 0xed, 0xe2, 0x36, 0xe2, 0xd0, 0x83, 0x1a, 0x1f, 0x04, 0x97,
	0x31, 0x32, 0xe3, 0xe2, 0x0d, 0x06, 0x19, 0x0b, 0x95, 0x29, 0x16, 0x52, 0xe5, 0x62, 0x01, 0x19,
	0x23, 0xc4, 0xa3, 0x87, 0xe4, 0x76, 0x29, 0x5e, 0x3d, 0x64, 0xb3, 0x95, 0x18, 0x92, 0xd8, 0xc0,
	0x5e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x16, 0x85, 0x22, 0x00, 0x01, 0x00, 0x00,
}
