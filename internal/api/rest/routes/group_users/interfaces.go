package group_users

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	GetList(ctx context.Context, userID string, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	Invite(ctx context.Context, adminID string, entity entities.GroupInviteUser) error
	Exclude(ctx context.Context, adminID string, entity entities.GroupExcludeUser) error
}
