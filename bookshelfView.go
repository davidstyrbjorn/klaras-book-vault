package main

import (
	"sort"
	"strings"

	g "github.com/AllenDang/giu"
)

type SortyByBook []Book

func (a SortyByBook) Len() int      { return len(a) }
func (a SortyByBook) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortyByBook) Less(i, j int) bool {
	switch state.currentSortingField {
	case 0:
		return strings.ToLower(a[i].Title) < strings.ToLower(a[j].Title)
	case 1:
		return strings.ToLower(a[i].Author) < strings.ToLower(a[j].Author)
	case 2:
		return a[i].Stars < a[j].Stars
	case 3:
		return a[i].Loaned && !a[j].Loaned
	case 4:
		return a[i].Read && !a[j].Read
	default:
		break
	}

	const BALLS = false
	return BALLS
}

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

func doesBookPassFilter(book Book) bool {
	if state.filterFlags&ONLY_EMPTY_ISBN != 0 {
		if book.ISBN == "" {
			return true
		} else {
			return false
		}
	}

	return true
}

func reverse() {
	lastIndex := len(state.books) - 1
	for i := range len(state.books) / 2 {
		j := lastIndex - i
		state.books[i], state.books[j] = state.books[j], state.books[i]
	}
}

func rowHeaderPressed(what uint8) {
	if state.sortingDirections[what] == g.DirectionUp {
		state.sortingDirections[what] = g.DirectionDown
	} else {
		state.sortingDirections[what] = g.DirectionUp
	}

	state.currentSortingField = what

	var sortyByBooks SortyByBook = state.books
	sort.Sort(sortyByBooks)
	if state.sortingDirections[state.currentSortingField] == g.DirectionUp {
		reverse()
	}
}

func buildBokhylla() []*g.TableRowWidget {
	var rows []*g.TableRowWidget

	// Column headers
	rows = append(rows, g.TableRow(
		g.Label(""),
		g.Row(g.Label("Titel"), g.ArrowButton(state.sortingDirections[0]).OnClick(func() {
			rowHeaderPressed(0)
		})),
		g.Row(g.Label("Författare"), g.ArrowButton(state.sortingDirections[1]).OnClick(func() {
			rowHeaderPressed(1)
		})),
		g.Row(g.Label("Betyg"), g.ArrowButton(state.sortingDirections[2]).OnClick(func() {
			rowHeaderPressed(2)
		})),
		g.Row(g.Label("Utlånad?"), g.ArrowButton(state.sortingDirections[3]).OnClick(func() {
			rowHeaderPressed(3)
		})),
		g.Row(g.Label("Utläst?"), g.ArrowButton(state.sortingDirections[4]).OnClick(func() {
			rowHeaderPressed(4)
		})),
	).Flags(g.TableRowFlagsHeaders))

	for i, book := range state.books {
		if !doesBookPassSearchCheck(book) {
			continue
		}

		if !doesBookPassFilter(book) {
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
		g.Row(
			g.RadioButton("Bara böcker med tomma ISBN fält", state.filterFlags&ONLY_EMPTY_ISBN != 0).OnChange(func() {
				if state.filterFlags&ONLY_EMPTY_ISBN != 0 {
					state.filterFlags &= ^ONLY_EMPTY_ISBN
				} else {
					state.filterFlags |= ONLY_EMPTY_ISBN
				}
			}),
		),
		g.Table().Rows(buildBokhylla()...),
	}
}
