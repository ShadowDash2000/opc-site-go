package usecase

import (
	"opc-site/app/model"
)

type UserRepository interface {
	Add(user *model.User) bool
	Authorize(user *model.User) bool
	GetUserById(userId int) *model.User
	ConvertPassword(password string) string
}
