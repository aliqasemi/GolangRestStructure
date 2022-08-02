package middleware

import (
	"basical-app/services/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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

func Authorize(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authUser := auth.AuthBuilder(c)
			if role == "superAdmin" {
				return next(c)
			} else if authUser.HasRole(role) {
				return next(c)
			} else {
				return &echo.HTTPError{
					Message: "Unauthorized",
					Code:    http.StatusForbidden,
				}
			}
		}
	}
}
