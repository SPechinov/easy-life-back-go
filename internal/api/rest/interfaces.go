package rest

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Register(router *echo.Group)
}
