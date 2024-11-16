package group

import (
	"context"
	"go-clean/internal/entities"
)

type Group struct {
	groupDatabase groupDatabase
	userService   userService
}

func New(groupDatabase groupDatabase, userService userService) *Group {
	return &Group{
		groupDatabase: groupDatabase,
		userService:   userService,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	group, err := g.groupDatabase.Add(ctx, entity)
	if err != nil {
		return nil, err
	}

	user, err := g.userService.GetUser(ctx, entities.UserGet{
		ID: entity.AdminID,
	})
	if err != nil {
		return nil, err
	}

	group.Admin = *user
	group.Users = append(group.Users, *user)

	return group, nil
}
