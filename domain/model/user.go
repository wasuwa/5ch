package model

import "time"

type User struct {
	Base
	Name     UserName
	Email    Email
	Password Password
}

func NewUser(name, email, password string) (*User, error) {
	n, err := newUserName(name)
	if err != nil {
		return nil, err
	}
	p, err := newPassword(password)
	if err != nil {
		return nil, err
	}
	e, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	u := &User{
		Name:     n,
		Email:    e,
		Password: p,
	}
	return u, nil
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) SetID(id uint) {
	u.ID = id
}

func (u *User) GetName() UserName {
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
