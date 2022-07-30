package router

import (
	"basical-app/controllers"
	"basical-app/middleware"
	"github.com/labstack/echo/v4"
)

func SetRout(e *echo.Echo) error {

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	authApi := e.Group("", middleware.AuthApi)
	authApi = authApi.Group("", middleware.ValidateToken())

	authApi.GET("/users", controllers.Index)
	authApi.GET("/users/:id", controllers.Show)
	authApi.POST("/users", controllers.Create)
	authApi.PUT("/users/:id", controllers.Update)
	authApi.DELETE("/users/:id", controllers.Delete)
	authApi.GET("/auth", controllers.Auth)

	return nil
}
