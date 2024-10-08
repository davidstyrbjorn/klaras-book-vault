package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"database/sql"
	"net/http"

	g "github.com/AllenDang/giu"
	_ "github.com/mattn/go-sqlite3"
)

type ISBNResponse struct {
	title string
}

func LookupBookFromISBN(isbn string) ISBNResponse {
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

	stringResult := string(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(stringResult), &data)
	if err != nil {
		log.Fatal(err)
	}

	var isbnResponse ISBNResponse
	for _, v := range data {
		details := v.(map[string]interface{})["details"].(map[string]interface{})
		isbnResponse.title = details["title"].(string)
	}

	return isbnResponse
}

func FuzzySearchBooks(books []Book, search string) []Book {
	return books
}

func _() {
	fmt.Println("Library Init")

	isbnResponse := LookupBookFromISBN("9780060598242")
	log.Printf(isbnResponse.title)

	isbnResponse = LookupBookFromISBN("9780553213690")
	log.Printf(isbnResponse.title)

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

func loop() {
	g.SingleWindow().Layout(
		g.Align(g.AlignCenter).To(g.Label("Klaras Bok Valv")),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Align(g.AlignCenter).To(g.Row(
			g.Button("Lägg till bok").OnClick(func() {
				state.addBookOpen = true
			}),
			g.Button("Bokhylla").OnClick(func() {
				state.bokhyllaOpen = true
			}),
		)),

		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Align(g.AlignCenter).To(g.Label("Lägg till bok - Skanna ISBN på en bok och klicka på lägg till")),
		g.Align(g.AlignCenter).To(g.Label("Bokhylla - Här visas alla böcker, som du skannat in")),

		g.Button("Lägg till 10 böcker").OnClick(func() {
			insert10Books()
		}),
	)

	windowAddBook()
	windowBokhylla()
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()

	go dbRoutine()

	w := g.NewMasterWindow("Library", 500, 300, g.MasterWindowFlagsNotResizable)
	w.Run(loop)
}
