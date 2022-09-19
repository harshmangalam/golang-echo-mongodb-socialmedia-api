package utils

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
)

const JWT_SECRET = "itssecret"

func GenerateJWTToken(userId string) string {
	// Create a new token object, specifying signing method and the claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(JWT_SECRET)

	if err != nil {
		log.Println(err)
	}

	return tokenString
}
