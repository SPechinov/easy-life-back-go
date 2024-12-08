package middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/pkg/client_error"
	"net/http"
)

var appErrorMapping = map[string]*rest_error.RestError{
	client_error.ErrNotAuthorized.Error():      rest_error.ErrNotAuthorized,
	client_error.ErrIncorrectPassword.Error():  rest_error.ErrIncorrectPassword,
	client_error.ErrUserNotFound.Error():       rest_error.ErrUserNotFound,
	client_error.ErrCodeIsNotInRedis.Error():   rest_error.ErrCodeDidNotSent,
	client_error.ErrUserExists.Error():         rest_error.ErrUserExists,
	client_error.ErrCodeMaxAttempts.Error():    rest_error.ErrCodeMaxAttempts,
	client_error.ErrCodesIsNotEqual.Error():    rest_error.ErrCodesIsNotEqual,
	client_error.ErrUserDeleted.Error():        rest_error.ErrUserDeleted,
	client_error.ErrUserNotAdminGroup.Error():  rest_error.ErrUserNotAdminGroup,
	client_error.ErrUserNotInGroup.Error():     rest_error.ErrUserNotInGroup,
	client_error.ErrUserInvited.Error():        rest_error.ErrUserInvited,
	client_error.ErrUserAdminGroup.Error():     rest_error.ErrUserAdminGroup,
	client_error.ErrGroupDeleted.Error():       rest_error.ErrGroupNotExist,
	client_error.ErrGroupNotExists.Error():     rest_error.ErrGroupNotExist,
	client_error.ErrNoteDeleted.Error():        rest_error.ErrNoteDeleted,
	client_error.ErrUserNotCreatorNote.Error(): rest_error.ErrUserNotCreatorNote,
}

func ResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		err := next(echoCtx)

		if err == nil {
			return nil
		}

		var restError *rest_error.RestError
		var validationError *rest_error.ValidationError
		var clientError *client_error.ClientError

		switch {
		// Rest error
		case errors.As(err, &restError):
			return echoCtx.JSON(restError.HttpCode, rest.NewResponseBad(restError.Code))

		// Validation error
		case errors.As(err, &validationError):
			return echoCtx.JSON(http.StatusBadRequest, rest.NewResponseBadValidation(validationError.Message))

		// Client error
		case errors.As(err, &clientError):
			value, exist := appErrorMapping[clientError.Code()]
			if !exist {
				return echoCtx.JSON(rest_error.ErrSomethingHappen.HttpCode, rest.NewResponseBad(rest_error.ErrSomethingHappen.Code))
			}
			return echoCtx.JSON(value.HttpCode, rest.NewResponseBad(value.Code))

		// Ops. PANIC
		default:
			return echoCtx.JSON(rest_error.ErrSomethingHappen.HttpCode, rest.NewResponseBad(rest_error.ErrSomethingHappen.Code))
		}
	}
}
