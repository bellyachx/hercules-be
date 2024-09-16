// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package exercisepb

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

// ExerciseServiceClient is the client API for ExerciseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseServiceClient interface {
	CreateExercise(ctx context.Context, in *Exercise, opts ...grpc.CallOption) (*ExerciseCreatedResponse, error)
	GetExercises(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetExercisesResponse, error)
}

type exerciseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseServiceClient(cc grpc.ClientConnInterface) ExerciseServiceClient {
	return &exerciseServiceClient{cc}
}

func (c *exerciseServiceClient) CreateExercise(ctx context.Context, in *Exercise, opts ...grpc.CallOption) (*ExerciseCreatedResponse, error) {
	out := new(ExerciseCreatedResponse)
	err := c.cc.Invoke(ctx, "/api.ExerciseService/CreateExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) GetExercises(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetExercisesResponse, error) {
	out := new(GetExercisesResponse)
	err := c.cc.Invoke(ctx, "/api.ExerciseService/GetExercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseServiceServer is the server API for ExerciseService service.
// All implementations must embed UnimplementedExerciseServiceServer
// for forward compatibility
type ExerciseServiceServer interface {
	CreateExercise(context.Context, *Exercise) (*ExerciseCreatedResponse, error)
	GetExercises(context.Context, *emptypb.Empty) (*GetExercisesResponse, error)
	mustEmbedUnimplementedExerciseServiceServer()
}

// UnimplementedExerciseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExerciseServiceServer struct {
}

func (UnimplementedExerciseServiceServer) CreateExercise(context.Context, *Exercise) (*ExerciseCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExercise not implemented")
}
func (UnimplementedExerciseServiceServer) GetExercises(context.Context, *emptypb.Empty) (*GetExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExercises not implemented")
}
func (UnimplementedExerciseServiceServer) mustEmbedUnimplementedExerciseServiceServer() {}

// UnsafeExerciseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseServiceServer will
// result in compilation errors.
type UnsafeExerciseServiceServer interface {
	mustEmbedUnimplementedExerciseServiceServer()
}

func RegisterExerciseServiceServer(s grpc.ServiceRegistrar, srv ExerciseServiceServer) {
	s.RegisterService(&ExerciseService_ServiceDesc, srv)
}

func _ExerciseService_CreateExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Exercise)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExerciseService/CreateExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, req.(*Exercise))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_GetExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).GetExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExerciseService/GetExercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).GetExercises(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseService_ServiceDesc is the grpc.ServiceDesc for ExerciseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ExerciseService",
	HandlerType: (*ExerciseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExercise",
			Handler:    _ExerciseService_CreateExercise_Handler,
		},
		{
			MethodName: "GetExercises",
			Handler:    _ExerciseService_GetExercises_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exercise.proto",
}
