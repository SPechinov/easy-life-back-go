package group_notes

import (
	"context"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type GroupNotes struct {
	groupNotesService groupNotesService
	groupsService     groupsService
}

func New(groupNoteService groupNotesService, groupService groupsService) *GroupNotes {
	return &GroupNotes{
		groupNotesService: groupNoteService,
		groupsService:     groupService,
	}
}

func (gn *GroupNotes) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	err := gn.groupsService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNotesService.GetList(ctx, entity)
}

func (gn *GroupNotes) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	err := gn.groupsService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNotesService.Get(ctx, entity)
}

func (gn *GroupNotes) Add(ctx context.Context, entity *entities.NoteAdd) (*entities.Note, error) {
	err := gn.groupsService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNotesService.Add(ctx, entity)
}

func (gn *GroupNotes) Patch(ctx context.Context, entity *entities.NotePatch) error {
	err := gn.groupsService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}

	return gn.groupNotesService.Patch(ctx, entity)
}

func (gn *GroupNotes) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	user, err := gn.groupsService.GetGroupUser(ctx, entity.UserID, entity.GroupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	isCreator := gn.groupNotesService.IsCreator(ctx, entity.UserID, entity.GroupID)
	isAdmin := user.Permission == constants.DefaultAdminPermission

	if !isCreator && !isAdmin {
		return client_error.ErrUserNotCreatorNote
	}

	return gn.groupNotesService.Delete(ctx, entity)
}
