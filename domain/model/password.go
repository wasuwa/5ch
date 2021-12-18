package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func newPassword(password string) (Password, error) {
	hp, err := hashPassword(password)
	if err != nil {
		return "", err
	}
	if password == "" {
		return "", errors.New("パスワードを入力してください")
	}
	if len(password) < 6 {
		return "", errors.New("パスワードは6文字以上入力してください")
	}
	return Password(hp), nil
}

func hashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(h), err
}
