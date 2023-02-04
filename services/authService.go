package services

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
	"github.com/renaldyhidayatt/redditgoent/repository"
)

type AuthService = interfaces.IAuthService

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) Register(input request.RegisterRequest) (*ent.User, error) {
	var request request.RegisterRequest

	request.Username = input.Username
	request.Email = input.Email
	request.Password = input.Password

	res, err := s.repository.Register(input)

	return res, err
}

func (s *authService) Login(input request.LoginRequest) (*ent.User, error) {
	var request request.LoginRequest

	request.Email = input.Email
	request.Password = input.Password

	res, err := s.repository.Login(input)

	return res, err
}

func (s *authService) GenerateVerificationToken(user *ent.User) (*ent.VerificationToken, error) {
	res, err := s.repository.GenerateVerificationToken(user)

	return res, err
}

func (s *authService) VerifyToken(verification string) (*ent.User, error) {
	res, err := s.repository.VerifyToken(verification)

	return res, err
}
