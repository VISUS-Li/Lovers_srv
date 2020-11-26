package Channel

import "github.com/gorilla/websocket"

/******
	Channle是消息推送的通道，用于消息推送端向消息处理端发送消息
一般是：(send msg)A->server(handle msg)->B(get msg)
******/
type Channel struct{
	WSConn			*websocket.Conn
	UserId  		string
	LoverId 		string
}

func NewChannel(wsConn *websocket.Conn, userId, loverId string) *Channel{
	c := new(Channel)
	c.WSConn 	 = wsConn
	c.UserId	 = userId
	c.LoverId	 = loverId
	return c
}
