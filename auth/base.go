package auth

import (
	"api/first-go/auth/application"
	"api/first-go/auth/infrastructure"
	"api/first-go/auth/presentation"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const apiVersion = "/api/v1"

func Init(router fiber.Router, db *sqlx.DB) {
	repo := infrastructure.NewRepository(db)
	usecase := application.NewAuthUseCase(repo)
	handler := presentation.NewAuthHandler(usecase)

	_ = handler

	authRouter := router.Group(apiVersion + "/auth")
	{
		authRouter.Post("/register", handler.Register)
	}
}