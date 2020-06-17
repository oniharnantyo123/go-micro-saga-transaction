package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	customer "customer/proto/customer"
)

type Customer struct{}

func (e *Customer) Handle(ctx context.Context, msg *customer.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *customer.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
