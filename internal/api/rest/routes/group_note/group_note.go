package group_note

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/controllers"
	"go-clean/internal/api/rest/middlewares"
)

const (
	urlGroupNotesList = "/notes"
	urlGroupAdd       = "/notes"
	urlGroupNote      = "/notes/:noteID"
	urlGroupPatch     = "/notes/:noteID"
	urlGroupDelete    = "/notes/:noteID"
)

type restGroupNoteController struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restGroupNoteController{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (controller *restGroupNoteController) Register(group *echo.Group) {
	router := group.Group("/groups/:groupID/notes")
	router.Use(middlewares.AuthMiddleware(controller.cfg))

	router.GET(
		urlGroupNotesList,
		controllers.NewControllerUserID(controller.GetList).Register,
	)

	router.GET(
		urlGroupNote,
		controllers.NewControllerUserID(controller.Get).Register,
	)

	router.POST(
		urlGroupAdd,
		controllers.NewControllerUserIDValidation(controller.Add, validateAddDTO).Register,
	)

	router.POST(
		urlGroupPatch,
		controllers.NewControllerUserIDValidation(controller.Patch, validatePatchDTO).Register,
	)

	router.POST(
		urlGroupDelete,
		controllers.NewControllerUserID(controller.Delete).Register,
	)
}
