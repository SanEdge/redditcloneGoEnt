package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
	"github.com/renaldyhidayatt/redditgoent/ent/verificationtoken"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
	"github.com/renaldyhidayatt/redditgoent/security"
	"github.com/renaldyhidayatt/redditgoent/utils"
)

type AuthRepository = interfaces.IAuthRepository

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, context context.Context) *authRepository {
	return &authRepository{db: db, context: context}
}

func (r *authRepository) Login(input request.LoginRequest) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).Where(user.EnabledEQ(true)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	checkPassword := security.VerifyPassword(user.Password, input.Password)

	if checkPassword != nil {
		return nil, fmt.Errorf("failed checkhash password: %w", err)

	}

	return user, nil
}

func (r *authRepository) Register(input request.RegisterRequest) (*ent.User, error) {

	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).Where(user.EnabledEQ(true)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("email already exits")
	}

	newUser, err := r.db.User.Create().SetUsername(input.Username).SetEmail(input.Email).SetPassword(security.HashPassword(input.Password)).SetEnabled(false).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create user: %w", err)
	}

	token, err := r.GenerateVerificationToken(newUser)

	if err != nil {
		return nil, fmt.Errorf("failed create verificationtoken: %w", err)
	}

	utils.SendMail(request.NotificationEmail{
		Subject:   "Please Activate your Account",
		Recipient: "dragon@gmail.com",
		Body: "Thank you for signing up to Spring Reddit, " +
			"please click on the below url to activate your account : " +
			"http://localhost:5000/api/auth/accountVerification/" + token.Token,
	})

	return newUser, nil

}

func (r *authRepository) GenerateVerificationToken(user *ent.User) (*ent.VerificationToken, error) {
	token := utils.RandStringBytes(30)
	verification, err := r.db.VerificationToken.Create().SetToken(token).AddUser(user).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create verificationtoken: %w", err)
	}

	return verification, nil

}

func (r *authRepository) VerifyToken(verication string) (*ent.User, error) {
	verify, err := r.db.VerificationToken.Query().Where(verificationtoken.TokenEQ(verication)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed fetch verificationtoken: %w", err)
	}
	verifyUser, err := verify.QueryUser().First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed fetch verificationtoken: %w", err)
	}

	user, err := r.db.User.Query().Where(user.EmailEQ(verifyUser.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed fetch user: %w", err)
	}
	user, err = r.db.User.UpdateOne(user).SetEnabled(true).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update user: %w", err)
	}

	return user, nil
}
