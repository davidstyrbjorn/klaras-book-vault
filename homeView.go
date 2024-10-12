package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func homeView() []g.Widget {
	return []g.Widget{
		g.Align(g.AlignCenter).To(g.Label("Klaras Bok Valv")),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Align(g.AlignCenter).To(g.Row(
			g.Button("Lägg till bok").OnClick(func() {
				switchView(ADD_BOOK)
			}),
			g.Button("Bokhylla").OnClick(func() {
				// switchView(ADD_BOOK)
				state.currentView = BOOK_SHELF
			}),
		)),

		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Align(g.AlignCenter).To(g.Label("Lägg till bok - Skanna ISBN på en bok och klicka på lägg till")),
		g.Align(g.AlignCenter).To(g.Label("Bokhylla - Här visas alla böcker, som du skannat in")),

		g.Button("Lägg till 10 böcker").OnClick(func() {
			insert10Books()
		}),

		g.Button("TEST").OnClick(func() {
			fmt.Println("TEST")
		}),
	}
}
