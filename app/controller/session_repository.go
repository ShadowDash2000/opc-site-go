package controller

import (
	_ "github.com/google/uuid"
	"log"
	"opc-site/app/model"
	"time"
)

type SessionRepository struct {
	SQLHandler SQLHandler
}

const sessionsTableName = "sessions"

func (sr *SessionRepository) AddToDb(session *model.Session) {
	query := "INSERT INTO " + sessionsTableName + " (userId, uuid, expirationTime) VALUES (?, ?, ?)"
	_, err := sr.SQLHandler.Db.Exec(
		query,
		session.UserId,
		session.UUID,
		session.ExpirationTime,
	)
	if err != nil {
		log.Panic(err)
	}
}

func (sr *SessionRepository) RemoveFromDb(uuid string) {
	query := "DELETE FROM " + sessionsTableName + " WHERE uuid=?"
	_, err := sr.SQLHandler.Db.Exec(
		query,
		uuid,
	)
	if err != nil {
		log.Panic(err)
	}
}

func (sr *SessionRepository) GetByUUID(uuid string) *model.Session {
	query := "SELECT * FROM " + sessionsTableName + " WHERE uuid=?"
	session := &model.Session{}
	sr.SQLHandler.Db.QueryRow(
		query,
		uuid,
	).Scan(
		&session.Id,
		&session.UserId,
		&session.ExpirationTime,
		&session.UUID,
	)

	return session
}

func (sr *SessionRepository) IsValidSession(uuid string) bool {
	session := sr.GetByUUID(uuid)
	if session.Id <= 0 {
		return false
	} else if session.ExpirationTime.Before(time.Now()) {
		return false
	}

	return true
}
