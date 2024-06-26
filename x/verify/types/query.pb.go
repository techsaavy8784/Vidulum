// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vidulum/verify/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetExternalAddressRequest struct {
	Vidulum string `protobuf:"bytes,1,opt,name=vidulum,proto3" json:"vidulum,omitempty"`
}

func (m *QueryGetExternalAddressRequest) Reset()         { *m = QueryGetExternalAddressRequest{} }
func (m *QueryGetExternalAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetExternalAddressRequest) ProtoMessage()    {}
func (*QueryGetExternalAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{2}
}
func (m *QueryGetExternalAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetExternalAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetExternalAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetExternalAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetExternalAddressRequest.Merge(m, src)
}
func (m *QueryGetExternalAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetExternalAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetExternalAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetExternalAddressRequest proto.InternalMessageInfo

func (m *QueryGetExternalAddressRequest) GetVidulum() string {
	if m != nil {
		return m.Vidulum
	}
	return ""
}

type QueryGetExternalAddressResponse struct {
	ExternalAddress ExternalAddress `protobuf:"bytes,1,opt,name=externalAddress,proto3" json:"externalAddress"`
}

func (m *QueryGetExternalAddressResponse) Reset()         { *m = QueryGetExternalAddressResponse{} }
func (m *QueryGetExternalAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetExternalAddressResponse) ProtoMessage()    {}
func (*QueryGetExternalAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{3}
}
func (m *QueryGetExternalAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetExternalAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetExternalAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetExternalAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetExternalAddressResponse.Merge(m, src)
}
func (m *QueryGetExternalAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetExternalAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetExternalAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetExternalAddressResponse proto.InternalMessageInfo

func (m *QueryGetExternalAddressResponse) GetExternalAddress() ExternalAddress {
	if m != nil {
		return m.ExternalAddress
	}
	return ExternalAddress{}
}

type QueryAllExternalAddressRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllExternalAddressRequest) Reset()         { *m = QueryAllExternalAddressRequest{} }
func (m *QueryAllExternalAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllExternalAddressRequest) ProtoMessage()    {}
func (*QueryAllExternalAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{4}
}
func (m *QueryAllExternalAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllExternalAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllExternalAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllExternalAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllExternalAddressRequest.Merge(m, src)
}
func (m *QueryAllExternalAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllExternalAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllExternalAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllExternalAddressRequest proto.InternalMessageInfo

func (m *QueryAllExternalAddressRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllExternalAddressResponse struct {
	ExternalAddress []ExternalAddress   `protobuf:"bytes,1,rep,name=externalAddress,proto3" json:"externalAddress"`
	Pagination      *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllExternalAddressResponse) Reset()         { *m = QueryAllExternalAddressResponse{} }
func (m *QueryAllExternalAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllExternalAddressResponse) ProtoMessage()    {}
func (*QueryAllExternalAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f302e47e6f3257, []int{5}
}
func (m *QueryAllExternalAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllExternalAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllExternalAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllExternalAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllExternalAddressResponse.Merge(m, src)
}
func (m *QueryAllExternalAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllExternalAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllExternalAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllExternalAddressResponse proto.InternalMessageInfo

func (m *QueryAllExternalAddressResponse) GetExternalAddress() []ExternalAddress {
	if m != nil {
		return m.ExternalAddress
	}
	return nil
}

func (m *QueryAllExternalAddressResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "vidulum.verify.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "vidulum.verify.QueryParamsResponse")
	proto.RegisterType((*QueryGetExternalAddressRequest)(nil), "vidulum.verify.QueryGetExternalAddressRequest")
	proto.RegisterType((*QueryGetExternalAddressResponse)(nil), "vidulum.verify.QueryGetExternalAddressResponse")
	proto.RegisterType((*QueryAllExternalAddressRequest)(nil), "vidulum.verify.QueryAllExternalAddressRequest")
	proto.RegisterType((*QueryAllExternalAddressResponse)(nil), "vidulum.verify.QueryAllExternalAddressResponse")
}

func init() { proto.RegisterFile("vidulum/verify/query.proto", fileDescriptor_36f302e47e6f3257) }

var fileDescriptor_36f302e47e6f3257 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xad, 0x46, 0x1c, 0xc1, 0xc2, 0x58, 0x42, 0x58, 0x65, 0x52, 0x46, 0xd4, 0x52,
	0x64, 0xc6, 0x46, 0x4f, 0xde, 0x52, 0xd0, 0x1e, 0x3c, 0x58, 0x73, 0xf4, 0x22, 0x13, 0x33, 0xae,
	0x0b, 0x9b, 0x9d, 0xcd, 0xce, 0x24, 0x34, 0x88, 0x17, 0x3f, 0x81, 0xe0, 0x07, 0xf0, 0xe6, 0x97,
	0xf0, 0x2c, 0xf4, 0x58, 0xf0, 0xe2, 0x49, 0x24, 0xf1, 0x83, 0x48, 0x66, 0xde, 0x62, 0x77, 0xcc,
	0x36, 0xd2, 0x5b, 0xb2, 0xef, 0xff, 0xff, 0xbf, 0xdf, 0xe3, 0xbd, 0x5d, 0x1c, 0x4d, 0x93, 0xe1,
	0x24, 0x9d, 0x8c, 0xc4, 0x54, 0x15, 0xc9, 0x9b, 0x99, 0x18, 0x4f, 0x54, 0x31, 0xe3, 0x79, 0xa1,
	0xad, 0x26, 0xd7, 0xa1, 0xc6, 0x7d, 0x2d, 0xda, 0x8e, 0x75, 0xac, 0x5d, 0x49, 0x2c, 0x7f, 0x79,
	0x55, 0x74, 0x2b, 0xd6, 0x3a, 0x4e, 0x95, 0x90, 0x79, 0x22, 0x64, 0x96, 0x69, 0x2b, 0x6d, 0xa2,
	0x33, 0x03, 0xd5, 0xbd, 0xd7, 0xda, 0x8c, 0xb4, 0x11, 0x03, 0x69, 0x94, 0x0f, 0x17, 0xd3, 0xfd,
	0x81, 0xb2, 0x72, 0x5f, 0xe4, 0x32, 0x4e, 0x32, 0x27, 0x06, 0xed, 0xcd, 0x80, 0x25, 0x97, 0x85,
	0x1c, 0x95, 0x41, 0x77, 0x82, 0xa2, 0x3a, 0xb6, 0xaa, 0xc8, 0x64, 0xfa, 0x4a, 0x0e, 0x87, 0x85,
	0x32, 0x20, 0x63, 0xdb, 0x98, 0xbc, 0x58, 0x76, 0x39, 0x72, 0xde, 0xbe, 0x1a, 0x4f, 0x94, 0xb1,
	0xec, 0x19, 0xbe, 0x51, 0x79, 0x6a, 0x72, 0x9d, 0x19, 0x45, 0x1e, 0xe1, 0xa6, 0xef, 0xd1, 0x46,
	0x3b, 0x68, 0xf7, 0x5a, 0xb7, 0xc5, 0xab, 0x13, 0x73, 0xaf, 0x3f, 0xb8, 0x74, 0xf2, 0xb3, 0xd3,
	0xe8, 0x83, 0x96, 0x3d, 0xc6, 0xd4, 0x85, 0x1d, 0x2a, 0xfb, 0x04, 0x20, 0x7a, 0x9e, 0x01, 0xda,
	0x91, 0x36, 0xbe, 0x02, 0x41, 0x2e, 0xf8, 0x6a, 0xbf, 0xfc, 0xcb, 0x0a, 0xdc, 0xa9, 0xf5, 0x02,
	0xd4, 0x73, 0xbc, 0xa5, 0xaa, 0x25, 0xa0, 0xeb, 0x84, 0x74, 0x41, 0x02, 0x60, 0x86, 0x6e, 0xf6,
	0x16, 0x78, 0x7b, 0x69, 0x5a, 0xc3, 0xfb, 0x14, 0xe3, 0xbf, 0xcb, 0x80, 0x6e, 0x77, 0xb9, 0xdf,
	0x1c, 0x5f, 0x6e, 0x8e, 0xfb, 0xb3, 0x80, 0xcd, 0xf1, 0x23, 0x19, 0x2b, 0xf0, 0xf6, 0xcf, 0x38,
	0xd9, 0x57, 0x04, 0xe3, 0xad, 0x6a, 0x75, 0xde, 0x78, 0x9b, 0x17, 0x1f, 0x8f, 0x1c, 0x56, 0xe0,
	0x37, 0x1c, 0xfc, 0xbd, 0xb5, 0xf0, 0x9e, 0xe6, 0x2c, 0x7d, 0xf7, 0xdb, 0x26, 0xbe, 0xec, 0xe8,
	0xc9, 0x18, 0x37, 0xfd, 0xe6, 0x09, 0x0b, 0xa1, 0xfe, 0x3d, 0xae, 0xe8, 0xf6, 0xb9, 0x1a, 0xdf,
	0x88, 0xd1, 0x0f, 0xdf, 0x7f, 0x7f, 0xda, 0x68, 0x93, 0x96, 0x58, 0x79, 0xe4, 0xe4, 0x0b, 0xc2,
	0x5b, 0xc1, 0xc0, 0x84, 0xaf, 0x0c, 0xae, 0x3d, 0xbb, 0x48, 0xfc, 0xb7, 0x1e, 0xa0, 0xba, 0x0e,
	0xea, 0x3e, 0xd9, 0x13, 0x6b, 0x5e, 0x2e, 0xf1, 0x0e, 0x04, 0xef, 0xc9, 0x67, 0x84, 0x49, 0x90,
	0xd7, 0x4b, 0xd3, 0x1a, 0xd6, 0xda, 0x93, 0xab, 0x61, 0xad, 0xbf, 0x1b, 0xb6, 0xeb, 0x58, 0x19,
	0xd9, 0x59, 0xc7, 0x7a, 0xf0, 0xe0, 0x64, 0x4e, 0xd1, 0xe9, 0x9c, 0xa2, 0x5f, 0x73, 0x8a, 0x3e,
	0x2e, 0x68, 0xe3, 0x74, 0x41, 0x1b, 0x3f, 0x16, 0xb4, 0xf1, 0xb2, 0x55, 0x5a, 0x8f, 0x4b, 0xb3,
	0x9d, 0xe5, 0xca, 0x0c, 0x9a, 0xee, 0xdb, 0xf1, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09,
	0xa7, 0xc1, 0x3f, 0x0d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ExternalAddress items.
	ExternalAddress(ctx context.Context, in *QueryGetExternalAddressRequest, opts ...grpc.CallOption) (*QueryGetExternalAddressResponse, error)
	ExternalAddressAll(ctx context.Context, in *QueryAllExternalAddressRequest, opts ...grpc.CallOption) (*QueryAllExternalAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/vidulum.verify.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ExternalAddress(ctx context.Context, in *QueryGetExternalAddressRequest, opts ...grpc.CallOption) (*QueryGetExternalAddressResponse, error) {
	out := new(QueryGetExternalAddressResponse)
	err := c.cc.Invoke(ctx, "/vidulum.verify.Query/ExternalAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ExternalAddressAll(ctx context.Context, in *QueryAllExternalAddressRequest, opts ...grpc.CallOption) (*QueryAllExternalAddressResponse, error) {
	out := new(QueryAllExternalAddressResponse)
	err := c.cc.Invoke(ctx, "/vidulum.verify.Query/ExternalAddressAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ExternalAddress items.
	ExternalAddress(context.Context, *QueryGetExternalAddressRequest) (*QueryGetExternalAddressResponse, error)
	ExternalAddressAll(context.Context, *QueryAllExternalAddressRequest) (*QueryAllExternalAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ExternalAddress(ctx context.Context, req *QueryGetExternalAddressRequest) (*QueryGetExternalAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExternalAddress not implemented")
}
func (*UnimplementedQueryServer) ExternalAddressAll(ctx context.Context, req *QueryAllExternalAddressRequest) (*QueryAllExternalAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExternalAddressAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vidulum.verify.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ExternalAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetExternalAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExternalAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vidulum.verify.Query/ExternalAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExternalAddress(ctx, req.(*QueryGetExternalAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ExternalAddressAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllExternalAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExternalAddressAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vidulum.verify.Query/ExternalAddressAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExternalAddressAll(ctx, req.(*QueryAllExternalAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vidulum.verify.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ExternalAddress",
			Handler:    _Query_ExternalAddress_Handler,
		},
		{
			MethodName: "ExternalAddressAll",
			Handler:    _Query_ExternalAddressAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vidulum/verify/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetExternalAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetExternalAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetExternalAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vidulum) > 0 {
		i -= len(m.Vidulum)
		copy(dAtA[i:], m.Vidulum)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Vidulum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetExternalAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetExternalAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetExternalAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ExternalAddress.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllExternalAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllExternalAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllExternalAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllExternalAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllExternalAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllExternalAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExternalAddress) > 0 {
		for iNdEx := len(m.ExternalAddress) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalAddress[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetExternalAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Vidulum)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetExternalAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ExternalAddress.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllExternalAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllExternalAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ExternalAddress) > 0 {
		for _, e := range m.ExternalAddress {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetExternalAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetExternalAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetExternalAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vidulum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vidulum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetExternalAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetExternalAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetExternalAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllExternalAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllExternalAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllExternalAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllExternalAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllExternalAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllExternalAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalAddress = append(m.ExternalAddress, ExternalAddress{})
			if err := m.ExternalAddress[len(m.ExternalAddress)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
