package group

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

type groupService interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error)
	GetFull(ctx context.Context, entity entities.GroupGet) (*entities.GroupFull, error)
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
	InviteUser(ctx context.Context, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error
	IsDeletedGroup(ctx context.Context, groupID string) bool
	IsGroupAdmin(ctx context.Context, userID, groupID string) bool
	Delete(ctx context.Context, entity entities.GroupDelete) error
}
