package main

import g "github.com/AllenDang/giu"

func buildBokhylla() []*g.TableRowWidget {
	rows := make([]*g.TableRowWidget, len(state.books))

	for i, book := range state.books {
		rows[i] = g.TableRow(
			g.Label(book.title),
			g.Label(book.author),
			g.Label(book.isbn),
		)
	}

	return rows
}

func windowBokhylla() {
	if !state.bokhyllaOpen {
		return
	}

	g.Window("Bokhylla").Size(500, 300).IsOpen(&state.bokhyllaOpen).Layout(
		g.Button("Ladda Bokhylla").OnClick(func() {
			dbState.performRead <- true
		}),
		g.Table().Rows(buildBokhylla()...),
	)
}
