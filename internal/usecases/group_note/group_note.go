package group_note

import (
	"context"
	"go-clean/internal/constants"
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
	err := gn.groupService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNoteService.GetList(ctx, entity)
}

func (gn *GroupNote) Get(ctx context.Context, entity *entities.NoteGet) (*entities.Note, error) {
	err := gn.groupService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return nil, err
	}

	return gn.groupNoteService.Get(ctx, entity)
}

func (gn *GroupNote) Add(ctx context.Context, entity *entities.NoteAdd) error {
	err := gn.groupService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}

	return gn.groupNoteService.Add(ctx, entity)
}

func (gn *GroupNote) Patch(ctx context.Context, entity *entities.NotePatch) error {
	err := gn.groupService.IsGroupUser(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}

	return gn.groupNoteService.Patch(ctx, entity)
}

func (gn *GroupNote) Delete(ctx context.Context, entity *entities.NoteDelete) error {
	user, err := gn.groupService.GetGroupUser(ctx, entity.UserID, entity.GroupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	isCreator := gn.groupNoteService.IsCreator(ctx, entity.UserID, entity.GroupID)
	isAdmin := user.Permission == constants.DefaultAdminPermission

	if !isCreator && !isAdmin {
		return client_error.ErrUserNotCreatorNote
	}

	return gn.groupNoteService.Delete(ctx, entity)
}
