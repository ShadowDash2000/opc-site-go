package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"opc-site/app/usecase"
)

type SocketController struct {
	SocketInteractor usecase.SocketInteractor
}

func NewSocketController() *SocketController {
	return &SocketController{
		SocketInteractor: usecase.SocketInteractor{
			SocketRepository: &SocketRepository{
				Clients: make(map[*websocket.Conn]bool),
			},
		},
	}
}

func (sc *SocketController) HandleConnection(w http.ResponseWriter, r *http.Request, callback func(string, *websocket.Conn)) {
	conn := sc.SocketInteractor.UpgradeConnection(w, r)
	defer conn.Close()

	sc.SocketInteractor.AddConnection(conn)
	defer sc.SocketInteractor.DeleteConnection(conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			log.Println("WebSocket read error:", err)
			break
		}

		callback(string(message), conn)
		fmt.Println("Received message:", string(message))
	}
}

func (sc *SocketController) SendMessage(message string) {
	sc.SocketInteractor.SendMessage(message)
}
