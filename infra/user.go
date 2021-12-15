package infra

import (
	"bbs/domain/model"
	"bbs/domain/repository"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) Create(u *model.User) (*model.User, error) {
	fmt.Println(u)
	if err := ur.Conn.Debug().Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
