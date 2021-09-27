// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protolca

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

// DataFetchServiceClient is the client API for DataFetchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataFetchServiceClient interface {
	// Get a data set by ID. You use this method if you are sure that a data set
	// of the requested type and ID exists in the database. An error is returned
	// if this is not the case.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataSet, error)
	// Get a data set by ID or name. Returns an empty data set if there is no such
	// data set in the database. If there are multiple data sets with the same
	// name in the database it returns the first of these data sets.
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DataSet, error)
	// Get (a page of) all data sets of a given type. If no page size is given
	// (page <= 0) it defaults to 100. An unset page (page <= 0) defaults to the
	// first page.
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	// Get all descriptors that match the given request.
	GetDescriptors(ctx context.Context, in *GetDescriptorsRequest, opts ...grpc.CallOption) (DataFetchService_GetDescriptorsClient, error)
	// Search for data sets.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DataFetchService_SearchClient, error)
	// Get the descriptors of the data sets that are contained in a given
	// category. The category can be identified by its full path or its ID where
	// an empty string or "/" identifies the root category (or "none-category") of
	// the respective model type.
	GetCategoryContent(ctx context.Context, in *GetCategoryContentRequest, opts ...grpc.CallOption) (DataFetchService_GetCategoryContentClient, error)
	// Get the full category tree for the given model type.
	GetCategoryTree(ctx context.Context, in *GetCategoryTreeRequest, opts ...grpc.CallOption) (*CategoryTree, error)
	// Get possible providers for the given flow. For products these are processes
	// with that product on the output side and for waste flows processes with
	// that waste flow on the input side. For elementary flows, an empty stream is
	// returned.
	GetProvidersFor(ctx context.Context, in *Ref, opts ...grpc.CallOption) (DataFetchService_GetProvidersForClient, error)
}

type dataFetchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataFetchServiceClient(cc grpc.ClientConnInterface) DataFetchServiceClient {
	return &dataFetchServiceClient{cc}
}

func (c *dataFetchServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataSet, error) {
	out := new(DataSet)
	err := c.cc.Invoke(ctx, "/protolca.services.DataFetchService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataFetchServiceClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DataSet, error) {
	out := new(DataSet)
	err := c.cc.Invoke(ctx, "/protolca.services.DataFetchService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataFetchServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/protolca.services.DataFetchService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataFetchServiceClient) GetDescriptors(ctx context.Context, in *GetDescriptorsRequest, opts ...grpc.CallOption) (DataFetchService_GetDescriptorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataFetchService_ServiceDesc.Streams[0], "/protolca.services.DataFetchService/GetDescriptors", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataFetchServiceGetDescriptorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataFetchService_GetDescriptorsClient interface {
	Recv() (*Ref, error)
	grpc.ClientStream
}

type dataFetchServiceGetDescriptorsClient struct {
	grpc.ClientStream
}

func (x *dataFetchServiceGetDescriptorsClient) Recv() (*Ref, error) {
	m := new(Ref)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataFetchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DataFetchService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataFetchService_ServiceDesc.Streams[1], "/protolca.services.DataFetchService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataFetchServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataFetchService_SearchClient interface {
	Recv() (*Ref, error)
	grpc.ClientStream
}

type dataFetchServiceSearchClient struct {
	grpc.ClientStream
}

func (x *dataFetchServiceSearchClient) Recv() (*Ref, error) {
	m := new(Ref)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataFetchServiceClient) GetCategoryContent(ctx context.Context, in *GetCategoryContentRequest, opts ...grpc.CallOption) (DataFetchService_GetCategoryContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataFetchService_ServiceDesc.Streams[2], "/protolca.services.DataFetchService/GetCategoryContent", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataFetchServiceGetCategoryContentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataFetchService_GetCategoryContentClient interface {
	Recv() (*Ref, error)
	grpc.ClientStream
}

type dataFetchServiceGetCategoryContentClient struct {
	grpc.ClientStream
}

func (x *dataFetchServiceGetCategoryContentClient) Recv() (*Ref, error) {
	m := new(Ref)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataFetchServiceClient) GetCategoryTree(ctx context.Context, in *GetCategoryTreeRequest, opts ...grpc.CallOption) (*CategoryTree, error) {
	out := new(CategoryTree)
	err := c.cc.Invoke(ctx, "/protolca.services.DataFetchService/GetCategoryTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataFetchServiceClient) GetProvidersFor(ctx context.Context, in *Ref, opts ...grpc.CallOption) (DataFetchService_GetProvidersForClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataFetchService_ServiceDesc.Streams[3], "/protolca.services.DataFetchService/GetProvidersFor", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataFetchServiceGetProvidersForClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataFetchService_GetProvidersForClient interface {
	Recv() (*Ref, error)
	grpc.ClientStream
}

type dataFetchServiceGetProvidersForClient struct {
	grpc.ClientStream
}

func (x *dataFetchServiceGetProvidersForClient) Recv() (*Ref, error) {
	m := new(Ref)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataFetchServiceServer is the server API for DataFetchService service.
// All implementations must embed UnimplementedDataFetchServiceServer
// for forward compatibility
type DataFetchServiceServer interface {
	// Get a data set by ID. You use this method if you are sure that a data set
	// of the requested type and ID exists in the database. An error is returned
	// if this is not the case.
	Get(context.Context, *GetRequest) (*DataSet, error)
	// Get a data set by ID or name. Returns an empty data set if there is no such
	// data set in the database. If there are multiple data sets with the same
	// name in the database it returns the first of these data sets.
	Find(context.Context, *FindRequest) (*DataSet, error)
	// Get (a page of) all data sets of a given type. If no page size is given
	// (page <= 0) it defaults to 100. An unset page (page <= 0) defaults to the
	// first page.
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	// Get all descriptors that match the given request.
	GetDescriptors(*GetDescriptorsRequest, DataFetchService_GetDescriptorsServer) error
	// Search for data sets.
	Search(*SearchRequest, DataFetchService_SearchServer) error
	// Get the descriptors of the data sets that are contained in a given
	// category. The category can be identified by its full path or its ID where
	// an empty string or "/" identifies the root category (or "none-category") of
	// the respective model type.
	GetCategoryContent(*GetCategoryContentRequest, DataFetchService_GetCategoryContentServer) error
	// Get the full category tree for the given model type.
	GetCategoryTree(context.Context, *GetCategoryTreeRequest) (*CategoryTree, error)
	// Get possible providers for the given flow. For products these are processes
	// with that product on the output side and for waste flows processes with
	// that waste flow on the input side. For elementary flows, an empty stream is
	// returned.
	GetProvidersFor(*Ref, DataFetchService_GetProvidersForServer) error
	mustEmbedUnimplementedDataFetchServiceServer()
}

// UnimplementedDataFetchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataFetchServiceServer struct {
}

func (UnimplementedDataFetchServiceServer) Get(context.Context, *GetRequest) (*DataSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDataFetchServiceServer) Find(context.Context, *FindRequest) (*DataSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedDataFetchServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedDataFetchServiceServer) GetDescriptors(*GetDescriptorsRequest, DataFetchService_GetDescriptorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDescriptors not implemented")
}
func (UnimplementedDataFetchServiceServer) Search(*SearchRequest, DataFetchService_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDataFetchServiceServer) GetCategoryContent(*GetCategoryContentRequest, DataFetchService_GetCategoryContentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCategoryContent not implemented")
}
func (UnimplementedDataFetchServiceServer) GetCategoryTree(context.Context, *GetCategoryTreeRequest) (*CategoryTree, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryTree not implemented")
}
func (UnimplementedDataFetchServiceServer) GetProvidersFor(*Ref, DataFetchService_GetProvidersForServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProvidersFor not implemented")
}
func (UnimplementedDataFetchServiceServer) mustEmbedUnimplementedDataFetchServiceServer() {}

// UnsafeDataFetchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataFetchServiceServer will
// result in compilation errors.
type UnsafeDataFetchServiceServer interface {
	mustEmbedUnimplementedDataFetchServiceServer()
}

func RegisterDataFetchServiceServer(s grpc.ServiceRegistrar, srv DataFetchServiceServer) {
	s.RegisterService(&DataFetchService_ServiceDesc, srv)
}

func _DataFetchService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFetchServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protolca.services.DataFetchService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFetchServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataFetchService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFetchServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protolca.services.DataFetchService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFetchServiceServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataFetchService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFetchServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protolca.services.DataFetchService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFetchServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataFetchService_GetDescriptors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDescriptorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFetchServiceServer).GetDescriptors(m, &dataFetchServiceGetDescriptorsServer{stream})
}

type DataFetchService_GetDescriptorsServer interface {
	Send(*Ref) error
	grpc.ServerStream
}

type dataFetchServiceGetDescriptorsServer struct {
	grpc.ServerStream
}

func (x *dataFetchServiceGetDescriptorsServer) Send(m *Ref) error {
	return x.ServerStream.SendMsg(m)
}

func _DataFetchService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFetchServiceServer).Search(m, &dataFetchServiceSearchServer{stream})
}

type DataFetchService_SearchServer interface {
	Send(*Ref) error
	grpc.ServerStream
}

type dataFetchServiceSearchServer struct {
	grpc.ServerStream
}

func (x *dataFetchServiceSearchServer) Send(m *Ref) error {
	return x.ServerStream.SendMsg(m)
}

func _DataFetchService_GetCategoryContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCategoryContentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFetchServiceServer).GetCategoryContent(m, &dataFetchServiceGetCategoryContentServer{stream})
}

type DataFetchService_GetCategoryContentServer interface {
	Send(*Ref) error
	grpc.ServerStream
}

type dataFetchServiceGetCategoryContentServer struct {
	grpc.ServerStream
}

func (x *dataFetchServiceGetCategoryContentServer) Send(m *Ref) error {
	return x.ServerStream.SendMsg(m)
}

func _DataFetchService_GetCategoryTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataFetchServiceServer).GetCategoryTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protolca.services.DataFetchService/GetCategoryTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataFetchServiceServer).GetCategoryTree(ctx, req.(*GetCategoryTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataFetchService_GetProvidersFor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ref)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFetchServiceServer).GetProvidersFor(m, &dataFetchServiceGetProvidersForServer{stream})
}

type DataFetchService_GetProvidersForServer interface {
	Send(*Ref) error
	grpc.ServerStream
}

type dataFetchServiceGetProvidersForServer struct {
	grpc.ServerStream
}

func (x *dataFetchServiceGetProvidersForServer) Send(m *Ref) error {
	return x.ServerStream.SendMsg(m)
}

// DataFetchService_ServiceDesc is the grpc.ServiceDesc for DataFetchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataFetchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protolca.services.DataFetchService",
	HandlerType: (*DataFetchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DataFetchService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _DataFetchService_Find_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _DataFetchService_GetAll_Handler,
		},
		{
			MethodName: "GetCategoryTree",
			Handler:    _DataFetchService_GetCategoryTree_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDescriptors",
			Handler:       _DataFetchService_GetDescriptors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Search",
			Handler:       _DataFetchService_Search_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCategoryContent",
			Handler:       _DataFetchService_GetCategoryContent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetProvidersFor",
			Handler:       _DataFetchService_GetProvidersFor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data_fetch.proto",
}
