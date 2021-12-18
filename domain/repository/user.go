package repository

import "bbs/domain/model"

type UserRepository interface {
	All() (*[]model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	Update(u *model.User, id uint) error
}
