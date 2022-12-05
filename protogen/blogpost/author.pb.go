// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/author.proto

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

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename           string   `protobuf:"bytes,4,opt,name=middlename,proto3" json:"middlename,omitempty"`
	Fullname             string   `protobuf:"bytes,5,opt,name=fullname,proto3" json:"fullname,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Author) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Author) GetMiddlename() string {
	if m != nil {
		return m.Middlename
	}
	return ""
}

func (m *Author) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *Author) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Author) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Author) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type CreateAuthorRequest struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename           string   `protobuf:"bytes,3,opt,name=middlename,proto3" json:"middlename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAuthorRequest) Reset()         { *m = CreateAuthorRequest{} }
func (m *CreateAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAuthorRequest) ProtoMessage()    {}
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{1}
}

func (m *CreateAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAuthorRequest.Unmarshal(m, b)
}
func (m *CreateAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAuthorRequest.Marshal(b, m, deterministic)
}
func (m *CreateAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAuthorRequest.Merge(m, src)
}
func (m *CreateAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAuthorRequest.Size(m)
}
func (m *CreateAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAuthorRequest proto.InternalMessageInfo

func (m *CreateAuthorRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *CreateAuthorRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *CreateAuthorRequest) GetMiddlename() string {
	if m != nil {
		return m.Middlename
	}
	return ""
}

type UpdateAuthorRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename           string   `protobuf:"bytes,4,opt,name=middlename,proto3" json:"middlename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAuthorRequest) Reset()         { *m = UpdateAuthorRequest{} }
func (m *UpdateAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAuthorRequest) ProtoMessage()    {}
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{2}
}

func (m *UpdateAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAuthorRequest.Unmarshal(m, b)
}
func (m *UpdateAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAuthorRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAuthorRequest.Merge(m, src)
}
func (m *UpdateAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAuthorRequest.Size(m)
}
func (m *UpdateAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAuthorRequest proto.InternalMessageInfo

func (m *UpdateAuthorRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateAuthorRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UpdateAuthorRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *UpdateAuthorRequest) GetMiddlename() string {
	if m != nil {
		return m.Middlename
	}
	return ""
}

type DeleteAuthorRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAuthorRequest) Reset()         { *m = DeleteAuthorRequest{} }
func (m *DeleteAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAuthorRequest) ProtoMessage()    {}
func (*DeleteAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{3}
}

func (m *DeleteAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAuthorRequest.Unmarshal(m, b)
}
func (m *DeleteAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAuthorRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAuthorRequest.Merge(m, src)
}
func (m *DeleteAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAuthorRequest.Size(m)
}
func (m *DeleteAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAuthorRequest proto.InternalMessageInfo

func (m *DeleteAuthorRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAuthorListRequest struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorListRequest) Reset()         { *m = GetAuthorListRequest{} }
func (m *GetAuthorListRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthorListRequest) ProtoMessage()    {}
func (*GetAuthorListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{4}
}

func (m *GetAuthorListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorListRequest.Unmarshal(m, b)
}
func (m *GetAuthorListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorListRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthorListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorListRequest.Merge(m, src)
}
func (m *GetAuthorListRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthorListRequest.Size(m)
}
func (m *GetAuthorListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorListRequest proto.InternalMessageInfo

func (m *GetAuthorListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetAuthorListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAuthorListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type AuthorList struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename           string   `protobuf:"bytes,4,opt,name=middlename,proto3" json:"middlename,omitempty"`
	Fullname             string   `protobuf:"bytes,5,opt,name=fullname,proto3" json:"fullname,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorList) Reset()         { *m = AuthorList{} }
func (m *AuthorList) String() string { return proto.CompactTextString(m) }
func (*AuthorList) ProtoMessage()    {}
func (*AuthorList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{5}
}

func (m *AuthorList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorList.Unmarshal(m, b)
}
func (m *AuthorList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorList.Marshal(b, m, deterministic)
}
func (m *AuthorList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorList.Merge(m, src)
}
func (m *AuthorList) XXX_Size() int {
	return xxx_messageInfo_AuthorList.Size(m)
}
func (m *AuthorList) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorList.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorList proto.InternalMessageInfo

func (m *AuthorList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthorList) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *AuthorList) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *AuthorList) GetMiddlename() string {
	if m != nil {
		return m.Middlename
	}
	return ""
}

func (m *AuthorList) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *AuthorList) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *AuthorList) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *AuthorList) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type GetAuthorListResponse struct {
	Authors              []*AuthorList `protobuf:"bytes,1,rep,name=Authors,proto3" json:"Authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAuthorListResponse) Reset()         { *m = GetAuthorListResponse{} }
func (m *GetAuthorListResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthorListResponse) ProtoMessage()    {}
func (*GetAuthorListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{6}
}

func (m *GetAuthorListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorListResponse.Unmarshal(m, b)
}
func (m *GetAuthorListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorListResponse.Marshal(b, m, deterministic)
}
func (m *GetAuthorListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorListResponse.Merge(m, src)
}
func (m *GetAuthorListResponse) XXX_Size() int {
	return xxx_messageInfo_GetAuthorListResponse.Size(m)
}
func (m *GetAuthorListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorListResponse proto.InternalMessageInfo

func (m *GetAuthorListResponse) GetAuthors() []*AuthorList {
	if m != nil {
		return m.Authors
	}
	return nil
}

type GetAuthorByIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorByIDRequest) Reset()         { *m = GetAuthorByIDRequest{} }
func (m *GetAuthorByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthorByIDRequest) ProtoMessage()    {}
func (*GetAuthorByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{7}
}

func (m *GetAuthorByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorByIDRequest.Unmarshal(m, b)
}
func (m *GetAuthorByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthorByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorByIDRequest.Merge(m, src)
}
func (m *GetAuthorByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthorByIDRequest.Size(m)
}
func (m *GetAuthorByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorByIDRequest proto.InternalMessageInfo

func (m *GetAuthorByIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAuthorByIDResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename           string   `protobuf:"bytes,4,opt,name=middlename,proto3" json:"middlename,omitempty"`
	Fullname             string   `protobuf:"bytes,5,opt,name=fullname,proto3" json:"fullname,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorByIDResponse) Reset()         { *m = GetAuthorByIDResponse{} }
func (m *GetAuthorByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthorByIDResponse) ProtoMessage()    {}
func (*GetAuthorByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc3958f4382bb72a, []int{8}
}

func (m *GetAuthorByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorByIDResponse.Unmarshal(m, b)
}
func (m *GetAuthorByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetAuthorByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorByIDResponse.Merge(m, src)
}
func (m *GetAuthorByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetAuthorByIDResponse.Size(m)
}
func (m *GetAuthorByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorByIDResponse proto.InternalMessageInfo

func (m *GetAuthorByIDResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetMiddlename() string {
	if m != nil {
		return m.Middlename
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetAuthorByIDResponse) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*Author)(nil), "Author")
	proto.RegisterType((*CreateAuthorRequest)(nil), "CreateAuthorRequest")
	proto.RegisterType((*UpdateAuthorRequest)(nil), "UpdateAuthorRequest")
	proto.RegisterType((*DeleteAuthorRequest)(nil), "DeleteAuthorRequest")
	proto.RegisterType((*GetAuthorListRequest)(nil), "GetAuthorListRequest")
	proto.RegisterType((*AuthorList)(nil), "AuthorList")
	proto.RegisterType((*GetAuthorListResponse)(nil), "GetAuthorListResponse")
	proto.RegisterType((*GetAuthorByIDRequest)(nil), "GetAuthorByIDRequest")
	proto.RegisterType((*GetAuthorByIDResponse)(nil), "GetAuthorByIDResponse")
}

func init() { proto.RegisterFile("protos/author.proto", fileDescriptor_cc3958f4382bb72a) }

var fileDescriptor_cc3958f4382bb72a = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0xd2, 0x25, 0xdd, 0xde, 0x36, 0x0e, 0x6e, 0x36, 0xaa, 0x08, 0xd0, 0x64, 0x69, 0x68,
	0x27, 0x57, 0x1a, 0x77, 0x44, 0xcb, 0x10, 0x9a, 0xc4, 0x61, 0x0a, 0xe2, 0x82, 0x90, 0x50, 0x16,
	0x3b, 0x9d, 0xa5, 0x24, 0x0e, 0xb1, 0x83, 0xd4, 0x13, 0x3f, 0x8d, 0x1f, 0xc4, 0x81, 0x3b, 0x27,
	0x54, 0xdb, 0xa1, 0x6e, 0x14, 0x7a, 0xe4, 0xc0, 0x8e, 0xef, 0xfb, 0xbe, 0xf7, 0xfc, 0x3e, 0xfb,
	0xbd, 0x04, 0x26, 0x75, 0x23, 0x94, 0x90, 0xb3, 0xb4, 0x55, 0xf7, 0xa2, 0x21, 0x3a, 0x8a, 0x3b,
	0x30, 0x13, 0x65, 0x29, 0x2a, 0x03, 0xe2, 0x1f, 0x1e, 0x84, 0x73, 0xad, 0x42, 0x8f, 0xc0, 0xe7,
	0x74, 0xea, 0x9d, 0x7b, 0x97, 0x87, 0x89, 0xcf, 0x29, 0x7a, 0x02, 0x87, 0x39, 0x6f, 0xa4, 0xaa,
	0xd2, 0x92, 0x4d, 0x7d, 0x0d, 0x6f, 0x00, 0x14, 0xc3, 0x41, 0x91, 0x5a, 0x72, 0xa4, 0xc9, 0x3f,
	0x31, 0x7a, 0x06, 0x50, 0x72, 0x4a, 0x0b, 0xa6, 0xd9, 0x7d, 0xcd, 0x3a, 0xc8, 0x3a, 0x37, 0x6f,
	0x8b, 0x42, 0xb3, 0x81, 0xc9, 0xed, 0x62, 0xf4, 0x14, 0x20, 0x6b, 0x58, 0xaa, 0x18, 0xfd, 0x9c,
	0xaa, 0x69, 0x68, 0x8e, 0xb5, 0xc8, 0x5c, 0xad, 0xe9, 0xb6, 0xa6, 0x1d, 0x3d, 0x36, 0xb4, 0x45,
	0x0c, 0x4d, 0x59, 0xc1, 0x2c, 0x7d, 0x60, 0x68, 0x8b, 0xcc, 0x15, 0x16, 0x30, 0x79, 0xad, 0x4b,
	0x19, 0xcb, 0x09, 0xfb, 0xd2, 0x32, 0xa9, 0xb6, 0x9d, 0x7a, 0xbb, 0x9c, 0xfa, 0x3b, 0x9d, 0x8e,
	0xfa, 0x4e, 0xf1, 0x37, 0x98, 0x7c, 0xd0, 0xcd, 0x6d, 0x1f, 0xf8, 0xcf, 0xae, 0x1a, 0x5f, 0xc0,
	0xe4, 0x5a, 0xdb, 0xdf, 0xd9, 0x00, 0xfe, 0x04, 0xd1, 0x5b, 0xa6, 0x8c, 0xe6, 0x1d, 0x97, 0xaa,
	0xd3, 0x9d, 0x41, 0x28, 0xf2, 0x5c, 0x32, 0xa5, 0xb5, 0x41, 0x62, 0x23, 0x14, 0x41, 0x50, 0xf0,
	0x92, 0x2b, 0xdd, 0x6c, 0x90, 0x98, 0x60, 0xad, 0x96, 0x2c, 0x6d, 0xb2, 0x7b, 0xdb, 0xa6, 0x8d,
	0xf0, 0x4f, 0x0f, 0x60, 0x53, 0xfb, 0x21, 0x0c, 0xda, 0x4b, 0x38, 0xed, 0xdd, 0xa7, 0xac, 0x45,
	0x25, 0x19, 0xba, 0x80, 0xb1, 0x41, 0xe5, 0xd4, 0x3b, 0x1f, 0x5d, 0x1e, 0x5d, 0x1d, 0x11, 0x47,
	0xd5, 0x71, 0xf8, 0xb9, 0xf3, 0x1e, 0x8b, 0xd5, 0xcd, 0xf5, 0xdf, 0xde, 0xed, 0x97, 0xe7, 0x1c,
	0x64, 0x84, 0xf6, 0xa0, 0xff, 0xff, 0x92, 0xaf, 0xbe, 0xfb, 0x70, 0x62, 0x9c, 0xbf, 0x67, 0xcd,
	0x57, 0x9e, 0x31, 0xf4, 0x18, 0xf6, 0x6f, 0x79, 0xb5, 0x44, 0x21, 0x79, 0x53, 0xd6, 0x6a, 0x15,
	0x07, 0xe4, 0x56, 0x54, 0x4b, 0xbc, 0x87, 0x66, 0x70, 0xec, 0x2e, 0x3e, 0x8a, 0xc8, 0xc0, 0x77,
	0x20, 0x1e, 0xdb, 0xb7, 0x30, 0x09, 0xee, 0xe2, 0xa2, 0x88, 0x0c, 0xec, 0x71, 0x2f, 0xc1, 0x5d,
	0x34, 0x14, 0x91, 0x81, 0xbd, 0x73, 0x13, 0x5e, 0xc1, 0xc9, 0xd6, 0x88, 0xa0, 0x53, 0x32, 0xb4,
	0x82, 0xf1, 0x19, 0x19, 0x9c, 0xa4, 0x5e, 0x85, 0xc5, 0xea, 0x86, 0xba, 0x15, 0x9c, 0xa1, 0x71,
	0x2b, 0xb8, 0x23, 0x82, 0xf7, 0x16, 0xc7, 0x1f, 0x81, 0xcc, 0xee, 0x0a, 0xb1, 0xac, 0x85, 0x54,
	0x77, 0xa1, 0xfe, 0x25, 0xbc, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xb0, 0xa6, 0x9f, 0x3e,
	0x06, 0x00, 0x00,
}