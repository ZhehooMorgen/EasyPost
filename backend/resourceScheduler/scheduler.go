package resourceScheduler

import (
	"backend/util"
	"context"
)

//Only suitable for those who know what res they will need
type Scheduler interface {
	//Register resource to Scheduler
	RegRes(resource Resource) util.Err
	//Request resource and then exec
	Request(ctx context.Context, logic func(), resources ...Resource) util.Err
}

