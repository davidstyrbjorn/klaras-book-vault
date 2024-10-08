package main

import (
	"database/sql"
	"fmt"
	"time"
)

// The book as it is represented in the application layer
type Book struct {
	id     uint
	isbn   string
	title  string
	author string
	read   bool
	loaned bool
	stars  uint8
	note   string
}

type DBState struct {
	db             *sql.DB
	createBookStmt *sql.Stmt
	readBookStmt   *sql.Stmt

	performRead chan (bool)
}

var dbState = DBState{
	db:             nil,
	createBookStmt: nil,
	readBookStmt:   nil,
	performRead:    make(chan bool),
}

func initDB() error {
	var err error
	dbState.db, err = sql.Open("sqlite3", DBName)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS Books(
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"isbn" TEXT NOT NULL,
		"title" TEXT NOT NULL,
		"author" TEXT NOT NULL,
		"read" BOOLEAN NOT NULL,
		"loaned" BOOLEAN NOT NULL,
		"stars" INTEGER CHECK (stars >= 0 AND stars <= 5) NOT NULL,
		"note" TEXT NOT NULL
	)`
	_, err = dbState.db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	dbState.createBookStmt, err = dbState.db.Prepare("INSERT INTO Books (isbn, title, author, read, loaned, stars, note) VALUES (?, ?, ?, 0, 0, 0, '')")
	if err != nil {
		return err
	}

	dbState.readBookStmt, err = dbState.db.Prepare("SELECT * FROM Books WHERE id = ?")
	if err != nil {
		return err
	}

	return nil
}

func closeDB() {
	dbState.createBookStmt.Close()
	dbState.readBookStmt.Close()
	dbState.db.Close()
}

func createBook(book Book) {
	_, err := dbState.createBookStmt.Exec(book.isbn, book.title, book.author)
	if err != nil {
		panic(err)
	}
}

func readAllBooks() ([]Book, error) {
	rows, err := dbState.db.Query("SELECT * FROM Books")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.id, &book.isbn, &book.title, &book.author, &book.read, &book.loaned, &book.stars, &book.note)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return books, nil
}

func fuzzySearchBook(search string) ([]Book, error) {
	// Execute query, get rows back
	rows, err := dbState.db.Query(fuzzySearchBookQuery, search)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Turn the rows into a slice of books
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.id, &book.isbn, &book.title, &book.author, &book.read, &book.loaned, &book.stars, &book.note)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return books, nil
}

func clearBookTable() {
	_, err := dbState.db.Exec("DELETE FROM Books")
	if err != nil {
		panic(err)
	}
}

func insert10Books() {
	createBook(Book{isbn: "9780062374888", title: "To Kill a Mockingbird", author: "Harper Lee"})
	createBook(Book{isbn: "9780451524935", title: "1984", author: "George Orwell"})
	createBook(Book{isbn: "9780316033803", title: "The Great Gatsby", author: "F. Scott Fitzgerald"})
	createBook(Book{isbn: "9780241380279", title: "Pride and Prejudice", author: "Jane Austen"})
	createBook(Book{isbn: "9780062401956", title: "The Catcher in the Rye", author: "J.D. Salinger"})
	createBook(Book{isbn: "9780451526342", title: "Animal Farm", author: "George Orwell"})
	createBook(Book{isbn: "9780316015844", title: "The Picture of Dorian Gray", author: "Oscar Wilde"})
	createBook(Book{isbn: "9780241954658", title: "Sense and Sensibility", author: "Jane Austen"})
	createBook(Book{isbn: "9780062801970", title: "The Lord of the Rings", author: "J.R.R. Tolkien"})
}

// func updateBook() {}

// func deleteBook() {}

func dbRoutine() {
	for {
		select {
		case <-dbState.performRead:
			books, err := readAllBooks()
			if err != nil {
				panic(err)
			}
			state.books = books
		default:
			// default case to prevent blocking
		}

		time.Sleep(1 * time.Second)
	}
}
