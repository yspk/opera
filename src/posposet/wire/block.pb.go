// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block.proto

package wire

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	Index                uint64   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Events               [][]byte `protobuf:"bytes,2,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetEvents() [][]byte {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "wire.Block")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xc9, 0x4f,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xcf, 0x2c, 0x4a, 0x55, 0x32, 0xe5,
	0x62, 0x75, 0x02, 0x09, 0x0a, 0x89, 0x70, 0xb1, 0x7a, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x04, 0x41, 0x38, 0x42, 0x62, 0x5c, 0x6c, 0xae, 0x65, 0xa9, 0x79, 0x25, 0xc5,
	0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x3c, 0x41, 0x50, 0x5e, 0x12, 0x1b, 0xd8, 0x0c, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0xb1, 0x9b, 0x13, 0x52, 0x00, 0x00, 0x00,
}