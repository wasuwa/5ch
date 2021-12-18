package infra

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

func (ur *UserRepository) All() (*[]model.User, error) {
	users := new([]model.User)
	result := ur.Conn.Find(users)
	if err := result.Error; err != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}

func (ur *UserRepository) Find(id uint) (*model.User, error) {
	u := new(model.User)
	result := ur.Conn.Where("id = ?", id).Find(u)
	if err := result.Error; err != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (ur *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.Conn.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) Update(u *model.User, id uint) error {
	result := ur.Conn.Where("id = ?", id).Updates(u)
	if err := result.Error; err != nil {
		return err
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
