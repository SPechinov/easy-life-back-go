package auth

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/middlewares"
)

const (
	urlLogin                 = "/login"
	urlRegistration          = "/registration"
	urlRegistrationConfirm   = "/registration-confirm"
	urlForgotPassword        = "/forgot-password"
	urlForgotPasswordConfirm = "/forgot-password-confirm"
	urlUpdateJWT             = "/update-jwt"
	urlLogout                = "/logout"
	urlLogoutAll             = "/logout-all"
)

type restAuthController struct {
	useCases useCases
	cfg      *config.Config
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restAuthController{cfg: cfg, useCases: useCases}
}

func (controller *restAuthController) Register(router *echo.Group) {
	authRouter := router.Group("/auth")

	authRouter.POST(urlLogin, controller.handlerLogin)
	authRouter.POST(urlRegistration, controller.handlerRegistration)
	authRouter.POST(urlRegistrationConfirm, controller.handlerRegistrationConfirm)
	authRouter.POST(urlForgotPassword, controller.handlerForgotPassword)
	authRouter.POST(urlForgotPasswordConfirm, controller.handlerForgotPasswordConfirm)
	authRouter.POST(urlUpdateJWT, controller.handlerUpdateJWT)

	authRouterWithAuth := authRouter.Group("")
	authRouterWithAuth.Use(middlewares.AuthMiddleware(controller.cfg))

	authRouterWithAuth.POST(urlLogout, controller.handlerLogout)
	authRouterWithAuth.POST(urlLogoutAll, controller.handlerLogoutAll)
}
