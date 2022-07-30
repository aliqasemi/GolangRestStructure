package middleware

import (
	"basical-app/services/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthApi(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authContext := &auth.AuthContext{Context: c}
		return next(authContext)
	}
}

func ValidateToken() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		SigningKey: auth.JwtKey,
		Claims:     &auth.JWTClaim{},
	}
	return middleware.JWTWithConfig(jwtConfig)
}
