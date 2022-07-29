package controllers

import (
	"basical-app/models"
	"basical-app/repository"
	"basical-app/services/auth"
	"basical-app/validation/user"
	"errors"
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := userService.Create(userModel)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, user)
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

func Register(c echo.Context) error {
	userInput := new(validation.UserInput)
	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userModel, err := userInput.ValidateAndBuildModel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&userModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := userModel.HashPassword(userModel.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userRepository := repository.UserRepositoryBuilder()
	user, err := userRepository.Create(userModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	loginInput := new(validation.LoginInput)
	if err := c.Bind(loginInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userModel, err := loginInput.ValidateAndBuildModel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRepository := repository.UserRepositoryBuilder()
	userData, err := userRepository.FindUserByEmailForLogin(userModel.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = userData.CheckPassword(loginInput.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errors.New("invalid password"))
	}

	tokenString, err := auth.GenerateJWT(userData.Email, userData.UserName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userData.Password = ""
	result := struct {
		BearerToken string      `json:"token,omitempty"`
		UserInfo    models.User `json:"user,omitempty"`
	}{
		BearerToken: "Bearer " + tokenString,
		UserInfo:    userData,
	}

	return c.JSON(http.StatusOK, &result)
}
