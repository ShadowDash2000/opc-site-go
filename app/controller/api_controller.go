package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"opc-site/app/model"
	"opc-site/app/usecase"
	"strconv"
	"strings"
)

type ApiController struct {
	SQLHandler        SQLHandler
	SessionController *SessionController
	SocketController  *SocketController
	ChatController    *ChatController
}

func NewApiController(sqlHandler SQLHandler) *ApiController {
	return &ApiController{
		SQLHandler: sqlHandler,
		SessionController: &SessionController{
			SessionInteractor: usecase.SessionInteractor{
				SessionRepository: &SessionRepository{
					SQLHandler: sqlHandler,
				},
			},
		},
		SocketController: NewSocketController(),
		ChatController:   NewChatController(sqlHandler),
	}
}

func (ac *ApiController) HandleApi(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	trimmedPath := path[len("/api/"):]
	trimmedPath = strings.TrimSuffix(trimmedPath, "/")

	switch trimmedPath {
	case "user/get-user":
		ac.HandleGetUserById(w, r)
	case "chat":
		ac.SendChatMessage(w, r)
	}

	fmt.Println("Запрос к:", trimmedPath)
}

func (ac *ApiController) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405", http.StatusMethodNotAllowed)
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil || userId <= 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	sessionCookie, _ := r.Cookie(sessionCookieName)
	session := ac.SessionController.SessionInteractor.GetByUUID(sessionCookie.Value)
	if session.UserId != userId {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	userInteractor := usecase.UserInteractor{
		UserRepository: &UserRepository{
			SQLHandler: ac.SQLHandler,
		},
	}
	user := userInteractor.GetUserById(userId)
	if user.Id <= 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

func (ac *ApiController) SendChatMessage(w http.ResponseWriter, r *http.Request) {
	ac.SocketController.HandleConnection(w, r, func(messageText string, conn *websocket.Conn) {
		sessionCookie, err := ac.SessionController.GetCookie(r)
		if err != nil {
			return
		}

		if !ac.SessionController.SessionInteractor.IsValidSession(sessionCookie.Value) {
			return
		}
		session := ac.SessionController.SessionInteractor.GetByUUID(sessionCookie.Value)

		message := &model.Message{}
		err = json.Unmarshal([]byte(messageText), &message)
		if err != nil {
			return
		}

		message = ac.ChatController.Send(session.UserId, message.Text)
		messageJson, err := json.Marshal(message)
		ac.SocketController.SendMessage(string(messageJson))
	})
}
