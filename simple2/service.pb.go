// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: service.proto

package xx

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x7e, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x31, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x32, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x1a, 0x0e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x65, 0x6e, 0x64, 0x6f, 0x63, 0x75, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x3b,
	0x78, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*ExampleMessage1)(nil), // 0: v1.ExampleMessage1
	(*ExampleMessage2)(nil), // 1: v1.ExampleMessage2
	(*ReturnType)(nil),      // 2: v1.ReturnType
}
var file_service_proto_depIdxs = []int32{
	0, // 0: v1.ExampleService.ExampleCall1:input_type -> v1.ExampleMessage1
	1, // 1: v1.ExampleService.ExampleCall2:input_type -> v1.ExampleMessage2
	2, // 2: v1.ExampleService.ExampleCall1:output_type -> v1.ReturnType
	2, // 3: v1.ExampleService.ExampleCall2:output_type -> v1.ReturnType
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	ExampleCall1(ctx context.Context, in *ExampleMessage1, opts ...grpc.CallOption) (*ReturnType, error)
	ExampleCall2(ctx context.Context, in *ExampleMessage2, opts ...grpc.CallOption) (*ReturnType, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) ExampleCall1(ctx context.Context, in *ExampleMessage1, opts ...grpc.CallOption) (*ReturnType, error) {
	out := new(ReturnType)
	err := c.cc.Invoke(ctx, "/v1.ExampleService/ExampleCall1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExampleCall2(ctx context.Context, in *ExampleMessage2, opts ...grpc.CallOption) (*ReturnType, error) {
	out := new(ReturnType)
	err := c.cc.Invoke(ctx, "/v1.ExampleService/ExampleCall2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	ExampleCall1(context.Context, *ExampleMessage1) (*ReturnType, error)
	ExampleCall2(context.Context, *ExampleMessage2) (*ReturnType, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) ExampleCall1(context.Context, *ExampleMessage1) (*ReturnType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleCall1 not implemented")
}
func (*UnimplementedExampleServiceServer) ExampleCall2(context.Context, *ExampleMessage2) (*ReturnType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleCall2 not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_ExampleCall1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleMessage1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleCall1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ExampleService/ExampleCall1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleCall1(ctx, req.(*ExampleMessage1))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ExampleCall2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleMessage2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleCall2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ExampleService/ExampleCall2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleCall2(ctx, req.(*ExampleMessage2))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleCall1",
			Handler:    _ExampleService_ExampleCall1_Handler,
		},
		{
			MethodName: "ExampleCall2",
			Handler:    _ExampleService_ExampleCall2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
