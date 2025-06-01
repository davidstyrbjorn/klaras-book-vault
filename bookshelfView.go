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
		return a[i].DateAdded.After(a[j].DateAdded)
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

// This function just takes the books we have and reverse them
func reverse() {
	lastIndex := len(state.books) - 1
	for i := range len(state.books) / 2 {
		j := lastIndex - i
		state.books[i], state.books[j] = state.books[j], state.books[i]
	}
}

func rowHeaderPressed(what int8) {
	// Did we click the one we are already on? if so apply some extra logic
	if what == state.currentSortingField {
		if !state.ascending {
			state.ascending = true
		} else {
			what = -1
		}
	} else {
		state.ascending = false
	}
	state.currentSortingField = what

	// Perform sort, reverse if ascending is false
	var sortyByBooks SortyByBook = state.books
	sort.Sort(sortyByBooks)
	if state.ascending {
		reverse()
	}

	g.Update()
}

func getArrowDirectionForRowHeader(what int8) g.Direction {
	if what != state.currentSortingField {
		return g.DirectionRight
	}

	if state.ascending {
		return g.DirectionUp
	} else {
		return g.DirectionDown
	}
}

func buildBokhylla() []*g.TableRowWidget {
	var rows []*g.TableRowWidget

	// Column headers
	rows = append(rows, g.TableRow(
		g.Label(""),
		g.Row(g.Label("Titel"), g.ArrowButton(getArrowDirectionForRowHeader(0)).OnClick(func() {
			rowHeaderPressed(0)
		})),
		g.Row(g.Label("Författare"), g.ArrowButton(getArrowDirectionForRowHeader(1)).OnClick(func() {
			rowHeaderPressed(1)
		})),
		g.Row(g.Label("Betyg"), g.ArrowButton(getArrowDirectionForRowHeader(2)).OnClick(func() {
			rowHeaderPressed(2)
		})),
		g.Row(g.Label("Utlånad?"), g.ArrowButton(getArrowDirectionForRowHeader(3)).OnClick(func() {
			rowHeaderPressed(3)
		})),
		g.Row(g.Label("Utläst?"), g.ArrowButton(getArrowDirectionForRowHeader(4)).OnClick(func() {
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
