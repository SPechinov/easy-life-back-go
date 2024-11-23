package group

import (
	"context"
	"go-clean/internal/entities"
)

type groupDatabase interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error)
	Delete(ctx context.Context, entity entities.GroupDeleteConfirm) error
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
}
