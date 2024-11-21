package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils"
	"go-clean/pkg/logger"
)

func StartLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		requestID := echoCtx.Response().Header().Get(constants.HeaderXRequestID)

		ctx := context.Background()
		ctx = logger.WithRequestID(ctx, requestID)
		ctx = logger.WithURL(ctx, echoCtx.Request().RequestURI)
		utils.SetCTXLoggerInEchoCTX(echoCtx, ctx)

		logger.Debug(ctx, "Start")
		err := next(echoCtx)
		logger.Debug(ctx, "Finish")

		return err
	}
}
