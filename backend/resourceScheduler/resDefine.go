package resourceScheduler

import (
	"github.com/satori/go.uuid"
)

type Type string

//Resource is something that can not be visited concurrently
//Register it to scheduler and write code following the guide can help protect the correctness of logic
type Resource interface {
	//resID should be it's own UUID, same resource instance should return same and different instance should return different
	//type marks the type of resource
	//If two resource instance share the same type and share the same zero UUID, scheduler will consider that they are replaceable
	Definition() (resType Type, resID uuid.UUID)
}
