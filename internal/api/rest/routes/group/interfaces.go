package group

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, adminID string, entity entities.GroupPatch) error
	Get(ctx context.Context, userID string, entity entities.GroupGetInfo) (*entities.Group, error)
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	Delete(ctx context.Context, adminID, groupID string) error
	DeleteConfirm(ctx context.Context, adminID, groupID, code string) error
}
