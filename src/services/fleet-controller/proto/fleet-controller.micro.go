// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: fleet-controller.proto

/*
Package fleet_controller is a generated protocol buffer package.

It is generated from these files:
	fleet-controller.proto

It has these top-level messages:
	Empty
	BookingRequest
	BookingResponse
	UnbookingRequest
	UnbookingResponse
	BeginRideRequest
	BeginRideResponse
	EndRideRequest
	EndRideResponse
	RetrieveReservationsResponse
	RetrieveUnbilledBookingsResponse
	AddInvoiceToBookingRequest
	AddInvoiceToBookingResponse
*/
package fleet_controller

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FleetController service

type FleetControllerService interface {
	Book(ctx context.Context, in *BookingRequest, opts ...client.CallOption) (*BookingResponse, error)
	Unbook(ctx context.Context, in *UnbookingRequest, opts ...client.CallOption) (*UnbookingResponse, error)
	BeginRide(ctx context.Context, in *BeginRideRequest, opts ...client.CallOption) (*BeginRideResponse, error)
	EndRide(ctx context.Context, in *EndRideRequest, opts ...client.CallOption) (*EndRideResponse, error)
	RetrieveReservations(ctx context.Context, in *Empty, opts ...client.CallOption) (*RetrieveReservationsResponse, error)
	RetrieveUnbilledBookings(ctx context.Context, in *Empty, opts ...client.CallOption) (*RetrieveUnbilledBookingsResponse, error)
	AddInvoiceToBooking(ctx context.Context, in *AddInvoiceToBookingRequest, opts ...client.CallOption) (*AddInvoiceToBookingResponse, error)
}

type fleetControllerService struct {
	c    client.Client
	name string
}

func NewFleetControllerService(name string, c client.Client) FleetControllerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "fleetcontroller"
	}
	return &fleetControllerService{
		c:    c,
		name: name,
	}
}

func (c *fleetControllerService) Book(ctx context.Context, in *BookingRequest, opts ...client.CallOption) (*BookingResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.Book", in)
	out := new(BookingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) Unbook(ctx context.Context, in *UnbookingRequest, opts ...client.CallOption) (*UnbookingResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.Unbook", in)
	out := new(UnbookingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) BeginRide(ctx context.Context, in *BeginRideRequest, opts ...client.CallOption) (*BeginRideResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.BeginRide", in)
	out := new(BeginRideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) EndRide(ctx context.Context, in *EndRideRequest, opts ...client.CallOption) (*EndRideResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.EndRide", in)
	out := new(EndRideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) RetrieveReservations(ctx context.Context, in *Empty, opts ...client.CallOption) (*RetrieveReservationsResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.RetrieveReservations", in)
	out := new(RetrieveReservationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) RetrieveUnbilledBookings(ctx context.Context, in *Empty, opts ...client.CallOption) (*RetrieveUnbilledBookingsResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.RetrieveUnbilledBookings", in)
	out := new(RetrieveUnbilledBookingsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetControllerService) AddInvoiceToBooking(ctx context.Context, in *AddInvoiceToBookingRequest, opts ...client.CallOption) (*AddInvoiceToBookingResponse, error) {
	req := c.c.NewRequest(c.name, "FleetController.AddInvoiceToBooking", in)
	out := new(AddInvoiceToBookingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FleetController service

type FleetControllerHandler interface {
	Book(context.Context, *BookingRequest, *BookingResponse) error
	Unbook(context.Context, *UnbookingRequest, *UnbookingResponse) error
	BeginRide(context.Context, *BeginRideRequest, *BeginRideResponse) error
	EndRide(context.Context, *EndRideRequest, *EndRideResponse) error
	RetrieveReservations(context.Context, *Empty, *RetrieveReservationsResponse) error
	RetrieveUnbilledBookings(context.Context, *Empty, *RetrieveUnbilledBookingsResponse) error
	AddInvoiceToBooking(context.Context, *AddInvoiceToBookingRequest, *AddInvoiceToBookingResponse) error
}

func RegisterFleetControllerHandler(s server.Server, hdlr FleetControllerHandler, opts ...server.HandlerOption) {
	type fleetController interface {
		Book(ctx context.Context, in *BookingRequest, out *BookingResponse) error
		Unbook(ctx context.Context, in *UnbookingRequest, out *UnbookingResponse) error
		BeginRide(ctx context.Context, in *BeginRideRequest, out *BeginRideResponse) error
		EndRide(ctx context.Context, in *EndRideRequest, out *EndRideResponse) error
		RetrieveReservations(ctx context.Context, in *Empty, out *RetrieveReservationsResponse) error
		RetrieveUnbilledBookings(ctx context.Context, in *Empty, out *RetrieveUnbilledBookingsResponse) error
		AddInvoiceToBooking(ctx context.Context, in *AddInvoiceToBookingRequest, out *AddInvoiceToBookingResponse) error
	}
	type FleetController struct {
		fleetController
	}
	h := &fleetControllerHandler{hdlr}
	s.Handle(s.NewHandler(&FleetController{h}, opts...))
}

type fleetControllerHandler struct {
	FleetControllerHandler
}

func (h *fleetControllerHandler) Book(ctx context.Context, in *BookingRequest, out *BookingResponse) error {
	return h.FleetControllerHandler.Book(ctx, in, out)
}

func (h *fleetControllerHandler) Unbook(ctx context.Context, in *UnbookingRequest, out *UnbookingResponse) error {
	return h.FleetControllerHandler.Unbook(ctx, in, out)
}

func (h *fleetControllerHandler) BeginRide(ctx context.Context, in *BeginRideRequest, out *BeginRideResponse) error {
	return h.FleetControllerHandler.BeginRide(ctx, in, out)
}

func (h *fleetControllerHandler) EndRide(ctx context.Context, in *EndRideRequest, out *EndRideResponse) error {
	return h.FleetControllerHandler.EndRide(ctx, in, out)
}

func (h *fleetControllerHandler) RetrieveReservations(ctx context.Context, in *Empty, out *RetrieveReservationsResponse) error {
	return h.FleetControllerHandler.RetrieveReservations(ctx, in, out)
}

func (h *fleetControllerHandler) RetrieveUnbilledBookings(ctx context.Context, in *Empty, out *RetrieveUnbilledBookingsResponse) error {
	return h.FleetControllerHandler.RetrieveUnbilledBookings(ctx, in, out)
}

func (h *fleetControllerHandler) AddInvoiceToBooking(ctx context.Context, in *AddInvoiceToBookingRequest, out *AddInvoiceToBookingResponse) error {
	return h.FleetControllerHandler.AddInvoiceToBooking(ctx, in, out)
}
