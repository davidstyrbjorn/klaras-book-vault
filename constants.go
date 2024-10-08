package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Database definitions
const DBName = "./library.db"
const ISBNurl = "http://openlibrary.org/api/books?bibkeys=ISBN:%v&jscmd=details&format=json"
const (
	UNREAD     = "unread"
	READ       = "read"
	LOANED_OUT = "loaned out"
)

// View definition
const (
	HOME = uint(iota)
	BOOK_SHELF
	ADD_BOOK
)

var BackgroundColor = rl.Beige

const TransitionDuration = 0.35

// Queries
const fuzzySearchBookQuery = `
SELECT id, isbn, title, author, read, loaned, stars, note
FROM books
WHERE isbn LIKE '%' || ?1 || '%'
   OR title LIKE '%' || ?1 || '%'
   OR author LIKE '%' || ?1 || '%'
ORDER BY
	CASE
		WHEN isbn = ?1 THEN 1
		WHEN title = ?1 THEN 2
		WHEN author = ?1 THEN 3
		ELSE 4
	END,
	length(isbn) - length(?1),
	length(title) - length(?1),
	length(author) - length(?1)
LIMIT 20;`
