// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: ydb_topic_v1.proto

package Ydb_Topic_V1

import (
	context "context"
	Ydb_Topic "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Topic"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	// Create Write Session
	// Pipeline example:
	// client                  server
	//         InitRequest(Topic, MessageGroupID, ...)
	//        ---------------->
	//         InitResponse(Partition, MaxSeqNo, ...)
	//        <----------------
	//         WriteRequest(data1, seqNo1)
	//        ---------------->
	//         WriteRequest(data2, seqNo2)
	//        ---------------->
	//         WriteResponse(seqNo1, offset1, ...)
	//        <----------------
	//         WriteRequest(data3, seqNo3)
	//        ---------------->
	//         WriteResponse(seqNo2, offset2, ...)
	//        <----------------
	//         [something went wrong] (status != SUCCESS, issues not empty)
	//        <----------------
	StreamWrite(ctx context.Context, opts ...grpc.CallOption) (TopicService_StreamWriteClient, error)
	// Create Read Session
	// Pipeline:
	// client                  server
	//         InitRequest(Topics, ClientId, ...)
	//        ---------------->
	//         InitResponse(SessionId)
	//        <----------------
	//         ReadRequest
	//        ---------------->
	//         ReadRequest
	//        ---------------->
	//         StartPartitionSessionRequest(Topic1, Partition1, PartitionSessionID1, ...)
	//        <----------------
	//         StartPartitionSessionRequest(Topic2, Partition2, PartitionSessionID2, ...)
	//        <----------------
	//         StartPartitionSessionResponse(PartitionSessionID1, ...)
	//             client must respond with this message to actually start recieving data messages from this partition
	//        ---------------->
	//         StopPartitionSessionRequest(PartitionSessionID1, ...)
	//        <----------------
	//         StopPartitionSessionResponse(PartitionSessionID1, ...)
	//             only after this response server will give this parittion to other session.
	//        ---------------->
	//         StartPartitionSessionResponse(PartitionSession2, ...)
	//        ---------------->
	//         ReadResponse(data, ...)
	//        <----------------
	//         CommitRequest(PartitionCommit1, ...)
	//        ---------------->
	//         CommitResponse(PartitionCommitAck1, ...)
	//        <----------------
	//         [something went wrong] (status != SUCCESS, issues not empty)
	//        <----------------
	StreamRead(ctx context.Context, opts ...grpc.CallOption) (TopicService_StreamReadClient, error)
	// Create topic command.
	CreateTopic(ctx context.Context, in *Ydb_Topic.CreateTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.CreateTopicResponse, error)
	// Describe topic command.
	DescribeTopic(ctx context.Context, in *Ydb_Topic.DescribeTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.DescribeTopicResponse, error)
	// Alter topic command.
	AlterTopic(ctx context.Context, in *Ydb_Topic.AlterTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.AlterTopicResponse, error)
	// Drop topic command.
	DropTopic(ctx context.Context, in *Ydb_Topic.DropTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.DropTopicResponse, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) StreamWrite(ctx context.Context, opts ...grpc.CallOption) (TopicService_StreamWriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &TopicService_ServiceDesc.Streams[0], "/Ydb.Topic.V1.TopicService/StreamWrite", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicServiceStreamWriteClient{stream}
	return x, nil
}

type TopicService_StreamWriteClient interface {
	Send(*Ydb_Topic.StreamWriteMessage_FromClient) error
	Recv() (*Ydb_Topic.StreamWriteMessage_FromServer, error)
	grpc.ClientStream
}

type topicServiceStreamWriteClient struct {
	grpc.ClientStream
}

func (x *topicServiceStreamWriteClient) Send(m *Ydb_Topic.StreamWriteMessage_FromClient) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topicServiceStreamWriteClient) Recv() (*Ydb_Topic.StreamWriteMessage_FromServer, error) {
	m := new(Ydb_Topic.StreamWriteMessage_FromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicServiceClient) StreamRead(ctx context.Context, opts ...grpc.CallOption) (TopicService_StreamReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &TopicService_ServiceDesc.Streams[1], "/Ydb.Topic.V1.TopicService/StreamRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicServiceStreamReadClient{stream}
	return x, nil
}

type TopicService_StreamReadClient interface {
	Send(*Ydb_Topic.StreamReadMessage_FromClient) error
	Recv() (*Ydb_Topic.StreamReadMessage_FromServer, error)
	grpc.ClientStream
}

type topicServiceStreamReadClient struct {
	grpc.ClientStream
}

func (x *topicServiceStreamReadClient) Send(m *Ydb_Topic.StreamReadMessage_FromClient) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topicServiceStreamReadClient) Recv() (*Ydb_Topic.StreamReadMessage_FromServer, error) {
	m := new(Ydb_Topic.StreamReadMessage_FromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *Ydb_Topic.CreateTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.CreateTopicResponse, error) {
	out := new(Ydb_Topic.CreateTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Topic.V1.TopicService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DescribeTopic(ctx context.Context, in *Ydb_Topic.DescribeTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.DescribeTopicResponse, error) {
	out := new(Ydb_Topic.DescribeTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Topic.V1.TopicService/DescribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) AlterTopic(ctx context.Context, in *Ydb_Topic.AlterTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.AlterTopicResponse, error) {
	out := new(Ydb_Topic.AlterTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Topic.V1.TopicService/AlterTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DropTopic(ctx context.Context, in *Ydb_Topic.DropTopicRequest, opts ...grpc.CallOption) (*Ydb_Topic.DropTopicResponse, error) {
	out := new(Ydb_Topic.DropTopicResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Topic.V1.TopicService/DropTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations must embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	// Create Write Session
	// Pipeline example:
	// client                  server
	//         InitRequest(Topic, MessageGroupID, ...)
	//        ---------------->
	//         InitResponse(Partition, MaxSeqNo, ...)
	//        <----------------
	//         WriteRequest(data1, seqNo1)
	//        ---------------->
	//         WriteRequest(data2, seqNo2)
	//        ---------------->
	//         WriteResponse(seqNo1, offset1, ...)
	//        <----------------
	//         WriteRequest(data3, seqNo3)
	//        ---------------->
	//         WriteResponse(seqNo2, offset2, ...)
	//        <----------------
	//         [something went wrong] (status != SUCCESS, issues not empty)
	//        <----------------
	StreamWrite(TopicService_StreamWriteServer) error
	// Create Read Session
	// Pipeline:
	// client                  server
	//         InitRequest(Topics, ClientId, ...)
	//        ---------------->
	//         InitResponse(SessionId)
	//        <----------------
	//         ReadRequest
	//        ---------------->
	//         ReadRequest
	//        ---------------->
	//         StartPartitionSessionRequest(Topic1, Partition1, PartitionSessionID1, ...)
	//        <----------------
	//         StartPartitionSessionRequest(Topic2, Partition2, PartitionSessionID2, ...)
	//        <----------------
	//         StartPartitionSessionResponse(PartitionSessionID1, ...)
	//             client must respond with this message to actually start recieving data messages from this partition
	//        ---------------->
	//         StopPartitionSessionRequest(PartitionSessionID1, ...)
	//        <----------------
	//         StopPartitionSessionResponse(PartitionSessionID1, ...)
	//             only after this response server will give this parittion to other session.
	//        ---------------->
	//         StartPartitionSessionResponse(PartitionSession2, ...)
	//        ---------------->
	//         ReadResponse(data, ...)
	//        <----------------
	//         CommitRequest(PartitionCommit1, ...)
	//        ---------------->
	//         CommitResponse(PartitionCommitAck1, ...)
	//        <----------------
	//         [something went wrong] (status != SUCCESS, issues not empty)
	//        <----------------
	StreamRead(TopicService_StreamReadServer) error
	// Create topic command.
	CreateTopic(context.Context, *Ydb_Topic.CreateTopicRequest) (*Ydb_Topic.CreateTopicResponse, error)
	// Describe topic command.
	DescribeTopic(context.Context, *Ydb_Topic.DescribeTopicRequest) (*Ydb_Topic.DescribeTopicResponse, error)
	// Alter topic command.
	AlterTopic(context.Context, *Ydb_Topic.AlterTopicRequest) (*Ydb_Topic.AlterTopicResponse, error)
	// Drop topic command.
	DropTopic(context.Context, *Ydb_Topic.DropTopicRequest) (*Ydb_Topic.DropTopicResponse, error)
	mustEmbedUnimplementedTopicServiceServer()
}

// UnimplementedTopicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) StreamWrite(TopicService_StreamWriteServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamWrite not implemented")
}
func (UnimplementedTopicServiceServer) StreamRead(TopicService_StreamReadServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRead not implemented")
}
func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *Ydb_Topic.CreateTopicRequest) (*Ydb_Topic.CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) DescribeTopic(context.Context, *Ydb_Topic.DescribeTopicRequest) (*Ydb_Topic.DescribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTopic not implemented")
}
func (UnimplementedTopicServiceServer) AlterTopic(context.Context, *Ydb_Topic.AlterTopicRequest) (*Ydb_Topic.AlterTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterTopic not implemented")
}
func (UnimplementedTopicServiceServer) DropTopic(context.Context, *Ydb_Topic.DropTopicRequest) (*Ydb_Topic.DropTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTopic not implemented")
}
func (UnimplementedTopicServiceServer) mustEmbedUnimplementedTopicServiceServer() {}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_StreamWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopicServiceServer).StreamWrite(&topicServiceStreamWriteServer{stream})
}

type TopicService_StreamWriteServer interface {
	Send(*Ydb_Topic.StreamWriteMessage_FromServer) error
	Recv() (*Ydb_Topic.StreamWriteMessage_FromClient, error)
	grpc.ServerStream
}

type topicServiceStreamWriteServer struct {
	grpc.ServerStream
}

func (x *topicServiceStreamWriteServer) Send(m *Ydb_Topic.StreamWriteMessage_FromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topicServiceStreamWriteServer) Recv() (*Ydb_Topic.StreamWriteMessage_FromClient, error) {
	m := new(Ydb_Topic.StreamWriteMessage_FromClient)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TopicService_StreamRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopicServiceServer).StreamRead(&topicServiceStreamReadServer{stream})
}

type TopicService_StreamReadServer interface {
	Send(*Ydb_Topic.StreamReadMessage_FromServer) error
	Recv() (*Ydb_Topic.StreamReadMessage_FromClient, error)
	grpc.ServerStream
}

type topicServiceStreamReadServer struct {
	grpc.ServerStream
}

func (x *topicServiceStreamReadServer) Send(m *Ydb_Topic.StreamReadMessage_FromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topicServiceStreamReadServer) Recv() (*Ydb_Topic.StreamReadMessage_FromClient, error) {
	m := new(Ydb_Topic.StreamReadMessage_FromClient)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Topic.CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Topic.V1.TopicService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*Ydb_Topic.CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DescribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Topic.DescribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DescribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Topic.V1.TopicService/DescribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DescribeTopic(ctx, req.(*Ydb_Topic.DescribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_AlterTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Topic.AlterTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).AlterTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Topic.V1.TopicService/AlterTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).AlterTopic(ctx, req.(*Ydb_Topic.AlterTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DropTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Topic.DropTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DropTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Topic.V1.TopicService/DropTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DropTopic(ctx, req.(*Ydb_Topic.DropTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Topic.V1.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "DescribeTopic",
			Handler:    _TopicService_DescribeTopic_Handler,
		},
		{
			MethodName: "AlterTopic",
			Handler:    _TopicService_AlterTopic_Handler,
		},
		{
			MethodName: "DropTopic",
			Handler:    _TopicService_DropTopic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWrite",
			Handler:       _TopicService_StreamWrite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamRead",
			Handler:       _TopicService_StreamRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ydb_topic_v1.proto",
}
