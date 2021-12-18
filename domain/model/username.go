package model

import "errors"

type UserName string

func newUserName(name string) (UserName, error) {
	if name == "" {
		return "", errors.New("名前を入力してください")
	}
	return UserName(name), nil
}
