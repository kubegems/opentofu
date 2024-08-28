// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.15.6
// source: cloudproto1.proto

package cloudproto1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CommandRequest is used to request the execution of a specific command with
// provided flags. It is the raw args from the terraform cloud command.
type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudproto1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudproto1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_cloudproto1_proto_rawDescGZIP(), []int{0}
}

func (x *CommandRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

// CommandResponse contains the result of the command execution, including any
// output or errors.
type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*CommandResponse_ExitCode
	//	*CommandResponse_Stdout
	//	*CommandResponse_Stderr
	Data isCommandResponse_Data `protobuf_oneof:"data"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudproto1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudproto1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_cloudproto1_proto_rawDescGZIP(), []int{1}
}

func (m *CommandResponse) GetData() isCommandResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CommandResponse) GetExitCode() int32 {
	if x, ok := x.GetData().(*CommandResponse_ExitCode); ok {
		return x.ExitCode
	}
	return 0
}

func (x *CommandResponse) GetStdout() []byte {
	if x, ok := x.GetData().(*CommandResponse_Stdout); ok {
		return x.Stdout
	}
	return nil
}

func (x *CommandResponse) GetStderr() []byte {
	if x, ok := x.GetData().(*CommandResponse_Stderr); ok {
		return x.Stderr
	}
	return nil
}

type isCommandResponse_Data interface {
	isCommandResponse_Data()
}

type CommandResponse_ExitCode struct {
	ExitCode int32 `protobuf:"varint,1,opt,name=exitCode,proto3,oneof"`
}

type CommandResponse_Stdout struct {
	Stdout []byte `protobuf:"bytes,2,opt,name=stdout,proto3,oneof"`
}

type CommandResponse_Stderr struct {
	Stderr []byte `protobuf:"bytes,3,opt,name=stderr,proto3,oneof"`
}

func (*CommandResponse_ExitCode) isCommandResponse_Data() {}

func (*CommandResponse_Stdout) isCommandResponse_Data() {}

func (*CommandResponse_Stderr) isCommandResponse_Data() {}

var File_cloudproto1_proto protoreflect.FileDescriptor

var file_cloudproto1_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x31,
	0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x6b, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x69,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x65,
	0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x5a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x6f, 0x66, 0x75, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x75, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudproto1_proto_rawDescOnce sync.Once
	file_cloudproto1_proto_rawDescData = file_cloudproto1_proto_rawDesc
)

func file_cloudproto1_proto_rawDescGZIP() []byte {
	file_cloudproto1_proto_rawDescOnce.Do(func() {
		file_cloudproto1_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudproto1_proto_rawDescData)
	})
	return file_cloudproto1_proto_rawDescData
}

var file_cloudproto1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloudproto1_proto_goTypes = []interface{}{
	(*CommandRequest)(nil),  // 0: cloudproto1.CommandRequest
	(*CommandResponse)(nil), // 1: cloudproto1.CommandResponse
}
var file_cloudproto1_proto_depIdxs = []int32{
	0, // 0: cloudproto1.CommandService.Execute:input_type -> cloudproto1.CommandRequest
	1, // 1: cloudproto1.CommandService.Execute:output_type -> cloudproto1.CommandResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloudproto1_proto_init() }
func file_cloudproto1_proto_init() {
	if File_cloudproto1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudproto1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_cloudproto1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
	file_cloudproto1_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommandResponse_ExitCode)(nil),
		(*CommandResponse_Stdout)(nil),
		(*CommandResponse_Stderr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudproto1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudproto1_proto_goTypes,
		DependencyIndexes: file_cloudproto1_proto_depIdxs,
		MessageInfos:      file_cloudproto1_proto_msgTypes,
	}.Build()
	File_cloudproto1_proto = out.File
	file_cloudproto1_proto_rawDesc = nil
	file_cloudproto1_proto_goTypes = nil
	file_cloudproto1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandServiceClient interface {
	// Execute runs a specific command with the provided flags and returns the result.
	Execute(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (CommandService_ExecuteClient, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) Execute(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (CommandService_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandService_serviceDesc.Streams[0], "/cloudproto1.CommandService/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandServiceExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommandService_ExecuteClient interface {
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type commandServiceExecuteClient struct {
	grpc.ClientStream
}

func (x *commandServiceExecuteClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServiceServer is the server API for CommandService service.
type CommandServiceServer interface {
	// Execute runs a specific command with the provided flags and returns the result.
	Execute(*CommandRequest, CommandService_ExecuteServer) error
}

// UnimplementedCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (*UnimplementedCommandServiceServer) Execute(*CommandRequest, CommandService_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterCommandServiceServer(s *grpc.Server, srv CommandServiceServer) {
	s.RegisterService(&_CommandService_serviceDesc, srv)
}

func _CommandService_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommandServiceServer).Execute(m, &commandServiceExecuteServer{stream})
}

type CommandService_ExecuteServer interface {
	Send(*CommandResponse) error
	grpc.ServerStream
}

type commandServiceExecuteServer struct {
	grpc.ServerStream
}

func (x *commandServiceExecuteServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudproto1.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _CommandService_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloudproto1.proto",
}
