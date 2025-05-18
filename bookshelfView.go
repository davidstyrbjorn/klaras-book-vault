package main

import (
	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
)

func starsToString(stars int32) string {
	result := ""
	for range stars {
		result += "*"
	}
	return result
}

func buildBokhylla() []*g.TableRowWidget {
	rows := make([]*g.TableRowWidget, len(state.books)+1)

	// Column headers
	rows[0] = g.TableRow(
		g.Label(""),
		g.Label("Titel"),
		g.Label("Författare"),
		g.Label("Betyg"),
		g.Label("Anteckning"),
	)

	for i, book := range state.books {
		rows[i+1] = g.TableRow(
			g.Button("Edit").OnClick(func() {
				state.bookToEdit = state.books[i]
				changeView(VIEW_EDIT_BOOK)
			}),
			g.Label(book.Title),
			g.Label(book.Author),
			g.Label(starsToString(book.Stars)),
			g.Label(book.Note),
		)
	}

	return rows
}

func bookshelfView() []giu.Widget {
	return giu.Layout{
		g.Button("Tillbaka").OnClick(func() {
			changeView(VIEW_HOME)
		}),
		g.Table().Rows(buildBokhylla()...),
	}
}
