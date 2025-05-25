package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "books.json"
const BACKUP_FILE_NAME = "backup_books.json"

func persistBooks(fileName string) {
	if fileName == "" {
		fileName = FILE_NAME
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create books file %v", err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(state.books, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal book list: %v", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write books to file %v", err)
	}

	fmt.Println("Books saved to", fileName)
}

func loadBooks() {
	_, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Printf("Failed to open book file")
		return
	}

	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatalf("Failed to read book file: %v", err)
	}

	err = json.Unmarshal(data, &state.books)
	if err != nil {
		log.Fatalf("Failed to unmarshal book list: %v", err)
	}
}

func createBackup() {

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
