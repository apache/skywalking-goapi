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
// source: ebpf/profiling/Process.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EBPFProcessReportList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*EBPFProcessProperties `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
	// An ID generated by eBPF agent, should be unique globally.
	EbpfAgentID string `protobuf:"bytes,2,opt,name=ebpfAgentID,proto3" json:"ebpfAgentID,omitempty"`
}

func (x *EBPFProcessReportList) Reset() {
	*x = EBPFProcessReportList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessReportList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessReportList) ProtoMessage() {}

func (x *EBPFProcessReportList) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessReportList.ProtoReflect.Descriptor instead.
func (*EBPFProcessReportList) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{0}
}

func (x *EBPFProcessReportList) GetProcesses() []*EBPFProcessProperties {
	if x != nil {
		return x.Processes
	}
	return nil
}

func (x *EBPFProcessReportList) GetEbpfAgentID() string {
	if x != nil {
		return x.EbpfAgentID
	}
	return ""
}

type EBPFProcessProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Process metadata
	//
	// Types that are assignable to Metadata:
	//	*EBPFProcessProperties_HostProcess
	//	*EBPFProcessProperties_K8SProcess
	Metadata isEBPFProcessProperties_Metadata `protobuf_oneof:"metadata"`
}

func (x *EBPFProcessProperties) Reset() {
	*x = EBPFProcessProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessProperties) ProtoMessage() {}

func (x *EBPFProcessProperties) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessProperties.ProtoReflect.Descriptor instead.
func (*EBPFProcessProperties) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{1}
}

func (m *EBPFProcessProperties) GetMetadata() isEBPFProcessProperties_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (x *EBPFProcessProperties) GetHostProcess() *EBPFHostProcessMetadata {
	if x, ok := x.GetMetadata().(*EBPFProcessProperties_HostProcess); ok {
		return x.HostProcess
	}
	return nil
}

func (x *EBPFProcessProperties) GetK8SProcess() *EBPFKubernetesProcessMetadata {
	if x, ok := x.GetMetadata().(*EBPFProcessProperties_K8SProcess); ok {
		return x.K8SProcess
	}
	return nil
}

type isEBPFProcessProperties_Metadata interface {
	isEBPFProcessProperties_Metadata()
}

type EBPFProcessProperties_HostProcess struct {
	HostProcess *EBPFHostProcessMetadata `protobuf:"bytes,1,opt,name=hostProcess,proto3,oneof"`
}

type EBPFProcessProperties_K8SProcess struct {
	K8SProcess *EBPFKubernetesProcessMetadata `protobuf:"bytes,2,opt,name=k8sProcess,proto3,oneof"`
}

func (*EBPFProcessProperties_HostProcess) isEBPFProcessProperties_Metadata() {}

func (*EBPFProcessProperties_K8SProcess) isEBPFProcessProperties_Metadata() {}

type EBPFHostProcessMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [required] Entity metadata
	// Must ensure that entity information is unique at the time of reporting
	Entity *EBPFProcessEntityMetadata `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// [required] Process full command line
	Cmd string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	// [required] The IP address of the host where the process resides
	HostIP string `protobuf:"bytes,3,opt,name=hostIP,proto3" json:"hostIP,omitempty"`
	// [required] The Process id of the host
	Pid int32 `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *EBPFHostProcessMetadata) Reset() {
	*x = EBPFHostProcessMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFHostProcessMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFHostProcessMetadata) ProtoMessage() {}

func (x *EBPFHostProcessMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFHostProcessMetadata.ProtoReflect.Descriptor instead.
func (*EBPFHostProcessMetadata) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{2}
}

func (x *EBPFHostProcessMetadata) GetEntity() *EBPFProcessEntityMetadata {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *EBPFHostProcessMetadata) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *EBPFHostProcessMetadata) GetHostIP() string {
	if x != nil {
		return x.HostIP
	}
	return ""
}

func (x *EBPFHostProcessMetadata) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// Process Entity metadata
type EBPFProcessEntityMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [required] Process belong layer name which define in the backend
	Layer string `protobuf:"bytes,1,opt,name=layer,proto3" json:"layer,omitempty"`
	// [required] Process belong service name
	ServiceName string `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	// [required] Process belong service instance name
	InstanceName string `protobuf:"bytes,3,opt,name=instanceName,proto3" json:"instanceName,omitempty"`
	// [required] Process name
	ProcessName string `protobuf:"bytes,4,opt,name=processName,proto3" json:"processName,omitempty"`
	// Process labels for aggregate from service
	Labels []string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *EBPFProcessEntityMetadata) Reset() {
	*x = EBPFProcessEntityMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessEntityMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessEntityMetadata) ProtoMessage() {}

func (x *EBPFProcessEntityMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessEntityMetadata.ProtoReflect.Descriptor instead.
func (*EBPFProcessEntityMetadata) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{3}
}

func (x *EBPFProcessEntityMetadata) GetLayer() string {
	if x != nil {
		return x.Layer
	}
	return ""
}

func (x *EBPFProcessEntityMetadata) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *EBPFProcessEntityMetadata) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *EBPFProcessEntityMetadata) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *EBPFProcessEntityMetadata) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// [WIP] Kubernetes process metadata
type EBPFKubernetesProcessMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EBPFKubernetesProcessMetadata) Reset() {
	*x = EBPFKubernetesProcessMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFKubernetesProcessMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFKubernetesProcessMetadata) ProtoMessage() {}

func (x *EBPFKubernetesProcessMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFKubernetesProcessMetadata.ProtoReflect.Descriptor instead.
func (*EBPFKubernetesProcessMetadata) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{4}
}

type EBPFReportProcessDownstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*EBPFProcessDownstream `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *EBPFReportProcessDownstream) Reset() {
	*x = EBPFReportProcessDownstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFReportProcessDownstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFReportProcessDownstream) ProtoMessage() {}

func (x *EBPFReportProcessDownstream) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFReportProcessDownstream.ProtoReflect.Descriptor instead.
func (*EBPFReportProcessDownstream) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{5}
}

func (x *EBPFReportProcessDownstream) GetProcesses() []*EBPFProcessDownstream {
	if x != nil {
		return x.Processes
	}
	return nil
}

type EBPFProcessDownstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generated process id
	ProcessId string `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	// Locate the process by basic information
	//
	// Types that are assignable to Process:
	//	*EBPFProcessDownstream_HostProcess
	//	*EBPFProcessDownstream_K8SProcess
	Process isEBPFProcessDownstream_Process `protobuf_oneof:"process"`
}

func (x *EBPFProcessDownstream) Reset() {
	*x = EBPFProcessDownstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessDownstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessDownstream) ProtoMessage() {}

func (x *EBPFProcessDownstream) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessDownstream.ProtoReflect.Descriptor instead.
func (*EBPFProcessDownstream) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{6}
}

func (x *EBPFProcessDownstream) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (m *EBPFProcessDownstream) GetProcess() isEBPFProcessDownstream_Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (x *EBPFProcessDownstream) GetHostProcess() *EBPFHostProcessDownstream {
	if x, ok := x.GetProcess().(*EBPFProcessDownstream_HostProcess); ok {
		return x.HostProcess
	}
	return nil
}

func (x *EBPFProcessDownstream) GetK8SProcess() *EBPFKubernetesProcessDownstream {
	if x, ok := x.GetProcess().(*EBPFProcessDownstream_K8SProcess); ok {
		return x.K8SProcess
	}
	return nil
}

type isEBPFProcessDownstream_Process interface {
	isEBPFProcessDownstream_Process()
}

type EBPFProcessDownstream_HostProcess struct {
	HostProcess *EBPFHostProcessDownstream `protobuf:"bytes,2,opt,name=hostProcess,proto3,oneof"`
}

type EBPFProcessDownstream_K8SProcess struct {
	K8SProcess *EBPFKubernetesProcessDownstream `protobuf:"bytes,3,opt,name=k8sProcess,proto3,oneof"`
}

func (*EBPFProcessDownstream_HostProcess) isEBPFProcessDownstream_Process() {}

func (*EBPFProcessDownstream_K8SProcess) isEBPFProcessDownstream_Process() {}

type EBPFHostProcessDownstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *EBPFHostProcessDownstream) Reset() {
	*x = EBPFHostProcessDownstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFHostProcessDownstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFHostProcessDownstream) ProtoMessage() {}

func (x *EBPFHostProcessDownstream) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFHostProcessDownstream.ProtoReflect.Descriptor instead.
func (*EBPFHostProcessDownstream) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{7}
}

func (x *EBPFHostProcessDownstream) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// [WIP] Kubernetes process downstream
type EBPFKubernetesProcessDownstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EBPFKubernetesProcessDownstream) Reset() {
	*x = EBPFKubernetesProcessDownstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFKubernetesProcessDownstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFKubernetesProcessDownstream) ProtoMessage() {}

func (x *EBPFKubernetesProcessDownstream) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFKubernetesProcessDownstream.ProtoReflect.Descriptor instead.
func (*EBPFKubernetesProcessDownstream) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{8}
}

type EBPFProcessPingPkgList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*EBPFProcessPingPkg `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *EBPFProcessPingPkgList) Reset() {
	*x = EBPFProcessPingPkgList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessPingPkgList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessPingPkgList) ProtoMessage() {}

func (x *EBPFProcessPingPkgList) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessPingPkgList.ProtoReflect.Descriptor instead.
func (*EBPFProcessPingPkgList) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{9}
}

func (x *EBPFProcessPingPkgList) GetProcesses() []*EBPFProcessPingPkg {
	if x != nil {
		return x.Processes
	}
	return nil
}

type EBPFProcessPingPkg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process entity
	EntityMetadata *EBPFProcessEntityMetadata `protobuf:"bytes,1,opt,name=entityMetadata,proto3" json:"entityMetadata,omitempty"`
}

func (x *EBPFProcessPingPkg) Reset() {
	*x = EBPFProcessPingPkg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ebpf_profiling_Process_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EBPFProcessPingPkg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBPFProcessPingPkg) ProtoMessage() {}

func (x *EBPFProcessPingPkg) ProtoReflect() protoreflect.Message {
	mi := &file_ebpf_profiling_Process_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBPFProcessPingPkg.ProtoReflect.Descriptor instead.
func (*EBPFProcessPingPkg) Descriptor() ([]byte, []int) {
	return file_ebpf_profiling_Process_proto_rawDescGZIP(), []int{10}
}

func (x *EBPFProcessPingPkg) GetEntityMetadata() *EBPFProcessEntityMetadata {
	if x != nil {
		return x.EntityMetadata
	}
	return nil
}

var File_ebpf_profiling_Process_proto protoreflect.FileDescriptor

var file_ebpf_profiling_Process_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x62, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x15, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45,
	0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x62, 0x70, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x62, 0x70, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0xbf, 0x01, 0x0a, 0x15, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x68,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x45, 0x42, 0x50, 0x46, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x6b, 0x38, 0x73, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x6b, 0x38, 0x73,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x17, 0x45, 0x42, 0x50, 0x46, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x40, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x50, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x50, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0xb1, 0x01,
	0x0a, 0x19, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x45, 0x42, 0x50, 0x46, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x61, 0x0a, 0x1b, 0x45, 0x42, 0x50, 0x46, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x42, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x15, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x4c, 0x0a,
	0x0b, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x0b,
	0x68, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x50, 0x0a, 0x0a, 0x6b,
	0x38, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x45, 0x42, 0x50, 0x46, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48,
	0x00, 0x52, 0x0a, 0x6b, 0x38, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x19, 0x45, 0x42, 0x50, 0x46,
	0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x1f, 0x45, 0x42, 0x50, 0x46, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x59, 0x0a, 0x16, 0x45, 0x42,
	0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6b, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6b, 0x67, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x12, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6b, 0x67, 0x12, 0x50, 0x0a, 0x0e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xca, 0x01,
	0x0a, 0x12, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42,
	0x50, 0x46, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44,
	0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x09, 0x6b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x42, 0x50, 0x46, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6b, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x42, 0x83, 0x01, 0x0a, 0x3b, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x65, 0x62, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x42, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x62, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ebpf_profiling_Process_proto_rawDescOnce sync.Once
	file_ebpf_profiling_Process_proto_rawDescData = file_ebpf_profiling_Process_proto_rawDesc
)

func file_ebpf_profiling_Process_proto_rawDescGZIP() []byte {
	file_ebpf_profiling_Process_proto_rawDescOnce.Do(func() {
		file_ebpf_profiling_Process_proto_rawDescData = protoimpl.X.CompressGZIP(file_ebpf_profiling_Process_proto_rawDescData)
	})
	return file_ebpf_profiling_Process_proto_rawDescData
}

var file_ebpf_profiling_Process_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ebpf_profiling_Process_proto_goTypes = []interface{}{
	(*EBPFProcessReportList)(nil),           // 0: skywalking.v3.EBPFProcessReportList
	(*EBPFProcessProperties)(nil),           // 1: skywalking.v3.EBPFProcessProperties
	(*EBPFHostProcessMetadata)(nil),         // 2: skywalking.v3.EBPFHostProcessMetadata
	(*EBPFProcessEntityMetadata)(nil),       // 3: skywalking.v3.EBPFProcessEntityMetadata
	(*EBPFKubernetesProcessMetadata)(nil),   // 4: skywalking.v3.EBPFKubernetesProcessMetadata
	(*EBPFReportProcessDownstream)(nil),     // 5: skywalking.v3.EBPFReportProcessDownstream
	(*EBPFProcessDownstream)(nil),           // 6: skywalking.v3.EBPFProcessDownstream
	(*EBPFHostProcessDownstream)(nil),       // 7: skywalking.v3.EBPFHostProcessDownstream
	(*EBPFKubernetesProcessDownstream)(nil), // 8: skywalking.v3.EBPFKubernetesProcessDownstream
	(*EBPFProcessPingPkgList)(nil),          // 9: skywalking.v3.EBPFProcessPingPkgList
	(*EBPFProcessPingPkg)(nil),              // 10: skywalking.v3.EBPFProcessPingPkg
	(*v3.Commands)(nil),                     // 11: skywalking.v3.Commands
}
var file_ebpf_profiling_Process_proto_depIdxs = []int32{
	1,  // 0: skywalking.v3.EBPFProcessReportList.processes:type_name -> skywalking.v3.EBPFProcessProperties
	2,  // 1: skywalking.v3.EBPFProcessProperties.hostProcess:type_name -> skywalking.v3.EBPFHostProcessMetadata
	4,  // 2: skywalking.v3.EBPFProcessProperties.k8sProcess:type_name -> skywalking.v3.EBPFKubernetesProcessMetadata
	3,  // 3: skywalking.v3.EBPFHostProcessMetadata.entity:type_name -> skywalking.v3.EBPFProcessEntityMetadata
	6,  // 4: skywalking.v3.EBPFReportProcessDownstream.processes:type_name -> skywalking.v3.EBPFProcessDownstream
	7,  // 5: skywalking.v3.EBPFProcessDownstream.hostProcess:type_name -> skywalking.v3.EBPFHostProcessDownstream
	8,  // 6: skywalking.v3.EBPFProcessDownstream.k8sProcess:type_name -> skywalking.v3.EBPFKubernetesProcessDownstream
	10, // 7: skywalking.v3.EBPFProcessPingPkgList.processes:type_name -> skywalking.v3.EBPFProcessPingPkg
	3,  // 8: skywalking.v3.EBPFProcessPingPkg.entityMetadata:type_name -> skywalking.v3.EBPFProcessEntityMetadata
	0,  // 9: skywalking.v3.EBPFProcessService.reportProcesses:input_type -> skywalking.v3.EBPFProcessReportList
	9,  // 10: skywalking.v3.EBPFProcessService.keepAlive:input_type -> skywalking.v3.EBPFProcessPingPkgList
	5,  // 11: skywalking.v3.EBPFProcessService.reportProcesses:output_type -> skywalking.v3.EBPFReportProcessDownstream
	11, // 12: skywalking.v3.EBPFProcessService.keepAlive:output_type -> skywalking.v3.Commands
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_ebpf_profiling_Process_proto_init() }
func file_ebpf_profiling_Process_proto_init() {
	if File_ebpf_profiling_Process_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ebpf_profiling_Process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessReportList); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessProperties); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFHostProcessMetadata); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessEntityMetadata); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFKubernetesProcessMetadata); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFReportProcessDownstream); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessDownstream); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFHostProcessDownstream); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFKubernetesProcessDownstream); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessPingPkgList); i {
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
		file_ebpf_profiling_Process_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EBPFProcessPingPkg); i {
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
	file_ebpf_profiling_Process_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EBPFProcessProperties_HostProcess)(nil),
		(*EBPFProcessProperties_K8SProcess)(nil),
	}
	file_ebpf_profiling_Process_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*EBPFProcessDownstream_HostProcess)(nil),
		(*EBPFProcessDownstream_K8SProcess)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ebpf_profiling_Process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ebpf_profiling_Process_proto_goTypes,
		DependencyIndexes: file_ebpf_profiling_Process_proto_depIdxs,
		MessageInfos:      file_ebpf_profiling_Process_proto_msgTypes,
	}.Build()
	File_ebpf_profiling_Process_proto = out.File
	file_ebpf_profiling_Process_proto_rawDesc = nil
	file_ebpf_profiling_Process_proto_goTypes = nil
	file_ebpf_profiling_Process_proto_depIdxs = nil
}
