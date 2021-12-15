package repository

import "bbs/domain/model"

type UserRepository interface {
	CreateUser(u *model.User) (*model.User, error)
}
