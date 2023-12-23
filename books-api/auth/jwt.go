package auth

import (
	Config "example/books-api/config"
	DB "example/books-api/db"
	Structs "example/books-api/types"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type User = Structs.User

func GetToken(user *User) (string, error) {
	// Generate token using user details
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	// claims["email"] = user.Email
	claims["username"] = user.Username
	claims["access"] = user.Access

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(Config.JWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(tokenString string) (string, string, error) {
	// Verify and decode the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JWTSecret()), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	userID := claims["id"].(string)
	auth := claims["access"].(string)

	_, err = DB.GetUser(userID)
	if err != nil {
		return "", "", err
	}
	return userID, auth, nil
}
