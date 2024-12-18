package domain

type RegisterUserSchema struct {
	Username, Email, Password string
}

type LoginUserSchema struct {
	Email, Password string
}