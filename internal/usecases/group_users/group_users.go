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

func (gu *GroupUsers) GetList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	user, err := gu.groupService.GetGroupUser(ctx, entity.UserID, entity.GroupID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}

	return gu.groupUsersService.GetList(ctx, entity)
}

func (gu *GroupUsers) Invite(ctx context.Context, entity entities.GroupInviteUser) error {
	err := gu.groupService.IsGroupAdmin(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}
	return gu.groupUsersService.InviteUser(ctx, entity)
}

func (gu *GroupUsers) Exclude(ctx context.Context, entity entities.GroupExcludeUser) error {
	err := gu.groupService.IsGroupAdmin(ctx, entity.UserID, entity.GroupID)
	if err != nil {
		return err
	}
	return gu.groupUsersService.ExcludeUser(ctx, entity)
}
