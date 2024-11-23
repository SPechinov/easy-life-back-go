package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupNotesDatabase "go-clean/internal/adapters/database/group_notes"
	groupsDatabase "go-clean/internal/adapters/database/groups"
	groupNotesRestHandler "go-clean/internal/api/rest/routes/group_notes"
	groupNotesService "go-clean/internal/services/group_notes"
	groupsService "go-clean/internal/services/groups"
	groupNotesUseCases "go-clean/internal/usecases/group_notes"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroupNote(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupsDatabase.New(postgres)
	gs := groupsService.New(gdb)

	gudb := groupNotesDatabase.New(postgres)
	gns := groupNotesService.New(gudb)
	gnc := groupNotesUseCases.New(gns, gs)

	groupNotesRestHandler.New(cfg, gnc).Register(router)
}
