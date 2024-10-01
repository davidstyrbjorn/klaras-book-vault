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

// State definition & decleration
type State struct {
	currentView uint
	textInput   string

	// Booksshelf state
	books []Book

	cursiveFont  rl.Font
	textFont     rl.Font
	isbnResponse ISBNResponse

	// Transition state
	transitioning    bool
	t                float32
	transitionShader rl.Shader
	fromTexture      rl.RenderTexture2D
	toTexture        rl.RenderTexture2D
	fromView         uint
	toView           uint
}

var state = State{
	currentView:   HOME,
	books:         []Book{},
	transitioning: false,
	t:             0,
	fromView:      HOME,
	toView:        HOME,
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

func _() {
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

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Klaras Bok Valv")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// var button bool

	state.cursiveFont = rl.LoadFontEx("fonts/Meditative.ttf", 128, nil, 0)
	state.textFont = rl.LoadFontEx("fonts/Virgil.ttf", 512, nil, 0)
	defer rl.UnloadFont(state.cursiveFont)
	defer rl.UnloadFont(state.textFont)

	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, int64(H2))
	rl.SetTextureFilter(state.textFont.Texture, rl.FilterPoint)
	gui.SetFont(state.textFont)

	// Setup the transition data
	state.fromTexture = rl.LoadRenderTexture(WINDOW_WIDTH, WINDOW_HEIGHT)
	state.toTexture = rl.LoadRenderTexture(WINDOW_WIDTH, WINDOW_HEIGHT)
	defer rl.UnloadRenderTexture(state.fromTexture)
	defer rl.UnloadRenderTexture(state.toTexture)

	state.transitionShader = rl.LoadShader("", "shaders/transition.fs") // use default vertex shader
	defer rl.UnloadShader(state.transitionShader)

	tLoc := rl.GetShaderLocation(state.transitionShader, "t")
	toTextureLoc := rl.GetShaderLocation(state.transitionShader, "texture1")

	for !rl.WindowShouldClose() {
		if !state.transitioning {
			rl.BeginDrawing()
			rl.ClearBackground(BackgroundColor)
			drawView(state.currentView)
			rl.EndDrawing()
		} else {
			// Draw the 'to' texture
			rl.BeginTextureMode(state.toTexture)
			rl.ClearBackground(BackgroundColor)
			drawView(state.toView)
			rl.EndTextureMode()

			// Draw the transition, using 'from' and 'to' textures + transition shader
			rl.BeginDrawing()
			rl.ClearBackground(BackgroundColor)
			rl.BeginShaderMode(state.transitionShader)

			rl.SetShaderValue(state.transitionShader, tLoc, []float32{state.t}, rl.ShaderUniformFloat)
			rl.SetShaderValueTexture(state.transitionShader, toTextureLoc, state.toTexture.Texture)
			rl.DrawTexturePro(state.fromTexture.Texture, rl.NewRectangle(0, 0, float32(state.fromTexture.Texture.Width), float32(-state.fromTexture.Texture.Height)), rl.NewRectangle(0, 0, float32(WINDOW_WIDTH), float32(WINDOW_HEIGHT)), rl.NewVector2(0, 0), 0, rl.White)

			rl.EndShaderMode()
			rl.EndDrawing()

			state.t += rl.GetFrameTime() / TransitionDuration
			if state.t >= 1.0 {
				state.t = 1.0
				state.transitioning = false
			}
		}
	}
}

func drawView(view uint) {
	switch view {
	case HOME:
		homeView()
	case BOOK_SHELF:
		bookShelfView()
	case ADD_BOOK:
		addBookView()
	}
}
