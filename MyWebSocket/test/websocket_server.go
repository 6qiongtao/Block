package main

import (
	"fmt"
	"net/http"
	"time"
	"vtoken_digiccy_go/MyWebSocket/impl"

	"github.com/gorilla/websocket"
)


//http
func WsHandler1(w http.ResponseWriter, r *http.Request)  {
	//w.Write([]byte("hello"))
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

//websocket
func WsHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("websocket call wsHandler")
	var (
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
		data []byte
	)

	if wsConn, err = uprGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	//websocket conn
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

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", WsHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}












