package domain

import (
	"github.com/gazmanzara/movdir/app/errs"
	"github.com/gazmanzara/movdir/domain/dto"
)

type Director struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Gender int    `db:"gender"`
}

func (d *Director) ToDTO() *dto.Director {
	return &dto.Director{
		Id:     d.Id,
		Name:   d.Name,
		Gender: d.Gender,
	}
}

type DirectorRepository interface {
	FindAll() ([]Director, *errs.AppError)
	FindById(id string) (*Director, *errs.AppError)
}
