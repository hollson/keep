// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

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

// 服务端响应操作类型
type PongAction int32

const (
	// 成员上线通知(带有染色等状态数据)
	PongAction_OnlineNotice PongAction = 0
	// 成员(主动)下线通知
	PongAction_OfflineNotice PongAction = 1
	// 地图分发数据(透传) packet
	PongAction_DataDistribute PongAction = 2
	// 地图副本数据(透传)
	PongAction_DataReplica PongAction = 3
	// 请求定时快照
	PongAction_Snapshot PongAction = 4
	// 暂停编辑(通知)
	PongAction_Hangup PongAction = 5
	// 销毁项目(通知)
	PongAction_Dispose PongAction = 6
)

var PongAction_name = map[int32]string{
	0: "OnlineNotice",
	1: "OfflineNotice",
	2: "DataDistribute",
	3: "DataReplica",
	4: "Snapshot",
	5: "Hangup",
	6: "Dispose",
}

var PongAction_value = map[string]int32{
	"OnlineNotice":   0,
	"OfflineNotice":  1,
	"DataDistribute": 2,
	"DataReplica":    3,
	"Snapshot":       4,
	"Hangup":         5,
	"Dispose":        6,
}

func (x PongAction) String() string {
	return proto.EnumName(PongAction_name, int32(x))
}

func (PongAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

type PongCode int32

const (
	// 正常
	PongCode_ok PongCode = 0
	// 签名错误
	PongCode_err_sign PongCode = 1
	// 参数错误
	PongCode_err_param PongCode = 2
)

var PongCode_name = map[int32]string{
	0: "ok",
	1: "err_sign",
	2: "err_param",
}

var PongCode_value = map[string]int32{
	"ok":        0,
	"err_sign":  1,
	"err_param": 2,
}

func (x PongCode) String() string {
	return proto.EnumName(PongCode_name, int32(x))
}

func (PongCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
}

type Member struct {
	// 用户编号(OpenId)
	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	// 用户名
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 上线序列/染色(可选)
	Sequence             uint32   `protobuf:"fixed32,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// 响应
type Response struct {
	// 状态码
	Code PongCode `protobuf:"varint,1,opt,name=Code,proto3,enum=PongCode" json:"Code,omitempty"`
	// 操作类型
	Action PongAction `protobuf:"varint,2,opt,name=Action,proto3,enum=PongAction" json:"Action,omitempty"`
	// 数据内容(可选),即透传数据包或Member对象
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
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

func (m *Response) GetCode() PongCode {
	if m != nil {
		return m.Code
	}
	return PongCode_ok
}

func (m *Response) GetAction() PongAction {
	if m != nil {
		return m.Action
	}
	return PongAction_OnlineNotice
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("PongAction", PongAction_name, PongAction_value)
	proto.RegisterEnum("PongCode", PongCode_name, PongCode_value)
	proto.RegisterType((*Member)(nil), "Member")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x6f, 0xe2, 0x30,
	0x10, 0x85, 0x49, 0x80, 0x10, 0x06, 0x08, 0x5e, 0x9f, 0xd0, 0x4a, 0x2b, 0x21, 0xf6, 0x82, 0x38,
	0x04, 0x89, 0xd5, 0x9e, 0x7a, 0x2a, 0xcd, 0xa1, 0x95, 0x5a, 0xa8, 0x4c, 0xb9, 0xf4, 0x52, 0x39,
	0x61, 0xa0, 0x16, 0xc4, 0x76, 0x9d, 0xa0, 0x9c, 0xfa, 0x87, 0xfa, 0x2b, 0x2b, 0x1b, 0x68, 0x7b,
	0x7b, 0xf3, 0xbd, 0x99, 0xf7, 0x64, 0x19, 0x22, 0x83, 0x85, 0x56, 0xb2, 0xc0, 0x58, 0x1b, 0x55,
	0xaa, 0xd1, 0x3d, 0x04, 0x0f, 0x98, 0xa7, 0x68, 0x28, 0x85, 0xc6, 0x7a, 0x7d, 0x97, 0x0c, 0xbc,
	0xa1, 0x37, 0x6e, 0x33, 0xa7, 0x2d, 0x5b, 0xf0, 0x1c, 0x07, 0xfe, 0x89, 0x59, 0x4d, 0x7f, 0x43,
	0xb8, 0xc2, 0xb7, 0x23, 0xca, 0x0c, 0x07, 0xf5, 0xa1, 0x37, 0x6e, 0xb1, 0xaf, 0x79, 0x94, 0x42,
	0xc8, 0xce, 0xf9, 0xf4, 0x0f, 0x34, 0x6e, 0xd4, 0x06, 0x5d, 0x5e, 0x34, 0x6b, 0xc7, 0x8f, 0x4a,
	0xee, 0x2c, 0x60, 0x0e, 0xd3, 0xbf, 0x10, 0x5c, 0x67, 0xa5, 0x50, 0xd2, 0x85, 0x47, 0xb3, 0x8e,
	0x5b, 0x38, 0x21, 0x76, 0xb6, 0x6c, 0x7f, 0xc2, 0x4b, 0xee, 0x7a, 0xba, 0xcc, 0xe9, 0xc9, 0x3b,
	0xc0, 0xf7, 0x26, 0x25, 0xd0, 0x5d, 0xca, 0x83, 0x90, 0xb8, 0x50, 0xa5, 0xc8, 0x90, 0xd4, 0xe8,
	0x2f, 0xe8, 0x2d, 0xb7, 0xdb, 0x1f, 0xc8, 0xa3, 0x14, 0x22, 0x7b, 0x9a, 0x88, 0xa2, 0x34, 0x22,
	0x3d, 0x96, 0x48, 0x7c, 0xda, 0x87, 0x8e, 0x65, 0x0c, 0xf5, 0x41, 0x64, 0x9c, 0xd4, 0x69, 0x17,
	0xc2, 0x95, 0xe4, 0xba, 0x78, 0x55, 0x25, 0x69, 0x50, 0x80, 0xe0, 0x96, 0xcb, 0xdd, 0x51, 0x93,
	0x26, 0xed, 0x40, 0x2b, 0x11, 0x85, 0x56, 0x05, 0x92, 0x60, 0x32, 0x85, 0xf0, 0xf2, 0x12, 0x1a,
	0x80, 0xaf, 0xf6, 0xa4, 0x66, 0x4f, 0xd1, 0x98, 0x97, 0x42, 0xec, 0x24, 0xf1, 0x68, 0x0f, 0xda,
	0x76, 0xd2, 0xdc, 0xf0, 0x9c, 0xf8, 0xf3, 0xff, 0xd0, 0xcf, 0x54, 0x1e, 0x1b, 0xac, 0x94, 0x39,
	0x6c, 0xe2, 0xaa, 0xaa, 0xe6, 0xe1, 0x13, 0xf2, 0xbc, 0x52, 0x66, 0xff, 0xdc, 0x8c, 0xa7, 0x57,
	0x3a, 0xfd, 0xf0, 0x09, 0x3b, 0xdb, 0x17, 0x27, 0x0d, 0xdc, 0xff, 0xfc, 0xfb, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0x7c, 0x82, 0xee, 0xb1, 0x01, 0x00, 0x00,
}
