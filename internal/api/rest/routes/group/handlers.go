package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *restGroupController) handlerGroupAdd(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}
