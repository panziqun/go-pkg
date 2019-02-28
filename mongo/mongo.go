package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func Setup() (err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		fmt.Println("new client err:%v", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	Database = client.Database("youpin")
	if err != nil {
		fmt.Println("client connect err:%v", err)
		return err
	}
	return err
}

func InsertOne(name string, data bson.M) (id interface{}, err error) {
	res, err := Database.Collection(name).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func One(name string, filter bson.D) (model interface{}, err error) {
	err = Database.Collection(name).FindOne(context.Background(), filter).Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func All(name string, filter bson.D) (list []interface{}, err error) {
	cur, err := Database.Collection(name).Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	count, _ := Database.Collection(name).CountDocuments(context.Background(), filter)
	list = make([]interface{}, count)
	for cur.Next(context.Background()) {
		list = append(list, cur.Current)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return list, err
}
