// Code generated by protoc-gen-go.
// source: server/pkg/metrics/metrics.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	server/pkg/metrics/metrics.proto

It has these top-level messages:
	Metrics
*/
package metrics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Metrics struct {
	ID        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Nodes     int64  `protobuf:"varint,2,opt,name=nodes" json:"nodes,omitempty"`
	Repos     int64  `protobuf:"varint,3,opt,name=repos" json:"repos,omitempty"`
	Commits   int64  `protobuf:"varint,4,opt,name=commits" json:"commits,omitempty"`
	Files     int64  `protobuf:"varint,5,opt,name=files" json:"files,omitempty"`
	Bytes     int64  `protobuf:"varint,6,opt,name=bytes" json:"bytes,omitempty"`
	Jobs      int64  `protobuf:"varint,7,opt,name=jobs" json:"jobs,omitempty"`
	Pipelines int64  `protobuf:"varint,8,opt,name=pipelines" json:"pipelines,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Metrics)(nil), "Metrics")
}

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0xcf, 0x3b, 0x0e, 0xc2, 0x30,
	0x0c, 0x80, 0x61, 0xf5, 0x4d, 0x3d, 0x30, 0x44, 0x0c, 0x19, 0x18, 0x2a, 0x26, 0x26, 0x3a, 0x70,
	0x0e, 0x96, 0x1e, 0xa1, 0x6d, 0x40, 0x86, 0xb6, 0x8e, 0xe2, 0x08, 0x89, 0xc3, 0x71, 0x37, 0x52,
	0x47, 0x15, 0x53, 0xfc, 0x7f, 0xb1, 0x22, 0x05, 0x1a, 0x36, 0xee, 0x6d, 0x5c, 0x6b, 0x5f, 0x8f,
	0x76, 0x36, 0xde, 0xe1, 0xc0, 0xdb, 0x79, 0xb1, 0x8e, 0x3c, 0x9d, 0xbe, 0x09, 0x54, 0xb7, 0x28,
	0x6a, 0x0f, 0x29, 0x8e, 0x3a, 0x69, 0x92, 0x73, 0xdd, 0x85, 0x49, 0x1d, 0xa0, 0x58, 0x68, 0x34,
	0xac, 0xd3, 0x40, 0x59, 0x17, 0x63, 0x55, 0x67, 0x2c, 0xb1, 0xce, 0xa2, 0x4a, 0x28, 0x0d, 0xd5,
	0x40, 0xf3, 0x8c, 0x9e, 0x75, 0x2e, 0xbe, 0xe5, 0xba, 0x7f, 0xc7, 0x29, 0xbc, 0x52, 0xc4, 0x7d,
	0x89, 0x55, 0xfb, 0x8f, 0x0f, 0x5a, 0x46, 0x95, 0x50, 0x0a, 0xf2, 0x27, 0xf5, 0xac, 0x2b, 0x41,
	0x99, 0xd5, 0x11, 0x6a, 0x8b, 0xd6, 0x4c, 0xb8, 0x84, 0xed, 0x9d, 0x5c, 0xfc, 0xa1, 0x2f, 0xe5,
	0x1b, 0xd7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x15, 0x70, 0x26, 0xea, 0x00, 0x00, 0x00,
}
