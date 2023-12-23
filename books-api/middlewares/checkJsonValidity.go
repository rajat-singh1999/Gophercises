package middlewares

import (
	"bytes"
	"encoding/json"
	Structs "example/books-api/types"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book = Structs.Book

// middlewares will be needed for presumablt all routes
// one middleware will be a default for all routes to check valididy of json checkJsonValidity.go

func TestMiddle(c *gin.Context) {
	var newBook Book
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "caught in middleware!"})
		c.Abort()
		return
	}

	// Replace the body so it can be read again in the handler
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "caught in middleware!"})
		c.Abort()
		return
	}

	id := newBook.ID
	title := newBook.Title

	fmt.Println("Hello middleware")
	fmt.Println("id:", id)
	fmt.Println("title:", title)
	c.Next()
}
