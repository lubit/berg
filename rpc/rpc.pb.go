// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

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

// The request message containing the user's name.
type JobRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addrs                []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobRequest) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

// The response message containing the greetings
type JobReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobReply) Reset()         { *m = JobReply{} }
func (m *JobReply) String() string { return proto.CompactTextString(m) }
func (*JobReply) ProtoMessage()    {}
func (*JobReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *JobReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobReply.Unmarshal(m, b)
}
func (m *JobReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobReply.Marshal(b, m, deterministic)
}
func (m *JobReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobReply.Merge(m, src)
}
func (m *JobReply) XXX_Size() int {
	return xxx_messageInfo_JobReply.Size(m)
}
func (m *JobReply) XXX_DiscardUnknown() {
	xxx_messageInfo_JobReply.DiscardUnknown(m)
}

var xxx_messageInfo_JobReply proto.InternalMessageInfo

func (m *JobReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*JobRequest)(nil), "JobRequest")
	proto.RegisterType((*JobReply)(nil), "JobReply")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe3, 0xe2, 0xf2, 0xca, 0x4f, 0x0a, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x13, 0x53, 0x52, 0x8a, 0x8a, 0x25, 0x98, 0x14,
	0x98, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x15, 0x2e, 0x0e, 0xb0, 0xbe, 0x82, 0x9c, 0x4a, 0x21,
	0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x46, 0x18, 0xd7, 0xe8, 0x01, 0x23,
	0x17, 0xbb, 0x7b, 0x51, 0x6a, 0x6a, 0x49, 0x6a, 0x91, 0x90, 0x0a, 0x17, 0x47, 0x70, 0x49, 0x62,
	0x51, 0x89, 0x57, 0x7e, 0x92, 0x10, 0xb7, 0x1e, 0xc2, 0x52, 0x29, 0x4e, 0x3d, 0x98, 0x49, 0x4a,
	0x0c, 0x42, 0xca, 0x5c, 0xec, 0xc1, 0x25, 0xf9, 0x05, 0x04, 0x15, 0xf9, 0xa6, 0xe6, 0x26, 0xa5,
	0x16, 0x15, 0xe3, 0x51, 0xa4, 0xc0, 0xc5, 0xe2, 0x95, 0x9f, 0x99, 0x87, 0x47, 0x85, 0x22, 0x17,
	0x2b, 0xd8, 0x45, 0xf8, 0x0d, 0x01, 0x39, 0x07, 0xb7, 0x8a, 0x24, 0x36, 0x70, 0x38, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x84, 0xa4, 0x45, 0xb2, 0x54, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	StartJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
	// Sends another greeting
	StopJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
	// List members
	Members(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
	// Join DDN
	Join(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
	// Start
	Start(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
	// Stop
	Stop(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) StartJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/StartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) StopJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/StopJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Members(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/Members", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Join(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Start(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Stop(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	out := new(JobReply)
	err := c.cc.Invoke(ctx, "/Greeter/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	StartJob(context.Context, *JobRequest) (*JobReply, error)
	// Sends another greeting
	StopJob(context.Context, *JobRequest) (*JobReply, error)
	// List members
	Members(context.Context, *JobRequest) (*JobReply, error)
	// Join DDN
	Join(context.Context, *JobRequest) (*JobReply, error)
	// Start
	Start(context.Context, *JobRequest) (*JobReply, error)
	// Stop
	Stop(context.Context, *JobRequest) (*JobReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) StartJob(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (*UnimplementedGreeterServer) StopJob(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (*UnimplementedGreeterServer) Members(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Members not implemented")
}
func (*UnimplementedGreeterServer) Join(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedGreeterServer) Start(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedGreeterServer) Stop(ctx context.Context, req *JobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).StartJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/StopJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).StopJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Members(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Join(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Start(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Stop(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _Greeter_StartJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _Greeter_StopJob_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Greeter_Members_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Greeter_Join_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Greeter_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Greeter_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
