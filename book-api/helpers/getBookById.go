package helpers

import (
	"errors"
	DB "example/books-api/db"
	Structs "example/books-api/types"
)

type Book = Structs.Book

func GetBook(id string) (*Book, error) {
	if books, err := DB.LoadBooksFromFile(); err != nil {
		return nil, err
	} else {
		for i, b := range books {
			if b.ID == id {
				return &books[i], nil
			}
		}
		return nil, errors.New("book not found")
	}
}
