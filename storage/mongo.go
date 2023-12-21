package storage

import (
	"fmt"
    "time"
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

func NewMongoClient(host string, port int) (KeyValueStore, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27018"))
    if err != nil {
        return nil,err
    }
    fmt.Printf("mongo connection: %v,%v",client,err)
    return MongoClient{client},nil
}

func(r MongoClient) Keys() ([]string, error) {
    return nil,nil
}

func(r MongoClient) Get(key string) (string, error) {
    return "",nil
}

func (r MongoClient) Set(key string, value interface{}, expiration time.Duration) error {
    return nil
}

func (r MongoClient) Exists(name string) (bool,error) {
    return false,nil
}
