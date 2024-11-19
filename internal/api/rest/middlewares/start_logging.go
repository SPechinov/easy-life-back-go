package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/pkg/logger"
)

func StartLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		requestID := echoCtx.Response().Header().Get(constants.HeaderXRequestID)

		ctx := context.Background()
		ctx = logger.WithRequestID(ctx, requestID)
		ctx = logger.WithURL(ctx, echoCtx.Request().RequestURI)

		// Group ID
		if groupID := echoCtx.Param("groupID"); groupID != "" {
			ctx = logger.WithGroupID(ctx, groupID)
		}
		echoCtx.Set(constants.CTXLoggerInCTX, ctx)

		logger.Debug(ctx, "Start")
		err := next(echoCtx)

		if err != nil {
			logger.Debug(ctx, "Err")
		} else {
			logger.Debug(ctx, "Finish")
		}
		return err
	}
}
