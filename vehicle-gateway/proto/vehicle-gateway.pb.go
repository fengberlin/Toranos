// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle-gateway.proto

package vehicle_gateway

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

type ReachVehicleRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=VehicleId,proto3" json:"VehicleId,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReachVehicleRequest) Reset()         { *m = ReachVehicleRequest{} }
func (m *ReachVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*ReachVehicleRequest) ProtoMessage()    {}
func (*ReachVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_gateway_996d90324d20f0c4, []int{0}
}
func (m *ReachVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReachVehicleRequest.Unmarshal(m, b)
}
func (m *ReachVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReachVehicleRequest.Marshal(b, m, deterministic)
}
func (dst *ReachVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReachVehicleRequest.Merge(dst, src)
}
func (m *ReachVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_ReachVehicleRequest.Size(m)
}
func (m *ReachVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReachVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReachVehicleRequest proto.InternalMessageInfo

func (m *ReachVehicleRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

func (m *ReachVehicleRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type ReachVehicleResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReachVehicleResponse) Reset()         { *m = ReachVehicleResponse{} }
func (m *ReachVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*ReachVehicleResponse) ProtoMessage()    {}
func (*ReachVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_gateway_996d90324d20f0c4, []int{1}
}
func (m *ReachVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReachVehicleResponse.Unmarshal(m, b)
}
func (m *ReachVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReachVehicleResponse.Marshal(b, m, deterministic)
}
func (dst *ReachVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReachVehicleResponse.Merge(dst, src)
}
func (m *ReachVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_ReachVehicleResponse.Size(m)
}
func (m *ReachVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReachVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReachVehicleResponse proto.InternalMessageInfo

func (m *ReachVehicleResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func init() {
	proto.RegisterType((*ReachVehicleRequest)(nil), "ReachVehicleRequest")
	proto.RegisterType((*ReachVehicleResponse)(nil), "ReachVehicleResponse")
}

func init() {
	proto.RegisterFile("vehicle-gateway.proto", fileDescriptor_vehicle_gateway_996d90324d20f0c4)
}

var fileDescriptor_vehicle_gateway_996d90324d20f0c4 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4b, 0xcd, 0xc8,
	0x4c, 0xce, 0x49, 0xd5, 0x4d, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0xf2, 0xe6, 0x12, 0x0e, 0x4a, 0x4d, 0x4c, 0xce, 0x08, 0x83, 0xc8, 0x06, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x70, 0x71, 0x42, 0x45, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x10, 0x02, 0x42, 0x62, 0x5c, 0x6c, 0xbe, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x12,
	0x4c, 0x60, 0x29, 0x28, 0x4f, 0xc9, 0x8c, 0x4b, 0x04, 0xd5, 0xb0, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x39, 0x2e, 0xae, 0xe0, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0xe2, 0xb4, 0xd2, 0x1c, 0xb0,
	0x71, 0x1c, 0x41, 0x48, 0x22, 0x46, 0xfe, 0x5c, 0x7c, 0x50, 0x2d, 0xee, 0x10, 0xc7, 0x09, 0xd9,
	0x72, 0xf1, 0x20, 0x9b, 0x24, 0x24, 0xa2, 0x87, 0xc5, 0x95, 0x52, 0xa2, 0x7a, 0xd8, 0xac, 0x53,
	0x62, 0x48, 0x62, 0x03, 0x7b, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x32, 0xd5, 0xae,
	0xf5, 0x00, 0x00, 0x00,
}
