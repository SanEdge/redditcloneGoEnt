package interfaces

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
)

type IAuthRepository interface {
	Login(input request.LoginRequest) (*ent.User, error)
	Register(input request.RegisterRequest) (*ent.User, error)
	GenerateVerificationToken(usr *ent.User) (*ent.VerificationToken, error)
	VerifyToken(verification string) (*ent.User, error)
}

type IAuthService interface {
	Login(input request.LoginRequest) (*ent.User, error)
	Register(input request.RegisterRequest) (*ent.User, error)
	GenerateVerificationToken(usr *ent.User) (*ent.VerificationToken, error)
	VerifyToken(verification string) (*ent.User, error)
}
