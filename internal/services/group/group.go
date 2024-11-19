package group

import (
	"context"
	"go-clean/internal/constants"
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
		ID: groupID,
	})
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) error {
	err := g.groupDatabase.Patch(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) GetFull(ctx context.Context, entity entities.GroupGet) (*entities.GroupFull, error) {
	groupChannel := make(chan *entities.Group, 1)
	usersChannel := make(chan []entities.GroupUser, 1)
	errChannel := make(chan error, 2)

	go func() {
		group, err := g.groupDatabase.Get(ctx, entities.GroupGet{ID: entity.ID})
		if err != nil {
			errChannel <- err
			return
		}
		groupChannel <- group
	}()

	go func() {
		users, err := g.groupDatabase.GetUsersList(ctx, entities.GroupGetUsersList{ID: entity.ID})
		if err != nil {
			errChannel <- err
			return
		}
		usersChannel <- users
	}()

	var groupInfo *entities.Group
	var users []entities.GroupUser

	select {
	case groupInfo = <-groupChannel:
	case err := <-errChannel:
		return nil, err
	}

	select {
	case users = <-usersChannel:
	case err := <-errChannel:
		return nil, err
	}

	group := entities.GroupFull{
		Group: *groupInfo,
		Users: users,
	}

	group.Users = users

	return &group, nil
}

func (g *Group) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupDatabase.GetList(ctx, entity)
}

func (g *Group) Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error) {
	groupInfo, err := g.groupDatabase.Get(ctx, entities.GroupGet{ID: entity.ID})
	if err != nil {
		return nil, err
	}

	return groupInfo, nil
}

func (g *Group) IsGroupAdmin(ctx context.Context, userID, groupID string) bool {
	user, err := g.groupDatabase.GetGroupUser(ctx, userID, groupID)
	if err != nil || user == nil {
		return false
	}

	return user.Permission == constants.DefaultAdminPermission
}

func (g *Group) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	return g.groupDatabase.GetGroupUser(ctx, userID, groupID)
}

func (g *Group) GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	return g.groupDatabase.GetUsersList(ctx, entity)
}

func (g *Group) InviteUser(ctx context.Context, entity entities.GroupInviteUser) error {
	return g.groupDatabase.InviteUser(ctx, entity)
}

func (g *Group) ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error {
	return g.groupDatabase.ExcludeUser(ctx, entity)
}

func (g *Group) IsDeletedGroup(ctx context.Context, groupID string) bool {
	group, err := g.groupDatabase.Get(ctx, entities.GroupGet{ID: groupID})
	if err != nil {
		return true
	}

	return group.Deleted()
}
