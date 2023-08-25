package controller

import (
	"github.com/google/uuid"
	"net/http"
	"opc-site/app/model"
	"opc-site/app/usecase"
	"time"
)

type SessionController struct {
	SessionInteractor usecase.SessionInteractor
}

const sessionCookieName = "session_uuid"

func NewSessionController(sqlHandler SQLHandler) *SessionController {
	return &SessionController{
		SessionInteractor: usecase.SessionInteractor{
			SessionRepository: &SessionRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (sc *SessionController) Set(w *http.ResponseWriter, userId int) {
	session := &model.Session{
		UserId:         userId,
		ExpirationTime: time.Now().Add(60 * 60 * 24 * 7 * time.Second),
		UUID:           uuid.NewString(),
	}

	sc.SessionInteractor.AddToDb(session)

	http.SetCookie(*w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   session.UUID,
		Expires: session.ExpirationTime,
	})
}

func (sc *SessionController) Unset(w *http.ResponseWriter, sessionUuid string) {
	sc.SessionInteractor.RemoveFromDb(sessionUuid)

	http.SetCookie(*w, &http.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
	})
}

func (sc *SessionController) GetCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(sessionCookieName)
}

func (sc *SessionController) CheckSession(w *http.ResponseWriter, r *http.Request) bool {
	sessionCookie, err := sc.GetCookie(r)
	if err != nil {
		return false
	}

	if sc.SessionInteractor.IsValidSession(sessionCookie.Value) {
		sc.Unset(w, sessionCookie.Value)
		return false
	}

	return true
}

func (sc *SessionController) GetCookieName() string {
	return sessionCookieName
}
