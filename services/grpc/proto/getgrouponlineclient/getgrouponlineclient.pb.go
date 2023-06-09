// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: getgrouponlineclient.proto

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

type GetGroupOnlineClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupOnlineClientRequest) Reset() {
	*x = GetGroupOnlineClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getgrouponlineclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupOnlineClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupOnlineClientRequest) ProtoMessage() {}

func (x *GetGroupOnlineClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getgrouponlineclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupOnlineClientRequest.ProtoReflect.Descriptor instead.
func (*GetGroupOnlineClientRequest) Descriptor() ([]byte, []int) {
	return file_getgrouponlineclient_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupOnlineClientRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type GetGroupOnlineClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clientid []string `protobuf:"bytes,1,rep,name=clientid,proto3" json:"clientid,omitempty"`
}

func (x *GetGroupOnlineClientResponse) Reset() {
	*x = GetGroupOnlineClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getgrouponlineclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupOnlineClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupOnlineClientResponse) ProtoMessage() {}

func (x *GetGroupOnlineClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_getgrouponlineclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupOnlineClientResponse.ProtoReflect.Descriptor instead.
func (*GetGroupOnlineClientResponse) Descriptor() ([]byte, []int) {
	return file_getgrouponlineclient_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupOnlineClientResponse) GetClientid() []string {
	if x != nil {
		return x.Clientid
	}
	return nil
}

var File_getgrouponlineclient_proto protoreflect.FileDescriptor

var file_getgrouponlineclient_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x65, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x69, 0x64, 0x32, 0x79, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_getgrouponlineclient_proto_rawDescOnce sync.Once
	file_getgrouponlineclient_proto_rawDescData = file_getgrouponlineclient_proto_rawDesc
)

func file_getgrouponlineclient_proto_rawDescGZIP() []byte {
	file_getgrouponlineclient_proto_rawDescOnce.Do(func() {
		file_getgrouponlineclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_getgrouponlineclient_proto_rawDescData)
	})
	return file_getgrouponlineclient_proto_rawDescData
}

var file_getgrouponlineclient_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_getgrouponlineclient_proto_goTypes = []interface{}{
	(*GetGroupOnlineClientRequest)(nil),  // 0: proto.GetGroupOnlineClientRequest
	(*GetGroupOnlineClientResponse)(nil), // 1: proto.GetGroupOnlineClientResponse
}
var file_getgrouponlineclient_proto_depIdxs = []int32{
	0, // 0: proto.GetGroupOnlineClient.GetGroupOnlineClient:input_type -> proto.GetGroupOnlineClientRequest
	1, // 1: proto.GetGroupOnlineClient.GetGroupOnlineClient:output_type -> proto.GetGroupOnlineClientResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_getgrouponlineclient_proto_init() }
func file_getgrouponlineclient_proto_init() {
	if File_getgrouponlineclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_getgrouponlineclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupOnlineClientRequest); i {
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
		file_getgrouponlineclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupOnlineClientResponse); i {
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
			RawDescriptor: file_getgrouponlineclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_getgrouponlineclient_proto_goTypes,
		DependencyIndexes: file_getgrouponlineclient_proto_depIdxs,
		MessageInfos:      file_getgrouponlineclient_proto_msgTypes,
	}.Build()
	File_getgrouponlineclient_proto = out.File
	file_getgrouponlineclient_proto_rawDesc = nil
	file_getgrouponlineclient_proto_goTypes = nil
	file_getgrouponlineclient_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GetGroupOnlineClientClient is the client API for GetGroupOnlineClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetGroupOnlineClientClient interface {
	GetGroupOnlineClient(ctx context.Context, in *GetGroupOnlineClientRequest, opts ...grpc.CallOption) (*GetGroupOnlineClientResponse, error)
}

type getGroupOnlineClientClient struct {
	cc grpc.ClientConnInterface
}

func NewGetGroupOnlineClientClient(cc grpc.ClientConnInterface) GetGroupOnlineClientClient {
	return &getGroupOnlineClientClient{cc}
}

func (c *getGroupOnlineClientClient) GetGroupOnlineClient(ctx context.Context, in *GetGroupOnlineClientRequest, opts ...grpc.CallOption) (*GetGroupOnlineClientResponse, error) {
	out := new(GetGroupOnlineClientResponse)
	err := c.cc.Invoke(ctx, "/proto.GetGroupOnlineClient/GetGroupOnlineClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetGroupOnlineClientServer is the server API for GetGroupOnlineClient service.
type GetGroupOnlineClientServer interface {
	GetGroupOnlineClient(context.Context, *GetGroupOnlineClientRequest) (*GetGroupOnlineClientResponse, error)
}

// UnimplementedGetGroupOnlineClientServer can be embedded to have forward compatible implementations.
type UnimplementedGetGroupOnlineClientServer struct {
}

func (*UnimplementedGetGroupOnlineClientServer) GetGroupOnlineClient(context.Context, *GetGroupOnlineClientRequest) (*GetGroupOnlineClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupOnlineClient not implemented")
}

func RegisterGetGroupOnlineClientServer(s *grpc.Server, srv GetGroupOnlineClientServer) {
	s.RegisterService(&_GetGroupOnlineClient_serviceDesc, srv)
}

func _GetGroupOnlineClient_GetGroupOnlineClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupOnlineClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGroupOnlineClientServer).GetGroupOnlineClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GetGroupOnlineClient/GetGroupOnlineClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGroupOnlineClientServer).GetGroupOnlineClient(ctx, req.(*GetGroupOnlineClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetGroupOnlineClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GetGroupOnlineClient",
	HandlerType: (*GetGroupOnlineClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupOnlineClient",
			Handler:    _GetGroupOnlineClient_GetGroupOnlineClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getgrouponlineclient.proto",
}
