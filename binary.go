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
