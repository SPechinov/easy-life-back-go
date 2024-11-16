package group

import (
	"context"
	"go-clean/config"
	"go-clean/internal/entities"
)

type Group struct {
	cfg          *config.Config
	groupService groupService
	userService  userService
}

func New(cfg *config.Config, groupService groupService, userService userService) Group {
	return Group{
		cfg:          cfg,
		groupService: groupService,
		userService:  userService,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupService.Add(ctx, entity)
}
