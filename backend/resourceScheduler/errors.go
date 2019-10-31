package resourceScheduler

import "backend/util"

func NewInvalidResourceError(e error)util.Err{
	return util.NewBasicError("res requested not found",404,e)
}

func NewUseOfNoneInitScheduler(e error)util.Err{
	return util.NewBasicError("use a scheduler before init, probably bad code logic",10,e)
}

func NewIllegalUUID(e error)util.Err{
	return util.NewBasicError("UUID is illegal or zero",10,e)
}
