package room

import (
	"github.com/gorilla/websocket"
)

type Musician struct {
	room *Room
	conn *websocket.Conn
	send chan []byte
}
