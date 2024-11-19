package utils

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils/rest_error"
)

type handler func(echoCtx echo.Context, ctx context.Context, userID string) error

func getLoggerCTXAndUserID(echoCtx echo.Context) (context.Context, string, error) {
	ctx, err := GetCTXLoggerFromEchoCTX(echoCtx)
	if err != nil {
		return nil, "", err
	}

	userID, err := GetUserIDFromEchoCTX(echoCtx)
	if err != nil {
		return nil, "", err
	}

	return ctx, userID, nil
}

func Handle(handler handler) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx, userID, err := getLoggerCTXAndUserID(echoCtx)
		if err != nil {
			return err
		}
		return handler(echoCtx, ctx, userID)
	}
}

type handlerWithValidate[T any] func(echoCtx echo.Context, ctx context.Context, dto *T, userID string) error

func HandleWithValidate[T any](validator func(*T) error, handler handlerWithValidate[T]) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx, userID, err := getLoggerCTXAndUserID(echoCtx)
		if err != nil {
			return err
		}

		dto := new(T)
		err = echoCtx.Bind(dto)
		if err != nil {
			return rest_error.ErrInvalidBodyData
		}

		err = validator(dto)
		if err != nil {
			return err
		}

		return handler(echoCtx, ctx, dto, userID)
	}
}
