package storage

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var ctx context.Context
// var cancel context.CancelFunc
//
// ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()

type MongoClient struct {
    client *mongo.Client
}

func NewMongoClient(host string, port int) (KeyValueDBClient, error) {
    hostStr := fmt.Sprintf("mongodb://%v:%v",host,port)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostStr))
    if err != nil {
        return nil,err
    }
    return MongoClient{client},nil
}

func(r MongoClient) Keys() ([]string, error) {
    return nil,nil
}

func(r MongoClient) Get(key string) (string, error) {
    collection := r.client.Database("toggles").Collection("flags")
    x := collection.FindOne(ctx,bson.D{})
    y,_ := x.Raw()
    fmt.Printf("strr -> %v",y.String())
    return "",nil
}

func (r MongoClient) Set(key string, value interface{}) error {
    return nil
}

func (r MongoClient) Exists(name string) (bool,error) {
    return false,nil
}
