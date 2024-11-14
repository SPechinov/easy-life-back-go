package auth

import (
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils"
)

func setResponseAuthData(c echo.Context, accessJWT, refreshJWT, sessionID string) {
	utils.SetResponseAccessJWT(c, accessJWT)
	utils.SetResponseRefreshJWT(c, refreshJWT)
	utils.SetResponseSessionID(c, sessionID)
}
