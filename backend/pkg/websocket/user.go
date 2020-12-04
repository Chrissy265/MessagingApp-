package websocket

import "github.com/gorilla/websocket"

type User struct {
	ID           int64
	conn         *websocket.Conn
	ConnectionId string
}
