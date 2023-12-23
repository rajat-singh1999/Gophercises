// this file will hold all the functions required to perform crud on user data
// CreateUser() -> create a new user
// DeleteUser() -> delete a user
// UserChangeAccess() -> toggle the access for the user, by default a user will not have an admin access
// UpdateUser() -> updates user info, access and id cant be updated
// ListUsers() -> shows info on all users

package Controllers

import (
	DB "example/books-api/db"
	Help "example/books-api/helpers"
	Structs "example/books-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User = Structs.User

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := DB.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"type": "error", "Error": gin.H{
			"message": err.Error(),
			"type":    "program created",
		}})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"type": "success",
		"user-details": gin.H{
			"id":       user.ID,
			"access":   user.Access,
			"username": user.Username,
			"email":    user.Email,
			"issued":   user.Issued,
		},
	})
}

func GetUsers(c *gin.Context) {
	users, err := DB.LoadUsersFromFile()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"type": "error", "Error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"type": "success", "users": users})
}

func CreateUser(newUser User) (User, error) {
	// load all users in a variable, then append

	if users, err := DB.LoadUsersFromFile(); err != nil {
		return newUser, err
	} else {
		// hash the password
		hashedPass, err := Help.CreatePasswordHash(newUser.Password)
		if err != nil {
			return newUser, err
		}
		newUser.Password = hashedPass
		users = append(users, newUser)
		if err := DB.SaveUsersToFile(users); err != nil {
			return newUser, err
		} else {
			return newUser, nil
		}
	}
}

func ChangeUserAuth(c *gin.Context) {
	toAccess, ok := c.GetQuery("access")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": gin.H{
				"message": "access parameter not given",
				"type":    "program created",
			},
		})
		return
	}
	id := c.Param("id")
	user, err := DB.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":           err.Error(),
			"program message": "user does not exist",
		})
		return
	}

	if toAccess != "user" && toAccess != "admin" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": gin.H{
				"message":               "incorrect access query received",
				"Access-Query-Received": toAccess,
				"type":                  "Program created",
			},
		})

	} else {
		user.Access = toAccess
		// delete user
		err := DB.DeleteUser(user.ID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		// load all users and append
		users, err := DB.LoadUsersFromFile()
		users = append(users, user)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		// save appended slice
		err = DB.SaveUsersToFile(users)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		// positive response
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "changed auth successfully",
			"user-details": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"access":   user.Access,
			},
		})
	}
}
