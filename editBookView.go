package main

import (
	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	clipboard "github.com/d-tsuji/clipboard"
)

func saveEditedBook() {
	for i := range state.books {
		if state.books[i].ISBN == state.bookToEdit.ISBN {
			state.books[i] = state.bookToEdit // Overwrite here
			return
		}
	}
}

func checkForEdits() bool {
	for i := range state.books {
		if state.books[i].ISBN == state.bookToEdit.ISBN {
			return state.books[i] != state.bookToEdit
		}
	}
	return false
}

func removeEditedBook() {
	for i := range state.books {
		if state.books[i].ISBN == state.bookToEdit.ISBN {
			state.books = append(state.books[:i], state.books[i+1:]...)
			return
		}
	}
}

func editBookView() []g.Widget {
	return g.Layout{
		giu.PopupModal("Du har osparade ändringar").Layout(
			giu.Align(giu.AlignCenter).To(
				giu.Button("Stäng och Spara").OnClick(func() {
					giu.CloseCurrentPopup()
					saveEditedBook()
					go DumpBookToFile()

					changeView(VIEW_BOOKSHELF)
				}),
				giu.Button("Stäng").OnClick(func() {
					changeView(VIEW_BOOKSHELF)
					giu.CloseCurrentPopup()
				}),
			),
		),
		g.Button("Tillbaka").OnClick(func() {
			if checkForEdits() {
				giu.OpenPopup("Du har osparade ändringar")
			} else {
				changeView(VIEW_BOOKSHELF)
			}
		}),
		g.Align(g.AlignCenter).To(g.Column(
			g.Label("Editera din bok"),
			g.Row(
				g.Labelf("ISBN = %v", state.bookToEdit.ISBN),
				g.Button("Kopiera ISBN").OnClick(func() {
					clipboard.Set(state.bookToEdit.ISBN)
				}),
			),
			g.InputText(&state.bookToEdit.Title).Hint("Titel"),
			g.InputText(&state.bookToEdit.Author).Hint("Författare"),
			g.InputText(&state.bookToEdit.Note).Hint("Egen Notering"),
			g.Checkbox("Utlånad?", &state.bookToEdit.Loaned),
			g.Checkbox("Utläst?", &state.bookToEdit.Read),
			g.InputInt(&state.bookToEdit.Stars),
			g.Row(
				g.Button("Spara").OnClick(func() {
					saveEditedBook()
					go DumpBookToFile()

					changeView(VIEW_BOOKSHELF)
				}),
				g.Button("Avbryt").OnClick(func() {
					changeView(VIEW_BOOKSHELF) // Do nothing and simply exit
				}),
				g.Button("Ta Bort Bok").OnClick(func() {
					removeEditedBook()
					go DumpBookToFile()

					changeView(VIEW_BOOKSHELF)
				}),
			),
		)),
	}
}
