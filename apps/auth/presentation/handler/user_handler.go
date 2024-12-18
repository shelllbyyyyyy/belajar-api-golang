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
		myErr, ok := common.ErrorMapping[err.Error()]
		if !ok {
			myErr = common.ErrorGeneral
		}

		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
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
		myErr, ok := common.ErrorMapping[err.Error()]
		if !ok {
			myErr = common.ErrorGeneral
		}

		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
		).Send(ctx)
	} 

	response := application.ToUserResponse(model)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("User Found"),
		common.WithData(response),
	).Send(ctx)
}