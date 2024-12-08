package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restAuthController) handlerLogin(echoCTX echo.Context, ctx context.Context, dto *LoginDTO) error {
	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithPassword(ctx, dto.Password)

	sessionID, accessJWT, refreshJWT, err := controller.useCases.Login(
		ctx,
		entities.UserLogin{
			AuthWay: entities.UserAuthWay{
				Email: dto.Email,
				Phone: dto.Phone,
			},
			Password: dto.Password,
		},
	)

	if err != nil {
		return err
	}

	setResponseAuthData(echoCTX, accessJWT, refreshJWT, sessionID)
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerRegistration(echoCTX echo.Context, ctx context.Context, dto *RegistrationDTO) error {
	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)

	err := controller.useCases.Registration(
		ctx,
		entities.UserAdd{
			AuthWay: entities.UserAuthWay{
				Email: dto.Email,
				Phone: dto.Phone,
			},
		},
	)
	if err != nil {
		return err
	}
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerRegistrationConfirm(echoCTX echo.Context, ctx context.Context, dto *RegistrationConfirmDTO) error {
	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithConfirmationCode(ctx, dto.Code)
	ctx = logger.WithPassword(ctx, dto.Password)

	err := controller.useCases.RegistrationConfirm(
		ctx,
		entities.UserAddConfirm{
			AuthWay: entities.UserAuthWay{
				Email: dto.Email,
				Phone: dto.Phone,
			},
			FirstName: dto.FirstName,
			Password:  dto.Password,
			Code:      dto.Code,
		},
	)
	if err != nil {
		return err
	}
	return echoCTX.NoContent(http.StatusCreated)
}

func (controller *restAuthController) handlerForgotPassword(echoCTX echo.Context, ctx context.Context, dto *ForgotPasswordDTO) error {
	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)

	err := controller.useCases.ForgotPassword(
		ctx,
		entities.UserForgotPassword{
			AuthWay: entities.UserAuthWay{
				Email: dto.Email,
				Phone: dto.Phone,
			},
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerForgotPasswordConfirm(echoCTX echo.Context, ctx context.Context, dto *ForgotPasswordConfirmDTO) error {
	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithPassword(ctx, dto.Password)

	err := controller.useCases.ForgotPasswordConfirm(
		ctx,
		entities.UserForgotPasswordConfirm{
			AuthWay: entities.UserAuthWay{
				Email: dto.Email,
				Phone: dto.Phone,
			},
			Password: dto.Password,
			Code:     dto.Code,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerUpdateJWT(echoCTX echo.Context, ctx context.Context) error {
	// Check UUID
	sessionID := utils.GetRequestSessionID(echoCTX)
	err := uuid.Validate(sessionID)
	if err != nil {
		logger.Warn(ctx, "has not got session id")
		return rest_error.ErrNotAuthorized
	}

	ctx = logger.WithSessionID(ctx, sessionID)

	// Check refreshJWT
	refreshJWT := utils.GetRequestRefreshJWT(echoCTX)
	isValid, token := helpers.IsValidJWT(controller.cfg.HTTPAuth.JWTSecretKey, refreshJWT)
	if !isValid {
		logger.Error(ctx, "refresh token invalid")
		return rest_error.ErrNotAuthorized
	}

	// Check refreshJWT data
	userID, ok := token.Claims.(jwt.MapClaims)[globalConstants.UserIDInJWTKey].(string)
	if !ok {
		logger.Error(ctx, "refresh token has not got correct data")
		return rest_error.ErrNotAuthorized
	}

	newSessionID, newAccessJWT, newRefreshJWT, err := controller.useCases.UpdateJWT(
		ctx,
		entities.UserUpdateJWT{
			ID:         userID,
			SessionID:  sessionID,
			RefreshJWT: refreshJWT,
		},
	)
	if err != nil {
		return rest_error.ErrNotAuthorized
	}

	setResponseAuthData(echoCTX, newAccessJWT, newRefreshJWT, newSessionID)

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerLogout(echoCTX echo.Context, ctx context.Context, userID string) error {
	// Check SessionID
	sessionID := utils.GetRequestSessionID(echoCTX)
	err := uuid.Validate(sessionID)
	if err != nil {
		return rest_error.ErrNotAuthorized
	}

	ctx = logger.WithSessionID(ctx, sessionID)

	controller.useCases.Logout(
		ctx,
		entities.UserLogout{
			ID:        userID,
			SessionID: sessionID,
		},
	)

	utils.ClearRefreshJWT(echoCTX)
	utils.ClearSessionID(echoCTX)
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerLogoutAll(echoCTX echo.Context, ctx context.Context, userID string) error {
	controller.useCases.LogoutAll(ctx, entities.UserLogoutAll{ID: userID})

	utils.ClearRefreshJWT(echoCTX)
	utils.ClearSessionID(echoCTX)

	return echoCTX.NoContent(http.StatusNoContent)
}
