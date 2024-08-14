package domain

import "github.com/gazmanzara/movdir/app/errs"

type Director struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}
type DirectorRepository interface {
	FindAll() ([]Director, *errs.AppError)
	FindById(id string) (*Director, *errs.AppError)
}
