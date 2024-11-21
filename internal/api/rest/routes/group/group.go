package group

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/controllers"
	"go-clean/internal/api/rest/middlewares"
)

const (
	urlGroupAdd           = ""
	urlGroupPatch         = "/:groupID"
	urlGroupDelete        = "/:groupID/delete"
	urlGroupDeleteConfirm = "/:groupID/delete-confirm"
	urlGroupsList         = ""
	urlGroup              = "/:groupID"
	urlGroupInfo          = "/:groupID/info"
	urlGroupUsers         = "/:groupID/users"
	urlGroupInviteUser    = "/:groupID/invite-user"
	urlGroupExcludeUser   = "/:groupID/exclude-user"
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

	authRouter.POST(
		urlGroupAdd,
		controllers.NewControllerUserIDValidation(controller.handlerAddGroup, validateAddDTO).Register,
	)
	authRouter.PATCH(
		urlGroupPatch,
		controllers.NewControllerUserIDValidation(controller.handlerPatchGroup, validatePatchDTO).Register,
	)
	authRouter.DELETE(
		urlGroupDelete,
		controllers.NewControllerUserID(controller.handlerDelete).Register,
	)
	authRouter.DELETE(
		urlGroupDeleteConfirm,
		controllers.NewControllerUserIDValidation(controller.handlerDeleteConfirm, validateDeleteDTO).Register,
	)

	authRouter.GET(
		urlGroupsList,
		controllers.NewControllerUserID(controller.handlerGetGroupsList).Register,
	)
	authRouter.GET(
		urlGroup,
		controllers.NewControllerUserID(controller.handlerGetFullGroup).Register,
	)
	authRouter.GET(
		urlGroupInfo,
		controllers.NewControllerUserID(controller.handlerGetGroupInfo).Register,
	)
	authRouter.GET(
		urlGroupUsers,
		controllers.NewControllerUserID(controller.handlerGetGroupUsers).Register,
	)

	authRouter.POST(
		urlGroupInviteUser,
		controllers.NewControllerUserIDValidation(controller.handlerInviteUserInGroup, validateInviteUserDTO).Register,
	)
	authRouter.POST(
		urlGroupExcludeUser,
		controllers.NewControllerUserIDValidation(controller.handlerExcludeUserFromGroup, validateExcludeUserDTO).Register,
	)
}
