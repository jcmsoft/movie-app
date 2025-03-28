// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: services.proto

package servicespb

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
	ClientProtoService_GetClientById_FullMethodName    = "/services.ClientProtoService/GetClientById"
	ClientProtoService_GetClientByEmail_FullMethodName = "/services.ClientProtoService/GetClientByEmail"
	ClientProtoService_IsClientExist_FullMethodName    = "/services.ClientProtoService/IsClientExist"
	ClientProtoService_AddClient_FullMethodName        = "/services.ClientProtoService/AddClient"
	ClientProtoService_UpdateClient_FullMethodName     = "/services.ClientProtoService/UpdateClient"
	ClientProtoService_DeleteClient_FullMethodName     = "/services.ClientProtoService/DeleteClient"
	ClientProtoService_GetAllClients_FullMethodName    = "/services.ClientProtoService/GetAllClients"
)

// ClientProtoServiceClient is the client API for ClientProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientProtoServiceClient interface {
	GetClientById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	GetClientByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	IsClientExist(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*IsExistResponse, error)
	AddClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	DeleteClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetAllClients(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ClientResponse], error)
}

type clientProtoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientProtoServiceClient(cc grpc.ClientConnInterface) ClientProtoServiceClient {
	return &clientProtoServiceClient{cc}
}

func (c *clientProtoServiceClient) GetClientById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_GetClientById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) GetClientByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_GetClientByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) IsClientExist(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*IsExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsExistResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_IsClientExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) AddClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_AddClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_UpdateClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) DeleteClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ClientProtoService_DeleteClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientProtoServiceClient) GetAllClients(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ClientResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ClientProtoService_ServiceDesc.Streams[0], ClientProtoService_GetAllClients_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EmptyRequest, ClientResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientProtoService_GetAllClientsClient = grpc.ServerStreamingClient[ClientResponse]

// ClientProtoServiceServer is the server API for ClientProtoService service.
// All implementations must embed UnimplementedClientProtoServiceServer
// for forward compatibility.
type ClientProtoServiceServer interface {
	GetClientById(context.Context, *IdRequest) (*ClientResponse, error)
	GetClientByEmail(context.Context, *EmailRequest) (*ClientResponse, error)
	IsClientExist(context.Context, *EmailRequest) (*IsExistResponse, error)
	AddClient(context.Context, *ClientRequest) (*ClientResponse, error)
	UpdateClient(context.Context, *ClientRequest) (*ClientResponse, error)
	DeleteClient(context.Context, *IdRequest) (*StatusResponse, error)
	GetAllClients(*EmptyRequest, grpc.ServerStreamingServer[ClientResponse]) error
	mustEmbedUnimplementedClientProtoServiceServer()
}

// UnimplementedClientProtoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientProtoServiceServer struct{}

func (UnimplementedClientProtoServiceServer) GetClientById(context.Context, *IdRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientById not implemented")
}
func (UnimplementedClientProtoServiceServer) GetClientByEmail(context.Context, *EmailRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientByEmail not implemented")
}
func (UnimplementedClientProtoServiceServer) IsClientExist(context.Context, *EmailRequest) (*IsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsClientExist not implemented")
}
func (UnimplementedClientProtoServiceServer) AddClient(context.Context, *ClientRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (UnimplementedClientProtoServiceServer) UpdateClient(context.Context, *ClientRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedClientProtoServiceServer) DeleteClient(context.Context, *IdRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientProtoServiceServer) GetAllClients(*EmptyRequest, grpc.ServerStreamingServer[ClientResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAllClients not implemented")
}
func (UnimplementedClientProtoServiceServer) mustEmbedUnimplementedClientProtoServiceServer() {}
func (UnimplementedClientProtoServiceServer) testEmbeddedByValue()                            {}

// UnsafeClientProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientProtoServiceServer will
// result in compilation errors.
type UnsafeClientProtoServiceServer interface {
	mustEmbedUnimplementedClientProtoServiceServer()
}

func RegisterClientProtoServiceServer(s grpc.ServiceRegistrar, srv ClientProtoServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientProtoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientProtoService_ServiceDesc, srv)
}

func _ClientProtoService_GetClientById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).GetClientById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_GetClientById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).GetClientById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_GetClientByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).GetClientByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_GetClientByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).GetClientByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_IsClientExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).IsClientExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_IsClientExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).IsClientExist(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_AddClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).AddClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_UpdateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).UpdateClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientProtoServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientProtoService_DeleteClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientProtoServiceServer).DeleteClient(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientProtoService_GetAllClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientProtoServiceServer).GetAllClients(m, &grpc.GenericServerStream[EmptyRequest, ClientResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientProtoService_GetAllClientsServer = grpc.ServerStreamingServer[ClientResponse]

// ClientProtoService_ServiceDesc is the grpc.ServiceDesc for ClientProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.ClientProtoService",
	HandlerType: (*ClientProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientById",
			Handler:    _ClientProtoService_GetClientById_Handler,
		},
		{
			MethodName: "GetClientByEmail",
			Handler:    _ClientProtoService_GetClientByEmail_Handler,
		},
		{
			MethodName: "IsClientExist",
			Handler:    _ClientProtoService_IsClientExist_Handler,
		},
		{
			MethodName: "AddClient",
			Handler:    _ClientProtoService_AddClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ClientProtoService_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientProtoService_DeleteClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllClients",
			Handler:       _ClientProtoService_GetAllClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}

const (
	MovieProtoService_GetMovieById_FullMethodName    = "/services.MovieProtoService/GetMovieById"
	MovieProtoService_GetMovieByTitle_FullMethodName = "/services.MovieProtoService/GetMovieByTitle"
	MovieProtoService_IsMovieExist_FullMethodName    = "/services.MovieProtoService/IsMovieExist"
	MovieProtoService_AddMovie_FullMethodName        = "/services.MovieProtoService/AddMovie"
	MovieProtoService_UpdateMovie_FullMethodName     = "/services.MovieProtoService/UpdateMovie"
	MovieProtoService_DeleteMovie_FullMethodName     = "/services.MovieProtoService/DeleteMovie"
	MovieProtoService_GetAllMovies_FullMethodName    = "/services.MovieProtoService/GetAllMovies"
)

// MovieProtoServiceClient is the client API for MovieProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieProtoServiceClient interface {
	GetMovieById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	GetMovieByTitle(ctx context.Context, in *TitleRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	IsMovieExist(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IsExistResponse, error)
	AddMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	UpdateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	DeleteMovie(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetAllMovies(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MovieResponse], error)
}

type movieProtoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieProtoServiceClient(cc grpc.ClientConnInterface) MovieProtoServiceClient {
	return &movieProtoServiceClient{cc}
}

func (c *movieProtoServiceClient) GetMovieById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_GetMovieById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) GetMovieByTitle(ctx context.Context, in *TitleRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_GetMovieByTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) IsMovieExist(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IsExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsExistResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_IsMovieExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) AddMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_AddMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) UpdateMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_UpdateMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) DeleteMovie(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, MovieProtoService_DeleteMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieProtoServiceClient) GetAllMovies(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MovieResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MovieProtoService_ServiceDesc.Streams[0], MovieProtoService_GetAllMovies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EmptyRequest, MovieResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MovieProtoService_GetAllMoviesClient = grpc.ServerStreamingClient[MovieResponse]

// MovieProtoServiceServer is the server API for MovieProtoService service.
// All implementations must embed UnimplementedMovieProtoServiceServer
// for forward compatibility.
type MovieProtoServiceServer interface {
	GetMovieById(context.Context, *IdRequest) (*MovieResponse, error)
	GetMovieByTitle(context.Context, *TitleRequest) (*MovieResponse, error)
	IsMovieExist(context.Context, *IdRequest) (*IsExistResponse, error)
	AddMovie(context.Context, *MovieRequest) (*MovieResponse, error)
	UpdateMovie(context.Context, *MovieRequest) (*MovieResponse, error)
	DeleteMovie(context.Context, *IdRequest) (*StatusResponse, error)
	GetAllMovies(*EmptyRequest, grpc.ServerStreamingServer[MovieResponse]) error
	mustEmbedUnimplementedMovieProtoServiceServer()
}

// UnimplementedMovieProtoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMovieProtoServiceServer struct{}

func (UnimplementedMovieProtoServiceServer) GetMovieById(context.Context, *IdRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieById not implemented")
}
func (UnimplementedMovieProtoServiceServer) GetMovieByTitle(context.Context, *TitleRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByTitle not implemented")
}
func (UnimplementedMovieProtoServiceServer) IsMovieExist(context.Context, *IdRequest) (*IsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsMovieExist not implemented")
}
func (UnimplementedMovieProtoServiceServer) AddMovie(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedMovieProtoServiceServer) UpdateMovie(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieProtoServiceServer) DeleteMovie(context.Context, *IdRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieProtoServiceServer) GetAllMovies(*EmptyRequest, grpc.ServerStreamingServer[MovieResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAllMovies not implemented")
}
func (UnimplementedMovieProtoServiceServer) mustEmbedUnimplementedMovieProtoServiceServer() {}
func (UnimplementedMovieProtoServiceServer) testEmbeddedByValue()                           {}

// UnsafeMovieProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieProtoServiceServer will
// result in compilation errors.
type UnsafeMovieProtoServiceServer interface {
	mustEmbedUnimplementedMovieProtoServiceServer()
}

func RegisterMovieProtoServiceServer(s grpc.ServiceRegistrar, srv MovieProtoServiceServer) {
	// If the following call pancis, it indicates UnimplementedMovieProtoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MovieProtoService_ServiceDesc, srv)
}

func _MovieProtoService_GetMovieById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).GetMovieById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_GetMovieById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).GetMovieById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_GetMovieByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).GetMovieByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_GetMovieByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).GetMovieByTitle(ctx, req.(*TitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_IsMovieExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).IsMovieExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_IsMovieExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).IsMovieExist(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_AddMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).AddMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_UpdateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).UpdateMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieProtoServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieProtoService_DeleteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieProtoServiceServer).DeleteMovie(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieProtoService_GetAllMovies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieProtoServiceServer).GetAllMovies(m, &grpc.GenericServerStream[EmptyRequest, MovieResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MovieProtoService_GetAllMoviesServer = grpc.ServerStreamingServer[MovieResponse]

// MovieProtoService_ServiceDesc is the grpc.ServiceDesc for MovieProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.MovieProtoService",
	HandlerType: (*MovieProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieById",
			Handler:    _MovieProtoService_GetMovieById_Handler,
		},
		{
			MethodName: "GetMovieByTitle",
			Handler:    _MovieProtoService_GetMovieByTitle_Handler,
		},
		{
			MethodName: "IsMovieExist",
			Handler:    _MovieProtoService_IsMovieExist_Handler,
		},
		{
			MethodName: "AddMovie",
			Handler:    _MovieProtoService_AddMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieProtoService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieProtoService_DeleteMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllMovies",
			Handler:       _MovieProtoService_GetAllMovies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}

const (
	FavoriteMovieProtoService_AddClientFavoriteMovie_FullMethodName    = "/services.FavoriteMovieProtoService/AddClientFavoriteMovie"
	FavoriteMovieProtoService_GetClientFavoriteMovies_FullMethodName   = "/services.FavoriteMovieProtoService/GetClientFavoriteMovies"
	FavoriteMovieProtoService_RemoveClientFavoriteMovie_FullMethodName = "/services.FavoriteMovieProtoService/RemoveClientFavoriteMovie"
)

// FavoriteMovieProtoServiceClient is the client API for FavoriteMovieProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteMovieProtoServiceClient interface {
	AddClientFavoriteMovie(ctx context.Context, in *FavoriteMovieRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetClientFavoriteMovies(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MovieResponse], error)
	RemoveClientFavoriteMovie(ctx context.Context, in *FavoriteMovieRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type favoriteMovieProtoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteMovieProtoServiceClient(cc grpc.ClientConnInterface) FavoriteMovieProtoServiceClient {
	return &favoriteMovieProtoServiceClient{cc}
}

func (c *favoriteMovieProtoServiceClient) AddClientFavoriteMovie(ctx context.Context, in *FavoriteMovieRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, FavoriteMovieProtoService_AddClientFavoriteMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteMovieProtoServiceClient) GetClientFavoriteMovies(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MovieResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FavoriteMovieProtoService_ServiceDesc.Streams[0], FavoriteMovieProtoService_GetClientFavoriteMovies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[IdRequest, MovieResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FavoriteMovieProtoService_GetClientFavoriteMoviesClient = grpc.ServerStreamingClient[MovieResponse]

func (c *favoriteMovieProtoServiceClient) RemoveClientFavoriteMovie(ctx context.Context, in *FavoriteMovieRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, FavoriteMovieProtoService_RemoveClientFavoriteMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteMovieProtoServiceServer is the server API for FavoriteMovieProtoService service.
// All implementations must embed UnimplementedFavoriteMovieProtoServiceServer
// for forward compatibility.
type FavoriteMovieProtoServiceServer interface {
	AddClientFavoriteMovie(context.Context, *FavoriteMovieRequest) (*StatusResponse, error)
	GetClientFavoriteMovies(*IdRequest, grpc.ServerStreamingServer[MovieResponse]) error
	RemoveClientFavoriteMovie(context.Context, *FavoriteMovieRequest) (*StatusResponse, error)
	mustEmbedUnimplementedFavoriteMovieProtoServiceServer()
}

// UnimplementedFavoriteMovieProtoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFavoriteMovieProtoServiceServer struct{}

func (UnimplementedFavoriteMovieProtoServiceServer) AddClientFavoriteMovie(context.Context, *FavoriteMovieRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClientFavoriteMovie not implemented")
}
func (UnimplementedFavoriteMovieProtoServiceServer) GetClientFavoriteMovies(*IdRequest, grpc.ServerStreamingServer[MovieResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetClientFavoriteMovies not implemented")
}
func (UnimplementedFavoriteMovieProtoServiceServer) RemoveClientFavoriteMovie(context.Context, *FavoriteMovieRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClientFavoriteMovie not implemented")
}
func (UnimplementedFavoriteMovieProtoServiceServer) mustEmbedUnimplementedFavoriteMovieProtoServiceServer() {
}
func (UnimplementedFavoriteMovieProtoServiceServer) testEmbeddedByValue() {}

// UnsafeFavoriteMovieProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteMovieProtoServiceServer will
// result in compilation errors.
type UnsafeFavoriteMovieProtoServiceServer interface {
	mustEmbedUnimplementedFavoriteMovieProtoServiceServer()
}

func RegisterFavoriteMovieProtoServiceServer(s grpc.ServiceRegistrar, srv FavoriteMovieProtoServiceServer) {
	// If the following call pancis, it indicates UnimplementedFavoriteMovieProtoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FavoriteMovieProtoService_ServiceDesc, srv)
}

func _FavoriteMovieProtoService_AddClientFavoriteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteMovieProtoServiceServer).AddClientFavoriteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteMovieProtoService_AddClientFavoriteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteMovieProtoServiceServer).AddClientFavoriteMovie(ctx, req.(*FavoriteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteMovieProtoService_GetClientFavoriteMovies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FavoriteMovieProtoServiceServer).GetClientFavoriteMovies(m, &grpc.GenericServerStream[IdRequest, MovieResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FavoriteMovieProtoService_GetClientFavoriteMoviesServer = grpc.ServerStreamingServer[MovieResponse]

func _FavoriteMovieProtoService_RemoveClientFavoriteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteMovieProtoServiceServer).RemoveClientFavoriteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteMovieProtoService_RemoveClientFavoriteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteMovieProtoServiceServer).RemoveClientFavoriteMovie(ctx, req.(*FavoriteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteMovieProtoService_ServiceDesc is the grpc.ServiceDesc for FavoriteMovieProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteMovieProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.FavoriteMovieProtoService",
	HandlerType: (*FavoriteMovieProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClientFavoriteMovie",
			Handler:    _FavoriteMovieProtoService_AddClientFavoriteMovie_Handler,
		},
		{
			MethodName: "RemoveClientFavoriteMovie",
			Handler:    _FavoriteMovieProtoService_RemoveClientFavoriteMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClientFavoriteMovies",
			Handler:       _FavoriteMovieProtoService_GetClientFavoriteMovies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}
