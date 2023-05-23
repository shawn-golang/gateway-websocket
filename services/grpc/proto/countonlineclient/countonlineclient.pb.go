// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: countonlineclient.proto

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

type CountOnlineClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountOnlineClientRequest) Reset() {
	*x = CountOnlineClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countonlineclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountOnlineClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountOnlineClientRequest) ProtoMessage() {}

func (x *CountOnlineClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_countonlineclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountOnlineClientRequest.ProtoReflect.Descriptor instead.
func (*CountOnlineClientRequest) Descriptor() ([]byte, []int) {
	return file_countonlineclient_proto_rawDescGZIP(), []int{0}
}

type CountOnlineClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountOnlineClientResponse) Reset() {
	*x = CountOnlineClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countonlineclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountOnlineClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountOnlineClientResponse) ProtoMessage() {}

func (x *CountOnlineClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_countonlineclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountOnlineClientResponse.ProtoReflect.Descriptor instead.
func (*CountOnlineClientResponse) Descriptor() ([]byte, []int) {
	return file_countonlineclient_proto_rawDescGZIP(), []int{1}
}

func (x *CountOnlineClientResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_countonlineclient_proto protoreflect.FileDescriptor

var file_countonlineclient_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x19,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0x6d, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_countonlineclient_proto_rawDescOnce sync.Once
	file_countonlineclient_proto_rawDescData = file_countonlineclient_proto_rawDesc
)

func file_countonlineclient_proto_rawDescGZIP() []byte {
	file_countonlineclient_proto_rawDescOnce.Do(func() {
		file_countonlineclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_countonlineclient_proto_rawDescData)
	})
	return file_countonlineclient_proto_rawDescData
}

var file_countonlineclient_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_countonlineclient_proto_goTypes = []interface{}{
	(*CountOnlineClientRequest)(nil),  // 0: proto.CountOnlineClientRequest
	(*CountOnlineClientResponse)(nil), // 1: proto.CountOnlineClientResponse
}
var file_countonlineclient_proto_depIdxs = []int32{
	0, // 0: proto.CountOnlineClient.CountOnlineClient:input_type -> proto.CountOnlineClientRequest
	1, // 1: proto.CountOnlineClient.CountOnlineClient:output_type -> proto.CountOnlineClientResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_countonlineclient_proto_init() }
func file_countonlineclient_proto_init() {
	if File_countonlineclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_countonlineclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountOnlineClientRequest); i {
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
		file_countonlineclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountOnlineClientResponse); i {
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
			RawDescriptor: file_countonlineclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_countonlineclient_proto_goTypes,
		DependencyIndexes: file_countonlineclient_proto_depIdxs,
		MessageInfos:      file_countonlineclient_proto_msgTypes,
	}.Build()
	File_countonlineclient_proto = out.File
	file_countonlineclient_proto_rawDesc = nil
	file_countonlineclient_proto_goTypes = nil
	file_countonlineclient_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountOnlineClientClient is the client API for CountOnlineClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountOnlineClientClient interface {
	CountOnlineClient(ctx context.Context, in *CountOnlineClientRequest, opts ...grpc.CallOption) (*CountOnlineClientResponse, error)
}

type countOnlineClientClient struct {
	cc grpc.ClientConnInterface
}

func NewCountOnlineClientClient(cc grpc.ClientConnInterface) CountOnlineClientClient {
	return &countOnlineClientClient{cc}
}

func (c *countOnlineClientClient) CountOnlineClient(ctx context.Context, in *CountOnlineClientRequest, opts ...grpc.CallOption) (*CountOnlineClientResponse, error) {
	out := new(CountOnlineClientResponse)
	err := c.cc.Invoke(ctx, "/proto.CountOnlineClient/CountOnlineClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountOnlineClientServer is the server API for CountOnlineClient service.
type CountOnlineClientServer interface {
	CountOnlineClient(context.Context, *CountOnlineClientRequest) (*CountOnlineClientResponse, error)
}

// UnimplementedCountOnlineClientServer can be embedded to have forward compatible implementations.
type UnimplementedCountOnlineClientServer struct {
}

func (*UnimplementedCountOnlineClientServer) CountOnlineClient(context.Context, *CountOnlineClientRequest) (*CountOnlineClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountOnlineClient not implemented")
}

func RegisterCountOnlineClientServer(s *grpc.Server, srv CountOnlineClientServer) {
	s.RegisterService(&_CountOnlineClient_serviceDesc, srv)
}

func _CountOnlineClient_CountOnlineClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountOnlineClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountOnlineClientServer).CountOnlineClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountOnlineClient/CountOnlineClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountOnlineClientServer).CountOnlineClient(ctx, req.(*CountOnlineClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountOnlineClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CountOnlineClient",
	HandlerType: (*CountOnlineClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountOnlineClient",
			Handler:    _CountOnlineClient_CountOnlineClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "countonlineclient.proto",
}
