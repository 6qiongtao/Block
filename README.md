#vtoken_digiccy_go 充提币（GO）

## Horizon Go中间层设计文档

**@author：Barry

**版权所有：泰雅通

在Horizon go SDK基础上做一个Go中间层，用于链接前段和Horizon层。本层的功能是，承接前端数据，提交Horizon层处理返回给前端。

### 项目框架

本目录包含一个http server，在http sever端分中做http router、gRPC client，其他微服务为 gRPC server

Python/App

​	↓

go route（RestAPI）

​	 ↓

go-micro client

​	↓

go-micro server

​	

### 包说明

本中间件，就用go语言调用官方提供的Horizon包实现来做前端中转，目的是为了让服务加速。

一个 go-micro web，其它为 go-micro srv（微服务 Server端）。前端通过Rest API调用到本模块的web端，web去通过grpc调用其它微服务去做业务处理。

首先是提供方法暴露的一方--服务器 httpServer。使用第github开源库httprouter做rest解析。（github.com/julienschmidt/httprouter）

grpc调用有两个参与者，分别是：**客户端（client）和服务器（server）**。



##### 一、HttpServer搭建

使用micro命令生成web框架，命令为HorizonRouter：

```go
lqt@u:~/work/go/src$ micro new --type "web" vtoken_digiccy_go/route
Creating service go.micro.web.vtoken_digiccy_go in /home/lqt/work/go/src/vtoken_digiccy_go/route

.
├── main.go
├── plugin.go
├── handler
│   └── handler.go
├── html
│   └── index.html
├── Dockerfile
├── Makefile
└── README.md

//shell启动服务 注册到consul
go run main.go --registry consul
```

上述命令是go micro 官方给出的方法，用于生成web服务，分别是：

* 1、进行改造成，使用开源httprouter做RESTAPI解析（"github.com/julienschmidt/httprouter"）

* 2、handler中写其他的微服务client

  

##### 二、微服务创建

**go micro**为我们提供了srv方法，如下所示：

```go
lqt@u:~/work/go/src$ micro new --type "srv" Block/vtoken_digiccy_go/test
Creating service go.micro.srv.test in /home/lqt/work/go/src/Block/vtoken_digiccy_go/test

.
├── main.go
├── plugin.go
├── handler
│   └── test.go
├── subscriber
│   └── test.go
├── proto/test
│   └── test.proto
├── Dockerfile
├── Makefile
└── README.md

/*
这一段不需要执行
download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file test.proto:
*/

修改protobuf，生成 go
cd /home/lqt/work/go/src/Block/vtoken_digiccy_go/test
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/test/test.proto

//启动服务 注册到consul
go run main.go --registry consul
```

经过服务注册和监听处理，gRPC调用过程中的服务端实现就已经完成了。接下来修改testGrpc的porobuf。



##### 三、关联go web 和 go srv 服务端

讲web和srv结合起来

1.  写
2.  
3.  



##### 四、依赖包管理

1. 安装govendor

   ```shell
   go get -u -v github.com/kardianos/govendor
   
   cd $GOPATH/src/github.com/kardianos/govendor
   go build
   sudo cp govendor /usr/bin
   ```

2. 使用govendor管理包

   ```shell
   cd /home/lqt/work/go/src/Block/hrozionMiddle/route
   govendor init
   #将GOPATH中本工程使用到的依赖包自动移动到vendor目录中
   #说明：如果本地GOPATH没有依赖包，先go get相应的依赖包
   govendor add +external
   或使用缩写： govendor add +e
   更新包
   govendor update
   ```

   

3. 注意 go-micro web项目使用govendor会出现错误

   govendor web项目中执行

   govendor init

   **先清空生成的 json内容，清空后再执行， srv项目不用** 

   govendor add +external

   

4. 参考文档

   https://www.cnblogs.com/liuzhongchao/p/9233177.html

   https://studygolang.com/articles/19248?fr=sidebar

   

   #####     五、Docker部署

   1. 编写docker-compose

      

   2. docker-compose 运行

   

   ##### 六、测试

   #####     











-