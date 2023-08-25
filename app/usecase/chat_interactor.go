package usecase

import "opc-site/app/model"

type ChatInteractor struct {
	ChatRepository ChatRepository
}

func (mi *ChatInteractor) Send(userId int, messageText string) *model.Message {
	return mi.ChatRepository.Send(userId, messageText)
}

func (mi *ChatInteractor) Get(messageId int) *model.Message {
	return mi.ChatRepository.Get(messageId)
}
