package group_note

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/postgres"
)

type GroupNote struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *GroupNote {
	return &GroupNote{
		postgres: postgres,
	}
}

func (g *GroupNote) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	return nil, nil
}
func (g *GroupNote) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	return nil, nil
}

func (g *GroupNote) Add(ctx context.Context, entity *entities.NoteAdd) error {
	return nil
}

func (g *GroupNote) Patch(ctx context.Context, entity *entities.NotePatch) error {
	return nil
}

func (g *GroupNote) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	return nil
}

func (g *GroupNote) IsCreator(ctx context.Context, userID, noteID string) bool {
	return false
}

func (g *GroupNote) IsDeleted(ctx context.Context, noteID string) bool {
	return false
}
