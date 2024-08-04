package main

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const DBName = "./library.db"
const (
	UNREAD     = "unread"
	READ       = "read"
	LOANED_OUT = "loaned out"
)

type Book struct {
	isbn        string
	title       string
	currentPage uint
	status      string
}

func LookupBookFromISBN(isbn string) Book {
	//TODO:  Find out the info we need to fill out a book struct

	return Book{
		isbn:        "",
		title:       "",
		currentPage: 0,
		status:      UNREAD,
	}
}

func FuzzySearchBooks(books []Book, search string) []Book {
	return books
}

func main() {
	fmt.Println("Library Init")

	database, err := sql.Open("sqlite3", DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS Books(
		"isbn" TEXT PRIMARY KEY,
		"title" TEXT,
		"page" INTEGER,
		"status" TEXT
	)`
	_, err = database.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	insertBookSQL := "INSERT INTO Books (title, page, status) VALUES (?, ?, ?)"
	insertBookStatement, err := database.Prepare(insertBookSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer insertBookStatement.Close()

	_, err = insertBookStatement.Exec("Edde", 0, UNREAD)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a book 'edde' into the Books table successfully!")
}
