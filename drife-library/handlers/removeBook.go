package handlers

import (
	"errors"
	"fmt"
)

func RemoveBook(title string) error {
	for i, book := range Library {
		if book.Title == title {
			Library = append(Library[:i], Library[i+1:]...)
			fmt.Println("Book is removed successfully")
			return nil
		}
	}
	return errors.New("book not found")
}
