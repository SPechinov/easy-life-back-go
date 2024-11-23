package rest_auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-clean/config"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/internal/constants/validation_rules"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"time"
)

func getKeyUserRegistrationCode(key string) string {
	return "http:rest-auth:reg-code:" + key
}

func getKeyUserForgotPasswordCode(key string) string {
	return "http:rest-auth:forgot-password-code:" + key
}

type RestAuth struct {
	store   store
	codes   codes
	service service
	cfg     *config.Config
}

func New(cfg *config.Config, store store, service service, codes codes) RestAuth {
	return RestAuth{cfg: cfg, store: store, service: service, codes: codes}
}

func generateSessionID() string {
	return uuid.New().String()
}

func (ra RestAuth) Login(ctx context.Context, data entities.UserLogin) (sessionID, accessJWT, refreshJWT string, err error) {
	// Check user
	user, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil {
		return "", "", "", err
	}
	userID := user.ID

	if !helpers.CheckPasswordHash(data.Password, user.Password) {
		logger.Debug(ctx, "Incorrect password")
		return "", "", "", client_error.ErrIncorrectPassword
	}

	// Create JWTs
	jwtData := ra.createJWTData(userID)

	accessJWT, refreshJWT, err = ra.createJWTPair(ctx, ra.cfg, jwtData)
	if err != nil {
		return "", "", "", err
	}
	sessionID = generateSessionID()

	err = ra.store.SetSession(ctx, userID, sessionID, refreshJWT)
	if err != nil {
		return "", "", "", err
	}

	return sessionID, accessJWT, refreshJWT, nil
}

func (ra RestAuth) Registration(ctx context.Context, data entities.UserAdd) error {
	deletedAt, err := ra.service.GetUserDeletedTime(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	// DB error
	if err != nil && !errors.Is(err, client_error.ErrUserNotFound) {
		return err
	}

	// User exist
	if deletedAt == nil && !errors.Is(err, client_error.ErrUserNotFound) {
		logger.Debug(ctx, "User exist")
		return client_error.ErrUserExists
	}

	if deletedAt != nil {
		logger.Debug(ctx, "User deleted: start restore")
	} else if errors.Is(err, client_error.ErrUserNotFound) {
		logger.Debug(ctx, "Registration start")
	}

	// Set code to store
	code := helpers.GenerateRandomCode(validation_rules.LenRegistrationCode)
	ctx = logger.WithConfirmationCode(ctx, code)
	logger.Debug(ctx, "Code sent")

	err = ra.codes.SetCode(
		ctx,
		getKeyUserRegistrationCode(data.AuthWay.GetAuthValue()),
		code,
		0,
		time.Minute*10,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ra RestAuth) RegistrationConfirm(ctx context.Context, data entities.UserAddConfirm) error {
	authWayValue := data.AuthWay.GetAuthValue()

	err := ra.codes.CompareCodes(ctx, getKeyUserRegistrationCode(authWayValue), data.Code)
	if err != nil {
		return err
	}

	deletedAt, err := ra.service.GetUserDeletedTime(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	// DB error
	if err != nil && !errors.Is(err, client_error.ErrUserNotFound) {
		return err
	}

	// User exist
	if deletedAt == nil && !errors.Is(err, client_error.ErrUserNotFound) {
		logger.Debug(ctx, "User exist")
		return client_error.ErrUserExists
	}

	if deletedAt != nil {
		logger.Debug(ctx, "User deleted: restored")
		err = ra.service.RestoreUser(ctx, data)
	} else if errors.Is(err, client_error.ErrUserNotFound) {
		logger.Debug(ctx, "Registration confirm")
		err = ra.service.AddUser(ctx, data)
	}

	return err
}

func (ra RestAuth) ForgotPassword(ctx context.Context, data entities.UserForgotPassword) error {
	// Check has user or not
	_, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil {
		return err
	}

	// Set code to store
	code := helpers.GenerateRandomCode(validation_rules.LenRegistrationCode)
	ctx = logger.WithConfirmationCode(ctx, code)
	logger.Debug(ctx, "Code sent")

	err = ra.codes.SetCode(
		ctx,
		getKeyUserForgotPasswordCode(data.AuthWay.GetAuthValue()),
		code,
		0,
		time.Minute*10,
	)
	if err != nil {
		return err
	}
	return err
}

func (ra RestAuth) ForgotPasswordConfirm(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	authWayValue := data.AuthWay.GetAuthValue()

	err := ra.codes.CompareCodes(ctx, getKeyUserForgotPasswordCode(authWayValue), data.Code)
	if err != nil {
		return err
	}

	// Check has user or not
	_, err = ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil {
		return err
	}

	err = ra.service.UpdatePasswordUser(ctx, data)
	return err
}

func (ra RestAuth) UpdateJWT(ctx context.Context, entity entities.UserUpdateJWT) (string, string, string, error) {
	// Check user in store
	value, err := ra.store.GetSession(ctx, entity.ID, entity.SessionID)
	if err != nil || value != entity.RefreshJWT {
		logger.Warn(ctx, "refreshJWT has not in redis")

		go ra.store.DeleteAllSessions(ctx, entity.ID)
		return "", "", "", client_error.ErrNotAuthorized
	}

	// Check user in DB
	user, err := ra.service.GetUser(ctx, entities.UserGet{ID: entity.ID})
	if err != nil || user == nil {
		logger.Warn(ctx, "user has not got in DB")
		return "", "", "", client_error.ErrNotAuthorized
	}

	// Check is valid refreshJWT
	isValid, _ := helpers.IsValidJWT(ra.cfg.HTTPAuth.JWTSecretKey, entity.RefreshJWT)
	if !isValid {
		logger.Warn(ctx, "refreshJWT is not valid")
		return "", "", "", rest_error.ErrNotAuthorized
	}

	// Delete old session
	go ra.store.DeleteSession(ctx, entity.ID, entity.SessionID)

	// Create JWTs
	jwtData := ra.createJWTData(entity.ID)
	newAccessJWT, newRefreshJWT, err := ra.createJWTPair(ctx, ra.cfg, jwtData)
	if err != nil {
		return "", "", "", err
	}
	newSessionID := generateSessionID()

	// Set new session
	err = ra.store.SetSession(ctx, entity.ID, newSessionID, newRefreshJWT)
	if err != nil {
		return "", "", "", err
	}

	return newSessionID, newAccessJWT, newRefreshJWT, nil
}

func (ra RestAuth) Logout(ctx context.Context, entity entities.UserLogout) {
	go ra.store.DeleteSession(ctx, entity.ID, entity.SessionID)
	return
}

func (ra RestAuth) LogoutAll(ctx context.Context, entity entities.UserLogoutAll) {
	go ra.store.DeleteAllSessions(ctx, entity.ID)
	return
}
