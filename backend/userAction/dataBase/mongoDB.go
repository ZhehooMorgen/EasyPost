package dataBase

import (
	"backend/dataBase"
	"backend/sensitive"
	"backend/util"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	accounts *mongo.Collection
)

func ConnectMongoDB() util.Err {
	if coll, err := dataBase.Connect(sensitive.COLLECTION_NAME_ACCOUNTS); err != nil {
		return err
	} else {
		accounts = coll
	}

	return nil
}
