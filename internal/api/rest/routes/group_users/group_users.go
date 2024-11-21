package group_users

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/controllers"
	"go-clean/internal/api/rest/middlewares"
)

const (
	urlGroupUsersList   = "/:groupID/users"
	urlGroupUserInvite  = "/:groupID/users/invite"
	urlGroupUserExclude = "/:groupID/users/exclude"
)

type restGroupUsersController struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restGroupUsersController{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (controller *restGroupUsersController) Register(group *echo.Group) {
	router := group.Group("/groups")
	router.Use(middlewares.AuthMiddleware(controller.cfg))

	router.GET(
		urlGroupUsersList,
		controllers.NewControllerUserID(controller.handlerGetList).Register,
	)
	router.POST(
		urlGroupUserInvite,
		controllers.NewControllerUserIDValidation(controller.handlerInviteUserInGroup, validateInviteUserDTO).Register,
	)
	router.POST(
		urlGroupUserExclude,
		controllers.NewControllerUserIDValidation(controller.handleExcludeUserInGroup, validateExcludeUserDTO).Register,
	)
}
