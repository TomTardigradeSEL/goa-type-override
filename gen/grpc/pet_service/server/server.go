// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Pet Service gRPC server
//
// Command:
// $ goa gen petsvc/design

package server

import (
	"context"
	pet_servicepb "petsvc/gen/grpc/pet_service/pb"
	petservice "petsvc/gen/pet_service"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the pet_servicepb.PetServiceServer interface.
type Server struct {
	SendDogsH    goagrpc.StreamHandler
	ReceiveDogsH goagrpc.StreamHandler
	SendCatH     goagrpc.UnaryHandler
	ReceiveCatH  goagrpc.UnaryHandler
	pet_servicepb.UnimplementedPetServiceServer
}

// SendDogsServerStream implements the petservice.SendDogsServerStream
// interface.
type SendDogsServerStream struct {
	stream pet_servicepb.PetService_SendDogsServer
}

// ReceiveDogsServerStream implements the petservice.ReceiveDogsServerStream
// interface.
type ReceiveDogsServerStream struct {
	stream pet_servicepb.PetService_ReceiveDogsServer
}

// New instantiates the server struct with the Pet Service service endpoints.
func New(e *petservice.Endpoints, uh goagrpc.UnaryHandler, sh goagrpc.StreamHandler) *Server {
	return &Server{
		SendDogsH:    NewSendDogsHandler(e.SendDogs, sh),
		ReceiveDogsH: NewReceiveDogsHandler(e.ReceiveDogs, sh),
		SendCatH:     NewSendCatHandler(e.SendCat, uh),
		ReceiveCatH:  NewReceiveCatHandler(e.ReceiveCat, uh),
	}
}

// NewSendDogsHandler creates a gRPC handler which serves the "Pet Service"
// service "SendDogs" endpoint.
func NewSendDogsHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, nil)
	}
	return h
}

// SendDogs implements the "SendDogs" method in pet_servicepb.PetServiceServer
// interface.
func (s *Server) SendDogs(stream pet_servicepb.PetService_SendDogsServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "SendDogs")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Pet Service")
	_, err := s.SendDogsH.Decode(ctx, nil)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &petservice.SendDogsEndpointInput{
		Stream: &SendDogsServerStream{stream: stream},
	}
	err = s.SendDogsH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewReceiveDogsHandler creates a gRPC handler which serves the "Pet Service"
// service "ReceiveDogs" endpoint.
func NewReceiveDogsHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, nil)
	}
	return h
}

// ReceiveDogs implements the "ReceiveDogs" method in
// pet_servicepb.PetServiceServer interface.
func (s *Server) ReceiveDogs(message *pet_servicepb.ReceiveDogsRequest, stream pet_servicepb.PetService_ReceiveDogsServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "ReceiveDogs")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Pet Service")
	_, err := s.ReceiveDogsH.Decode(ctx, message)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &petservice.ReceiveDogsEndpointInput{
		Stream: &ReceiveDogsServerStream{stream: stream},
	}
	err = s.ReceiveDogsH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewSendCatHandler creates a gRPC handler which serves the "Pet Service"
// service "SendCat" endpoint.
func NewSendCatHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSendCatRequest, EncodeSendCatResponse)
	}
	return h
}

// SendCat implements the "SendCat" method in pet_servicepb.PetServiceServer
// interface.
func (s *Server) SendCat(ctx context.Context, message *pet_servicepb.SendCatRequest) (*pet_servicepb.SendCatResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "SendCat")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Pet Service")
	resp, err := s.SendCatH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pet_servicepb.SendCatResponse), nil
}

// NewReceiveCatHandler creates a gRPC handler which serves the "Pet Service"
// service "ReceiveCat" endpoint.
func NewReceiveCatHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeReceiveCatResponse)
	}
	return h
}

// ReceiveCat implements the "ReceiveCat" method in
// pet_servicepb.PetServiceServer interface.
func (s *Server) ReceiveCat(ctx context.Context, message *pet_servicepb.ReceiveCatRequest) (*pet_servicepb.ReceiveCatResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "ReceiveCat")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Pet Service")
	resp, err := s.ReceiveCatH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pet_servicepb.ReceiveCatResponse), nil
}

// Recv reads instances of "pet_servicepb.SendDogsStreamingRequest" from the
// "SendDogs" endpoint gRPC stream.
func (s *SendDogsServerStream) Recv() (*petservice.DogPayload, error) {
	var res *petservice.DogPayload
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewSendDogsStreamingRequestDogPayload(v), nil
}

func (s *SendDogsServerStream) Close() error {
	// synchronize stream
	return s.stream.SendAndClose(&pet_servicepb.SendDogsResponse{})
}

// Send streams instances of "pet_servicepb.ReceiveDogsResponse" to the
// "ReceiveDogs" endpoint gRPC stream.
func (s *ReceiveDogsServerStream) Send(res *petservice.DogPayload) error {
	v := NewProtoDogPayloadReceiveDogsResponse(res)
	return s.stream.Send(v)
}

func (s *ReceiveDogsServerStream) Close() error {
	// nothing to do here
	return nil
}