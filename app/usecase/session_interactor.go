package usecase

import (
	"opc-site/app/model"
)

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (si *SessionInteractor) AddToDb(session *model.Session) {
	si.SessionRepository.AddToDb(session)
}

func (si *SessionInteractor) RemoveFromDb(uuid string) {
	si.SessionRepository.RemoveFromDb(uuid)
}

func (si *SessionInteractor) GetByUUID(uuid string) *model.Session {
	return si.SessionRepository.GetByUUID(uuid)
}

func (si *SessionInteractor) IsValidSession(uuid string) bool {
	return si.SessionRepository.IsValidSession(uuid)
}
