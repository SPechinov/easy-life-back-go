package group_notes

import (
	"context"
	"go-clean/internal/entities"
)

type GroupNotes struct {
	groupNotesDatabase groupNotesDatabase
}

func New(groupNotesDatabase groupNotesDatabase) *GroupNotes {
	return &GroupNotes{
		groupNotesDatabase: groupNotesDatabase,
	}
}

func (gn *GroupNotes) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	return gn.groupNotesDatabase.GetList(ctx, entity)
}

func (gn *GroupNotes) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	return gn.groupNotesDatabase.Get(ctx, entity)
}

func (gn *GroupNotes) Add(ctx context.Context, entity *entities.NoteAdd) (*entities.Note, error) {
	return gn.groupNotesDatabase.Add(ctx, entity)
}

func (gn *GroupNotes) Patch(ctx context.Context, entity *entities.NotePatch) error {
	return gn.groupNotesDatabase.Patch(ctx, entity)
}

func (gn *GroupNotes) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	return gn.groupNotesDatabase.Delete(ctx, entity)
}

func (gn *GroupNotes) IsCreator(ctx context.Context, userID, noteID string) bool {
	return gn.groupNotesDatabase.IsCreator(ctx, userID, noteID)
}

func (gn *GroupNotes) IsDeleted(ctx context.Context, noteID string) bool {
	return gn.groupNotesDatabase.IsDeleted(ctx, noteID)
}
