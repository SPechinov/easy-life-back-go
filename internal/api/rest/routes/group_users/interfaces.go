package group_users

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	GetList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	Invite(ctx context.Context, entity entities.GroupInviteUser) error
	Exclude(ctx context.Context, entity entities.GroupExcludeUser) error
}
