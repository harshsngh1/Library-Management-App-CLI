package main

import (
	"bufio"
	"drife-library/handlers"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n Welcome to Drife Library")

	for {
		fmt.Println("\n Please choose an option to proceed with :")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Search for a book")
		fmt.Println("4. List books by genre")
		fmt.Println("5. List all books")
		fmt.Println("6. Exit")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Enter book title:")
			scanner.Scan()
			title := scanner.Text()

			fmt.Println("Enter book author:")
			scanner.Scan()
			author := scanner.Text()

			fmt.Println("Enter book genre:")
			scanner.Scan()
			genre := scanner.Text()

			handlers.AddBook(title, author, genre)
		case "2":
			fmt.Println("\n Enter the book title to remove")
			scanner.Scan()
			title := scanner.Text()
			err := handlers.RemoveBook(title)
			if err != nil {
				fmt.Println(err)
			}
		case "3":
			fmt.Println("\n Enter the search query (title name or author name)")
			scanner.Scan()
			query := scanner.Text()
			result := handlers.SearchBook(query)
			for _, book := range result {
				fmt.Printf("Title: %s, Author: %s, Genre: %s\n", book.Title, book.Author, book.Genre)
			}
		case "4":
			fmt.Println("\n Enter the genre to list:")
			scanner.Scan()
			genre := scanner.Text()
			result := handlers.ListBookByGenre(genre)
			for _, book := range result {
				fmt.Printf("Title: %s, Author: %s, Genre: %s\n", book.Title, book.Author, book.Genre)
			}
		case "5":
			fmt.Println("\n List of all books we have : ")
			result := handlers.ListAllBooks()
			for _, book := range result {
				fmt.Printf("Title: %s, Author: %s, Genre: %s\n", book.Title, book.Author, book.Genre)
			}
		case "6":
			fmt.Println("\n Exiting the Library")
			return

		default:
			fmt.Println("\n Invalid choice. Please choose a valid option.")
		}
	}
}
