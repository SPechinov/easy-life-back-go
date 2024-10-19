package auth

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group, controller *Controller) {
	g := e.Group("/auth")

	g.POST("/sign-in", controller.SignIn)
	g.POST("/registration", controller.Registration)
	g.POST("/registration-success", controller.RegistrationSuccess)
}
