package application

import (
	"api/first-go/apps/auth/domain"
	"api/first-go/apps/auth/infrastructure"
	"api/first-go/common"
	"api/first-go/configs"

	"context"
)

type AuthUseCase struct {
	repo infrastructure.UserRepository 
}

func NewAuthUseCase(repo infrastructure.UserRepository) (AuthUseCase) {
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

func (u AuthUseCase) Login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	payload := domain.LoginUserSchema{
		Email: req.Email,
		Password: req.Password,
	}

	model, err := u.repo.FindByEmail(ctx, payload.Email)
	if err != nil { 
		err = common.ErrNotFound
		
		return
	}

	if err = model.ComparePassword(payload.Password); err != nil {
		err = common.ErrPasswordNotMatch

		return
	}
	

	token, err = model.GenerateToken()
	return 
}