package infrastructure

import (
	"api/first-go/apps/auth/domain"
	"context"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindById(ctx context.Context, id string) (model domain.User, err error)
	CreateAuth(ctx context.Context, model domain.User) (err error)
}