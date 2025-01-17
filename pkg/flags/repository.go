package flags

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type FlagRepository interface {
	Get(key string) (bool, error)
	Exists(name string) (bool, error)
	Set(key string, value any) error
	List() ([]Flag, error)
	Delete(name string) error
}

type FlagMongoRepository struct {
	collection *mongo.Collection
}

func NewFlagMongoRepository(host string, port uint) (FlagRepository, error) {
	hostStr := fmt.Sprintf("mongodb://%v:%v", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostStr))
	if err != nil {
		return nil, err
	}
	collection := client.Database("toggles").Collection("flags")
	return FlagMongoRepository{collection}, nil
}

func (repository FlagMongoRepository) List() ([]Flag, error) {
	cur, err := repository.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	result := make([]Flag, 0)
	for cur.Next(ctx) {
		var flag Flag
		err := cur.Decode(&flag)
		if err != nil {
			return nil, err
		}
		result = append(result, flag)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (repository FlagMongoRepository) Get(key string) (bool, error) {
	x := repository.collection.FindOne(ctx, bson.D{{Key: "name", Value: key}})
	var f Flag
	err := x.Decode(&f)
	if err == mongo.ErrNoDocuments {
		return false, ErrFlagNotFound
	} else if err != nil {
		return false, err
	}
	return f.Value, nil
}

func (repository FlagMongoRepository) Set(key string, value any) error {
	filter := bson.D{{"name", key}}
	update := bson.D{{"$set", bson.D{{"value", value}}}}
	opts := options.Update().SetUpsert(true)
	_, err := repository.collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (repository FlagMongoRepository) Exists(name string) (bool, error) {
	_, err := repository.Get(name)
	if err == nil {
		return true, nil
	} else if err == ErrFlagNotFound {
		return false, nil
	} else {
		return false, err
	}
}

func (repository FlagMongoRepository) Delete(name string) error {
	_, err := repository.collection.DeleteOne(ctx, bson.D{{Key: "name", Value: name}})
	if err != nil {
		return err
	}
	return nil
}
