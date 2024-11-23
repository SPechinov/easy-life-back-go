package group_notes

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/postgres"
)

type GroupNotes struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *GroupNotes {
	return &GroupNotes{
		postgres: postgres,
	}
}

func (g *GroupNotes) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	return nil, nil
}
func (g *GroupNotes) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	return nil, nil
}

func (g *GroupNotes) Add(ctx context.Context, entity *entities.NoteAdd) error {
	return nil
}

func (g *GroupNotes) Patch(ctx context.Context, entity *entities.NotePatch) error {
	return nil
}

func (g *GroupNotes) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	return nil
}

func (g *GroupNotes) IsCreator(ctx context.Context, userID, noteID string) bool {
	return false
}

func (g *GroupNotes) IsDeleted(ctx context.Context, noteID string) bool {
	return false
}
