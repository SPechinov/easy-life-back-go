package group

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/middlewares"
	"go-clean/internal/api/rest/utils"
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

	authRouter.POST(urlGroupAdd, utils.HandleWithValidate[AddDTO](validateAddDTO, controller.handlerAddGroup))
	authRouter.PATCH(urlGroupPatch, utils.HandleWithValidate[PatchDTO](validatePatchDTO, controller.handlerPatchGroup))
	authRouter.DELETE(urlGroupDelete, utils.Handle(controller.handlerDelete))
	authRouter.DELETE(urlGroupDeleteConfirm, utils.HandleWithValidate[DeleteDTO](validateDeleteDTO, controller.handlerDeleteConfirm))

	authRouter.GET(urlGroupsList, utils.Handle(controller.handlerGetGroupsList))
	authRouter.GET(urlGroup, utils.Handle(controller.handlerGetFullGroup))
	authRouter.GET(urlGroupInfo, utils.Handle(controller.handlerGetGroupInfo))
	authRouter.GET(urlGroupUsers, utils.Handle(controller.handlerGetGroupUsers))

	authRouter.POST(urlGroupInviteUser, utils.HandleWithValidate[InviteUserDTO](validateInviteUserDTO, controller.handlerInviteUserInGroup))
	authRouter.POST(urlGroupExcludeUser, utils.HandleWithValidate[ExcludeUserDTO](validateExcludeUserDTO, controller.handlerExcludeUserFromGroup))
}
