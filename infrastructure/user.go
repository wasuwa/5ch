package infrastructure

import (
	"bbs/domain/model"
	"bbs/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	if err := ur.Conn.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) FindUser(id uint) (*model.User, error) {

}

func (ur *UserRepository) UpdateUser(u *model.User) (*model.User, error) {

}

func (ur *UserRepository) DestroyUser(id uint) error {

}
