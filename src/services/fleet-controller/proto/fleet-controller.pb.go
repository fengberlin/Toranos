// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleet-controller.proto

package fleet_controller

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type BookingRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=VehicleId,proto3" json:"VehicleId,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingRequest) Reset()         { *m = BookingRequest{} }
func (m *BookingRequest) String() string { return proto.CompactTextString(m) }
func (*BookingRequest) ProtoMessage()    {}
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{1}
}
func (m *BookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingRequest.Unmarshal(m, b)
}
func (m *BookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingRequest.Marshal(b, m, deterministic)
}
func (dst *BookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingRequest.Merge(dst, src)
}
func (m *BookingRequest) XXX_Size() int {
	return xxx_messageInfo_BookingRequest.Size(m)
}
func (m *BookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookingRequest proto.InternalMessageInfo

func (m *BookingRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

func (m *BookingRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type BookingResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	ReservedTime         uint32   `protobuf:"varint,2,opt,name=ReservedTime,proto3" json:"ReservedTime,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingResponse) Reset()         { *m = BookingResponse{} }
func (m *BookingResponse) String() string { return proto.CompactTextString(m) }
func (*BookingResponse) ProtoMessage()    {}
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{2}
}
func (m *BookingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingResponse.Unmarshal(m, b)
}
func (m *BookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingResponse.Marshal(b, m, deterministic)
}
func (dst *BookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingResponse.Merge(dst, src)
}
func (m *BookingResponse) XXX_Size() int {
	return xxx_messageInfo_BookingResponse.Size(m)
}
func (m *BookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookingResponse proto.InternalMessageInfo

func (m *BookingResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *BookingResponse) GetReservedTime() uint32 {
	if m != nil {
		return m.ReservedTime
	}
	return 0
}

func (m *BookingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UnbookingRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=VehicleId,proto3" json:"VehicleId,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbookingRequest) Reset()         { *m = UnbookingRequest{} }
func (m *UnbookingRequest) String() string { return proto.CompactTextString(m) }
func (*UnbookingRequest) ProtoMessage()    {}
func (*UnbookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{3}
}
func (m *UnbookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbookingRequest.Unmarshal(m, b)
}
func (m *UnbookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbookingRequest.Marshal(b, m, deterministic)
}
func (dst *UnbookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbookingRequest.Merge(dst, src)
}
func (m *UnbookingRequest) XXX_Size() int {
	return xxx_messageInfo_UnbookingRequest.Size(m)
}
func (m *UnbookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnbookingRequest proto.InternalMessageInfo

func (m *UnbookingRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

func (m *UnbookingRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type UnbookingResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbookingResponse) Reset()         { *m = UnbookingResponse{} }
func (m *UnbookingResponse) String() string { return proto.CompactTextString(m) }
func (*UnbookingResponse) ProtoMessage()    {}
func (*UnbookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{4}
}
func (m *UnbookingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbookingResponse.Unmarshal(m, b)
}
func (m *UnbookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbookingResponse.Marshal(b, m, deterministic)
}
func (dst *UnbookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbookingResponse.Merge(dst, src)
}
func (m *UnbookingResponse) XXX_Size() int {
	return xxx_messageInfo_UnbookingResponse.Size(m)
}
func (m *UnbookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnbookingResponse proto.InternalMessageInfo

func (m *UnbookingResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *UnbookingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RetrieveReservationsResponse struct {
	Reservations         []*RetrieveReservationsResponse_Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
	Error                string                                      `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *RetrieveReservationsResponse) Reset()         { *m = RetrieveReservationsResponse{} }
func (m *RetrieveReservationsResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveReservationsResponse) ProtoMessage()    {}
func (*RetrieveReservationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{5}
}
func (m *RetrieveReservationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveReservationsResponse.Unmarshal(m, b)
}
func (m *RetrieveReservationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveReservationsResponse.Marshal(b, m, deterministic)
}
func (dst *RetrieveReservationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveReservationsResponse.Merge(dst, src)
}
func (m *RetrieveReservationsResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveReservationsResponse.Size(m)
}
func (m *RetrieveReservationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveReservationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveReservationsResponse proto.InternalMessageInfo

func (m *RetrieveReservationsResponse) GetReservations() []*RetrieveReservationsResponse_Reservation {
	if m != nil {
		return m.Reservations
	}
	return nil
}

func (m *RetrieveReservationsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RetrieveReservationsResponse_Reservation struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Vehicle              string   `protobuf:"bytes,3,opt,name=Vehicle,proto3" json:"Vehicle,omitempty"`
	Customer             string   `protobuf:"bytes,4,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Status               uint32   `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveReservationsResponse_Reservation) Reset() {
	*m = RetrieveReservationsResponse_Reservation{}
}
func (m *RetrieveReservationsResponse_Reservation) String() string { return proto.CompactTextString(m) }
func (*RetrieveReservationsResponse_Reservation) ProtoMessage()    {}
func (*RetrieveReservationsResponse_Reservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_a676a5a05ce10420, []int{5, 0}
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Unmarshal(m, b)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Marshal(b, m, deterministic)
}
func (dst *RetrieveReservationsResponse_Reservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveReservationsResponse_Reservation.Merge(dst, src)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Size() int {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Size(m)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveReservationsResponse_Reservation.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveReservationsResponse_Reservation proto.InternalMessageInfo

func (m *RetrieveReservationsResponse_Reservation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RetrieveReservationsResponse_Reservation) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *RetrieveReservationsResponse_Reservation) GetVehicle() string {
	if m != nil {
		return m.Vehicle
	}
	return ""
}

func (m *RetrieveReservationsResponse_Reservation) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *RetrieveReservationsResponse_Reservation) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*BookingRequest)(nil), "BookingRequest")
	proto.RegisterType((*BookingResponse)(nil), "BookingResponse")
	proto.RegisterType((*UnbookingRequest)(nil), "UnbookingRequest")
	proto.RegisterType((*UnbookingResponse)(nil), "UnbookingResponse")
	proto.RegisterType((*RetrieveReservationsResponse)(nil), "RetrieveReservationsResponse")
	proto.RegisterType((*RetrieveReservationsResponse_Reservation)(nil), "RetrieveReservationsResponse.Reservation")
}

func init() {
	proto.RegisterFile("fleet-controller.proto", fileDescriptor_fleet_controller_a676a5a05ce10420)
}

var fileDescriptor_fleet_controller_a676a5a05ce10420 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x0e, 0xd2, 0x30,
	0x18, 0x65, 0x03, 0x06, 0x7c, 0xfc, 0x37, 0x84, 0x2c, 0x0b, 0x1a, 0xd2, 0x2b, 0x8c, 0xb1, 0x89,
	0xf8, 0x00, 0x46, 0x09, 0x26, 0xbb, 0xd0, 0x98, 0xa2, 0xde, 0xc3, 0xf6, 0xa1, 0x0b, 0x63, 0xc5,
	0xb6, 0x23, 0xf1, 0x11, 0xbc, 0xf2, 0x5d, 0x7c, 0x42, 0xb3, 0x32, 0xc6, 0x40, 0x42, 0xbc, 0xf0,
	0xf2, 0x3b, 0xcd, 0x39, 0x3d, 0x3d, 0xdf, 0x29, 0x8c, 0xb7, 0x31, 0xa2, 0x7e, 0x11, 0x88, 0x44,
	0x4b, 0x11, 0xc7, 0x28, 0xd9, 0x41, 0x0a, 0x2d, 0x68, 0x03, 0xea, 0xcb, 0xfd, 0x41, 0xff, 0xa0,
	0x1f, 0xa0, 0xf7, 0x56, 0x88, 0x5d, 0x94, 0x7c, 0xe5, 0xf8, 0x3d, 0x45, 0xa5, 0xc9, 0x04, 0x5a,
	0x5f, 0xf0, 0x5b, 0x14, 0xc4, 0xe8, 0x87, 0xae, 0x35, 0xb5, 0x66, 0x2d, 0x7e, 0x01, 0xc8, 0x53,
	0x80, 0x45, 0xaa, 0xb4, 0xd8, 0xa3, 0xf4, 0x43, 0xd7, 0x36, 0xc7, 0x25, 0x84, 0xee, 0xa0, 0x5f,
	0xe8, 0xa9, 0x83, 0x48, 0x14, 0x66, 0x94, 0x55, 0x1a, 0x04, 0xa8, 0xd4, 0x36, 0x8d, 0x8d, 0x62,
	0x93, 0x97, 0x10, 0x42, 0xa1, 0xc3, 0x51, 0xa1, 0x3c, 0x62, 0xf8, 0x29, 0xda, 0xa3, 0x11, 0xed,
	0xf2, 0x2b, 0x8c, 0x8c, 0xa0, 0xbe, 0x94, 0x52, 0x48, 0xb7, 0x6a, 0x6e, 0x3c, 0x0d, 0xf4, 0x23,
	0x0c, 0x3e, 0x27, 0x9b, 0xff, 0x69, 0xdf, 0x87, 0x61, 0x49, 0xf1, 0x1f, 0x1f, 0x50, 0x98, 0xb3,
	0xcb, 0xe6, 0x7e, 0xd9, 0x30, 0xe1, 0xa8, 0x65, 0x84, 0x47, 0x3c, 0xbd, 0x65, 0xad, 0x23, 0x91,
	0xa8, 0x42, 0xf6, 0x3d, 0x74, 0x64, 0x09, 0x77, 0xad, 0x69, 0x75, 0xd6, 0x9e, 0x3f, 0x63, 0x8f,
	0x48, 0xac, 0x04, 0xf2, 0x2b, 0xfa, 0x7d, 0x17, 0xde, 0x4f, 0x0b, 0xda, 0x25, 0x0e, 0xe9, 0x81,
	0x9d, 0xe7, 0xd2, 0xe5, 0xb6, 0x1f, 0x66, 0x71, 0x2d, 0x24, 0xae, 0x35, 0x86, 0x6f, 0xb4, 0x61,
	0x56, 0xf9, 0x05, 0x20, 0x2e, 0x34, 0xf2, 0xec, 0xf2, 0xe0, 0xcf, 0x23, 0xf1, 0xa0, 0x79, 0x8e,
	0xcd, 0xad, 0x99, 0xa3, 0x62, 0x26, 0x63, 0x70, 0x56, 0x7a, 0xad, 0x53, 0xe5, 0xd6, 0xcd, 0x3d,
	0xf9, 0x34, 0xff, 0x6d, 0x41, 0xff, 0x5d, 0xd6, 0xc7, 0x45, 0x51, 0x47, 0xf2, 0x1c, 0x6a, 0x59,
	0x5f, 0x48, 0x9f, 0x5d, 0xd7, 0xd0, 0x1b, 0xb0, 0x9b, 0x1e, 0xd1, 0x0a, 0x79, 0x09, 0xce, 0x69,
	0x3b, 0x64, 0xc8, 0x6e, 0x17, 0xef, 0x11, 0xf6, 0xd7, 0xe6, 0x68, 0x85, 0xbc, 0x86, 0xd1, 0xbd,
	0x3c, 0x89, 0xc3, 0x4c, 0xff, 0xbd, 0x27, 0x0f, 0xe3, 0xa6, 0x95, 0x8d, 0x63, 0x3e, 0xcc, 0xab,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x7b, 0xde, 0x05, 0x4a, 0x03, 0x00, 0x00,
}
