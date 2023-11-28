package main

import (
	Controller "example/books-api/controllers"
	Middleware "example/books-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", Controller.GetBooks)                          // http://localhost:8080/book
	router.GET("/book/:id", Controller.BookById)                       // http://localhost:8080/book/3
	router.GET("/return", Controller.ReturnBook)                       // http://localhost:8080/return?id=3
	router.POST("/addbook", Middleware.TestMiddle, Controller.AddBook) // http://localhost:8080/addbook
	router.POST("/update", Controller.UpdateBook)                      // http://localhost:8080/update
	router.PATCH("/checkout", Controller.CheckOut)                     //http://localhost:8080/checkout?id=3
	router.DELETE("/delete/:id", Controller.DeleteBook)
	router.Run("localhost:8080")
}
