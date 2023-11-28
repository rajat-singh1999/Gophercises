package Controllers

import (
	"errors"
	DB "example/books-api/db"
	Helper "example/books-api/helpers"
	Structs "example/books-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book = Structs.Book

func BookById(c *gin.Context) {
	id := c.Param("id") // get id parameter

	// calling helper function that does the actual job
	book, err := Helper.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book with book id does not exist."})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book found.", "book-details": book})
}

func GetBooks(c *gin.Context) {
	// gets all books
	if books, err := DB.LoadBooksFromFile(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err, "type": "program created"})
	} else {
		c.IndentedJSON(http.StatusOK, books)
	}
}

func AddBook(c *gin.Context) {
	var newBook Book
	// trying to bind the received body data to the newBook variable
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// all the fields should be filled (business logic)
	if newBook.ID == "" || newBook.Title == "" || newBook.Author == "" || newBook.Quantity == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": errors.New("book parameters not complete").Error(), "type": "program created"}})
		return
	}

	// checking if a book with this id already exists
	if _, err := Helper.GetBook(newBook.ID); err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": errors.New("book with id already exists, try to use '/update' route").Error(), "type": "program created"}})
		return
	}

	// loading all the books from file to a go slice,
	// appending the newBook to the slice
	// saving the the new slice in the file
	if books, err := DB.LoadBooksFromFile(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err, "type": "program created"})
	} else {
		books = append(books, newBook)
		if err := DB.SaveBooksToFile(books); err != nil {
			c.IndentedJSON(http.StatusNotModified, gin.H{"message": "Book was not inserted.", "book-details": newBook, "Error": err})
		} else {
			c.IndentedJSON(http.StatusOK, newBook)
		}
	}
}

func UpdateBook(c *gin.Context) {
	var newBook Book
	// binding update parameters to the newBook variable
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error(), "type": "program created"}})
		return
	}

	// getting the old book for reference,
	// the below code also checks while getting if a book with
	// the passed id is present or not
	var oldBook, err = Helper.GetBook(newBook.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error(), "type": "program created"}})
	}

	// the code below this tries to validate the data sent in the request
	newBook.Quantity = oldBook.Quantity
	// ID field must not be empty
	if newBook.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": errors.New("id parameter not provided").Error(), "type": "program created"}})
		return
	}
	// if both title and author data is missing than it is caught here
	if newBook.Title == "" && newBook.Author == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": errors.New("you may change either title or author, but you are changing none").Error(), "type": "program created"}})
		return
	}
	if newBook.Title == "" {
		newBook.Title = oldBook.Title
	}
	if newBook.Author == "" {
		newBook.Author = oldBook.Author
	}

	// data is corrected and checked for fatal errors, now we can move for update
	if err := DB.DeleteBook(newBook.ID); err != nil {
		c.IndentedJSON(http.StatusNotModified, gin.H{"Error": gin.H{"message": err.Error(), "type": "program created"}})
	} else {
		if books, err := DB.LoadBooksFromFile(); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": gin.H{"message": err.Error(), "type": "program created"}})
		} else {
			books = append(books, newBook)
			if err := DB.SaveBooksToFile(books); err != nil {
				c.IndentedJSON(http.StatusNotModified, gin.H{"message": "Book was not inserted.", "book-details": newBook, "Error": gin.H{"message": err.Error(), "type": "program created"}})
			} else {
				c.IndentedJSON(http.StatusOK, newBook)
			}
		}
	}
}

func CheckOut(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query parameters not passed"})
		return
	}

	if err := DB.BookCheckout(id); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"Error": gin.H{"message": err.Error(), "type": "program created"}})
	} else {
		book, _ := Helper.GetBook(id)
		c.IndentedJSON(http.StatusOK, gin.H{"book-details": book, "message": "Book issued."})
	}
}

func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query parameters not passed"})
		return
	}

	if err := DB.ReturnBook(id); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"Error": gin.H{"message": err.Error(), "type": "program created"}})
	} else {
		book, _ := Helper.GetBook(id)
		c.IndentedJSON(http.StatusOK, gin.H{"book-details": book, "message": "Book Returned."})
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := DB.DeleteBook(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": gin.H{"message": err.Error(), "type": "program generated"}})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book with id: " + id + " deleted"})
}
