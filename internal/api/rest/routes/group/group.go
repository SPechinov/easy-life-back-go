package group

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/middlewares"
	"go-clean/internal/api/rest/utils"
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

	authRouter.GET("", utils.Handle(controller.handlerGetGroupsList))
	authRouter.POST("", utils.HandleWithValidate[AddDTO](validateAddDTO, controller.handlerAddGroup))
	authRouter.GET("/:groupID", utils.Handle(controller.handlerGetGroup))
	authRouter.GET("/:groupID/info", utils.Handle(controller.handlerGetGroupInfo))
	authRouter.GET("/:groupID/users", utils.Handle(controller.handlerGetGroupUsers))
	authRouter.PATCH("/:groupID", utils.HandleWithValidate[PatchDTO](validatePatchDTO, controller.handlerPatchGroup))
	authRouter.POST("/:groupID/invite-user", utils.HandleWithValidate[InviteUserDTO](validateInviteUserDTO, controller.handlerInviteUserInGroup))
	authRouter.POST("/:groupID/exclude-user", utils.HandleWithValidate[ExcludeUserDTO](validateExcludeUserDTO, controller.handlerExcludeUserFromGroup))
}
