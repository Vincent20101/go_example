// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pcms.proto

package pcms

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CcrMsg struct {
	// pcms send DNN, appId to aaa-client, used to select the diameter peer and route
	DiameterId *DiameterId `protobuf:"bytes,1,opt,name=diameterId,proto3" json:"diameterId,omitempty"`
	// this is msgPack encoded CCR go structure
	Avps                 []byte   `protobuf:"bytes,2,opt,name=avps,proto3" json:"avps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CcrMsg) Reset()         { *m = CcrMsg{} }
func (m *CcrMsg) String() string { return proto.CompactTextString(m) }
func (*CcrMsg) ProtoMessage()    {}
func (*CcrMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{0}
}

func (m *CcrMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CcrMsg.Unmarshal(m, b)
}
func (m *CcrMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CcrMsg.Marshal(b, m, deterministic)
}
func (m *CcrMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CcrMsg.Merge(m, src)
}
func (m *CcrMsg) XXX_Size() int {
	return xxx_messageInfo_CcrMsg.Size(m)
}
func (m *CcrMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CcrMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CcrMsg proto.InternalMessageInfo

func (m *CcrMsg) GetDiameterId() *DiameterId {
	if m != nil {
		return m.DiameterId
	}
	return nil
}

func (m *CcrMsg) GetAvps() []byte {
	if m != nil {
		return m.Avps
	}
	return nil
}

type CcaMsg struct {
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	//status -1, means GRPC failed, i.e it didnot reach the other service, or GRPC timeout
	//*for sending request message, gx-servce should handle this case
	//*for receving request from client, diamSvc should handle this and coded proper error response message back
	// to client side
	// this is msgPack encoded CCA go structure
	Avps                 []byte   `protobuf:"bytes,2,opt,name=avps,proto3" json:"avps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CcaMsg) Reset()         { *m = CcaMsg{} }
func (m *CcaMsg) String() string { return proto.CompactTextString(m) }
func (*CcaMsg) ProtoMessage()    {}
func (*CcaMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{1}
}

func (m *CcaMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CcaMsg.Unmarshal(m, b)
}
func (m *CcaMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CcaMsg.Marshal(b, m, deterministic)
}
func (m *CcaMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CcaMsg.Merge(m, src)
}
func (m *CcaMsg) XXX_Size() int {
	return xxx_messageInfo_CcaMsg.Size(m)
}
func (m *CcaMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CcaMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CcaMsg proto.InternalMessageInfo

func (m *CcaMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CcaMsg) GetAvps() []byte {
	if m != nil {
		return m.Avps
	}
	return nil
}

type RarMsg struct {
	// this is msgPack encoded RAR go structure
	Avps                 []byte   `protobuf:"bytes,1,opt,name=avps,proto3" json:"avps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RarMsg) Reset()         { *m = RarMsg{} }
func (m *RarMsg) String() string { return proto.CompactTextString(m) }
func (*RarMsg) ProtoMessage()    {}
func (*RarMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{2}
}

func (m *RarMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RarMsg.Unmarshal(m, b)
}
func (m *RarMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RarMsg.Marshal(b, m, deterministic)
}
func (m *RarMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RarMsg.Merge(m, src)
}
func (m *RarMsg) XXX_Size() int {
	return xxx_messageInfo_RarMsg.Size(m)
}
func (m *RarMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RarMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RarMsg proto.InternalMessageInfo

func (m *RarMsg) GetAvps() []byte {
	if m != nil {
		return m.Avps
	}
	return nil
}

type RaaMsg struct {
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// this is msgPack encoded RAA go structure
	Avps                 []byte   `protobuf:"bytes,2,opt,name=avps,proto3" json:"avps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaaMsg) Reset()         { *m = RaaMsg{} }
func (m *RaaMsg) String() string { return proto.CompactTextString(m) }
func (*RaaMsg) ProtoMessage()    {}
func (*RaaMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{3}
}

func (m *RaaMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaaMsg.Unmarshal(m, b)
}
func (m *RaaMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaaMsg.Marshal(b, m, deterministic)
}
func (m *RaaMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaaMsg.Merge(m, src)
}
func (m *RaaMsg) XXX_Size() int {
	return xxx_messageInfo_RaaMsg.Size(m)
}
func (m *RaaMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RaaMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RaaMsg proto.InternalMessageInfo

func (m *RaaMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RaaMsg) GetAvps() []byte {
	if m != nil {
		return m.Avps
	}
	return nil
}

type HelloReq struct {
	Hello                string   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{4}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

type HelloRsp struct {
	Hello                string   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRsp) Reset()         { *m = HelloRsp{} }
func (m *HelloRsp) String() string { return proto.CompactTextString(m) }
func (*HelloRsp) ProtoMessage()    {}
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{5}
}

func (m *HelloRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRsp.Unmarshal(m, b)
}
func (m *HelloRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRsp.Marshal(b, m, deterministic)
}
func (m *HelloRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRsp.Merge(m, src)
}
func (m *HelloRsp) XXX_Size() int {
	return xxx_messageInfo_HelloRsp.Size(m)
}
func (m *HelloRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRsp proto.InternalMessageInfo

func (m *HelloRsp) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

type DiameterId struct {
	Dnn                  string   `protobuf:"bytes,1,opt,name=dnn,proto3" json:"dnn,omitempty"`
	AppId                int32    `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiameterId) Reset()         { *m = DiameterId{} }
func (m *DiameterId) String() string { return proto.CompactTextString(m) }
func (*DiameterId) ProtoMessage()    {}
func (*DiameterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed238852961c25a8, []int{6}
}

func (m *DiameterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterId.Unmarshal(m, b)
}
func (m *DiameterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterId.Marshal(b, m, deterministic)
}
func (m *DiameterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterId.Merge(m, src)
}
func (m *DiameterId) XXX_Size() int {
	return xxx_messageInfo_DiameterId.Size(m)
}
func (m *DiameterId) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterId.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterId proto.InternalMessageInfo

func (m *DiameterId) GetDnn() string {
	if m != nil {
		return m.Dnn
	}
	return ""
}

func (m *DiameterId) GetAppId() int32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func init() {
	proto.RegisterType((*CcrMsg)(nil), "pcms.CcrMsg")
	proto.RegisterType((*CcaMsg)(nil), "pcms.CcaMsg")
	proto.RegisterType((*RarMsg)(nil), "pcms.RarMsg")
	proto.RegisterType((*RaaMsg)(nil), "pcms.RaaMsg")
	proto.RegisterType((*HelloReq)(nil), "pcms.HelloReq")
	proto.RegisterType((*HelloRsp)(nil), "pcms.HelloRsp")
	proto.RegisterType((*DiameterId)(nil), "pcms.DiameterId")
}

func init() { proto.RegisterFile("pcms.proto", fileDescriptor_ed238852961c25a8) }

var fileDescriptor_ed238852961c25a8 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0xb5, 0x09, 0x3a, 0x96, 0x52, 0x06, 0x91, 0x12, 0x3c, 0x84, 0x9c, 0xaa, 0x87, 0x22,
	0xb5, 0x3f, 0x20, 0xf5, 0x60, 0x0f, 0x8a, 0x8c, 0x5f, 0x30, 0x6e, 0x16, 0x2d, 0x34, 0xcd, 0xba,
	0x1b, 0xfb, 0x23, 0xfe, 0xb0, 0xec, 0x6e, 0xba, 0x31, 0xa0, 0x87, 0xde, 0xde, 0xdb, 0x79, 0xef,
	0xcd, 0xec, 0x0c, 0x80, 0x96, 0x95, 0x9d, 0x69, 0x53, 0x37, 0x35, 0x0e, 0x1c, 0x2e, 0x9e, 0x21,
	0x5d, 0x4a, 0xf3, 0x64, 0xdf, 0xf1, 0x16, 0xa0, 0x5c, 0x73, 0xa5, 0x1a, 0x65, 0x56, 0xe5, 0x44,
	0xe4, 0x62, 0x7a, 0x3e, 0x1f, 0xcf, 0xbc, 0xe1, 0x21, 0xbe, 0xd3, 0x2f, 0x0d, 0x22, 0x0c, 0x78,
	0xa7, 0xed, 0xe4, 0x38, 0x17, 0xd3, 0x21, 0x79, 0x5c, 0x2c, 0x5c, 0x1e, 0xbb, 0xbc, 0x4b, 0x48,
	0x6d, 0xc3, 0xcd, 0x97, 0xf5, 0x59, 0x09, 0xb5, 0xec, 0x4f, 0xd7, 0x15, 0xa4, 0xc4, 0x7e, 0x8a,
	0x7d, 0x55, 0xf4, 0x33, 0x89, 0x0f, 0xce, 0xcc, 0xe1, 0xf4, 0x51, 0x6d, 0x36, 0x35, 0xa9, 0x4f,
	0xbc, 0x80, 0xe4, 0xc3, 0x61, 0x6f, 0x3b, 0xa3, 0x40, 0x3a, 0x85, 0xd5, 0xff, 0x28, 0x16, 0x00,
	0xdd, 0xdf, 0x71, 0x0c, 0x27, 0xe5, 0x76, 0xdb, 0x2a, 0x1c, 0x74, 0x2e, 0xd6, 0x7a, 0x55, 0xfa,
	0xc6, 0x09, 0x05, 0x32, 0xff, 0x16, 0x30, 0x7a, 0x91, 0x95, 0xbd, 0x67, 0x7e, 0x55, 0x66, 0xb7,
	0x96, 0x0a, 0xaf, 0x21, 0xf1, 0xad, 0x70, 0x14, 0x36, 0xba, 0x9f, 0x2c, 0xeb, 0x71, 0xab, 0x8b,
	0x23, 0xbc, 0x01, 0x68, 0xcd, 0x4b, 0x69, 0x70, 0x18, 0xea, 0xe1, 0x46, 0x59, 0x64, 0x6e, 0x1b,
	0x3d, 0x2d, 0x71, 0xd4, 0x86, 0x4d, 0x66, 0x91, 0x05, 0xed, 0x5b, 0xea, 0xcf, 0x7e, 0xf7, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x5d, 0x8d, 0xde, 0xdc, 0x04, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PcmsAaaServiceClient is the client API for PcmsAaaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PcmsAaaServiceClient interface {
	// PCMS service sending Hello message to aaa-client service
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	PcmsAaaCcr(ctx context.Context, in *CcrMsg, opts ...grpc.CallOption) (*CcaMsg, error)
	PcmsAaaRar(ctx context.Context, in *RarMsg, opts ...grpc.CallOption) (*RaaMsg, error)
}

type pcmsAaaServiceClient struct {
	cc *grpc.ClientConn
}

func NewPcmsAaaServiceClient(cc *grpc.ClientConn) PcmsAaaServiceClient {
	return &pcmsAaaServiceClient{cc}
}

func (c *pcmsAaaServiceClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/pcms.PcmsAaaService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcmsAaaServiceClient) PcmsAaaCcr(ctx context.Context, in *CcrMsg, opts ...grpc.CallOption) (*CcaMsg, error) {
	out := new(CcaMsg)
	err := c.cc.Invoke(ctx, "/pcms.PcmsAaaService/PcmsAaaCcr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcmsAaaServiceClient) PcmsAaaRar(ctx context.Context, in *RarMsg, opts ...grpc.CallOption) (*RaaMsg, error) {
	out := new(RaaMsg)
	err := c.cc.Invoke(ctx, "/pcms.PcmsAaaService/PcmsAaaRar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PcmsAaaServiceServer is the server API for PcmsAaaService service.
type PcmsAaaServiceServer interface {
	// PCMS service sending Hello message to aaa-client service
	Hello(context.Context, *HelloReq) (*HelloRsp, error)
	PcmsAaaCcr(context.Context, *CcrMsg) (*CcaMsg, error)
	PcmsAaaRar(context.Context, *RarMsg) (*RaaMsg, error)
}

func RegisterPcmsAaaServiceServer(s *grpc.Server, srv PcmsAaaServiceServer) {
	s.RegisterService(&_PcmsAaaService_serviceDesc, srv)
}

func _PcmsAaaService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcmsAaaServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcms.PcmsAaaService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcmsAaaServiceServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcmsAaaService_PcmsAaaCcr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CcrMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcmsAaaServiceServer).PcmsAaaCcr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcms.PcmsAaaService/PcmsAaaCcr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcmsAaaServiceServer).PcmsAaaCcr(ctx, req.(*CcrMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcmsAaaService_PcmsAaaRar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RarMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcmsAaaServiceServer).PcmsAaaRar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcms.PcmsAaaService/PcmsAaaRar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcmsAaaServiceServer).PcmsAaaRar(ctx, req.(*RarMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _PcmsAaaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pcms.PcmsAaaService",
	HandlerType: (*PcmsAaaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _PcmsAaaService_Hello_Handler,
		},
		{
			MethodName: "PcmsAaaCcr",
			Handler:    _PcmsAaaService_PcmsAaaCcr_Handler,
		},
		{
			MethodName: "PcmsAaaRar",
			Handler:    _PcmsAaaService_PcmsAaaRar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pcms.proto",
}
