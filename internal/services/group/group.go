package group

import (
	"context"
	"go-clean/internal/entities"
)

type Group struct {
	groupDatabase groupDatabase
}

func New(groupDatabase groupDatabase) *Group {
	return &Group{
		groupDatabase: groupDatabase,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupDatabase.Add(ctx, entity)
}
