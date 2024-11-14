package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	userRestHandler "go-clean/internal/api/rest/routes/user"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewRestUser(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres *postgres.Postgres) {
	userRestHandler.New(cfg).Register(router)
}
