package usecase

import (
	"opc-site/app/model"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Login(user *model.User) bool {
	return ui.UserRepository.Authorize(user)
}

func (ui *UserInteractor) Register(user *model.User) bool {
	return ui.UserRepository.Add(user)
}

func (ui *UserInteractor) GetUserById(userId int) *model.User {
	return ui.UserRepository.GetUserById(userId)
}
