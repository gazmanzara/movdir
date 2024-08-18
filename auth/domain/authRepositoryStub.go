package domain

import "github.com/gazmanzara/movdir/auth/errs"

type AuthRepositoryStub struct {
	token string
}

func (a *AuthRepositoryStub) Login(username string, password string) (string, *errs.AppError) {
	if username == password {
		return "", errs.NewBadRequestError("Invalid username or password")
	}
	return a.token, nil
}

func NewAuthRepositoryStub() AuthRepository {
	return &AuthRepositoryStub{
		token: "token",
	}
}
