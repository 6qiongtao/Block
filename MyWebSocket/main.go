package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"vtoken_digiccy_go/MyWebSocket/handler"
	"vtoken_digiccy_go/MyWebSocket/subscriber"

	MYWebSocket "vtoken_digiccy_go/MyWebSocket/proto/MyWebSocket"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.MyWebSocket"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	MYWebSocket.RegisterMyWebSocketHandler(service.Server(), new(handler.MyWebSocket))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.MyWebSocket", service.Server(), new(subscriber.MyWebSocket))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.MyWebSocket", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
