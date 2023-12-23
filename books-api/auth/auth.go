// this file will handle login and signup processes
// it will be complemented by the jwt.go file in the same package
// passwords will be hashed

// while logging in the user shall be provided with a auth token
// jwt.go will have a middleware to take care of that
// another middleware in jwt.go should be for validating the token
package auth

import (
	"errors"
	Controllers "example/books-api/controllers"
	DB "example/books-api/db"
	Helper "example/books-api/helpers"
	Structs "example/books-api/types"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// log in route
	// username and password is passed
	// GetUserByUsername username and id are unique fields so...
	// if all ok, jwt token is gerated and sent GetToken(user User)
	// else error message is sent
	type LoginInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginInp LoginInput
	if err := c.BindJSON(&loginInp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": err.Error()}})
		return
	}

	var credential string
	if loginInp.Username == "" {
		credential = loginInp.Email
	} else if loginInp.Email == "" {
		credential = loginInp.Username
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": "invalid credentials", "type": "program created"}})
		return
	}
	type User = Structs.User

	user, err := DB.GetUserByUsername(credential)
	if err != nil {
		user, err = DB.GetUserByEmail(credential)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": "invalid credentials", "type": "program created"}})
			return
		}
	}
	errorInPasswordCheck := Helper.CheckPasswordHash(loginInp.Password, user.Password)
	if errorInPasswordCheck == nil {
		token, err := GetToken(user)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"type": "error", "error": gin.H{"message": err.Error()}})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"type":    "success",
			"message": "logged in sucessfully.",
			"token":   token,
		})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": "invalid credentials", "type": "program created"}})
	}
}

func getValidId() string {
	for {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(10000) + 1
		randomeNumber := strconv.Itoa(randomNumber)
		if _, err := DB.GetUser(randomeNumber); err == nil {
			continue
		}
		return randomeNumber
	}
}

func SignUp(c *gin.Context) {
	// sign up route
	// edit create user route
	type User = Structs.User

	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": err.Error()}})
		return
	}

	if newUser.Email == "" || newUser.Password == "" || newUser.Username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": "paramter(s) are missing", "type": "user generated"}})
		return
	}

	if _, err := DB.GetUserByUsername(newUser.Username); err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": errors.New("user with this username already exists").Error(), "type": "program created"}})
		return
	}

	if _, err := DB.GetUserByEmail(newUser.Email); err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "error": gin.H{"message": errors.New("user with this email already exists").Error(), "type": "program created"}})
		return
	}

	newUser.ID = getValidId()
	newUser.Access = "user"
	var tempSlice []string
	newUser.Issued = tempSlice

	if temp, err := Controllers.CreateUser(newUser); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"type": "error", "error": gin.H{"message": err.Error()}})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"type": "success",
			"user-details": gin.H{
				"id":       temp.ID,
				"username": temp.Username,
				"email":    temp.Email,
				"password": "[PROTECTED]",
			},
			"message": "user created",
		})
	}

}
