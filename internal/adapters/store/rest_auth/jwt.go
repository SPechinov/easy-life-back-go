package rest_auth

import (
	"context"
	"go-clean/internal/constants"
	"go-clean/pkg/logger"
)

func getKeySession(userID, sessionID string) string {
	return "http:rest-auth:session:" + userID + ":" + sessionID
}

func getKeyAllUserIDSession(userID string) string {
	return "http:rest-auth:session:" + userID + ":*"
}

func (ra *RestAuth) SetSession(ctx context.Context, userID, sessionID, refreshJWT string) error {
	err := ra.redis.SetWithTTL(getKeySession(userID, sessionID), refreshJWT, constants.RestAuthRefreshWTDuration)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

func (ra *RestAuth) GetSession(ctx context.Context, userID, sessionID string) (string, error) {
	value, err := ra.redis.Get(getKeySession(userID, sessionID))
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}
	return value, nil
}

func (ra *RestAuth) DeleteSession(ctx context.Context, userID, sessionID string) {
	err := ra.redis.Delete(getKeySession(userID, sessionID))
	if err != nil {
		logger.Error(ctx, err)
	}
	return
}

func (ra *RestAuth) DeleteAllSessions(ctx context.Context, userID string) {
	var cursor uint64

	for {
		keys, newCursor, err := ra.redis.Scan(cursor, getKeyAllUserIDSession(userID), 0)
		if err != nil {
			logger.Error(ctx, err)
		}

		for _, key := range keys {
			err = ra.redis.Delete(key)
			if err != nil {
				logger.Error(ctx, err)
			}
		}

		if newCursor == 0 {
			break
		}

		cursor = newCursor
	}

	return
}
