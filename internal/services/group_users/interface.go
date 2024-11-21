package group_users

import (
	"context"
	"go-clean/internal/entities"
)

type groupUsersDatabase interface {
	GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	InviteUser(ctx context.Context, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error
}
