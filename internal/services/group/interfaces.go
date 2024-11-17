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
	GetUsersList(ctx context.Context, entity entities.GroupUsersListGet) ([]entities.GroupUser, error)
	GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error)
}
