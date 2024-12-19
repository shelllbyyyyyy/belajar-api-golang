package application

type RegisterRequestPayload struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequestPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenPayload struct {
	Id, Role string
}

type UpdatePayload struct {
	Username  *string `json:"username" validate:"optional"`
	Email     *string `json:"email" validate:"optional"`
	Password  *string `json:"password" validate:"optional"`
	Role      *string `json:"role" validate:"optional"`
	IsDeleted *bool   `json:"is_deleted" validate:"optional"`
}

type DeleteUser struct {
	Password string `json:"password" validate:"required,min=6"`
}