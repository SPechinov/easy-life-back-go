package user

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/middlewares"
)

type restUserController struct {
	cfg *config.Config
}

func New(cfg *config.Config) rest.Handler {
	return &restUserController{cfg: cfg}
}

func (c *restUserController) Register(router *echo.Group) {
	authRouter := router.Group("/user")
	authRouter.Use(middlewares.AuthMiddleware(c.cfg))

	authRouter.GET("", c.handlerUser)
}
