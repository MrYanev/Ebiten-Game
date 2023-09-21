package internal

import (
	"github.com/hajimehoshi/ebiten"
)

var startMenuImage *ebiten.Image
var isGameStarted bool

func updateStartMenu(screen *ebiten.Image) error {
	//Check for input to start the game
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		isGameStarted = true
	}
	return nil
}

func drawStartMenu(screen *ebiten.Image) {
	//Draw the start menu background
	screen.DrawImage(startMenuImage, nil)

	//To draw text and other elements here
}

func switchToGame() {
	isGameStarted = true
}

func initializeStartMenu() {
	//Load up the start menu background

	//Any additional setup goes here
}
