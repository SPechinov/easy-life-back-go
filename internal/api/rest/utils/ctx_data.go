package utils

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
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

func GetUserIDFromEchoCTX(echoCtx echo.Context) (string, error) {
	userID, ok := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return "", rest_error.ErrNotAuthorized
	}

	return userID, nil
}

func SetUserIDInEchoCTX(echoCtx echo.Context, userID string) {
	echoCtx.Set(globalConstants.CTXUserIDKey, userID)
}
