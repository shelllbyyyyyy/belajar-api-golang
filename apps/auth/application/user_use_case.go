package application

import (
	"api/first-go/apps/auth/domain"
	"api/first-go/apps/auth/infrastructure"
	"api/first-go/common"

	"context"
)

type UserUseCase struct {
	repo infrastructure.UserRepository
}

func NewUserUseCase(repo infrastructure.UserRepository) UserUseCase {
	return UserUseCase{
		repo: repo,
	}
}

func (u UserUseCase) FindByEmail(ctx context.Context, email string) (model domain.User, err error) {
	model, err = u.repo.FindByEmail(ctx, email)
	if err != nil {
		if err != common.ErrNotFound {
			return
		}
	}

	return
}

func (u UserUseCase) FindById(ctx context.Context, id string) (model domain.User, err error) {
	model, err = u.repo.FindById(ctx, id)
	if err != nil {
		if err != common.ErrNotFound {
			return
		}
	}

	return
}
