package repository

import "bbs/domain/model"

type UserRepository interface {
	Index() (*[]model.User, error)
	Create(u *model.User) (*model.User, error)
}
