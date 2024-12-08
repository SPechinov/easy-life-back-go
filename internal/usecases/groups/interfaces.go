package groups

import (
	"context"
	"go-clean/internal/entities"
	"time"
)

type codes interface {
	SetCode(ctx context.Context, key, code string, attempts int, ttl time.Duration) error
	GetCode(ctx context.Context, key string) (string, int, error)
	CompareCodes(ctx context.Context, key, code string) error
	DeleteCode(ctx context.Context, key string) error
}

type groupsService interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error)
	Delete(ctx context.Context, entity entities.GroupDeleteConfirm) error
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	IsGroupAdmin(ctx context.Context, userID, groupID string) error
	IsGroupUser(ctx context.Context, userID, groupID string) error
}
