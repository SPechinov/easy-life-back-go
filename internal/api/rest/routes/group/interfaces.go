package group

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Get(ctx context.Context, userID string, entity entities.GroupGet) (*entities.Group, error)
	GetUsersList(ctx context.Context, userID string, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	Patch(ctx context.Context, adminID string, entity entities.GroupPatch) error
	InviteUser(ctx context.Context, adminID string, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, adminID string, entity entities.GroupExcludeUser) error
}
