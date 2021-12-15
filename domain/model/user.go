package model

import (
	"errors"
	"regexp"
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
	u := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return u, nil
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
