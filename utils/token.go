package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const JWT_SECRET = "itssecret"

func GenerateJWTToken(userId string) (string, error) {
	// Create a new token object, specifying signing method and the claims

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()
	claims["id"] = userId

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	return tokenString, err
}

func VerifyJWTToken(tokenString string) (interface{}, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"], nil
	} else {
		return nil, err
	}

}
