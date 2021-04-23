// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pbuser

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type User struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	IsBoy                bool                  `protobuf:"varint,3,opt,name=is_boy,json=isBoy,proto3" json:"is_boy,omitempty"`
	IsGirl               *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=is_girl,json=isGirl,proto3" json:"is_girl,omitempty"`
	Job                  *wrappers.StringValue `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	HomeUrl              string                `protobuf:"bytes,6,opt,name=home_url,json=homeUrl,proto3" json:"home_url,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetIsBoy() bool {
	if m != nil {
		return m.IsBoy
	}
	return false
}

func (m *User) GetIsGirl() *wrappers.BoolValue {
	if m != nil {
		return m.IsGirl
	}
	return nil
}

func (m *User) GetJob() *wrappers.StringValue {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *User) GetHomeUrl() string {
	if m != nil {
		return m.HomeUrl
	}
	return ""
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.user.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x95, 0x36, 0xff, 0xfa, 0xba, 0x59, 0x20, 0x99, 0x08, 0x41, 0xc4, 0x94, 0xc9, 0x95,
	0xe8, 0xc4, 0x48, 0x16, 0xf6, 0x40, 0x19, 0x58, 0x22, 0x5b, 0x35, 0xc1, 0xc8, 0xa9, 0xad, 0x67,
	0x47, 0xa8, 0x2b, 0x9f, 0x1c, 0xc5, 0xa6, 0x0b, 0xa8, 0x9b, 0xef, 0xf9, 0x77, 0xba, 0xd3, 0x01,
	0x4c, 0x4e, 0x22, 0xb3, 0x68, 0xbc, 0x21, 0x85, 0x15, 0x6c, 0x96, 0xd5, 0xed, 0x60, 0xcc, 0xa0,
	0xe5, 0x26, 0x9c, 0xc5, 0xf4, 0xbe, 0xf1, 0x6a, 0x94, 0xce, 0xf3, 0xd1, 0x46, 0xb2, 0xba, 0xf9,
	0x0b, 0x7c, 0x21, 0xb7, 0x56, 0xa2, 0x8b, 0xff, 0x77, 0xdf, 0x0b, 0x48, 0x77, 0x4e, 0x22, 0x21,
	0x90, 0x1e, 0xf8, 0x28, 0x69, 0x52, 0x27, 0xcd, 0xaa, 0x0b, 0x6f, 0x72, 0x01, 0x99, 0x1c, 0xb9,
	0xd2, 0x74, 0x11, 0x8e, 0x51, 0x90, 0x4b, 0xc8, 0x95, 0xeb, 0x85, 0x39, 0xd2, 0x65, 0x9d, 0x34,
	0x65, 0x97, 0x29, 0xd7, 0x9a, 0x23, 0xd9, 0x42, 0xa1, 0x5c, 0x3f, 0x28, 0xd4, 0x34, 0xad, 0x93,
	0x66, 0x7d, 0x5f, 0xb1, 0x98, 0xcd, 0x4e, 0xd9, 0xac, 0x35, 0x46, 0xbf, 0x72, 0x3d, 0xc9, 0x2e,
	0x57, 0xee, 0x49, 0xa1, 0x26, 0x0c, 0x96, 0x9f, 0x46, 0xd0, 0x2c, 0x18, 0xae, 0xff, 0x19, 0x9e,
	0x3d, 0xaa, 0xc3, 0x10, 0x2d, 0x33, 0x48, 0xae, 0xa0, 0xfc, 0x30, 0xa3, 0xec, 0x27, 0xd4, 0x34,
	0x0f, 0xa5, 0x8a, 0x59, 0xef, 0x50, 0x93, 0x07, 0x80, 0xc9, 0xee, 0xb9, 0x97, 0xfb, 0x9e, 0x7b,
	0xba, 0x3e, 0x53, 0xe1, 0xe5, 0xb4, 0x4f, 0xb7, 0xfa, 0xa5, 0x1f, 0x7d, 0x5b, 0xbe, 0xe5, 0x56,
	0xcc, 0x7b, 0x8a, 0x3c, 0x80, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x8d, 0x69, 0x83,
	0x6d, 0x01, 0x00, 0x00,
}