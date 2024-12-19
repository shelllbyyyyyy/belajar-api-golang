package application

import (
	"api/first-go/apps/auth/domain"
	"api/first-go/apps/auth/infrastructure"
	"api/first-go/common"
	"api/first-go/configs"
	"api/first-go/util"

	"context"
)

type Token struct {
	AccessToken string 	`json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

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

func (u AuthUseCase) Login(ctx context.Context, req LoginRequestPayload) (*Token, error) {
	payload := domain.LoginUserSchema{
		Email: req.Email,
		Password: req.Password,
	}

	model, err := u.repo.FindByEmail(ctx, payload.Email)
	if err != nil { 
		err = common.ErrNotFound
		
		return nil, err
	}

	if err = model.ComparePassword(payload.Password); err != nil {
		err = common.ErrPasswordNotMatch

		return nil, err
	}

	access_token, err := model.GenerateToken(15)
	if err != nil {
		return nil, err
	}

	refresh_token, err := model.GenerateToken(60 * 7 * 24)
	if err != nil {
		return nil, err
	}
	
	return &Token{AccessToken: access_token, RefreshToken: refresh_token}, nil
}

func (u AuthUseCase) Refresh(ctx context.Context, req TokenPayload) (string, error) {
	token, err := util.GenerateToken(req.Id, req.Email, 15)
	if err != nil {
		return "", err
	}
	
	return token, nil	
}