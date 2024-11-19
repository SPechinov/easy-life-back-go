package utils

import (
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils/rest_error"
)

func BindAndValidate[T any](echoCtx echo.Context, validator func(*T) error) (*T, error) {
	dto := new(T)
	if err := echoCtx.Bind(dto); err != nil {
		return nil, rest_error.ErrInvalidBodyData
	}
	if err := validator(dto); err != nil {
		return nil, err
	}

	return dto, nil
}
