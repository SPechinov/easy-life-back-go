package group

import (
	"context"
	"go-clean/internal/entities"
)

type useCases interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, adminID string, entity entities.GroupPatch) (*entities.Group, error)
}
