package rest_auth

import (
	"context"
	"go-clean/config"
	"go-clean/internal/constants"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
)

func (ra RestAuth) createJWTData(userID string) map[string]string {
	return map[string]string{
		constants.UserIDInJWTKey: userID,
	}
}

func (ra RestAuth) createJWTPair(ctx context.Context, cfg *config.Config, jwtData map[string]string) (string, string, error) {
	accessJWT, err := helpers.CreateJWT(cfg.HTTPAuth.JWTSecretKey, constants.RestAuthAccessJWTDuration, jwtData)
	if err != nil {
		logger.Error(ctx, err)
		return "", "", err
	}

	refreshJWT, err := helpers.CreateJWT(cfg.HTTPAuth.JWTSecretKey, constants.RestAuthRefreshWTDuration, jwtData)
	if err != nil {
		logger.Error(ctx, err)
		return "", "", err
	}

	return accessJWT, refreshJWT, nil
}
