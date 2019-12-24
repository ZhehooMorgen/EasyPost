package dataBase

import (
	"backend/resourceScheduler"
	"backend/resourceScheduler/schs"
	"backend/sensitive"
	"backend/util"
	"context"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
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

type DBCollection struct {
	*mongo.Collection
	id         uuid.UUID
}

func (c *DBCollection) Definition() (t resourceScheduler.Type, id uuid.UUID) {
	return resourceScheduler.Type(reflect.TypeOf(c).String()), id
}

var (
	Accounts *DBCollection
	Settings *DBCollection
)

func connectAndReg(collName string, collR **DBCollection) util.Err {
	if coll, err := Connect(collName); err != nil {
		return err
	} else {
		*collR = &DBCollection{
			coll,
			 uuid.NewV4(),
		}
		_ = schs.Scheduler.RegRes(*collR)
	}
	return nil
}

func ConnectAndGetAllCollections() util.Err {
	if err := connectAndReg(sensitive.COLLECTION_NAME_ACCOUNTS, &Accounts); err != nil {
		return err
	}
	if err := connectAndReg(sensitive.COLLECTION_NAME_SETTINGS, &Settings); err != nil {
		return err
	}
	return nil
}
