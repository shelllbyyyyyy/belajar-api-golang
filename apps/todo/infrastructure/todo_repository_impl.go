package infrastructure

import (
	"api/first-go/apps/todo/domain"
	"api/first-go/common"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, model domain.Todo) error {
	query := `
		INSERT INTO public.todos (
			id, user_id, name, status, created_at, updated_at
		) VALUES (
			:id, :user_id, :name, :status, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)

	return err
}

func (r repository) FindById(ctx context.Context, id string) (*domain.Todo, error) {
	var todo domain.Todo

	query := `
		SELECT 
			id, user_id, name, status, is_archived, created_at, started_at, paused_at, finished_at, updated_at 
		FROM public.todos WHERE id = $1
	`

	err := r.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r repository) FindByUserId(ctx context.Context, userId string) ([]domain.Todo, error) {
	var todos []domain.Todo

	query := `
		SELECT 
			id, user_id, name, status, is_archived, created_at, started_at, paused_at, finished_at, updated_at 
		FROM public.todos WHERE user_id = $1
	`

	err := r.db.SelectContext(ctx, &todos, query, userId)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r repository) Update(ctx context.Context, payload *domain.UpdateToDoPayload) (bool, error) {
	query := `
		UPDATE public.todos 
		SET name = COALESCE($1, name),
    		status = COALESCE($2, status),
    		started_at = COALESCE($3, started_at),
    		paused_at = COALESCE($4, paused_at),
    		finished_at = COALESCE($5, finished_at),
    		is_archived = COALESCE($6, is_archived),
			updated_at = COALESCE($7, updated_at)
		WHERE id = $8
		RETURNING id, user_id, name, status, is_archived, created_at, started_at, paused_at, finished_at, updated_at
	`

	var todo domain.Todo
	err := r.db.GetContext(ctx, &todo, query, 
		payload.Name, 
		payload.Status, 
		payload.StartedAt, 
		payload.PausedAt, 
		payload.FinishedAt, 
		payload.IsArchived, 
		payload.UpdatedAt, 
		payload.Id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			
			return false, common.ErrUpdateToDoFailed
		}

		return false, err
	}

	return true, nil
}