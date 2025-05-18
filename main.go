package main

import (
	g "github.com/AllenDang/giu"
	_ "github.com/mattn/go-sqlite3"
)

// func _() {
// 	fmt.Println("Library Init")

// 	isbnResponse := LookupBookFromISBN("9780060598242")
// 	log.Printf(isbnResponse.title)

// 	isbnResponse = LookupBookFromISBN("9780553213690")
// 	log.Printf(isbnResponse.title)

// 	return

// 	database, err := sql.Open("sqlite3", DBName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer database.Close()

// 	createTableSQL := `CREATE TABLE IF NOT EXISTS Books(
// 		"id" PRIMARY KEY INC INTEGER
// 		"isbn" TEXT,
// 		"title" TEXT,
// 		"page" INTEGER,
// 		"status" TEXT
// 	)`
// 	_, err = database.Exec(createTableSQL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	insertBookSQL := "INSERT INTO Books (title, page, status) VALUES (?, ?, ?)"
// 	insertBookStatement, err := database.Prepare(insertBookSQL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer insertBookStatement.Close()

// 	_, err = insertBookStatement.Exec("Edde", 0, UNREAD)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted a book 'edde' into the Books table successfully!")
// }

func viewReducer() []g.Widget {
	switch state.currentView {
	case VIEW_HOME:
		return homeView()
	case VIEW_ADD_BOOK:
		return addBookView()
	case VIEW_BOOKSHELF:
		return bookshelfView()
	case VIEW_EDIT_BOOK:
		return editBookView()
	}

	return []g.Widget{}
}

func loop() {
	g.SingleWindow().Layout(
		viewReducer()...,
	)
}

func main() {
	// isbnResponse, err := LookupBookFromISBN("9789113039527")
	// log.Printf(isbnResponse.title)

	if err := LoadBooksFromFile(); err != nil {
		panic(err)
	}

	w := g.NewMasterWindow("Klaras Bok Valv", 800, 800, 0)
	w.Run(loop)
}
