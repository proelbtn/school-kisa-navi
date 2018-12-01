// Code generated by protoc-gen-go. DO NOT EDIT.
// source: food.proto

package food

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type FoodRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodRequest) Reset()         { *m = FoodRequest{} }
func (m *FoodRequest) String() string { return proto.CompactTextString(m) }
func (*FoodRequest) ProtoMessage()    {}
func (*FoodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c7773ae8ad4a83, []int{0}
}

func (m *FoodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodRequest.Unmarshal(m, b)
}
func (m *FoodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodRequest.Marshal(b, m, deterministic)
}
func (m *FoodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodRequest.Merge(m, src)
}
func (m *FoodRequest) XXX_Size() int {
	return xxx_messageInfo_FoodRequest.Size(m)
}
func (m *FoodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FoodRequest proto.InternalMessageInfo

func (m *FoodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FoodResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodResponse) Reset()         { *m = FoodResponse{} }
func (m *FoodResponse) String() string { return proto.CompactTextString(m) }
func (*FoodResponse) ProtoMessage()    {}
func (*FoodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c7773ae8ad4a83, []int{1}
}

func (m *FoodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodResponse.Unmarshal(m, b)
}
func (m *FoodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodResponse.Marshal(b, m, deterministic)
}
func (m *FoodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodResponse.Merge(m, src)
}
func (m *FoodResponse) XXX_Size() int {
	return xxx_messageInfo_FoodResponse.Size(m)
}
func (m *FoodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FoodResponse proto.InternalMessageInfo

func (m *FoodResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*FoodRequest)(nil), "FoodRequest")
	proto.RegisterType((*FoodResponse)(nil), "FoodResponse")
}

func init() { proto.RegisterFile("food.proto", fileDescriptor_99c7773ae8ad4a83) }

var fileDescriptor_99c7773ae8ad4a83 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcf, 0x4f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2, 0x76, 0xcb, 0xcf, 0x4f, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x94, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0xb1, 0xa9, 0x31, 0xd2, 0xe7, 0x62, 0x01, 0xa9, 0x11, 0x52, 0xe7, 0x62, 0x0b, 0x4e,
	0x4d, 0x2c, 0x4a, 0xce, 0x10, 0xe2, 0xd1, 0x43, 0x32, 0x57, 0x8a, 0x57, 0x0f, 0xd9, 0x08, 0x25,
	0x86, 0x24, 0x36, 0xb0, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x08, 0x46, 0x3a,
	0x8c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FoodClient is the client API for Food service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoodClient interface {
	Search(ctx context.Context, in *FoodRequest, opts ...grpc.CallOption) (*FoodResponse, error)
}

type foodClient struct {
	cc *grpc.ClientConn
}

func NewFoodClient(cc *grpc.ClientConn) FoodClient {
	return &foodClient{cc}
}

func (c *foodClient) Search(ctx context.Context, in *FoodRequest, opts ...grpc.CallOption) (*FoodResponse, error) {
	out := new(FoodResponse)
	err := c.cc.Invoke(ctx, "/Food/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServer is the server API for Food service.
type FoodServer interface {
	Search(context.Context, *FoodRequest) (*FoodResponse, error)
}

func RegisterFoodServer(s *grpc.Server, srv FoodServer) {
	s.RegisterService(&_Food_serviceDesc, srv)
}

func _Food_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Food/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).Search(ctx, req.(*FoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Food_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Food",
	HandlerType: (*FoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Food_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food.proto",
}
