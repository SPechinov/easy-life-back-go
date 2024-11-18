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

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) error {
	err := g.groupDatabase.Patch(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error) {
	groupChannel := make(chan *entities.Group, 1)
	usersChannel := make(chan []entities.GroupUser, 1)
	errChannel := make(chan error, 2)

	go func() {
		group, err := g.groupDatabase.Get(ctx, entities.GroupGet{GroupID: entity.GroupID})
		if err != nil {
			errChannel <- err
			return
		}
		groupChannel <- group
	}()

	go func() {
		users, err := g.groupDatabase.GetUsersList(ctx, entities.GroupUsersListGet{GroupID: entity.GroupID})
		if err != nil {
			errChannel <- err
			return
		}
		usersChannel <- users
	}()

	var group *entities.Group
	var users []entities.GroupUser

	select {
	case group = <-groupChannel:
	case err := <-errChannel:
		return nil, err
	}

	select {
	case users = <-usersChannel:
	case err := <-errChannel:
		return nil, err
	}

	group.Users = users

	return group, nil
}

func (g *Group) IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error) {
	return g.groupDatabase.IsGroupAdmin(ctx, userID, groupID)
}

func (g *Group) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	return g.groupDatabase.GetGroupUser(ctx, userID, groupID)
}
