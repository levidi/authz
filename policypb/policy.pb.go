// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: policypb/policy.proto

package policypb

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

type PolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PolicyRequest) Reset() {
	*x = PolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policypb_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRequest) ProtoMessage() {}

func (x *PolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policypb_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRequest.ProtoReflect.Descriptor instead.
func (*PolicyRequest) Descriptor() ([]byte, []int) {
	return file_policypb_policy_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *PolicyRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type PolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PolicyResponse) Reset() {
	*x = PolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policypb_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyResponse) ProtoMessage() {}

func (x *PolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policypb_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyResponse.ProtoReflect.Descriptor instead.
func (*PolicyResponse) Descriptor() ([]byte, []int) {
	return file_policypb_policy_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PolicyResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_policypb_policy_proto protoreflect.FileDescriptor

var file_policypb_policy_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22,
	0x3b, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x0e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0x4d, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x15,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policypb_policy_proto_rawDescOnce sync.Once
	file_policypb_policy_proto_rawDescData = file_policypb_policy_proto_rawDesc
)

func file_policypb_policy_proto_rawDescGZIP() []byte {
	file_policypb_policy_proto_rawDescOnce.Do(func() {
		file_policypb_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policypb_policy_proto_rawDescData)
	})
	return file_policypb_policy_proto_rawDescData
}

var file_policypb_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_policypb_policy_proto_goTypes = []interface{}{
	(*PolicyRequest)(nil),  // 0: policy.PolicyRequest
	(*PolicyResponse)(nil), // 1: policy.PolicyResponse
}
var file_policypb_policy_proto_depIdxs = []int32{
	0, // 0: policy.PolicyService.SetPolicy:input_type -> policy.PolicyRequest
	1, // 1: policy.PolicyService.SetPolicy:output_type -> policy.PolicyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_policypb_policy_proto_init() }
func file_policypb_policy_proto_init() {
	if File_policypb_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policypb_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyRequest); i {
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
		file_policypb_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyResponse); i {
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
			RawDescriptor: file_policypb_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policypb_policy_proto_goTypes,
		DependencyIndexes: file_policypb_policy_proto_depIdxs,
		MessageInfos:      file_policypb_policy_proto_msgTypes,
	}.Build()
	File_policypb_policy_proto = out.File
	file_policypb_policy_proto_rawDesc = nil
	file_policypb_policy_proto_goTypes = nil
	file_policypb_policy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolicyServiceClient interface {
	SetPolicy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyResponse, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) SetPolicy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyResponse, error) {
	out := new(PolicyResponse)
	err := c.cc.Invoke(ctx, "/policy.PolicyService/SetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
type PolicyServiceServer interface {
	SetPolicy(context.Context, *PolicyRequest) (*PolicyResponse, error)
}

// UnimplementedPolicyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (*UnimplementedPolicyServiceServer) SetPolicy(context.Context, *PolicyRequest) (*PolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPolicy not implemented")
}

func RegisterPolicyServiceServer(s *grpc.Server, srv PolicyServiceServer) {
	s.RegisterService(&_PolicyService_serviceDesc, srv)
}

func _PolicyService_SetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).SetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy.PolicyService/SetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).SetPolicy(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "policy.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPolicy",
			Handler:    _PolicyService_SetPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policypb/policy.proto",
}