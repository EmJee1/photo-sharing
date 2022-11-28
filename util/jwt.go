package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

// TODO: extract into .env file
var jwtKey = []byte("secret-key")

func GenerateJWT(email string) (time.Time, string, error) {
	expiresAt := time.Now().Add(time.Hour)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return expiresAt, tokenString, err
}

func ParseJWT(tokenString string) (string, error) {
	claims := &Claims{}
	var err error
	token, pErr := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if pErr != nil || !token.Valid {
		err = errors.New("Invalid JWT token")
	}

	return claims.Email, err
}
