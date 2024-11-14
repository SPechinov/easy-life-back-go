package rest_auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-clean/config"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/internal/constants"
	"go-clean/internal/constants/validation_rules"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
)

type RestAuth struct {
	store   store
	service service
	cfg     *config.Config
}

func New(cfg *config.Config, store store, service service) RestAuth {
	return RestAuth{cfg: cfg, store: store, service: service}
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

	ctx = logger.LogWithUserID(ctx, userID)

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
	// Check has user or not
	dbUser, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil && !errors.Is(err, client_error.ErrUserNotFound) {
		return err
	}

	// If user deleted restore
	if dbUser != nil {
		if !dbUser.Deleted() {
			logger.Debug(ctx, "User exist")
			return client_error.ErrUserExists
		}
		logger.Debug(ctx, "User deleted: start restore")
	}

	// Set code to store
	code, err := helpers.GenerateRandomCode(validation_rules.LenRegistrationCode)
	if err != nil {
		logger.Debug(ctx, err)
		return err
	}

	ctx = logger.LogWithConfirmationCode(ctx, code)

	logger.Debug(ctx, "Code sent")

	err = ra.store.SetRegistrationCode(ctx, data.AuthWay.GetAuthValue(), code)
	return err
}

func (ra RestAuth) RegistrationConfirm(ctx context.Context, data entities.UserAddConfirm) error {
	authWayValue := data.AuthWay.GetAuthValue()

	// Get code from store and compare
	storeCode, attempt, err := ra.store.GetRegistrationCode(ctx, authWayValue)
	if err != nil {
		return err
	}

	if storeCode != data.Code {
		if attempt+1 >= constants.MaxCodeCompareAttempt {
			logger.Debug(ctx, "Max code attempt")
			_ = ra.store.DeleteRegistrationCode(ctx, authWayValue)
			return client_error.ErrCodeMaxAttempts
		}

		logger.Debug(ctx, "Codes not equal")
		_ = ra.store.UpdateRegistrationCode(ctx, authWayValue, attempt+1)
		return client_error.ErrCodesIsNotEqual
	}

	_ = ra.store.DeleteRegistrationCode(ctx, authWayValue)

	// Check has user or not
	dbUser, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil && !errors.Is(err, client_error.ErrUserNotFound) {
		return err
	}

	if dbUser != nil {
		// If user deleted restore
		if dbUser.Deleted() {
			logger.Debug(ctx, "User deleted: restored")
			err = ra.service.RestoreUser(ctx, data)
			return err
		}
		logger.Debug(ctx, "User exist")
		return client_error.ErrUserExists
	}

	// Add user
	err = ra.service.AddUser(ctx, data)
	return err
}

func (ra RestAuth) ForgotPassword(ctx context.Context, data entities.UserForgotPassword) error {
	// Check has user or not
	dbUser, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil {
		return err
	}

	ctx = logger.LogWithUserID(ctx, dbUser.ID)

	if dbUser.Deleted() {
		logger.Debug(ctx, "User deleted")
		return client_error.ErrUserDeleted
	}

	// Set code to store
	code, err := helpers.GenerateRandomCode(validation_rules.LenRegistrationCode)
	if err != nil {
		logger.Debug(ctx, err)
		return err
	}

	ctx = logger.LogWithConfirmationCode(ctx, code)
	logger.Debug(ctx, "Code sent")

	err = ra.store.SetForgotPasswordCode(ctx, data.AuthWay.GetAuthValue(), code)
	return err
}

func (ra RestAuth) ForgotPasswordConfirm(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	authWayValue := data.AuthWay.GetAuthValue()
	// Get code from store and compare
	storeCode, attempt, err := ra.store.GetForgotPasswordCode(ctx, authWayValue)
	if err != nil {
		return err
	}

	if storeCode != data.Code {
		if attempt+1 >= constants.MaxCodeCompareAttempt {
			logger.Debug(ctx, "Max code attempt")
			_ = ra.store.DeleteForgotPasswordCode(ctx, authWayValue)
			return client_error.ErrCodeMaxAttempts
		}

		logger.Debug(ctx, "Codes not equal")
		_ = ra.store.UpdateForgotPasswordCode(ctx, authWayValue, attempt+1)
		return client_error.ErrCodesIsNotEqual
	}

	_ = ra.store.DeleteForgotPasswordCode(ctx, authWayValue)

	// Check has user or not
	dbUser, err := ra.service.GetUser(ctx, entities.UserGet{Email: data.AuthWay.Email, Phone: data.AuthWay.Phone})
	if err != nil {
		return err
	}

	ctx = logger.LogWithUserID(ctx, dbUser.ID)

	if dbUser.Deleted() {
		logger.Debug(ctx, "User deleted")
		return client_error.ErrUserDeleted
	}

	err = ra.service.UpdatePasswordUser(ctx, data)
	return err
}

func (ra RestAuth) UpdateJWT(ctx context.Context, userID, sessionID, refreshJWT string) (string, string, string, error) {
	// Check user in store
	value, err := ra.store.GetSession(ctx, userID, sessionID)
	if err != nil || value != refreshJWT {
		logger.Warn(ctx, "refreshJWT has not in redis")

		go ra.store.DeleteAllSessions(ctx, userID)
		return "", "", "", client_error.ErrNotAuthorized
	}

	// Check user in DB
	user, err := ra.service.GetUser(ctx, entities.UserGet{ID: userID})
	if err != nil || user == nil {
		logger.Warn(ctx, "user has not got in DB")
		return "", "", "", client_error.ErrNotAuthorized
	}

	// Check is valid refreshJWT
	isValid, _ := helpers.IsValidJWT(ra.cfg.HTTPAuth.JWTSecretKey, refreshJWT)
	if !isValid {
		logger.Warn(ctx, "refreshJWT is not valid")
		return "", "", "", rest_error.ErrNotAuthorized
	}

	// Delete old session
	go ra.store.DeleteSession(ctx, userID, sessionID)

	// Create JWTs
	jwtData := ra.createJWTData(userID)
	newAccessJWT, newRefreshJWT, err := ra.createJWTPair(ctx, ra.cfg, jwtData)
	if err != nil {
		return "", "", "", err
	}
	newSessionID := generateSessionID()

	// Set new session
	err = ra.store.SetSession(ctx, userID, newSessionID, newRefreshJWT)
	if err != nil {
		return "", "", "", err
	}

	return newSessionID, newAccessJWT, newRefreshJWT, nil
}

func (ra RestAuth) Logout(ctx context.Context, userID, sessionID string) {
	go ra.store.DeleteSession(ctx, userID, sessionID)
	return
}

func (ra RestAuth) LogoutAll(ctx context.Context, userID string) {
	go ra.store.DeleteAllSessions(ctx, userID)
	return
}
