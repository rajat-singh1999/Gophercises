// for the following routes there should be a checkAuth middleware
// /books, /book/:id -> admin, user
// /return, /addboon, /update, /checkout, /delete/:id -> admin
// for this there should be two different middleware functions i think
// checkUserAuth, checkAdminAuth -> these two can be in the same file

package middlewares

import (
	Auth "example/books-api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check user authentication here
		// If authenticated, call the next handler
		// If not, respond with an error or redirect
		authHeader := c.GetHeader("Authorization")
		_, access, err := Auth.CheckToken(authHeader)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "tokenerror", "message": err.Error()})
			c.Abort()
			return
		}

		if access == "user" || access == "admin" {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusForbidden, gin.H{"type": "tokenerror", "message": "unauthorized"})
			c.Abort()
			return
		}
	}
}

func CheckAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin authentication here
		// If authenticated, call the next handler
		// If not, respond with an error or redirect
		authHeader := c.GetHeader("Authorization")
		_, access, err := Auth.CheckToken(authHeader)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "tokenerror", "message": err.Error()})
			c.Abort()
			return
		}

		if access == "admin" {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusForbidden, gin.H{"type": "tokenerror", "message": "unauthorized"})
			c.Abort()
			return
		}
	}
}
