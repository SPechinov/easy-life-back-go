package group

import (
	"context"
	"go-clean/internal/entities"
)

type groupService interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error)
	IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error)
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
	GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	InviteUser(ctx context.Context, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error
	IsDeletedGroup(ctx context.Context, groupID string) bool
}
