// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_common.proto

package api

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
//const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SdkRequest struct {
	ContractId           *ContractID `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Method               string      `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Payload              string      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	ChannelName          string      `protobuf:"bytes,4,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	AppId                string      `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Sign                 string      `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	File                 []byte      `protobuf:"bytes,7,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SdkRequest) Reset()         { *m = SdkRequest{} }
func (m *SdkRequest) String() string { return proto.CompactTextString(m) }
func (*SdkRequest) ProtoMessage()    {}
func (*SdkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bed5d3d43538692, []int{0}
}

func (m *SdkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdkRequest.Unmarshal(m, b)
}
func (m *SdkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdkRequest.Marshal(b, m, deterministic)
}
func (m *SdkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdkRequest.Merge(m, src)
}
func (m *SdkRequest) XXX_Size() int {
	return xxx_messageInfo_SdkRequest.Size(m)
}
func (m *SdkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SdkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SdkRequest proto.InternalMessageInfo

func (m *SdkRequest) GetContractId() *ContractID {
	if m != nil {
		return m.ContractId
	}
	return nil
}

func (m *SdkRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SdkRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *SdkRequest) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *SdkRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SdkRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *SdkRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type RpcReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	TransactionHash      string   `protobuf:"bytes,4,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	File                 []byte   `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcReply) Reset()         { *m = RpcReply{} }
func (m *RpcReply) String() string { return proto.CompactTextString(m) }
func (*RpcReply) ProtoMessage()    {}
func (*RpcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bed5d3d43538692, []int{1}
}

func (m *RpcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcReply.Unmarshal(m, b)
}
func (m *RpcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcReply.Marshal(b, m, deterministic)
}
func (m *RpcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcReply.Merge(m, src)
}
func (m *RpcReply) XXX_Size() int {
	return xxx_messageInfo_RpcReply.Size(m)
}
func (m *RpcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcReply.DiscardUnknown(m)
}

var xxx_messageInfo_RpcReply proto.InternalMessageInfo

func (m *RpcReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RpcReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RpcReply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *RpcReply) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *RpcReply) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*SdkRequest)(nil), "api.SdkRequest")
	proto.RegisterType((*RpcReply)(nil), "api.RpcReply")
}

func init() { proto.RegisterFile("rpc_common.proto", fileDescriptor_3bed5d3d43538692) }

var fileDescriptor_3bed5d3d43538692 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x65, 0xda, 0xa4, 0xf0, 0x5a, 0xa9, 0x91, 0x25, 0x90, 0xc5, 0x14, 0x3a, 0x85, 0x25,
	0x42, 0x70, 0x04, 0x18, 0xe8, 0xc2, 0x60, 0x0e, 0x10, 0x3d, 0x6c, 0xd3, 0x58, 0xc4, 0x7f, 0x88,
	0xcd, 0xd0, 0x63, 0x70, 0x35, 0x4e, 0x84, 0x62, 0x92, 0xc0, 0xc4, 0xf6, 0xbe, 0xdf, 0xf7, 0x6c,
	0xfd, 0x6c, 0x28, 0x7a, 0x2f, 0x1a, 0xe1, 0x8c, 0x71, 0xb6, 0xf6, 0xbd, 0x8b, 0x8e, 0x2e, 0xd0,
	0xeb, 0xcb, 0x42, 0x38, 0x1b, 0x7b, 0x14, 0x71, 0xff, 0xf0, 0x83, 0x77, 0x5f, 0x04, 0xe0, 0x59,
	0xbe, 0x71, 0xf5, 0xfe, 0xa1, 0x42, 0xa4, 0x37, 0xb0, 0x9e, 0x56, 0x1a, 0x2d, 0x19, 0x29, 0x49,
	0xb5, 0xbe, 0xdd, 0xd6, 0xe8, 0x75, 0x7d, 0x3f, 0x1f, 0xe5, 0x30, 0x5f, 0x23, 0xe9, 0x05, 0xe4,
	0x46, 0xc5, 0xd6, 0x49, 0x76, 0x52, 0x92, 0xea, 0x8c, 0x8f, 0x89, 0x32, 0x58, 0x79, 0x3c, 0x76,
	0x0e, 0x25, 0x5b, 0xa4, 0x62, 0x8a, 0xf4, 0x0a, 0x36, 0xa2, 0x45, 0x6b, 0x55, 0xd7, 0x58, 0x34,
	0x8a, 0x2d, 0x53, 0xbd, 0x1e, 0xd9, 0x13, 0x1a, 0x45, 0xcf, 0x21, 0x47, 0xef, 0x07, 0x83, 0x2c,
	0x95, 0x19, 0x7a, 0xbf, 0x97, 0x94, 0xc2, 0x32, 0xe8, 0x83, 0x65, 0x79, 0x82, 0x69, 0x1e, 0xd8,
	0xab, 0xee, 0x14, 0x5b, 0x95, 0xa4, 0xda, 0xf0, 0x34, 0xef, 0x3e, 0x09, 0x9c, 0x72, 0x2f, 0xb8,
	0xf2, 0xdd, 0x71, 0x58, 0x10, 0x4e, 0xaa, 0xf4, 0x96, 0x8c, 0xa7, 0x79, 0x90, 0x33, 0x2a, 0x04,
	0x3c, 0xa8, 0xd1, 0x7a, 0x8a, 0xff, 0x68, 0x5f, 0x43, 0x11, 0x7b, 0xb4, 0x01, 0x45, 0xd4, 0xce,
	0x36, 0x2d, 0x86, 0x76, 0x54, 0xdf, 0xfe, 0xe1, 0x8f, 0x18, 0xda, 0xd9, 0x29, 0xfb, 0x75, 0x7a,
	0xc9, 0xd3, 0x7f, 0xdf, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x34, 0x5e, 0xd5, 0x62, 0x9a, 0x01,
	0x00, 0x00,
}
