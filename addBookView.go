package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func onAddBookClick() {

	if state.isbnInput == "" {
		state.isbnError = fmt.Errorf("var snäll fyll i ett ISBN nummer tack")
		return
	}

	isbnResponse, err := LookupBookFromISBN(state.isbnInput)
	if err != nil {
		state.isbnError = err
		return
	}

	state.isbnResponse = isbnResponse
	state.bookAlreadyExists = false

	for i := range state.books {
		if state.books[i].ISBN == state.isbnInput {
			state.bookAlreadyExists = true
		}
	}

	// If we have no books with this ISBN, we can just insert it and be done
	if !state.bookAlreadyExists {
		state.books = append(state.books, Book{Title: state.isbnResponse.title, ISBN: state.isbnResponse.isbn})
		go DumpBookToFile()
	}

	state.isbnInput = ""
	g.Update()
}

func addBookView() []g.Widget {
	return []g.Widget{
		g.Button("Tillbaka").OnClick(func() { changeView(VIEW_HOME) }),
		g.Align(g.AlignCenter).To(g.Column(g.Label("Lägg till bok"),
			g.InputText(&state.isbnInput).Hint("Skanna ISBN Streckkod"),
			g.Button("Sök").OnClick(func() {
				go onAddBookClick()
			}),
			g.Separator(),
			g.Label("Resultat"),
			// g.Condition(state.isbnError != nil, g.Labelf("Error = %v", state.isbnError.Error()), nil),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == nil, g.Labelf("Titel = %v", state.isbnResponse.title), nil),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == nil, g.Labelf("ISBN = %v", state.isbnResponse.isbn), nil),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == nil && !state.bookAlreadyExists,
				g.Style().SetColor(g.StyleColorText, SuccessGreen).To(g.Label("Boken tillagd")),
				g.Label(""),
			),
			g.Condition(state.isbnResponse.title != "" && state.isbnError == nil && state.bookAlreadyExists,
				g.Style().SetColor(g.StyleColorText, FailedRed).To(g.Label("Du har redan denna bok tillagd")),
				g.Label(""),
			),
		)),
	}
}
