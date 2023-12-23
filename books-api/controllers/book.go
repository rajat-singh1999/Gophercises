package Controllers

import (
	"errors"
	DB "example/books-api/db"
	Structs "example/books-api/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book = Structs.Book

func BookById(c *gin.Context) {
	id := c.Param("id") // get id parameter

	// calling helper function that does the actual job
	book, err := DB.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"type": "error", "error": gin.H{"message": "book with book id does not exist."}})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"type": "success", "message": "Book found.", "book-details": book})
}

func GetBooks(c *gin.Context) {
	// gets all books
	if books, err := DB.LoadBooksFromFile(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": gin.H{"message": err.Error()}, "type": "error"})
	} else {
		c.IndentedJSON(http.StatusOK, books)
	}
}

func AddBook(c *gin.Context) {
	var newBook Book
	// trying to bind the received body data to the newBook variable
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "type": "error"})
		return
	}

	// all the fields should be filled (business logic)
	if newBook.ID == "" || newBook.Title == "" || newBook.Author == "" || newBook.Total == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": errors.New("book parameters not complete").Error(), "type": "program created"}, "type": "error"})
		return
	}

	newBook.Quantity = newBook.Total

	// checking if a book with this id already exists=
	if oldBook, err := DB.GetBook(newBook.ID); err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"book-details": oldBook, "message": errors.New("book with id already exists, try to use '/update' route").Error(), "type": "program created"}, "type": "error"})
		return
	}

	// loading all the books from file to a go slice,
	// appending the newBook to the slice
	// saving the the new slice in the file
	if books, err := DB.LoadBooksFromFile(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "type": "error"})
	} else {
		books = append(books, newBook)
		if err := DB.SaveBooksToFile(books); err != nil {
			c.IndentedJSON(http.StatusNotModified, gin.H{"error": gin.H{"message": "Book was not inserted.", "book-details": newBook, "Error": err}, "type": "error"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"type": "success", "book-details": newBook, "message": "Book Added Successfully!"})
		}
	}
}

func AddExistingBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	book, err := DB.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Error": gin.H{
				"message": err.Error(),
				"type":    "Program Created",
			},
			"type": "error",
		})
		return
	}

	qualityQuery, ok := c.GetQuery("c")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": gin.H{
				"message": errors.New("no query parameter passed"),
				"type":    "Caller Created",
			},
			"type": "error",
		})
		return
	}

	if adding, err := strconv.Atoi(qualityQuery); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": gin.H{
				"message": errors.New("no query parameter passed"),
				"type":    "Caller Created",
			},
			"type": "error",
		})
	} else {
		book.Quantity = book.Quantity + adding
		book.Total = book.Total + adding
	}

	_, err = DB.SaveBookToBooks(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error": gin.H{
				"message": err.Error(),
				"type":    "go created",
			},
			"type": "error",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":      "books added",
		"book-details": book,
		"type":         "success",
	})
}

func UpdateBook(c *gin.Context) {
	var newBook Book
	// binding update parameters to the newBook variable
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": err.Error(),
				"type":    "program created",
			},
			"type": "error",
		})
		return
	}

	// getting the old book for reference,
	// the below code also checks while getting if a book with
	// the passed id is present or not
	var oldBook, err = DB.GetBook(newBook.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": err.Error(),
				"type":    "program created",
			},
			"type": "error",
		})
		return
	}

	// the code below this tries to validate the data sent in the request
	newBook.Quantity = oldBook.Quantity
	newBook.Total = oldBook.Total
	// ID field must not be empty
	if newBook.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": errors.New("id parameter not provided").Error(),
				"type":    "program created",
			},
			"type": "error",
		})
		return
	}
	// if both title and author data is missing than it is caught here
	if newBook.Title == "" && newBook.Author == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": errors.New("you may change either title or author, but you are changing none").Error(),
				"type":    "program created",
			},
			"type": "error",
		})
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
		c.IndentedJSON(http.StatusNotModified, gin.H{"type": "error", "Error": gin.H{"message": err.Error(), "type": "program created"}})
	} else {
		if books, err := DB.LoadBooksFromFile(); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"type": "error", "Error": gin.H{"message": err.Error(), "type": "program created"}})
		} else {
			books = append(books, newBook)
			if err := DB.SaveBooksToFile(books); err != nil {
				c.IndentedJSON(http.StatusNotModified, gin.H{"type": "error", "message": "Book was not inserted.", "book-details": newBook, "Error": gin.H{"message": err.Error(), "type": "program created"}})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"type": "success", "book-details": newBook, "message": "book updated"})
			}
		}
	}
}

func CheckOut(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "query parameters not passed", "type": "error"})
		return
	}

	uid := c.Param("uid")
	user, err := DB.GetUser(uid) // checking if user exists
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "Error": gin.H{"message": err.Error(), "type": "program created"}})
		return
	}

	// checking if the user issued array already has this book or not
	if err := DB.CheckIfBookIssuedByUser(uid, id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"type": "error",
			"Error": gin.H{
				"message": err.Error(),
				"type":    "program created",
			},
		})
		return
	}

	if err := DB.BookCheckout(id); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"type": "error", "Error": gin.H{"message": err.Error(), "type": "program created"}})
		return
	} else {
		book, _ := DB.GetBook(id)
		user.Issued = append(user.Issued, book.ID)
		_, err = DB.SaveUserToUsers(user)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"type":         "success",
			"book-details": book,
			"user-details": gin.H{
				"id":       user.ID,
				"username": user.Username,
			},
			"message": "Book issued.",
		})
	}
}

func ReturnBook(c *gin.Context) {
	uid := c.Param("uid")

	bid, ok := c.GetQuery("bid")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"type": "error", "Error": "Query parameter not found!"})
		return
	}

	user, internalErr, badRequestErr := DB.UserReturnsBook(uid, bid) //update user for return
	if internalErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"type":  "error",
			"Error": internalErr.Error(),
		})
		return
	}
	if badRequestErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"type":  "error",
			"Error": badRequestErr.Error(),
		})
		return
	}

	book, err := DB.ReturnBook(bid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"type":  "error",
			"Error": err.Error(),
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"type":         "success",
		"message":      "Book returned successfully!",
		"book-details": book,
		"user-details": gin.H{
			"id":                     uid,
			"username":               user.Username,
			"Issued-books-with-user": user.Issued,
		},
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if conflictError, err := DB.DeleteBookSafely(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error(), "type": "program generated"}, "type": "error"})
		return
	} else if conflictError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error(), "type": "program generated"}, "type": "error"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "book with id: " + id + " deleted", "type": "success"})
	}
}
