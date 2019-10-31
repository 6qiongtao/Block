package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	impl "vtoken_digiccy_go/route/socket"

	"vtoken_digiccy_go/route/cache"
	"vtoken_digiccy_go/route/config"
	"vtoken_digiccy_go/route/merror"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/service/grpc"

	//导入要调用的的proto
	//MYWebSocket "vtoken_digiccy_go/MyWebSocket/proto/MyWebSocket"
	test "vtoken_digiccy_go/test/proto/test"

	"github.com/micro/go-micro/util/log"
)

func TestCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode the incoming request as json
	//log.Info("handler POST TestCall /example/call")
	merror.Log(config.AppLogLevel, "====>docker-compose env AppLogLevel:", config.AppLogLevel)

	//cache.ErrLog.Println("handler POST TestCall /example/call")
	/*
	//demo mysql and redis
	if config.RunMode == "test" {
	//mysql
	isEmpty, err := store.Mysql.IsTableEmpty(new(model.Admin))
	if err != nil {
			log.Debug("IsTableEmpty err:", err)
		}
		log.Info("MysqlTest:Table Is Empty: ", isEmpty)
		//redis
		get, _ := store.RedisString("","test")
		log.Info("Redis Test:RedisGetString: ", get)
		//consul
		pair, _, err := config.ConsulKV.Get("test", nil)
		if err != nil {
			log.Info("client.KV().Get err:", err)
		}
		if pair != nil {
			log.Infof("Consul Test: K: %v, V: %s\n", pair.Key, pair.Value)
		}
	}
	*/

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
		return
	}

	//初始化grpc
	service := grpc.NewService()
	service.Init()

	//调用服务 返回句柄
	testClient := test.NewTestService("go.micro.srv.test", service.Client())
	//log.Info("name:",request["name"].(string))
	//通过句柄 调用call函数
	rsp, err := testClient.TestCall(context.TODO(), &test.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
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
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
		return
	}

	//merror.Log(0, "URI:", r.Host, r.URL, "err:", nil)
	//merror.Warn(0, "URI:", r.Host, r.URL, "err:", nil)
	//merror.Crash(merror.CrashLevel, "URI:", r.Host, r.URL, "err:", nil)

	//log
	cache.ApiLog.Println("URI:", r.Host, r.URL,
						"Method:", r.Method,
						"RequestBody:", request,
						"ResponseBody:", response,)
}



//http 升级websocket
var (
	uprGrader = websocket.Upgrader {
		CheckOrigin: func(r *http.Request) bool {
			//允许跨域
			return true
		},
	}
)
func SocketTestCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	log.Info("handler POST SocketTestCall /example/ws")
	var (
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
		data []byte
	)
	//升级到 websocket
	if wsConn, err = uprGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	//websocket conn init
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	//每隔一秒发送给客户端一个心跳
	go func() {
		var (
			err error
		)
		for  {
			if err = conn.WriteMessage([]byte("heatbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	//循环读取前端消息，回复消息
	for {
		//读取前端消息，
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		//TODO: 调用grpc处理

		//回消息
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

	/*
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
		return
	}

	//初始化grpc
	service := grpc.NewService()
	service.Init()

	//调用服务 返回句柄
	testClient := MYWebSocket.NewMyWebSocketService("go.micro.srv.MyWebSocket", service.Client())
	//log.Info("name:",request["name"].(string))
	//通过句柄 调用call函数
	rsp, err := testClient.MyWebSocketCall(context.TODO(), &MYWebSocket.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
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
		merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
		return
	}

	//merror.Log(0, "URI:", r.Host, r.URL, "err:", nil)
	//merror.Warn(0, "URI:", r.Host, r.URL, "err:", nil)
	//merror.Crash(merror.CrashLevel, "URI:", r.Host, r.URL, "err:", nil)

	//log
	cache.ApiLog.Println("URI:", r.Host, r.URL,
						"Method:", r.Method,
						"RequestBody:", request,
						"ResponseBody:", response,)

	 */
}