package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupDatabase "go-clean/internal/adapters/database/group"
	groupRestHandler "go-clean/internal/api/rest/routes/group"
	"go-clean/internal/services/codes"
	groupService "go-clean/internal/services/group"
	groupUseCases "go-clean/internal/usecases/group"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroup(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupDatabase.New(postgres)
	gs := groupService.New(gdb)
	c := codes.New(redis)
	guc := groupUseCases.New(cfg, gs, c)

	groupRestHandler.New(cfg, &guc).Register(router)
}
