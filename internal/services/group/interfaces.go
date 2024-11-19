package group

import (
	"context"
	"go-clean/internal/entities"
)

type userService interface {
	GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error)
}

type groupDatabase interface {
	Add(ctx context.Context, entity entities.GroupAdd) (groupID string, err error)
	Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error)
	Patch(ctx context.Context, entity entities.GroupPatch) error
	IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error)
	GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error)
	GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error)
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
	InviteUser(ctx context.Context, entity entities.GroupInviteUser) error
	ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error
	IsDeletedGroup(ctx context.Context, groupID string) bool
}
