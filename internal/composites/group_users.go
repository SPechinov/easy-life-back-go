package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupDatabase "go-clean/internal/adapters/database/group"
	groupUsersDatabase "go-clean/internal/adapters/database/group_users"
	groupUsersRestHandler "go-clean/internal/api/rest/routes/group_users"
	groupService "go-clean/internal/services/group"
	groupUsersService "go-clean/internal/services/group_users"
	groupUsersUseCases "go-clean/internal/usecases/group_users"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroupUsers(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupDatabase.New(postgres)
	gs := groupService.New(gdb)

	gudb := groupUsersDatabase.New(postgres)
	gus := groupUsersService.New(gudb)
	guc := groupUsersUseCases.New(cfg, gus, gs)

	groupUsersRestHandler.New(cfg, guc).Register(router)
}
