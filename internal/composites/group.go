package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupDatabase "go-clean/internal/adapters/database/group"
	userDatabase "go-clean/internal/adapters/database/user"
	groupRestHandler "go-clean/internal/api/rest/routes/group"
	groupService "go-clean/internal/services/group"
	userService "go-clean/internal/services/user"
	groupUseCases "go-clean/internal/usecases/group"
	"go-clean/pkg/postgres"
)

func NewGroup(cfg *config.Config, router *echo.Group, postgres postgres.Client) {
	udb := userDatabase.New(postgres)
	us := userService.New(udb)

	gdb := groupDatabase.New(postgres)
	gs := groupService.New(gdb, us)
	guc := groupUseCases.New(cfg, gs)

	groupRestHandler.New(cfg, &guc).Register(router)
}
