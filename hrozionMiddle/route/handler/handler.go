package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/service/grpc"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	//"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	//导入要调用的的proto
	test "Block/hrozionMiddle/test/proto/test"

)

func TestCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode the incoming request as json
	log.Info("handler POST ExampleCall /example/call")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	service := grpc.NewService()
	service.Init()

	// call the backend service
	//调用服务 返回句柄
	testClient := test.NewTestService("go.micro.srv.test", service.Client())
	//name := ps.ByName("name")
	log.Info("name:",request["name"].(string))
	//通过句柄 调用call函数
	rsp, err := testClient.TestCall(context.TODO(), &test.Request{
		Name: request["name"].(string),
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