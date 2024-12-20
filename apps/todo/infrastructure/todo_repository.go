package infrastructure

import (
	"api/first-go/apps/todo/domain"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, model domain.Todo) error
	FindById(ctx context.Context, id string) (*domain.Todo, error)
	FindByUserId(ctx context.Context, userId string) ([]domain.Todo, error)
	Update(ctx context.Context, payload *domain.UpdateToDoPayload) (bool, error)
}