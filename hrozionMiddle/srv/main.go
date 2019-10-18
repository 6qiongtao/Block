package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"micro/rpc/srv/handler"
	"micro/rpc/srv/subscriber"

	example "micro/rpc/srv/proto/example"
)

func main() {
	// regist to consul
	myRegAddr := "127.0.0.1:8500"
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{myRegAddr, }
	})
	//New Service
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
