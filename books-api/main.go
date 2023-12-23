package main

import (
	Auth "example/books-api/auth"
	Controller "example/books-api/controllers"
	MiddleWare "example/books-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.POST("/signup", Auth.SignUp)
	router.POST("/login", Auth.Login)
	router.Use(MiddleWare.CheckUserAuth())
	{
		router.GET("/books", Controller.GetBooks)    // http://localhost:8080/book
		router.GET("/book/:id", Controller.BookById) // http://localhost:8080/book/3
		router.GET("/user/:id", Controller.GetUserByID)
	}
	router.Use(MiddleWare.CheckAdminAuth())
	{
		router.GET("/users", Controller.GetUsers)

		router.GET("/return/:uid", Controller.ReturnBook)              // http://localhost:8080/return/uid?bid=3
		router.GET("/addexistingbook/:id", Controller.AddExistingBook) // http://localhost:8080/addexistingbook/id?c=4
		router.POST("/addbook", Controller.AddBook)                    // http://localhost:8080/addbook
		router.POST("/update", Controller.UpdateBook)                  // http://localhost:8080/update
		router.PATCH("/checkout/:uid", Controller.CheckOut)            // http://localhost:8080/checkout/uid?id=3
		router.DELETE("/delete/:id", Controller.DeleteBook)

		router.GET("/changeuserauth/:id", Controller.ChangeUserAuth) // http://localhost:8080/changeuserauth/:id?access=2
	}
	router.Run("localhost:8080")
}
