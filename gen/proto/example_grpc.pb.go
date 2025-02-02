// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/example.proto

package examplepost

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Answer_SayAddress_FullMethodName = "/example.Answer/SayAddress"
)

// AnswerClient is the client API for Answer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnswerClient interface {
	SayAddress(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type answerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnswerClient(cc grpc.ClientConnInterface) AnswerClient {
	return &answerClient{cc}
}

func (c *answerClient) SayAddress(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, Answer_SayAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswerServer is the server API for Answer service.
// All implementations must embed UnimplementedAnswerServer
// for forward compatibility.
type AnswerServer interface {
	SayAddress(context.Context, *PostRequest) (*PostResponse, error)
	mustEmbedUnimplementedAnswerServer()
}

// UnimplementedAnswerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnswerServer struct{}

func (UnimplementedAnswerServer) SayAddress(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayAddress not implemented")
}
func (UnimplementedAnswerServer) mustEmbedUnimplementedAnswerServer() {}
func (UnimplementedAnswerServer) testEmbeddedByValue()                {}

// UnsafeAnswerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnswerServer will
// result in compilation errors.
type UnsafeAnswerServer interface {
	mustEmbedUnimplementedAnswerServer()
}

func RegisterAnswerServer(s grpc.ServiceRegistrar, srv AnswerServer) {
	// If the following call pancis, it indicates UnimplementedAnswerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Answer_ServiceDesc, srv)
}

func _Answer_SayAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServer).SayAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Answer_SayAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServer).SayAddress(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Answer_ServiceDesc is the grpc.ServiceDesc for Answer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Answer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Answer",
	HandlerType: (*AnswerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayAddress",
			Handler:    _Answer_SayAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/example.proto",
}
