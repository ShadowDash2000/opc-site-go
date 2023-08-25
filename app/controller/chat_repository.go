package controller

import (
	"log"
	"opc-site/app/model"
	"time"
)

type ChatRepository struct {
	SQLHandler SQLHandler
}

const messagesTableName = "messages"

func (cr *ChatRepository) Send(userId int, messageText string) *model.Message {
	message := &model.Message{
		UserId: userId,
		Text:   messageText,
		Time:   time.Now(),
	}

	query := "INSERT INTO " + messagesTableName + " (userId, text, time) VALUES (?, ?, ?)"
	_, err := cr.SQLHandler.Db.Exec(
		query,
		message.UserId,
		message.Text,
		message.Time,
	)
	if err != nil {
		log.Panic(err.Error())
	}

	return message
}

func (cr *ChatRepository) Get(messageId int) *model.Message {
	return &model.Message{
		Id:     1,
		UserId: 1,
		Text:   "GET TEST MESSAGE",
		Time:   time.Now(),
	}
}
