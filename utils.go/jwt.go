package utils

import (
	"time"
    "fmt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("thisisnotakey") 

// GenerateToken creates a JWT token for a user
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{ //creates a map to store key value pairs ,here we store username and exp
"username": username,
"exp"     : time.Now().Add(time.Minute *1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)//create a new token using the signmethod and attaches a claims
	return token.SignedString(jwtSecret) //sign the token with secret key to ensure integrity
}

// ValidateToken checks if the JWT token is valid
func ValidateToken(tokenString string) (jwt.MapClaims, bool) { // a JWT token string as input and returns the decoded claims and a boolean indicating if the token is valid.
	claims := jwt.MapClaims{} //Creates an empty MapClaims to store decoded JWT claims.

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { //Parses the JWT token, extracts claims, and verifies it using the secret key.
		return jwtSecret, nil  //Provides the secret key (jwtSecret) needed to verify the token.
	})

	if err != nil || !token.Valid {
		fmt.Println("Token validation failed:",err)
		return nil, false
	}

	return claims, true  // If the token is valid, it returns the decoded claims and true.
}