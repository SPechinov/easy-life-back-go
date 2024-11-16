package group

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/postgres"
)

type Group struct {
	postgres *postgres.Postgres
}

func New(postgres *postgres.Postgres) *Group {
	return &Group{
		postgres: postgres,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return &entities.Group{}, nil
}
