package usecase

import (
	"opc-site/app/model"
)

type UserRepository interface {
	Add(user *model.User) bool
	Authorize(user *model.User) bool
	Logout() bool
}
