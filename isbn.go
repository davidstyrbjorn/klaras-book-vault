package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ISBN related functions, mainly the one that takes a ISBN code and returns a response
func LookupBookFromISBN(isbn string) (ISBNResponse, error) {
	var isbnResponse ISBNResponse

	resp, err := http.Get(fmt.Sprintf(ISBNurl, isbn))
	if err != nil {
		return isbnResponse, fmt.Errorf("Failed to perform GET request")
	}

	if resp.StatusCode != http.StatusOK {
		return isbnResponse, fmt.Errorf("Kunde inte hitta bok")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return isbnResponse, fmt.Errorf("Failed to read response body")
	}

	stringResult := string(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(stringResult), &data)
	if err != nil {
		return isbnResponse, fmt.Errorf("Failed to unmarshal json response1")
	}

	for _, v := range data {
		details := v.(map[string]interface{})["details"].(map[string]interface{})
		isbnResponse.title = details["title"].(string)
		// TODO, parse out the rest of the stuff
	}

	isbnResponse.isbn = isbn

	return isbnResponse, nil
}
