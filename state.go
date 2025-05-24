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
	VIEW_STATS
)

// Filter flag is a bitmask where each bit represents some type of filter
const (
	ONLY_EMPTY_ISBN = uint8(1 << 0)
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
	title  string
	Author string
	isbn   string
}

type State struct {
	currentView int

	searchString      string
	books             []Book
	bookToEdit        Book
	placeholderAuthor string
	filterFlags       uint8

	isbnInput         string
	isbnResponse      ISBNResponse
	isbnError         string
	isbnLoading       bool
	manualInputTitle  string
	manualInputAuthor string
}

var state = State{
	currentView:  VIEW_HOME,
	books:        []Book{}, // Loaded from DB and kept in memory
	isbnResponse: ISBNResponse{title: ""},
	isbnInput:    "",
	isbnError:    "",
	searchString: "",
	isbnLoading:  false,
	filterFlags:  0,
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
		putFocusOnIsbnInput = true
	}

	if to == VIEW_EDIT_BOOK {
		pickRandomPlaceholderAuthor()
	}

	state.currentView = to

	g.Update()
}
