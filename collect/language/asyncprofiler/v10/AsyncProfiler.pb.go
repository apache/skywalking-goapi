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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: asyncprofiler/AsyncProfiler.proto

package v10

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

type AsyncProfilerCollectType int32

const (
	// PROFILING_SUCCESS means the Java Agent has finished the execution
	AsyncProfilerCollectType_PROFILING_SUCCESS AsyncProfilerCollectType = 0
	// EXECUTION_TASK_ERROR means potential execution error caused by the Java Agent, such as an error in the sent task parameters.
	AsyncProfilerCollectType_EXECUTION_TASK_ERROR AsyncProfilerCollectType = 1
)

// Enum value maps for AsyncProfilerCollectType.
var (
	AsyncProfilerCollectType_name = map[int32]string{
		0: "PROFILING_SUCCESS",
		1: "EXECUTION_TASK_ERROR",
	}
	AsyncProfilerCollectType_value = map[string]int32{
		"PROFILING_SUCCESS":    0,
		"EXECUTION_TASK_ERROR": 1,
	}
)

func (x AsyncProfilerCollectType) Enum() *AsyncProfilerCollectType {
	p := new(AsyncProfilerCollectType)
	*p = x
	return p
}

func (x AsyncProfilerCollectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AsyncProfilerCollectType) Descriptor() protoreflect.EnumDescriptor {
	return file_asyncprofiler_AsyncProfiler_proto_enumTypes[0].Descriptor()
}

func (AsyncProfilerCollectType) Type() protoreflect.EnumType {
	return &file_asyncprofiler_AsyncProfiler_proto_enumTypes[0]
}

func (x AsyncProfilerCollectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AsyncProfilerCollectType.Descriptor instead.
func (AsyncProfilerCollectType) EnumDescriptor() ([]byte, []int) {
	return file_asyncprofiler_AsyncProfiler_proto_rawDescGZIP(), []int{0}
}

type AsyncProfilerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// metaData of the AsyncProfiler task and its result data, only sent in the first request.
	MetaData *AsyncProfilerMetaData `protobuf:"bytes,1,opt,name=metaData,proto3" json:"metaData,omitempty"`
	// Types that are assignable to Result:
	//
	//	*AsyncProfilerData_ErrorMessage
	//	*AsyncProfilerData_Content
	Result isAsyncProfilerData_Result `protobuf_oneof:"result"`
}

func (x *AsyncProfilerData) Reset() {
	*x = AsyncProfilerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncProfilerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncProfilerData) ProtoMessage() {}

func (x *AsyncProfilerData) ProtoReflect() protoreflect.Message {
	mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncProfilerData.ProtoReflect.Descriptor instead.
func (*AsyncProfilerData) Descriptor() ([]byte, []int) {
	return file_asyncprofiler_AsyncProfiler_proto_rawDescGZIP(), []int{0}
}

func (x *AsyncProfilerData) GetMetaData() *AsyncProfilerMetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (m *AsyncProfilerData) GetResult() isAsyncProfilerData_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *AsyncProfilerData) GetErrorMessage() string {
	if x, ok := x.GetResult().(*AsyncProfilerData_ErrorMessage); ok {
		return x.ErrorMessage
	}
	return ""
}

func (x *AsyncProfilerData) GetContent() []byte {
	if x, ok := x.GetResult().(*AsyncProfilerData_Content); ok {
		return x.Content
	}
	return nil
}

type isAsyncProfilerData_Result interface {
	isAsyncProfilerData_Result()
}

type AsyncProfilerData_ErrorMessage struct {
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3,oneof"`
}

type AsyncProfilerData_Content struct {
	// JFR binary content
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3,oneof"`
}

func (*AsyncProfilerData_ErrorMessage) isAsyncProfilerData_Result() {}

func (*AsyncProfilerData_Content) isAsyncProfilerData_Result() {}

type AsyncProfilerMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// async-profiler task id
	TaskId string `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// AsyncProfilerCollectType indicates the overall status of the AsyncProfiler task, i.e. success or failure
	Type AsyncProfilerCollectType `protobuf:"varint,4,opt,name=type,proto3,enum=skywalking.v10.AsyncProfilerCollectType" json:"type,omitempty"`
	// if type is success then it will be the size of the JFR file, otherwise it will be 0
	ContentSize int32 `protobuf:"varint,5,opt,name=contentSize,proto3" json:"contentSize,omitempty"`
}

func (x *AsyncProfilerMetaData) Reset() {
	*x = AsyncProfilerMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncProfilerMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncProfilerMetaData) ProtoMessage() {}

func (x *AsyncProfilerMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncProfilerMetaData.ProtoReflect.Descriptor instead.
func (*AsyncProfilerMetaData) Descriptor() ([]byte, []int) {
	return file_asyncprofiler_AsyncProfiler_proto_rawDescGZIP(), []int{1}
}

func (x *AsyncProfilerMetaData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AsyncProfilerMetaData) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *AsyncProfilerMetaData) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *AsyncProfilerMetaData) GetType() AsyncProfilerCollectType {
	if x != nil {
		return x.Type
	}
	return AsyncProfilerCollectType_PROFILING_SUCCESS
}

func (x *AsyncProfilerMetaData) GetContentSize() int32 {
	if x != nil {
		return x.ContentSize
	}
	return 0
}

type AsyncProfilerTaskCommandQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// service name of the Java process
	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// lastCommandTime is the timestamp of the last AsyncProfiler task received
	LastCommandTime int64 `protobuf:"varint,3,opt,name=lastCommandTime,proto3" json:"lastCommandTime,omitempty"`
}

func (x *AsyncProfilerTaskCommandQuery) Reset() {
	*x = AsyncProfilerTaskCommandQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncProfilerTaskCommandQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncProfilerTaskCommandQuery) ProtoMessage() {}

func (x *AsyncProfilerTaskCommandQuery) ProtoReflect() protoreflect.Message {
	mi := &file_asyncprofiler_AsyncProfiler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncProfilerTaskCommandQuery.ProtoReflect.Descriptor instead.
func (*AsyncProfilerTaskCommandQuery) Descriptor() ([]byte, []int) {
	return file_asyncprofiler_AsyncProfiler_proto_rawDescGZIP(), []int{2}
}

func (x *AsyncProfilerTaskCommandQuery) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AsyncProfilerTaskCommandQuery) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *AsyncProfilerTaskCommandQuery) GetLastCommandTime() int64 {
	if x != nil {
		return x.LastCommandTime
	}
	return 0
}

var File_asyncprofiler_AsyncProfiler_proto protoreflect.FileDescriptor

var file_asyncprofiler_AsyncProfiler_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x30, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x41, 0x73,
	0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x41, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x30, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xd3,
	0x01, 0x0a, 0x15, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x1d, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x2a, 0x4b, 0x0a, 0x18, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x32, 0xc4, 0x01, 0x0a, 0x11, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x28, 0x01,
	0x12, 0x66, 0x0a, 0x1c, 0x67, 0x65, 0x74, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x2d, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0xa6, 0x01, 0x0a, 0x3c, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x30, 0x50, 0x01, 0x5a, 0x43, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x30,
	0xaa, 0x02, 0x1e, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x31,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asyncprofiler_AsyncProfiler_proto_rawDescOnce sync.Once
	file_asyncprofiler_AsyncProfiler_proto_rawDescData = file_asyncprofiler_AsyncProfiler_proto_rawDesc
)

func file_asyncprofiler_AsyncProfiler_proto_rawDescGZIP() []byte {
	file_asyncprofiler_AsyncProfiler_proto_rawDescOnce.Do(func() {
		file_asyncprofiler_AsyncProfiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_asyncprofiler_AsyncProfiler_proto_rawDescData)
	})
	return file_asyncprofiler_AsyncProfiler_proto_rawDescData
}

var file_asyncprofiler_AsyncProfiler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_asyncprofiler_AsyncProfiler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_asyncprofiler_AsyncProfiler_proto_goTypes = []interface{}{
	(AsyncProfilerCollectType)(0),         // 0: skywalking.v10.AsyncProfilerCollectType
	(*AsyncProfilerData)(nil),             // 1: skywalking.v10.AsyncProfilerData
	(*AsyncProfilerMetaData)(nil),         // 2: skywalking.v10.AsyncProfilerMetaData
	(*AsyncProfilerTaskCommandQuery)(nil), // 3: skywalking.v10.AsyncProfilerTaskCommandQuery
	(*v3.Commands)(nil),                   // 4: skywalking.v3.Commands
}
var file_asyncprofiler_AsyncProfiler_proto_depIdxs = []int32{
	2, // 0: skywalking.v10.AsyncProfilerData.metaData:type_name -> skywalking.v10.AsyncProfilerMetaData
	0, // 1: skywalking.v10.AsyncProfilerMetaData.type:type_name -> skywalking.v10.AsyncProfilerCollectType
	1, // 2: skywalking.v10.AsyncProfilerTask.collect:input_type -> skywalking.v10.AsyncProfilerData
	3, // 3: skywalking.v10.AsyncProfilerTask.getAsyncProfilerTaskCommands:input_type -> skywalking.v10.AsyncProfilerTaskCommandQuery
	4, // 4: skywalking.v10.AsyncProfilerTask.collect:output_type -> skywalking.v3.Commands
	4, // 5: skywalking.v10.AsyncProfilerTask.getAsyncProfilerTaskCommands:output_type -> skywalking.v3.Commands
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_asyncprofiler_AsyncProfiler_proto_init() }
func file_asyncprofiler_AsyncProfiler_proto_init() {
	if File_asyncprofiler_AsyncProfiler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asyncprofiler_AsyncProfiler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncProfilerData); i {
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
		file_asyncprofiler_AsyncProfiler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncProfilerMetaData); i {
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
		file_asyncprofiler_AsyncProfiler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncProfilerTaskCommandQuery); i {
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
	file_asyncprofiler_AsyncProfiler_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AsyncProfilerData_ErrorMessage)(nil),
		(*AsyncProfilerData_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_asyncprofiler_AsyncProfiler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_asyncprofiler_AsyncProfiler_proto_goTypes,
		DependencyIndexes: file_asyncprofiler_AsyncProfiler_proto_depIdxs,
		EnumInfos:         file_asyncprofiler_AsyncProfiler_proto_enumTypes,
		MessageInfos:      file_asyncprofiler_AsyncProfiler_proto_msgTypes,
	}.Build()
	File_asyncprofiler_AsyncProfiler_proto = out.File
	file_asyncprofiler_AsyncProfiler_proto_rawDesc = nil
	file_asyncprofiler_AsyncProfiler_proto_goTypes = nil
	file_asyncprofiler_AsyncProfiler_proto_depIdxs = nil
}