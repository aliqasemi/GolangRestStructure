package router

import (
	"basical-app/controllers"
	"github.com/labstack/echo/v4"
)

func SetRout(e *echo.Echo) error {
	e.GET("/users", controllers.Index)
	e.GET("/users/:id", controllers.Show)
	e.POST("/users", controllers.Create)
	e.PUT("/users/:id", controllers.Update)
	e.DELETE("/users/:id", controllers.Delete)
	return nil
}
