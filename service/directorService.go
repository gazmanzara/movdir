package service

import (
	"github.com/gazmanzara/movdir/domain"
	"github.com/gazmanzara/movdir/dto"
	"github.com/gazmanzara/movdir/errs"
)

type DirectorService interface {
	GetAllDirectors() ([]dto.Director, *errs.AppError)
	GetDirectorById(id string) (*dto.Director, *errs.AppError)
	CreateDirector(director dto.Director) (*dto.Director, *errs.AppError)
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

func (s DefaultDirectorService) CreateDirector(director dto.Director) (*dto.Director, *errs.AppError) {
	err := director.Validate()
	if err != nil {
		return nil, err
	}

	p := domain.Director{
		Id:     director.Id,
		Name:   director.Name,
		Gender: director.Gender,
	}

	d, err := s.repo.Save(p)
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
