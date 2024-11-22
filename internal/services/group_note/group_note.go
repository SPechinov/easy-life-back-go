package group_note

import (
	"context"
	"go-clean/internal/entities"
)

type GroupNote struct {
	groupNoteDatabase groupNoteDatabase
}

func New(groupNoteDatabase groupNoteDatabase) *GroupNote {
	return &GroupNote{
		groupNoteDatabase: groupNoteDatabase,
	}
}

func (gn *GroupNote) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	return gn.groupNoteDatabase.GetList(ctx, entity)
}

func (gn *GroupNote) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	return gn.groupNoteDatabase.Get(ctx, entity)
}

func (gn *GroupNote) Add(ctx context.Context, entity *entities.NoteAdd) error {
	return gn.groupNoteDatabase.Add(ctx, entity)
}

func (gn *GroupNote) Patch(ctx context.Context, entity *entities.NotePatch) error {
	return gn.groupNoteDatabase.Patch(ctx, entity)
}

func (gn *GroupNote) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	return gn.groupNoteDatabase.Delete(ctx, entity)
}

func (gn *GroupNote) IsCreator(ctx context.Context, userID, noteID string) bool {
	return gn.groupNoteDatabase.IsCreator(ctx, userID, noteID)
}

func (gn *GroupNote) IsDeleted(ctx context.Context, noteID string) bool {
	return gn.groupNoteDatabase.IsDeleted(ctx, noteID)
}
