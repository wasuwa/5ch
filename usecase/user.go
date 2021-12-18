package usecase

import (
	"bbs/domain/model"
	"bbs/domain/repository"
)

type UserUsecase interface {
	Index() (*[]model.User, error)
	Find(id uint) (*model.User, error)
	Create(name, email, password string) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (uu *userUsecase) Index() (*[]model.User, error) {
	users, err := uu.userRepository.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) Find(id uint) (*model.User, error) {
	u, err := uu.userRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return u, nil
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
