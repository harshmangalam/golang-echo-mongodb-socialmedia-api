package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const JWT_SECRET = "itssecret"

func GenerateJWTToken(userId string) (string, error) {
	fmt.Println(userId)
	// Create a new token object, specifying signing method and the claims

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()
	claims["id"] = userId

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	return tokenString, err
}

func ParseTokenFunc(tokenString string, c echo.Context) (interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string), nil
}
