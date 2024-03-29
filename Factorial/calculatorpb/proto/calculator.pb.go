// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CalculatorRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatorAverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorAverageResponse) Reset()         { *m = CalculatorAverageResponse{} }
func (m *CalculatorAverageResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorAverageResponse) ProtoMessage()    {}
func (*CalculatorAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *CalculatorAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorAverageResponse.Unmarshal(m, b)
}
func (m *CalculatorAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorAverageResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorAverageResponse.Merge(m, src)
}
func (m *CalculatorAverageResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorAverageResponse.Size(m)
}
func (m *CalculatorAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorAverageResponse proto.InternalMessageInfo

func (m *CalculatorAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatorOddEvenResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorOddEvenResponse) Reset()         { *m = CalculatorOddEvenResponse{} }
func (m *CalculatorOddEvenResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorOddEvenResponse) ProtoMessage()    {}
func (*CalculatorOddEvenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *CalculatorOddEvenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorOddEvenResponse.Unmarshal(m, b)
}
func (m *CalculatorOddEvenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorOddEvenResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorOddEvenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorOddEvenResponse.Merge(m, src)
}
func (m *CalculatorOddEvenResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorOddEvenResponse.Size(m)
}
func (m *CalculatorOddEvenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorOddEvenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorOddEvenResponse proto.InternalMessageInfo

func (m *CalculatorOddEvenResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CalculatorOddEvenResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type CalculatorFactorialResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorFactorialResponse) Reset()         { *m = CalculatorFactorialResponse{} }
func (m *CalculatorFactorialResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorFactorialResponse) ProtoMessage()    {}
func (*CalculatorFactorialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *CalculatorFactorialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorFactorialResponse.Unmarshal(m, b)
}
func (m *CalculatorFactorialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorFactorialResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorFactorialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorFactorialResponse.Merge(m, src)
}
func (m *CalculatorFactorialResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorFactorialResponse.Size(m)
}
func (m *CalculatorFactorialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorFactorialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorFactorialResponse proto.InternalMessageInfo

func (m *CalculatorFactorialResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
	proto.RegisterType((*CalculatorAverageResponse)(nil), "calculator.CalculatorAverageResponse")
	proto.RegisterType((*CalculatorOddEvenResponse)(nil), "calculator.CalculatorOddEvenResponse")
	proto.RegisterType((*CalculatorFactorialResponse)(nil), "calculator.CalculatorFactorialResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x69, 0x73, 0x09, 0x3a, 0xc3, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62,
	0x5c, 0x6c, 0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50,
	0x9e, 0x92, 0x0e, 0x97, 0x10, 0xb2, 0xe2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0xea, 0xa2,
	0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x98, 0x6a, 0x08, 0x4f, 0xc9, 0x98, 0x4b, 0x12, 0xa1, 0xda, 0xb1,
	0x2c, 0xb5, 0x28, 0x31, 0x3d, 0x15, 0x87, 0x26, 0x26, 0xb8, 0x26, 0x77, 0x64, 0x4d, 0xfe, 0x29,
	0x29, 0xae, 0x65, 0xa9, 0x79, 0x84, 0x6c, 0x12, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x4c, 0xb9, 0xa4, 0x11, 0x06, 0xb9, 0x25,
	0x26, 0x97, 0xe4, 0x17, 0x65, 0x26, 0xe6, 0x10, 0x32, 0xca, 0x68, 0x0b, 0x33, 0x72, 0x80, 0x04,
	0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x79, 0x72, 0xb1, 0x05, 0x17, 0x96, 0x26, 0x16, 0xa5,
	0x0a, 0xc9, 0xea, 0x21, 0x05, 0x27, 0x46, 0xc8, 0x49, 0xc9, 0xe1, 0x92, 0x86, 0x58, 0xab, 0xc4,
	0x20, 0x14, 0xc8, 0xc5, 0x13, 0x50, 0x94, 0x99, 0x9b, 0x0a, 0x71, 0x52, 0x31, 0xc5, 0x06, 0x1a,
	0x30, 0x0a, 0x05, 0x73, 0xb1, 0x43, 0x83, 0x97, 0x90, 0x69, 0xaa, 0xd8, 0xa5, 0xd1, 0x22, 0x47,
	0x89, 0x41, 0x83, 0x51, 0x28, 0x94, 0x8b, 0x1d, 0x1a, 0xfc, 0x64, 0x1a, 0x8a, 0x16, 0x79, 0x20,
	0x43, 0x0d, 0x40, 0xc6, 0x72, 0xc2, 0x23, 0x83, 0x90, 0xc1, 0xea, 0xd8, 0xa5, 0x31, 0x22, 0x53,
	0x89, 0xc1, 0x89, 0x2f, 0x8a, 0x07, 0xa1, 0xb6, 0x20, 0x29, 0x89, 0x0d, 0x9c, 0xd2, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xca, 0x61, 0xc3, 0xfd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Square(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// Server Streaming
	PrimeFactors(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorsClient, error)
	// Client Streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	// BiDirectional Streaming
	OddEven(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_OddEvenClient, error)
	Factorial(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorFactorialResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Square(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Square", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeFactors(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeFactors", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeFactorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeFactorsClient interface {
	Recv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeFactorsClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeFactorsClient) Recv() (*CalculatorResponse, error) {
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*CalculatorRequest) error
	CloseAndRecv() (*CalculatorAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*CalculatorAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculatorAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) OddEven(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_OddEvenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/OddEven", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceOddEvenClient{stream}
	return x, nil
}

type CalculatorService_OddEvenClient interface {
	Send(*CalculatorRequest) error
	Recv() (*CalculatorOddEvenResponse, error)
	grpc.ClientStream
}

type calculatorServiceOddEvenClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceOddEvenClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceOddEvenClient) Recv() (*CalculatorOddEvenResponse, error) {
	m := new(CalculatorOddEvenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Factorial(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorFactorialResponse, error) {
	out := new(CalculatorFactorialResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Factorial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Square(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// Server Streaming
	PrimeFactors(*CalculatorRequest, CalculatorService_PrimeFactorsServer) error
	// Client Streaming
	Average(CalculatorService_AverageServer) error
	// BiDirectional Streaming
	OddEven(CalculatorService_OddEvenServer) error
	Factorial(context.Context, *CalculatorRequest) (*CalculatorFactorialResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Square(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Square not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeFactors(req *CalculatorRequest, srv CalculatorService_PrimeFactorsServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeFactors not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(srv CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServiceServer) OddEven(srv CalculatorService_OddEvenServer) error {
	return status.Errorf(codes.Unimplemented, "method OddEven not implemented")
}
func (*UnimplementedCalculatorServiceServer) Factorial(ctx context.Context, req *CalculatorRequest) (*CalculatorFactorialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Factorial not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Square_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Square(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Square",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Square(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeFactors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeFactors(m, &calculatorServicePrimeFactorsServer{stream})
}

type CalculatorService_PrimeFactorsServer interface {
	Send(*CalculatorResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeFactorsServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeFactorsServer) Send(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*CalculatorAverageResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *CalculatorAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_OddEven_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).OddEven(&calculatorServiceOddEvenServer{stream})
}

type CalculatorService_OddEvenServer interface {
	Send(*CalculatorOddEvenResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorServiceOddEvenServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceOddEvenServer) Send(m *CalculatorOddEvenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceOddEvenServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Factorial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Factorial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Factorial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Factorial(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Square",
			Handler:    _CalculatorService_Square_Handler,
		},
		{
			MethodName: "Factorial",
			Handler:    _CalculatorService_Factorial_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeFactors",
			Handler:       _CalculatorService_PrimeFactors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OddEven",
			Handler:       _CalculatorService_OddEven_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
