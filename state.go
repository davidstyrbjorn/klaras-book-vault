package main

import g "github.com/AllenDang/giu"

const (
	VIEW_HOME = iota
	VIEW_ADD_BOOK
	VIEW_BOOKSHELF
	VIEW_EDIT_BOOK
)

// The book as it is represented in the application layer
type Book struct {
	ISBN   string
	Title  string
	Author string
	Read   bool
	Loaned bool
	Stars  int32
	Note   string
}

type ISBNResponse struct {
	title string
	isbn  string
}

type State struct {
	currentView int

	books      []Book
	bookToEdit Book

	bookAlreadyExists bool
	isbnInput         string
	isbnResponse      ISBNResponse
	isbnError         error
}

var state = State{
	currentView:  VIEW_HOME,
	books:        []Book{}, // Loaded from DB and kept in memory
	isbnResponse: ISBNResponse{title: ""},
	isbnInput:    "",
	isbnError:    nil,
}

func resetAddBookState() {
	state.isbnInput = ""
	state.isbnError = nil
	state.isbnResponse.title = ""
}

func changeView(to int) {
	// oldView := state.currentView

	if to == VIEW_ADD_BOOK {
		resetAddBookState()
	}

	state.currentView = to

	g.Update()
}
