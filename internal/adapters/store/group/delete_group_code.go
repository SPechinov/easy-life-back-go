package group

import (
	"context"
	"errors"
	"go-clean/pkg/client_error"
	"go-clean/pkg/logger"
	"go-clean/pkg/redis"
	"time"
)

func getKeyDeleteGroup(groupID string) string {
	return "group:delete:" + groupID
}

func (g *Group) SetGroupDeleteCode(ctx context.Context, groupID, code string) error {
	err := g.redis.SetCode(getKeyDeleteGroup(groupID), code, 0, 10*time.Minute)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (g *Group) GetGroupDeleteCode(ctx context.Context, groupID string) (string, int, error) {
	code, attempts, err := g.redis.GetCode(getKeyDeleteGroup(groupID))
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return "", 0, client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return "", 0, err
	}

	return code, attempts, nil
}

func (g *Group) UpdateGroupDeleteCode(ctx context.Context, groupID string, attempts int) error {
	err := g.redis.UpdateCode(getKeyDeleteGroup(groupID), attempts)
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (g *Group) DeleteGroupDeleteCode(ctx context.Context, groupID string) error {
	err := g.redis.Delete(getKeyDeleteGroup(groupID))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
