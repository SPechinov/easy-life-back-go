package server

import (
	envInt "easy-life-back-go/internal/env"
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/pkg/redis"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log/slog"
)

func Start() {
	// Load env
	env, err := envInt.Init("../../env.yaml")
	if err != nil {
		slog.Error(err.Error())
	}

	// Start redis
	r := redis.NewClient(
		fmt.Sprintf("%v:%v", env.Redis.Host, env.Redis.Port),
		env.Redis.Password,
		env.Redis.DB,
	)
	slog.Info("Redis started")

	// Get echo and disable echo logger
	e := echo.New()
	e.Logger.SetOutput(io.Discard)

	// Start routing
	registerRoutes(&common.RoutesParams{
		Echo:  e.Group("/api"),
		Redis: r,
	})

	// Start server
	slog.Info("Server started on port: " + env.Server.Port)
	err = e.Start(":" + env.Server.Port)
	if err != nil {
		slog.Error("Server crashed...")
	}
}
