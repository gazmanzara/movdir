package service

import "github.com/gazmanzara/movdir/domain"

type DirectorService interface {
	GetAllDirectors() ([]domain.Director, error)
}

type DefaultDirectorService struct {
	repo domain.DirectorRepository
}

func (s DefaultDirectorService) GetAllDirectors() ([]domain.Director, error) {
	return s.repo.FindAll()
}

func NewDirectorService(repo domain.DirectorRepository) DefaultDirectorService {
	return DefaultDirectorService{
		repo: repo,
	}
}
