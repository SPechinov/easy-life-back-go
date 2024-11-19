package group

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/middlewares"
)

type restGroupController struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restGroupController{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (controller *restGroupController) Register(router *echo.Group) {
	authRouter := router.Group("/groups")
	authRouter.Use(middlewares.AuthMiddleware(controller.cfg))

	authRouter.GET("", controller.handlerGetGroupsList)
	authRouter.POST("", controller.handlerAddGroup)
	authRouter.GET("/:groupID", controller.handlerGetGroup)
	authRouter.GET("/:groupID/info", controller.handlerGetGroupInfo)
	authRouter.PATCH("/:groupID", controller.handlerPatchGroup)
	authRouter.GET("/:groupID/users", controller.handlerGetGroupUsers)
	authRouter.POST("/:groupID/invite-user", controller.handlerInviteUserInGroup)
	authRouter.POST("/:groupID/exclude-user", controller.handlerExcludeUserFromGroup)
}
