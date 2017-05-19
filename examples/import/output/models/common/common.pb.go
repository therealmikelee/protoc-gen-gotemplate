// Code generated by protoc-gen-go.
// source: proto/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	proto/common.proto

It has these top-level messages:
	GetArticle
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetArticle struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Tenant string `protobuf:"bytes,2,opt,name=tenant" json:"tenant,omitempty"`
}

func (m *GetArticle) Reset()                    { *m = GetArticle{} }
func (m *GetArticle) String() string            { return proto.CompactTextString(m) }
func (*GetArticle) ProtoMessage()               {}
func (*GetArticle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetArticle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetArticle) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func init() {
	proto.RegisterType((*GetArticle)(nil), "common.GetArticle")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x3c, 0x25,
	0x13, 0x2e, 0x2e, 0xf7, 0xd4, 0x12, 0xc7, 0xa2, 0x92, 0xcc, 0xe4, 0x9c, 0x54, 0x21, 0x3e, 0x2e,
	0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x31, 0x2e,
	0xb6, 0x92, 0xd4, 0xbc, 0xc4, 0xbc, 0x12, 0x09, 0x26, 0xb0, 0x18, 0x94, 0xe7, 0x24, 0x16, 0x25,
	0x92, 0x9b, 0x9f, 0x92, 0x9a, 0x53, 0x0c, 0x35, 0xd4, 0x1a, 0x42, 0x25, 0xb1, 0x81, 0x0d, 0x37,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x89, 0x5a, 0x1d, 0x73, 0x72, 0x00, 0x00, 0x00,
}