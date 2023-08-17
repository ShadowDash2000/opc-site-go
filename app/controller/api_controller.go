package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opc-site/app/usecase"
	"strconv"
	"strings"
)

type ApiController struct {
	SQLHandler        SQLHandler
	SessionController *SessionController
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
	}
}

func (ac *ApiController) HandleApi(w http.ResponseWriter, r *http.Request) {
	if !ac.SessionController.CheckSession(&w, r) {
		http.Error(w, "401", http.StatusUnauthorized)
	}

	path := r.URL.Path
	trimmedPath := path[len("/api/"):]
	trimmedPath = strings.TrimSuffix(trimmedPath, "/")

	switch trimmedPath {
	case "get-user":
		ac.HandleGetUserById(w, r)
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
