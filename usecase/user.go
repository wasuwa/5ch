package usecase

import (
	"bbs/domain/model"
	"bbs/domain/repository"
)

type UserUsecase interface {
	Create(name, email, password string) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (uu *userUsecase) Create(name, email, password string) (*model.User, error) {
	u, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	u, err = uu.userRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
