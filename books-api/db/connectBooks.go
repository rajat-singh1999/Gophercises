package db

import (
	"encoding/json"
	"errors"
	Config "example/books-api/config"
	"example/books-api/types"
	"io/ioutil"
	"os"
)

type Book = types.Book

var path = Config.DataPath()

func GetBook(id string) (Book, error) {
	var books []Book
	var temp Book
	books, err := LoadBooksFromFile()
	if err != nil {
		return temp, err
	} else {
		for i, b := range books {
			if b.ID == id {
				return books[i], nil
			}
		}
		return temp, errors.New("book not found")
	}
}

func LoadBooksFromFile() ([]Book, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var books []Book
	json.Unmarshal(byteValue, &books)
	return books, nil
}

func SaveBooksToFile(books []Book) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func SaveBookToBooks(book Book) ([]Book, error) {
	err := DeleteBook(book.ID)
	if err != nil {
		return nil, err
	}

	books, err := LoadBooksFromFile()
	if err != nil {
		return nil, err
	}

	books = append(books, book)

	err = SaveBooksToFile(books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func BookCheckout(id string) error {
	if books, err := LoadBooksFromFile(); err != nil {
		return err
	} else {
		for i, b := range books {
			if b.ID == id {
				if books[i].Quantity <= 0 {
					return errors.New("no more books available, try another book")
				}
				books[i].Quantity -= 1
				if err = SaveBooksToFile(books); err != nil {
					return errors.New("error while saving to file")
				} else {
					return nil
				}
			}
		}
		return errors.New("book not found")
	}
}

func ReturnBook(id string) (Book, error) {
	if books, err := LoadBooksFromFile(); err != nil {
		return books[0], err
	} else {
		for i, b := range books {
			if b.ID == id {
				books[i].Quantity += 1
				SaveBooksToFile(books) //book side update
				return books[i], nil
			}
		}
		return books[0], errors.New("book not found")
	}
}

func DeleteBook(id string) error {
	if books, err := LoadBooksFromFile(); err != nil {
		return err
	} else {
		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				if err = SaveBooksToFile(books); err != nil {
					return errors.New("error while saving to file")
				} else {
					return nil
				}
			}
		}
		return errors.New("no book matching the passed ID found")
	}
}

/* the below method checks before deleting a book if a copy of that book has been issued to any user */
func DeleteBookSafely(id string) (error, error) {
	if books, err := LoadBooksFromFile(); err != nil {
		return nil, err
	} else {
		for i, b := range books {
			if b.ID == id {
				if b.Quantity == b.Total {
					books = append(books[:i], books[i+1:]...)
					if err = SaveBooksToFile(books); err != nil {
						return nil, errors.New("error while saving to file")
					} else {
						return nil, nil
					}
				} else {
					return errors.New("book cannot be deleted, some copies are issued to users"), nil
				}
			}
		}
		return nil, errors.New("no book matching the passed ID found")
	}
}
