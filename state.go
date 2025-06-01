package main

import (
	"fmt"
	"math/rand"
	"time"

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

const (
	NUM_SORTABLE_FIELDS = 5
)

// The book as it is represented in the application layer
type Book struct {
	ISBN       string    `json:"isbn"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Read       bool      `json:"read"`
	Loaned     bool      `json:"loaned"`
	Stars      int32     `json:"stars"`
	Note       string    `json:"note"`
	DateAdded  time.Time `json:"dateAdded"`
	DateRead   time.Time `json:"dateRead"`   // The date when Read went from false to true
	DateLoaned time.Time `json:"dateLoaned"` // The date when Loaned went from false to true
}

type ISBNResponse struct {
	title  string
	Author string
	isbn   string
}

type State struct {
	currentView int

	searchString        string
	books               []Book
	bookToEdit          Book
	placeholderAuthor   string
	filterFlags         uint8
	currentSortingField int8
	ascending           bool // true = down, false = up

	isbnInput         string
	isbnResponse      ISBNResponse
	isbnError         string
	isbnLoading       bool
	manualInputTitle  string
	manualInputAuthor string
}

var state = State{
	currentView:         VIEW_HOME,
	books:               []Book{}, // Loaded from DB and kept in memory
	isbnResponse:        ISBNResponse{title: ""},
	isbnInput:           "",
	isbnError:           "",
	searchString:        "",
	isbnLoading:         false,
	filterFlags:         0,
	currentSortingField: -1,
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

	if to == VIEW_BOOKSHELF {
	}

	state.currentView = to

	g.Update()
}
