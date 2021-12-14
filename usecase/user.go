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

}
