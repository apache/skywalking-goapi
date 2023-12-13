// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: ebpf/accesslog.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EBPFAccessLogServiceClient is the client API for EBPFAccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EBPFAccessLogServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (EBPFAccessLogService_CollectClient, error)
}

type eBPFAccessLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEBPFAccessLogServiceClient(cc grpc.ClientConnInterface) EBPFAccessLogServiceClient {
	return &eBPFAccessLogServiceClient{cc}
}

func (c *eBPFAccessLogServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (EBPFAccessLogService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &EBPFAccessLogService_ServiceDesc.Streams[0], "/skywalking.v3.EBPFAccessLogService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &eBPFAccessLogServiceCollectClient{stream}
	return x, nil
}

type EBPFAccessLogService_CollectClient interface {
	Send(*EBPFAccessLogMessage) error
	CloseAndRecv() (*EBPFAccessLogDownstream, error)
	grpc.ClientStream
}

type eBPFAccessLogServiceCollectClient struct {
	grpc.ClientStream
}

func (x *eBPFAccessLogServiceCollectClient) Send(m *EBPFAccessLogMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eBPFAccessLogServiceCollectClient) CloseAndRecv() (*EBPFAccessLogDownstream, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EBPFAccessLogDownstream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EBPFAccessLogServiceServer is the server API for EBPFAccessLogService service.
// All implementations must embed UnimplementedEBPFAccessLogServiceServer
// for forward compatibility
type EBPFAccessLogServiceServer interface {
	Collect(EBPFAccessLogService_CollectServer) error
	mustEmbedUnimplementedEBPFAccessLogServiceServer()
}

// UnimplementedEBPFAccessLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEBPFAccessLogServiceServer struct {
}

func (UnimplementedEBPFAccessLogServiceServer) Collect(EBPFAccessLogService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedEBPFAccessLogServiceServer) mustEmbedUnimplementedEBPFAccessLogServiceServer() {}

// UnsafeEBPFAccessLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EBPFAccessLogServiceServer will
// result in compilation errors.
type UnsafeEBPFAccessLogServiceServer interface {
	mustEmbedUnimplementedEBPFAccessLogServiceServer()
}

func RegisterEBPFAccessLogServiceServer(s grpc.ServiceRegistrar, srv EBPFAccessLogServiceServer) {
	s.RegisterService(&EBPFAccessLogService_ServiceDesc, srv)
}

func _EBPFAccessLogService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EBPFAccessLogServiceServer).Collect(&eBPFAccessLogServiceCollectServer{stream})
}

type EBPFAccessLogService_CollectServer interface {
	SendAndClose(*EBPFAccessLogDownstream) error
	Recv() (*EBPFAccessLogMessage, error)
	grpc.ServerStream
}

type eBPFAccessLogServiceCollectServer struct {
	grpc.ServerStream
}

func (x *eBPFAccessLogServiceCollectServer) SendAndClose(m *EBPFAccessLogDownstream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eBPFAccessLogServiceCollectServer) Recv() (*EBPFAccessLogMessage, error) {
	m := new(EBPFAccessLogMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EBPFAccessLogService_ServiceDesc is the grpc.ServiceDesc for EBPFAccessLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EBPFAccessLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.EBPFAccessLogService",
	HandlerType: (*EBPFAccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _EBPFAccessLogService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ebpf/accesslog.proto",
}
