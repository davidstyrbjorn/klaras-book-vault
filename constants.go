package main

import (
	"image/color"
)

const BookFilename = "./books.klara"
const ISBNurl = "http://openlibrary.org/api/books?bibkeys=ISBN:%v&jscmd=details&format=json"

// View definition
const (
	HOME = uint(iota)
	BOOK_SHELF
	ADD_BOOK
)

var SuccessGreen = color.RGBA{20, 220, 0, 255}
var FailedRed = color.RGBA{200, 0, 10, 255}
