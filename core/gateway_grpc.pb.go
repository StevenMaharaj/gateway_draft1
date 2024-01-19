// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/gateway.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	SendEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Gateway_SendEventClient, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) SendEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Gateway_SendEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[0], "/gateway_draft1.Gateway/SendEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySendEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SendEventClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type gatewaySendEventClient struct {
	grpc.ClientStream
}

func (x *gatewaySendEventClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	SendEvent(*emptypb.Empty, Gateway_SendEventServer) error
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) SendEvent(*emptypb.Empty, Gateway_SendEventServer) error {
	return status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_SendEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).SendEvent(m, &gatewaySendEventServer{stream})
}

type Gateway_SendEventServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type gatewaySendEventServer struct {
	grpc.ServerStream
}

func (x *gatewaySendEventServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway_draft1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendEvent",
			Handler:       _Gateway_SendEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gateway.proto",
}
