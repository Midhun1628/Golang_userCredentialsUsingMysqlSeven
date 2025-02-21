package utils

import (
	"time"
"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("thisisnotakey") // Change this to a secure key

// GenerateToken creates a JWT token for a user
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Second *20).Unix(), // Token expires in 1 minute
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken checks if the JWT token is valid
func ValidateToken(tokenString string) (jwt.MapClaims, bool) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Token validation failed:",err)
		return nil, false
	}

	return claims, true
}