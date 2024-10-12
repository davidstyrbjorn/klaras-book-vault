package main

type State struct {
	// Booksshelf state
	books        []Book
	isbnResponse ISBNResponse

	isbnInput string

	currentView uint
}

var state = State{
	books:        []Book{},
	isbnResponse: ISBNResponse{title: ""},
	isbnInput:    "",
	currentView:  HOME,
}

func switchView(newView uint) {
	// oldView := state.currentView
	if newView == BOOK_SHELF {
		dbState.performRead <- true
	}
	state.currentView = newView
}
