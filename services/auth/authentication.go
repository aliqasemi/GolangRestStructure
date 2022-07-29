package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"user-name"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

type AuthContext struct {
	echo.Context
}

func GenerateJWT(email string, username string, id string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}

func (c *AuthContext) GetUserId() string {
	t := c.Get("user")
	fmt.Println(t)
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*JWTClaim)
	return claim.ID
}
