package main

import (
	"errors"
	Config "example/mypackage/config"
	Structs "example/mypackage/types"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type User = Structs.User

func GenerateJWT(user User) (string, error) {
	var secret = Config.JWTSecret()
	token := jwt.New(jwt.SigningMethodEd25519)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(2 * time.Hour)
	claims["authorized"] = true
	claims["userid"] = user.ID
	claims["access"] = user.Access

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	var secret = Config.JWTSecret()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	var user User
	user.ID = "1"
	user.Username = "rajat"
	user.Access = "admin"
	user.Email = "rsingh1734@gmail.com"
	user.Issued = make([]string, 0)
	user.Password = "user123"

	if token, err := GenerateJWT(user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token)
	}
}
