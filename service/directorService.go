package service

import (
	"github.com/gazmanzara/movdir/app/errs"
	"github.com/gazmanzara/movdir/domain"
)

type DirectorService interface {
	GetAllDirectors() ([]domain.Director, *errs.AppError)
	GetDirectorById(id string) (*domain.Director, *errs.AppError)
}

type DefaultDirectorService struct {
	repo domain.DirectorRepository
}

func (s DefaultDirectorService) GetAllDirectors() ([]domain.Director, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultDirectorService) GetDirectorById(id string) (*domain.Director, *errs.AppError) {
	return s.repo.FindById(id)
}

func NewDirectorService(repo domain.DirectorRepository) DefaultDirectorService {
	return DefaultDirectorService{
		repo: repo,
	}
}
