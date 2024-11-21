package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils"
	"go-clean/pkg/logger"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		requestID := uuid.New().String()
		echoCtx.Response().Header().Set(constants.HeaderXRequestID, requestID)

		// Log request ID
		ctx, err := utils.GetCTXLoggerFromEchoCTX(echoCtx)
		if err != nil {
			return err
		}

		ctx = logger.WithRequestID(ctx, requestID)
		utils.SetCTXLoggerInEchoCTX(echoCtx, ctx)
		return next(echoCtx)
	}
}
