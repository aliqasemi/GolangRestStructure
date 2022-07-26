package services

import (
	models "basical-app/models"
	"basical-app/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInterface interface {
	Index() ([]models.User, error)
	Show(id string) (models.User, error)
	Create(user models.User) (*mongo.InsertOneResult, error)
	Update(user models.User, id string) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type userService struct {
}

func UserServiceBuilder() UserInterface {
	return userService{}
}

func (userService) Index() ([]models.User, error) {
	users, _ := repository.UserRepositoryBuilder().IndexUser()
	return users, nil
}

func (userService) Show(id string) (models.User, error) {
	user, _ := repository.UserRepositoryBuilder().ShowUser(id)
	return user, nil
}

func (userService) Create(user models.User) (*mongo.InsertOneResult, error) {
	userCreate, _ := repository.UserRepositoryBuilder().CreateUser(user)
	return userCreate, nil
}

func (userService) Update(user models.User, id string) (*mongo.UpdateResult, error) {
	userCreate, _ := repository.UserRepositoryBuilder().UpdateUser(user, id)
	return userCreate, nil
}

func (userService) Delete(id string) (*mongo.DeleteResult, error) {
	userCreate, _ := repository.UserRepositoryBuilder().DeleteUser(id)
	return userCreate, nil
}
