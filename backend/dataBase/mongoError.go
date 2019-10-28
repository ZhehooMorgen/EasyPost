package dataBase

import "backend/util"

func NewMongoDBError(err error)util.Err{
	return util.NewBasicError("error occured during mongo query",-1,err)
}
