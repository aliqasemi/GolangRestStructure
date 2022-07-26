package controllers

import (
	"basical-app/models"
	"basical-app/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Show(c echo.Context) error {
	userService := services.UserServiceBuilder()
	user, err := userService.Show(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func Index(c echo.Context) error {
	userService := services.UserServiceBuilder()
	users, err := userService.Index()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	userService := services.UserServiceBuilder()
	users, err := userService.Create(models.User{FirstName: "matin", LastName: "shahri", Age: 24})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, users)
}

func Update(c echo.Context) error {
	userService := services.UserServiceBuilder()
	users, err := userService.Update(models.User{Age: 25}, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, users)
}

func Delete(c echo.Context) error {
	userService := services.UserServiceBuilder()
	users, err := userService.Delete(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, users)
}
