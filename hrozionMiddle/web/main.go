package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"

	"net/http"

	//"github.com/micro/go-web"
	//"github.com/micro/go-micro/web"
	"hrozionMiddle/web/handler"
)

func main() {
	// create new web service
    myRegAddr := "127.0.0.1:8500"
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{myRegAddr, }
	})

	//创建1个新的web服务
	service := web.NewService(
		web.Registry(reg),
		//服务名称
		web.Name("go.micro.web.web"),
		//版本号
		web.Version("latest"),
		//设置服务的端口号
		web.Address(":7777"),
	)

	// initialise service
	//初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))
	// register html handler
	service.Handle("/", rou)

	//注册当前的web前段页面
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//注册call的请求
	//service.HandleFunc("/example/call", handler.ExampleCall)
	rou.POST("/example/call", handler.ExampleCall)
	// run service
	//运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
