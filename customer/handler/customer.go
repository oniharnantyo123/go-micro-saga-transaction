package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	customer "customer/proto/customer"
)

type Customer struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Customer) Detail(ctx context.Context, req *customer.Request, rsp *customer.Response) error {
	log.Info("Received Customer.Call request")

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Customer) Stream(ctx context.Context, req *customer.StreamingRequest, stream customer.Customer_StreamStream) error {
	log.Infof("Received Customer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&customer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Customer) PingPong(ctx context.Context, stream customer.Customer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&customer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
