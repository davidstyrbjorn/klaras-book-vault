package main

import (
	"fmt"
	"log"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func homeView() {
	key := rl.GetCharPressed()
	for {
		if key <= 0 {
			break
		}

		if key > 0 {
			if key >= 32 && key <= 125 {
				state.textInput += string(rune(key))
			}
		}

		key = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) && len(state.textInput) != 0 {
		state.textInput = state.textInput[0 : len(state.textInput)-1]
	}

	welcomeMsg := "Klaras Bok Valv"
	msgPosition := rl.Vector2{}
	msgPosition.X = float32(WINDOW_WIDTH)/2 - rl.MeasureTextEx(state.cursiveFont, welcomeMsg, H1, 0).X/2
	msgPosition.Y = 20

	isbnPosition := rl.Vector2{}

	buttonWidth := float32(120)
	buttonHeight := float32(60)
	buttonX := float32(WINDOW_WIDTH)/2 - buttonWidth/2
	buttonY := float32(WINDOW_HEIGHT)/2 - buttonHeight/2 + 120
	button := gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Search")
	if button {
		go func() {
			state.isbnResponse = LookupBookFromISBN(state.textInput)
			isbnPosition.X = float32(WINDOW_WIDTH)/2 - rl.MeasureTextEx(state.cursiveFont, state.isbnResponse.title, H1, 0).X/2
			isbnPosition.Y = msgPosition.Y + 130
		}()
	}

	buttonX = 20
	buttonY = 20
	button = gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Bokhylla")
	if button {
		changeView(BOOK_SHELF)
	}

	// Some debug butttons
	// Clear book table
	// Insert 10 books
	buttonY = 100
	buttonX = 20
	button = gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Clear")
	if button {
		clearBookTable()
	}

	buttonY = 180
	buttonX = 20
	button = gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Fill")
	if button {
		insert10Books()
	}

	buttonY = 260
	buttonX = 20
	button = gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Read")
	if button {
		go func() {
			books, err := readAllBooks()
			if err != nil {
				log.Println(err)
			}
			if len(books) > 0 {
				fmt.Println(books)
			} else {
				fmt.Println("No books found")
			}
		}()
	}

	rl.DrawTextEx(state.cursiveFont, "Klaras Bok Valv", msgPosition, H1, 0, rl.Black)

	rl.DrawTextEx(state.textFont, state.isbnResponse.title, isbnPosition, H2, 0, rl.Black)

	inputPosition := rl.Vector2{}
	inputPosition.X = float32(WINDOW_WIDTH)/2 - rl.MeasureTextEx(state.textFont, state.textInput, H1, 0).X/2
	inputPosition.Y = float32(WINDOW_HEIGHT)/2 - rl.MeasureTextEx(state.textFont, state.textInput, H1, 0).Y/2
	rl.DrawTextEx(state.textFont, state.textInput, inputPosition, H1, 0, rl.Black)

	// Add a black line underneath the text input
	lineStartX := inputPosition.X
	lineEndX := inputPosition.X + rl.MeasureTextEx(state.textFont, state.textInput, H1, 0).X
	lineY := inputPosition.Y + rl.MeasureTextEx(state.textFont, state.textInput, H1, 0).Y + 5
	rl.DrawLine(int32(lineStartX), int32(lineY), int32(lineEndX), int32(lineY), rl.Black)

}

func updateBookShelfView() {
	// Search for books
	if len(state.textInput) > 0 {
		go func() {
			books, err := fuzzySearchBook(state.textInput)
			if err != nil {
				log.Println(err)
			}
			state.books = books
		}()
	} else {
		go func() {
			books, err := readAllBooks()
			if err != nil {
				log.Println(err)
			}
			state.books = books
		}()
	}
}

func changeView(view uint) {
	if view == state.currentView {
		return
	}

	if view == BOOK_SHELF {
		state.textInput = ""
		updateBookShelfView()
	} else if view == HOME {
		state.textInput = ""
	}

	state.currentView = view
}

func bookShelfView() {
	titleMsg := "Bokhyllan"
	msgPosition := rl.Vector2{}
	msgPosition.X = float32(WINDOW_WIDTH)/2 - rl.MeasureTextEx(state.cursiveFont, titleMsg, H1, 0).X/2
	msgPosition.Y = 20
	rl.DrawTextEx(state.cursiveFont, titleMsg, msgPosition, H1, 0, rl.Black)

	buttonWidth := float32(120)
	buttonHeight := float32(60)
	buttonX := float32(20)
	buttonY := float32(20)
	button := gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "Back")
	if button {
		changeView(HOME)
	}

	centerX := int32(WINDOW_WIDTH) / 2
	rl.DrawLine(centerX-200, 200, centerX+200, 200, rl.Black)

	key := rl.GetCharPressed()
	for {
		if key <= 0 {
			break
		}

		if key > 0 {
			if key >= 32 && key <= 125 {
				state.textInput += string(rune(key))
				updateBookShelfView()
			}
		}

		key = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) && len(state.textInput) != 0 {
		state.textInput = state.textInput[0 : len(state.textInput)-1]
		updateBookShelfView()
	}

	if state.textInput == "" {
		rl.DrawTextEx(state.textFont, "Beskriv boken du letar efter", rl.Vector2{X: float32(centerX) - 200, Y: 175}, H2, 0, rl.Gray)
	} else {
		rl.DrawTextEx(state.textFont, state.textInput, rl.Vector2{X: float32(centerX) - 200, Y: 175}, H2, 0, rl.Black)
	}

	for idx, book := range state.books {
		rl.DrawTextEx(state.textFont, book.title, rl.Vector2{X: float32(centerX) - 200, Y: 210 + float32(idx*20)}, H2, 0, rl.Black)
	}
}

func addBookView() {

}
