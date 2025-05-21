package main

import (
	g "github.com/AllenDang/giu"
)

func homeView() []g.Widget {
	return g.Layout{
		g.Align(g.AlignCenter).To(g.Label("Klaras Bok Valv")),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Column(g.Align(g.AlignCenter).To(
			g.Button("Lägg till bok").OnClick(func() {
				changeView(VIEW_ADD_BOOK)
			}),
			g.Button("Bokhylla").OnClick(func() {
				changeView(VIEW_BOOKSHELF)
			}),
		)),

		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),
		g.Spacing(), g.Spacing(), g.Spacing(),

		g.Align(g.AlignCenter).To(g.Label("Lägg till bok - Skanna ISBN på en bok och klicka på lägg till")),
		g.Align(g.AlignCenter).To(g.Label("Bokhylla - Här visas alla böcker, som du skannat in")),
	}
}
