package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type UserRepository interface {
	FindAll() ([]*User, error)
	FindByUserName(userName string) (*User, error)
	Create(user User) error
	Update(user *User) error
	Upsert(user User) error
	GetPermissions(userName string) ([]string, error)
	AddPermission(userName, permission string) error
	RemovePermission(userName, permission string) error
}

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(host string, port uint) (UserRepository, error) {
	hostStr := fmt.Sprintf("mongodb://%v:%v", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostStr))
	if err != nil {
		return nil, err
	}
	collection := client.Database("toggles").Collection("users")
	return UserMongoRepository{collection}, nil
}

func (repository UserMongoRepository) Upsert(user User) error {
	_, err := repository.collection.ReplaceOne(ctx, bson.D{{Key: "username", Value: user.UserName}}, user, options.Replace().SetUpsert(true))
	return err
}

func (repository UserMongoRepository) FindAll() ([]*User, error) {
	cur, err := repository.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	result := make([]*User, 0)
	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, &user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (repository UserMongoRepository) FindByUserName(userName string) (*User, error) {
	x := repository.collection.FindOne(ctx, bson.D{{Key: "username", Value: userName}})
	var user User
	err := x.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository UserMongoRepository) Create(user User) error {
	_, err := repository.collection.InsertOne(ctx, user)
	return err
}

func (repository UserMongoRepository) Update(user *User) error {
	_, err := repository.collection.ReplaceOne(ctx, bson.D{{Key: "username", Value: user.UserName}}, user)
	return err
}

func (repository UserMongoRepository) GetPermissions(userName string) ([]string, error) {
	opts := options.FindOne().SetProjection(bson.D{{Key: "permissions", Value: 1}})
	permissions := repository.collection.FindOne(ctx, bson.D{{Key: "username", Value: userName}}, opts)

	var u User
	err := permissions.Decode(&u)
	if err != nil {
		return nil, err
	}

	return u.Permissions, nil
}

func (repository UserMongoRepository) AddPermission(userName, permission string) error {
	_, err := repository.collection.UpdateOne(ctx, bson.D{{Key: "username", Value: userName}}, bson.D{{Key: "$push", Value: bson.D{{Key: "permissions", Value: permission}}}})
	return err
}

func (repository UserMongoRepository) RemovePermission(userName, permission string) error {
	_, err := repository.collection.UpdateOne(ctx, bson.D{{Key: "username", Value: userName}}, bson.D{{Key: "$pull", Value: bson.D{{Key: "permissions", Value: permission}}}})
	return err
}
