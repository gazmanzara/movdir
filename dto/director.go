package dto

import (
	"github.com/gazmanzara/movdir/errs"
	"strings"
)

type Director struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

func (d *Director) Validate() *errs.AppError {
	var errors []string

	if d.Id == 0 {
		errors = append(errors, "Director id is required")
	}
	if d.Name == "" {
		errors = append(errors, "Director name is required")
	}
	if d.Gender == 0 {
		errors = append(errors, "Director gender is required")
	}

	if len(errors) > 0 {
		return errs.NewUnprocessableEntityError(strings.Join(errors, ", "))
	} else {
		return nil
	}
}
