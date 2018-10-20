// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package models

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x81, 0x72, 0xf9, 0x8a, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0xa1, 0xd2, 0x46, 0xf3,
	0x19, 0xb9, 0xb8, 0x3c, 0xf3, 0x52, 0x52, 0x2b, 0x52, 0x8b, 0x1c, 0x03, 0x3c, 0x85, 0xdc, 0xb8,
	0x84, 0x82, 0x4b, 0x93, 0x72, 0x33, 0x4b, 0xc0, 0x62, 0x41, 0x10, 0xad, 0x42, 0xa2, 0x7a, 0x30,
	0x43, 0x90, 0x85, 0xa5, 0xc4, 0xf5, 0xe0, 0x86, 0x41, 0xc5, 0x21, 0x3c, 0x25, 0x06, 0x21, 0x4f,
	0x2e, 0x61, 0x88, 0x39, 0xc1, 0xa9, 0x89, 0x45, 0xc9, 0x19, 0x30, 0x83, 0xc4, 0xe0, 0x06, 0xa1,
	0x88, 0x4b, 0x49, 0x20, 0x4c, 0x82, 0x49, 0xc0, 0x8c, 0x4a, 0x62, 0x03, 0x3b, 0xd4, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x11, 0x8c, 0x96, 0xd8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexerAPIClient is the client API for IndexerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexerAPIClient interface {
	// SubmitIndexRequest is used to submit content to be indexed by the lens system
	SubmitIndexRequest(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	// SubmitSearchRequest is used to search the lens system
	SubmitSearchRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type indexerAPIClient struct {
	cc *grpc.ClientConn
}

func NewIndexerAPIClient(cc *grpc.ClientConn) IndexerAPIClient {
	return &indexerAPIClient{cc}
}

func (c *indexerAPIClient) SubmitIndexRequest(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/IndexerAPI/SubmitIndexRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerAPIClient) SubmitSearchRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/IndexerAPI/SubmitSearchRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerAPIServer is the server API for IndexerAPI service.
type IndexerAPIServer interface {
	// SubmitIndexRequest is used to submit content to be indexed by the lens system
	SubmitIndexRequest(context.Context, *IndexRequest) (*IndexResponse, error)
	// SubmitSearchRequest is used to search the lens system
	SubmitSearchRequest(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterIndexerAPIServer(s *grpc.Server, srv IndexerAPIServer) {
	s.RegisterService(&_IndexerAPI_serviceDesc, srv)
}

func _IndexerAPI_SubmitIndexRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerAPIServer).SubmitIndexRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IndexerAPI/SubmitIndexRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerAPIServer).SubmitIndexRequest(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexerAPI_SubmitSearchRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerAPIServer).SubmitSearchRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IndexerAPI/SubmitSearchRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerAPIServer).SubmitSearchRequest(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexerAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "IndexerAPI",
	HandlerType: (*IndexerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitIndexRequest",
			Handler:    _IndexerAPI_SubmitIndexRequest_Handler,
		},
		{
			MethodName: "SubmitSearchRequest",
			Handler:    _IndexerAPI_SubmitSearchRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}