package controllers

import (
	"basical-app/models"
	"basical-app/repository"
	"basical-app/validation"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Show(c echo.Context) error {
	userService := repository.UserRepositoryBuilder()
	user, err := userService.Show(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func Index(c echo.Context) error {
	userService := repository.UserRepositoryBuilder()
	users, err := userService.Index()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	userService := repository.UserRepositoryBuilder()
	userInput := new(validation.UserInput)
	err := c.Bind(userInput)
	if err != nil {
		log.Fatal(err)
	}
	userModel, err := userInput.ValidateAndBuildModel()
	if err != nil {
		return c.JSON(http.StatusFound, err.Error())
	}
	users, err := userService.Create(userModel)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, users)
}

func Update(c echo.Context) error {
	userService := repository.UserRepositoryBuilder()
	users, err := userService.Update(models.User{Age: 25}, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, users)
}

func Delete(c echo.Context) error {
	userService := repository.UserRepositoryBuilder()
	users, err := userService.Delete(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, users)
}
