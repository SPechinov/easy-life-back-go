package server

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

func Start(port string) {
	e := echo.New()

	registerRoutes(e.Group("/api"))

	err := e.Start(":" + port)
	if err != nil {
		slog.Error("Server crashed...")
	}
}
