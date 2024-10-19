package server

import (
	"easy-life-back-go/internal/server/routes/auth"
	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Group) {
	registerAuthRoutes(e)
}

func registerAuthRoutes(e *echo.Group) {
	service := auth.NewService()
	controller := auth.NewController(service)
	auth.RegisterRoutes(e, controller)
}
