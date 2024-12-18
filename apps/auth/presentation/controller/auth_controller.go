package controller

import (
	"api/first-go/apps/auth/application"
	"api/first-go/apps/auth/infrastructure"
	handler "api/first-go/apps/auth/presentation/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func AuthRoute(router fiber.Router, db *sqlx.DB) {
	repo := infrastructure.NewUserRepository(db)
	usecase := application.NewAuthUseCase(repo)
	handler := handler.NewAuthHandler(usecase)

	_ = handler

	authRouter := router.Group("/api/v1/auth")
	{
		authRouter.Post("/register", handler.Register)
		authRouter.Post("/login", handler.Login)
	}
}