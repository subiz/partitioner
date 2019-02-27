// Code generated by protoc-gen-go. DO NOT EDIT.
// source: partitioner.proto

package header

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Configuration struct {
	Version              string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32                  `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Cluster              string                 `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	TotalPartitions      int32                  `protobuf:"varint,5,opt,name=total_partitions,json=totalPartitions,proto3" json:"total_partitions,omitempty"`
	NextTerm             int32                  `protobuf:"varint,6,opt,name=next_term,json=nextTerm,proto3" json:"next_term,omitempty"`
	Workers              map[string]*WorkerInfo `protobuf:"bytes,7,rep,name=workers,proto3" json:"workers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Configuration) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Configuration) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Configuration) GetTotalPartitions() int32 {
	if m != nil {
		return m.TotalPartitions
	}
	return 0
}

func (m *Configuration) GetNextTerm() int32 {
	if m != nil {
		return m.NextTerm
	}
	return 0
}

func (m *Configuration) GetWorkers() map[string]*WorkerInfo {
	if m != nil {
		return m.Workers
	}
	return nil
}

type WorkerInfo struct {
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Partitions           []int32  `protobuf:"varint,7,rep,packed,name=partitions,proto3" json:"partitions,omitempty"`
	Host                 string   `protobuf:"bytes,8,opt,name=host,proto3" json:"host,omitempty"`
	NotifyHost           string   `protobuf:"bytes,9,opt,name=notify_host,json=notifyHost,proto3" json:"notify_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerInfo) Reset()         { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()    {}
func (*WorkerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{1}
}
func (m *WorkerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerInfo.Unmarshal(m, b)
}
func (m *WorkerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerInfo.Marshal(b, m, deterministic)
}
func (dst *WorkerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerInfo.Merge(dst, src)
}
func (m *WorkerInfo) XXX_Size() int {
	return xxx_messageInfo_WorkerInfo.Size(m)
}
func (m *WorkerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerInfo proto.InternalMessageInfo

func (m *WorkerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerInfo) GetPartitions() []int32 {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *WorkerInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *WorkerInfo) GetNotifyHost() string {
	if m != nil {
		return m.NotifyHost
	}
	return ""
}

type GetConfigRequest struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32    `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{2}
}
func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(dst, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

func (m *GetConfigRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetConfigRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *GetConfigRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type LeaveRequest struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32    `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{3}
}
func (m *LeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRequest.Unmarshal(m, b)
}
func (m *LeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRequest.Marshal(b, m, deterministic)
}
func (dst *LeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRequest.Merge(dst, src)
}
func (m *LeaveRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveRequest.Size(m)
}
func (m *LeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRequest proto.InternalMessageInfo

func (m *LeaveRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LeaveRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LeaveRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *LeaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JoinRequest struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32    `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	NotifyHost           string   `protobuf:"bytes,7,opt,name=notify_host,json=notifyHost,proto3" json:"notify_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{4}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (dst *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(dst, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JoinRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *JoinRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *JoinRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JoinRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *JoinRequest) GetNotifyHost() string {
	if m != nil {
		return m.NotifyHost
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_partitioner_f10fee30da69b5dc, []int{5}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Configuration)(nil), "header.Configuration")
	proto.RegisterMapType((map[string]*WorkerInfo)(nil), "header.Configuration.WorkersEntry")
	proto.RegisterType((*WorkerInfo)(nil), "header.WorkerInfo")
	proto.RegisterType((*GetConfigRequest)(nil), "header.GetConfigRequest")
	proto.RegisterType((*LeaveRequest)(nil), "header.LeaveRequest")
	proto.RegisterType((*JoinRequest)(nil), "header.JoinRequest")
	proto.RegisterType((*Empty)(nil), "header.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Configuration, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Empty, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Empty, error)
}

type coordinatorClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorClient(cc *grpc.ClientConn) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/header.Coordinator/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/header.Coordinator/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/header.Coordinator/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	GetConfig(context.Context, *GetConfigRequest) (*Configuration, error)
	Leave(context.Context, *LeaveRequest) (*Empty, error)
	Join(context.Context, *JoinRequest) (*Empty, error)
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Coordinator/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Coordinator/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Coordinator/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "header.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Coordinator_GetConfig_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Coordinator_Leave_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Coordinator_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	Notify(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Configuration, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Notify(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/header.Worker/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/header.Worker/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	Notify(context.Context, *Configuration) (*Empty, error)
	GetConfig(context.Context, *GetConfigRequest) (*Configuration, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Worker/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Notify(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Worker/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "header.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Worker_Notify_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Worker_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner.proto",
}

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Hello(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*WorkerInfo, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Hello(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*WorkerInfo, error) {
	out := new(WorkerInfo)
	err := c.cc.Invoke(ctx, "/header.Hello/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Hello(context.Context, *GetConfigRequest) (*WorkerInfo, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.Hello/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Hello(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "header.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Hello_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner.proto",
}

func init() { proto.RegisterFile("partitioner.proto", fileDescriptor_partitioner_f10fee30da69b5dc) }

var fileDescriptor_partitioner_f10fee30da69b5dc = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0x6e, 0x7e, 0xdc, 0x5d, 0x6f, 0xd2, 0x6a, 0x1c, 0x15, 0x96, 0x08, 0x5a, 0xf6, 0x29, 0x82,
	0x1c, 0x72, 0x3e, 0x28, 0x22, 0x22, 0x94, 0x62, 0x15, 0x29, 0x72, 0x08, 0x82, 0x2f, 0xe1, 0x34,
	0x1b, 0x7b, 0xf4, 0xba, 0x9b, 0x6e, 0x36, 0xd1, 0xf8, 0xa7, 0xf8, 0x2f, 0x88, 0xff, 0xa3, 0xbb,
	0xb3, 0x6e, 0x3c, 0x9b, 0xf4, 0xc9, 0xf8, 0x36, 0xfb, 0xcd, 0x37, 0x33, 0xdf, 0x7c, 0xbb, 0x77,
	0x70, 0x63, 0x5a, 0x6a, 0x53, 0x99, 0x4a, 0x49, 0xa1, 0xb3, 0xa9, 0x56, 0x46, 0x61, 0x7c, 0x2a,
	0xca, 0xb1, 0xd0, 0xfc, 0x67, 0x1b, 0xf6, 0x0f, 0x95, 0x9c, 0x54, 0x9f, 0xe7, 0xba, 0x74, 0x0c,
	0x64, 0x90, 0x2c, 0x84, 0x9e, 0xd9, 0x90, 0xb5, 0x0f, 0x5a, 0xc3, 0xb4, 0x08, 0x47, 0x44, 0xe8,
	0x1a, 0xa1, 0xcf, 0x59, 0xc7, 0xc2, 0x51, 0x41, 0xb1, 0x63, 0x7f, 0xaa, 0xe7, 0x33, 0x1b, 0xb3,
	0xae, 0x67, 0xff, 0x3e, 0xe2, 0x7d, 0xe8, 0x1b, 0x65, 0xca, 0x7a, 0xb4, 0x1a, 0x3e, 0x63, 0x11,
	0x55, 0x5e, 0x27, 0xfc, 0xed, 0x0a, 0xc6, 0x3b, 0x90, 0x4a, 0xf1, 0xd5, 0x8c, 0xa8, 0x7b, 0x4c,
	0x9c, 0x5d, 0x07, 0xbc, 0x73, 0x13, 0x9e, 0x41, 0xf2, 0x45, 0xe9, 0x33, 0xab, 0x81, 0x25, 0x07,
	0x9d, 0x61, 0x2f, 0xe7, 0x99, 0xd7, 0x9e, 0xfd, 0xa5, 0x3b, 0x7b, 0xef, 0x49, 0x47, 0xd2, 0xe8,
	0x65, 0x11, 0x4a, 0x06, 0x27, 0xb0, 0xd7, 0x4c, 0x60, 0x1f, 0x3a, 0x67, 0x62, 0xc9, 0x5a, 0xa4,
	0xd5, 0x85, 0x38, 0x84, 0x68, 0x51, 0xd6, 0x73, 0x41, 0xdb, 0xf6, 0x72, 0x0c, 0xdd, 0x7d, 0xd9,
	0x2b, 0x39, 0x51, 0x85, 0x27, 0x3c, 0x6d, 0x3f, 0x69, 0xf1, 0x0b, 0x80, 0x3f, 0x09, 0xbc, 0x06,
	0xed, 0x6a, 0x4c, 0x5b, 0xa5, 0x85, 0x8d, 0xf0, 0x2e, 0x40, 0x63, 0x5b, 0x27, 0x37, 0x2a, 0x1a,
	0x88, 0x73, 0xf0, 0x54, 0xcd, 0x0c, 0xdb, 0xa5, 0x0a, 0x8a, 0xf1, 0x1e, 0xf4, 0xa4, 0x32, 0xd5,
	0x64, 0x39, 0xa2, 0x54, 0x4a, 0x29, 0xf0, 0xd0, 0xb1, 0x45, 0xf8, 0x07, 0xe8, 0xbf, 0x14, 0xc6,
	0x2f, 0x5b, 0x88, 0x8b, 0xb9, 0xb0, 0x45, 0x5b, 0xba, 0x24, 0x3e, 0x81, 0xbd, 0x37, 0xa2, 0x5c,
	0x88, 0x2d, 0xf7, 0xbd, 0x6c, 0x0c, 0xff, 0xde, 0x82, 0xde, 0x6b, 0x55, 0xc9, 0xff, 0x3c, 0x67,
	0x65, 0x70, 0x7c, 0xb5, 0xc1, 0xc9, 0x9a, 0xc1, 0x09, 0x44, 0x47, 0xe7, 0x53, 0xb3, 0xcc, 0x7f,
	0x58, 0x95, 0x87, 0x4a, 0xe9, 0x71, 0x25, 0x4b, 0xa3, 0x34, 0x3e, 0x87, 0x74, 0xe5, 0x3c, 0xb2,
	0xf0, 0x30, 0x2e, 0x5f, 0xc6, 0xe0, 0xf6, 0xc6, 0x07, 0xc9, 0x77, 0x30, 0x83, 0x88, 0xdc, 0xc5,
	0x5b, 0x81, 0xd1, 0x34, 0x7b, 0xb0, 0x1f, 0x50, 0x9a, 0x6e, 0xf9, 0x0f, 0xa0, 0xeb, 0x4c, 0xc2,
	0x9b, 0x21, 0xd1, 0xb0, 0x6c, 0x8d, 0x9d, 0x7f, 0x83, 0xd8, 0x3f, 0x45, 0x7c, 0x08, 0xf1, 0x09,
	0xad, 0x83, 0x9b, 0xa5, 0xac, 0x4f, 0xfa, 0xc7, 0xcd, 0xf2, 0x17, 0x10, 0x1d, 0x8b, 0xba, 0x56,
	0xf8, 0x38, 0x04, 0x57, 0x37, 0xd9, 0xf0, 0x45, 0xf1, 0x9d, 0x8f, 0x31, 0xfd, 0x87, 0x1e, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x70, 0xc9, 0x58, 0x9c, 0x04, 0x00, 0x00,
}
