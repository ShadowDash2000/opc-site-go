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

func (ui *UserInteractor) Logout() bool {
	return ui.UserRepository.Logout()
}
