package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/client"
	//倒入的是案例的proto
	//example "github.com/micro/examples/template/srv/proto/example"
	example "micro/rpc/srv/proto/example"
)

func ExampleCall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// decode the incoming request as json
	log.Info("handler POST ExampleCall /example/call")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	//调用服务 返回句柄
	exampleClient := example.NewExampleService("go.micro.srv.srv", client.DefaultClient)
	name := ps.ByName("name")
	//通过句柄 调用call函数
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: name,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
