package group_notes

import (
	"context"
	"go-clean/internal/entities"
)

type groupNotesService interface {
	GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error)
	Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error)
	Add(ctx context.Context, entity *entities.NoteAdd) (*entities.Note, error)
	Patch(ctx context.Context, entity *entities.NotePatch) error
	Delete(ctx context.Context, entity *entities.NoteDelete) error
	IsCreator(ctx context.Context, userID, noteID string) bool
}

type groupsService interface {
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
	IsGroupUser(ctx context.Context, userID, groupID string) error
}
