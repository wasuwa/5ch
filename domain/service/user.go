package service

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"
)

type UserService struct {
	email model.Email
}

func NewUser(email model.Email) UserService {
	u := UserService{email: email}
	return u
}

func (us *UserService) getEmail() model.Email {
	return us.email
}

func (us *UserService) Exists(ur repository.UserRepository, u *model.User) (bool, error) {
	_, err := ur.FindByEmail(u)
	if err != nil {
		return true, err
	}
	return false, nil
}
