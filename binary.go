package main

import (
	"encoding/gob"
	"os"
)

func DumpBookToFile() error {
	file, err := os.Create(BookFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(state.books); err != nil {
		return err
	}

	return nil
}

func LoadBooksFromFile() error {
	file, err := os.Open(BookFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&state.books); err != nil {
		return err
	}

	return nil
}

func insert10DummyBooks() {
	state.books = append(state.books, Book{ISBN: "9780451524935", Title: "1984", Author: "George Orwell", Note: "Så dålig lol", Stars: 1})
	state.books = append(state.books, Book{ISBN: "9780316033803", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Note: "funkar OK", Stars: 2})
	state.books = append(state.books, Book{ISBN: "9780241380279", Title: "Pride and Prejudice", Author: "Jane Austen", Note: "sexig som fan", Stars: 4})
	state.books = append(state.books, Book{ISBN: "9780062401956", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Note: "JÄVLAR VAD BRA", Stars: 5})
	state.books = append(state.books, Book{ISBN: "9780451526342", Title: "Animal Farm", Author: "George Orwell", Note: "OH. MY. GOD", Stars: 5})
	state.books = append(state.books, Book{ISBN: "9780316015844", Title: "The Picture of Dorian Gray", Author: "Oscar Wilde", Note: "ville somna till den", Stars: 2})
	state.books = append(state.books, Book{ISBN: "9780241954658", Title: "Sense and Sensibility", Author: "Jane Austen", Note: "som en star wars film", Stars: 3})
	state.books = append(state.books, Book{ISBN: "9780062801970", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Note: "overrated crap", Stars: 1})
}
