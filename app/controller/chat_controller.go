package controller

import (
	"opc-site/app/model"
	"opc-site/app/usecase"
)

type ChatController struct {
	ChatInteractor    usecase.ChatInteractor
	SessionController *SessionController
}

func NewChatController(sqlHandler SQLHandler) *ChatController {
	return &ChatController{
		ChatInteractor: usecase.ChatInteractor{
			ChatRepository: &ChatRepository{
				SQLHandler: sqlHandler,
			},
		},
		SessionController: NewSessionController(sqlHandler),
	}
}

func (cc *ChatController) Send(userId int, messageText string) *model.Message {
	return cc.ChatInteractor.Send(userId, messageText)
}

func (cc *ChatController) Get(messageId int) *model.Message {
	return cc.ChatInteractor.Get(messageId)
}
