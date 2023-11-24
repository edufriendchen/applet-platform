package common

import "errors"

var (
	ErrInternal = errors.New("internal err")
	ErrEmptyRow = errors.New("empty row")

	ErrUnauthorized          = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrUserBlocked           = errors.New("user already blocked")
	ErrUserDeactivate        = errors.New("user already deactivate")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserInvited           = errors.New("please complete registration")
	ErrInvalidUserCredential = errors.New("invalid user credential")
	ErrNotThirdPartyAccount  = errors.New("not bound to a third-party account")
)
