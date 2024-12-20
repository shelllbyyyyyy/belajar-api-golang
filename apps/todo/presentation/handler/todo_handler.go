package handler

import (
	"api/first-go/apps/todo/application"
	"api/first-go/apps/todo/domain"
	"api/first-go/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


var Actions = []string{
	"changeName",
	"working",
	"paused",
	"complete",
	"archive",
	"unarchive",
}

type TodoHandler struct {
	useCase application.TodoUseCase
}

func NewTodoHandler(useCase application.TodoUseCase) TodoHandler {
	return TodoHandler{
		useCase: useCase,
	}
}

func (h TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	var req = application.CreateToDoRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
			common.WithHttpCode(http.StatusBadRequest),
			common.WithMessage("Update failed"),
		).Send(ctx)
	}

	payload := domain.CreateToDoPayload{
		Name: req.Name,
		UserId: ctx.Locals("id").(string),
	}

	err := h.useCase.Create(ctx.UserContext(), payload)
	if err != nil {
		myErr, ok := common.ErrorMapping[err.Error()]
		if !ok {
			myErr = common.ErrorGeneral
		}

		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
		).Send(ctx)
	}

	return common.NewResponse(
		common.WithHttpCode(http.StatusCreated),
		common.WithMessage("Create todo success"),
	).Send(ctx)
}

func (h TodoHandler) GetTodoById(ctx *fiber.Ctx) error {
	var param = ctx.Params("id")

	model, err := h.useCase.FindById(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("Todo not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	response := application.ToResponse(model)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("Todo Found"),
		common.WithData(response),
	).Send(ctx)
}

func (h TodoHandler) GetTodoByUserId(ctx *fiber.Ctx) error {
	var param = ctx.Locals("id").(string)

	model, err := h.useCase.FindByUserId(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("Todo not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	response := application.ToResponseList(model)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("Todo Found"),
		common.WithData(response),
	).Send(ctx)
}

func (h TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	var param = ctx.Params("id")
	var action = ctx.Params("action")


	actionExists := false
	for _, a := range Actions {
		if a == action {
			actionExists = true
			break
		}
	}

	if !actionExists {
		return common.NewResponse(
			common.WithMessage("Actions not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	model, err := h.useCase.FindById(ctx.UserContext(), param)
	if err != nil {
		return common.NewResponse(
			common.WithMessage("Todo not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	return h.Case(ctx, action, model)
}

func (h TodoHandler) Case(ctx *fiber.Ctx, action string, model *domain.Todo) error {
	switch action {
	case "changeName":
		var req = application.UpdateToDoRequest{}

		if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			   common.WithMessage(err.Error()),
			   common.WithError(myErr),
			   common.WithHttpCode(http.StatusBadRequest),
		   	   common.WithMessage("Update failed"),
			).Send(ctx)
		}
		data, err := h.useCase.UpdateName(ctx.UserContext(), model, *req.Name)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update name success"),
			common.WithData(data),
		).Send(ctx)

	case "working":
		data, err := h.useCase.Working(ctx.UserContext(), model)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update status success"),
			common.WithData(data),
		).Send(ctx)
	
	case "paused":
		data, err := h.useCase.Paused(ctx.UserContext(), model)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update status success"),
			common.WithData(data),
		).Send(ctx)
	
	case "complete":
		data, err := h.useCase.Finished(ctx.UserContext(), model)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update status success"),
			common.WithData(data),
		).Send(ctx)
	
	case "archive":
		data, err := h.useCase.Archived(ctx.UserContext(), model)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update status success"),
			common.WithData(data),
		).Send(ctx)
		
	case "unarchive":
		data, err := h.useCase.UnArchived(ctx.UserContext(), model)
		if err != nil {
			myErr, ok := common.ErrorMapping[err.Error()]
			if !ok {
				myErr = common.ErrorGeneral
			}

			return common.NewResponse(
				common.WithMessage(err.Error()),
				common.WithError(myErr),
			).Send(ctx)
		}

		return common.NewResponse(
			common.WithHttpCode(http.StatusOK),
			common.WithMessage("Update status success"),
			common.WithData(data),
		).Send(ctx)
	
	default:
		return common.NewResponse(
			common.WithMessage("Actions not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
}
}