// Code generated by protoc-gen-go.
// source: registry.proto
// DO NOT EDIT!

/*
Package regproto is a generated protocol buffer package.

It is generated from these files:
	registry.proto

It has these top-level messages:
	Chunk
	Received
*/
package regproto

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

// Chunk of a file
type Chunk struct {
	// TODO AuthToken auth = 1;
	FileType string `protobuf:"bytes,1,opt,name=file_type,json=fileType" json:"file_type,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Acknowledgement - maybe stream this for partial upload recovery
// Do we need this or can we just use gRPC errors?
type Received struct {
	Received bool `protobuf:"varint,1,opt,name=received" json:"received,omitempty"`
}

func (m *Received) Reset()                    { *m = Received{} }
func (m *Received) String() string            { return proto.CompactTextString(m) }
func (*Received) ProtoMessage()               {}
func (*Received) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Chunk)(nil), "regproto.Chunk")
	proto.RegisterType((*Received)(nil), "regproto.Received")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Registry service

type RegistryClient interface {
	Push(ctx context.Context, opts ...grpc.CallOption) (Registry_PushClient, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Push(ctx context.Context, opts ...grpc.CallOption) (Registry_PushClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Registry_serviceDesc.Streams[0], c.cc, "/regproto.Registry/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryPushClient{stream}
	return x, nil
}

type Registry_PushClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Received, error)
	grpc.ClientStream
}

type registryPushClient struct {
	grpc.ClientStream
}

func (x *registryPushClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registryPushClient) CloseAndRecv() (*Received, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Received)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Registry service

type RegistryServer interface {
	Push(Registry_PushServer) error
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Push(&registryPushServer{stream})
}

type Registry_PushServer interface {
	SendAndClose(*Received) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type registryPushServer struct {
	grpc.ServerStream
}

func (x *registryPushServer) SendAndClose(m *Received) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registryPushServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regproto.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Push",
			Handler:       _Registry_Push_Handler,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xf2, 0xc1, 0x2c,
	0x25, 0x1f, 0x2e, 0x56, 0xe7, 0x8c, 0xd2, 0xbc, 0x6c, 0x21, 0x69, 0x2e, 0xce, 0xb4, 0xcc, 0x9c,
	0xd4, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x0e, 0x90, 0x40,
	0x08, 0x90, 0x2f, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x07, 0xb3,
	0x41, 0x62, 0x29, 0x89, 0x25, 0x89, 0x12, 0xcc, 0x40, 0x31, 0x9e, 0x20, 0x30, 0x5b, 0x49, 0x8d,
	0x8b, 0x23, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0x35, 0x45, 0x48, 0x8a, 0x0b, 0x68, 0x0b, 0x84,
	0x0d, 0x36, 0x8f, 0x23, 0x08, 0xce, 0x37, 0xb2, 0x06, 0xa9, 0x83, 0xb8, 0x48, 0x48, 0x9f, 0x8b,
	0x25, 0xa0, 0xb4, 0x38, 0x43, 0x88, 0x5f, 0x0f, 0xe6, 0x28, 0x3d, 0xb0, 0x8b, 0xa4, 0x84, 0x10,
	0x02, 0x30, 0x43, 0x95, 0x18, 0x34, 0x18, 0x93, 0xd8, 0xc0, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x52, 0xc0, 0xe8, 0x84, 0xd5, 0x00, 0x00, 0x00,
}
