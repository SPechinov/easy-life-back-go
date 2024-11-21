package group_users

import (
	"context"
	"go-clean/internal/entities"
)

type GroupUsers struct {
	groupUsersDatabase groupUsersDatabase
}

func New(groupUsersDatabase groupUsersDatabase) *GroupUsers {
	return &GroupUsers{
		groupUsersDatabase: groupUsersDatabase,
	}
}

func (ug *GroupUsers) GetList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	return ug.groupUsersDatabase.GetUsersList(ctx, entity)
}

func (ug *GroupUsers) InviteUser(ctx context.Context, entity entities.GroupInviteUser) error {
	return ug.groupUsersDatabase.InviteUser(ctx, entity)
}

func (ug *GroupUsers) ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error {
	return ug.groupUsersDatabase.ExcludeUser(ctx, entity)
}
