package utils

import (
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	globalConstants "go-clean/internal/constants"
	"net/http"
	"strings"
	"time"
)

func GetRequestAccessJWT(ctx echo.Context) string {
	accessJWTDirty := ctx.Request().Header["Authorization"]

	if len(accessJWTDirty) == 0 || len(accessJWTDirty[0]) == 0 {
		return ""
	}

	return strings.TrimPrefix(accessJWTDirty[0], "Bearer ")
}

func GetRequestRefreshJWT(ctx echo.Context) string {
	refreshJWTCookie, err := ctx.Cookie(constants.CookieRefreshJWT)
	if err != nil || len(refreshJWTCookie.Value) == 0 {
		return ""
	}

	return refreshJWTCookie.Value
}

func GetRequestSessionID(ctx echo.Context) string {
	refreshJWTCookie, err := ctx.Cookie(constants.CookieSessionID)
	if err != nil || len(refreshJWTCookie.Value) == 0 {
		return ""
	}

	return refreshJWTCookie.Value
}

func SetResponseAccessJWT(ctx echo.Context, accessJWT string) {
	ctx.Response().Header().Add(constants.HeaderResponseAccessJWT, accessJWT)
}

func SetResponseRefreshJWT(ctx echo.Context, refreshJWT string) {
	ctx.SetCookie(&http.Cookie{
		Name:     constants.CookieRefreshJWT,
		Value:    refreshJWT,
		Path:     "/",
		MaxAge:   int(globalConstants.RestAuthRefreshWTDuration / time.Second),
		Secure:   true,
		HttpOnly: true,
	})
}

func SetResponseSessionID(ctx echo.Context, sessionID string) {
	ctx.SetCookie(&http.Cookie{
		Name:     constants.CookieSessionID,
		Value:    sessionID,
		Path:     "/",
		MaxAge:   int(globalConstants.RestAuthRefreshWTDuration / time.Second),
		Secure:   true,
		HttpOnly: true,
	})
}

func ClearRefreshJWT(ctx echo.Context) {
	ctx.SetCookie(&http.Cookie{
		Name:     constants.CookieRefreshJWT,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}

func ClearSessionID(ctx echo.Context) {
	ctx.SetCookie(&http.Cookie{
		Name:     constants.CookieSessionID,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}
