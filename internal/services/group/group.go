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
	groupID, err := g.groupDatabase.Add(ctx, entity)
	if err != nil {
		return nil, err
	}

	group, err := g.groupDatabase.Get(ctx, entities.GroupGet{
		GroupID: groupID,
	})
	if err != nil {
		return nil, err
	}

	group.Users = make([]entities.GroupUser, 0)

	return group, nil
}

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) (*entities.Group, error) {
	err := g.groupDatabase.Patch(ctx, entity)
	if err != nil {
		return nil, err
	}

	group, err := g.groupDatabase.Get(ctx, entities.GroupGet{
		GroupID: entity.GroupID,
	})
	if err != nil {
		return nil, err
	}

	users, err := g.groupDatabase.GetUsersList(ctx, entities.GroupUsersListGet{
		GroupID: entity.GroupID,
	})
	if err != nil {
		return nil, err
	}
	group.Users = users

	return group, nil
}

func (g *Group) IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error) {
	return g.groupDatabase.IsGroupAdmin(ctx, userID, groupID)
}
