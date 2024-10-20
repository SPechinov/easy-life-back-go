package common

import (
	"easy-life-back-go/pkg/redis"
	"github.com/labstack/echo/v4"
)

type RoutesParams struct {
	Echo  *echo.Group
	Redis redis.Client
}
