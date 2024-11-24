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
	errKeyUserInvited       = "userInvited"
	errKeyInvalidParams     = "invalidParams"
	errKeyUserNotAdminGroup = "userNotAdminGroup"
	errKeyUserNotInGroup    = "userNotInGroup"
	errKeyUserAdminGroup    = "userAdminGroup"
	errKeyGroupNotExist     = "groupNotExist"
	errKeyNoteDeleted       = "noteDeleted"
	errKeyNoteNotExists     = "noteNotExists"
	errKeyUserNotCreateNote = "userNotCreateNote"
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
var ErrUserInvited = New(http.StatusBadRequest, errKeyUserInvited)
var ErrInvalidParams = New(http.StatusBadRequest, errKeyInvalidParams)
var ErrUserNotAdminGroup = New(http.StatusBadRequest, errKeyUserNotAdminGroup)
var ErrUserAdminGroup = New(http.StatusBadRequest, errKeyUserAdminGroup)
var ErrUserNotInGroup = New(http.StatusBadRequest, errKeyUserNotInGroup)
var ErrGroupNotExist = New(http.StatusBadRequest, errKeyGroupNotExist)
var ErrNoteDeleted = New(http.StatusBadRequest, errKeyNoteDeleted)
var ErrNoteNotExists = New(http.StatusBadRequest, errKeyNoteNotExists)
var ErrUserNotCreatorNote = New(http.StatusBadRequest, errKeyUserNotCreateNote)
var ErrSomethingHappen = New(http.StatusInternalServerError, errKeySomethingHappen)
