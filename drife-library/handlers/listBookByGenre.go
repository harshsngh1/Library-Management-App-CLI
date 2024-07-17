package handlers

import (
	"drife-library/models"
	"fmt"
)

func ListBookByGenre(genre string) []models.Book {
	var result []models.Book
	for _, book := range Library {
		if book.Genre == genre {
			result = append(result, book)
		}
	}
	if len(result) == 0 {
		fmt.Println("No books found in this genre")
	}
	return result
}
