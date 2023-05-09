// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Pet Service gRPC client
//
// Command:
// $ goa gen petsvc/design

package client

import (
	"context"
	pet_servicepb "petsvc/gen/grpc/pet_service/pb"
	petservice "petsvc/gen/pet_service"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli pet_servicepb.PetServiceClient
	opts    []grpc.CallOption
}

// SendDogsClientStream implements the petservice.SendDogsClientStream
// interface.
type SendDogsClientStream struct {
	stream pet_servicepb.PetService_SendDogsClient
}

// ReceiveDogsClientStream implements the petservice.ReceiveDogsClientStream
// interface.
type ReceiveDogsClientStream struct {
	stream pet_servicepb.PetService_ReceiveDogsClient
}

// NewClient instantiates gRPC client for all the Pet Service service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: pet_servicepb.NewPetServiceClient(cc),
		opts:    opts,
	}
}

// SendDogs calls the "SendDogs" function in pet_servicepb.PetServiceClient
// interface.
func (c *Client) SendDogs() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSendDogsFunc(c.grpccli, c.opts...),
			nil,
			DecodeSendDogsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// ReceiveDogs calls the "ReceiveDogs" function in
// pet_servicepb.PetServiceClient interface.
func (c *Client) ReceiveDogs() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildReceiveDogsFunc(c.grpccli, c.opts...),
			nil,
			DecodeReceiveDogsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// SendCat calls the "SendCat" function in pet_servicepb.PetServiceClient
// interface.
func (c *Client) SendCat() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSendCatFunc(c.grpccli, c.opts...),
			EncodeSendCatRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// ReceiveCat calls the "ReceiveCat" function in pet_servicepb.PetServiceClient
// interface.
func (c *Client) ReceiveCat() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildReceiveCatFunc(c.grpccli, c.opts...),
			nil,
			DecodeReceiveCatResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Send streams instances of "pet_servicepb.SendDogsStreamingRequest" to the
// "SendDogs" endpoint gRPC stream.
func (s *SendDogsClientStream) Send(res *petservice.DogPayload) error {
	v := NewProtoDogPayloadSendDogsStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *SendDogsClientStream) Close() error {
	// synchronize and report any server error
	_, err := s.stream.CloseAndRecv()
	return err
}

// Recv reads instances of "pet_servicepb.ReceiveDogsResponse" from the
// "ReceiveDogs" endpoint gRPC stream.
func (s *ReceiveDogsClientStream) Recv() (*petservice.DogPayload, error) {
	var res *petservice.DogPayload
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewReceiveDogsResponseDogPayload(v), nil
}