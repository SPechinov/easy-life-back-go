package group_note

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type GroupNote struct {
	groupNoteService groupNoteService
	groupService     groupService
}

func New(groupNoteService groupNoteService, groupService groupService) *GroupNote {
	return &GroupNote{
		groupNoteService: groupNoteService,
		groupService:     groupService,
	}
}

func (gn *GroupNote) GetList(ctx context.Context, entity *entities.NoteGetList) ([]entities.Note, error) {
	err := gn.isAvailableGroup(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNoteService.GetList(ctx, entity)
}

func (gn *GroupNote) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	err := gn.isAvailableGroup(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	if gn.groupNoteService.IsDeleted(ctx, entity.ID) {
		return nil, client_error.ErrNoteDeleted
	}

	return gn.groupNoteService.Get(ctx, entity)
}

func (gn *GroupNote) Add(ctx context.Context, entity *entities.NoteAdd) error {
	err := gn.isAvailableGroup(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}
	return gn.groupNoteService.Add(ctx, entity)
}

func (gn *GroupNote) Patch(ctx context.Context, entity *entities.NotePatch) error {
	err := gn.isAvailableGroup(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}

	if gn.groupNoteService.IsDeleted(ctx, entity.ID) {
		return client_error.ErrNoteDeleted
	}

	return gn.groupNoteService.Patch(ctx, entity)
}

func (gn *GroupNote) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	err := gn.isAvailableGroup(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}

	isCreator := gn.groupNoteService.IsCreator(ctx, entity.UserID, entity.GroupID)
	isAdmin := gn.groupService.IsGroupAdmin(ctx, entity.UserID, entity.GroupID)

	if gn.groupNoteService.IsDeleted(ctx, entity.ID) {
		return client_error.ErrNoteDeleted
	}

	if !isCreator && !isAdmin {
		return client_error.ErrUserNotCreatorNote
	}

	return gn.groupNoteService.Delete(ctx, entity)
}

func (gn *GroupNote) isAvailableGroup(ctx context.Context, userID, groupID string) error {
	if gn.groupService.IsDeletedGroup(ctx, groupID) {
		return client_error.ErrGroupDeleted
	}

	user, err := gn.groupService.GetGroupUser(ctx, userID, groupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	return nil
}
