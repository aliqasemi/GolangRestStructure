package middleware

import (
	"basical-app/services/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthApi(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header["Authorization"]
		if len(c.Request().Header["Authorization"]) < 0 {
			return c.JSON(http.StatusForbidden, "You don't have token")
		}
		err := auth.ValidateToken(tokenString[0][7:])
		if err != nil {
			return c.JSON(http.StatusForbidden, "You are not authorized!")
		}
		authContext := &auth.AuthContext{Context: c}
		return next(authContext)
	}
}
