package groups

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
)

type restGroupsController struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restGroupsController{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (controller *restGroupsController) Register(group *echo.Group) {
	router := group.Group("/groups")
	router.Use(middlewares.AuthMiddleware(controller.cfg))

	router.POST(
		urlGroupAdd,
		controllers.NewControllerUserIDValidation(controller.handlerAddGroup, validateAddDTO).Register,
	)
	router.PATCH(
		urlGroupPatch,
		controllers.NewControllerUserIDValidation(controller.handlerPatchGroup, validatePatchDTO).Register,
	)
	router.POST(
		urlGroupDelete,
		controllers.NewControllerUserID(controller.handlerDelete).Register,
	)
	router.POST(
		urlGroupDeleteConfirm,
		controllers.NewControllerUserIDValidation(controller.handlerDeleteConfirm, validateDeleteConfirmDTO).Register,
	)

	router.GET(
		urlGroupsList,
		controllers.NewControllerUserID(controller.handlerGetGroupsList).Register,
	)
	router.GET(
		urlGroup,
		controllers.NewControllerUserID(controller.handlerGetGroup).Register,
	)
}
