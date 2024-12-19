package presentation

import (
	"api/first-go/apps/auth/application"
	"api/first-go/common"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    auth application.AuthUseCase
    user application.UserUseCase
}

func NewAuthHandler(au application.AuthUseCase, uu application.UserUseCase) AuthHandler {
    return AuthHandler{
		auth: au,
		user: uu,
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

	_, err := h.user.FindByEmail(ctx.UserContext(), req.Email)
	if err == nil {
		return common.NewResponse(
			common.WithMessage("Email has already been registered"),
			common.WithError(common.ErrorEmailAlreadyUsed),
			common.WithHttpCode(http.StatusConflict),
		).Send(ctx)
	}

    if err := h.auth.Register(ctx.UserContext(), req); err != nil {
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

	model, err := h.user.FindByEmail(ctx.UserContext(), req.Email)
	if  err != nil { 
		return common.NewResponse(
			common.WithMessage("Email not registered"),
			common.WithError(common.ErrorNotFound),
		).Send(ctx)
	}

	token, err := h.auth.Login(ctx.UserContext(), model, req.Password)
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

	cookie := new(fiber.Cookie)
  	cookie.Name = "refresh_token"
  	cookie.Value = token.RefreshToken
  	cookie.Expires = time.Now().Add(7 *24 * time.Hour)
	cookie.Secure = true
	cookie.HTTPOnly = true

	ctx.Cookie(cookie)

	return common.NewResponse(
		common.WithHttpCode(http.StatusOK),
		common.WithMessage("Login successfully"),
		common.WithData(map[string]interface{}{
			"access_token": token.AccessToken,
		}),
	).Send(ctx)
}

func (h AuthHandler) Refresh(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")
	role := ctx.Locals("role")

	req := application.TokenPayload{
		Id: id.(string),
		Role: role.(string),
	}

	token, err := h.auth.Refresh(ctx.UserContext(), req)
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