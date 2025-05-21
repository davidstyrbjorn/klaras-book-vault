package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

		authors := []string{}
		if details["authors"] != nil {
			for _, author := range details["authors"].([]interface{}) {
				authorString := author.(map[string]interface{})["name"].(string)
				authors = append(authors, authorString)
			}
		}
		if len(authors) == 1 {
			isbnResponse.Author = authors[0]
		} else if len(authors) > 1 {
			isbnResponse.Author = strings.Join(authors, ", ")
		}
	}

	isbnResponse.isbn = isbn

	return isbnResponse, nil
}
