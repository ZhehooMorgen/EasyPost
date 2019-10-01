package userinfo

import (
	"errors"
)

// ErrUserNotFound : cannot get the target user using user ID
var ErrUserNotFound = errors.New("user not found")
