package server

import (
	"easy-life-back-go/pkg/redis"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func Start(port string) {
	e := echo.New()
	r := redis.NewClient("localhost:6379", "", 0)

	registerRoutes(&RoutesParams{
		echo:  e.Group("/api"),
		redis: &r,
	})

	err := e.Start(":" + port)
	if err != nil {
		slog.Error("Server crashed...")
	}
}
