// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Question/Answer.proto

package Data_Question

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

type Answer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_be86e061646c4a51, []int{0}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Answer) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Answer) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*Answer)(nil), "Data.Question.Answer")
}

func init() { proto.RegisterFile("Question/Answer.proto", fileDescriptor_be86e061646c4a51) }

var fileDescriptor_be86e061646c4a51 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0d, 0x2c, 0x4d, 0x2d,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x77, 0xcc, 0x2b, 0x2e, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x75, 0x49, 0x2c, 0x49, 0xd4, 0x83, 0xc9, 0x29, 0x39, 0x71, 0xb1, 0x41, 0xa4,
	0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53,
	0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xdc, 0xc4, 0xf4, 0x54, 0x09, 0x66, 0xb0, 0x20, 0x84, 0x93,
	0xc4, 0x06, 0x36, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xb0, 0xe7, 0x8f, 0x72, 0x00,
	0x00, 0x00,
}