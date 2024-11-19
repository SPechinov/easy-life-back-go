package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-clean/config"
	restConstants "go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/internal/constants"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
)

func isValidJWTPair(secretKey, accessJWT, refreshJWT string) (accessToken *jwt.Token, refreshToken *jwt.Token, err error) {
	isValid, accessToken := helpers.IsValidJWT(secretKey, accessJWT)
	if !isValid {
		return nil, nil, rest_error.ErrNotAuthorized
	}

	isValid, refreshToken = helpers.IsValidJWT(secretKey, refreshJWT)
	if !isValid {
		return nil, nil, rest_error.ErrNotAuthorized
	}

	return accessToken, refreshToken, nil
}

func isValidDataInJWTPair(accessToken *jwt.Token, refreshToken *jwt.Token) (userID string, err error) {
	userIDFromAccess, ok := accessToken.Claims.(jwt.MapClaims)[constants.UserIDInJWTKey].(string)
	if !ok || userIDFromAccess == "" {
		return "", rest_error.ErrNotAuthorized
	}

	userIDFromRefresh, ok := refreshToken.Claims.(jwt.MapClaims)[constants.UserIDInJWTKey].(string)
	if !ok || userIDFromRefresh == "" {
		return "", rest_error.ErrNotAuthorized
	}

	if userIDFromAccess != userIDFromRefresh {
		return "", rest_error.ErrNotAuthorized
	}

	return userIDFromAccess, nil
}

// AuthMiddleware check: should be valid sessionID and data in accessJWT and refreshJWT should be the same
// UserID, accessJWT, refreshJWT, sessionID contain in echo context
func AuthMiddleware(cfg *config.Config) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			// Check SessionID
			sessionID := utils.GetRequestSessionID(echoCtx)
			err := uuid.Validate(sessionID)
			if err != nil {
				return rest_error.ErrNotAuthorized
			}

			// Check JWTs
			accessJWT := utils.GetRequestAccessJWT(echoCtx)
			refreshJWT := utils.GetRequestRefreshJWT(echoCtx)

			accessToken, refreshToken, err := isValidJWTPair(cfg.HTTPAuth.JWTSecretKey, accessJWT, refreshJWT)
			if err != nil {
				return err
			}

			userID, err := isValidDataInJWTPair(accessToken, refreshToken)
			if err != nil {
				return err
			}

			// Save in context
			echoCtx.Set(constants.CTXUserIDKey, userID)

			// Logging UserID
			if loggerCtx, ok := echoCtx.Get(restConstants.CTXLoggerInCTX).(context.Context); !ok {
				logger.Error(loggerCtx, "No context")
			} else {
				loggerCtx = logger.WithUserID(loggerCtx, userID)
				echoCtx.Set(restConstants.CTXLoggerInCTX, loggerCtx)
			}

			return next(echoCtx)
		}
	}
}
