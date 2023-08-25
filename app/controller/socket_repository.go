package controller

import (
	_ "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type SocketRepository struct {
	Clients map[*websocket.Conn]bool
}

func (sr *SocketRepository) UpgradeConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	return connection
}

func (sr *SocketRepository) AddConnection(conn *websocket.Conn) {
	sr.Clients[conn] = true
}

func (sr *SocketRepository) DeleteConnection(conn *websocket.Conn) {
	delete(sr.Clients, conn)
}

func (sr *SocketRepository) SendMessage(message string) {
	for conn := range sr.Clients {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
