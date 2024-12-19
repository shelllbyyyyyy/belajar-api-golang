package application

import (
	"api/first-go/apps/auth/domain"
	"api/first-go/apps/auth/infrastructure"
	"api/first-go/configs"

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

func (u UserUseCase) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	model, err := u.repo.FindByEmail(ctx, email)
	if  err != nil {
		return nil, err
	}

	return model, nil
}

func (u UserUseCase) FindById(ctx context.Context, id string) (*domain.User, error) {
	model, err := u.repo.FindById(ctx, id)
	if  err != nil {
		return nil, err
	}

	return model, nil
}

func (u UserUseCase) Update(ctx context.Context, model *domain.User, payload UpdatePayload) (bool, error) {
	if payload.Password != nil {
		if err := model.ComparePassword(*payload.Password); err != nil {
			return false, err
		}

		model.Password = *payload.Password
		if err := model.EncryptPassword(int(configs.Cfg.App.Encryption.Salt)); err != nil {
			return false, err
		}
	}
	
	var update *domain.UpdateUserSchema = &domain.UpdateUserSchema{
		Username: payload.Username,
		Email: payload.Email,
		IsDeleted: payload.IsDeleted,
		Role: payload.Role,
		Password: func() *string {
			if payload.Password != nil {
				return &model.Password
			}
			return nil
		}(),
	}

	result, err := u.repo.Update(ctx, model.Id,  update)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (u UserUseCase) Delete(ctx context.Context, model *domain.User, password string) (bool, error) {
	if err := model.ComparePassword(password); err != nil {
		return false, err
	}

	isDeletedTrue := true
	
	var update *domain.UpdateUserSchema = &domain.UpdateUserSchema{
		Username: nil,
		Email: nil,
		IsDeleted: &isDeletedTrue,
		Role: nil,
		Password: nil,
	}

	result, err := u.repo.Update(ctx, model.Id,  update)
	if err != nil {
		return false, err
	}

	return result, nil
}
