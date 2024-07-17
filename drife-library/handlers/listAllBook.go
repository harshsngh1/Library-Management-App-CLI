package handlers

import (
	"drife-library/models"
	"fmt"
)

func ListAllBooks() []models.Book {
	if len(Library) == 0 {
		fmt.Println("\n No books are available")
	}
	return Library
}
