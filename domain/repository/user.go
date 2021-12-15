package repository

import "bbs/domain/model"

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
}
