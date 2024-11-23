package groups

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error)
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	Delete(ctx context.Context, entity entities.GroupDelete) error
	DeleteConfirm(ctx context.Context, entity entities.GroupDeleteConfirm) error
}
