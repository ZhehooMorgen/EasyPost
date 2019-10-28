package querys

import (
	"backend/dataBase"
	"backend/sensitive"
	"backend/util"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"sync"
)

const (
	ACCOUNT_TITLE = "title"
	SETTING_NAME = "name"
	SETTING_VALUE = "value"
	SETTING_ACCOUNT_NUM_ACCU = "accountNumAccumulate"
)

var (
	accounts           *mongo.Collection
	settings           *mongo.Collection
	createAccountMutex sync.Mutex
)

func InitMongoDBConnection() util.Err {
	if coll, err := dataBase.Connect(sensitive.COLLECTION_NAME_ACCOUNTS); err != nil {
		return err
	} else {
		accounts = coll
	}
	if coll, err := dataBase.Connect(sensitive.COLLECTION_NAME_SETTINGS); err != nil {
		return err
	} else {
		settings = coll
	}
	return nil
}

func ExistAccountTitle(ctx context.Context, title string) (bool, util.Err) {
	Cursor, e := accounts.Find(ctx, bson.D{
		{ACCOUNT_TITLE, title},
	})
	if e != nil {
		return false, dataBase.NewMongoDBError(e)
	}
	for Cursor.Next(ctx) {
		return true, nil
	}
	return false, nil
}

func CreateAccount(ctx context.Context, title string, pwd string) (userID string, err util.Err) {
	createAccountMutex.Lock()
	defer createAccountMutex.Unlock()
	result := struct {
		Key   string
		Value int64
	}{}
	if e := settings.FindOne(ctx, bson.D{
		{SETTING_NAME, SETTING_ACCOUNT_NUM_ACCU},
	}).Decode(&result); e != nil {
		return "", dataBase.NewMongoDBError(e)
	}
	result.Value+=rand.Int63n(9999)
	if _, e := settings.UpdateOne(ctx, bson.D{
		{SETTING_NAME, SETTING_ACCOUNT_NUM_ACCU},
	},
	bson.D{
		{"$set",bson.D{
			{SETTING_VALUE,result.Value},
		} },
	}); e != nil {
		return "", dataBase.NewMongoDBError(e)
	}
	if _, e := accounts.InsertOne(ctx, bson.D{
		{"num",result.Value},
		{"name", title},
		{"pwd", pwd},
	}); e != nil {
		return "", dataBase.NewMongoDBError(e)
	}
	return "fsdfs", nil
}
