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
		return isbnResponse, fmt.Errorf("failed to perform GET request")
	}

	if resp.StatusCode != http.StatusOK {
		return isbnResponse, fmt.Errorf("kunde inte hitta bok")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return isbnResponse, fmt.Errorf("failed to read response body")
	}

	stringResult := string(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(stringResult), &data)
	if err != nil {
		return isbnResponse, fmt.Errorf("failed to unmarshal json response, kontakta utvecklare och get isbn numret du försökte använda")
	}

	// If we got a zero map back, we have a bad ISBN number
	if len(data) == 0 {
		return isbnResponse, fmt.Errorf("isbn '%v' gav inget resultat tyvärr", isbn)
	}

	for _, v := range data {
		details := v.(map[string]interface{})["details"].(map[string]interface{})
		isbnResponse.title = details["title"].(string)
		// TODO, parse out the rest of the stuff
	}

	isbnResponse.isbn = isbn

	return isbnResponse, nil
}
