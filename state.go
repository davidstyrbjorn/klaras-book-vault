package main

import (
	"fmt"
	"math/rand"

	g "github.com/AllenDang/giu"
)

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

	searchString      string
	books             []Book
	bookToEdit        Book
	placeholderAuthor string

	isbnInput    string
	isbnResponse ISBNResponse
	isbnError    string
	isbnLoading  bool
}

var state = State{
	currentView:  VIEW_HOME,
	books:        []Book{}, // Loaded from DB and kept in memory
	isbnResponse: ISBNResponse{title: ""},
	isbnInput:    "",
	isbnError:    "",
	searchString: "",
	isbnLoading:  false,
}

func resetAddBookState() {
	fmt.Println("Resetting add book state")
	state.isbnInput = ""
	state.isbnError = ""
	state.isbnResponse.title = ""
	state.isbnLoading = false
}

func pickRandomPlaceholderAuthor() {
	x := rand.Intn(100)
	if x <= 25 {
		state.placeholderAuthor = "Anders Andersson"
	} else if x <= 50 {
		state.placeholderAuthor = "Klara Klarasson"
	} else if x <= 75 {
		state.placeholderAuthor = "Erik Eriksson"
	} else {
		state.placeholderAuthor = "Sara Sarasson"
	}
}

func changeView(to int) {
	// oldView := state.currentView

	if to == VIEW_ADD_BOOK {
		resetAddBookState()
	}

	if to == VIEW_EDIT_BOOK {
		pickRandomPlaceholderAuthor()
	}

	state.currentView = to

	g.Update()
}
