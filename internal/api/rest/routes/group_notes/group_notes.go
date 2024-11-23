package group_notes

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/controllers"
	"go-clean/internal/api/rest/middlewares"
)

const (
	urlGroupNotesList = ""
	urlGroupAdd       = ""
	urlGroupNote      = "/:noteID"
	urlGroupPatch     = "/:noteID"
	urlGroupDelete    = "/:noteID"
)

type restGroupNotesController struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) rest.Handler {
	return &restGroupNotesController{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (controller *restGroupNotesController) Register(group *echo.Group) {
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

	router.PATCH(
		urlGroupPatch,
		controllers.NewControllerUserIDValidation(controller.Patch, validatePatchDTO).Register,
	)

	router.DELETE(
		urlGroupDelete,
		controllers.NewControllerUserID(controller.Delete).Register,
	)
}
