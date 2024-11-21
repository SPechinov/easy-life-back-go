package group

import (
	"context"
	"go-clean/internal/constants"
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

func (g *Group) IsDeletedGroup(ctx context.Context, groupID string) bool {
	group, err := g.groupDatabase.Get(ctx, entities.GroupGet{ID: groupID})
	if err != nil {
		return true
	}

	return group.Deleted()
}

func (g *Group) Delete(ctx context.Context, entity entities.GroupDelete) error {
	return g.groupDatabase.Delete(ctx, entity)
}
