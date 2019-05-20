// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app.proto

package internal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ToServer struct {
	// Types that are valid to be assigned to Event:
	//	*ToServer_Tx_
	//	*ToServer_Answer_
	Event                isToServer_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToServer) Reset()         { *m = ToServer{} }
func (m *ToServer) String() string { return proto.CompactTextString(m) }
func (*ToServer) ProtoMessage()    {}
func (*ToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{0}
}

func (m *ToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer.Unmarshal(m, b)
}
func (m *ToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer.Marshal(b, m, deterministic)
}
func (m *ToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer.Merge(m, src)
}
func (m *ToServer) XXX_Size() int {
	return xxx_messageInfo_ToServer.Size(m)
}
func (m *ToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer proto.InternalMessageInfo

type isToServer_Event interface {
	isToServer_Event()
}

type ToServer_Tx_ struct {
	Tx *ToServer_Tx `protobuf:"bytes,1,opt,name=tx,proto3,oneof"`
}

type ToServer_Answer_ struct {
	Answer *ToServer_Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

func (*ToServer_Tx_) isToServer_Event() {}

func (*ToServer_Answer_) isToServer_Event() {}

func (m *ToServer) GetEvent() isToServer_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToServer) GetTx() *ToServer_Tx {
	if x, ok := m.GetEvent().(*ToServer_Tx_); ok {
		return x.Tx
	}
	return nil
}

func (m *ToServer) GetAnswer() *ToServer_Answer {
	if x, ok := m.GetEvent().(*ToServer_Answer_); ok {
		return x.Answer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ToServer) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ToServer_OneofMarshaler, _ToServer_OneofUnmarshaler, _ToServer_OneofSizer, []interface{}{
		(*ToServer_Tx_)(nil),
		(*ToServer_Answer_)(nil),
	}
}

func _ToServer_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ToServer)
	// event
	switch x := m.Event.(type) {
	case *ToServer_Tx_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tx); err != nil {
			return err
		}
	case *ToServer_Answer_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Answer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ToServer.Event has unexpected type %T", x)
	}
	return nil
}

func _ToServer_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ToServer)
	switch tag {
	case 1: // event.tx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ToServer_Tx)
		err := b.DecodeMessage(msg)
		m.Event = &ToServer_Tx_{msg}
		return true, err
	case 2: // event.answer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ToServer_Answer)
		err := b.DecodeMessage(msg)
		m.Event = &ToServer_Answer_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ToServer_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ToServer)
	// event
	switch x := m.Event.(type) {
	case *ToServer_Tx_:
		s := proto.Size(x.Tx)
		n++ // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToServer_Answer_:
		s := proto.Size(x.Answer)
		n++ // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ToServer_Tx struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToServer_Tx) Reset()         { *m = ToServer_Tx{} }
func (m *ToServer_Tx) String() string { return proto.CompactTextString(m) }
func (*ToServer_Tx) ProtoMessage()    {}
func (*ToServer_Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{0, 0}
}

func (m *ToServer_Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Tx.Unmarshal(m, b)
}
func (m *ToServer_Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Tx.Marshal(b, m, deterministic)
}
func (m *ToServer_Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Tx.Merge(m, src)
}
func (m *ToServer_Tx) XXX_Size() int {
	return xxx_messageInfo_ToServer_Tx.Size(m)
}
func (m *ToServer_Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Tx proto.InternalMessageInfo

func (m *ToServer_Tx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToServer_Answer struct {
	Uid []byte `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ToServer_Answer_Data
	//	*ToServer_Answer_Error
	Payload              isToServer_Answer_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ToServer_Answer) Reset()         { *m = ToServer_Answer{} }
func (m *ToServer_Answer) String() string { return proto.CompactTextString(m) }
func (*ToServer_Answer) ProtoMessage()    {}
func (*ToServer_Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{0, 1}
}

func (m *ToServer_Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Answer.Unmarshal(m, b)
}
func (m *ToServer_Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Answer.Marshal(b, m, deterministic)
}
func (m *ToServer_Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Answer.Merge(m, src)
}
func (m *ToServer_Answer) XXX_Size() int {
	return xxx_messageInfo_ToServer_Answer.Size(m)
}
func (m *ToServer_Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Answer proto.InternalMessageInfo

func (m *ToServer_Answer) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

type isToServer_Answer_Payload interface {
	isToServer_Answer_Payload()
}

type ToServer_Answer_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type ToServer_Answer_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ToServer_Answer_Data) isToServer_Answer_Payload() {}

func (*ToServer_Answer_Error) isToServer_Answer_Payload() {}

func (m *ToServer_Answer) GetPayload() isToServer_Answer_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ToServer_Answer) GetData() []byte {
	if x, ok := m.GetPayload().(*ToServer_Answer_Data); ok {
		return x.Data
	}
	return nil
}

func (m *ToServer_Answer) GetError() string {
	if x, ok := m.GetPayload().(*ToServer_Answer_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ToServer_Answer) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ToServer_Answer_OneofMarshaler, _ToServer_Answer_OneofUnmarshaler, _ToServer_Answer_OneofSizer, []interface{}{
		(*ToServer_Answer_Data)(nil),
		(*ToServer_Answer_Error)(nil),
	}
}

func _ToServer_Answer_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ToServer_Answer)
	// payload
	switch x := m.Payload.(type) {
	case *ToServer_Answer_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Data)
	case *ToServer_Answer_Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Error)
	case nil:
	default:
		return fmt.Errorf("ToServer_Answer.Payload has unexpected type %T", x)
	}
	return nil
}

func _ToServer_Answer_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ToServer_Answer)
	switch tag {
	case 2: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &ToServer_Answer_Data{x}
		return true, err
	case 3: // payload.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &ToServer_Answer_Error{x}
		return true, err
	default:
		return false, nil
	}
}

func _ToServer_Answer_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ToServer_Answer)
	// payload
	switch x := m.Payload.(type) {
	case *ToServer_Answer_Data:
		n++ // tag and wire
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *ToServer_Answer_Error:
		n++ // tag and wire
		n += proto.SizeVarint(uint64(len(x.Error)))
		n += len(x.Error)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ToClient struct {
	// Types that are valid to be assigned to Event:
	//	*ToClient_Block_
	//	*ToClient_Query_
	//	*ToClient_Restore_
	Event                isToClient_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToClient) Reset()         { *m = ToClient{} }
func (m *ToClient) String() string { return proto.CompactTextString(m) }
func (*ToClient) ProtoMessage()    {}
func (*ToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{1}
}

func (m *ToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient.Unmarshal(m, b)
}
func (m *ToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient.Marshal(b, m, deterministic)
}
func (m *ToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient.Merge(m, src)
}
func (m *ToClient) XXX_Size() int {
	return xxx_messageInfo_ToClient.Size(m)
}
func (m *ToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient proto.InternalMessageInfo

type isToClient_Event interface {
	isToClient_Event()
}

type ToClient_Block_ struct {
	Block *ToClient_Block `protobuf:"bytes,1,opt,name=block,proto3,oneof"`
}

type ToClient_Query_ struct {
	Query *ToClient_Query `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

type ToClient_Restore_ struct {
	Restore *ToClient_Restore `protobuf:"bytes,3,opt,name=restore,proto3,oneof"`
}

func (*ToClient_Block_) isToClient_Event() {}

func (*ToClient_Query_) isToClient_Event() {}

func (*ToClient_Restore_) isToClient_Event() {}

func (m *ToClient) GetEvent() isToClient_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToClient) GetBlock() *ToClient_Block {
	if x, ok := m.GetEvent().(*ToClient_Block_); ok {
		return x.Block
	}
	return nil
}

func (m *ToClient) GetQuery() *ToClient_Query {
	if x, ok := m.GetEvent().(*ToClient_Query_); ok {
		return x.Query
	}
	return nil
}

func (m *ToClient) GetRestore() *ToClient_Restore {
	if x, ok := m.GetEvent().(*ToClient_Restore_); ok {
		return x.Restore
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ToClient) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ToClient_OneofMarshaler, _ToClient_OneofUnmarshaler, _ToClient_OneofSizer, []interface{}{
		(*ToClient_Block_)(nil),
		(*ToClient_Query_)(nil),
		(*ToClient_Restore_)(nil),
	}
}

func _ToClient_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ToClient)
	// event
	switch x := m.Event.(type) {
	case *ToClient_Block_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *ToClient_Query_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Query); err != nil {
			return err
		}
	case *ToClient_Restore_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Restore); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ToClient.Event has unexpected type %T", x)
	}
	return nil
}

func _ToClient_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ToClient)
	switch tag {
	case 1: // event.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ToClient_Block)
		err := b.DecodeMessage(msg)
		m.Event = &ToClient_Block_{msg}
		return true, err
	case 2: // event.query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ToClient_Query)
		err := b.DecodeMessage(msg)
		m.Event = &ToClient_Query_{msg}
		return true, err
	case 3: // event.restore
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ToClient_Restore)
		err := b.DecodeMessage(msg)
		m.Event = &ToClient_Restore_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ToClient_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ToClient)
	// event
	switch x := m.Event.(type) {
	case *ToClient_Block_:
		s := proto.Size(x.Block)
		n++ // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_Query_:
		s := proto.Size(x.Query)
		n++ // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_Restore_:
		s := proto.Size(x.Restore)
		n++ // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ToClient_Block struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Block) Reset()         { *m = ToClient_Block{} }
func (m *ToClient_Block) String() string { return proto.CompactTextString(m) }
func (*ToClient_Block) ProtoMessage()    {}
func (*ToClient_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{1, 0}
}

func (m *ToClient_Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Block.Unmarshal(m, b)
}
func (m *ToClient_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Block.Marshal(b, m, deterministic)
}
func (m *ToClient_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Block.Merge(m, src)
}
func (m *ToClient_Block) XXX_Size() int {
	return xxx_messageInfo_ToClient_Block.Size(m)
}
func (m *ToClient_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Block.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Block proto.InternalMessageInfo

func (m *ToClient_Block) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToClient_Query struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Index                int64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Query) Reset()         { *m = ToClient_Query{} }
func (m *ToClient_Query) String() string { return proto.CompactTextString(m) }
func (*ToClient_Query) ProtoMessage()    {}
func (*ToClient_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{1, 1}
}

func (m *ToClient_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Query.Unmarshal(m, b)
}
func (m *ToClient_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Query.Marshal(b, m, deterministic)
}
func (m *ToClient_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Query.Merge(m, src)
}
func (m *ToClient_Query) XXX_Size() int {
	return xxx_messageInfo_ToClient_Query.Size(m)
}
func (m *ToClient_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Query.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Query proto.InternalMessageInfo

func (m *ToClient_Query) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Query) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ToClient_Restore struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Restore) Reset()         { *m = ToClient_Restore{} }
func (m *ToClient_Restore) String() string { return proto.CompactTextString(m) }
func (*ToClient_Restore) ProtoMessage()    {}
func (*ToClient_Restore) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf5b050d57d8bb, []int{1, 2}
}

func (m *ToClient_Restore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Restore.Unmarshal(m, b)
}
func (m *ToClient_Restore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Restore.Marshal(b, m, deterministic)
}
func (m *ToClient_Restore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Restore.Merge(m, src)
}
func (m *ToClient_Restore) XXX_Size() int {
	return xxx_messageInfo_ToClient_Restore.Size(m)
}
func (m *ToClient_Restore) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Restore.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Restore proto.InternalMessageInfo

func (m *ToClient_Restore) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Restore) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ToServer)(nil), "internal.ToServer")
	proto.RegisterType((*ToServer_Tx)(nil), "internal.ToServer.Tx")
	proto.RegisterType((*ToServer_Answer)(nil), "internal.ToServer.Answer")
	proto.RegisterType((*ToClient)(nil), "internal.ToClient")
	proto.RegisterType((*ToClient_Block)(nil), "internal.ToClient.Block")
	proto.RegisterType((*ToClient_Query)(nil), "internal.ToClient.Query")
	proto.RegisterType((*ToClient_Restore)(nil), "internal.ToClient.Restore")
}

func init() { proto.RegisterFile("internal/app.proto", fileDescriptor_bcdf5b050d57d8bb) }

var fileDescriptor_bcdf5b050d57d8bb = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xaa, 0x40,
	0x14, 0x86, 0x01, 0x2f, 0xa2, 0x47, 0x17, 0x37, 0x27, 0xf7, 0x36, 0x94, 0x95, 0x71, 0x53, 0x37,
	0x45, 0xa3, 0x49, 0xbb, 0xae, 0x6e, 0x58, 0x74, 0x53, 0xf4, 0x05, 0x46, 0x39, 0x49, 0x49, 0xc9,
	0x0c, 0x1d, 0x46, 0x8b, 0x6f, 0xd7, 0x47, 0xe9, 0xa3, 0x34, 0x33, 0x0c, 0x8d, 0x89, 0x2c, 0xba,
	0x9b, 0xe1, 0x7c, 0x3f, 0xff, 0xf9, 0x7f, 0x00, 0xcc, 0xb9, 0x22, 0xc9, 0x59, 0x31, 0x67, 0x65,
	0x19, 0x97, 0x52, 0x28, 0x81, 0x83, 0xf6, 0xd9, 0xf4, 0xcb, 0x85, 0xc1, 0x4e, 0x6c, 0x49, 0x9e,
	0x48, 0xe2, 0x1d, 0x78, 0xaa, 0x0e, 0xdd, 0x89, 0x3b, 0x1b, 0x2d, 0xff, 0xc7, 0x2d, 0x13, 0xb7,
	0xf3, 0x78, 0x57, 0x27, 0x4e, 0xea, 0xa9, 0x1a, 0x57, 0xd0, 0x67, 0xbc, 0xfa, 0x20, 0x19, 0x7a,
	0x06, 0xbe, 0xed, 0x80, 0x9f, 0x0c, 0x90, 0x38, 0xa9, 0x45, 0xa3, 0x10, 0xbc, 0x5d, 0x8d, 0x08,
	0x7f, 0x32, 0xa6, 0x98, 0x71, 0x19, 0xa7, 0xe6, 0x1c, 0x6d, 0xa1, 0xdf, 0xd0, 0xf8, 0x17, 0x7a,
	0xc7, 0x3c, 0xb3, 0x43, 0x7d, 0xc4, 0x7f, 0x96, 0xd7, 0x46, 0xe3, 0xc4, 0x69, 0x14, 0x78, 0x03,
	0x3e, 0x49, 0x29, 0x64, 0xd8, 0x9b, 0xb8, 0xb3, 0x61, 0xe2, 0xa4, 0xcd, 0x75, 0x3d, 0x84, 0xa0,
	0x64, 0xe7, 0x42, 0xb0, 0x6c, 0x1d, 0x80, 0x4f, 0x27, 0xe2, 0x6a, 0xfa, 0xe9, 0xe9, 0x88, 0x9b,
	0x22, 0x27, 0xae, 0x70, 0x01, 0xfe, 0xbe, 0x10, 0x87, 0x37, 0x9b, 0x32, 0xbc, 0x5c, 0xbc, 0x41,
	0xe2, 0xb5, 0x9e, 0xeb, 0x57, 0x1a, 0x50, 0x2b, 0xde, 0x8f, 0x24, 0xcf, 0x36, 0x6a, 0x97, 0xe2,
	0x45, 0xcf, 0xb5, 0xc2, 0x80, 0xf8, 0x00, 0x81, 0xa4, 0x4a, 0x09, 0x49, 0x66, 0xbd, 0xd1, 0x32,
	0xea, 0xd0, 0xa4, 0x0d, 0x91, 0x38, 0x69, 0x0b, 0x47, 0xf7, 0xe0, 0x1b, 0xef, 0x8e, 0x16, 0xf0,
	0xb2, 0x05, 0xdb, 0xda, 0x1c, 0x7c, 0x63, 0xdc, 0x59, 0x9a, 0x9f, 0xf3, 0x8c, 0x6a, 0xc3, 0xf7,
	0xd2, 0xe6, 0x12, 0xcd, 0x21, 0xb0, 0xae, 0xbf, 0x73, 0xf8, 0xa9, 0x70, 0xb9, 0x81, 0xc1, 0x33,
	0x3b, 0xbc, 0x52, 0x95, 0x57, 0xf8, 0x08, 0xc1, 0x46, 0x70, 0x4e, 0x07, 0x85, 0x78, 0xfd, 0xd9,
	0x23, 0xbc, 0xce, 0x3a, 0x75, 0x66, 0xee, 0xc2, 0xdd, 0xf7, 0xcd, 0xbf, 0xb7, 0xfa, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x96, 0x0d, 0x68, 0xb9, 0x91, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LachesisClient is the client API for Lachesis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LachesisClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Lachesis_ConnectClient, error)
}

type lachesisClient struct {
	cc *grpc.ClientConn
}

func NewLachesisClient(cc *grpc.ClientConn) LachesisClient {
	return &lachesisClient{cc}
}

func (c *lachesisClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Lachesis_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Lachesis_serviceDesc.Streams[0], "/internal.Lachesis/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &lachesisConnectClient{stream}
	return x, nil
}

type Lachesis_ConnectClient interface {
	Send(*ToServer) error
	Recv() (*ToClient, error)
	grpc.ClientStream
}

type lachesisConnectClient struct {
	grpc.ClientStream
}

func (x *lachesisConnectClient) Send(m *ToServer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lachesisConnectClient) Recv() (*ToClient, error) {
	m := new(ToClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LachesisServer is the server API for Lachesis service.
type LachesisServer interface {
	Connect(Lachesis_ConnectServer) error
}

func RegisterLachesisServer(s *grpc.Server, srv LachesisServer) {
	s.RegisterService(&_Lachesis_serviceDesc, srv)
}

func _Lachesis_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LachesisServer).Connect(&lachesisConnectServer{stream})
}

type Lachesis_ConnectServer interface {
	Send(*ToClient) error
	Recv() (*ToServer, error)
	grpc.ServerStream
}

type lachesisConnectServer struct {
	grpc.ServerStream
}

func (x *lachesisConnectServer) Send(m *ToClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lachesisConnectServer) Recv() (*ToServer, error) {
	m := new(ToServer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Lachesis_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Lachesis",
	HandlerType: (*LachesisServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Lachesis_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/app.proto",
}
