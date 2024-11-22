package composites

import (
	"github.com/labstack/echo/v4"
	"go-clean/config"
	groupDatabase "go-clean/internal/adapters/database/group"
	groupNoteDatabase "go-clean/internal/adapters/database/group_note"
	groupNoteRestHandler "go-clean/internal/api/rest/routes/group_note"
	groupService "go-clean/internal/services/group"
	groupNoteService "go-clean/internal/services/group_note"
	groupNoteUseCases "go-clean/internal/usecases/group_note"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
)

func NewGroupNote(cfg *config.Config, router *echo.Group, redis *redis.Redis, postgres postgres.Client) {
	gdb := groupDatabase.New(postgres)
	gs := groupService.New(gdb)

	gudb := groupNoteDatabase.New(postgres)
	gns := groupNoteService.New(gudb)
	gnc := groupNoteUseCases.New(gns, gs)

	groupNoteRestHandler.New(cfg, gnc).Register(router)
}
