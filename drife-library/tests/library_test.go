package tests

import (
	"drife-library/handlers"
	"drife-library/models"
	"testing"
)

func resetLibrary() {
	handlers.Library = []models.Book{}
}

func TestAddBook(t *testing.T) {
	resetLibrary()
	err := handlers.AddBook("Test Title", "Test Author", "Test Genre")
	if err != nil {
		t.Errorf("Error adding book: %v", err)
	}
	if len(handlers.Library) != 1 {
		t.Errorf("Expected 1 book, found %d", len(handlers.Library))
	}
}

func TestRemoveBook(t *testing.T) {
	resetLibrary()
	handlers.AddBook("Test Title", "Test Author", "Test Genre")
	err := handlers.RemoveBook("Test Title")
	if err != nil {
		t.Errorf("Error removing book: %v", err)
	}
	if len(handlers.Library) != 0 {
		t.Errorf("Expected 0 books, found %d", len(handlers.Library))
	}
}

func TestSearchBook(t *testing.T) {
	resetLibrary()
	handlers.AddBook("Test Title", "Test Author", "Test Genre")
	results := handlers.SearchBook("Test Title")
	if len(results) == 0 {
		t.Errorf("Expected to find book, found none")
	}
	if results[0].Title != "Test Title" {
		t.Errorf("Expected 'Test Title', found %s", results[0].Title)
	}
}

func TestListBooksByGenre(t *testing.T) {
	resetLibrary()
	handlers.AddBook("Test Title", "Test Author", "Test Genre")
	results := handlers.ListBookByGenre("Test Genre")
	if len(results) == 0 {
		t.Errorf("Expected to find books, found none")
	}
	if results[0].Genre != "Test Genre" {
		t.Errorf("Expected 'Test Genre', found %s", results[0].Genre)
	}
}

func TestListAllBooks(t *testing.T) {
	resetLibrary()
	handlers.AddBook("Test Title", "Test Author", "Test Genre")
	results := handlers.ListAllBooks()
	if len(results) == 0 {
		t.Errorf("Expected to find books, found none")
	}
	if results[0].Title != "Test Title" {
		t.Errorf("Expected 'Test Title', found %s", results[0].Title)
	}
}
