package usecase

import "opc-site/app/model"

type ChatRepository interface {
	Send(userId int, messageText string) *model.Message
	Get(messageId int) *model.Message
}
