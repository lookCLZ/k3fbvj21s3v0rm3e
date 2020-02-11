package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"micro/rpc/srv/handler"
	"micro/rpc/srv/subscriber"

	srv "micro/rpc/srv/proto/srv"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	srv.RegisterSrvHandler(service.Server(), new(handler.Srv))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), new(subscriber.Srv))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
