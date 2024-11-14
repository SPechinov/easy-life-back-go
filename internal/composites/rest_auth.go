package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	userDatabase "go-clean/internal/adapters/database/user"
	restAuthStore "go-clean/internal/adapters/store/rest_auth"
	authRestHandler "go-clean/internal/api/rest/routes/auth"
	userService "go-clean/internal/services/user"
	restAuthUseCases "go-clean/internal/usecases/rest_auth"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewRestAuth(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres *postgres.Postgres) {
	udb := userDatabase.New(postgres)
	srvc := userService.New(udb)
	st := restAuthStore.New(redis)
	uc := restAuthUseCases.New(cfg, &st, srvc)
	authRestHandler.New(cfg, uc).Register(router)
}
