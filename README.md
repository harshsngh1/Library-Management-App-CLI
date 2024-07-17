# Drife-library

This is a simple CLI library service written in Golang

## Introduction

This project implements an interactive CLI approach for managing the library. It allows users to add books, remove books, search book based on title name and author name and list all books or books with specific genere.

## Technologies Used

- Language: Go
- Framework: None
- Database: None, used in memory map to store all book info
- Test Framework: Go testing (standard library for writing tests in Go)


## Project Structure
The project is structured as follows:

```
drife-library/
├── handlers/
│   ├── addBook.go
│   ├── removeBook.go
│   ├── searchBook.go
│   ├── listBooksByGenre.go
│   └── listAllBooks.go
├── models/
│   └── book.go
├── tests/
│   └── library_test.go
├── main.go
├── go.mod
└── README.md

```

## Installation

`Golang requirement` : Go 1.16+
1. Clone the repository :

```
git clone https://github.com/harshsngh1/Drife-library 
```

2. Navigate to the project directory:

```
cd drife-library/
```

3. Build code(optional)

```
go build -o library-cli
```
4. Run the code(only if you have build the binary file)
```
./library-cli
```
or directly:
```
go run main.go
```