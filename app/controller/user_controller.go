package controller

import (
	_ "database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"opc-site/app/model"
	_ "opc-site/app/model"
	"opc-site/app/usecase"
)

type UserController struct {
	UserInteractor    usecase.UserInteractor
	SessionController *SessionController
}

const usersTableName = "users"

func NewUserController(sqlHandler SQLHandler) *UserController {
	return &UserController{
		UserInteractor: usecase.UserInteractor{
			UserRepository: &UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		SessionController: &SessionController{
			SessionInteractor: usecase.SessionInteractor{
				SessionRepository: &SessionRepository{
					SQLHandler: sqlHandler,
				},
			},
		},
	}
}

func (uc *UserController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405", http.StatusMethodNotAllowed)
		return
	}

	sessionCookie, err := r.Cookie(uc.SessionController.GetCookieName())
	if err != nil || len(sessionCookie.Value) > 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user *model.User
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	if len(user.Name) <= 0 || len(user.Password) <= 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	if uc.UserInteractor.Login(user) {
		uc.SessionController.Set(&w, user.Id)
		return
	}

	http.Error(w, "401", http.StatusUnauthorized)
}

func (uc *UserController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if !uc.SessionController.CheckSession(&w, r) {
		http.Error(w, "401", http.StatusUnauthorized)
		return
	}

	sessionCookie, err := r.Cookie(uc.SessionController.GetCookieName())
	if err != nil {
		log.Panic(err)
	}
	uc.SessionController.Unset(&w, sessionCookie.Value)
}
