package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret string = "itssecret"

func GenerateJWTToken(userId string) string {
	// Create a new token object, specifying signing method and the claims

	claims := &jwt.RegisteredClaims{
		Issuer:    "Harsh",
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
		ExpiresAt: jwt.NewNumericDate(time.Now()),
		Subject:   userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Println(err)
	}

	return tokenString
}

func VerifyJWTToken(tokenString string) (interface{}, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"], nil
	} else {
		return nil, err
	}

}
