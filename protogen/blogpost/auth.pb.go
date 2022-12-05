// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/auth.proto

package blogpost

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginUserRequest struct {
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserRequest) Reset()         { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()    {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{0}
}

func (m *LoginUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserRequest.Unmarshal(m, b)
}
func (m *LoginUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserRequest.Marshal(b, m, deterministic)
}
func (m *LoginUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserRequest.Merge(m, src)
}
func (m *LoginUserRequest) XXX_Size() int {
	return xxx_messageInfo_LoginUserRequest.Size(m)
}
func (m *LoginUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserRequest proto.InternalMessageInfo

func (m *LoginUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{1}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRequest) Reset()         { *m = TokenRequest{} }
func (m *TokenRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()    {}
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{2}
}

func (m *TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRequest.Unmarshal(m, b)
}
func (m *TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRequest.Marshal(b, m, deterministic)
}
func (m *TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRequest.Merge(m, src)
}
func (m *TokenRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRequest.Size(m)
}
func (m *TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRequest proto.InternalMessageInfo

func (m *TokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type HasAccessResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	HasAccess            bool     `protobuf:"varint,2,opt,name=has_access,json=hasAccess,proto3" json:"has_access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasAccessResponse) Reset()         { *m = HasAccessResponse{} }
func (m *HasAccessResponse) String() string { return proto.CompactTextString(m) }
func (*HasAccessResponse) ProtoMessage()    {}
func (*HasAccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{3}
}

func (m *HasAccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasAccessResponse.Unmarshal(m, b)
}
func (m *HasAccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasAccessResponse.Marshal(b, m, deterministic)
}
func (m *HasAccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasAccessResponse.Merge(m, src)
}
func (m *HasAccessResponse) XXX_Size() int {
	return xxx_messageInfo_HasAccessResponse.Size(m)
}
func (m *HasAccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasAccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasAccessResponse proto.InternalMessageInfo

func (m *HasAccessResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *HasAccessResponse) GetHasAccess() bool {
	if m != nil {
		return m.HasAccess
	}
	return false
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{4}
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

func (m *User) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateUserRequest struct {
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{5}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{6}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{7}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserListRequest struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{8}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

func (m *GetUserListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetUserListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetUserListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type GetUserListResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListResponse) Reset()         { *m = GetUserListResponse{} }
func (m *GetUserListResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserListResponse) ProtoMessage()    {}
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{9}
}

func (m *GetUserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListResponse.Unmarshal(m, b)
}
func (m *GetUserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListResponse.Marshal(b, m, deterministic)
}
func (m *GetUserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListResponse.Merge(m, src)
}
func (m *GetUserListResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserListResponse.Size(m)
}
func (m *GetUserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListResponse proto.InternalMessageInfo

func (m *GetUserListResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserByIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIDRequest) Reset()         { *m = GetUserByIDRequest{} }
func (m *GetUserByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIDRequest) ProtoMessage()    {}
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4668a8c0df8f6333, []int{10}
}

func (m *GetUserByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIDRequest.Unmarshal(m, b)
}
func (m *GetUserByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIDRequest.Merge(m, src)
}
func (m *GetUserByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIDRequest.Size(m)
}
func (m *GetUserByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIDRequest proto.InternalMessageInfo

func (m *GetUserByIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginUserRequest)(nil), "LoginUserRequest")
	proto.RegisterType((*TokenResponse)(nil), "TokenResponse")
	proto.RegisterType((*TokenRequest)(nil), "TokenRequest")
	proto.RegisterType((*HasAccessResponse)(nil), "HasAccessResponse")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "DeleteUserRequest")
	proto.RegisterType((*GetUserListRequest)(nil), "GetUserListRequest")
	proto.RegisterType((*GetUserListResponse)(nil), "GetUserListResponse")
	proto.RegisterType((*GetUserByIDRequest)(nil), "GetUserByIDRequest")
}

func init() { proto.RegisterFile("protos/auth.proto", fileDescriptor_4668a8c0df8f6333) }

var fileDescriptor_4668a8c0df8f6333 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x93, 0x34, 0x8e, 0x9a, 0xe9, 0x87, 0xf0, 0xa4, 0x02, 0xe3, 0x0a, 0xa9, 0x5a, 0x8a,
	0x04, 0x02, 0x6d, 0x51, 0xb8, 0x71, 0x41, 0x09, 0x45, 0x50, 0x54, 0xa4, 0xca, 0xb4, 0x97, 0x5e,
	0x22, 0xd7, 0x9e, 0x26, 0x16, 0xb1, 0xd7, 0x78, 0x37, 0x20, 0x3f, 0x12, 0x6f, 0xc7, 0x23, 0xa0,
	0x5d, 0x3b, 0x76, 0x52, 0x87, 0x1e, 0x10, 0x37, 0xcf, 0xcc, 0x6f, 0x67, 0xc6, 0xff, 0xd9, 0x59,
	0xb0, 0xd3, 0x4c, 0x28, 0x21, 0x4f, 0xfc, 0x85, 0x9a, 0x71, 0xf3, 0xed, 0x0e, 0x4a, 0x57, 0x20,
	0xe2, 0x58, 0x24, 0x85, 0x93, 0x7d, 0x86, 0x07, 0xe7, 0x62, 0x1a, 0x25, 0x57, 0x92, 0x32, 0x8f,
	0xbe, 0x2f, 0x48, 0x2a, 0x74, 0x61, 0x7b, 0x21, 0x29, 0x4b, 0xfc, 0x98, 0x9c, 0xce, 0x51, 0xfb,
	0x79, 0xdf, 0xab, 0x6c, 0x1d, 0x4b, 0x7d, 0x29, 0x7f, 0x8a, 0x2c, 0x74, 0xb6, 0x8a, 0xd8, 0xd2,
	0x66, 0xcf, 0x60, 0xef, 0x52, 0x7c, 0xa3, 0xc4, 0x23, 0x99, 0x8a, 0x44, 0x12, 0x1e, 0x80, 0x65,
	0x1c, 0x4e, 0xdb, 0x90, 0x85, 0xc1, 0x8e, 0x61, 0xb7, 0xc4, 0x8a, 0x72, 0x9b, 0xa9, 0x2f, 0x60,
	0x7f, 0xf2, 0xe5, 0x28, 0x08, 0x48, 0xca, 0x2a, 0xe1, 0x63, 0xe8, 0xea, 0x4e, 0x0c, 0xb9, 0x33,
	0xb4, 0xb8, 0xe9, 0xda, 0xb8, 0xf0, 0x09, 0xc0, 0xcc, 0x97, 0x13, 0xdf, 0x1c, 0x30, 0x6d, 0x6f,
	0x7b, 0xfd, 0xd9, 0x32, 0x03, 0xfb, 0xd5, 0x86, 0xae, 0xa6, 0x71, 0x1f, 0x3a, 0x51, 0x58, 0x96,
	0xea, 0x44, 0xe1, 0xbf, 0xfe, 0x2c, 0x1e, 0x42, 0x5f, 0x73, 0x13, 0x95, 0xa7, 0xe4, 0x74, 0xeb,
	0x83, 0x97, 0x79, 0x4a, 0xba, 0x99, 0x20, 0x23, 0x5f, 0x51, 0x38, 0xf1, 0x95, 0x63, 0x99, 0x68,
	0xbf, 0xf4, 0x8c, 0x94, 0x0e, 0x2f, 0xd2, 0x70, 0x19, 0xee, 0x15, 0xe1, 0xd2, 0x33, 0x52, 0x6c,
	0x06, 0xf6, 0x7b, 0xc3, 0xfe, 0x87, 0xa1, 0xdc, 0xdb, 0x27, 0x7b, 0x07, 0xf6, 0x95, 0x29, 0xbb,
	0x5a, 0x69, 0x83, 0x42, 0x55, 0xf6, 0xce, 0x9d, 0x91, 0x3f, 0x05, 0xfb, 0x94, 0xe6, 0x74, 0x6f,
	0x02, 0x76, 0x0d, 0xf8, 0x91, 0x94, 0x26, 0xce, 0x23, 0xa9, 0x96, 0xd4, 0x43, 0xe8, 0x89, 0xdb,
	0x5b, 0x49, 0xca, 0x90, 0x96, 0x57, 0x5a, 0xfa, 0x3a, 0xcc, 0xa3, 0x38, 0x52, 0xa6, 0x96, 0xe5,
	0x15, 0x86, 0xa6, 0x25, 0xf9, 0x59, 0x30, 0x2b, 0x7f, 0xb0, 0xb4, 0xd8, 0x10, 0x06, 0x6b, 0xb9,
	0xcb, 0x8b, 0x72, 0x08, 0x96, 0xfe, 0x49, 0xe9, 0xb4, 0x8f, 0xb6, 0xea, 0x9b, 0x52, 0xf8, 0xd8,
	0x71, 0xd5, 0xcf, 0x38, 0x3f, 0x3b, 0xfd, 0x4b, 0xd7, 0xc3, 0xdf, 0x1d, 0xd8, 0xd1, 0xcc, 0x57,
	0xca, 0x7e, 0x44, 0x01, 0xe1, 0x23, 0xe8, 0x5e, 0x44, 0xc9, 0x14, 0x7b, 0xfc, 0x43, 0x9c, 0xaa,
	0xdc, 0xb5, 0xf8, 0x85, 0x48, 0xa6, 0xac, 0x85, 0x2f, 0x00, 0xea, 0x71, 0x21, 0xf2, 0xc6, 0xec,
	0xdc, 0xa2, 0x7c, 0x81, 0xd6, 0x7a, 0x23, 0xf2, 0x86, 0xf8, 0x6b, 0x68, 0xad, 0x2c, 0x22, 0x6f,
	0xc8, 0x5c, 0xa3, 0x6f, 0x61, 0x67, 0x45, 0x03, 0x1c, 0xf0, 0xa6, 0xda, 0xee, 0x01, 0xdf, 0x20,
	0x13, 0x6b, 0xe1, 0xcb, 0xea, 0xec, 0x38, 0x3f, 0x0b, 0xeb, 0xb3, 0x2b, 0xca, 0xd4, 0x85, 0x5e,
	0x81, 0x65, 0x1e, 0x0b, 0xb4, 0xf9, 0xdd, 0x47, 0xc3, 0xdd, 0xe7, 0x6b, 0xbb, 0xcf, 0x5a, 0xf8,
	0x1a, 0xfa, 0xd5, 0x06, 0xe3, 0x1e, 0x5f, 0xdd, 0x79, 0x17, 0x79, 0x63, 0xb9, 0x59, 0x6b, 0xbc,
	0x7b, 0x0d, 0xfc, 0xe4, 0x66, 0x2e, 0xa6, 0xa9, 0x90, 0xea, 0xa6, 0x67, 0x5e, 0xa8, 0x37, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x4b, 0x33, 0xd6, 0xcb, 0x04, 0x00, 0x00,
}