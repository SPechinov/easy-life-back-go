package group_users

import (
	"context"
	"go-clean/config"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type GroupUsers struct {
	cfg               *config.Config
	groupService      groupService
	groupUsersService groupUsersService
}

func New(cfg *config.Config, groupUsersService groupUsersService, groupService groupService) *GroupUsers {
	return &GroupUsers{
		cfg:               cfg,
		groupService:      groupService,
		groupUsersService: groupUsersService,
	}
}

func (gu *GroupUsers) GetList(ctx context.Context, userID string, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	if gu.groupService.IsDeletedGroup(ctx, entity.GroupID) {
		return nil, client_error.ErrGroupDeleted
	}

	user, err := gu.groupService.GetGroupUser(ctx, userID, entity.GroupID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}

	return gu.groupUsersService.GetList(ctx, entity)
}

func (gu *GroupUsers) Invite(ctx context.Context, adminID string, entity entities.GroupInviteUser) error {
	if gu.groupService.IsDeletedGroup(ctx, entity.GroupID) {
		return client_error.ErrGroupDeleted
	}
	if !gu.groupService.IsGroupAdmin(ctx, adminID, entity.GroupID) {
		return client_error.ErrUserNotAdminGroup
	}

	return gu.groupUsersService.InviteUser(ctx, entity)
}

func (gu *GroupUsers) Exclude(ctx context.Context, adminID string, entity entities.GroupExcludeUser) error {
	if gu.groupService.IsDeletedGroup(ctx, entity.GroupID) {
		return client_error.ErrGroupDeleted
	}
	if !gu.groupService.IsGroupAdmin(ctx, adminID, entity.GroupID) {
		return client_error.ErrUserNotAdminGroup
	}

	return gu.groupUsersService.ExcludeUser(ctx, entity)
}
