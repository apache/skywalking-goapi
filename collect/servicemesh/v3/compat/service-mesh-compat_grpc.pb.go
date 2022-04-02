// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package compat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking.apache.org/repo/goapi/collect/servicemesh/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceMeshMetricServiceClient is the client API for ServiceMeshMetricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceMeshMetricServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (ServiceMeshMetricService_CollectClient, error)
}

type serviceMeshMetricServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceMeshMetricServiceClient(cc grpc.ClientConnInterface) ServiceMeshMetricServiceClient {
	return &serviceMeshMetricServiceClient{cc}
}

func (c *serviceMeshMetricServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (ServiceMeshMetricService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceMeshMetricService_ServiceDesc.Streams[0], "/ServiceMeshMetricService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMeshMetricServiceCollectClient{stream}
	return x, nil
}

type ServiceMeshMetricService_CollectClient interface {
	Send(*v3.ServiceMeshMetric) error
	CloseAndRecv() (*v3.MeshProbeDownstream, error)
	grpc.ClientStream
}

type serviceMeshMetricServiceCollectClient struct {
	grpc.ClientStream
}

func (x *serviceMeshMetricServiceCollectClient) Send(m *v3.ServiceMeshMetric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceMeshMetricServiceCollectClient) CloseAndRecv() (*v3.MeshProbeDownstream, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(v3.MeshProbeDownstream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceMeshMetricServiceServer is the server API for ServiceMeshMetricService service.
// All implementations must embed UnimplementedServiceMeshMetricServiceServer
// for forward compatibility
type ServiceMeshMetricServiceServer interface {
	Collect(ServiceMeshMetricService_CollectServer) error
	mustEmbedUnimplementedServiceMeshMetricServiceServer()
}

// UnimplementedServiceMeshMetricServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceMeshMetricServiceServer struct {
}

func (UnimplementedServiceMeshMetricServiceServer) Collect(ServiceMeshMetricService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedServiceMeshMetricServiceServer) mustEmbedUnimplementedServiceMeshMetricServiceServer() {
}

// UnsafeServiceMeshMetricServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceMeshMetricServiceServer will
// result in compilation errors.
type UnsafeServiceMeshMetricServiceServer interface {
	mustEmbedUnimplementedServiceMeshMetricServiceServer()
}

func RegisterServiceMeshMetricServiceServer(s grpc.ServiceRegistrar, srv ServiceMeshMetricServiceServer) {
	s.RegisterService(&ServiceMeshMetricService_ServiceDesc, srv)
}

func _ServiceMeshMetricService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceMeshMetricServiceServer).Collect(&serviceMeshMetricServiceCollectServer{stream})
}

type ServiceMeshMetricService_CollectServer interface {
	SendAndClose(*v3.MeshProbeDownstream) error
	Recv() (*v3.ServiceMeshMetric, error)
	grpc.ServerStream
}

type serviceMeshMetricServiceCollectServer struct {
	grpc.ServerStream
}

func (x *serviceMeshMetricServiceCollectServer) SendAndClose(m *v3.MeshProbeDownstream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceMeshMetricServiceCollectServer) Recv() (*v3.ServiceMeshMetric, error) {
	m := new(v3.ServiceMeshMetric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceMeshMetricService_ServiceDesc is the grpc.ServiceDesc for ServiceMeshMetricService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceMeshMetricService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceMeshMetricService",
	HandlerType: (*ServiceMeshMetricServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _ServiceMeshMetricService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service-mesh-probe/service-mesh-compat.proto",
}