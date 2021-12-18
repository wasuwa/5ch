package model

import (
	"errors"
	"regexp"
)

type Email string

func newEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New("メールアドレスを入力してください")
	}
	if !isTypeValid(email) {
		return "", errors.New("正しい形式でメールアドレスを入力してください")
	}
	return Email(email), nil
}

func isTypeValid(email string) bool {
	s := `^[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z$`
	r := regexp.MustCompile(s)
	return r.MatchString(email)
}
