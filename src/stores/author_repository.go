package stores

import (
	"errors"
	"go-books/src/models"
)

var authors []models.Author

func GetAuthors() []models.Author {
	return authors
}

func GetAuthorByID(id int) (models.Author, error) {
	for _, author := range authors {
		if author.ID == id {
			return author, nil
		}
	}
	return models.Author{}, errors.New("author not found")
}

func CreateAuthor(author models.Author) {
	authors = append(authors, author)
}

func UpdateAuthor(id int, updatedAuthor models.Author) error {
	for i, author := range authors {
		if author.ID == id {
			authors[i] = updatedAuthor
			return nil
		}
	}
	return errors.New("author not found")
}

func DeleteAuthor(id int) error {
	for i, author := range authors {
		if author.ID == id {
			authors = append(authors[:i], authors[i+1:]...)
			return nil
		}
	}
	return errors.New("author not found")
}
