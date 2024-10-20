package server

import (
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/pkg/redis"
	"github.com/labstack/echo/v4"
	"log/slog"
)

func Start(port string) {
	e := echo.New()
	r := redis.NewClient("localhost:6379", "", 0)

	registerRoutes(&common.RoutesParams{
		Echo:  e.Group("/api"),
		Redis: r,
	})

	err := e.Start(":" + port)
	if err != nil {
		slog.Error("Server crashed...")
	}
}
