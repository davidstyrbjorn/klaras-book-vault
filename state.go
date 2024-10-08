package main

type State struct {
	// Booksshelf state
	books        []Book
	isbnResponse ISBNResponse

	isbnInput string

	addBookOpen  bool
	bokhyllaOpen bool
}

var state = State{
	books:        []Book{},
	isbnResponse: ISBNResponse{title: ""},
	addBookOpen:  false,
	bokhyllaOpen: false,
	isbnInput:    "",
}
