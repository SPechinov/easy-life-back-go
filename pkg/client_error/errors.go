package client_error

var ErrNotAuthorized = New("notAuthorized")
var ErrIncorrectPassword = New("incorrectPassword")
var ErrUserNotFound = New("userNotFound")
var ErrCodeIsNotInRedis = New("codeIsNotInRedis")
var ErrUserExists = New("userExists")
var ErrCodeMaxAttempts = New("codeMaxAttempts")
var ErrCodesIsNotEqual = New("codesIsNotEqual")
var ErrUserDeleted = New("userDeleted")
var ErrUserNotAdminGroup = New("userNotAdminGroup")
var ErrUserNotInGroup = New("userNotInGroup")
