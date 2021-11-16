// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationClient interface {
	// Sends a greeting for a new server participant
	NewServer(ctx context.Context, in *NewServerRequest, opts ...grpc.CallOption) (*NewServerReply, error)
	// Sends a greeting for a new frontend participant
	NewFrontEnd(ctx context.Context, in *NewFrontEndRequest, opts ...grpc.CallOption) (*NewFrontEndReply, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) NewServer(ctx context.Context, in *NewServerRequest, opts ...grpc.CallOption) (*NewServerReply, error) {
	out := new(NewServerReply)
	err := c.cc.Invoke(ctx, "/communication.Replication/NewServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) NewFrontEnd(ctx context.Context, in *NewFrontEndRequest, opts ...grpc.CallOption) (*NewFrontEndReply, error) {
	out := new(NewFrontEndReply)
	err := c.cc.Invoke(ctx, "/communication.Replication/NewFrontEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations must embed UnimplementedReplicationServer
// for forward compatibility
type ReplicationServer interface {
	// Sends a greeting for a new server participant
	NewServer(context.Context, *NewServerRequest) (*NewServerReply, error)
	// Sends a greeting for a new frontend participant
	NewFrontEnd(context.Context, *NewFrontEndRequest) (*NewFrontEndReply, error)
	mustEmbedUnimplementedReplicationServer()
}

// UnimplementedReplicationServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (UnimplementedReplicationServer) NewServer(context.Context, *NewServerRequest) (*NewServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewServer not implemented")
}
func (UnimplementedReplicationServer) NewFrontEnd(context.Context, *NewFrontEndRequest) (*NewFrontEndReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewFrontEnd not implemented")
}
func (UnimplementedReplicationServer) mustEmbedUnimplementedReplicationServer() {}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	s.RegisterService(&Replication_ServiceDesc, srv)
}

func _Replication_NewServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).NewServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.Replication/NewServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).NewServer(ctx, req.(*NewServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_NewFrontEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewFrontEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).NewFrontEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.Replication/NewFrontEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).NewFrontEnd(ctx, req.(*NewFrontEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Replication_ServiceDesc is the grpc.ServiceDesc for Replication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communication.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewServer",
			Handler:    _Replication_NewServer_Handler,
		},
		{
			MethodName: "NewFrontEnd",
			Handler:    _Replication_NewFrontEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Replication/protobuf/communication.proto",
}
