package controller

import (
	"api/first-go/apps/todo/application"
	"api/first-go/apps/todo/infrastructure"
	handler "api/first-go/apps/todo/presentation/handler"
	"api/first-go/common"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TodoRoute(router fiber.Router, db *sqlx.DB) {
	repo := infrastructure.NewTodoRepository(db)
	todo := application.NewTodoUseCase(repo)
	handler := handler.NewTodoHandler(todo)

	_ = handler

	todoRouter := router.Group("/api/v1/todo")
	{
		todoRouter.Post("/", common.CheckAuth(), handler.CreateTodo)
		todoRouter.Get("/:id", common.CheckAuth(), handler.GetTodoById)
		todoRouter.Get("/", common.CheckAuth(), handler.GetTodoByUserId)
		todoRouter.Patch("/:action/:id", common.CheckAuth(), handler.UpdateTodo)
	}
}