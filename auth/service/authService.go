package service

import (
	"github.com/gazmanzara/movdir/auth/domain"
	"github.com/gazmanzara/movdir/auth/dto"
	"github.com/gazmanzara/movdir/auth/errs"
)

type AuthService interface {
	Login(payload dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (s *DefaultAuthService) Login(payload dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	res, err := s.repo.Login(payload.Username, payload.Password)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: res,
	}, nil
}

func NewAuthService(repo domain.AuthRepository) AuthService {
	return &DefaultAuthService{
		repo: repo,
	}
}
