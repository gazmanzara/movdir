package service

import (
	"github.com/gazmanzara/movdir/app/errs"
	"github.com/gazmanzara/movdir/domain"
	"github.com/gazmanzara/movdir/domain/dto"
)

type DirectorService interface {
	GetAllDirectors() ([]dto.Director, *errs.AppError)
	GetDirectorById(id string) (*dto.Director, *errs.AppError)
}

type DefaultDirectorService struct {
	repo domain.DirectorRepository
}

func (s DefaultDirectorService) GetAllDirectors() ([]dto.Director, *errs.AppError) {
	d, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	directorsDTO := make([]dto.Director, len(d))
	for i, director := range d {
		directorsDTO[i] = *director.ToDTO()
	}

	return directorsDTO, nil
}

func (s DefaultDirectorService) GetDirectorById(id string) (*dto.Director, *errs.AppError) {
	d, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return d.ToDTO(), nil
}

func NewDirectorService(repo domain.DirectorRepository) DefaultDirectorService {
	return DefaultDirectorService{
		repo: repo,
	}
}
