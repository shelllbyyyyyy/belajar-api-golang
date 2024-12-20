package application

import "api/first-go/apps/todo/domain"

func ToResponse(model *domain.Todo) TodoResponse {
	return TodoResponse{
		Id:         model.Id,
		UserId:     model.UserId,
		Name:       model.Name,
		Status:     string(model.Status),
		IsArchived: model.IsArchived,
		CreatedAt:  model.CreatedAt,
		StartedAt:  model.StartedAt,
		PausedAt:   model.PausedAt,
		FinishedAt: model.FinishedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func ToResponseList(models []domain.Todo) []TodoResponse {
	var res []TodoResponse
	for _, model := range models {
		res = append(res, ToResponse(&model))
	}
	return res
}