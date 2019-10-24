package dataBase

import (
	"backend/sensitive"
	"backend/util"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect() util.Err {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+sensitive.MONGODB_ADDR))
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+sensitive.MONGODB_ADDR).SetAuth(options.Credential{
		AuthSource: "admin",
		Username:   "compass",
		Password:   "compass",
	}))

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return util.NewBasicError("can not connect to mongodb", 0)
	}
	collection := client.Database("easypost").Collection("settings")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return util.NewBasicError("mongodb find error", 404)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func Start() util.Err {
	return Connect()
}
