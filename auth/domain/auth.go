package domain

import "github.com/gazmanzara/movdir/auth/errs"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type AuthRepository interface {
	FindOne(username string, password string) (*User, *errs.AppError)
	Save(user User) (*User, *errs.AppError)
}
