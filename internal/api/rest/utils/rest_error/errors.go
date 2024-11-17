package rest_error

import "net/http"

const (
	errKeyNotAuthorized     = "notAuthorized"
	errKeyInvalidBodyData   = "invalidBodyData"
	errKeyIncorrectPassword = "incorrectPassword"
	errKeyUserNotFound      = "userNotFound"
	errKeyCodeDidNotSend    = "codeDidNotSend"
	errKeyUserExists        = "userExists"
	errKeyCodeMaxAttempts   = "codeMaxAttempts"
	errKeyCodesIsNotEqual   = "codesIsNotEqual"
	errKeyUserDeleted       = "userDeleted"
	errKeyInvalidParams     = "invalidParams"
	errKeyUserNotGroupAdmin = "userNotGroupAdmin"
	errKeySomethingHappen   = "somethingHappen"
)

var ErrNotAuthorized = New(http.StatusUnauthorized, errKeyNotAuthorized)
var ErrInvalidBodyData = New(http.StatusBadRequest, errKeyInvalidBodyData)
var ErrIncorrectPassword = New(http.StatusBadRequest, errKeyIncorrectPassword)
var ErrUserNotFound = New(http.StatusBadRequest, errKeyUserNotFound)
var ErrCodeDidNotSent = New(http.StatusBadRequest, errKeyCodeDidNotSend)
var ErrUserExists = New(http.StatusBadRequest, errKeyUserExists)
var ErrCodeMaxAttempts = New(http.StatusBadRequest, errKeyCodeMaxAttempts)
var ErrCodesIsNotEqual = New(http.StatusBadRequest, errKeyCodesIsNotEqual)
var ErrUserDeleted = New(http.StatusBadRequest, errKeyUserDeleted)
var ErrInvalidParams = New(http.StatusBadRequest, errKeyInvalidParams)
var ErrUserNotAdminGroup = New(http.StatusBadRequest, errKeyUserNotGroupAdmin)
var ErrSomethingHappen = New(http.StatusInternalServerError, errKeySomethingHappen)
