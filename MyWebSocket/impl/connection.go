package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

/*
SendMessage 将消息投递到out channel
ReadMessage 从 in channel 读取消息
起两个协程
	读协程，循环读取websocket，将消息投递到in channel
	写协程，循环读取 out channel，将消息写回 websocket
*/
type Connection struct {
	wsConn *websocket.Conn
	inChan  chan []byte
	outChan chan []byte
	closeChan chan byte
	mutex sync.Mutex
	isClosed bool
}

//init
func InitConnection (wsConn *websocket.Conn) ( conn *Connection, err error) {
	conn = &Connection{
		wsConn:  wsConn,
		inChan:  make(chan []byte, 1024),
		outChan: make(chan []byte, 1024),
		closeChan: make(chan byte, 1),
	}
	//读协程，
	go conn.readLoop()
	//写协程
	go conn.writeLoop()
	return
}
//读消息
func (conn *Connection) ReadMessage() (data []byte, err error)  {
	select {
	case data = <-conn.inChan:
	case <- conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

//写消息
func (conn *Connection) WriteMessage (data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <- conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

//关闭
func (conn *Connection) Close() {
	conn.wsConn.Close()
	//只执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//读消息
func (conn *Connection) readLoop() {
	var (
		data []byte
		err error
	)
	for  {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞，容量1024
		select {
			case conn.inChan <- data:
			case <- conn.closeChan: //close chann
				goto ERR
		}
	}
ERR:
	conn.Close()
}



//写回
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err error
	)
	for  {
		select {
		case data = <- conn.outChan :
		case <-conn.closeChan  :
			goto ERR
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
















