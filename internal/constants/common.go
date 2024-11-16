package constants

import "time"

const (
	MaxCodeCompareAttempt = 5
	UserIDInJWTKey        = "UserID"
	CTXUserIDKey          = "UserID"

	RestAuthAccessJWTDuration = time.Minute * 5
	RestAuthRefreshWTDuration = time.Hour * 720

	DefaultAdminPermission = 777
)
