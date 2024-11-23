package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupUsersDatabase "go-clean/internal/adapters/database/group_users"
	groupsDatabase "go-clean/internal/adapters/database/groups"
	groupUsersRestHandler "go-clean/internal/api/rest/routes/group_users"
	groupUsersService "go-clean/internal/services/group_users"
	groupsService "go-clean/internal/services/groups"
	groupUsersUseCases "go-clean/internal/usecases/group_users"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroupUsers(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupsDatabase.New(postgres)
	gs := groupsService.New(gdb)

	gudb := groupUsersDatabase.New(postgres)
	gus := groupUsersService.New(gudb)
	guc := groupUsersUseCases.New(cfg, gus, gs)

	groupUsersRestHandler.New(cfg, guc).Register(router)
}
