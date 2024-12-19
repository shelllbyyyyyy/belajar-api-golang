package infrastructure

import (
	"context"
	"database/sql"

	"api/first-go/apps/auth/domain"
	"api/first-go/common"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateAuth(ctx context.Context, model domain.User) (err error) {
	query := `
		INSERT INTO public.users (
			id, username, email, password, created_at, updated_at
		) VALUES (
			:id, :username, :email, :password, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)

	return
}

func (r repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
	SELECT id, username, email, password, role, created_at, updated_at
	FROM public.users
	WHERE email = $1`

	var user domain.User
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			
			return nil, common.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r repository) FindById(ctx context.Context, id string) (model domain.User, err error) {
	query := `
	SELECT id, username, email, password, role, created_at, updated_at
	FROM public.users
	WHERE id = $1`

	err = r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = common.ErrNotFound
			return
		}

		return
	}

	return
}