package usecase

import (
	"opc-site/app/model"
)

type SessionRepository interface {
	AddToDb(session *model.Session)
	RemoveFromDb(uuid string)
	GetByUUID(uuid string) *model.Session
	IsValidSession(uuid string) bool
}
