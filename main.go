package main

import (
	"time"

	g "github.com/AllenDang/giu"
	_ "github.com/mattn/go-sqlite3"
)

func viewReducer() []g.Widget {
	switch state.currentView {
	case VIEW_HOME:
		return homeView()
	case VIEW_ADD_BOOK:
		return addBookView()
	case VIEW_BOOKSHELF:
		return bookshelfView()
	case VIEW_EDIT_BOOK:
		return editBookView()
	case VIEW_STATS:
		return statsView()
	}

	return []g.Widget{}
}

func loop() {
	g.SingleWindow().Layout(
		viewReducer()...,
	)
}

func onAnyKeyPressed(_ g.Key, _ g.Modifier, action g.Action) {
}

func main() {
	loadBooks()

	for i, _ := range state.books {
		state.books[i].DateAdded = time.Now()
	}

	// persistBooks("")

	w := g.NewMasterWindow("Klaras Bok Valv", 800, 800, 0)
	w.SetAdditionalInputHandlerCallback(onAnyKeyPressed)

	w.Run(loop)
}
