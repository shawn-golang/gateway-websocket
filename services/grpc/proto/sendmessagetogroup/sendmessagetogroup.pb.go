// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: sendmessagetogroup.proto

package proto

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

type SendMessageToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageToGroupRequest) Reset() {
	*x = SendMessageToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessagetogroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageToGroupRequest) ProtoMessage() {}

func (x *SendMessageToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessagetogroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageToGroupRequest.ProtoReflect.Descriptor instead.
func (*SendMessageToGroupRequest) Descriptor() ([]byte, []int) {
	return file_sendmessagetogroup_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageToGroupRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SendMessageToGroupRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMessageToGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Arrive int32 `protobuf:"varint,2,opt,name=arrive,proto3" json:"arrive,omitempty"`
}

func (x *SendMessageToGroupResponse) Reset() {
	*x = SendMessageToGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendmessagetogroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageToGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageToGroupResponse) ProtoMessage() {}

func (x *SendMessageToGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sendmessagetogroup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageToGroupResponse.ProtoReflect.Descriptor instead.
func (*SendMessageToGroupResponse) Descriptor() ([]byte, []int) {
	return file_sendmessagetogroup_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageToGroupResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *SendMessageToGroupResponse) GetArrive() int32 {
	if x != nil {
		return x.Arrive
	}
	return 0
}

var File_sendmessagetogroup_proto protoreflect.FileDescriptor

var file_sendmessagetogroup_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x74, 0x6f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4b, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4c,
	0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x32, 0x71, 0x0a, 0x12,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x5b, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sendmessagetogroup_proto_rawDescOnce sync.Once
	file_sendmessagetogroup_proto_rawDescData = file_sendmessagetogroup_proto_rawDesc
)

func file_sendmessagetogroup_proto_rawDescGZIP() []byte {
	file_sendmessagetogroup_proto_rawDescOnce.Do(func() {
		file_sendmessagetogroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_sendmessagetogroup_proto_rawDescData)
	})
	return file_sendmessagetogroup_proto_rawDescData
}

var file_sendmessagetogroup_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sendmessagetogroup_proto_goTypes = []interface{}{
	(*SendMessageToGroupRequest)(nil),  // 0: proto.SendMessageToGroupRequest
	(*SendMessageToGroupResponse)(nil), // 1: proto.SendMessageToGroupResponse
}
var file_sendmessagetogroup_proto_depIdxs = []int32{
	0, // 0: proto.SendMessageToGroup.SendMessageToGroup:input_type -> proto.SendMessageToGroupRequest
	1, // 1: proto.SendMessageToGroup.SendMessageToGroup:output_type -> proto.SendMessageToGroupResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sendmessagetogroup_proto_init() }
func file_sendmessagetogroup_proto_init() {
	if File_sendmessagetogroup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sendmessagetogroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageToGroupRequest); i {
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
		file_sendmessagetogroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageToGroupResponse); i {
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
			RawDescriptor: file_sendmessagetogroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sendmessagetogroup_proto_goTypes,
		DependencyIndexes: file_sendmessagetogroup_proto_depIdxs,
		MessageInfos:      file_sendmessagetogroup_proto_msgTypes,
	}.Build()
	File_sendmessagetogroup_proto = out.File
	file_sendmessagetogroup_proto_rawDesc = nil
	file_sendmessagetogroup_proto_goTypes = nil
	file_sendmessagetogroup_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendMessageToGroupClient is the client API for SendMessageToGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendMessageToGroupClient interface {
	SendMessageToGroup(ctx context.Context, in *SendMessageToGroupRequest, opts ...grpc.CallOption) (*SendMessageToGroupResponse, error)
}

type sendMessageToGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewSendMessageToGroupClient(cc grpc.ClientConnInterface) SendMessageToGroupClient {
	return &sendMessageToGroupClient{cc}
}

func (c *sendMessageToGroupClient) SendMessageToGroup(ctx context.Context, in *SendMessageToGroupRequest, opts ...grpc.CallOption) (*SendMessageToGroupResponse, error) {
	out := new(SendMessageToGroupResponse)
	err := c.cc.Invoke(ctx, "/proto.SendMessageToGroup/SendMessageToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendMessageToGroupServer is the server API for SendMessageToGroup service.
type SendMessageToGroupServer interface {
	SendMessageToGroup(context.Context, *SendMessageToGroupRequest) (*SendMessageToGroupResponse, error)
}

// UnimplementedSendMessageToGroupServer can be embedded to have forward compatible implementations.
type UnimplementedSendMessageToGroupServer struct {
}

func (*UnimplementedSendMessageToGroupServer) SendMessageToGroup(context.Context, *SendMessageToGroupRequest) (*SendMessageToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToGroup not implemented")
}

func RegisterSendMessageToGroupServer(s *grpc.Server, srv SendMessageToGroupServer) {
	s.RegisterService(&_SendMessageToGroup_serviceDesc, srv)
}

func _SendMessageToGroup_SendMessageToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendMessageToGroupServer).SendMessageToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SendMessageToGroup/SendMessageToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendMessageToGroupServer).SendMessageToGroup(ctx, req.(*SendMessageToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendMessageToGroup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendMessageToGroup",
	HandlerType: (*SendMessageToGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToGroup",
			Handler:    _SendMessageToGroup_SendMessageToGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sendmessagetogroup.proto",
}
