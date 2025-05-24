package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "books.json"
const BACKUP_FILE_NAME = "backup_books.json"

func persistBooks() {
	file, err := os.Create(FILE_NAME)
	if err != nil {
		log.Fatalf("Failed to create books file %v", err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(state.books, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal book list: %v", err)
	}

	err = os.WriteFile(FILE_NAME, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write books to file %v", err)
	}

	fmt.Println("Books saved to", FILE_NAME)
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
