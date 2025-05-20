package main

import (
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
	}

	return []g.Widget{}
}

func loop() {
	g.SingleWindow().Layout(
		viewReducer()...,
	)
}

func main() {
	if err := LoadBooksFromFile(); err != nil {
		println("Can't find books binary file!")
	}

	w := g.NewMasterWindow("Klaras Bok Valv", 800, 800, 0)
	w.Run(loop)
}
