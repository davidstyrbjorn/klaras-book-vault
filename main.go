package main

import (
	"fmt"
	"io"
	"log"

	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const DBName = "./library.db"
const ISBNurl = "http://openlibrary.org/api/books?bibkeys=ISBN:%v&jscmd=details&format=json"
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
	// http://openlibrary.org/api/books?bibkeys=ISBN:1931498717&jscmd=details&format=json

	//TODO:  Find out the info we need to fill out a book struct
	resp, err := http.Get(fmt.Sprintf(ISBNurl, isbn))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Printf(sb)

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

	LookupBookFromISBN("9780060598242")

	return

	database, err := sql.Open("sqlite3", DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS Books(
		"id" PRIMARY KEY INC INTEGER
		"isbn" TEXT,
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
