package communityAction

import (
	"backend/util"
)

// ErrUserNotFound : cannot get the target user using user ID
var ErrUserNotFound = util.NewBasicError("cannot get the target user using user ID",404, nil)
