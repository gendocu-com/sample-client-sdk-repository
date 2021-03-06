// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: example.proto

package example

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ExampleMessage1 - Example Leading Comment for ExampleMessage1
type ExampleMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyString string `protobuf:"bytes,1,opt,name=MyString,proto3" json:"MyString,omitempty"`
}

func (x *ExampleMessage1) Reset() {
	*x = ExampleMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage1) ProtoMessage() {}

func (x *ExampleMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage1.ProtoReflect.Descriptor instead.
func (*ExampleMessage1) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleMessage1) GetMyString() string {
	if x != nil {
		return x.MyString
	}
	return ""
}

//
//ExampleMessage2 - Example Leading Comment for ExampleMessage2
type ExampleMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyInt  int32                          `protobuf:"varint,1,opt,name=MyInt,proto3" json:"MyInt,omitempty"`
	Nested *ExampleMessage2_ExampleNested `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *ExampleMessage2) Reset() {
	*x = ExampleMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage2) ProtoMessage() {}

func (x *ExampleMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage2.ProtoReflect.Descriptor instead.
func (*ExampleMessage2) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleMessage2) GetMyInt() int32 {
	if x != nil {
		return x.MyInt
	}
	return 0
}

func (x *ExampleMessage2) GetNested() *ExampleMessage2_ExampleNested {
	if x != nil {
		return x.Nested
	}
	return nil
}

//
//ReturnType - Empty Structure Placeholder
type ReturnType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReturnType) Reset() {
	*x = ReturnType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnType) ProtoMessage() {}

func (x *ReturnType) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnType.ProtoReflect.Descriptor instead.
func (*ReturnType) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

// MyInt - Example trailing Comment
type ExampleMessage2_ExampleNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExampleMessage2_ExampleNested) Reset() {
	*x = ExampleMessage2_ExampleNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage2_ExampleNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage2_ExampleNested) ProtoMessage() {}

func (x *ExampleMessage2_ExampleNested) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage2_ExampleNested.ProtoReflect.Descriptor instead.
func (*ExampleMessage2_ExampleNested) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ExampleMessage2_ExampleNested) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2d, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x84,
	0x01, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x79, 0x49, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4d, 0x79, 0x49, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x1a, 0x23, 0x0a, 0x0d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x32, 0x72, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x31, 0x12, 0x10, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x32, 0x12, 0x10, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_proto_goTypes = []interface{}{
	(*ExampleMessage1)(nil),               // 0: ExampleMessage1
	(*ExampleMessage2)(nil),               // 1: ExampleMessage2
	(*ReturnType)(nil),                    // 2: ReturnType
	(*ExampleMessage2_ExampleNested)(nil), // 3: ExampleMessage2.ExampleNested
}
var file_example_proto_depIdxs = []int32{
	3, // 0: ExampleMessage2.nested:type_name -> ExampleMessage2.ExampleNested
	0, // 1: ExampleService.ExampleCall1:input_type -> ExampleMessage1
	1, // 2: ExampleService.ExampleCall2:input_type -> ExampleMessage2
	2, // 3: ExampleService.ExampleCall1:output_type -> ReturnType
	2, // 4: ExampleService.ExampleCall2:output_type -> ReturnType
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage1); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage2); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnType); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage2_ExampleNested); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
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
	err := c.cc.Invoke(ctx, "/ExampleService/ExampleCall1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExampleCall2(ctx context.Context, in *ExampleMessage2, opts ...grpc.CallOption) (*ReturnType, error) {
	out := new(ReturnType)
	err := c.cc.Invoke(ctx, "/ExampleService/ExampleCall2", in, out, opts...)
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
		FullMethod: "/ExampleService/ExampleCall1",
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
		FullMethod: "/ExampleService/ExampleCall2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleCall2(ctx, req.(*ExampleMessage2))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExampleService",
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
	Metadata: "example.proto",
}
