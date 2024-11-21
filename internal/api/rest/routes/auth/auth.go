package auth

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/controllers"
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

	authRouter.POST(
		urlLogin,
		controllers.NewControllerValidation(controller.handlerLogin, validateLoginDTO).Register,
	)
	authRouter.POST(
		urlRegistration,
		controllers.NewControllerValidation(controller.handlerRegistration, validateRegistrationDTO).Register,
	)
	authRouter.POST(
		urlRegistrationConfirm,
		controllers.NewControllerValidation(controller.handlerRegistrationConfirm, validateRegistrationConfirmDTO).Register,
	)
	authRouter.POST(
		urlForgotPassword,
		controllers.NewControllerValidation(controller.handlerForgotPassword, validateForgotPasswordDTO).Register,
	)
	authRouter.POST(
		urlForgotPasswordConfirm,
		controllers.NewControllerValidation(controller.handlerForgotPasswordConfirm, validateForgotPasswordConfirmDTO).Register,
	)
	authRouter.POST(
		urlUpdateJWT,
		controllers.NewController(controller.handlerUpdateJWT).Register,
	)

	authRouterWithAuth := authRouter.Group("")
	authRouterWithAuth.Use(middlewares.AuthMiddleware(controller.cfg))

	authRouterWithAuth.POST(
		urlLogout, controllers.NewControllerUserID(controller.handlerLogout).Register,
	)
	authRouterWithAuth.POST(
		urlLogoutAll, controllers.NewControllerUserID(controller.handlerLogoutAll).Register,
	)
}
