// Code generated by protoc-gen-go. DO NOT EDIT.
// source: daobackup.proto

/*
Package daobackup is a generated protocol buffer package.

It is generated from these files:
	daobackup.proto

It has these top-level messages:
	Chunk
	Status
	ChunkStatus
*/
package daobackup

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Chunk struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Chunk) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Chunk) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type Status struct {
	Ok      bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Status) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChunkStatus struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Chunk  *Chunk  `protobuf:"bytes,2,opt,name=chunk" json:"chunk,omitempty"`
}

func (m *ChunkStatus) Reset()                    { *m = ChunkStatus{} }
func (m *ChunkStatus) String() string            { return proto.CompactTextString(m) }
func (*ChunkStatus) ProtoMessage()               {}
func (*ChunkStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChunkStatus) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ChunkStatus) GetChunk() *Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*Chunk)(nil), "Chunk")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*ChunkStatus)(nil), "ChunkStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DOABackup service

type DOABackupClient interface {
	PutChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Status, error)
	GetChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*ChunkStatus, error)
}

type dOABackupClient struct {
	cc *grpc.ClientConn
}

func NewDOABackupClient(cc *grpc.ClientConn) DOABackupClient {
	return &dOABackupClient{cc}
}

func (c *dOABackupClient) PutChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/DOABackup/PutChunk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dOABackupClient) GetChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*ChunkStatus, error) {
	out := new(ChunkStatus)
	err := grpc.Invoke(ctx, "/DOABackup/GetChunk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DOABackup service

type DOABackupServer interface {
	PutChunk(context.Context, *Chunk) (*Status, error)
	GetChunk(context.Context, *Chunk) (*ChunkStatus, error)
}

func RegisterDOABackupServer(s *grpc.Server, srv DOABackupServer) {
	s.RegisterService(&_DOABackup_serviceDesc, srv)
}

func _DOABackup_PutChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DOABackupServer).PutChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DOABackup/PutChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DOABackupServer).PutChunk(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _DOABackup_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DOABackupServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DOABackup/GetChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DOABackupServer).GetChunk(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

var _DOABackup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DOABackup",
	HandlerType: (*DOABackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutChunk",
			Handler:    _DOABackup_PutChunk_Handler,
		},
		{
			MethodName: "GetChunk",
			Handler:    _DOABackup_GetChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daobackup.proto",
}

func init() { proto.RegisterFile("daobackup.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0xcc, 0x4f,
	0x4a, 0x4c, 0xce, 0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe7, 0x62, 0x75,
	0xce, 0x28, 0xcd, 0xcb, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x49, 0x39, 0xf9, 0x49, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x3c, 0x41, 0x60, 0xb6, 0x92, 0x11, 0x17, 0x5b, 0x70, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x10, 0x1f,
	0x17, 0x53, 0x7e, 0x36, 0x58, 0x3d, 0x47, 0x10, 0x53, 0x7e, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x6e,
	0x6a, 0x71, 0x71, 0x62, 0x7a, 0x2a, 0x58, 0x03, 0x67, 0x10, 0x8c, 0xab, 0xe4, 0xc3, 0xc5, 0x0d,
	0xb6, 0x04, 0xaa, 0x51, 0x9e, 0x8b, 0xad, 0x18, 0xcc, 0x02, 0x6b, 0xe6, 0x36, 0x62, 0xd7, 0x83,
	0x48, 0x04, 0x41, 0x85, 0x85, 0x64, 0xb8, 0x58, 0x93, 0x41, 0xea, 0xc1, 0xe6, 0x70, 0x1b, 0xb1,
	0xe9, 0x81, 0x75, 0x07, 0x41, 0x04, 0x8d, 0xfc, 0xb8, 0x38, 0x5d, 0xfc, 0x1d, 0x9d, 0xc0, 0xbe,
	0x10, 0x92, 0xe5, 0xe2, 0x08, 0x28, 0x2d, 0x81, 0x78, 0x01, 0xaa, 0x4e, 0x0a, 0x66, 0x9e, 0x12,
	0x83, 0x90, 0x12, 0x17, 0x87, 0x7b, 0x2a, 0x9a, 0x34, 0x8f, 0x1e, 0x92, 0x63, 0x94, 0x18, 0x92,
	0xd8, 0xc0, 0x21, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x11, 0xf0, 0x4b, 0x1c, 0x01,
	0x00, 0x00,
}