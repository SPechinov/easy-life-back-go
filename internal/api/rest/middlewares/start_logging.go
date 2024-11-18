package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/pkg/logger"
)

func StartLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := c.Response().Header().Get(constants.HeaderXRequestID)

		ctx := context.Background()
		ctx = logger.WithRequestID(ctx, requestID)
		ctx = logger.WithURL(ctx, c.Request().RequestURI)

		// Group ID
		if groupID := c.Param("groupID"); groupID != "" {
			ctx = logger.WithGroupID(ctx, groupID)
		}

		c.Set(constants.CTXLoggerInCTX, ctx)
		return next(c)
	}
}
