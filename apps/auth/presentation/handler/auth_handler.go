package presentation

import (
	"api/first-go/apps/auth/application"
	"api/first-go/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    usecase application.AuthUseCase
}

func NewAuthHandler(uc application.AuthUseCase) AuthHandler {
    return AuthHandler{
		usecase: uc,
	}
}

func (h AuthHandler) Register(ctx *fiber.Ctx) error {
    var req = application.RegisterRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
			common.WithHttpCode(http.StatusBadRequest),
			common.WithMessage("register fail"),
		).Send(ctx)
	}

    if err := h.usecase.Register(ctx.UserContext(), req); err != nil {
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
		common.WithMessage("Register successfully"),
	).Send(ctx)
}

func (h AuthHandler) Login(ctx *fiber.Ctx) error {
    var req = application.LoginRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := common.ErrorBadRequest
		return common.NewResponse(
			common.WithMessage(err.Error()),
			common.WithError(myErr),
			common.WithHttpCode(http.StatusBadRequest),
			common.WithMessage("Login fail"),
		).Send(ctx)
	}

	token, err := h.usecase.Login(ctx.UserContext(), req)
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
		common.WithMessage("Login successfully"),
		common.WithData(map[string]interface{}{
			"access_token": token,
		}),
	).Send(ctx)
}