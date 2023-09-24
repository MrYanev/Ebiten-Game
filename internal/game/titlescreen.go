package game

import (
	"bytes"
	"image"
	"log"

	gameImages "Ebiten-Game/assets/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imageBackground *ebiten.Image
	isButtonPressed bool
)

func init() {
	//Background Image
	background, _, err = image.Decode(bytes.NewReader(gameImages.Background.png))
	if err != nil {
		log.Fatal(err)
	}
	imageBackground = ebiten.NewImage(background)
}

type TitleScreen struct {
	cont int
}

func anyButtonPressed(screen *ebiten.Image) bool {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		isButtonPressed = true
	}
	return false
}

func (s *TitleScreen) Update(state *GameState) error
