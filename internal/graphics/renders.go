package graphics

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	imageGameBG   *ebiten.Image
	imageWindows  = ebiten.NewImage(ScreenWidth, ScreenHeight)
	imageGameOver = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

func textBoxWidth() int {
	x, y = nextWindow
}

func init() {

	//Player Image
	knight, _, err = ebitenutil.NewImageFromFile("D:\\Code-projects\\Ebiten-Game\\assets\\images\\_Idle.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{knight, screenWidth / 2.0, screenHeight / 2.0, 4}
}
