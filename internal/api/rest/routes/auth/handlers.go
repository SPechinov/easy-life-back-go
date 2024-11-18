package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restAuthController) handlerLogin(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(LoginDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithPassword(ctx, dto.Password)
	logger.Debug(ctx, "Start")

	err = validateLoginDTO(dto)
	if err != nil {
		return err
	}

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

	setResponseAuthData(c, accessJWT, refreshJWT, sessionID)

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerRegistration(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(RegistrationDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	logger.Debug(ctx, "Start")

	err = validateRegistrationDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.Registration(
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

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerRegistrationConfirm(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(RegistrationConfirmDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithConfirmationCode(ctx, dto.Code)
	ctx = logger.WithPassword(ctx, dto.Password)
	logger.Debug(ctx, "Start")

	err = validateRegistrationConfirmDTO(dto)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = controller.useCases.RegistrationConfirm(
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

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusCreated)
}

func (controller *restAuthController) handlerForgotPassword(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(ForgotPasswordDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	logger.Debug(ctx, "Start")

	err = validateForgotPasswordDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.ForgotPassword(
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

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerForgotPasswordConfirm(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(ForgotPasswordConfirmDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithRestAuthData(ctx, dto.Email, dto.Phone)
	ctx = logger.WithPassword(ctx, dto.Password)
	logger.Debug(ctx, "Start")

	err = validateForgotPasswordConfirmDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.ForgotPasswordConfirm(
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

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerUpdateJWT(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	logger.Debug(ctx, "Start")
	// Check UUID
	sessionID := utils.GetRequestSessionID(c)
	err := uuid.Validate(sessionID)
	if err != nil {
		logger.Warn(ctx, "has not got session id")
		return rest_error.ErrNotAuthorized
	}

	ctx = logger.WithSessionID(ctx, sessionID)

	// Check refreshJWT
	refreshJWT := utils.GetRequestRefreshJWT(c)
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

	newSessionID, newAccessJWT, newRefreshJWT, err := controller.useCases.UpdateJWT(ctx, userID, sessionID, refreshJWT)
	if err != nil {
		return rest_error.ErrNotAuthorized
	}

	setResponseAuthData(c, newAccessJWT, newRefreshJWT, newSessionID)

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerLogout(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	logger.Debug(ctx, "Start")

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	// Check SessionID
	sessionID := utils.GetRequestSessionID(c)
	err := uuid.Validate(sessionID)
	if err != nil {
		return rest_error.ErrNotAuthorized
	}

	ctx = logger.WithSessionID(ctx, sessionID)

	controller.useCases.Logout(ctx, userID, sessionID)

	utils.ClearRefreshJWT(c)
	utils.ClearSessionID(c)

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restAuthController) handlerLogoutAll(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	logger.Debug(ctx, "Start")
	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	controller.useCases.LogoutAll(ctx, userID)

	utils.ClearRefreshJWT(c)
	utils.ClearSessionID(c)

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}
