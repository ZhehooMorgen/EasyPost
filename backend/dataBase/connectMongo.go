package dataBase

import (
	"backend/sensitive"
	"backend/util"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func Connect(collectionName string) (*mongo.Collection, util.Err) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+sensitive.MONGODB_ADDR).SetAuth(options.Credential{
		AuthSource: sensitive.MONGOBD_DBNAME,
		Username:   sensitive.MONGOBD_USERNAME,
		Password:   sensitive.MONGOBD_PASSWORD,
	}))

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, util.NewBasicError("can not connect to mongodb", 0, err)
	}
	collection := client.Database(sensitive.MONGOBD_DBNAME).Collection(collectionName)
	_, err = collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, util.NewBasicError("can not exec CURD on specific DB collection", 500, err)
	}
	return collection, nil
}
