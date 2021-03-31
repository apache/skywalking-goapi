//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service-mesh-probe/service-mesh.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking/network/common/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Protocol int32

const (
	Protocol_HTTP Protocol = 0
	Protocol_gRPC Protocol = 1
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "HTTP",
		1: "gRPC",
	}
	Protocol_value = map[string]int32{
		"HTTP": 0,
		"gRPC": 1,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_service_mesh_probe_service_mesh_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_service_mesh_probe_service_mesh_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_service_mesh_probe_service_mesh_proto_rawDescGZIP(), []int{0}
}

type ServiceMeshMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start timestamp in milliseconds of this RPC,
	// measured between the current time and midnight, January 1, 1970 UTC.
	StartTime int64 `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// End timestamp in milliseconds of this RPC,
	// measured between the current time and midnight, January 1, 1970 UTC.
	EndTime               int64  `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	SourceServiceName     string `protobuf:"bytes,3,opt,name=sourceServiceName,proto3" json:"sourceServiceName,omitempty"`
	SourceServiceInstance string `protobuf:"bytes,4,opt,name=sourceServiceInstance,proto3" json:"sourceServiceInstance,omitempty"`
	DestServiceName       string `protobuf:"bytes,5,opt,name=destServiceName,proto3" json:"destServiceName,omitempty"`
	DestServiceInstance   string `protobuf:"bytes,6,opt,name=destServiceInstance,proto3" json:"destServiceInstance,omitempty"`
	Endpoint              string `protobuf:"bytes,7,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Latency               int32  `protobuf:"varint,8,opt,name=latency,proto3" json:"latency,omitempty"`
	ResponseCode          int32  `protobuf:"varint,9,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
	// Status represents the response status of this calling.
	Status      bool           `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Protocol    Protocol       `protobuf:"varint,11,opt,name=protocol,proto3,enum=skywalking.v3.Protocol" json:"protocol,omitempty"`
	DetectPoint v3.DetectPoint `protobuf:"varint,12,opt,name=detectPoint,proto3,enum=skywalking.v3.DetectPoint" json:"detectPoint,omitempty"`
	// NONE, mTLS, or TLS
	TlsMode string `protobuf:"bytes,13,opt,name=tlsMode,proto3" json:"tlsMode,omitempty"`
	// The sidecar/proxy internal error code, the value bases on the implementation.
	// The envoy internal error codes are listed here, https://www.envoyproxy.io/docs/envoy/latest/api-v2/data/accesslog/v2/accesslog.proto#data-accesslog-v2-responseflags
	InternalErrorCode string `protobuf:"bytes,14,opt,name=internalErrorCode,proto3" json:"internalErrorCode,omitempty"`
}

func (x *ServiceMeshMetric) Reset() {
	*x = ServiceMeshMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mesh_probe_service_mesh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMeshMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMeshMetric) ProtoMessage() {}

func (x *ServiceMeshMetric) ProtoReflect() protoreflect.Message {
	mi := &file_service_mesh_probe_service_mesh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMeshMetric.ProtoReflect.Descriptor instead.
func (*ServiceMeshMetric) Descriptor() ([]byte, []int) {
	return file_service_mesh_probe_service_mesh_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceMeshMetric) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ServiceMeshMetric) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ServiceMeshMetric) GetSourceServiceName() string {
	if x != nil {
		return x.SourceServiceName
	}
	return ""
}

func (x *ServiceMeshMetric) GetSourceServiceInstance() string {
	if x != nil {
		return x.SourceServiceInstance
	}
	return ""
}

func (x *ServiceMeshMetric) GetDestServiceName() string {
	if x != nil {
		return x.DestServiceName
	}
	return ""
}

func (x *ServiceMeshMetric) GetDestServiceInstance() string {
	if x != nil {
		return x.DestServiceInstance
	}
	return ""
}

func (x *ServiceMeshMetric) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ServiceMeshMetric) GetLatency() int32 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *ServiceMeshMetric) GetResponseCode() int32 {
	if x != nil {
		return x.ResponseCode
	}
	return 0
}

func (x *ServiceMeshMetric) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ServiceMeshMetric) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_HTTP
}

func (x *ServiceMeshMetric) GetDetectPoint() v3.DetectPoint {
	if x != nil {
		return x.DetectPoint
	}
	return v3.DetectPoint_client
}

func (x *ServiceMeshMetric) GetTlsMode() string {
	if x != nil {
		return x.TlsMode
	}
	return ""
}

func (x *ServiceMeshMetric) GetInternalErrorCode() string {
	if x != nil {
		return x.InternalErrorCode
	}
	return ""
}

type MeshProbeDownstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MeshProbeDownstream) Reset() {
	*x = MeshProbeDownstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mesh_probe_service_mesh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshProbeDownstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshProbeDownstream) ProtoMessage() {}

func (x *MeshProbeDownstream) ProtoReflect() protoreflect.Message {
	mi := &file_service_mesh_probe_service_mesh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshProbeDownstream.ProtoReflect.Descriptor instead.
func (*MeshProbeDownstream) Descriptor() ([]byte, []int) {
	return file_service_mesh_probe_service_mesh_proto_rawDescGZIP(), []int{1}
}

var File_service_mesh_probe_service_mesh_proto protoreflect.FileDescriptor

var file_service_mesh_probe_service_mesh_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x04, 0x0a, 0x11,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x0b, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2a, 0x1e, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x67, 0x52, 0x50, 0x43, 0x10, 0x01, 0x32, 0x6f, 0x0a,
	0x18, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x28, 0x01, 0x42, 0x77,
	0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x33, 0x50, 0x01, 0x5a, 0x21, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x33, 0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_mesh_probe_service_mesh_proto_rawDescOnce sync.Once
	file_service_mesh_probe_service_mesh_proto_rawDescData = file_service_mesh_probe_service_mesh_proto_rawDesc
)

func file_service_mesh_probe_service_mesh_proto_rawDescGZIP() []byte {
	file_service_mesh_probe_service_mesh_proto_rawDescOnce.Do(func() {
		file_service_mesh_probe_service_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_mesh_probe_service_mesh_proto_rawDescData)
	})
	return file_service_mesh_probe_service_mesh_proto_rawDescData
}

var file_service_mesh_probe_service_mesh_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_mesh_probe_service_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_mesh_probe_service_mesh_proto_goTypes = []interface{}{
	(Protocol)(0),               // 0: skywalking.v3.Protocol
	(*ServiceMeshMetric)(nil),   // 1: skywalking.v3.ServiceMeshMetric
	(*MeshProbeDownstream)(nil), // 2: skywalking.v3.MeshProbeDownstream
	(v3.DetectPoint)(0),         // 3: skywalking.v3.DetectPoint
}
var file_service_mesh_probe_service_mesh_proto_depIdxs = []int32{
	0, // 0: skywalking.v3.ServiceMeshMetric.protocol:type_name -> skywalking.v3.Protocol
	3, // 1: skywalking.v3.ServiceMeshMetric.detectPoint:type_name -> skywalking.v3.DetectPoint
	1, // 2: skywalking.v3.ServiceMeshMetricService.collect:input_type -> skywalking.v3.ServiceMeshMetric
	2, // 3: skywalking.v3.ServiceMeshMetricService.collect:output_type -> skywalking.v3.MeshProbeDownstream
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_mesh_probe_service_mesh_proto_init() }
func file_service_mesh_probe_service_mesh_proto_init() {
	if File_service_mesh_probe_service_mesh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_mesh_probe_service_mesh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMeshMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_mesh_probe_service_mesh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshProbeDownstream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_mesh_probe_service_mesh_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mesh_probe_service_mesh_proto_goTypes,
		DependencyIndexes: file_service_mesh_probe_service_mesh_proto_depIdxs,
		EnumInfos:         file_service_mesh_probe_service_mesh_proto_enumTypes,
		MessageInfos:      file_service_mesh_probe_service_mesh_proto_msgTypes,
	}.Build()
	File_service_mesh_probe_service_mesh_proto = out.File
	file_service_mesh_probe_service_mesh_proto_rawDesc = nil
	file_service_mesh_probe_service_mesh_proto_goTypes = nil
	file_service_mesh_probe_service_mesh_proto_depIdxs = nil
}
