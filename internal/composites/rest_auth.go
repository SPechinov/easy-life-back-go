package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	usersDatabase "go-clean/internal/adapters/database/users"
	restAuthStore "go-clean/internal/adapters/store/rest_auth"
	authRestHandler "go-clean/internal/api/rest/routes/auth"
	"go-clean/internal/services/codes"
	usersService "go-clean/internal/services/users"
	restAuthUseCases "go-clean/internal/usecases/rest_auth"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewRestAuth(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	udb := usersDatabase.New(postgres)
	srvc := usersService.New(udb)
	st := restAuthStore.New(redis)
	c := codes.New(redis)
	uc := restAuthUseCases.New(cfg, &st, srvc, c)
	authRestHandler.New(cfg, uc).Register(router)
}
