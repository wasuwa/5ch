package model

import (
	"errors"
	"time"
)

type User struct {
	Base
	Name     string
	Email    Email
	Password Password
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("名前を入力してください")
	}
	e, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	p, err := newPassword(password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Name:     name,
		Email:    e,
		Password: p,
	}
	return u, nil
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() Email {
	return u.Email
}

func (u *User) GetPassword() Password {
	return u.Password
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.CreatedAt
}
