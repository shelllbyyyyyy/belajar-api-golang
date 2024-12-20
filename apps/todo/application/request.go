package application

type CreateToDoRequest struct {
	Name string `json:"name" validate:"required,min=4,max=100"`
}

type UpdateToDoRequest struct {
	Name *string `json:"name" validate:"optional,min=4,max=100"`
}