package group_users

import (
	"context"
	"go-clean/internal/entities"
)

type groupUsersService interface {
	GetList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	InviteUser(ctx context.Context, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error
}

type groupService interface {
	IsGroupAdmin(ctx context.Context, userID, groupID string) error
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
}
