package application

import (
	"api/first-go/apps/todo/domain"
	"api/first-go/apps/todo/infrastructure"
	"api/first-go/common"
	"context"
)

type TodoUseCase struct {
	repo infrastructure.TodoRepository
}

func NewTodoUseCase(repo infrastructure.TodoRepository) TodoUseCase {
	return TodoUseCase{
		repo: repo,
	}
}

func (t TodoUseCase) Create(ctx context.Context, payload domain.CreateToDoPayload) error {
	todo, err := domain.NewTodo(payload)
	if err != nil {
		err = common.ErrCreateToDoFailed
		return err
	}

	return t.repo.Create(ctx, *todo)
}

func (t TodoUseCase) FindById(ctx context.Context, id string) (*domain.Todo, error) {
	model, err := t.repo.FindById(ctx, id)
	if err != nil {
		err = common.ErrNotFound
		return nil, err
	}

	return model, nil
}

func (t TodoUseCase) FindByUserId(ctx context.Context, userId string) ([]domain.Todo, error) {
	models, err := t.repo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return []domain.Todo{}, nil
	}

	return models, nil
}

func (t TodoUseCase) UpdateName(ctx context.Context, model *domain.Todo, name string) (bool, error) {
	if err := model.Update(name); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: &name,
		Status: nil,
		StartedAt: nil,
		PausedAt: nil,
		FinishedAt: nil,
		IsArchived: nil,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}

func (t TodoUseCase) Working(ctx context.Context, model *domain.Todo) (bool, error) {
	if err := model.Working(); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: nil,
		Status: &model.Status,
		StartedAt: model.StartedAt,
		PausedAt: nil,
		FinishedAt: nil,
		IsArchived: nil,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}

func (t TodoUseCase) Paused(ctx context.Context, model *domain.Todo) (bool, error) {
	if err := model.Paused(); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: nil,
		Status: &model.Status,
		StartedAt: nil,
		PausedAt: model.PausedAt,
		FinishedAt: nil,
		IsArchived: nil,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}

func (t TodoUseCase) Finished(ctx context.Context, model *domain.Todo) (bool, error) {
	if err := model.Complete(); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: nil,
		Status: &model.Status,
		StartedAt: nil,
		PausedAt: nil,
		FinishedAt: model.FinishedAt,
		IsArchived: nil,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}

func (t TodoUseCase) Archived(ctx context.Context, model *domain.Todo) (bool, error) {
	if err := model.Archived(); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: nil,
		Status: nil,
		StartedAt: nil,
		PausedAt: nil,
		FinishedAt: nil,
		IsArchived: &model.IsArchived,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}

func (t TodoUseCase) UnArchived(ctx context.Context, model *domain.Todo) (bool, error) {
	if err := model.UnArchived(); err != nil {
		return false, err
	}

	payload := &domain.UpdateToDoPayload{
		Name: nil,
		Status: nil,
		StartedAt: nil,
		PausedAt: nil,
		FinishedAt: nil,
		IsArchived: &model.IsArchived,
		UpdatedAt: model.UpdatedAt,
		Id: model.Id,
	}

	_, err := t.repo.Update(ctx, payload)
	if err != nil {
		err = common.ErrUpdateToDoFailed
		return false, err
	}

	return true, nil
}