package model

import (
	"errors"
	"regexp"
)

type User struct {
	base
	name     string
	email    string
	password string
}

func New(n, e, p string) (*User, error) {
	if err := validation(n, e, p); err != nil {
		return nil, err
	}
	u := &User{
		name:     n,
		email:    e,
		password: p,
	}
	return u, nil
}

func validation(n, e, p string) error {
	if n == "" {
		return errors.New("名前を入力してください")
	}
	if e == "" {
		return errors.New("メールアドレスを入力してください")
	}
	if isEmailTypeValid(e) {
		return errors.New("正しい形式でメールアドレスを入力してください")
	}
	if p == "" {
		return errors.New("パスワードを入力してください")
	}
	if len(n) >= 6 {
		return errors.New("パスワードは6文字以上入力してください")
	}
	return nil
}

func isEmailTypeValid(e string) bool {
	str := `^[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z$`
	r := regexp.MustCompile(str)
	return r.MatchString(e)
}
