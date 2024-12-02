// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.15.5
// source: proto/acts.proto

package acts_grpc

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
	ActsService_Send_FullMethodName      = "/acts.grpc.ActsService/Send"
	ActsService_OnMessage_FullMethodName = "/acts.grpc.ActsService/OnMessage"
)

// ActsServiceClient is the client API for ActsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// acts service
type ActsServiceClient interface {
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	// rpc OnFlow(Message) returns (stream Message) {}
	// rpc OnStep(Message) returns (stream Message) {}
	// rpc OnAct(Message) returns (stream Message) {}
	OnMessage(ctx context.Context, in *MessageOptions, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error)
}

type actsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActsServiceClient(cc grpc.ClientConnInterface) ActsServiceClient {
	return &actsServiceClient{cc}
}

func (c *actsServiceClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, ActsService_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actsServiceClient) OnMessage(ctx context.Context, in *MessageOptions, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActsService_ServiceDesc.Streams[0], ActsService_OnMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MessageOptions, Message]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActsService_OnMessageClient = grpc.ServerStreamingClient[Message]

// ActsServiceServer is the server API for ActsService service.
// All implementations must embed UnimplementedActsServiceServer
// for forward compatibility.
//
// acts service
type ActsServiceServer interface {
	Send(context.Context, *Message) (*Message, error)
	// rpc OnFlow(Message) returns (stream Message) {}
	// rpc OnStep(Message) returns (stream Message) {}
	// rpc OnAct(Message) returns (stream Message) {}
	OnMessage(*MessageOptions, grpc.ServerStreamingServer[Message]) error
	mustEmbedUnimplementedActsServiceServer()
}

// UnimplementedActsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActsServiceServer struct{}

func (UnimplementedActsServiceServer) Send(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedActsServiceServer) OnMessage(*MessageOptions, grpc.ServerStreamingServer[Message]) error {
	return status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedActsServiceServer) mustEmbedUnimplementedActsServiceServer() {}
func (UnimplementedActsServiceServer) testEmbeddedByValue()                     {}

// UnsafeActsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActsServiceServer will
// result in compilation errors.
type UnsafeActsServiceServer interface {
	mustEmbedUnimplementedActsServiceServer()
}

func RegisterActsServiceServer(s grpc.ServiceRegistrar, srv ActsServiceServer) {
	// If the following call pancis, it indicates UnimplementedActsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActsService_ServiceDesc, srv)
}

func _ActsService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActsServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActsService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActsServiceServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActsService_OnMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActsServiceServer).OnMessage(m, &grpc.GenericServerStream[MessageOptions, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActsService_OnMessageServer = grpc.ServerStreamingServer[Message]

// ActsService_ServiceDesc is the grpc.ServiceDesc for ActsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "acts.grpc.ActsService",
	HandlerType: (*ActsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _ActsService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnMessage",
			Handler:       _ActsService_OnMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/acts.proto",
}
