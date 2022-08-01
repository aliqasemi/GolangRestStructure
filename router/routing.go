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

	authApi.GET("/users", controllers.Index, middleware.Authorize("user"))
	authApi.GET("/users/:id", controllers.Show, middleware.Authorize("user"))
	authApi.POST("/users", controllers.Create, middleware.Authorize("admin"))
	authApi.PUT("/users/:id", controllers.Update, middleware.Authorize("admin"))
	authApi.DELETE("/users/:id", controllers.Delete, middleware.Authorize("admin"))
	authApi.GET("/auth", controllers.Auth, middleware.Authorize("user"))

	return nil
}
