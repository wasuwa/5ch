package service

import "bbs/domain/repository"

type UserService struct {
	email string
}

func NewUser(email string) UserService {
	u := UserService{email: email}
	return u
}

func (us *UserService) getEmail() string {
	return us.email
}

func (us *UserService) Exists(ur repository.UserRepository) (bool, error) {
	u, err := ur.FindByEmail(us.getEmail())
	if err != nil {
		return false, err
	}
	if u != nil {
		return true, nil
	}
	return false, nil
}
