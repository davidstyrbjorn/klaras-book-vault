package main

import g "github.com/AllenDang/giu"

func windowAddBook() {
	if !state.addBookOpen {
		return
	}

	g.Window("Lägg till Bok").Size(500, 300).IsOpen(&state.addBookOpen).Layout(
		g.Align(g.AlignCenter).To(g.Column(g.Label("Lägg till bok"),
			g.InputText(&state.isbnInput).Hint("Skanna ISBN Streckkod"),
			g.Button("Lägg till").OnClick(func() {}),
			g.Separator(),
			g.Label("Resultat"),
		)),
	)
}
