package repository

import "bbs/domain/model"

type UserRepository interface {
	All() (*[]model.User, error)
	Find(id uint) (*model.User, error)
	Create(u *model.User) (*model.User, error)
}
