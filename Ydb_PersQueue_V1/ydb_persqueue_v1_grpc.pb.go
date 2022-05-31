// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: ydb_persqueue_v1.proto

package Ydb_PersQueue_V1

import (
	context "context"
	Ydb_PersQueue_ClusterDiscovery "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_ClusterDiscovery"
	Ydb_PersQueue_V1 "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_V1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PersQueueServiceClient is the client API for PersQueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersQueueServiceClient interface {
	StreamingWrite(ctx context.Context, opts ...grpc.CallOption) (PersQueueService_StreamingWriteClient, error)
	//*
	// Creates Read Session
	// Pipeline:
	// client                  server
	//         Init(Topics, ClientId, ...)
	//        ---------------->
	//         Init(SessionId)
	//        <----------------
	//         read1
	//        ---------------->
	//         read2
	//        ---------------->
	//         assign(Topic1, Cluster, Partition1, ...) - assigns and releases are optional
	//        <----------------
	//         assign(Topic2, Clutster, Partition2, ...)
	//        <----------------
	//         start_read(Topic1, Partition1, ...) - client must respond to assign request with this message. Only after this client will start recieving messages from this partition
	//        ---------------->
	//         release(Topic1, Partition1, ...)
	//        <----------------
	//         released(Topic1, Partition1, ...) - only after released server will give this parittion to other session.
	//        ---------------->
	//         start_read(Topic2, Partition2, ...) - client must respond to assign request with this message. Only after this client will start recieving messages from this partition
	//        ---------------->
	//         read data(data, ...)
	//        <----------------
	//         commit(cookie1)
	//        ---------------->
	//         committed(cookie1)
	//        <----------------
	//         issue(description, ...)
	//        <----------------
	MigrationStreamingRead(ctx context.Context, opts ...grpc.CallOption) (PersQueueService_MigrationStreamingReadClient, error)
	// Get information about reading
	GetReadSessionsInfo(ctx context.Context, in *Ydb_PersQueue_V1.ReadInfoRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.ReadInfoResponse, error)
	//
	// Describe topic command.
	DescribeTopic(ctx context.Context, in *Ydb_PersQueue_V1.DescribeTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.DescribeTopicResponse, error)
	//
	// Drop topic command.
	DropTopic(ctx context.Context, in *Ydb_PersQueue_V1.DropTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.DropTopicResponse, error)
	//
	// Create topic command.
	CreateTopic(ctx context.Context, in *Ydb_PersQueue_V1.CreateTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.CreateTopicResponse, error)
	//
	// Alter topic command.
	AlterTopic(ctx context.Context, in *Ydb_PersQueue_V1.AlterTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.AlterTopicResponse, error)
	//
	// Add read rule command.
	AddReadRule(ctx context.Context, in *Ydb_PersQueue_V1.AddReadRuleRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.AddReadRuleResponse, error)
	//
	// Remove read rule command.
	RemoveReadRule(ctx context.Context, in *Ydb_PersQueue_V1.RemoveReadRuleRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.RemoveReadRuleResponse, error)
}

type persQueueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersQueueServiceClient(cc grpc.ClientConnInterface) PersQueueServiceClient {
	return &persQueueServiceClient{cc}
}

func (c *persQueueServiceClient) StreamingWrite(ctx context.Context, opts ...grpc.CallOption) (PersQueueService_StreamingWriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersQueueService_ServiceDesc.Streams[0], "/Ydb.PersQueue.V1.PersQueueService/StreamingWrite", opts...)
	if err != nil {
		return nil, err
	}
	x := &persQueueServiceStreamingWriteClient{stream}
	return x, nil
}

type PersQueueService_StreamingWriteClient interface {
	Send(*Ydb_PersQueue_V1.StreamingWriteClientMessage) error
	Recv() (*Ydb_PersQueue_V1.StreamingWriteServerMessage, error)
	grpc.ClientStream
}

type persQueueServiceStreamingWriteClient struct {
	grpc.ClientStream
}

func (x *persQueueServiceStreamingWriteClient) Send(m *Ydb_PersQueue_V1.StreamingWriteClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *persQueueServiceStreamingWriteClient) Recv() (*Ydb_PersQueue_V1.StreamingWriteServerMessage, error) {
	m := new(Ydb_PersQueue_V1.StreamingWriteServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *persQueueServiceClient) MigrationStreamingRead(ctx context.Context, opts ...grpc.CallOption) (PersQueueService_MigrationStreamingReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersQueueService_ServiceDesc.Streams[1], "/Ydb.PersQueue.V1.PersQueueService/MigrationStreamingRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &persQueueServiceMigrationStreamingReadClient{stream}
	return x, nil
}

type PersQueueService_MigrationStreamingReadClient interface {
	Send(*Ydb_PersQueue_V1.MigrationStreamingReadClientMessage) error
	Recv() (*Ydb_PersQueue_V1.MigrationStreamingReadServerMessage, error)
	grpc.ClientStream
}

type persQueueServiceMigrationStreamingReadClient struct {
	grpc.ClientStream
}

func (x *persQueueServiceMigrationStreamingReadClient) Send(m *Ydb_PersQueue_V1.MigrationStreamingReadClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *persQueueServiceMigrationStreamingReadClient) Recv() (*Ydb_PersQueue_V1.MigrationStreamingReadServerMessage, error) {
	m := new(Ydb_PersQueue_V1.MigrationStreamingReadServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *persQueueServiceClient) GetReadSessionsInfo(ctx context.Context, in *Ydb_PersQueue_V1.ReadInfoRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.ReadInfoResponse, error) {
	out := new(Ydb_PersQueue_V1.ReadInfoResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/GetReadSessionsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) DescribeTopic(ctx context.Context, in *Ydb_PersQueue_V1.DescribeTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.DescribeTopicResponse, error) {
	out := new(Ydb_PersQueue_V1.DescribeTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/DescribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) DropTopic(ctx context.Context, in *Ydb_PersQueue_V1.DropTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.DropTopicResponse, error) {
	out := new(Ydb_PersQueue_V1.DropTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/DropTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) CreateTopic(ctx context.Context, in *Ydb_PersQueue_V1.CreateTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.CreateTopicResponse, error) {
	out := new(Ydb_PersQueue_V1.CreateTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) AlterTopic(ctx context.Context, in *Ydb_PersQueue_V1.AlterTopicRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.AlterTopicResponse, error) {
	out := new(Ydb_PersQueue_V1.AlterTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/AlterTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) AddReadRule(ctx context.Context, in *Ydb_PersQueue_V1.AddReadRuleRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.AddReadRuleResponse, error) {
	out := new(Ydb_PersQueue_V1.AddReadRuleResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/AddReadRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persQueueServiceClient) RemoveReadRule(ctx context.Context, in *Ydb_PersQueue_V1.RemoveReadRuleRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_V1.RemoveReadRuleResponse, error) {
	out := new(Ydb_PersQueue_V1.RemoveReadRuleResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.PersQueueService/RemoveReadRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersQueueServiceServer is the server API for PersQueueService service.
// All implementations must embed UnimplementedPersQueueServiceServer
// for forward compatibility
type PersQueueServiceServer interface {
	StreamingWrite(PersQueueService_StreamingWriteServer) error
	//*
	// Creates Read Session
	// Pipeline:
	// client                  server
	//         Init(Topics, ClientId, ...)
	//        ---------------->
	//         Init(SessionId)
	//        <----------------
	//         read1
	//        ---------------->
	//         read2
	//        ---------------->
	//         assign(Topic1, Cluster, Partition1, ...) - assigns and releases are optional
	//        <----------------
	//         assign(Topic2, Clutster, Partition2, ...)
	//        <----------------
	//         start_read(Topic1, Partition1, ...) - client must respond to assign request with this message. Only after this client will start recieving messages from this partition
	//        ---------------->
	//         release(Topic1, Partition1, ...)
	//        <----------------
	//         released(Topic1, Partition1, ...) - only after released server will give this parittion to other session.
	//        ---------------->
	//         start_read(Topic2, Partition2, ...) - client must respond to assign request with this message. Only after this client will start recieving messages from this partition
	//        ---------------->
	//         read data(data, ...)
	//        <----------------
	//         commit(cookie1)
	//        ---------------->
	//         committed(cookie1)
	//        <----------------
	//         issue(description, ...)
	//        <----------------
	MigrationStreamingRead(PersQueueService_MigrationStreamingReadServer) error
	// Get information about reading
	GetReadSessionsInfo(context.Context, *Ydb_PersQueue_V1.ReadInfoRequest) (*Ydb_PersQueue_V1.ReadInfoResponse, error)
	//
	// Describe topic command.
	DescribeTopic(context.Context, *Ydb_PersQueue_V1.DescribeTopicRequest) (*Ydb_PersQueue_V1.DescribeTopicResponse, error)
	//
	// Drop topic command.
	DropTopic(context.Context, *Ydb_PersQueue_V1.DropTopicRequest) (*Ydb_PersQueue_V1.DropTopicResponse, error)
	//
	// Create topic command.
	CreateTopic(context.Context, *Ydb_PersQueue_V1.CreateTopicRequest) (*Ydb_PersQueue_V1.CreateTopicResponse, error)
	//
	// Alter topic command.
	AlterTopic(context.Context, *Ydb_PersQueue_V1.AlterTopicRequest) (*Ydb_PersQueue_V1.AlterTopicResponse, error)
	//
	// Add read rule command.
	AddReadRule(context.Context, *Ydb_PersQueue_V1.AddReadRuleRequest) (*Ydb_PersQueue_V1.AddReadRuleResponse, error)
	//
	// Remove read rule command.
	RemoveReadRule(context.Context, *Ydb_PersQueue_V1.RemoveReadRuleRequest) (*Ydb_PersQueue_V1.RemoveReadRuleResponse, error)
	mustEmbedUnimplementedPersQueueServiceServer()
}

// UnimplementedPersQueueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersQueueServiceServer struct {
}

func (UnimplementedPersQueueServiceServer) StreamingWrite(PersQueueService_StreamingWriteServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingWrite not implemented")
}
func (UnimplementedPersQueueServiceServer) MigrationStreamingRead(PersQueueService_MigrationStreamingReadServer) error {
	return status.Errorf(codes.Unimplemented, "method MigrationStreamingRead not implemented")
}
func (UnimplementedPersQueueServiceServer) GetReadSessionsInfo(context.Context, *Ydb_PersQueue_V1.ReadInfoRequest) (*Ydb_PersQueue_V1.ReadInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadSessionsInfo not implemented")
}
func (UnimplementedPersQueueServiceServer) DescribeTopic(context.Context, *Ydb_PersQueue_V1.DescribeTopicRequest) (*Ydb_PersQueue_V1.DescribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTopic not implemented")
}
func (UnimplementedPersQueueServiceServer) DropTopic(context.Context, *Ydb_PersQueue_V1.DropTopicRequest) (*Ydb_PersQueue_V1.DropTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTopic not implemented")
}
func (UnimplementedPersQueueServiceServer) CreateTopic(context.Context, *Ydb_PersQueue_V1.CreateTopicRequest) (*Ydb_PersQueue_V1.CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedPersQueueServiceServer) AlterTopic(context.Context, *Ydb_PersQueue_V1.AlterTopicRequest) (*Ydb_PersQueue_V1.AlterTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterTopic not implemented")
}
func (UnimplementedPersQueueServiceServer) AddReadRule(context.Context, *Ydb_PersQueue_V1.AddReadRuleRequest) (*Ydb_PersQueue_V1.AddReadRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReadRule not implemented")
}
func (UnimplementedPersQueueServiceServer) RemoveReadRule(context.Context, *Ydb_PersQueue_V1.RemoveReadRuleRequest) (*Ydb_PersQueue_V1.RemoveReadRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReadRule not implemented")
}
func (UnimplementedPersQueueServiceServer) mustEmbedUnimplementedPersQueueServiceServer() {}

// UnsafePersQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersQueueServiceServer will
// result in compilation errors.
type UnsafePersQueueServiceServer interface {
	mustEmbedUnimplementedPersQueueServiceServer()
}

func RegisterPersQueueServiceServer(s grpc.ServiceRegistrar, srv PersQueueServiceServer) {
	s.RegisterService(&PersQueueService_ServiceDesc, srv)
}

func _PersQueueService_StreamingWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersQueueServiceServer).StreamingWrite(&persQueueServiceStreamingWriteServer{stream})
}

type PersQueueService_StreamingWriteServer interface {
	Send(*Ydb_PersQueue_V1.StreamingWriteServerMessage) error
	Recv() (*Ydb_PersQueue_V1.StreamingWriteClientMessage, error)
	grpc.ServerStream
}

type persQueueServiceStreamingWriteServer struct {
	grpc.ServerStream
}

func (x *persQueueServiceStreamingWriteServer) Send(m *Ydb_PersQueue_V1.StreamingWriteServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *persQueueServiceStreamingWriteServer) Recv() (*Ydb_PersQueue_V1.StreamingWriteClientMessage, error) {
	m := new(Ydb_PersQueue_V1.StreamingWriteClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersQueueService_MigrationStreamingRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersQueueServiceServer).MigrationStreamingRead(&persQueueServiceMigrationStreamingReadServer{stream})
}

type PersQueueService_MigrationStreamingReadServer interface {
	Send(*Ydb_PersQueue_V1.MigrationStreamingReadServerMessage) error
	Recv() (*Ydb_PersQueue_V1.MigrationStreamingReadClientMessage, error)
	grpc.ServerStream
}

type persQueueServiceMigrationStreamingReadServer struct {
	grpc.ServerStream
}

func (x *persQueueServiceMigrationStreamingReadServer) Send(m *Ydb_PersQueue_V1.MigrationStreamingReadServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *persQueueServiceMigrationStreamingReadServer) Recv() (*Ydb_PersQueue_V1.MigrationStreamingReadClientMessage, error) {
	m := new(Ydb_PersQueue_V1.MigrationStreamingReadClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersQueueService_GetReadSessionsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.ReadInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).GetReadSessionsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/GetReadSessionsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).GetReadSessionsInfo(ctx, req.(*Ydb_PersQueue_V1.ReadInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_DescribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.DescribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).DescribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/DescribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).DescribeTopic(ctx, req.(*Ydb_PersQueue_V1.DescribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_DropTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.DropTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).DropTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/DropTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).DropTopic(ctx, req.(*Ydb_PersQueue_V1.DropTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).CreateTopic(ctx, req.(*Ydb_PersQueue_V1.CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_AlterTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.AlterTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).AlterTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/AlterTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).AlterTopic(ctx, req.(*Ydb_PersQueue_V1.AlterTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_AddReadRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.AddReadRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).AddReadRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/AddReadRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).AddReadRule(ctx, req.(*Ydb_PersQueue_V1.AddReadRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersQueueService_RemoveReadRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_V1.RemoveReadRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersQueueServiceServer).RemoveReadRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.PersQueueService/RemoveReadRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersQueueServiceServer).RemoveReadRule(ctx, req.(*Ydb_PersQueue_V1.RemoveReadRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PersQueueService_ServiceDesc is the grpc.ServiceDesc for PersQueueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersQueueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.PersQueue.V1.PersQueueService",
	HandlerType: (*PersQueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReadSessionsInfo",
			Handler:    _PersQueueService_GetReadSessionsInfo_Handler,
		},
		{
			MethodName: "DescribeTopic",
			Handler:    _PersQueueService_DescribeTopic_Handler,
		},
		{
			MethodName: "DropTopic",
			Handler:    _PersQueueService_DropTopic_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _PersQueueService_CreateTopic_Handler,
		},
		{
			MethodName: "AlterTopic",
			Handler:    _PersQueueService_AlterTopic_Handler,
		},
		{
			MethodName: "AddReadRule",
			Handler:    _PersQueueService_AddReadRule_Handler,
		},
		{
			MethodName: "RemoveReadRule",
			Handler:    _PersQueueService_RemoveReadRule_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingWrite",
			Handler:       _PersQueueService_StreamingWrite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MigrationStreamingRead",
			Handler:       _PersQueueService_MigrationStreamingRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ydb_persqueue_v1.proto",
}

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	// Get PQ clusters which are eligible for the specified Write or Read Sessions
	DiscoverClusters(ctx context.Context, in *Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterDiscoveryServiceClient(cc grpc.ClientConnInterface) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) DiscoverClusters(ctx context.Context, in *Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest, opts ...grpc.CallOption) (*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse, error) {
	out := new(Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse)
	err := c.cc.Invoke(ctx, "/Ydb.PersQueue.V1.ClusterDiscoveryService/DiscoverClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterDiscoveryServiceServer is the server API for ClusterDiscoveryService service.
// All implementations must embed UnimplementedClusterDiscoveryServiceServer
// for forward compatibility
type ClusterDiscoveryServiceServer interface {
	// Get PQ clusters which are eligible for the specified Write or Read Sessions
	DiscoverClusters(context.Context, *Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest) (*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse, error)
	mustEmbedUnimplementedClusterDiscoveryServiceServer()
}

// UnimplementedClusterDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterDiscoveryServiceServer struct {
}

func (UnimplementedClusterDiscoveryServiceServer) DiscoverClusters(context.Context, *Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest) (*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverClusters not implemented")
}
func (UnimplementedClusterDiscoveryServiceServer) mustEmbedUnimplementedClusterDiscoveryServiceServer() {
}

// UnsafeClusterDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterDiscoveryServiceServer will
// result in compilation errors.
type UnsafeClusterDiscoveryServiceServer interface {
	mustEmbedUnimplementedClusterDiscoveryServiceServer()
}

func RegisterClusterDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&ClusterDiscoveryService_ServiceDesc, srv)
}

func _ClusterDiscoveryService_DiscoverClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).DiscoverClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.PersQueue.V1.ClusterDiscoveryService/DiscoverClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).DiscoverClusters(ctx, req.(*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ClusterDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.PersQueue.V1.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverClusters",
			Handler:    _ClusterDiscoveryService_DiscoverClusters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_persqueue_v1.proto",
}
