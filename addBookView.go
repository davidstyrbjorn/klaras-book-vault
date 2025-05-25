package main

import (
	"time"
	"unicode"

	g "github.com/AllenDang/giu"
)

var putFocusOnIsbnInput = false

func verifyIsbnString(isbn string) (bool, string) {
	result := ""
	// Strip isbn of any hyphens first thing
	for _, char := range state.isbnInput {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}
	isbn = result

	length := len(isbn)
	if length != 10 && length != 13 {
		return false, isbn
	}

	for _, char := range isbn {
		if !unicode.IsDigit(char) {
			return false, isbn
		}
	}

	return true, isbn
}

func onAddBookClick() {
	state.isbnLoading = true
	defer func() {
		state.isbnLoading = false
		putFocusOnIsbnInput = true
		g.Update()
	}()
	g.Update()

	if state.isbnInput == "" {
		state.isbnError = "Var snäll fyll i ett ISBN nummer"
		return
	}

	// Verify isbn string, also returns a pure isbn version in case it contains any extra technically allowed characters
	correct, newIsbn := verifyIsbnString(state.isbnInput)
	if !correct {
		state.isbnError = "Inte ett giltigt ISBN nummer"
		return
	}
	state.isbnInput = newIsbn

	isbnResponse, err := LookupBookFromISBN(state.isbnInput)
	if err != nil {
		state.isbnError = err.Error()
		return
	}

	state.isbnResponse = isbnResponse

	for i := range state.books {
		if state.books[i].ISBN == state.isbnInput {
			state.isbnError = "Du har redan en bok med detta isbn nummer!"
			return
		}
	}

	// If we get to here, insert the book, dump into binary and set that we have no error!
	state.books = append(state.books, Book{Title: state.isbnResponse.title, ISBN: state.isbnResponse.isbn, Author: state.isbnResponse.Author, DateAdded: time.Now()})
	go persistBooks("")
	state.isbnError = "" // No error to show if we got to here!
	state.isbnInput = ""
}

func manuallyAddBook() {
	state.books = append(state.books, Book{
		ISBN:   state.isbnInput,
		Title:  state.manualInputTitle,
		Author: state.manualInputAuthor,
		Read:   false,
		Loaned: false,
		Stars:  0,
		Note:   "",
	})
	go persistBooks("")
	state.isbnInput = ""
	state.isbnError = ""
	putFocusOnIsbnInput = true
	g.Update()
}

func addBookView() []g.Widget {
	return []g.Widget{
		g.Button("Tillbaka").OnClick(func() { changeView(VIEW_HOME) }),
		g.Align(g.AlignCenter).To(g.Column(
			g.Label("Lägg till bok"),
			g.Custom(func() {
				if putFocusOnIsbnInput {
					putFocusOnIsbnInput = false
					g.SetKeyboardFocusHere()
				}
			}),
			g.InputText(&state.isbnInput).Hint("Skanna ISBN Streckkod").ID("isbn-input"),
			g.Button("Sök").OnClick(func() {
				go onAddBookClick()
			}).Disabled(state.isbnLoading),
			g.Separator(),
			g.Label("Hittad bok!"),
			g.Condition(state.isbnError != "",
				g.Column(
					g.Style().SetColor(g.StyleColorText, FailedRed).To(g.Labelf("%v", state.isbnError)),
					g.Label("Lägg till manuellt?"),
					g.Row(
						g.InputText(&state.manualInputTitle),
						g.Label("Titel"),
					),
					g.Row(
						g.InputText(&state.manualInputAuthor),
						g.Label("Författare"),
					),
					g.Button("Lägg till manuellt").OnClick(func() {
						go manuallyAddBook()
					}),
				),
				nil,
			),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == "", g.Labelf("Titel = %v", state.isbnResponse.title), nil),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == "", g.Labelf("ISBN = %v", state.isbnResponse.isbn), nil),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == "",
				g.Style().SetColor(g.StyleColorText, SuccessGreen).To(g.Label("Boken tillagd")),
				nil,
			),
		)),
	}
}
