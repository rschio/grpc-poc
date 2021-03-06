// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackerClient interface {
	// This comment is used to document the stub generated by gRPC
	// and also the openAPI auto generated.
	//
	// Login authenticates a user and return an access token.
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error)
	// Search searches by a tracking code and returns it's events.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// SearchBatch searches one or more tracking codes and returns a stream of
	// SearchResponses.
	SearchBatch(ctx context.Context, in *SearchBatchRequest, opts ...grpc.CallOption) (Tracker_SearchBatchClient, error)
}

type trackerClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackerClient(cc grpc.ClientConnInterface) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.Tracker/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/proto.Tracker/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) SearchBatch(ctx context.Context, in *SearchBatchRequest, opts ...grpc.CallOption) (Tracker_SearchBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tracker_ServiceDesc.Streams[0], "/proto.Tracker/SearchBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackerSearchBatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tracker_SearchBatchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type trackerSearchBatchClient struct {
	grpc.ClientStream
}

func (x *trackerSearchBatchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrackerServer is the server API for Tracker service.
// All implementations must embed UnimplementedTrackerServer
// for forward compatibility
type TrackerServer interface {
	// This comment is used to document the stub generated by gRPC
	// and also the openAPI auto generated.
	//
	// Login authenticates a user and return an access token.
	Login(context.Context, *Credentials) (*Token, error)
	// Search searches by a tracking code and returns it's events.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// SearchBatch searches one or more tracking codes and returns a stream of
	// SearchResponses.
	SearchBatch(*SearchBatchRequest, Tracker_SearchBatchServer) error
	mustEmbedUnimplementedTrackerServer()
}

// UnimplementedTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedTrackerServer struct {
}

func (UnimplementedTrackerServer) Login(context.Context, *Credentials) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedTrackerServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedTrackerServer) SearchBatch(*SearchBatchRequest, Tracker_SearchBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchBatch not implemented")
}
func (UnimplementedTrackerServer) mustEmbedUnimplementedTrackerServer() {}

// UnsafeTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackerServer will
// result in compilation errors.
type UnsafeTrackerServer interface {
	mustEmbedUnimplementedTrackerServer()
}

func RegisterTrackerServer(s grpc.ServiceRegistrar, srv TrackerServer) {
	s.RegisterService(&Tracker_ServiceDesc, srv)
}

func _Tracker_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Tracker/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Tracker/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_SearchBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchBatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackerServer).SearchBatch(m, &trackerSearchBatchServer{stream})
}

type Tracker_SearchBatchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type trackerSearchBatchServer struct {
	grpc.ServerStream
}

func (x *trackerSearchBatchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Tracker_ServiceDesc is the grpc.ServiceDesc for Tracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Tracker_Login_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Tracker_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchBatch",
			Handler:       _Tracker_SearchBatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tracker.proto",
}
