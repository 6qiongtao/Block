package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/util/log"
	"net/http"
	//"github.com/micro/go-micro/registry"
	//"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"

	"vtoken_digiccy_go/route/handler"
)

func main() {
	// create new web service
	//myRegAddr := "127.0.0.1:8500"
	//reg := consul.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{myRegAddr, }
	//})

	//new web service
	service := web.NewService(
		//web.Registry(reg),
		//服务名称
		web.Name("go.micro.web.route"),
		//版本号
		web.Version("latest"),
		//port
		web.Address(":9999"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	//http
	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("static"))
	// register html handler
	service.Handle("/", rou)
	// register call handler
	rou.POST("/example/call", handler.TestCall)

	//websocket
	//func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	//rou.HandlerFunc("", "/example/ws", handler.SocketTestCall)
	//HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	service.HandleFunc("/ws", handler.SocketTestCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

