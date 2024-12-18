package application

import (
	"api/first-go/auth/domain"
	"api/first-go/common"
	"api/first-go/configs"

	"context"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (model domain.User, err error)
	CreateAuth(ctx context.Context, model domain.User) (err error)
}

type AuthUseCase struct {
	repo Repository 
}

func NewAuthUseCase(repo Repository) (AuthUseCase) {
	return AuthUseCase{
		repo: repo,
	}
}

func (u AuthUseCase) Register(ctx context.Context, req RegisterRequestPayload) (err error) {
	payload := domain.RegisterUserSchema{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}

	user, err := domain.NewUser(payload)
	if err = user.Validate(); err != nil {
		return
	}

	if err = user.EncryptPassword(int(configs.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	model, err := u.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		if err != common.ErrNotFound {
			return
		}
	}

	if model.IsExists() {
		return common.ErrEmailAlreadyUsed
	}

	return u.repo.CreateAuth(ctx, *user)
}