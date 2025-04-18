// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: dbpb/dbpb.proto

package dbpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Postgres_CreateTask_FullMethodName = "/example.Postgres/CreateTask"
	Postgres_GetTask_FullMethodName    = "/example.Postgres/GetTask"
	Postgres_GetTasks_FullMethodName   = "/example.Postgres/GetTasks"
	Postgres_UpdateTask_FullMethodName = "/example.Postgres/UpdateTask"
	Postgres_DeleteTask_FullMethodName = "/example.Postgres/DeleteTask"
	Postgres_MarkAsDone_FullMethodName = "/example.Postgres/MarkAsDone"
)

// PostgresClient is the client API for Postgres service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClient interface {
	// Создать задачу
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Получить задачу по IP
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Получение всех задач
	GetTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Task], error)
	// обновить задачу
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// Удалить задачу
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Отметить задачу как выполненную
	MarkAsDone(ctx context.Context, in *MarkAsDoneRequest, opts ...grpc.CallOption) (*Task, error)
}

type postgresClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClient(cc grpc.ClientConnInterface) PostgresClient {
	return &postgresClient{cc}
}

func (c *postgresClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, Postgres_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, Postgres_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClient) GetTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Task], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Postgres_ServiceDesc.Streams[0], Postgres_GetTasks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, Task]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Postgres_GetTasksClient = grpc.ServerStreamingClient[Task]

func (c *postgresClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, Postgres_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Postgres_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClient) MarkAsDone(ctx context.Context, in *MarkAsDoneRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, Postgres_MarkAsDone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresServer is the server API for Postgres service.
// All implementations must embed UnimplementedPostgresServer
// for forward compatibility.
type PostgresServer interface {
	// Создать задачу
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	// Получить задачу по IP
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	// Получение всех задач
	GetTasks(*emptypb.Empty, grpc.ServerStreamingServer[Task]) error
	// обновить задачу
	UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error)
	// Удалить задачу
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	// Отметить задачу как выполненную
	MarkAsDone(context.Context, *MarkAsDoneRequest) (*Task, error)
	mustEmbedUnimplementedPostgresServer()
}

// UnimplementedPostgresServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostgresServer struct{}

func (UnimplementedPostgresServer) CreateTask(context.Context, *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedPostgresServer) GetTask(context.Context, *GetTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedPostgresServer) GetTasks(*emptypb.Empty, grpc.ServerStreamingServer[Task]) error {
	return status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedPostgresServer) UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedPostgresServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedPostgresServer) MarkAsDone(context.Context, *MarkAsDoneRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsDone not implemented")
}
func (UnimplementedPostgresServer) mustEmbedUnimplementedPostgresServer() {}
func (UnimplementedPostgresServer) testEmbeddedByValue()                  {}

// UnsafePostgresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresServer will
// result in compilation errors.
type UnsafePostgresServer interface {
	mustEmbedUnimplementedPostgresServer()
}

func RegisterPostgresServer(s grpc.ServiceRegistrar, srv PostgresServer) {
	// If the following call pancis, it indicates UnimplementedPostgresServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Postgres_ServiceDesc, srv)
}

func _Postgres_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Postgres_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Postgres_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Postgres_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Postgres_GetTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostgresServer).GetTasks(m, &grpc.GenericServerStream[emptypb.Empty, Task]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Postgres_GetTasksServer = grpc.ServerStreamingServer[Task]

func _Postgres_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Postgres_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Postgres_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Postgres_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Postgres_MarkAsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServer).MarkAsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Postgres_MarkAsDone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServer).MarkAsDone(ctx, req.(*MarkAsDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Postgres_ServiceDesc is the grpc.ServiceDesc for Postgres service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Postgres_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Postgres",
	HandlerType: (*PostgresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _Postgres_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Postgres_GetTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Postgres_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Postgres_DeleteTask_Handler,
		},
		{
			MethodName: "MarkAsDone",
			Handler:    _Postgres_MarkAsDone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTasks",
			Handler:       _Postgres_GetTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dbpb/dbpb.proto",
}
