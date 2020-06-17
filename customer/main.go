package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"customer/handler"
	"customer/subscriber"
	"github.com/micro/go-micro/v2/service"
	databaseConnect "github.com/oniharnantyo123/go-micro-saga-transaction/database"

	customer "customer/proto/customer"
)

func main() {
	database := databaseConnect.

	db, err := database.

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.customer"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	customer.RegisterCustomerHandler(service.Server(), new(handler.Customer))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.customer", service.Server(), new(subscriber.Customer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
