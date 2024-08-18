package domain

import "github.com/gazmanzara/movdir/auth/errs"

type AuthRepositoryStub struct{}

func (a *AuthRepositoryStub) FindOne(username string, password string) (*User, *errs.AppError) {
	if username == password {
		return nil, errs.NewBadRequestError("Invalid username or password")
	}
	return &User{
		Id:       1,
		Username: username,
		Password: password,
		Role:     1,
	}, nil
}

func (a *AuthRepositoryStub) Save(user User) (*User, *errs.AppError) {
	return &user, nil
}

func NewAuthRepositoryStub() AuthRepository {
	return &AuthRepositoryStub{}
}
