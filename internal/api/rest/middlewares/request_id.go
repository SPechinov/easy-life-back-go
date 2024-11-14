package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := uuid.New().String()
		c.Response().Header().Set(constants.HeaderXRequestID, requestID)
		return next(c)
	}
}
