package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/laughmaker/go-pkg/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Mongo struct {
	DB *mongo.Database
}

func Setup() {
	var uri string
	if conf.Mongodb.User != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Mongodb.User, conf.Mongodb.Password, conf.Mongodb.Host, conf.Mongodb.Port)
	} else {
		uri = fmt.Sprintf("mongodb://@%s:%d", conf.Mongodb.Host, conf.Mongodb.Port)
	}
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("new client err:%v", err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("client connect err:%v", err)
		panic(err)
	}
}

func GetDB(dbname string) *mongo.Database {
	return client.Database(dbname)
}

func DefaultDB() *mongo.Database {
	return client.Database(conf.Mongodb.Name)
}

func (m *Mongo) InsertOne(name string, data interface{}) (id interface{}, err error) {
	res, err := m.DB.Collection(name).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (m *Mongo) One(name string, filter interface{}, args ...*options.FindOneOptions) (model interface{}, err error) {
	err = m.DB.Collection(name).FindOne(context.Background(), filter, args...).Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *Mongo) All(name string, filter interface{}, args ...*options.FindOptions) (list []interface{}, err error) {
	cur, err := m.DB.Collection(name).Find(context.Background(), filter, args...)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result bson.M
		cur.Decode(&result)
		list = append(list, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
