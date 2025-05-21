package main

import (
	"strings"

	g "github.com/AllenDang/giu"
)

func starsToString(stars int32) string {
	result := ""
	for range stars {
		result += "*"
	}
	return result
}

func doesBookPassSearchCheck(book Book) bool {
	s := strings.ToLower(state.searchString)

	// Concatenate searchable fields into one string
	content := strings.ToLower(
		book.ISBN + " " +
			book.Title + " " +
			book.Author + " " +
			book.Note,
	)

	// Basic fuzzy-ish check: does the search string appear in the content?
	return strings.Contains(content, s)
}

func buildBokhylla() []*g.TableRowWidget {
	var rows []*g.TableRowWidget

	// Column headers
	rows = append(rows, g.TableRow(
		g.Label(""),
		g.Label("Titel"),
		g.Label("Författare"),
		g.Label("Betyg"),
		g.Label("Utlånad?"),
		g.Label("Utläst?"),
	).Flags(g.TableRowFlagsHeaders))

	for i, book := range state.books {
		if !doesBookPassSearchCheck(book) {
			continue
		}

		// Closure capture fix: shadow the loop variable
		b := book

		rows = append(rows, g.TableRow(
			g.Row(
				g.Labelf("%v", i+1),
				g.Button("Edit").OnClick(func() {
					state.bookToEdit = b
					changeView(VIEW_EDIT_BOOK)
				}),
			),
			g.Label(b.Title),
			g.Label(b.Author),
			g.Label(starsToString(b.Stars)),
			g.Condition(b.Loaned, g.Label("Utlånad"), g.Label("Hemma")),
			g.Condition(b.Read, g.Label("Utläst"), g.Label("TBR")),
		))
	}

	return rows
}

func bookshelfView() []g.Widget {
	return g.Layout{
		g.Row(
			g.Button("Tillbaka").OnClick(func() {
				changeView(VIEW_HOME)
			}),
			g.Spacing(),
			g.Spacing(),
			g.Label("Sök"),
			g.InputText(&state.searchString),
			g.Spacing(),
		),
		g.Table().Rows(buildBokhylla()...),
	}
}
