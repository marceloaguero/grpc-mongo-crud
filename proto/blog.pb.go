// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

package blogpb

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogReq) Reset()         { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()    {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{1}
}

func (m *CreateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogReq.Unmarshal(m, b)
}
func (m *CreateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogReq.Marshal(b, m, deterministic)
}
func (m *CreateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogReq.Merge(m, src)
}
func (m *CreateBlogReq) XXX_Size() int {
	return xxx_messageInfo_CreateBlogReq.Size(m)
}
func (m *CreateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogReq proto.InternalMessageInfo

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRes) Reset()         { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()    {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{2}
}

func (m *CreateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRes.Unmarshal(m, b)
}
func (m *CreateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRes.Marshal(b, m, deterministic)
}
func (m *CreateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRes.Merge(m, src)
}
func (m *CreateBlogRes) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRes.Size(m)
}
func (m *CreateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRes proto.InternalMessageInfo

func (m *CreateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlobReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlobReq) Reset()         { *m = ReadBlobReq{} }
func (m *ReadBlobReq) String() string { return proto.CompactTextString(m) }
func (*ReadBlobReq) ProtoMessage()    {}
func (*ReadBlobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{3}
}

func (m *ReadBlobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlobReq.Unmarshal(m, b)
}
func (m *ReadBlobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlobReq.Marshal(b, m, deterministic)
}
func (m *ReadBlobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlobReq.Merge(m, src)
}
func (m *ReadBlobReq) XXX_Size() int {
	return xxx_messageInfo_ReadBlobReq.Size(m)
}
func (m *ReadBlobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlobReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlobReq proto.InternalMessageInfo

func (m *ReadBlobReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRes) Reset()         { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()    {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{4}
}

func (m *ReadBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRes.Unmarshal(m, b)
}
func (m *ReadBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRes.Marshal(b, m, deterministic)
}
func (m *ReadBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRes.Merge(m, src)
}
func (m *ReadBlogRes) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRes.Size(m)
}
func (m *ReadBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRes proto.InternalMessageInfo

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogReq) Reset()         { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()    {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{5}
}

func (m *UpdateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogReq.Unmarshal(m, b)
}
func (m *UpdateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogReq.Marshal(b, m, deterministic)
}
func (m *UpdateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogReq.Merge(m, src)
}
func (m *UpdateBlogReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogReq.Size(m)
}
func (m *UpdateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogReq proto.InternalMessageInfo

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRes) Reset()         { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()    {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{6}
}

func (m *UpdateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRes.Unmarshal(m, b)
}
func (m *UpdateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRes.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRes.Merge(m, src)
}
func (m *UpdateBlogRes) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRes.Size(m)
}
func (m *UpdateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRes proto.InternalMessageInfo

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogReq) Reset()         { *m = DeleteBlogReq{} }
func (m *DeleteBlogReq) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogReq) ProtoMessage()    {}
func (*DeleteBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{7}
}

func (m *DeleteBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogReq.Unmarshal(m, b)
}
func (m *DeleteBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogReq.Marshal(b, m, deterministic)
}
func (m *DeleteBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogReq.Merge(m, src)
}
func (m *DeleteBlogReq) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogReq.Size(m)
}
func (m *DeleteBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogReq proto.InternalMessageInfo

func (m *DeleteBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBlogRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRes) Reset()         { *m = DeleteBlogRes{} }
func (m *DeleteBlogRes) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRes) ProtoMessage()    {}
func (*DeleteBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{8}
}

func (m *DeleteBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRes.Unmarshal(m, b)
}
func (m *DeleteBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRes.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRes.Merge(m, src)
}
func (m *DeleteBlogRes) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRes.Size(m)
}
func (m *DeleteBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRes proto.InternalMessageInfo

func (m *DeleteBlogRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListBlogReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogReq) Reset()         { *m = ListBlogReq{} }
func (m *ListBlogReq) String() string { return proto.CompactTextString(m) }
func (*ListBlogReq) ProtoMessage()    {}
func (*ListBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{9}
}

func (m *ListBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogReq.Unmarshal(m, b)
}
func (m *ListBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogReq.Marshal(b, m, deterministic)
}
func (m *ListBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogReq.Merge(m, src)
}
func (m *ListBlogReq) XXX_Size() int {
	return xxx_messageInfo_ListBlogReq.Size(m)
}
func (m *ListBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogReq proto.InternalMessageInfo

type ListBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRes) Reset()         { *m = ListBlogRes{} }
func (m *ListBlogRes) String() string { return proto.CompactTextString(m) }
func (*ListBlogRes) ProtoMessage()    {}
func (*ListBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{10}
}

func (m *ListBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRes.Unmarshal(m, b)
}
func (m *ListBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRes.Marshal(b, m, deterministic)
}
func (m *ListBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRes.Merge(m, src)
}
func (m *ListBlogRes) XXX_Size() int {
	return xxx_messageInfo_ListBlogRes.Size(m)
}
func (m *ListBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRes proto.InternalMessageInfo

func (m *ListBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*ReadBlobReq)(nil), "blog.ReadBlobReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*DeleteBlogReq)(nil), "blog.DeleteBlogReq")
	proto.RegisterType((*DeleteBlogRes)(nil), "blog.DeleteBlogRes")
	proto.RegisterType((*ListBlogReq)(nil), "blog.ListBlogReq")
	proto.RegisterType((*ListBlogRes)(nil), "blog.ListBlogRes")
}

func init() { proto.RegisterFile("proto/blog.proto", fileDescriptor_fc5203cdc85000bc) }

var fileDescriptor_fc5203cdc85000bc = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0x84, 0x30,
	0x10, 0xcd, 0xe2, 0xba, 0xc2, 0x10, 0x8c, 0x56, 0x0f, 0x04, 0xe3, 0x47, 0x38, 0xe9, 0xc1, 0x65,
	0xb3, 0x9b, 0xec, 0x0f, 0x40, 0x2f, 0x26, 0x9e, 0x30, 0x5e, 0xbc, 0x18, 0xa0, 0xcd, 0xda, 0x84,
	0x6c, 0x91, 0x76, 0xfd, 0x35, 0xfe, 0x58, 0xd3, 0xc1, 0xca, 0x87, 0x24, 0xe8, 0xad, 0xf3, 0x66,
	0xde, 0xbc, 0xd7, 0xd7, 0x14, 0x8e, 0xca, 0x4a, 0x28, 0x11, 0x65, 0x85, 0xd8, 0xcc, 0xf1, 0x48,
	0xa6, 0xfa, 0x1c, 0xe6, 0x30, 0x8d, 0x0b, 0xb1, 0x21, 0x87, 0x60, 0x71, 0xea, 0x4f, 0xae, 0x26,
	0xd7, 0x4e, 0x62, 0x71, 0x4a, 0xce, 0xc0, 0x49, 0x77, 0xea, 0x4d, 0x54, 0xaf, 0x9c, 0xfa, 0x16,
	0xc2, 0x76, 0x0d, 0x3c, 0x50, 0x72, 0x0a, 0xfb, 0x8a, 0xab, 0x82, 0xf9, 0x7b, 0xd8, 0xa8, 0x0b,
	0xe2, 0xc3, 0x41, 0x2e, 0xb6, 0x8a, 0x6d, 0x95, 0x3f, 0x45, 0xdc, 0x94, 0x61, 0x04, 0xde, 0x5d,
	0xc5, 0x52, 0xc5, 0xb4, 0x54, 0xc2, 0xde, 0xc9, 0x05, 0xa0, 0x3a, 0xea, 0xb9, 0x4b, 0x98, 0xa3,
	0x2d, 0x6c, 0xd6, 0xae, 0x7a, 0x04, 0x39, 0x4a, 0x38, 0x07, 0x37, 0x61, 0x29, 0x8d, 0x0b, 0x91,
	0xe9, 0xfd, 0xbd, 0xdb, 0x84, 0xb7, 0x3f, 0xed, 0x3f, 0x6d, 0x8b, 0xc0, 0x7b, 0x2e, 0xe9, 0xff,
	0xfc, 0xb6, 0x09, 0xe3, 0x0a, 0x97, 0xe0, 0xdd, 0xb3, 0x82, 0x35, 0x0a, 0x7d, 0xc7, 0x37, 0xdd,
	0x01, 0xa9, 0xd3, 0x95, 0xbb, 0x3c, 0x67, 0x52, 0xe2, 0x94, 0x9d, 0x98, 0x32, 0xf4, 0xc0, 0x7d,
	0xe4, 0x52, 0x7d, 0x6f, 0xd2, 0x77, 0x6d, 0xca, 0x51, 0x27, 0xcb, 0x4f, 0x0b, 0x5c, 0x5d, 0x3e,
	0xb1, 0xea, 0x83, 0xe7, 0x8c, 0xac, 0x01, 0x9a, 0xe8, 0xc9, 0x49, 0x3d, 0xdf, 0x79, 0xbd, 0x60,
	0x00, 0x94, 0x64, 0x01, 0xb6, 0x89, 0x98, 0x1c, 0xd7, 0x03, 0xad, 0x17, 0x09, 0xba, 0x10, 0x32,
	0xd6, 0x00, 0x4d, 0x68, 0x46, 0xa9, 0x93, 0x7b, 0x30, 0x00, 0x22, 0xaf, 0x89, 0xc6, 0xf0, 0x3a,
	0x69, 0x06, 0x03, 0xa0, 0x24, 0x2b, 0x70, 0x4c, 0x30, 0xd2, 0x58, 0x6c, 0x05, 0x17, 0xfc, 0x82,
	0xe4, 0x62, 0x12, 0xdb, 0x2f, 0x33, 0x8d, 0x96, 0x59, 0x36, 0xc3, 0x6f, 0xb3, 0xfa, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x95, 0xf0, 0xa2, 0x1b, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlobReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error)
	ListBlogs(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlobReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error) {
	out := new(DeleteBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlogs(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogsClient interface {
	Recv() (*ListBlogRes, error)
	grpc.ClientStream
}

type blogServiceListBlogsClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogsClient) Recv() (*ListBlogRes, error) {
	m := new(ListBlogRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlobReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error)
	ListBlogs(*ListBlogReq, BlogService_ListBlogsServer) error
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogReq) (*CreateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlobReq) (*ReadBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogReq) (*UpdateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) DeleteBlog(ctx context.Context, req *DeleteBlogReq) (*DeleteBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlogs(req *ListBlogReq, srv BlogService_ListBlogsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlogs not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlogs(m, &blogServiceListBlogsServer{stream})
}

type BlogService_ListBlogsServer interface {
	Send(*ListBlogRes) error
	grpc.ServerStream
}

type blogServiceListBlogsServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogsServer) Send(m *ListBlogRes) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlogs",
			Handler:       _BlogService_ListBlogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/blog.proto",
}
