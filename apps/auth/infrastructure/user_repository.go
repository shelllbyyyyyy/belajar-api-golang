package infrastructure

import (
	"api/first-go/apps/auth/domain"
	"context"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindById(ctx context.Context, id string) (*domain.User, error)
	CreateAuth(ctx context.Context, model domain.User) (err error)
	Update(ctx context.Context, id string, payload *domain.UpdateUserSchema) (bool, error)
}