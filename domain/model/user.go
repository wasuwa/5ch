package model

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	if err := validation(name, email, password); err != nil {
		return nil, err
	}
	p, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Name:     name,
		Email:    email,
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

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.CreatedAt
}

func validation(name, email, password string) error {
	if name == "" {
		return errors.New("名前を入力してください")
	}
	if email == "" {
		return errors.New("メールアドレスを入力してください")
	}
	if isEmailTypeInvalid(email) {
		return errors.New("正しい形式でメールアドレスを入力してください")
	}
	if password == "" {
		return errors.New("パスワードを入力してください")
	}
	if len(password) < 6 {
		return errors.New("パスワードは6文字以上入力してください")
	}
	return nil
}

func isEmailTypeInvalid(email string) bool {
	s := `^[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z$`
	r := regexp.MustCompile(s)
	return !r.MatchString(email)
}

func hashPassword(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(h), err
}
