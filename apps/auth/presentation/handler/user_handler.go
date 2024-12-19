package presentation

import (
	"api/first-go/apps/auth/application"
	"api/first-go/common"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    usecase application.UserUseCase
}

func NewUserHandler(uc application.UserUseCase) UserHandler {
    return UserHandler{
		usecase: uc,
	}
}

func (h UserHandler) FindByEmail(ctx *fiber.Ctx) error {
    var param = ctx.Params("email")

	model, err := h.usecase.FindByEmail(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("User not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	response := application.ToUserResponse(model)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("User Found"),
		common.WithData(response),
	).Send(ctx)
}

func (h UserHandler) FindById(ctx *fiber.Ctx) error {
    var param = ctx.Params("id")

	model, err := h.usecase.FindById(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("User not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	response := application.ToUserResponse(model)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("User Found"),
		common.WithData(response),
	).Send(ctx)
}

func (h UserHandler) Update(ctx *fiber.Ctx) error {
    var param = ctx.Locals("id").(string)
	var req = application.UpdatePayload{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
			common.WithHttpCode(http.StatusBadRequest),
			common.WithMessage("register fail"),
		).Send(ctx)
	}

	model, err := h.usecase.FindById(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("User not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	result, err := h.usecase.Update(ctx.UserContext(), model, req)
	if err != nil {
		return common.NewResponse(
			common.WithMessage("Update user failed"),
			common.WithError(common.ErrorBadRequest),
		).Send(ctx)
	}

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("Update user success"),
		common.WithData(result),
	).Send(ctx)
}

func (h UserHandler) Delete(ctx *fiber.Ctx) error {
    var param = ctx.Locals("id").(string)
	var req = application.DeleteUser{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
			common.WithHttpCode(http.StatusBadRequest),
			common.WithMessage("register fail"),
		).Send(ctx)
	}

	model, err := h.usecase.FindById(ctx.UserContext(), param)
    if err != nil {
		return common.NewResponse(
			common.WithMessage("User not found"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	result, err := h.usecase.Delete(ctx.UserContext(), model, req.Password)
	if err != nil {
		return common.NewResponse(
			common.WithMessage("Delete user failed"),
			common.WithError(common.ErrorBadRequest),
		).Send(ctx)
	}

	ctx.ClearCookie()

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("Delete user success"),
		common.WithData(result),
	).Send(ctx)
}