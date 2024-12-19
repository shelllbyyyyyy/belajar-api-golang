package domain

type RegisterUserSchema struct {
	Username, Email, Password string
}

type LoginUserSchema struct {
	Email, Password string
}

type UpdateUserSchema struct {
	Username  *string
	Email     *string
	Password  *string
	Role      *string
	IsDeleted *bool
}