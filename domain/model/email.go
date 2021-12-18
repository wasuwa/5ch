package model

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	if isEmailTypeInvalid(email) {
		return "", errors.New("正しい形式でメールアドレスを入力してください")
	}
	return Email(email), nil
}

func isEmailTypeInvalid(email string) bool {
	s := `^[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z$`
	r := regexp.MustCompile(s)
	return !r.MatchString(email)
}
