package domain

import "github.com/gazmanzara/movdir/auth/errs"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthRepository interface {
	Login(username string, password string) (string, *errs.AppError)
}
