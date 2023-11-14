package service

import (
	"AILN/app/common/jwt"
	"AILN/app/model/user"
	"fmt"
)

type AuthService struct{}

func (a *AuthService) Login(username, password string) (token string, err error) {
	if !user.ExistUP(username, password) {
		return "", fmt.Errorf("username or password is wrong")
	}

	result, err := user.FindOneByUP(username, password)
	if err != nil {
		return "", err
	}

	return jwt.SignToken(result.ID)
}
