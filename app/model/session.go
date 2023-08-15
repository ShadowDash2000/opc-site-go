package model

import (
	"time"
)

type Session struct {
	Id             int       `json:"id" db:"id"`
	UserId         int       `json:"userId" db:"userId"`
	ExpirationTime time.Time `json:"expirationTime" db:"expirationTime"`
	UUID           string    `json:"uuid" db:"uuid"`
}
