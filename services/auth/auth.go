package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

var JwtKey = []byte("secret")

type JWTClaim struct {
	Username string `json:"user-name"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type AuthContext struct {
	echo.Context
}

type AuthInterface interface {
	GetUserId() string
	HasRole(role string) bool
}

func AuthBuilder(c echo.Context) AuthInterface {
	return c.(*AuthContext)
}

func GenerateJWT(email string, username string, id string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		ID:       id,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JwtKey)
	return
}

func (c *AuthContext) GetUserId() string {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*JWTClaim)
	return claim.ID
}

func (c *AuthContext) HasRole(role string) bool {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*JWTClaim)
	if claim.Role == role || claim.Role == "admin" || claim.Role == "superAdmin" {
		return true
	} else {
		return false
	}
}
