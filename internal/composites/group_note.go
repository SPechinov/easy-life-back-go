package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupNoteDatabase "go-clean/internal/adapters/database/group_note"
	groupsDatabase "go-clean/internal/adapters/database/groups"
	groupNoteRestHandler "go-clean/internal/api/rest/routes/group_note"
	groupNoteService "go-clean/internal/services/group_note"
	groupsService "go-clean/internal/services/groups"
	groupNoteUseCases "go-clean/internal/usecases/group_note"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroupNote(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupsDatabase.New(postgres)
	gs := groupsService.New(gdb)

	gudb := groupNoteDatabase.New(postgres)
	gns := groupNoteService.New(gudb)
	gnc := groupNoteUseCases.New(gns, gs)

	groupNoteRestHandler.New(cfg, gnc).Register(router)
}
