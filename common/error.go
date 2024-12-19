package common

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimimum 6 character")
	ErrUsernameRequired      = errors.New("username is required")
	ErrUsernameInvalidLength = errors.New("username must have minimimum 4 character")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral         = NewError("Internal Server Error", "500", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("Bad Request", "400", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "404", http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), "401", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "403", http.StatusForbidden)
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "400", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "400", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "400", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "400", http.StatusBadRequest)
	ErrorUsernameRequired      = NewError(ErrUsernameRequired.Error(), "400", http.StatusBadRequest)
	ErrorUsernameInvalidLength = NewError(ErrUsernameInvalidLength.Error(), "400", http.StatusBadRequest)

	ErrorAuthIsNotExists  = NewError(ErrAuthIsNotExists.Error(), "404", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "409", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "401", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrUsernameRequired.Error():         ErrorUsernameRequired,
		ErrUsernameInvalidLength.Error():          ErrorUsernameInvalidLength,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error():       ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrUnauthorized.Error():          ErrorUnauthorized,
		ErrForbiddenAccess.Error():       ErrorForbiddenAccess,
	}
)