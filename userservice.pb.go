// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userservice.proto

package userservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type User_Gender int32

const (
	User_MALE   User_Gender = 0
	User_FEMALE User_Gender = 1
)

var User_Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var User_Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}

func (User_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0, 0}
}

type User struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string      `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Gender               User_Gender `protobuf:"varint,16,opt,name=gender,proto3,enum=userservice.User_Gender" json:"gender,omitempty"`
	Avatar               string      `protobuf:"bytes,17,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() User_Gender {
	if m != nil {
		return m.Gender
	}
	return User_MALE
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type CreateRequest struct {
	User                 *User                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CreateResponse struct {
	User                 *User                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type QueryRequest struct {
	UserId               []string `protobuf:"bytes,1,rep,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{3}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetUserId() []string {
	if m != nil {
		return m.UserId
	}
	return nil
}

type QueryResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{4}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterEnum("userservice.User_Gender", User_Gender_name, User_Gender_value)
	proto.RegisterType((*User)(nil), "userservice.User")
	proto.RegisterType((*CreateRequest)(nil), "userservice.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "userservice.CreateResponse")
	proto.RegisterType((*QueryRequest)(nil), "userservice.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "userservice.QueryResponse")
}

func init() { proto.RegisterFile("userservice.proto", fileDescriptor_68a7ca558839fd2b) }

var fileDescriptor_68a7ca558839fd2b = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0xff, 0x42, 0x69, 0xe0, 0xf6, 0x87, 0xc0, 0xc4, 0x90, 0x71, 0x4c, 0x94, 0x34, 0x51,
	0x59, 0x15, 0x53, 0x37, 0xac, 0x4c, 0x88, 0x41, 0x63, 0xa2, 0x0b, 0xab, 0x3e, 0xc0, 0x60, 0xaf,
	0xa4, 0x09, 0xa5, 0x65, 0xa6, 0xc5, 0xf8, 0x16, 0x3e, 0x99, 0xcf, 0x64, 0x66, 0xa6, 0x45, 0x88,
	0x6c, 0xdd, 0xf5, 0xcc, 0x77, 0x66, 0xee, 0xb9, 0xa7, 0xd0, 0x2b, 0x24, 0x0a, 0x89, 0x62, 0x1d,
	0xbf, 0xa2, 0x9f, 0x89, 0x34, 0x4f, 0x89, 0xbb, 0x75, 0xc4, 0x4e, 0xe6, 0x69, 0x3a, 0x5f, 0xe0,
	0x48, 0xa3, 0x59, 0xf1, 0x36, 0xca, 0xe3, 0x04, 0x65, 0xce, 0x93, 0xcc, 0xb8, 0xbd, 0x2f, 0x0b,
	0xec, 0x17, 0x89, 0x82, 0x74, 0xa0, 0x16, 0x47, 0xd4, 0x1a, 0x58, 0xc3, 0x56, 0x58, 0x8b, 0x23,
	0x72, 0x00, 0x0d, 0x4c, 0x78, 0xbc, 0xa0, 0x35, 0x7d, 0x64, 0x04, 0x61, 0xd0, 0x54, 0xcf, 0x2f,
	0x79, 0x82, 0xb4, 0xae, 0xc1, 0x46, 0x2b, 0x96, 0x71, 0x29, 0xdf, 0x53, 0x11, 0x51, 0xdb, 0xb0,
	0x4a, 0x93, 0x0b, 0x70, 0xe6, 0xb8, 0x8c, 0x50, 0xd0, 0xee, 0xc0, 0x1a, 0x76, 0x02, 0xea, 0x6f,
	0x07, 0x57, 0x01, 0xfc, 0x5b, 0xcd, 0xc3, 0xd2, 0x47, 0xfa, 0xe0, 0xf0, 0x35, 0xcf, 0xb9, 0xa0,
	0x3d, 0xfd, 0x56, 0xa9, 0xbc, 0x63, 0x70, 0x8c, 0x93, 0x34, 0xc1, 0x7e, 0x98, 0xdc, 0x4f, 0xbb,
	0xff, 0x08, 0x80, 0x73, 0x33, 0xd5, 0xdf, 0x96, 0x97, 0x41, 0xfb, 0x5a, 0x20, 0xcf, 0x31, 0xc4,
	0x55, 0x81, 0x32, 0x27, 0xa7, 0x60, 0xab, 0x59, 0x7a, 0x35, 0x37, 0xe8, 0xfd, 0x1a, 0x1c, 0x6a,
	0x4c, 0xc6, 0xd0, 0xda, 0x74, 0xa3, 0x77, 0x76, 0x03, 0xe6, 0x9b, 0xf6, 0xfc, 0xaa, 0x3d, 0xff,
	0xb9, 0x72, 0x84, 0x3f, 0x66, 0x6f, 0x05, 0x9d, 0x6a, 0xa2, 0xcc, 0xd2, 0xa5, 0xc4, 0xbf, 0x1f,
	0x79, 0x06, 0xff, 0x1f, 0x0b, 0x14, 0x1f, 0xd5, 0x8e, 0x7d, 0x70, 0xd4, 0x8b, 0x77, 0xea, 0x07,
	0xd6, 0x55, 0x59, 0x46, 0x79, 0x63, 0x68, 0x97, 0xbe, 0x32, 0xd9, 0x39, 0x34, 0x74, 0x18, 0xed,
	0xdb, 0x1b, 0xcd, 0xf0, 0xe0, 0xd3, 0x02, 0x57, 0xe9, 0x27, 0xc3, 0xc8, 0x04, 0x1c, 0xb3, 0x24,
	0x61, 0x3b, 0x77, 0x76, 0xba, 0x66, 0x47, 0x7b, 0x59, 0x39, 0xfb, 0x0a, 0x1a, 0x3a, 0x0c, 0x39,
	0xdc, 0x71, 0x6d, 0x2f, 0xc2, 0xd8, 0x3e, 0x64, 0xee, 0xcf, 0x1c, 0xdd, 0xc9, 0xe5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x87, 0xa6, 0x09, 0x2c, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _UserService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}
