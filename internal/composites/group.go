package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupDatabase "go-clean/internal/adapters/database/group"
	userDatabase "go-clean/internal/adapters/database/user"
	groupStore "go-clean/internal/adapters/store/group"
	groupRestHandler "go-clean/internal/api/rest/routes/group"
	groupService "go-clean/internal/services/group"
	userService "go-clean/internal/services/user"
	groupUseCases "go-clean/internal/usecases/group"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroup(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	udb := userDatabase.New(postgres)
	us := userService.New(udb)

	gdb := groupDatabase.New(postgres)
	gs := groupService.New(gdb, us)
	gStore := groupStore.New(redis)
	guc := groupUseCases.New(cfg, gs, gStore)

	groupRestHandler.New(cfg, &guc).Register(router)
}
