package internal

import (
	"Ebiten-Game/assets/images"
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	img, _, err := image.Decode(bytes.NewReader(images.Grassy_Mountains.png))
	if err != nil {
		log.Fatal(err)
	}
	startMenuImage = ebiten.NewImageFromImage(img)

	//Any additional setup goes here
}
