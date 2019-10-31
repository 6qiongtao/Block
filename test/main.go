package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"vtoken_digiccy_go/test/handler"
	"vtoken_digiccy_go/test/subscriber"

	test "vtoken_digiccy_go/test/proto/test"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.test"),
		micro.Version("latest"),
		//micro.Address(":8888"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test", service.Server(), new(subscriber.Test))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
