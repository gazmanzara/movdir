package service

import (
	"github.com/gazmanzara/movdir/auth/domain"
	"github.com/gazmanzara/movdir/auth/dto"
	"github.com/gazmanzara/movdir/auth/errs"
)

type AuthService interface {
	Login(payload dto.LoginRequest) (*dto.AuthResponse, *errs.AppError)
	Register(payload dto.RegisterRequest) (*dto.AuthResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (s *DefaultAuthService) Login(payload dto.LoginRequest) (*dto.AuthResponse, *errs.AppError) {
	res, err := s.repo.FindOne(payload.Username, payload.Password)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: res.Username + res.Password,
	}, nil
}

func (s *DefaultAuthService) Register(payload dto.RegisterRequest) (*dto.AuthResponse, *errs.AppError) {
	newUser := domain.User{
		Username: payload.Username,
		Password: payload.Password,
		Role:     1,
	}

	res, err := s.repo.Save(newUser)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: res.Username + res.Password,
	}, nil
}

func NewAuthService(repo domain.AuthRepository) AuthService {
	return &DefaultAuthService{
		repo: repo,
	}
}
