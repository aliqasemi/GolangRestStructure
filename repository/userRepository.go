package repository

import (
	"basical-app/database"
	models "basical-app/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	Index() ([]models.User, error)
	Show(id string) (models.User, error)
	Create(user models.User) (*mongo.InsertOneResult, error)
	Update(user models.User, id string) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	FindUserByPhoneForLogin(email string) (models.User, error)
}

type userRepository struct {
	db database.MongoDb
}

func UserRepositoryBuilder() UserRepositoryInterface {
	return userRepository{db: database.GetClient()}
}

func (user userRepository) Index() ([]models.User, error) {
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

func (user userRepository) Show(id string) (models.User, error) {
	userCollection := user.db.GetModelCollection("users")
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	var userFind models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", primitiveId}}).Decode(&userFind)
	if err != nil {
		return models.User{}, err
	}
	return userFind, nil
}

func (user userRepository) Create(userData models.User) (*mongo.InsertOneResult, error) {
	userCollection := user.db.GetModelCollection("users")
	userCreated, err := userCollection.InsertOne(context.TODO(), userData)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

func (user userRepository) Update(userData models.User, id string) (*mongo.UpdateResult, error) {
	userCollection := user.db.GetModelCollection("users")
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	userUpdated, err := userCollection.UpdateOne(context.TODO(), bson.D{{"_id", primitiveId}}, bson.D{{"$set", userData}})
	if err != nil {
		return nil, err
	}
	return userUpdated, nil
}

func (user userRepository) Delete(id string) (*mongo.DeleteResult, error) {
	userCollection := user.db.GetModelCollection("users")
	primitiveId, _ := primitive.ObjectIDFromHex(id)
	userCreated, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", primitiveId}})
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

func (user userRepository) FindUserByPhoneForLogin(phoneNumber string) (models.User, error) {
	userCollection := user.db.GetModelCollection("users")
	var userFind models.User
	err := userCollection.FindOne(context.TODO(), bson.D{{"phone-number", phoneNumber}}).Decode(&userFind)
	if err != nil {
		return models.User{}, errors.New("user with this Phone Number doesn't exist")
	} else {
		return userFind, nil
	}
}
