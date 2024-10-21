package server

import (
	envInt "easy-life-back-go/internal/env"
	"easy-life-back-go/internal/pkg/store_codes"
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/pkg/store"
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

	// Start store
	s := store.NewClient(
		fmt.Sprintf("%v:%v", env.Redis.Host, env.Redis.Port),
		env.Redis.Password,
		env.Redis.DB,
	)
	sCodes := store_codes.NewStoreCodes(s)
	slog.Info("Redis started")

	// Get echo and disable echo logger
	e := echo.New()
	e.Logger.SetOutput(io.Discard)

	// Start routing
	registerRoutes(&common.RoutesParams{
		Echo:       e.Group("/api"),
		Store:      s,
		StoreCodes: sCodes,
	})

	// Start server
	slog.Info("Server started on port: " + env.Server.Port)
	err = e.Start(":" + env.Server.Port)
	if err != nil {
		slog.Error("Server crashed...")
	}
}
