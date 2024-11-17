package group

import (
	"context"
	"go-clean/internal/entities"
)

type groupService interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) (*entities.Group, error)
	IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error)
}
