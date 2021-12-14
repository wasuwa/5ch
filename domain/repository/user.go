package repository

import "bbs/domain/model"

type UserRepository interface {
	CreateUser(u *model.User) (*model.User, error)
	FindUser(id uint) (*model.User, error)
	UpdateUser(u *model.User) (*model.User, error)
	DestroyUser(id uint) error
}
