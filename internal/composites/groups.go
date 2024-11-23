package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupsDatabase "go-clean/internal/adapters/database/groups"
	groupsRestHandler "go-clean/internal/api/rest/routes/groups"
	"go-clean/internal/services/codes"
	groupsService "go-clean/internal/services/groups"
	groupsUseCases "go-clean/internal/usecases/groups"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroups(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupsDatabase.New(postgres)
	gs := groupsService.New(gdb)
	c := codes.New(redis)
	guc := groupsUseCases.New(cfg, gs, c)

	groupsRestHandler.New(cfg, &guc).Register(router)
}
