package group

import (
	"context"
	"go-clean/internal/entities"
)

type groupDatabase interface {
	Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error)
}