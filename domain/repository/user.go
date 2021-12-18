package repository

import "ddd-sample/domain/model"

type UserRepository interface {
	All() (*[]model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByEmail(u *model.User) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	Update(u *model.User, id uint) error
}
