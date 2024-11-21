package utils

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/pkg/logger"
)

func SetCTXLoggerInEchoCTX(echoCtx echo.Context, ctx context.Context) {
	echoCtx.Set(constants.CTXLoggerInCTX, ctx)
}

func GetCTXLoggerFromEchoCTX(echoCtx echo.Context) (context.Context, error) {
	ctx, ok := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No echo context")
		return nil, rest_error.ErrSomethingHappen
	}

	return ctx, nil
}
