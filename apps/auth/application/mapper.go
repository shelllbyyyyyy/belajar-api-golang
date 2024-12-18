package application

import "api/first-go/apps/auth/domain"

func ToUserResponse(model domain.User) UserResponse {
	return UserResponse{
		Id:   model.Id,
		Username: model.Username,
		Email: model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}