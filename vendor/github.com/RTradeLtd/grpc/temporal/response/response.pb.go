// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

/*
Package response is a generated protocol buffer package.

It is generated from these files:
	response.proto

It has these top-level messages:
	SignResponse
*/
package response

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

type SignResponse struct {
	H       string `protobuf:"bytes,1,opt,name=h" json:"h,omitempty"`
	R       string `protobuf:"bytes,2,opt,name=r" json:"r,omitempty"`
	S       string `protobuf:"bytes,3,opt,name=s" json:"s,omitempty"`
	V       string `protobuf:"bytes,4,opt,name=v" json:"v,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Hash    string `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	Sig     string `protobuf:"bytes,7,opt,name=sig" json:"sig,omitempty"`
}

func (m *SignResponse) Reset()                    { *m = SignResponse{} }
func (m *SignResponse) String() string            { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()               {}
func (*SignResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignResponse) GetH() string {
	if m != nil {
		return m.H
	}
	return ""
}

func (m *SignResponse) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *SignResponse) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

func (m *SignResponse) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func (m *SignResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SignResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SignResponse) GetSig() string {
	if m != nil {
		return m.Sig
	}
	return ""
}

func init() {
	proto.RegisterType((*SignResponse)(nil), "response.SignResponse")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0xb1, 0x0e, 0x82, 0x40,
	0x0c, 0xc6, 0xf1, 0x54, 0x10, 0xb4, 0x21, 0xc6, 0x74, 0xea, 0x68, 0x9c, 0x9c, 0x5c, 0x7c, 0x13,
	0x7c, 0x02, 0x0c, 0x17, 0x8e, 0x85, 0x23, 0xad, 0xe1, 0x11, 0x7c, 0x6e, 0xd3, 0x1e, 0xb7, 0x7d,
	0xbf, 0xff, 0xf4, 0xe1, 0x45, 0x82, 0xae, 0x69, 0xd1, 0xf0, 0x5c, 0x25, 0x7d, 0x13, 0x9d, 0x8a,
	0xef, 0x3f, 0xc0, 0xee, 0x3d, 0x4f, 0x4b, 0xbf, 0x07, 0xea, 0x10, 0x22, 0xc3, 0x0d, 0x1e, 0xe7,
	0x1e, 0xa2, 0x49, 0xf8, 0x90, 0x25, 0x26, 0xe5, 0x2a, 0x4b, 0x4d, 0x1b, 0xd7, 0x59, 0x1b, 0x31,
	0xb6, 0xc3, 0x38, 0x4a, 0x50, 0xe5, 0xa3, 0xb7, 0x42, 0x22, 0xac, 0xe3, 0xa0, 0x91, 0x1b, 0xcf,
	0xbe, 0xe9, 0x8a, 0x95, 0xce, 0x13, 0xb7, 0x9e, 0x6c, 0x7e, 0x1a, 0x7f, 0xf6, 0xfa, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x17, 0x21, 0x01, 0xab, 0x00, 0x00, 0x00,
}
