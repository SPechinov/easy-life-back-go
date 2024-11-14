package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *restUserController) handlerUser(ctx echo.Context) error {
	println("hello")
	return ctx.NoContent(http.StatusNoContent)
}
