package handlers

import (
	"drife-library/models"
	"errors"
	"fmt"
)

var Library []models.Book

func AddBook(title, author, genre string) error {
	if title == "" || author == "" || genre == "" {
		return errors.New("book information is incomplete")
	}
	book := models.Book{Title: title, Author: author, Genre: genre}
	Library = append(Library, book)
	fmt.Println("Your Book is added successfully")
	return nil
}
