// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: zk_prover.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ZKProverServiceClient is the client API for ZKProverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZKProverServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GenProof(ctx context.Context, in *GenProofRequest, opts ...grpc.CallOption) (*GenProofResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	GetProof(ctx context.Context, opts ...grpc.CallOption) (ZKProverService_GetProofClient, error)
}

type zKProverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZKProverServiceClient(cc grpc.ClientConnInterface) ZKProverServiceClient {
	return &zKProverServiceClient{cc}
}

func (c *zKProverServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/zkprover.v1.ZKProverService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKProverServiceClient) GenProof(ctx context.Context, in *GenProofRequest, opts ...grpc.CallOption) (*GenProofResponse, error) {
	out := new(GenProofResponse)
	err := c.cc.Invoke(ctx, "/zkprover.v1.ZKProverService/GenProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKProverServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/zkprover.v1.ZKProverService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKProverServiceClient) GetProof(ctx context.Context, opts ...grpc.CallOption) (ZKProverService_GetProofClient, error) {
	stream, err := c.cc.NewStream(ctx, &ZKProverService_ServiceDesc.Streams[0], "/zkprover.v1.ZKProverService/GetProof", opts...)
	if err != nil {
		return nil, err
	}
	x := &zKProverServiceGetProofClient{stream}
	return x, nil
}

type ZKProverService_GetProofClient interface {
	Send(*GetProofRequest) error
	Recv() (*GetProofResponse, error)
	grpc.ClientStream
}

type zKProverServiceGetProofClient struct {
	grpc.ClientStream
}

func (x *zKProverServiceGetProofClient) Send(m *GetProofRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zKProverServiceGetProofClient) Recv() (*GetProofResponse, error) {
	m := new(GetProofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ZKProverServiceServer is the server API for ZKProverService service.
// All implementations must embed UnimplementedZKProverServiceServer
// for forward compatibility
type ZKProverServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GenProof(context.Context, *GenProofRequest) (*GenProofResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	GetProof(ZKProverService_GetProofServer) error
	mustEmbedUnimplementedZKProverServiceServer()
}

// UnimplementedZKProverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedZKProverServiceServer struct {
}

func (UnimplementedZKProverServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedZKProverServiceServer) GenProof(context.Context, *GenProofRequest) (*GenProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenProof not implemented")
}
func (UnimplementedZKProverServiceServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedZKProverServiceServer) GetProof(ZKProverService_GetProofServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedZKProverServiceServer) mustEmbedUnimplementedZKProverServiceServer() {}

// UnsafeZKProverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZKProverServiceServer will
// result in compilation errors.
type UnsafeZKProverServiceServer interface {
	mustEmbedUnimplementedZKProverServiceServer()
}

func RegisterZKProverServiceServer(s grpc.ServiceRegistrar, srv ZKProverServiceServer) {
	s.RegisterService(&ZKProverService_ServiceDesc, srv)
}

func _ZKProverService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKProverServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkprover.v1.ZKProverService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKProverServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZKProverService_GenProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKProverServiceServer).GenProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkprover.v1.ZKProverService/GenProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKProverServiceServer).GenProof(ctx, req.(*GenProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZKProverService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKProverServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkprover.v1.ZKProverService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKProverServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZKProverService_GetProof_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZKProverServiceServer).GetProof(&zKProverServiceGetProofServer{stream})
}

type ZKProverService_GetProofServer interface {
	Send(*GetProofResponse) error
	Recv() (*GetProofRequest, error)
	grpc.ServerStream
}

type zKProverServiceGetProofServer struct {
	grpc.ServerStream
}

func (x *zKProverServiceGetProofServer) Send(m *GetProofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zKProverServiceGetProofServer) Recv() (*GetProofRequest, error) {
	m := new(GetProofRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ZKProverService_ServiceDesc is the grpc.ServiceDesc for ZKProverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZKProverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zkprover.v1.ZKProverService",
	HandlerType: (*ZKProverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ZKProverService_GetStatus_Handler,
		},
		{
			MethodName: "GenProof",
			Handler:    _ZKProverService_GenProof_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _ZKProverService_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProof",
			Handler:       _ZKProverService_GetProof_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "zk_prover.proto",
}
