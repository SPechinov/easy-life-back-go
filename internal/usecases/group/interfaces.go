package group

import (
	"context"
	"go-clean/internal/entities"
)

type userService interface {
	GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error)
}

type groupService interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
}
