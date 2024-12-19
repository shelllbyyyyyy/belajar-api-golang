package controller

import (
	"api/first-go/apps/auth/application"
	"api/first-go/apps/auth/infrastructure"
	handlers "api/first-go/apps/auth/presentation/handler"
	"api/first-go/common"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoute(router fiber.Router, db *sqlx.DB) {
	repo := infrastructure.NewUserRepository(db)
	usecase := application.NewUserUseCase(repo)
	handler := handlers.NewUserHandler(usecase)

	_ = handler

	authRouter := router.Group("/api/v1/users")
	{
		authRouter.Get("/:email", common.CheckAuth(), handler.FindByEmail )
		authRouter.Get("/:id", common.CheckAuth(), handler.FindById)
		authRouter.Patch("/", common.CheckAuth(), handler.Update)
		authRouter.Delete("/", common.CheckAuth(), handler.Delete)
	}
}