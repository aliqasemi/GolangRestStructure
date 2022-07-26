package repository

import (
	"basical-app/database"
	models "basical-app/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	IndexUser() ([]models.User, error)
	ShowUser(id string) (models.User, error)
	CreateUser(user models.User) (*mongo.InsertOneResult, error)
	UpdateUser(user models.User, id string) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
}

type userRepository struct {
	db database.MongoDb
}

func UserRepositoryBuilder() UserRepositoryInterface {
	return userRepository{db: database.GetClient()}
}

func (user userRepository) IndexUser() ([]models.User, error) {
	userCollection := user.db.GetModelCollection("users")
	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (user userRepository) ShowUser(id string) (models.User, error) {
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	userCollection := user.db.GetModelCollection("users")
	var userFind models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", primitiveId}}).Decode(&userFind)
	if err != nil {
		return models.User{}, err
	}
	return userFind, nil
}

func (user userRepository) CreateUser(userData models.User) (*mongo.InsertOneResult, error) {
	userCollection := user.db.GetModelCollection("users")
	userCreated, err := userCollection.InsertOne(context.TODO(), userData)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

func (user userRepository) UpdateUser(userData models.User, id string) (*mongo.UpdateResult, error) {
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	userCollection := user.db.GetModelCollection("users")
	userUpdated, err := userCollection.UpdateOne(context.TODO(), bson.D{{"_id", primitiveId}}, bson.D{{"$set", userData}})
	if err != nil {
		return nil, err
	}
	return userUpdated, nil
}

func (user userRepository) DeleteUser(id string) (*mongo.DeleteResult, error) {
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	userCollection := user.db.GetModelCollection("users")
	userCreated, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", primitiveId}})
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}
