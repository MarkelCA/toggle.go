package flags

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var ctx = context.Background()

type FlagRepository interface {
    Get(key string) (bool, error)
    Keys() ([]string, error)
    Exists(name string) (bool, error)
    Set(key string, value interface{}) error
    List() ([]Flag, error)
}

type FlagMongoRepository struct {
    collection *mongo.Collection
}

func NewFlagMongoRepository(host string, port int) (FlagRepository, error) {
    hostStr := fmt.Sprintf("mongodb://%v:%v",host,port)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostStr))
    if err != nil {
        return nil,err
    }
    collection := client.Database("toggles").Collection("flags")
    return FlagMongoRepository{collection},nil
}

func (repository FlagMongoRepository) List() ([]Flag, error) {
    cur,err := repository.collection.Find(ctx,bson.D{})
    if err != nil { 
        return nil,err
    }
    defer cur.Close(ctx)
    result := make([]Flag,0)
    for cur.Next(ctx) {
        var flag Flag
        err := cur.Decode(&flag)
        if err != nil { 
            return nil,err
        }
        result = append(result,flag)
    }
    if err := cur.Err(); err != nil {
        return nil,err
    }
    return result,nil
}

func(repository FlagMongoRepository) Keys() ([]string, error) {
    return nil,nil
}

func(repository FlagMongoRepository) Get(key string) (bool, error) {
    x := repository.collection.FindOne(ctx,bson.D{{Key:"name",Value:key}})
    var f Flag
    err := x.Decode(&f)
    if err == mongo.ErrNoDocuments {
        return false,FlagNotFoundError
    } else if err != nil {
        return false,err
    }
    return f.Value,nil
}

func (repository FlagMongoRepository) Set(key string, value interface{}) error {
    filter := bson.D{{"name",key}}
    update := bson.D{{"$set", bson.D{{"value",value}}}}
    opts := options.Update().SetUpsert(true)
    _, err := repository.collection.UpdateOne(context.TODO(), filter, update, opts)

    if err != nil {
        return err
    } else {
        return nil
    }
}

func (repository FlagMongoRepository) Exists(name string) (bool,error) {
    _,err := repository.Get(name)
    log.Printf("aaaa ver: %v",err)
    if err == nil {
        return true,nil
    } else if err == FlagNotFoundError {
        return false,nil
    } else {
        return false,err
    }
}
