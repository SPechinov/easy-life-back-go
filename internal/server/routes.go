package server

import (
	"easy-life-back-go/internal/server/routes/auth"
	"easy-life-back-go/pkg/redis"
	"github.com/labstack/echo/v4"
)

type RoutesParams struct {
	echo  *echo.Group
	redis *redis.Client
}

func registerRoutes(params *RoutesParams) {
	auth.RegisterRoutes(params)
}
