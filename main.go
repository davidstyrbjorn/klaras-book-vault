package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"database/sql"
	"net/http"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/mattn/go-sqlite3"
)

const DBName = "./library.db"
const ISBNurl = "http://openlibrary.org/api/books?bibkeys=ISBN:%v&jscmd=details&format=json"
const (
	UNREAD     = "unread"
	READ       = "read"
	LOANED_OUT = "loaned out"
)

type Book struct {
	isbn        string
	title       string
	currentPage uint
	status      string
}

type ISBNResponse struct {
	title string
}

func LookupBookFromISBN(isbn string) ISBNResponse {
	// http://openlibrary.org/api/books?bibkeys=ISBN:1931498717&jscmd=details&format=json

	//TODO:  Find out the info we need to fill out a book struct
	resp, err := http.Get(fmt.Sprintf(ISBNurl, isbn))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	stringResult := string(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(stringResult), &data)
	if err != nil {
		log.Fatal(err)
	}

	var isbnResponse ISBNResponse
	for _, v := range data {
		details := v.(map[string]interface{})["details"].(map[string]interface{})
		isbnResponse.title = details["title"].(string)
	}

	return isbnResponse
}

func FuzzySearchBooks(books []Book, search string) []Book {
	return books
}

func main2() {
	fmt.Println("Library Init")

	isbnResponse := LookupBookFromISBN("9780060598242")
	log.Printf(isbnResponse.title)

	isbnResponse = LookupBookFromISBN("9780553213690")
	log.Printf(isbnResponse.title)

	return

	database, err := sql.Open("sqlite3", DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS Books(
		"id" PRIMARY KEY INC INTEGER
		"isbn" TEXT,
		"title" TEXT,
		"page" INTEGER,
		"status" TEXT
	)`
	_, err = database.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	insertBookSQL := "INSERT INTO Books (title, page, status) VALUES (?, ?, ?)"
	insertBookStatement, err := database.Prepare(insertBookSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer insertBookStatement.Close()

	_, err = insertBookStatement.Exec("Edde", 0, UNREAD)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a book 'edde' into the Books table successfully!")
}

const Width = 1280
const Height = 720

func main() {
	rl.InitWindow(Width, Height, "Klaras Bok Valv")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// var button bool

	fontTtf := rl.LoadFont("fonts/Meditative.ttf")
	welcomeMsg := "Klaras Bok Valv"
	msgPosition := rl.Vector2{}
	scale := float32(fontTtf.BaseSize)
	msgPosition.X = float32(Width)/2 - rl.MeasureTextEx(fontTtf, welcomeMsg, scale, 0).X/2
	msgPosition.Y = float32(Height)/2 - rl.MeasureTextEx(fontTtf, welcomeMsg, scale, 0).X/2

	textInput := ""

	var button bool
	buttonWidth := float32(120)
	buttonHeight := float32(60)
	buttonX := float32(Width)/2 - buttonWidth/2
	buttonY := float32(Height)/2 - buttonHeight/2 + 120

	var isbnResponse ISBNResponse
	isbnPosition := rl.Vector2{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Beige)

		key := rl.GetCharPressed()
		for {
			if key <= 0 {
				break
			}

			if key > 0 {
				if key >= 32 && key <= 125 {
					textInput += string(rune(key))
				}
			}

			key = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyBackspace) && len(textInput) != 0 {
			textInput = textInput[0 : len(textInput)-1]
		}

		button = gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Search")
		if button {
			go func() {
				isbnResponse = LookupBookFromISBN(textInput)
				isbnPosition.X = float32(Width)/2 - rl.MeasureTextEx(fontTtf, isbnResponse.title, scale, 0).X/2
				isbnPosition.Y = msgPosition.Y + 130
			}()
		}

		rl.DrawTextEx(fontTtf, "Klaras Bok Valv", msgPosition, scale, 0, rl.Black)
		rl.DrawTextEx(fontTtf, isbnResponse.title, isbnPosition, scale, 0, rl.White)

		inputPosition := rl.Vector2{}
		inputPosition.X = float32(Width)/2 - rl.MeasureTextEx(fontTtf, textInput, scale, 0).X/2
		inputPosition.Y = float32(Height)/2 - rl.MeasureTextEx(fontTtf, textInput, scale, 0).Y/2
		rl.DrawTextEx(fontTtf, textInput, inputPosition, scale, 0, rl.White)

		rl.EndDrawing()
	}
}
