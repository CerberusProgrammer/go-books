package stores

import (
	"errors"
	"go-books/src/models"
)

var books []models.Book

func GetBooks() []models.Book {
	return books
}

func GetBookByID(id int) (models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func CreateBook(book models.Book) {
	books = append(books, book)
}

func UpdateBook(id int, updatedBook models.Book) error {
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			return nil
		}
	}
	return errors.New("book not found")
}

func DeleteBook(id int) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
