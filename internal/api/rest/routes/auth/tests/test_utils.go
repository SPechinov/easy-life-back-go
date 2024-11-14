package tests

import (
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/middlewares"
)

func StartTestServer() (*echo.Echo, *echo.Group) {
	restServer := echo.New()
	restServer.Use(middlewares.RequestIDMiddleware)
	restServer.Use(middlewares.StartLogging)
	restServer.Use(middlewares.ResponseMiddleware)
	return restServer, restServer.Group("")
}
