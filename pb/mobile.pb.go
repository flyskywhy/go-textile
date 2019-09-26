// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mobile.proto

package pb

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

type MobileEventType int32

const (
	MobileEventType_NODE_START               MobileEventType = 0
	MobileEventType_NODE_ONLINE              MobileEventType = 1
	MobileEventType_NODE_STOP                MobileEventType = 2
	MobileEventType_ACCOUNT_UPDATE           MobileEventType = 10
	MobileEventType_THREAD_UPDATE            MobileEventType = 11
	MobileEventType_NOTIFICATION             MobileEventType = 12
	MobileEventType_QUERY_RESPONSE           MobileEventType = 20
	MobileEventType_CAFE_SYNC_GROUP_UPDATE   MobileEventType = 30
	MobileEventType_CAFE_SYNC_GROUP_COMPLETE MobileEventType = 31
	MobileEventType_CAFE_SYNC_GROUP_FAILED   MobileEventType = 32
)

var MobileEventType_name = map[int32]string{
	0:  "NODE_START",
	1:  "NODE_ONLINE",
	2:  "NODE_STOP",
	10: "ACCOUNT_UPDATE",
	11: "THREAD_UPDATE",
	12: "NOTIFICATION",
	20: "QUERY_RESPONSE",
	30: "CAFE_SYNC_GROUP_UPDATE",
	31: "CAFE_SYNC_GROUP_COMPLETE",
	32: "CAFE_SYNC_GROUP_FAILED",
}

var MobileEventType_value = map[string]int32{
	"NODE_START":               0,
	"NODE_ONLINE":              1,
	"NODE_STOP":                2,
	"ACCOUNT_UPDATE":           10,
	"THREAD_UPDATE":            11,
	"NOTIFICATION":             12,
	"QUERY_RESPONSE":           20,
	"CAFE_SYNC_GROUP_UPDATE":   30,
	"CAFE_SYNC_GROUP_COMPLETE": 31,
	"CAFE_SYNC_GROUP_FAILED":   32,
}

func (x MobileEventType) String() string {
	return proto.EnumName(MobileEventType_name, int32(x))
}

func (MobileEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3486309221f3b440, []int{0}
}

type MobileQueryEvent_Type int32

const (
	MobileQueryEvent_DATA  MobileQueryEvent_Type = 0
	MobileQueryEvent_DONE  MobileQueryEvent_Type = 1
	MobileQueryEvent_ERROR MobileQueryEvent_Type = 2
)

var MobileQueryEvent_Type_name = map[int32]string{
	0: "DATA",
	1: "DONE",
	2: "ERROR",
}

var MobileQueryEvent_Type_value = map[string]int32{
	"DATA":  0,
	"DONE":  1,
	"ERROR": 2,
}

func (x MobileQueryEvent_Type) String() string {
	return proto.EnumName(MobileQueryEvent_Type_name, int32(x))
}

func (MobileQueryEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3486309221f3b440, []int{1, 0}
}

type MobileWalletAccount struct {
	Seed                 string   `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileWalletAccount) Reset()         { *m = MobileWalletAccount{} }
func (m *MobileWalletAccount) String() string { return proto.CompactTextString(m) }
func (*MobileWalletAccount) ProtoMessage()    {}
func (*MobileWalletAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3486309221f3b440, []int{0}
}

func (m *MobileWalletAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileWalletAccount.Unmarshal(m, b)
}
func (m *MobileWalletAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileWalletAccount.Marshal(b, m, deterministic)
}
func (m *MobileWalletAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileWalletAccount.Merge(m, src)
}
func (m *MobileWalletAccount) XXX_Size() int {
	return xxx_messageInfo_MobileWalletAccount.Size(m)
}
func (m *MobileWalletAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileWalletAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MobileWalletAccount proto.InternalMessageInfo

func (m *MobileWalletAccount) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *MobileWalletAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MobileQueryEvent struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 MobileQueryEvent_Type `protobuf:"varint,2,opt,name=type,proto3,enum=MobileQueryEvent_Type" json:"type,omitempty"`
	Data                 *QueryResult          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Error                *Error                `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileQueryEvent) Reset()         { *m = MobileQueryEvent{} }
func (m *MobileQueryEvent) String() string { return proto.CompactTextString(m) }
func (*MobileQueryEvent) ProtoMessage()    {}
func (*MobileQueryEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3486309221f3b440, []int{1}
}

func (m *MobileQueryEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileQueryEvent.Unmarshal(m, b)
}
func (m *MobileQueryEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileQueryEvent.Marshal(b, m, deterministic)
}
func (m *MobileQueryEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileQueryEvent.Merge(m, src)
}
func (m *MobileQueryEvent) XXX_Size() int {
	return xxx_messageInfo_MobileQueryEvent.Size(m)
}
func (m *MobileQueryEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileQueryEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MobileQueryEvent proto.InternalMessageInfo

func (m *MobileQueryEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MobileQueryEvent) GetType() MobileQueryEvent_Type {
	if m != nil {
		return m.Type
	}
	return MobileQueryEvent_DATA
}

func (m *MobileQueryEvent) GetData() *QueryResult {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MobileQueryEvent) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("MobileEventType", MobileEventType_name, MobileEventType_value)
	proto.RegisterEnum("MobileQueryEvent_Type", MobileQueryEvent_Type_name, MobileQueryEvent_Type_value)
	proto.RegisterType((*MobileWalletAccount)(nil), "MobileWalletAccount")
	proto.RegisterType((*MobileQueryEvent)(nil), "MobileQueryEvent")
}

func init() { proto.RegisterFile("mobile.proto", fileDescriptor_3486309221f3b440) }

var fileDescriptor_3486309221f3b440 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x6b, 0xcf, 0xed, 0x96, 0x93, 0x3f, 0xd5, 0x4e, 0x47, 0x31, 0xa5, 0x6c, 0x21, 0x30,
	0x28, 0xbd, 0xf0, 0x45, 0xf7, 0x04, 0x9a, 0xad, 0x6c, 0x86, 0x54, 0x72, 0x15, 0x85, 0xd1, 0xdd,
	0x18, 0xbb, 0x16, 0xc3, 0xe0, 0xd6, 0x9e, 0xad, 0x8c, 0xe5, 0xc1, 0xf6, 0x46, 0x7b, 0x90, 0x61,
	0x25, 0xbe, 0x09, 0xbb, 0x3b, 0xdf, 0xf9, 0xbe, 0xdf, 0xe7, 0x03, 0x16, 0x4c, 0x9e, 0xeb, 0xbc,
	0xac, 0x74, 0xd0, 0xb4, 0xb5, 0xa9, 0xaf, 0xc6, 0x3f, 0xb7, 0xba, 0xdd, 0x1d, 0xc4, 0xf4, 0x59,
	0x77, 0x5d, 0xf6, 0xe3, 0xe0, 0x2d, 0x42, 0xb8, 0xb8, 0xb7, 0xd9, 0x6f, 0x59, 0x55, 0x69, 0x43,
	0x9f, 0x9e, 0xea, 0xed, 0x8b, 0x41, 0x04, 0xaf, 0xd3, 0xba, 0xf0, 0x9d, 0xb9, 0x73, 0x33, 0x92,
	0x76, 0x46, 0x1f, 0x5e, 0x67, 0x45, 0xd1, 0xea, 0xae, 0xf3, 0x5d, 0xbb, 0x1e, 0xe4, 0xe2, 0x8f,
	0x03, 0x64, 0xdf, 0xf2, 0xd0, 0x7f, 0x89, 0xfd, 0xd2, 0x2f, 0x06, 0x67, 0xe0, 0x96, 0x43, 0x81,
	0x5b, 0x16, 0x78, 0x0b, 0x9e, 0xd9, 0x35, 0xda, 0xb2, 0xb3, 0xbb, 0xcb, 0xe0, 0x18, 0x08, 0xd4,
	0xae, 0xd1, 0xd2, 0x66, 0x70, 0x0e, 0x5e, 0x91, 0x99, 0xcc, 0x7f, 0x35, 0x77, 0x6e, 0xc6, 0x77,
	0x93, 0xc0, 0xa6, 0xa4, 0xee, 0xb6, 0x95, 0x91, 0xd6, 0xc1, 0x6b, 0x38, 0xd5, 0x6d, 0x5b, 0xb7,
	0xbe, 0x67, 0x23, 0x67, 0x01, 0xeb, 0x95, 0xdc, 0x2f, 0x17, 0x1f, 0xc1, 0xeb, 0xdb, 0xf0, 0x0d,
	0x78, 0x11, 0x55, 0x94, 0x9c, 0xd8, 0x49, 0x70, 0x46, 0x1c, 0x1c, 0xc1, 0x29, 0x93, 0x52, 0x48,
	0xe2, 0xde, 0xfe, 0x75, 0xe0, 0x7c, 0x7f, 0x86, 0xbd, 0xc0, 0x22, 0x33, 0x00, 0x2e, 0x22, 0x96,
	0xae, 0x15, 0x95, 0x8a, 0x9c, 0xe0, 0x39, 0x8c, 0xad, 0x16, 0x7c, 0x15, 0x5b, 0x7e, 0x0a, 0xa3,
	0x43, 0x40, 0x24, 0xc4, 0x45, 0x84, 0x19, 0x0d, 0x43, 0xb1, 0xe1, 0x2a, 0xdd, 0x24, 0x11, 0x55,
	0x8c, 0x00, 0xbe, 0x85, 0xa9, 0xfa, 0x2a, 0x19, 0x8d, 0x86, 0xd5, 0x18, 0x09, 0x4c, 0xb8, 0x50,
	0xf1, 0x32, 0x0e, 0xa9, 0x8a, 0x05, 0x27, 0x93, 0x1e, 0x7c, 0xd8, 0x30, 0xf9, 0x98, 0x4a, 0xb6,
	0x4e, 0x04, 0x5f, 0x33, 0xf2, 0x0e, 0xaf, 0xe0, 0x32, 0xa4, 0x4b, 0x96, 0xae, 0x1f, 0x79, 0x98,
	0x7e, 0x91, 0x62, 0x93, 0x0c, 0x0d, 0xef, 0xf1, 0x1a, 0xfc, 0x63, 0x2f, 0x14, 0xf7, 0xc9, 0x8a,
	0x29, 0x46, 0x3e, 0xfc, 0x8f, 0x5c, 0xd2, 0x78, 0xc5, 0x22, 0x32, 0xff, 0x7c, 0x01, 0xd3, 0xb2,
	0x0e, 0x8c, 0xfe, 0x6d, 0xec, 0x9b, 0xc8, 0xbf, 0xbb, 0x4d, 0x9e, 0x9f, 0xd9, 0xff, 0xff, 0xe9,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x14, 0x44, 0x17, 0x2b, 0x02, 0x00, 0x00,
}
