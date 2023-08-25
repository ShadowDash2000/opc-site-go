package model

import "time"

type Message struct {
	Id     int       `json:"id,omitempty"`
	UserId int       `json:"userId,omitempty"`
	Text   string    `json:"text,omitempty"`
	Time   time.Time `json:"time"`
}
