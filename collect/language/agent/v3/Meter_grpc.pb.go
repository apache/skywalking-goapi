// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: language-agent/Meter.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeterReportServiceClient is the client API for MeterReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeterReportServiceClient interface {
	// Meter data is reported in a certain period. The agent/SDK should report all collected metrics in this period through one stream.
	// The whole stream is an input data set, client should onComplete the stream per report period.
	Collect(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectClient, error)
	// Reporting meter data in bulk mode as MeterDataCollection.
	// By using this, each one in the stream would be treated as a complete input for MAL engine,
	// comparing to `collect (stream MeterData)`, which is using one stream as an input data set.
	CollectBatch(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectBatchClient, error)
}

type meterReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeterReportServiceClient(cc grpc.ClientConnInterface) MeterReportServiceClient {
	return &meterReportServiceClient{cc}
}

func (c *meterReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &MeterReportService_ServiceDesc.Streams[0], "/skywalking.v3.MeterReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &meterReportServiceCollectClient{stream}
	return x, nil
}

type MeterReportService_CollectClient interface {
	Send(*MeterData) error
	CloseAndRecv() (*v3.Commands, error)
	grpc.ClientStream
}

type meterReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *meterReportServiceCollectClient) Send(m *MeterData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *meterReportServiceCollectClient) CloseAndRecv() (*v3.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(v3.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *meterReportServiceClient) CollectBatch(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &MeterReportService_ServiceDesc.Streams[1], "/skywalking.v3.MeterReportService/collectBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &meterReportServiceCollectBatchClient{stream}
	return x, nil
}

type MeterReportService_CollectBatchClient interface {
	Send(*MeterDataCollection) error
	CloseAndRecv() (*v3.Commands, error)
	grpc.ClientStream
}

type meterReportServiceCollectBatchClient struct {
	grpc.ClientStream
}

func (x *meterReportServiceCollectBatchClient) Send(m *MeterDataCollection) error {
	return x.ClientStream.SendMsg(m)
}

func (x *meterReportServiceCollectBatchClient) CloseAndRecv() (*v3.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(v3.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeterReportServiceServer is the server API for MeterReportService service.
// All implementations must embed UnimplementedMeterReportServiceServer
// for forward compatibility
type MeterReportServiceServer interface {
	// Meter data is reported in a certain period. The agent/SDK should report all collected metrics in this period through one stream.
	// The whole stream is an input data set, client should onComplete the stream per report period.
	Collect(MeterReportService_CollectServer) error
	// Reporting meter data in bulk mode as MeterDataCollection.
	// By using this, each one in the stream would be treated as a complete input for MAL engine,
	// comparing to `collect (stream MeterData)`, which is using one stream as an input data set.
	CollectBatch(MeterReportService_CollectBatchServer) error
	mustEmbedUnimplementedMeterReportServiceServer()
}

// UnimplementedMeterReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeterReportServiceServer struct {
}

func (UnimplementedMeterReportServiceServer) Collect(MeterReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedMeterReportServiceServer) CollectBatch(MeterReportService_CollectBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method CollectBatch not implemented")
}
func (UnimplementedMeterReportServiceServer) mustEmbedUnimplementedMeterReportServiceServer() {}

// UnsafeMeterReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeterReportServiceServer will
// result in compilation errors.
type UnsafeMeterReportServiceServer interface {
	mustEmbedUnimplementedMeterReportServiceServer()
}

func RegisterMeterReportServiceServer(s grpc.ServiceRegistrar, srv MeterReportServiceServer) {
	s.RegisterService(&MeterReportService_ServiceDesc, srv)
}

func _MeterReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeterReportServiceServer).Collect(&meterReportServiceCollectServer{stream})
}

type MeterReportService_CollectServer interface {
	SendAndClose(*v3.Commands) error
	Recv() (*MeterData, error)
	grpc.ServerStream
}

type meterReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *meterReportServiceCollectServer) SendAndClose(m *v3.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *meterReportServiceCollectServer) Recv() (*MeterData, error) {
	m := new(MeterData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MeterReportService_CollectBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeterReportServiceServer).CollectBatch(&meterReportServiceCollectBatchServer{stream})
}

type MeterReportService_CollectBatchServer interface {
	SendAndClose(*v3.Commands) error
	Recv() (*MeterDataCollection, error)
	grpc.ServerStream
}

type meterReportServiceCollectBatchServer struct {
	grpc.ServerStream
}

func (x *meterReportServiceCollectBatchServer) SendAndClose(m *v3.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *meterReportServiceCollectBatchServer) Recv() (*MeterDataCollection, error) {
	m := new(MeterDataCollection)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeterReportService_ServiceDesc is the grpc.ServiceDesc for MeterReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeterReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.MeterReportService",
	HandlerType: (*MeterReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _MeterReportService_Collect_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "collectBatch",
			Handler:       _MeterReportService_CollectBatch_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent/Meter.proto",
}
