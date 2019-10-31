package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//http
func wsHandler1(w http.ResponseWriter, r *http.Request)  {
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
func wsHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("websocket call wsHandler")
	var (
		conn *websocket.Conn
		err error
		//msgType int
		data []byte
	)

	if conn, err = uprGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	//websocket conn
	for  {
		//text(json), Binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}


}