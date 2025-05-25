package main

import (
	g "github.com/AllenDang/giu"
)

func statsView() []g.Widget {

	readBooks := 0
	loanedBooks := 0
	for _, book := range state.books {
		if book.Read {
			readBooks = readBooks + 1
		}
		if book.Loaned {
			loanedBooks = loanedBooks + 1
		}
	}

	percntageReadBooks := float32(float32(readBooks)/float32(len(state.books))) * 100
	percentageLoanedBooks := float32(loanedBooks) / float32(len(state.books)) * 100

	return g.Layout{
		g.Button("Tillbaka").OnClick(func() {
			changeView(VIEW_HOME)
		}),
		g.Align(g.AlignCenter).To(
			g.Column(
				g.Style().SetFontSize(32).To(g.Label("STATISTIK")),
				g.Row(
					g.Label("Totalt antal böcker: "),
					g.Labelf("%v", len(state.books)),
				),
				g.Row(
					g.Label("Utlästa böcker: "),
					g.Labelf("%v (%.2f%%)", readBooks, percntageReadBooks),
				),
				g.Row(
					g.Label("Utlånade böcker: "),
					g.Labelf("%v (%.2f%%)", loanedBooks, percentageLoanedBooks),
				),
			),
		),
	}
}
