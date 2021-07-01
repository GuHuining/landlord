package hub

import "github.com/gorilla/websocket"

type Player struct {
	UserID   int
	Nickname string
	Conn     *websocket.Conn
}
