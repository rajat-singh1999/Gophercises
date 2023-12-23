package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID   string
	Email    string
	Password string
	Auth     string
}

func getToken(user User) (string, error) {
	// Generate token using user details
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = user.UserID
	claims["email"] = user.Email
	claims["auth"] = user.Auth

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkToken(tokenString string) (string, string, error) {
	// Verify and decode the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	userID := claims["userid"].(string)
	auth := claims["auth"].(string)

	return userID, auth, nil
}

func main() {
	// Pass user details
	user := User{
		UserID:   "123",
		Email:    "example@example.com",
		Password: "password123",
		Auth:     "user",
	}

	// Get token
	token, err := getToken(user)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Token:", token)

	// Check token
	userID, auth, err := checkToken(token)
	if err != nil {
		fmt.Println("Error checking token:", err)
		return
	}

	fmt.Println("UserID:", userID)
	fmt.Println("Auth:", auth)
}

// package main

// import (
// 	"fmt"
// )

// func getToken() {
// 	// pass details and jwt get token and validity
// 	// use only userid, email, password and an auth field to genrate token
// }

// func checkToken() {
// 	// decode token and tell the userid and auth field of the user
// 	// also perform some check on the decoded details with the original details
// 	// decode token and send back the decoded details
// }

// func main() {
// 	// pass userid, username, emailid, password, auth(user, admin)
// 	fmt.Println("Working.")
// }
